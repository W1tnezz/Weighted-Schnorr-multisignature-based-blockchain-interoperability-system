// SPDX-License-Identifier: MIT
pragma solidity  ^0.8.0;

import "./RegistryContract.sol";
import "./crypto/Schnorr.sol";
import "./crypto/BN256G1.sol";

contract OracleContract {

    // 当前是否有请求验证的标志
    bool private isValidateTime = false;

    // 请求验证需要的钱
    uint256 public constant BASE_FEE = 0.1 ether;
    uint256 public constant AGGREGATE_FEE = 0.5 ether;

    uint256 private currentRank;
    uint256 private currentSize;
    uint256 public constant PUBKEY_LENGTH = 33;

    // 保存验证结果的映射；
    mapping(bytes32 => bool) private blockValidationResults;
    mapping(bytes32 => bool) private txValidationResults;

    // 验证类型的枚举：未知，区块存在验证，交易存在验证;
    enum ValidationType { UNKNOWN, BLOCK, TRANSACTION }

    // indexed属性是为了方便在日志结构中查找，这个是一个事件，会存储到交易的日志中，就是类似于挖矿上链
    event ValidationRequest(ValidationType typ, address indexed from, bytes32 hash, uint256 size, uint256 minRank);

    RegistryContract private registryContract;
    constructor(address registry) {
        registryContract = RegistryContract(registry);
    }

    modifier minFee(uint size) {
        require(msg.value >= BASE_FEE * size + AGGREGATE_FEE, "too few fee amount");
        _;
    }

    function validateBlock(bytes32 _message, uint256 size, uint256 minRank) external payable minFee(size) {
        require(!isValidateTime, "Another validate is in progress!");
        isValidateTime = true;
        emit ValidationRequest(ValidationType.BLOCK, msg.sender, _message, size, minRank);
    }

    function validateTransaction(bytes32 _message, uint256 size, uint256 minRank) external payable minFee(size) {
        require(!isValidateTime, "Another validate is in progress!");
        isValidateTime = true;
        emit ValidationRequest(ValidationType.TRANSACTION, msg.sender, _message, size, minRank);
    }


//-----------------------------------------------------------------------------------------------------------------------------------------------


    function submitBlockValidationResult(bool _result, bytes32 message, uint256 signature, uint256 rx , uint256 ry, uint256 _hash, address[] memory validators) external {
        require(isValidateTime, "Not validate time!");
        submitValidationResult(ValidationType.BLOCK, _result, message, signature, rx, ry, _hash, validators);
        isValidateTime = false;
    }

    function submitTransactionValidationResult(bool _result, bytes32 message, uint256 signature, uint256 rx , uint256 ry, uint256 _hash, address[] memory validators) external {
        require(isValidateTime, "Not validate time!");
        submitValidationResult(ValidationType.TRANSACTION, _result, message, signature, rx, ry, _hash, validators);
        isValidateTime = false;
    }

    function submitValidationResult(
        ValidationType _typ,
        bool _result,
        bytes32 message,
        uint256 signature, uint256 rx , uint256 ry, uint256 _hash, 
        address[] memory validators
    ) private {

        require(_typ != ValidationType.UNKNOWN, "unknown validation type");
        require(registryContract.getAggregator() == msg.sender, "not the aggregator");  //判断当前合约的调用者是不是聚合器
        
        
        uint256 totalRank = 0;
        bytes[][] memory allPubKeys = new bytes[][](validators.length);
        for(uint32 i = 0 ; i < validators.length ; i++){
            // 验证单个节点的信誉值；
            uint256 rank = registryContract.getNodeRank(validators[i]);
            require(rank >= currentRank, "low singal rank");
            totalRank += rank;
            bytes[] memory pubKeys = registryContract.getNodePublicKeys(validators[i]);
            allPubKeys[i] = pubKeys;
        }
        require(totalRank >= currentRank, "low total rank");
        
        // TODO:公钥重新聚合
        bytes memory S = new bytes((totalRank + 1) * PUBKEY_LENGTH);
        uint256 index = PUBKEY_LENGTH;
        for(uint32 i = 0 ; i < allPubKeys.length ; i++){
            for(uint32 j = 0; j < allPubKeys[i].length; j++){
                for(uint32 k = 0; k < allPubKeys[i][j].length; k++){
                    S[index] = allPubKeys[i][j][k];
                    index++;
                }
            }
        }

        uint256 pubKeyX = 0;
        uint256 pubKeyY = 0;

        for(uint32 i = 0 ; i < allPubKeys.length ; i++){
            for(uint32 j = 0; j < allPubKeys[i].length; j++){
                uint256 tempX;
                uint256 tempY;
                (tempX, tempY) = BN256G1.fromCompressed(allPubKeys[i][j]);
                for(uint32 k = 0; k < PUBKEY_LENGTH; k++){
                    S[k] = allPubKeys[i][j][k];
                }
                uint256 res = bytesToUint256(sha256(S));
                (tempX, tempY) = BN256G1.mulPoint([tempX, tempY, res]);

                (pubKeyX, pubKeyY) = BN256G1.addPoint([tempX, tempY, pubKeyX, pubKeyY]);
            }
        }


        /*Schnorr签名的验证*/
        require(Schnorr.verify(signature, pubKeyX, pubKeyY, rx, ry, _hash), "sig: address doesn't match");

        if (_typ == ValidationType.BLOCK) {
            blockValidationResults[message] = _result;
        } else if (_typ == ValidationType.TRANSACTION) {
            txValidationResults[message] = _result;
        }

        // 给当前合约的调用者（聚合器）转账 
        payable(msg.sender).transfer(AGGREGATE_FEE);     //此处完成给聚合器的报酬转账
        // 给所有的参与验证的验证器节点转账

        for(uint32 i = 0 ; i < validators.length ; i++){
            if(address(this).balance >= BASE_FEE * registryContract.getNodeRank(validators[i])){
                payable(validators[i]).transfer(BASE_FEE * registryContract.getNodeRank(validators[i])); 
            } else{
                payable(validators[i]).transfer(address(this).balance); 
            }
        }
    }

    function bytesToUint256(bytes32 b) public pure returns (uint256){
        uint256 number;
        for(uint i= 0; i < b.length; i++){
            number = number + uint8(b[i])*(2**(8*(b.length-(i+1))));
        }
        return  number;
    }
}
