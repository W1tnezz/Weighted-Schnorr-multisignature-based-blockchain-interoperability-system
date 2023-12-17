const fs = require("fs");
const createCsvWriter = require("csv-writer").createObjectCsvWriter;
const OracleContract = artifacts.require("OracleContract");

module.exports = async function () {

  let oracleContract = await OracleContract.deployed();
  let tx = "0xf4bb9a6843194a5f9d10c70202b445733ad8feb5b5f9bce2c699e4cd0abb4bd2";
  let size = 10;   // 总量阈值
  let minRank = 3;  // 个人阈值
  let fee = await oracleContract.TOTAL_FEE();

  await oracleContract.validateTransaction(tx, size, minRank,{
    value: fee,
  });

};
