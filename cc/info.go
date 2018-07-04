// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cc

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// InfoContractABI is the input ABI used to generate the binding from.
const InfoContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_channel\",\"type\":\"uint256\"}],\"name\":\"setChannel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_fName\",\"type\":\"string\"},{\"name\":\"_age\",\"type\":\"uint256\"}],\"name\":\"setInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getChannel\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"channel\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"age\",\"type\":\"uint256\"}],\"name\":\"EventSetInfo\",\"type\":\"event\"}]"

// InfoContractBin is the compiled bytecode used for deploying new contracts.
const InfoContractBin = `0x60806040526000600255600060035534801561001a57600080fd5b506103f18061002a6000396000f3006080604052600436106100615763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663148dfbd781146100665780635a9b0b89146100805780638262963b1461011b578063bf835cba14610176575b600080fd5b34801561007257600080fd5b5061007e60043561019d565b005b34801561008c57600080fd5b506100956101a2565b6040518084815260200180602001838152602001828103825284818151815260200191508051906020019080838360005b838110156100de5781810151838201526020016100c6565b50505050905090810190601f16801561010b5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b34801561012757600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261007e94369492936024939284019190819084018382808284375094975050933594506102579350505050565b34801561018257600080fd5b5061018b610326565b60408051918252519081900360200190f35b600355565b6000606060006002546000600154818054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156102455780601f1061021a57610100808354040283529160200191610245565b820191906000526020600020905b81548152906001019060200180831161022857829003601f168201915b50505050509150925092509250909192565b815161026a90600090602085019061032d565b50806001819055506002546001016002819055506003547fe20534f956fdea37d7b8fa04e966222e377dfaaf8dd1cabb88d1fecc883da79583836040518080602001838152602001828103825284818151815260200191508051906020019080838360005b838110156102e75781810151838201526020016102cf565b50505050905090810190601f1680156103145780820380516001836020036101000a031916815260200191505b50935050505060405180910390a25050565b6003545b90565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061036e57805160ff191683800117855561039b565b8280016001018555821561039b579182015b8281111561039b578251825591602001919060010190610380565b506103a79291506103ab565b5090565b61032a91905b808211156103a757600081556001016103b15600a165627a7a723058206fb4ef932ea5d067bb12fae5b4fbb5d283a275b574c33212978590c7d38179460029`

// DeployInfoContract deploys a new Ethereum contract, binding an instance of InfoContract to it.
func DeployInfoContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *InfoContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InfoContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(InfoContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &InfoContract{InfoContractCaller: InfoContractCaller{contract: contract}, InfoContractTransactor: InfoContractTransactor{contract: contract}, InfoContractFilterer: InfoContractFilterer{contract: contract}}, nil
}

// InfoContract is an auto generated Go binding around an Ethereum contract.
type InfoContract struct {
	InfoContractCaller     // Read-only binding to the contract
	InfoContractTransactor // Write-only binding to the contract
	InfoContractFilterer   // Log filterer for contract events
}

// InfoContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type InfoContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InfoContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InfoContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InfoContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InfoContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InfoContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InfoContractSession struct {
	Contract     *InfoContract     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InfoContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InfoContractCallerSession struct {
	Contract *InfoContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// InfoContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InfoContractTransactorSession struct {
	Contract     *InfoContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// InfoContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type InfoContractRaw struct {
	Contract *InfoContract // Generic contract binding to access the raw methods on
}

// InfoContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InfoContractCallerRaw struct {
	Contract *InfoContractCaller // Generic read-only contract binding to access the raw methods on
}

// InfoContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InfoContractTransactorRaw struct {
	Contract *InfoContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInfoContract creates a new instance of InfoContract, bound to a specific deployed contract.
func NewInfoContract(address common.Address, backend bind.ContractBackend) (*InfoContract, error) {
	contract, err := bindInfoContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InfoContract{InfoContractCaller: InfoContractCaller{contract: contract}, InfoContractTransactor: InfoContractTransactor{contract: contract}, InfoContractFilterer: InfoContractFilterer{contract: contract}}, nil
}

// NewInfoContractCaller creates a new read-only instance of InfoContract, bound to a specific deployed contract.
func NewInfoContractCaller(address common.Address, caller bind.ContractCaller) (*InfoContractCaller, error) {
	contract, err := bindInfoContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InfoContractCaller{contract: contract}, nil
}

// NewInfoContractTransactor creates a new write-only instance of InfoContract, bound to a specific deployed contract.
func NewInfoContractTransactor(address common.Address, transactor bind.ContractTransactor) (*InfoContractTransactor, error) {
	contract, err := bindInfoContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InfoContractTransactor{contract: contract}, nil
}

// NewInfoContractFilterer creates a new log filterer instance of InfoContract, bound to a specific deployed contract.
func NewInfoContractFilterer(address common.Address, filterer bind.ContractFilterer) (*InfoContractFilterer, error) {
	contract, err := bindInfoContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InfoContractFilterer{contract: contract}, nil
}

// bindInfoContract binds a generic wrapper to an already deployed contract.
func bindInfoContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InfoContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InfoContract *InfoContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _InfoContract.Contract.InfoContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InfoContract *InfoContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfoContract.Contract.InfoContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InfoContract *InfoContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InfoContract.Contract.InfoContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InfoContract *InfoContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _InfoContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InfoContract *InfoContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfoContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InfoContract *InfoContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InfoContract.Contract.contract.Transact(opts, method, params...)
}

// GetChannel is a free data retrieval call binding the contract method 0xbf835cba.
//
// Solidity: function getChannel() constant returns(uint256)
func (_InfoContract *InfoContractCaller) GetChannel(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _InfoContract.contract.Call(opts, out, "getChannel")
	return *ret0, err
}

// GetChannel is a free data retrieval call binding the contract method 0xbf835cba.
//
// Solidity: function getChannel() constant returns(uint256)
func (_InfoContract *InfoContractSession) GetChannel() (*big.Int, error) {
	return _InfoContract.Contract.GetChannel(&_InfoContract.CallOpts)
}

// GetChannel is a free data retrieval call binding the contract method 0xbf835cba.
//
// Solidity: function getChannel() constant returns(uint256)
func (_InfoContract *InfoContractCallerSession) GetChannel() (*big.Int, error) {
	return _InfoContract.Contract.GetChannel(&_InfoContract.CallOpts)
}

// GetInfo is a free data retrieval call binding the contract method 0x5a9b0b89.
//
// Solidity: function getInfo() constant returns(uint256, string, uint256)
func (_InfoContract *InfoContractCaller) GetInfo(opts *bind.CallOpts) (*big.Int, string, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(string)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _InfoContract.contract.Call(opts, out, "getInfo")
	return *ret0, *ret1, *ret2, err
}

// GetInfo is a free data retrieval call binding the contract method 0x5a9b0b89.
//
// Solidity: function getInfo() constant returns(uint256, string, uint256)
func (_InfoContract *InfoContractSession) GetInfo() (*big.Int, string, *big.Int, error) {
	return _InfoContract.Contract.GetInfo(&_InfoContract.CallOpts)
}

// GetInfo is a free data retrieval call binding the contract method 0x5a9b0b89.
//
// Solidity: function getInfo() constant returns(uint256, string, uint256)
func (_InfoContract *InfoContractCallerSession) GetInfo() (*big.Int, string, *big.Int, error) {
	return _InfoContract.Contract.GetInfo(&_InfoContract.CallOpts)
}

// SetChannel is a paid mutator transaction binding the contract method 0x148dfbd7.
//
// Solidity: function setChannel(_channel uint256) returns()
func (_InfoContract *InfoContractTransactor) SetChannel(opts *bind.TransactOpts, _channel *big.Int) (*types.Transaction, error) {
	return _InfoContract.contract.Transact(opts, "setChannel", _channel)
}

// SetChannel is a paid mutator transaction binding the contract method 0x148dfbd7.
//
// Solidity: function setChannel(_channel uint256) returns()
func (_InfoContract *InfoContractSession) SetChannel(_channel *big.Int) (*types.Transaction, error) {
	return _InfoContract.Contract.SetChannel(&_InfoContract.TransactOpts, _channel)
}

// SetChannel is a paid mutator transaction binding the contract method 0x148dfbd7.
//
// Solidity: function setChannel(_channel uint256) returns()
func (_InfoContract *InfoContractTransactorSession) SetChannel(_channel *big.Int) (*types.Transaction, error) {
	return _InfoContract.Contract.SetChannel(&_InfoContract.TransactOpts, _channel)
}

// SetInfo is a paid mutator transaction binding the contract method 0x8262963b.
//
// Solidity: function setInfo(_fName string, _age uint256) returns()
func (_InfoContract *InfoContractTransactor) SetInfo(opts *bind.TransactOpts, _fName string, _age *big.Int) (*types.Transaction, error) {
	return _InfoContract.contract.Transact(opts, "setInfo", _fName, _age)
}

// SetInfo is a paid mutator transaction binding the contract method 0x8262963b.
//
// Solidity: function setInfo(_fName string, _age uint256) returns()
func (_InfoContract *InfoContractSession) SetInfo(_fName string, _age *big.Int) (*types.Transaction, error) {
	return _InfoContract.Contract.SetInfo(&_InfoContract.TransactOpts, _fName, _age)
}

// SetInfo is a paid mutator transaction binding the contract method 0x8262963b.
//
// Solidity: function setInfo(_fName string, _age uint256) returns()
func (_InfoContract *InfoContractTransactorSession) SetInfo(_fName string, _age *big.Int) (*types.Transaction, error) {
	return _InfoContract.Contract.SetInfo(&_InfoContract.TransactOpts, _fName, _age)
}

// InfoContractEventSetInfoIterator is returned from FilterEventSetInfo and is used to iterate over the raw logs and unpacked data for EventSetInfo events raised by the InfoContract contract.
type InfoContractEventSetInfoIterator struct {
	Event *InfoContractEventSetInfo // Event containing the contract specifics and raw log

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
func (it *InfoContractEventSetInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfoContractEventSetInfo)
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
		it.Event = new(InfoContractEventSetInfo)
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
func (it *InfoContractEventSetInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfoContractEventSetInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfoContractEventSetInfo represents a EventSetInfo event raised by the InfoContract contract.
type InfoContractEventSetInfo struct {
	Channel *big.Int
	Name    string
	Age     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEventSetInfo is a free log retrieval operation binding the contract event 0xe20534f956fdea37d7b8fa04e966222e377dfaaf8dd1cabb88d1fecc883da795.
//
// Solidity: e EventSetInfo(channel indexed uint256, name string, age uint256)
func (_InfoContract *InfoContractFilterer) FilterEventSetInfo(opts *bind.FilterOpts, channel []*big.Int) (*InfoContractEventSetInfoIterator, error) {

	var channelRule []interface{}
	for _, channelItem := range channel {
		channelRule = append(channelRule, channelItem)
	}

	logs, sub, err := _InfoContract.contract.FilterLogs(opts, "EventSetInfo", channelRule)
	if err != nil {
		return nil, err
	}
	return &InfoContractEventSetInfoIterator{contract: _InfoContract.contract, event: "EventSetInfo", logs: logs, sub: sub}, nil
}

// WatchEventSetInfo is a free log subscription operation binding the contract event 0xe20534f956fdea37d7b8fa04e966222e377dfaaf8dd1cabb88d1fecc883da795.
//
// Solidity: e EventSetInfo(channel indexed uint256, name string, age uint256)
func (_InfoContract *InfoContractFilterer) WatchEventSetInfo(opts *bind.WatchOpts, sink chan<- *InfoContractEventSetInfo, channel []*big.Int) (event.Subscription, error) {

	var channelRule []interface{}
	for _, channelItem := range channel {
		channelRule = append(channelRule, channelItem)
	}

	logs, sub, err := _InfoContract.contract.WatchLogs(opts, "EventSetInfo", channelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfoContractEventSetInfo)
				if err := _InfoContract.contract.UnpackLog(event, "EventSetInfo", log); err != nil {
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
