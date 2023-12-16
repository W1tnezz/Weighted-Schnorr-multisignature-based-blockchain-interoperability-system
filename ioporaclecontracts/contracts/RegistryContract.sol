// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract RegistryContract {
    struct OracleNode {
        address addr;  // 链上地址
        string ipAddr; // 节点IP地址
        bytes[] pubKeys;  // schnorr公钥；
        uint256 stake; // 质押
        uint256 rank;  // 可信等级，即公钥数量
        uint256 index;
    }

    uint256 public constant MIN_STAKE = 1 ether;
    bool private hasAggregator;
    address private aggregatorAddr;
    string private aggregatorIP;

    mapping(address => OracleNode) private oracleNodes;

    address[] private oracleNodeIndices;

    event RegisterOracleNode(address indexed sender);

    function registerOracleNode(string calldata _ipAddr, bytes[] calldata _pubKey, uint256 rank)
        external
        payable
    {
        require(!oracleNodeIsRegistered(msg.sender), "already registered");
        require(msg.value >= rank * MIN_STAKE, "low stake");
        require(_pubKey.length == rank, "key number error");
        if(!hasAggregator){
            hasAggregator = true;
            aggregatorAddr = msg.sender;
            aggregatorIP = _ipAddr;
        }
        OracleNode storage iopNode = oracleNodes[msg.sender];
        iopNode.addr = msg.sender;
        iopNode.ipAddr = _ipAddr;
        iopNode.pubKeys = _pubKey;
        iopNode.stake = msg.value;
        iopNode.rank = rank;
        iopNode.index = oracleNodeIndices.length;
        oracleNodeIndices.push(iopNode.addr);

        emit RegisterOracleNode(msg.sender);
    }

    function oracleNodeIsRegistered(address _addr) public view returns (bool) {
        if (oracleNodeIndices.length == 0) return false;
        return (oracleNodeIndices[oracleNodes[_addr].index] == _addr);
    }

    function findOracleNodeByAddress(address _addr)
        public
        view
        returns (OracleNode memory)
    {
        require(oracleNodeIsRegistered(_addr), "not found");
        return oracleNodes[_addr];
    }

    function findOracleNodeByIndex(uint256 _index)
        public
        view
        returns (OracleNode memory)
    {
        require(_index >= 0 && _index < oracleNodeIndices.length, "not found");
        return oracleNodes[oracleNodeIndices[_index]];
    }

    function unregister(address unregisterAddr) 
        public
    {
        require(msg.sender == unregisterAddr, "Only allow unregister yourself!");
        require(oracleNodeIsRegistered(unregisterAddr), "Haven't registered!");
        payable(unregisterAddr).transfer(oracleNodes[unregisterAddr].stake); // 退回押金
        delete oracleNodeIndices[oracleNodes[unregisterAddr].index]; // 删除数组地址
        delete oracleNodes[unregisterAddr]; // 删除map键值对
    }

    function deleteNode(address addr)
        external
    {
        delete oracleNodeIndices[oracleNodes[addr].index]; // 删除数组地址
        delete oracleNodes[addr]; // 删除map键值对
    }

    function getNodeRank(address addr) public view returns (uint256){
        return oracleNodes[addr].rank;
    }

    function getNodePublicKeys(address addr) public view returns (bytes[] memory){
        return oracleNodes[addr].pubKeys;
    }

    function isAggregator(address _addr) public view returns (bool) {
        return _addr == aggregatorAddr;
    }

    function getAggregator() public view returns (address) {
        require(hasAggregator, "no aggregator");
        return aggregatorAddr;
    }

    function getAggregatorIP() public view returns (string memory) {
        require(hasAggregator, "no aggregator");
        return aggregatorIP;
    }
}
