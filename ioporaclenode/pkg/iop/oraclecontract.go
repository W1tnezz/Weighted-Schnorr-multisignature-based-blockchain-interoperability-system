// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iop

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// OracleContractMetaData contains all meta data concerning the OracleContract contract.
var OracleContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumOracleContract.ValidationType\",\"name\":\"typ\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minRank\",\"type\":\"uint256\"}],\"name\":\"ValidationRequest\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AGGREGATE_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BASE_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PUBKEY_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"b\",\"type\":\"bytes32\"}],\"name\":\"bytesToUint256\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_result\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"signature\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_hash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubKeyX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubKeyY\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"submitBlockValidationResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_result\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"signature\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_hash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubKeyX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubKeyY\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"submitTransactionValidationResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"totalFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_message\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minRank\",\"type\":\"uint256\"}],\"name\":\"validateBlock\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_message\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minRank\",\"type\":\"uint256\"}],\"name\":\"validateTransaction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// OracleContractABI is the input ABI used to generate the binding from.
// Deprecated: Use OracleContractMetaData.ABI instead.
var OracleContractABI = OracleContractMetaData.ABI

// OracleContract is an auto generated Go binding around an Ethereum contract.
type OracleContract struct {
	OracleContractCaller     // Read-only binding to the contract
	OracleContractTransactor // Write-only binding to the contract
	OracleContractFilterer   // Log filterer for contract events
}

// OracleContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleContractSession struct {
	Contract     *OracleContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleContractCallerSession struct {
	Contract *OracleContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// OracleContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleContractTransactorSession struct {
	Contract     *OracleContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// OracleContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleContractRaw struct {
	Contract *OracleContract // Generic contract binding to access the raw methods on
}

// OracleContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleContractCallerRaw struct {
	Contract *OracleContractCaller // Generic read-only contract binding to access the raw methods on
}

// OracleContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleContractTransactorRaw struct {
	Contract *OracleContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracleContract creates a new instance of OracleContract, bound to a specific deployed contract.
func NewOracleContract(address common.Address, backend bind.ContractBackend) (*OracleContract, error) {
	contract, err := bindOracleContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OracleContract{OracleContractCaller: OracleContractCaller{contract: contract}, OracleContractTransactor: OracleContractTransactor{contract: contract}, OracleContractFilterer: OracleContractFilterer{contract: contract}}, nil
}

// NewOracleContractCaller creates a new read-only instance of OracleContract, bound to a specific deployed contract.
func NewOracleContractCaller(address common.Address, caller bind.ContractCaller) (*OracleContractCaller, error) {
	contract, err := bindOracleContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleContractCaller{contract: contract}, nil
}

// NewOracleContractTransactor creates a new write-only instance of OracleContract, bound to a specific deployed contract.
func NewOracleContractTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleContractTransactor, error) {
	contract, err := bindOracleContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleContractTransactor{contract: contract}, nil
}

// NewOracleContractFilterer creates a new log filterer instance of OracleContract, bound to a specific deployed contract.
func NewOracleContractFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleContractFilterer, error) {
	contract, err := bindOracleContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleContractFilterer{contract: contract}, nil
}

// bindOracleContract binds a generic wrapper to an already deployed contract.
func bindOracleContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OracleContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleContract *OracleContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OracleContract.Contract.OracleContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleContract *OracleContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleContract.Contract.OracleContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleContract *OracleContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleContract.Contract.OracleContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleContract *OracleContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OracleContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleContract *OracleContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleContract *OracleContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleContract.Contract.contract.Transact(opts, method, params...)
}

// AGGREGATEFEE is a free data retrieval call binding the contract method 0x1471866c.
//
// Solidity: function AGGREGATE_FEE() view returns(uint256)
func (_OracleContract *OracleContractCaller) AGGREGATEFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "AGGREGATE_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AGGREGATEFEE is a free data retrieval call binding the contract method 0x1471866c.
//
// Solidity: function AGGREGATE_FEE() view returns(uint256)
func (_OracleContract *OracleContractSession) AGGREGATEFEE() (*big.Int, error) {
	return _OracleContract.Contract.AGGREGATEFEE(&_OracleContract.CallOpts)
}

// AGGREGATEFEE is a free data retrieval call binding the contract method 0x1471866c.
//
// Solidity: function AGGREGATE_FEE() view returns(uint256)
func (_OracleContract *OracleContractCallerSession) AGGREGATEFEE() (*big.Int, error) {
	return _OracleContract.Contract.AGGREGATEFEE(&_OracleContract.CallOpts)
}

// BASEFEE is a free data retrieval call binding the contract method 0x3d18651e.
//
// Solidity: function BASE_FEE() view returns(uint256)
func (_OracleContract *OracleContractCaller) BASEFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "BASE_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASEFEE is a free data retrieval call binding the contract method 0x3d18651e.
//
// Solidity: function BASE_FEE() view returns(uint256)
func (_OracleContract *OracleContractSession) BASEFEE() (*big.Int, error) {
	return _OracleContract.Contract.BASEFEE(&_OracleContract.CallOpts)
}

// BASEFEE is a free data retrieval call binding the contract method 0x3d18651e.
//
// Solidity: function BASE_FEE() view returns(uint256)
func (_OracleContract *OracleContractCallerSession) BASEFEE() (*big.Int, error) {
	return _OracleContract.Contract.BASEFEE(&_OracleContract.CallOpts)
}

// PUBKEYLENGTH is a free data retrieval call binding the contract method 0xa4d55d1d.
//
// Solidity: function PUBKEY_LENGTH() view returns(uint256)
func (_OracleContract *OracleContractCaller) PUBKEYLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "PUBKEY_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PUBKEYLENGTH is a free data retrieval call binding the contract method 0xa4d55d1d.
//
// Solidity: function PUBKEY_LENGTH() view returns(uint256)
func (_OracleContract *OracleContractSession) PUBKEYLENGTH() (*big.Int, error) {
	return _OracleContract.Contract.PUBKEYLENGTH(&_OracleContract.CallOpts)
}

// PUBKEYLENGTH is a free data retrieval call binding the contract method 0xa4d55d1d.
//
// Solidity: function PUBKEY_LENGTH() view returns(uint256)
func (_OracleContract *OracleContractCallerSession) PUBKEYLENGTH() (*big.Int, error) {
	return _OracleContract.Contract.PUBKEYLENGTH(&_OracleContract.CallOpts)
}

// BytesToUint256 is a free data retrieval call binding the contract method 0x53d62ff7.
//
// Solidity: function bytesToUint256(bytes32 b) pure returns(uint256)
func (_OracleContract *OracleContractCaller) BytesToUint256(opts *bind.CallOpts, b [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "bytesToUint256", b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BytesToUint256 is a free data retrieval call binding the contract method 0x53d62ff7.
//
// Solidity: function bytesToUint256(bytes32 b) pure returns(uint256)
func (_OracleContract *OracleContractSession) BytesToUint256(b [32]byte) (*big.Int, error) {
	return _OracleContract.Contract.BytesToUint256(&_OracleContract.CallOpts, b)
}

// BytesToUint256 is a free data retrieval call binding the contract method 0x53d62ff7.
//
// Solidity: function bytesToUint256(bytes32 b) pure returns(uint256)
func (_OracleContract *OracleContractCallerSession) BytesToUint256(b [32]byte) (*big.Int, error) {
	return _OracleContract.Contract.BytesToUint256(&_OracleContract.CallOpts, b)
}

// TotalFee is a free data retrieval call binding the contract method 0xcea74edf.
//
// Solidity: function totalFee(uint256 size) pure returns(uint256)
func (_OracleContract *OracleContractCaller) TotalFee(opts *bind.CallOpts, size *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OracleContract.contract.Call(opts, &out, "totalFee", size)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalFee is a free data retrieval call binding the contract method 0xcea74edf.
//
// Solidity: function totalFee(uint256 size) pure returns(uint256)
func (_OracleContract *OracleContractSession) TotalFee(size *big.Int) (*big.Int, error) {
	return _OracleContract.Contract.TotalFee(&_OracleContract.CallOpts, size)
}

// TotalFee is a free data retrieval call binding the contract method 0xcea74edf.
//
// Solidity: function totalFee(uint256 size) pure returns(uint256)
func (_OracleContract *OracleContractCallerSession) TotalFee(size *big.Int) (*big.Int, error) {
	return _OracleContract.Contract.TotalFee(&_OracleContract.CallOpts, size)
}

// SubmitBlockValidationResult is a paid mutator transaction binding the contract method 0xe21312de.
//
// Solidity: function submitBlockValidationResult(bool _result, bytes32 message, uint256 signature, uint256 rx, uint256 ry, uint256 _hash, uint256 pubKeyX, uint256 pubKeyY, address[] validators) returns()
func (_OracleContract *OracleContractTransactor) SubmitBlockValidationResult(opts *bind.TransactOpts, _result bool, message [32]byte, signature *big.Int, rx *big.Int, ry *big.Int, _hash *big.Int, pubKeyX *big.Int, pubKeyY *big.Int, validators []common.Address) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "submitBlockValidationResult", _result, message, signature, rx, ry, _hash, pubKeyX, pubKeyY, validators)
}

// SubmitBlockValidationResult is a paid mutator transaction binding the contract method 0xe21312de.
//
// Solidity: function submitBlockValidationResult(bool _result, bytes32 message, uint256 signature, uint256 rx, uint256 ry, uint256 _hash, uint256 pubKeyX, uint256 pubKeyY, address[] validators) returns()
func (_OracleContract *OracleContractSession) SubmitBlockValidationResult(_result bool, message [32]byte, signature *big.Int, rx *big.Int, ry *big.Int, _hash *big.Int, pubKeyX *big.Int, pubKeyY *big.Int, validators []common.Address) (*types.Transaction, error) {
	return _OracleContract.Contract.SubmitBlockValidationResult(&_OracleContract.TransactOpts, _result, message, signature, rx, ry, _hash, pubKeyX, pubKeyY, validators)
}

// SubmitBlockValidationResult is a paid mutator transaction binding the contract method 0xe21312de.
//
// Solidity: function submitBlockValidationResult(bool _result, bytes32 message, uint256 signature, uint256 rx, uint256 ry, uint256 _hash, uint256 pubKeyX, uint256 pubKeyY, address[] validators) returns()
func (_OracleContract *OracleContractTransactorSession) SubmitBlockValidationResult(_result bool, message [32]byte, signature *big.Int, rx *big.Int, ry *big.Int, _hash *big.Int, pubKeyX *big.Int, pubKeyY *big.Int, validators []common.Address) (*types.Transaction, error) {
	return _OracleContract.Contract.SubmitBlockValidationResult(&_OracleContract.TransactOpts, _result, message, signature, rx, ry, _hash, pubKeyX, pubKeyY, validators)
}

// SubmitTransactionValidationResult is a paid mutator transaction binding the contract method 0x2ef31c3c.
//
// Solidity: function submitTransactionValidationResult(bool _result, bytes32 message, uint256 signature, uint256 rx, uint256 ry, uint256 _hash, uint256 pubKeyX, uint256 pubKeyY, address[] validators) returns()
func (_OracleContract *OracleContractTransactor) SubmitTransactionValidationResult(opts *bind.TransactOpts, _result bool, message [32]byte, signature *big.Int, rx *big.Int, ry *big.Int, _hash *big.Int, pubKeyX *big.Int, pubKeyY *big.Int, validators []common.Address) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "submitTransactionValidationResult", _result, message, signature, rx, ry, _hash, pubKeyX, pubKeyY, validators)
}

// SubmitTransactionValidationResult is a paid mutator transaction binding the contract method 0x2ef31c3c.
//
// Solidity: function submitTransactionValidationResult(bool _result, bytes32 message, uint256 signature, uint256 rx, uint256 ry, uint256 _hash, uint256 pubKeyX, uint256 pubKeyY, address[] validators) returns()
func (_OracleContract *OracleContractSession) SubmitTransactionValidationResult(_result bool, message [32]byte, signature *big.Int, rx *big.Int, ry *big.Int, _hash *big.Int, pubKeyX *big.Int, pubKeyY *big.Int, validators []common.Address) (*types.Transaction, error) {
	return _OracleContract.Contract.SubmitTransactionValidationResult(&_OracleContract.TransactOpts, _result, message, signature, rx, ry, _hash, pubKeyX, pubKeyY, validators)
}

// SubmitTransactionValidationResult is a paid mutator transaction binding the contract method 0x2ef31c3c.
//
// Solidity: function submitTransactionValidationResult(bool _result, bytes32 message, uint256 signature, uint256 rx, uint256 ry, uint256 _hash, uint256 pubKeyX, uint256 pubKeyY, address[] validators) returns()
func (_OracleContract *OracleContractTransactorSession) SubmitTransactionValidationResult(_result bool, message [32]byte, signature *big.Int, rx *big.Int, ry *big.Int, _hash *big.Int, pubKeyX *big.Int, pubKeyY *big.Int, validators []common.Address) (*types.Transaction, error) {
	return _OracleContract.Contract.SubmitTransactionValidationResult(&_OracleContract.TransactOpts, _result, message, signature, rx, ry, _hash, pubKeyX, pubKeyY, validators)
}

// ValidateBlock is a paid mutator transaction binding the contract method 0xce9d294b.
//
// Solidity: function validateBlock(bytes32 _message, uint256 size, uint256 minRank) payable returns()
func (_OracleContract *OracleContractTransactor) ValidateBlock(opts *bind.TransactOpts, _message [32]byte, size *big.Int, minRank *big.Int) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "validateBlock", _message, size, minRank)
}

// ValidateBlock is a paid mutator transaction binding the contract method 0xce9d294b.
//
// Solidity: function validateBlock(bytes32 _message, uint256 size, uint256 minRank) payable returns()
func (_OracleContract *OracleContractSession) ValidateBlock(_message [32]byte, size *big.Int, minRank *big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.ValidateBlock(&_OracleContract.TransactOpts, _message, size, minRank)
}

// ValidateBlock is a paid mutator transaction binding the contract method 0xce9d294b.
//
// Solidity: function validateBlock(bytes32 _message, uint256 size, uint256 minRank) payable returns()
func (_OracleContract *OracleContractTransactorSession) ValidateBlock(_message [32]byte, size *big.Int, minRank *big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.ValidateBlock(&_OracleContract.TransactOpts, _message, size, minRank)
}

// ValidateTransaction is a paid mutator transaction binding the contract method 0x3e170dd9.
//
// Solidity: function validateTransaction(bytes32 _message, uint256 size, uint256 minRank) payable returns()
func (_OracleContract *OracleContractTransactor) ValidateTransaction(opts *bind.TransactOpts, _message [32]byte, size *big.Int, minRank *big.Int) (*types.Transaction, error) {
	return _OracleContract.contract.Transact(opts, "validateTransaction", _message, size, minRank)
}

// ValidateTransaction is a paid mutator transaction binding the contract method 0x3e170dd9.
//
// Solidity: function validateTransaction(bytes32 _message, uint256 size, uint256 minRank) payable returns()
func (_OracleContract *OracleContractSession) ValidateTransaction(_message [32]byte, size *big.Int, minRank *big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.ValidateTransaction(&_OracleContract.TransactOpts, _message, size, minRank)
}

// ValidateTransaction is a paid mutator transaction binding the contract method 0x3e170dd9.
//
// Solidity: function validateTransaction(bytes32 _message, uint256 size, uint256 minRank) payable returns()
func (_OracleContract *OracleContractTransactorSession) ValidateTransaction(_message [32]byte, size *big.Int, minRank *big.Int) (*types.Transaction, error) {
	return _OracleContract.Contract.ValidateTransaction(&_OracleContract.TransactOpts, _message, size, minRank)
}

// OracleContractValidationRequestIterator is returned from FilterValidationRequest and is used to iterate over the raw logs and unpacked data for ValidationRequest events raised by the OracleContract contract.
type OracleContractValidationRequestIterator struct {
	Event *OracleContractValidationRequest // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OracleContractValidationRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleContractValidationRequest)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OracleContractValidationRequest)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OracleContractValidationRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleContractValidationRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleContractValidationRequest represents a ValidationRequest event raised by the OracleContract contract.
type OracleContractValidationRequest struct {
	Typ     uint8
	From    common.Address
	Hash    [32]byte
	Size    *big.Int
	MinRank *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterValidationRequest is a free log retrieval operation binding the contract event 0xa2a630edc48d1fc0a90b61e08bbc34ce73957ce9d59455207cd0591006bd4d4b.
//
// Solidity: event ValidationRequest(uint8 typ, address indexed from, bytes32 hash, uint256 size, uint256 minRank)
func (_OracleContract *OracleContractFilterer) FilterValidationRequest(opts *bind.FilterOpts, from []common.Address) (*OracleContractValidationRequestIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _OracleContract.contract.FilterLogs(opts, "ValidationRequest", fromRule)
	if err != nil {
		return nil, err
	}
	return &OracleContractValidationRequestIterator{contract: _OracleContract.contract, event: "ValidationRequest", logs: logs, sub: sub}, nil
}

// WatchValidationRequest is a free log subscription operation binding the contract event 0xa2a630edc48d1fc0a90b61e08bbc34ce73957ce9d59455207cd0591006bd4d4b.
//
// Solidity: event ValidationRequest(uint8 typ, address indexed from, bytes32 hash, uint256 size, uint256 minRank)
func (_OracleContract *OracleContractFilterer) WatchValidationRequest(opts *bind.WatchOpts, sink chan<- *OracleContractValidationRequest, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _OracleContract.contract.WatchLogs(opts, "ValidationRequest", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleContractValidationRequest)
				if err := _OracleContract.contract.UnpackLog(event, "ValidationRequest", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidationRequest is a log parse operation binding the contract event 0xa2a630edc48d1fc0a90b61e08bbc34ce73957ce9d59455207cd0591006bd4d4b.
//
// Solidity: event ValidationRequest(uint8 typ, address indexed from, bytes32 hash, uint256 size, uint256 minRank)
func (_OracleContract *OracleContractFilterer) ParseValidationRequest(log types.Log) (*OracleContractValidationRequest, error) {
	event := new(OracleContractValidationRequest)
	if err := _OracleContract.contract.UnpackLog(event, "ValidationRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
