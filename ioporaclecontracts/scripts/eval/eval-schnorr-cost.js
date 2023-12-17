const fs = require("fs");
const createCsvWriter = require("csv-writer").createObjectCsvWriter;
const OracleContract = artifacts.require("OracleContract");

module.exports = async function () {

  let oracleContract = await OracleContract.deployed();
  let tx = "0x5df34dd9c92065367162bbc91b91da617a1bda0c260bddc77c1fdf4c4b98f2ed";
  let size = 10;   // 总量阈值
  let minRank = 3;  // 个人阈值
  let fee = await oracleContract.totalFee(size);

  await oracleContract.validateTransaction(tx, size, minRank,{
    value: fee,
  });

};
