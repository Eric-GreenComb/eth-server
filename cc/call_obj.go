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

// CallObjectABI is the input ABI used to generate the binding from.
const CallObjectABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"addrObj\",\"type\":\"address\"},{\"name\":\"val\",\"type\":\"string\"}],\"name\":\"ModifyValue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CallObjectBin is the compiled bytecode used for deploying new contracts.
const CallObjectBin = `0x608060405234801561001057600080fd5b506101cf806100206000396000f3006080604052600436106100405763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166340d85bfe8114610045575b600080fd5b34801561005157600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526100b995833573ffffffffffffffffffffffffffffffffffffffff169536956044949193909101919081908401838280828437509497506100bb9650505050505050565b005b6040517fbc45be6300000000000000000000000000000000000000000000000000000000815260206004820181815283516024840152835173ffffffffffffffffffffffffffffffffffffffff86169363bc45be639386939283926044019185019080838360005b8381101561013b578181015183820152602001610123565b50505050905090810190601f1680156101685780820380516001836020036101000a031916815260200191505b5092505050600060405180830381600087803b15801561018757600080fd5b505af115801561019b573d6000803e3d6000fd5b5050505050505600a165627a7a723058201fd57ba1fb26c70c7041cd6d0e177c26ba9a2627d1dc9e170ff8a400858dd3eb0029`

// DeployCallObject deploys a new Ethereum contract, binding an instance of CallObject to it.
func DeployCallObject(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CallObject, error) {
	parsed, err := abi.JSON(strings.NewReader(CallObjectABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CallObjectBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CallObject{CallObjectCaller: CallObjectCaller{contract: contract}, CallObjectTransactor: CallObjectTransactor{contract: contract}, CallObjectFilterer: CallObjectFilterer{contract: contract}}, nil
}

// CallObject is an auto generated Go binding around an Ethereum contract.
type CallObject struct {
	CallObjectCaller     // Read-only binding to the contract
	CallObjectTransactor // Write-only binding to the contract
	CallObjectFilterer   // Log filterer for contract events
}

// CallObjectCaller is an auto generated read-only Go binding around an Ethereum contract.
type CallObjectCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallObjectTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CallObjectTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallObjectFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CallObjectFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallObjectSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CallObjectSession struct {
	Contract     *CallObject       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CallObjectCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CallObjectCallerSession struct {
	Contract *CallObjectCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CallObjectTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CallObjectTransactorSession struct {
	Contract     *CallObjectTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CallObjectRaw is an auto generated low-level Go binding around an Ethereum contract.
type CallObjectRaw struct {
	Contract *CallObject // Generic contract binding to access the raw methods on
}

// CallObjectCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CallObjectCallerRaw struct {
	Contract *CallObjectCaller // Generic read-only contract binding to access the raw methods on
}

// CallObjectTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CallObjectTransactorRaw struct {
	Contract *CallObjectTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCallObject creates a new instance of CallObject, bound to a specific deployed contract.
func NewCallObject(address common.Address, backend bind.ContractBackend) (*CallObject, error) {
	contract, err := bindCallObject(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CallObject{CallObjectCaller: CallObjectCaller{contract: contract}, CallObjectTransactor: CallObjectTransactor{contract: contract}, CallObjectFilterer: CallObjectFilterer{contract: contract}}, nil
}

// NewCallObjectCaller creates a new read-only instance of CallObject, bound to a specific deployed contract.
func NewCallObjectCaller(address common.Address, caller bind.ContractCaller) (*CallObjectCaller, error) {
	contract, err := bindCallObject(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CallObjectCaller{contract: contract}, nil
}

// NewCallObjectTransactor creates a new write-only instance of CallObject, bound to a specific deployed contract.
func NewCallObjectTransactor(address common.Address, transactor bind.ContractTransactor) (*CallObjectTransactor, error) {
	contract, err := bindCallObject(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CallObjectTransactor{contract: contract}, nil
}

// NewCallObjectFilterer creates a new log filterer instance of CallObject, bound to a specific deployed contract.
func NewCallObjectFilterer(address common.Address, filterer bind.ContractFilterer) (*CallObjectFilterer, error) {
	contract, err := bindCallObject(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CallObjectFilterer{contract: contract}, nil
}

// bindCallObject binds a generic wrapper to an already deployed contract.
func bindCallObject(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CallObjectABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CallObject *CallObjectRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CallObject.Contract.CallObjectCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CallObject *CallObjectRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CallObject.Contract.CallObjectTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CallObject *CallObjectRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CallObject.Contract.CallObjectTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CallObject *CallObjectCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CallObject.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CallObject *CallObjectTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CallObject.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CallObject *CallObjectTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CallObject.Contract.contract.Transact(opts, method, params...)
}

// ModifyValue is a paid mutator transaction binding the contract method 0x40d85bfe.
//
// Solidity: function ModifyValue(addrObj address, val string) returns()
func (_CallObject *CallObjectTransactor) ModifyValue(opts *bind.TransactOpts, addrObj common.Address, val string) (*types.Transaction, error) {
	return _CallObject.contract.Transact(opts, "ModifyValue", addrObj, val)
}

// ModifyValue is a paid mutator transaction binding the contract method 0x40d85bfe.
//
// Solidity: function ModifyValue(addrObj address, val string) returns()
func (_CallObject *CallObjectSession) ModifyValue(addrObj common.Address, val string) (*types.Transaction, error) {
	return _CallObject.Contract.ModifyValue(&_CallObject.TransactOpts, addrObj, val)
}

// ModifyValue is a paid mutator transaction binding the contract method 0x40d85bfe.
//
// Solidity: function ModifyValue(addrObj address, val string) returns()
func (_CallObject *CallObjectTransactorSession) ModifyValue(addrObj common.Address, val string) (*types.Transaction, error) {
	return _CallObject.Contract.ModifyValue(&_CallObject.TransactOpts, addrObj, val)
}

// ObjectInterfaceABI is the input ABI used to generate the binding from.
const ObjectInterfaceABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"val\",\"type\":\"string\"}],\"name\":\"ModifyValue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ObjectInterfaceBin is the compiled bytecode used for deploying new contracts.
const ObjectInterfaceBin = `0x`

// DeployObjectInterface deploys a new Ethereum contract, binding an instance of ObjectInterface to it.
func DeployObjectInterface(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ObjectInterface, error) {
	parsed, err := abi.JSON(strings.NewReader(ObjectInterfaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ObjectInterfaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ObjectInterface{ObjectInterfaceCaller: ObjectInterfaceCaller{contract: contract}, ObjectInterfaceTransactor: ObjectInterfaceTransactor{contract: contract}, ObjectInterfaceFilterer: ObjectInterfaceFilterer{contract: contract}}, nil
}

// ObjectInterface is an auto generated Go binding around an Ethereum contract.
type ObjectInterface struct {
	ObjectInterfaceCaller     // Read-only binding to the contract
	ObjectInterfaceTransactor // Write-only binding to the contract
	ObjectInterfaceFilterer   // Log filterer for contract events
}

// ObjectInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ObjectInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ObjectInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ObjectInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ObjectInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ObjectInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ObjectInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ObjectInterfaceSession struct {
	Contract     *ObjectInterface  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ObjectInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ObjectInterfaceCallerSession struct {
	Contract *ObjectInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ObjectInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ObjectInterfaceTransactorSession struct {
	Contract     *ObjectInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ObjectInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ObjectInterfaceRaw struct {
	Contract *ObjectInterface // Generic contract binding to access the raw methods on
}

// ObjectInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ObjectInterfaceCallerRaw struct {
	Contract *ObjectInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// ObjectInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ObjectInterfaceTransactorRaw struct {
	Contract *ObjectInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewObjectInterface creates a new instance of ObjectInterface, bound to a specific deployed contract.
func NewObjectInterface(address common.Address, backend bind.ContractBackend) (*ObjectInterface, error) {
	contract, err := bindObjectInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ObjectInterface{ObjectInterfaceCaller: ObjectInterfaceCaller{contract: contract}, ObjectInterfaceTransactor: ObjectInterfaceTransactor{contract: contract}, ObjectInterfaceFilterer: ObjectInterfaceFilterer{contract: contract}}, nil
}

// NewObjectInterfaceCaller creates a new read-only instance of ObjectInterface, bound to a specific deployed contract.
func NewObjectInterfaceCaller(address common.Address, caller bind.ContractCaller) (*ObjectInterfaceCaller, error) {
	contract, err := bindObjectInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ObjectInterfaceCaller{contract: contract}, nil
}

// NewObjectInterfaceTransactor creates a new write-only instance of ObjectInterface, bound to a specific deployed contract.
func NewObjectInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*ObjectInterfaceTransactor, error) {
	contract, err := bindObjectInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ObjectInterfaceTransactor{contract: contract}, nil
}

// NewObjectInterfaceFilterer creates a new log filterer instance of ObjectInterface, bound to a specific deployed contract.
func NewObjectInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*ObjectInterfaceFilterer, error) {
	contract, err := bindObjectInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ObjectInterfaceFilterer{contract: contract}, nil
}

// bindObjectInterface binds a generic wrapper to an already deployed contract.
func bindObjectInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ObjectInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ObjectInterface *ObjectInterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ObjectInterface.Contract.ObjectInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ObjectInterface *ObjectInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ObjectInterface.Contract.ObjectInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ObjectInterface *ObjectInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ObjectInterface.Contract.ObjectInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ObjectInterface *ObjectInterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ObjectInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ObjectInterface *ObjectInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ObjectInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ObjectInterface *ObjectInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ObjectInterface.Contract.contract.Transact(opts, method, params...)
}

// ModifyValue is a paid mutator transaction binding the contract method 0xbc45be63.
//
// Solidity: function ModifyValue(val string) returns()
func (_ObjectInterface *ObjectInterfaceTransactor) ModifyValue(opts *bind.TransactOpts, val string) (*types.Transaction, error) {
	return _ObjectInterface.contract.Transact(opts, "ModifyValue", val)
}

// ModifyValue is a paid mutator transaction binding the contract method 0xbc45be63.
//
// Solidity: function ModifyValue(val string) returns()
func (_ObjectInterface *ObjectInterfaceSession) ModifyValue(val string) (*types.Transaction, error) {
	return _ObjectInterface.Contract.ModifyValue(&_ObjectInterface.TransactOpts, val)
}

// ModifyValue is a paid mutator transaction binding the contract method 0xbc45be63.
//
// Solidity: function ModifyValue(val string) returns()
func (_ObjectInterface *ObjectInterfaceTransactorSession) ModifyValue(val string) (*types.Transaction, error) {
	return _ObjectInterface.Contract.ModifyValue(&_ObjectInterface.TransactOpts, val)
}
