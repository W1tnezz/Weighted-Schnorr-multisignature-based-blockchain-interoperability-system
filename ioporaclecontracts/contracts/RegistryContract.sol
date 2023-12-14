// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract RegistryContract {
    struct OracleNode {
        address addr;  // 链上地址
        string ipAddr; // 节点IP地址
        bytes[] pubKey;  // schnorr公钥；
        uint256 stake; // 质押
        uint256 rank;  // 可信等级，即公钥数量
        uint256 index;
    }

    struct AggregatorNode {
        address addr;  // 链上地址
        string ipAddr; // 节点IP地址
    }

    uint256 public constant MIN_STAKE = 1 ether;
    bool private hasAggregator;

    mapping(address => OracleNode) private oracleNodes;
    mapping(address => AggregatorNode) private AggregatorNodes;
    address[] private oracleNodeIndices;

    event RegisterOracleNode(address indexed sender);

    function registerOracleNode(string calldata _ipAddr, bytes[] calldata _pubKey, uint256 rank)
        external
        payable
    {
        require(!oracleNodeIsRegistered(msg.sender), "already registered");
        require(msg.value >= rank * MIN_STAKE, "low stake");
        require(_pubKey.length == rank, "key number error");

        OracleNode storage iopNode = oracleNodes[msg.sender];
        iopNode.addr = msg.sender;
        iopNode.ipAddr = _ipAddr;
        iopNode.pubKey = _pubKey;
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

    function registerAggregatorNode(string calldata ipAddr)
        external
        payable
    {
        require(!hasAggregator, "already registered");
        hasAggregator = true;
        AggregatorNode storage aggregator = AggregatorNodes[msg.sender];
        aggregator.addr = msg.sender;
        aggregator.ipAddr = ipAddr;
        emit RegisterOracleNode(msg.sender);
    }
}
