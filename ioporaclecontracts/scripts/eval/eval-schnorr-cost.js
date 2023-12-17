const fs = require("fs");
const createCsvWriter = require("csv-writer").createObjectCsvWriter;
const OracleContract = artifacts.require("OracleContract");

module.exports = async function () {

  let oracleContract = await OracleContract.deployed();
  let tx = "0x034897cba1b98c3c5d4bb491bd5cef7fac8ff87bf0b97ce563894dbd61bcd008";
  let size = 10;   // 总量阈值
  let minRank = 3;  // 个人阈值
  let fee = await oracleContract.totalFee(size);

  await oracleContract.validateTransaction(tx, size, minRank,{
    value: fee,
  });

};
