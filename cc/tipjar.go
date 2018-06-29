// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cc

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// TipJarABI is the input ABI used to generate the binding from.
const TipJarABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// TipJarBin is the compiled bytecode used for deploying new contracts.
const TipJarBin = `0x608060405234801561001057600080fd5b5060008054600160a060020a033316600160a060020a03199091161790556102068061003d6000396000f3006080604052600436106100615763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166312065fe081146100665780633ccfd60b1461008d578063a6f9dae1146100a4578063b6b55f25146100d2575b600080fd5b34801561007257600080fd5b5061007b6100dd565b60408051918252519081900360200190f35b34801561009957600080fd5b506100a26100f8565b005b3480156100b057600080fd5b506100a273ffffffffffffffffffffffffffffffffffffffff6004351661016a565b6100a26004356101ce565b73ffffffffffffffffffffffffffffffffffffffff30163190565b6000543373ffffffffffffffffffffffffffffffffffffffff90811691161461012057600080fd5b60405173ffffffffffffffffffffffffffffffffffffffff33811691309091163180156108fc02916000818181858888f19350505050158015610167573d6000803e3d6000fd5b50565b6000543373ffffffffffffffffffffffffffffffffffffffff90811691161461019257600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b34811461016757600080fd00a165627a7a723058201bb262a308e23bf00430013e506edb8b668ae381c51af5b20055ae4f3479b8570029`

// DeployTipJar deploys a new Ethereum contract, binding an instance of TipJar to it.
func DeployTipJar(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TipJar, error) {
	parsed, err := abi.JSON(strings.NewReader(TipJarABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TipJarBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TipJar{TipJarCaller: TipJarCaller{contract: contract}, TipJarTransactor: TipJarTransactor{contract: contract}, TipJarFilterer: TipJarFilterer{contract: contract}}, nil
}

// TipJar is an auto generated Go binding around an Ethereum contract.
type TipJar struct {
	TipJarCaller     // Read-only binding to the contract
	TipJarTransactor // Write-only binding to the contract
	TipJarFilterer   // Log filterer for contract events
}

// TipJarCaller is an auto generated read-only Go binding around an Ethereum contract.
type TipJarCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TipJarTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TipJarTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TipJarFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TipJarFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TipJarSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TipJarSession struct {
	Contract     *TipJar           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TipJarCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TipJarCallerSession struct {
	Contract *TipJarCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TipJarTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TipJarTransactorSession struct {
	Contract     *TipJarTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TipJarRaw is an auto generated low-level Go binding around an Ethereum contract.
type TipJarRaw struct {
	Contract *TipJar // Generic contract binding to access the raw methods on
}

// TipJarCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TipJarCallerRaw struct {
	Contract *TipJarCaller // Generic read-only contract binding to access the raw methods on
}

// TipJarTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TipJarTransactorRaw struct {
	Contract *TipJarTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTipJar creates a new instance of TipJar, bound to a specific deployed contract.
func NewTipJar(address common.Address, backend bind.ContractBackend) (*TipJar, error) {
	contract, err := bindTipJar(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TipJar{TipJarCaller: TipJarCaller{contract: contract}, TipJarTransactor: TipJarTransactor{contract: contract}, TipJarFilterer: TipJarFilterer{contract: contract}}, nil
}

// NewTipJarCaller creates a new read-only instance of TipJar, bound to a specific deployed contract.
func NewTipJarCaller(address common.Address, caller bind.ContractCaller) (*TipJarCaller, error) {
	contract, err := bindTipJar(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TipJarCaller{contract: contract}, nil
}

// NewTipJarTransactor creates a new write-only instance of TipJar, bound to a specific deployed contract.
func NewTipJarTransactor(address common.Address, transactor bind.ContractTransactor) (*TipJarTransactor, error) {
	contract, err := bindTipJar(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TipJarTransactor{contract: contract}, nil
}

// NewTipJarFilterer creates a new log filterer instance of TipJar, bound to a specific deployed contract.
func NewTipJarFilterer(address common.Address, filterer bind.ContractFilterer) (*TipJarFilterer, error) {
	contract, err := bindTipJar(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TipJarFilterer{contract: contract}, nil
}

// bindTipJar binds a generic wrapper to an already deployed contract.
func bindTipJar(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TipJarABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TipJar *TipJarRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TipJar.Contract.TipJarCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TipJar *TipJarRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TipJar.Contract.TipJarTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TipJar *TipJarRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TipJar.Contract.TipJarTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TipJar *TipJarCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TipJar.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TipJar *TipJarTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TipJar.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TipJar *TipJarTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TipJar.Contract.contract.Transact(opts, method, params...)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() constant returns(uint256)
func (_TipJar *TipJarCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TipJar.contract.Call(opts, out, "getBalance")
	return *ret0, err
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() constant returns(uint256)
func (_TipJar *TipJarSession) GetBalance() (*big.Int, error) {
	return _TipJar.Contract.GetBalance(&_TipJar.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() constant returns(uint256)
func (_TipJar *TipJarCallerSession) GetBalance() (*big.Int, error) {
	return _TipJar.Contract.GetBalance(&_TipJar.CallOpts)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(newOwner address) returns()
func (_TipJar *TipJarTransactor) ChangeOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TipJar.contract.Transact(opts, "changeOwner", newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(newOwner address) returns()
func (_TipJar *TipJarSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _TipJar.Contract.ChangeOwner(&_TipJar.TransactOpts, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(newOwner address) returns()
func (_TipJar *TipJarTransactorSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _TipJar.Contract.ChangeOwner(&_TipJar.TransactOpts, newOwner)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(amount uint256) returns()
func (_TipJar *TipJarTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _TipJar.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(amount uint256) returns()
func (_TipJar *TipJarSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _TipJar.Contract.Deposit(&_TipJar.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(amount uint256) returns()
func (_TipJar *TipJarTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _TipJar.Contract.Deposit(&_TipJar.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_TipJar *TipJarTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TipJar.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_TipJar *TipJarSession) Withdraw() (*types.Transaction, error) {
	return _TipJar.Contract.Withdraw(&_TipJar.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_TipJar *TipJarTransactorSession) Withdraw() (*types.Transaction, error) {
	return _TipJar.Contract.Withdraw(&_TipJar.TransactOpts)
}
