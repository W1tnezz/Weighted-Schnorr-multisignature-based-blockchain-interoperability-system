const OracleContract = artifacts.require("OracleContract");

module.exports = function (deployer) {
    return deployer.deploy(
        OracleContract,
    );
    
};
