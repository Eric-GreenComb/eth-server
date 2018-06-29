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

// SavingsABI is the input ABI used to generate the binding from.
const SavingsABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"numberOfDays\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"}]"

// SavingsBin is the compiled bytecode used for deploying new contracts.
const SavingsBin = `0x608060405260405160208061016c833981016040525160008054600160a060020a033316600160a060020a031990911617905542620151809091020160015561011f8061004d6000396000f30060806040526004361060485763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633ccfd60b8114604d578063b6b55f25146061575b600080fd5b348015605857600080fd5b50605f606a565b005b605f60043560e8565b6000543373ffffffffffffffffffffffffffffffffffffffff908116911614609157600080fd5b600154421015609f57600080fd5b60405173ffffffffffffffffffffffffffffffffffffffff33811691309091163180156108fc02916000818181858888f1935050505015801560e5573d6000803e3d6000fd5b50565b34811460e557600080fd00a165627a7a72305820507b28a8d2c9df1b7e2b9ee0b1bded1b18dd8e3f8b5ac2c975401851ae08c97b0029`

// DeploySavings deploys a new Ethereum contract, binding an instance of Savings to it.
func DeploySavings(auth *bind.TransactOpts, backend bind.ContractBackend, numberOfDays *big.Int) (common.Address, *types.Transaction, *Savings, error) {
	parsed, err := abi.JSON(strings.NewReader(SavingsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SavingsBin), backend, numberOfDays)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Savings{SavingsCaller: SavingsCaller{contract: contract}, SavingsTransactor: SavingsTransactor{contract: contract}, SavingsFilterer: SavingsFilterer{contract: contract}}, nil
}

// Savings is an auto generated Go binding around an Ethereum contract.
type Savings struct {
	SavingsCaller     // Read-only binding to the contract
	SavingsTransactor // Write-only binding to the contract
	SavingsFilterer   // Log filterer for contract events
}

// SavingsCaller is an auto generated read-only Go binding around an Ethereum contract.
type SavingsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SavingsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SavingsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SavingsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SavingsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SavingsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SavingsSession struct {
	Contract     *Savings          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SavingsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SavingsCallerSession struct {
	Contract *SavingsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// SavingsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SavingsTransactorSession struct {
	Contract     *SavingsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SavingsRaw is an auto generated low-level Go binding around an Ethereum contract.
type SavingsRaw struct {
	Contract *Savings // Generic contract binding to access the raw methods on
}

// SavingsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SavingsCallerRaw struct {
	Contract *SavingsCaller // Generic read-only contract binding to access the raw methods on
}

// SavingsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SavingsTransactorRaw struct {
	Contract *SavingsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSavings creates a new instance of Savings, bound to a specific deployed contract.
func NewSavings(address common.Address, backend bind.ContractBackend) (*Savings, error) {
	contract, err := bindSavings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Savings{SavingsCaller: SavingsCaller{contract: contract}, SavingsTransactor: SavingsTransactor{contract: contract}, SavingsFilterer: SavingsFilterer{contract: contract}}, nil
}

// NewSavingsCaller creates a new read-only instance of Savings, bound to a specific deployed contract.
func NewSavingsCaller(address common.Address, caller bind.ContractCaller) (*SavingsCaller, error) {
	contract, err := bindSavings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SavingsCaller{contract: contract}, nil
}

// NewSavingsTransactor creates a new write-only instance of Savings, bound to a specific deployed contract.
func NewSavingsTransactor(address common.Address, transactor bind.ContractTransactor) (*SavingsTransactor, error) {
	contract, err := bindSavings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SavingsTransactor{contract: contract}, nil
}

// NewSavingsFilterer creates a new log filterer instance of Savings, bound to a specific deployed contract.
func NewSavingsFilterer(address common.Address, filterer bind.ContractFilterer) (*SavingsFilterer, error) {
	contract, err := bindSavings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SavingsFilterer{contract: contract}, nil
}

// bindSavings binds a generic wrapper to an already deployed contract.
func bindSavings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SavingsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Savings *SavingsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Savings.Contract.SavingsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Savings *SavingsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Savings.Contract.SavingsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Savings *SavingsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Savings.Contract.SavingsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Savings *SavingsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Savings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Savings *SavingsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Savings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Savings *SavingsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Savings.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(amount uint256) returns()
func (_Savings *SavingsTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Savings.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(amount uint256) returns()
func (_Savings *SavingsSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Savings.Contract.Deposit(&_Savings.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(amount uint256) returns()
func (_Savings *SavingsTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Savings.Contract.Deposit(&_Savings.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Savings *SavingsTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Savings.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Savings *SavingsSession) Withdraw() (*types.Transaction, error) {
	return _Savings.Contract.Withdraw(&_Savings.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Savings *SavingsTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Savings.Contract.Withdraw(&_Savings.TransactOpts)
}
