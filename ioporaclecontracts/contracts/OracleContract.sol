// SPDX-License-Identifier: MIT
pragma solidity  >0.8.0;

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

    function totalFee(uint256 size) public pure returns (uint256){
        return size * BASE_FEE + AGGREGATE_FEE;
    }

    function validateBlock(bytes32 _message, uint256 size, uint256 minRank) external payable minFee(size) {
        require(!isValidateTime, "Another validate is in progress!");
        isValidateTime = true;
        emit ValidationRequest(ValidationType.BLOCK, msg.sender, _message, size, minRank);
    }

    function validateTransaction(bytes32 _message, uint256 size, uint256 minRank) external payable minFee(size) {
        // require(!isValidateTime, "Another validate is in progress!");
        isValidateTime = true;
        emit ValidationRequest(ValidationType.TRANSACTION, msg.sender, _message, size, minRank);
    }


//-----------------------------------------------------------------------------------------------------------------------------------------------


    function submitBlockValidationResult(bool _result, bytes32 message, uint256 signature,uint256 pubKeyX, uint256 pubKeyY, uint256 rx , uint256 ry, uint256 _hash, address[] memory validators, uint256[2][] memory pubKeyArray) external {
        require(isValidateTime, "Not validate time!");
        // submitValidationResult(ValidationType.BLOCK, _result, message, signature, rx, ry, _hash, validators);
        submitValidationResult2(ValidationType.BLOCK, _result, message, signature, pubKeyX, pubKeyY, rx, ry, _hash, validators, pubKeyArray);
        isValidateTime = false;
    }

    function submitTransactionValidationResult(bool _result, bytes32 message, uint256 signature,uint256 pubKeyX, uint256 pubKeyY, uint256 rx , uint256 ry, uint256 _hash, address[] memory validators, uint256[2][] memory pubKeyArray) external {
        require(isValidateTime, "Not validate time!");
        // submitValidationResult(ValidationType.TRANSACTION, _result, message, signature, rx, ry, _hash, validators);
        submitValidationResult2(ValidationType.BLOCK, _result, message, signature, pubKeyX, pubKeyY, rx, ry, _hash, validators, pubKeyArray);
        isValidateTime = false;
    }

    // 点加进行Hash恢复公钥，gas: 440616
    function submitValidationResult0(
        ValidationType _typ,
        bool _result,
        bytes32 message,
        uint256 signature, uint256 rx , uint256 ry, uint256 _hash, 
        address[] memory validators
    ) private {
        require(_typ != ValidationType.UNKNOWN, "unknown validation type");
        require(registryContract.getAggregator() == msg.sender, "not the aggregator");  //判断当前合约的调用者是不是聚合器
    
        uint256 totalRank = 0;
        for(uint16 i = 0 ; i < validators.length ; i++){
            // 验证单个节点的信誉值；
            uint256 rank = registryContract.getNodeRank(validators[i]);
            require(rank >= currentRank, "low singal rank");
            totalRank += rank;
        }
        require(totalRank >= currentRank, "low total rank");
            
        //公钥重新聚合
        uint256 Sx = 0;
        uint256 Sy = 0;
        for(uint16 i = 0; i < validators.length; i++){
            uint256[2][] memory keys = registryContract.getNodePublicKeys(validators[i]);
            for(uint16 j = 0 ; j < keys.length; j++){
                (Sx, Sy) = BN256G1.addPoint([Sx, Sy, keys[j][0], keys[j][1]]);
            }
        }
        
        uint256 pubKeyX = 0;
        uint256 pubKeyY = 0;

        for(uint16 i = 0; i < validators.length; i++){
            uint256[2][] memory keys = registryContract.getNodePublicKeys(validators[i]);
            for(uint16 j = 0 ; j < keys.length ; j++){
                uint256 tempX = keys[j][0];
                uint256 tempY = keys[j][1];

                uint256 pkX;
                uint256 pkY;
                (pkX, pkY) = BN256G1.addPoint([tempX, tempY, Sx, Sy]);
                
                uint256 res = uint256(sha256(abi.encode(pkX, pkY)));
                (tempX, tempY) = BN256G1.mulPoint([tempX, tempY, res]);
                (pubKeyX, pubKeyY) = BN256G1.addPoint([tempX, tempY, pubKeyX, pubKeyY]);
            }
        }

        /*Schnorr签名的验证*/
        require(Schnorr.verify(signature, pubKeyX, pubKeyY, rx, ry, _hash), "sig: address doesn't match");
        // require(Schnorr.verify(signature, keyX, keyY, rx, ry, _hash), "sig: address doesn't match");

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

    // 字节序列拼接构造公钥，gas: 574237
    function submitValidationResult1(
        ValidationType _typ,
        bool _result,
        bytes32 message,
        uint256 signature, uint256 rx , uint256 ry, uint256 _hash, 
        address[] memory validators
    ) private {
        require(_typ != ValidationType.UNKNOWN, "unknown validation type");
        require(registryContract.getAggregator() == msg.sender, "not the aggregator");  //判断当前合约的调用者是不是聚合器
    
        uint256 totalRank = 0;
        for(uint16 i = 0 ; i < validators.length ; i++){
            // 验证单个节点的信誉值；
            uint256 rank = registryContract.getNodeRank(validators[i]);
            require(rank >= currentRank, "low singal rank");
            totalRank += rank;
        }
        require(totalRank >= currentRank, "low total rank");
            
        //公钥重新聚合
        uint256[2][] memory keys1 = registryContract.getNodePublicKeys(validators[0]);
        bytes memory S = abi.encodePacked(bytes32(keys1[0][0]), bytes32(keys1[0][1]));

        for(uint16 i = 0; i < validators.length; i++){
            uint256[2][] memory keys = registryContract.getNodePublicKeys(validators[i]);
            for(uint16 j = 0 ; j < keys.length; j++){
                if(i == 0 && j == 0){
                    continue;
                } else{
                    S = abi.encodePacked(S, bytes32(keys[j][0]), bytes32(keys[j][1]));
                }
            }
        }
        
        uint256 pubKeyX = 0;
        uint256 pubKeyY = 0;

        for(uint16 i = 0; i < validators.length; i++){
            uint256[2][] memory keys = registryContract.getNodePublicKeys(validators[i]);
            for(uint16 j = 0 ; j < keys.length ; j++){
                uint256 tempX = keys[j][0];
                uint256 tempY = keys[j][1];
                
                uint256 res = uint256(sha256(abi.encodePacked(bytes32(keys[j][0]), bytes32(keys[j][1]), S)));
                (tempX, tempY) = BN256G1.mulPoint([tempX, tempY, res]);
                (pubKeyX, pubKeyY) = BN256G1.addPoint([tempX, tempY, pubKeyX, pubKeyY]);
            }
        }

        /*Schnorr签名的验证*/
        require(Schnorr.verify(signature, pubKeyX, pubKeyY, rx, ry, _hash), "sig: address doesn't match");
        // require(Schnorr.verify(signature, keyX, keyY, rx, ry, _hash), "sig: address doesn't match");

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


    // 直接上传多重公钥，以及签名者公钥集合 gas: 273551
    function submitValidationResult2(
        ValidationType _typ,
        bool _result,
        bytes32 message,
        uint256 signature,uint256 pubKeyX, uint256 pubKeyY, uint256 rx , uint256 ry, uint256 _hash, 
        address[] memory validators, uint256[2][] memory pubKeyArray
    ) private {
        require(_typ != ValidationType.UNKNOWN, "unknown validation type");
        require(registryContract.getAggregator() == msg.sender, "not the aggregator");  //判断当前合约的调用者是不是聚合器
        require(pubKeyArray.length >= currentRank, "low total rank");

        uint32 index = 0;
        for(uint16 i = 0 ; i < validators.length ; i++){
            // 验证单个节点的信誉值；
            uint256[2][] memory keys = registryContract.getNodePublicKeys(validators[i]);
            require(keys.length >= currentRank, "low singal rank");
            
            for(uint16 j = 0; j < keys.length; j++){
                require(pubKeyArray[index][0] == keys[j][0] && pubKeyArray[index][1] == keys[j][1], "pubkey not equal!");
                index++;
            }
        }

        /*Schnorr签名的验证*/
        require(Schnorr.verify(signature, pubKeyX, pubKeyY, rx, ry, _hash), "sig: address doesn't match");
        // require(Schnorr.verify(signature, keyX, keyY, rx, ry, _hash), "sig: address doesn't match");

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
}
