// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cc

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ObjectABI is the input ABI used to generate the binding from.
const ObjectABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"key\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"value\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"val\",\"type\":\"string\"}],\"name\":\"ModifyValue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ObjectBin is the compiled bytecode used for deploying new contracts.
const ObjectBin = `0x608060405234801561001057600080fd5b5061031b806100206000396000f3006080604052600436106100565763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633943380c811461005b5780633fa4f245146100e5578063bc45be63146100fa575b600080fd5b34801561006757600080fd5b50610070610155565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100aa578181015183820152602001610092565b50505050905090810190601f1680156100d75780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156100f157600080fd5b506100706101e2565b34801561010657600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261015394369492936024939284019190819084018382808284375094975061023d9650505050505050565b005b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156101da5780601f106101af576101008083540402835291602001916101da565b820191906000526020600020905b8154815290600101906020018083116101bd57829003601f168201915b505050505081565b6000805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156101da5780601f106101af576101008083540402835291602001916101da565b8051610250906000906020840190610254565b5050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061029557805160ff19168380011785556102c2565b828001600101855582156102c2579182015b828111156102c25782518255916020019190600101906102a7565b506102ce9291506102d2565b5090565b6102ec91905b808211156102ce57600081556001016102d8565b905600a165627a7a72305820c00f1884ea5437abbc82a1c4cd5ed7f4e6e80082b6c9504e763ae1e5368d41030029`

// DeployObject deploys a new Ethereum contract, binding an instance of Object to it.
func DeployObject(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Object, error) {
	parsed, err := abi.JSON(strings.NewReader(ObjectABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ObjectBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Object{ObjectCaller: ObjectCaller{contract: contract}, ObjectTransactor: ObjectTransactor{contract: contract}, ObjectFilterer: ObjectFilterer{contract: contract}}, nil
}

// Object is an auto generated Go binding around an Ethereum contract.
type Object struct {
	ObjectCaller     // Read-only binding to the contract
	ObjectTransactor // Write-only binding to the contract
	ObjectFilterer   // Log filterer for contract events
}

// ObjectCaller is an auto generated read-only Go binding around an Ethereum contract.
type ObjectCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ObjectTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ObjectTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ObjectFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ObjectFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ObjectSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ObjectSession struct {
	Contract     *Object           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ObjectCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ObjectCallerSession struct {
	Contract *ObjectCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ObjectTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ObjectTransactorSession struct {
	Contract     *ObjectTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ObjectRaw is an auto generated low-level Go binding around an Ethereum contract.
type ObjectRaw struct {
	Contract *Object // Generic contract binding to access the raw methods on
}

// ObjectCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ObjectCallerRaw struct {
	Contract *ObjectCaller // Generic read-only contract binding to access the raw methods on
}

// ObjectTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ObjectTransactorRaw struct {
	Contract *ObjectTransactor // Generic write-only contract binding to access the raw methods on
}

// NewObject creates a new instance of Object, bound to a specific deployed contract.
func NewObject(address common.Address, backend bind.ContractBackend) (*Object, error) {
	contract, err := bindObject(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Object{ObjectCaller: ObjectCaller{contract: contract}, ObjectTransactor: ObjectTransactor{contract: contract}, ObjectFilterer: ObjectFilterer{contract: contract}}, nil
}

// NewObjectCaller creates a new read-only instance of Object, bound to a specific deployed contract.
func NewObjectCaller(address common.Address, caller bind.ContractCaller) (*ObjectCaller, error) {
	contract, err := bindObject(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ObjectCaller{contract: contract}, nil
}

// NewObjectTransactor creates a new write-only instance of Object, bound to a specific deployed contract.
func NewObjectTransactor(address common.Address, transactor bind.ContractTransactor) (*ObjectTransactor, error) {
	contract, err := bindObject(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ObjectTransactor{contract: contract}, nil
}

// NewObjectFilterer creates a new log filterer instance of Object, bound to a specific deployed contract.
func NewObjectFilterer(address common.Address, filterer bind.ContractFilterer) (*ObjectFilterer, error) {
	contract, err := bindObject(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ObjectFilterer{contract: contract}, nil
}

// bindObject binds a generic wrapper to an already deployed contract.
func bindObject(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ObjectABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Object *ObjectRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Object.Contract.ObjectCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Object *ObjectRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Object.Contract.ObjectTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Object *ObjectRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Object.Contract.ObjectTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Object *ObjectCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Object.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Object *ObjectTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Object.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Object *ObjectTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Object.Contract.contract.Transact(opts, method, params...)
}

// Key is a free data retrieval call binding the contract method 0x3943380c.
//
// Solidity: function key() constant returns(string)
func (_Object *ObjectCaller) Key(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Object.contract.Call(opts, out, "key")
	return *ret0, err
}

// Key is a free data retrieval call binding the contract method 0x3943380c.
//
// Solidity: function key() constant returns(string)
func (_Object *ObjectSession) Key() (string, error) {
	return _Object.Contract.Key(&_Object.CallOpts)
}

// Key is a free data retrieval call binding the contract method 0x3943380c.
//
// Solidity: function key() constant returns(string)
func (_Object *ObjectCallerSession) Key() (string, error) {
	return _Object.Contract.Key(&_Object.CallOpts)
}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() constant returns(string)
func (_Object *ObjectCaller) Value(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Object.contract.Call(opts, out, "value")
	return *ret0, err
}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() constant returns(string)
func (_Object *ObjectSession) Value() (string, error) {
	return _Object.Contract.Value(&_Object.CallOpts)
}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() constant returns(string)
func (_Object *ObjectCallerSession) Value() (string, error) {
	return _Object.Contract.Value(&_Object.CallOpts)
}

// ModifyValue is a paid mutator transaction binding the contract method 0xbc45be63.
//
// Solidity: function ModifyValue(val string) returns()
func (_Object *ObjectTransactor) ModifyValue(opts *bind.TransactOpts, val string) (*types.Transaction, error) {
	return _Object.contract.Transact(opts, "ModifyValue", val)
}

// ModifyValue is a paid mutator transaction binding the contract method 0xbc45be63.
//
// Solidity: function ModifyValue(val string) returns()
func (_Object *ObjectSession) ModifyValue(val string) (*types.Transaction, error) {
	return _Object.Contract.ModifyValue(&_Object.TransactOpts, val)
}

// ModifyValue is a paid mutator transaction binding the contract method 0xbc45be63.
//
// Solidity: function ModifyValue(val string) returns()
func (_Object *ObjectTransactorSession) ModifyValue(val string) (*types.Transaction, error) {
	return _Object.Contract.ModifyValue(&_Object.TransactOpts, val)
}
