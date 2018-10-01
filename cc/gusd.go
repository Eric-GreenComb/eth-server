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

// CustodianUpgradeableABI is the input ABI used to generate the binding from.
const CustodianUpgradeableABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_proposedCustodian\",\"type\":\"address\"}],\"name\":\"requestCustodianChange\",\"outputs\":[{\"name\":\"lockId\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"custodian\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_lockId\",\"type\":\"bytes32\"}],\"name\":\"confirmCustodianChange\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lockRequestCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"custodianChangeReqs\",\"outputs\":[{\"name\":\"proposedNew\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_custodian\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_lockId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_msgSender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_proposedCustodian\",\"type\":\"address\"}],\"name\":\"CustodianChangeRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_lockId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_newCustodian\",\"type\":\"address\"}],\"name\":\"CustodianChangeConfirmed\",\"type\":\"event\"}]"

// CustodianUpgradeableBin is the compiled bytecode used for deploying new contracts.
const CustodianUpgradeableBin = `0x608060405234801561001057600080fd5b5060405160208061039783398101604052516000805560018054600160a060020a03909216600160a060020a0319909216919091179055610341806100566000396000f30060806040526004361061006c5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166315b210828114610071578063375b74c3146100a45780633a8343ee146100d5578063cb81fecf146100ef578063cf6e448814610104575b600080fd5b34801561007d57600080fd5b50610092600160a060020a036004351661011c565b60408051918252519081900360200190f35b3480156100b057600080fd5b506100b96101cf565b60408051600160a060020a039092168252519081900360200190f35b3480156100e157600080fd5b506100ed6004356101de565b005b3480156100fb57600080fd5b5061009261027f565b34801561011057600080fd5b506100b9600435610285565b6000600160a060020a038216151561013357600080fd5b61013b6102a0565b6040805160208181018352600160a060020a0380871680845260008681526002845285902093518454921673ffffffffffffffffffffffffffffffffffffffff1990921691909117909255825184815233918101919091528083019190915290519192507fd76fc900a7e1a6fcf11d54b7ba943918df6c53a3128140658c389b3da1e997ba919081900360600190a1919050565b600154600160a060020a031681565b600154600160a060020a031633146101f557600080fd5b6101fe816102e1565b60018054600160a060020a0392831673ffffffffffffffffffffffffffffffffffffffff1991821617825560008481526002602090815260409182902080549093169092559154825185815293169083015280517f9a99272c0f6b7a30ef9e76e684a7cd408bfd4f11a72f36a8e276253c920e442d9281900390910190a150565b60005481565b600260205260009081526040902054600160a060020a031681565b60008054600101908190556040805160001943014081526c010000000000000000000000003002602082015260348101929092525190819003605401902090565b60008181526002602052604081208054600160a060020a0316151561030557600080fd5b54600160a060020a0316929150505600a165627a7a72305820a6c9d10c713f0f9f5bfaf08431403a8d887d377489c8196b855106386469c06c0029`

// DeployCustodianUpgradeable deploys a new Ethereum contract, binding an instance of CustodianUpgradeable to it.
func DeployCustodianUpgradeable(auth *bind.TransactOpts, backend bind.ContractBackend, _custodian common.Address) (common.Address, *types.Transaction, *CustodianUpgradeable, error) {
	parsed, err := abi.JSON(strings.NewReader(CustodianUpgradeableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CustodianUpgradeableBin), backend, _custodian)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CustodianUpgradeable{CustodianUpgradeableCaller: CustodianUpgradeableCaller{contract: contract}, CustodianUpgradeableTransactor: CustodianUpgradeableTransactor{contract: contract}, CustodianUpgradeableFilterer: CustodianUpgradeableFilterer{contract: contract}}, nil
}

// CustodianUpgradeable is an auto generated Go binding around an Ethereum contract.
type CustodianUpgradeable struct {
	CustodianUpgradeableCaller     // Read-only binding to the contract
	CustodianUpgradeableTransactor // Write-only binding to the contract
	CustodianUpgradeableFilterer   // Log filterer for contract events
}

// CustodianUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type CustodianUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CustodianUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CustodianUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CustodianUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CustodianUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CustodianUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CustodianUpgradeableSession struct {
	Contract     *CustodianUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CustodianUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CustodianUpgradeableCallerSession struct {
	Contract *CustodianUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// CustodianUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CustodianUpgradeableTransactorSession struct {
	Contract     *CustodianUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// CustodianUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type CustodianUpgradeableRaw struct {
	Contract *CustodianUpgradeable // Generic contract binding to access the raw methods on
}

// CustodianUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CustodianUpgradeableCallerRaw struct {
	Contract *CustodianUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// CustodianUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CustodianUpgradeableTransactorRaw struct {
	Contract *CustodianUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCustodianUpgradeable creates a new instance of CustodianUpgradeable, bound to a specific deployed contract.
func NewCustodianUpgradeable(address common.Address, backend bind.ContractBackend) (*CustodianUpgradeable, error) {
	contract, err := bindCustodianUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CustodianUpgradeable{CustodianUpgradeableCaller: CustodianUpgradeableCaller{contract: contract}, CustodianUpgradeableTransactor: CustodianUpgradeableTransactor{contract: contract}, CustodianUpgradeableFilterer: CustodianUpgradeableFilterer{contract: contract}}, nil
}

// NewCustodianUpgradeableCaller creates a new read-only instance of CustodianUpgradeable, bound to a specific deployed contract.
func NewCustodianUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*CustodianUpgradeableCaller, error) {
	contract, err := bindCustodianUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CustodianUpgradeableCaller{contract: contract}, nil
}

// NewCustodianUpgradeableTransactor creates a new write-only instance of CustodianUpgradeable, bound to a specific deployed contract.
func NewCustodianUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*CustodianUpgradeableTransactor, error) {
	contract, err := bindCustodianUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CustodianUpgradeableTransactor{contract: contract}, nil
}

// NewCustodianUpgradeableFilterer creates a new log filterer instance of CustodianUpgradeable, bound to a specific deployed contract.
func NewCustodianUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*CustodianUpgradeableFilterer, error) {
	contract, err := bindCustodianUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CustodianUpgradeableFilterer{contract: contract}, nil
}

// bindCustodianUpgradeable binds a generic wrapper to an already deployed contract.
func bindCustodianUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CustodianUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CustodianUpgradeable *CustodianUpgradeableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CustodianUpgradeable.Contract.CustodianUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CustodianUpgradeable *CustodianUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CustodianUpgradeable.Contract.CustodianUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CustodianUpgradeable *CustodianUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CustodianUpgradeable.Contract.CustodianUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CustodianUpgradeable *CustodianUpgradeableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CustodianUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CustodianUpgradeable *CustodianUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CustodianUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CustodianUpgradeable *CustodianUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CustodianUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// Custodian is a free data retrieval call binding the contract method 0x375b74c3.
//
// Solidity: function custodian() constant returns(address)
func (_CustodianUpgradeable *CustodianUpgradeableCaller) Custodian(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CustodianUpgradeable.contract.Call(opts, out, "custodian")
	return *ret0, err
}

// Custodian is a free data retrieval call binding the contract method 0x375b74c3.
//
// Solidity: function custodian() constant returns(address)
func (_CustodianUpgradeable *CustodianUpgradeableSession) Custodian() (common.Address, error) {
	return _CustodianUpgradeable.Contract.Custodian(&_CustodianUpgradeable.CallOpts)
}

// Custodian is a free data retrieval call binding the contract method 0x375b74c3.
//
// Solidity: function custodian() constant returns(address)
func (_CustodianUpgradeable *CustodianUpgradeableCallerSession) Custodian() (common.Address, error) {
	return _CustodianUpgradeable.Contract.Custodian(&_CustodianUpgradeable.CallOpts)
}

// CustodianChangeReqs is a free data retrieval call binding the contract method 0xcf6e4488.
//
// Solidity: function custodianChangeReqs( bytes32) constant returns(proposedNew address)
func (_CustodianUpgradeable *CustodianUpgradeableCaller) CustodianChangeReqs(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CustodianUpgradeable.contract.Call(opts, out, "custodianChangeReqs", arg0)
	return *ret0, err
}

// CustodianChangeReqs is a free data retrieval call binding the contract method 0xcf6e4488.
//
// Solidity: function custodianChangeReqs( bytes32) constant returns(proposedNew address)
func (_CustodianUpgradeable *CustodianUpgradeableSession) CustodianChangeReqs(arg0 [32]byte) (common.Address, error) {
	return _CustodianUpgradeable.Contract.CustodianChangeReqs(&_CustodianUpgradeable.CallOpts, arg0)
}

// CustodianChangeReqs is a free data retrieval call binding the contract method 0xcf6e4488.
//
// Solidity: function custodianChangeReqs( bytes32) constant returns(proposedNew address)
func (_CustodianUpgradeable *CustodianUpgradeableCallerSession) CustodianChangeReqs(arg0 [32]byte) (common.Address, error) {
	return _CustodianUpgradeable.Contract.CustodianChangeReqs(&_CustodianUpgradeable.CallOpts, arg0)
}

// LockRequestCount is a free data retrieval call binding the contract method 0xcb81fecf.
//
// Solidity: function lockRequestCount() constant returns(uint256)
func (_CustodianUpgradeable *CustodianUpgradeableCaller) LockRequestCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CustodianUpgradeable.contract.Call(opts, out, "lockRequestCount")
	return *ret0, err
}

// LockRequestCount is a free data retrieval call binding the contract method 0xcb81fecf.
//
// Solidity: function lockRequestCount() constant returns(uint256)
func (_CustodianUpgradeable *CustodianUpgradeableSession) LockRequestCount() (*big.Int, error) {
	return _CustodianUpgradeable.Contract.LockRequestCount(&_CustodianUpgradeable.CallOpts)
}

// LockRequestCount is a free data retrieval call binding the contract method 0xcb81fecf.
//
// Solidity: function lockRequestCount() constant returns(uint256)
func (_CustodianUpgradeable *CustodianUpgradeableCallerSession) LockRequestCount() (*big.Int, error) {
	return _CustodianUpgradeable.Contract.LockRequestCount(&_CustodianUpgradeable.CallOpts)
}

// ConfirmCustodianChange is a paid mutator transaction binding the contract method 0x3a8343ee.
//
// Solidity: function confirmCustodianChange(_lockId bytes32) returns()
func (_CustodianUpgradeable *CustodianUpgradeableTransactor) ConfirmCustodianChange(opts *bind.TransactOpts, _lockId [32]byte) (*types.Transaction, error) {
	return _CustodianUpgradeable.contract.Transact(opts, "confirmCustodianChange", _lockId)
}

// ConfirmCustodianChange is a paid mutator transaction binding the contract method 0x3a8343ee.
//
// Solidity: function confirmCustodianChange(_lockId bytes32) returns()
func (_CustodianUpgradeable *CustodianUpgradeableSession) ConfirmCustodianChange(_lockId [32]byte) (*types.Transaction, error) {
	return _CustodianUpgradeable.Contract.ConfirmCustodianChange(&_CustodianUpgradeable.TransactOpts, _lockId)
}

// ConfirmCustodianChange is a paid mutator transaction binding the contract method 0x3a8343ee.
//
// Solidity: function confirmCustodianChange(_lockId bytes32) returns()
func (_CustodianUpgradeable *CustodianUpgradeableTransactorSession) ConfirmCustodianChange(_lockId [32]byte) (*types.Transaction, error) {
	return _CustodianUpgradeable.Contract.ConfirmCustodianChange(&_CustodianUpgradeable.TransactOpts, _lockId)
}

// RequestCustodianChange is a paid mutator transaction binding the contract method 0x15b21082.
//
// Solidity: function requestCustodianChange(_proposedCustodian address) returns(lockId bytes32)
func (_CustodianUpgradeable *CustodianUpgradeableTransactor) RequestCustodianChange(opts *bind.TransactOpts, _proposedCustodian common.Address) (*types.Transaction, error) {
	return _CustodianUpgradeable.contract.Transact(opts, "requestCustodianChange", _proposedCustodian)
}

// RequestCustodianChange is a paid mutator transaction binding the contract method 0x15b21082.
//
// Solidity: function requestCustodianChange(_proposedCustodian address) returns(lockId bytes32)
func (_CustodianUpgradeable *CustodianUpgradeableSession) RequestCustodianChange(_proposedCustodian common.Address) (*types.Transaction, error) {
	return _CustodianUpgradeable.Contract.RequestCustodianChange(&_CustodianUpgradeable.TransactOpts, _proposedCustodian)
}

// RequestCustodianChange is a paid mutator transaction binding the contract method 0x15b21082.
//
// Solidity: function requestCustodianChange(_proposedCustodian address) returns(lockId bytes32)
func (_CustodianUpgradeable *CustodianUpgradeableTransactorSession) RequestCustodianChange(_proposedCustodian common.Address) (*types.Transaction, error) {
	return _CustodianUpgradeable.Contract.RequestCustodianChange(&_CustodianUpgradeable.TransactOpts, _proposedCustodian)
}

// CustodianUpgradeableCustodianChangeConfirmedIterator is returned from FilterCustodianChangeConfirmed and is used to iterate over the raw logs and unpacked data for CustodianChangeConfirmed events raised by the CustodianUpgradeable contract.
type CustodianUpgradeableCustodianChangeConfirmedIterator struct {
	Event *CustodianUpgradeableCustodianChangeConfirmed // Event containing the contract specifics and raw log

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
func (it *CustodianUpgradeableCustodianChangeConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustodianUpgradeableCustodianChangeConfirmed)
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
		it.Event = new(CustodianUpgradeableCustodianChangeConfirmed)
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
func (it *CustodianUpgradeableCustodianChangeConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CustodianUpgradeableCustodianChangeConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CustodianUpgradeableCustodianChangeConfirmed represents a CustodianChangeConfirmed event raised by the CustodianUpgradeable contract.
type CustodianUpgradeableCustodianChangeConfirmed struct {
	LockId       [32]byte
	NewCustodian common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCustodianChangeConfirmed is a free log retrieval operation binding the contract event 0x9a99272c0f6b7a30ef9e76e684a7cd408bfd4f11a72f36a8e276253c920e442d.
//
// Solidity: e CustodianChangeConfirmed(_lockId bytes32, _newCustodian address)
func (_CustodianUpgradeable *CustodianUpgradeableFilterer) FilterCustodianChangeConfirmed(opts *bind.FilterOpts) (*CustodianUpgradeableCustodianChangeConfirmedIterator, error) {

	logs, sub, err := _CustodianUpgradeable.contract.FilterLogs(opts, "CustodianChangeConfirmed")
	if err != nil {
		return nil, err
	}
	return &CustodianUpgradeableCustodianChangeConfirmedIterator{contract: _CustodianUpgradeable.contract, event: "CustodianChangeConfirmed", logs: logs, sub: sub}, nil
}

// WatchCustodianChangeConfirmed is a free log subscription operation binding the contract event 0x9a99272c0f6b7a30ef9e76e684a7cd408bfd4f11a72f36a8e276253c920e442d.
//
// Solidity: e CustodianChangeConfirmed(_lockId bytes32, _newCustodian address)
func (_CustodianUpgradeable *CustodianUpgradeableFilterer) WatchCustodianChangeConfirmed(opts *bind.WatchOpts, sink chan<- *CustodianUpgradeableCustodianChangeConfirmed) (event.Subscription, error) {

	logs, sub, err := _CustodianUpgradeable.contract.WatchLogs(opts, "CustodianChangeConfirmed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CustodianUpgradeableCustodianChangeConfirmed)
				if err := _CustodianUpgradeable.contract.UnpackLog(event, "CustodianChangeConfirmed", log); err != nil {
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

// CustodianUpgradeableCustodianChangeRequestedIterator is returned from FilterCustodianChangeRequested and is used to iterate over the raw logs and unpacked data for CustodianChangeRequested events raised by the CustodianUpgradeable contract.
type CustodianUpgradeableCustodianChangeRequestedIterator struct {
	Event *CustodianUpgradeableCustodianChangeRequested // Event containing the contract specifics and raw log

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
func (it *CustodianUpgradeableCustodianChangeRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustodianUpgradeableCustodianChangeRequested)
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
		it.Event = new(CustodianUpgradeableCustodianChangeRequested)
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
func (it *CustodianUpgradeableCustodianChangeRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CustodianUpgradeableCustodianChangeRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CustodianUpgradeableCustodianChangeRequested represents a CustodianChangeRequested event raised by the CustodianUpgradeable contract.
type CustodianUpgradeableCustodianChangeRequested struct {
	LockId            [32]byte
	MsgSender         common.Address
	ProposedCustodian common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCustodianChangeRequested is a free log retrieval operation binding the contract event 0xd76fc900a7e1a6fcf11d54b7ba943918df6c53a3128140658c389b3da1e997ba.
//
// Solidity: e CustodianChangeRequested(_lockId bytes32, _msgSender address, _proposedCustodian address)
func (_CustodianUpgradeable *CustodianUpgradeableFilterer) FilterCustodianChangeRequested(opts *bind.FilterOpts) (*CustodianUpgradeableCustodianChangeRequestedIterator, error) {

	logs, sub, err := _CustodianUpgradeable.contract.FilterLogs(opts, "CustodianChangeRequested")
	if err != nil {
		return nil, err
	}
	return &CustodianUpgradeableCustodianChangeRequestedIterator{contract: _CustodianUpgradeable.contract, event: "CustodianChangeRequested", logs: logs, sub: sub}, nil
}

// WatchCustodianChangeRequested is a free log subscription operation binding the contract event 0xd76fc900a7e1a6fcf11d54b7ba943918df6c53a3128140658c389b3da1e997ba.
//
// Solidity: e CustodianChangeRequested(_lockId bytes32, _msgSender address, _proposedCustodian address)
func (_CustodianUpgradeable *CustodianUpgradeableFilterer) WatchCustodianChangeRequested(opts *bind.WatchOpts, sink chan<- *CustodianUpgradeableCustodianChangeRequested) (event.Subscription, error) {

	logs, sub, err := _CustodianUpgradeable.contract.WatchLogs(opts, "CustodianChangeRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CustodianUpgradeableCustodianChangeRequested)
				if err := _CustodianUpgradeable.contract.UnpackLog(event, "CustodianChangeRequested", log); err != nil {
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

// ERC20ImplABI is the input ABI used to generate the binding from.
const ERC20ImplABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"pendingPrintMap\",\"outputs\":[{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proposedCustodian\",\"type\":\"address\"}],\"name\":\"requestCustodianChange\",\"outputs\":[{\"name\":\"lockId\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApprovalWithSender\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"custodian\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_lockId\",\"type\":\"bytes32\"}],\"name\":\"confirmPrint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_lockId\",\"type\":\"bytes32\"}],\"name\":\"confirmCustodianChange\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFromWithSender\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApprovalWithSender\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"erc20Proxy\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tos\",\"type\":\"address[]\"},{\"name\":\"_values\",\"type\":\"uint256[]\"}],\"name\":\"batchTransfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approveWithSender\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sweeper\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_vs\",\"type\":\"uint8[]\"},{\"name\":\"_rs\",\"type\":\"bytes32[]\"},{\"name\":\"_ss\",\"type\":\"bytes32[]\"},{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"enableSweep\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"requestPrint\",\"outputs\":[{\"name\":\"lockId\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lockRequestCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"custodianChangeReqs\",\"outputs\":[{\"name\":\"proposedNew\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sweepMsg\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferWithSender\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"sweptSet\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_froms\",\"type\":\"address[]\"},{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"replaySweep\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"erc20Store\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_erc20Proxy\",\"type\":\"address\"},{\"name\":\"_erc20Store\",\"type\":\"address\"},{\"name\":\"_custodian\",\"type\":\"address\"},{\"name\":\"_sweeper\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_lockId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"PrintingLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_lockId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"PrintingConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_lockId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_msgSender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_proposedCustodian\",\"type\":\"address\"}],\"name\":\"CustodianChangeRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_lockId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_newCustodian\",\"type\":\"address\"}],\"name\":\"CustodianChangeConfirmed\",\"type\":\"event\"}]"

// ERC20ImplBin is the compiled bytecode used for deploying new contracts.
const ERC20ImplBin = `0x608060405234801561001057600080fd5b5060405160808061233683398101604090815281516020830151918301516060909301516000805560018054600160a060020a031916600160a060020a03808716919091179091559193918116151561006857600080fd5b60038054600160a060020a03958616600160a060020a0319918216179091556004805494861694821694909417909355600580549190941692169190911790915550604080516c01000000000000000000000000300281527f737765657000000000000000000000000000000000000000000000000000000060148201529051908190036019019020600655612233806101036000396000f30060806040526004361061012f5763ffffffff60e060020a60003504166304ff7d3f811461013457806315b210821461016f57806318160ddd146101a25780632e0179b5146101b7578063375b74c3146101f5578063380ba30c146102265780633a8343ee1461024057806342966c68146102585780635d5e22cd1461027057806361e1077d146102a057806370a08231146102ca5780637f555b03146102eb57806388d695b21461030057806389064fd21461038e5780639189a59e146103b8578063b7279ca6146103cd578063be23d2911461049f578063cb81fecf146104c3578063cf6e4488146104d8578063dd62ed3e146104f0578063de5007ff14610517578063dfe0f0ca1461052c578063ea5fbfd514610556578063eb55b2a314610577578063f602c312146105d7575b600080fd5b34801561014057600080fd5b5061014c6004356105ec565b60408051600160a060020a03909316835260208301919091528051918290030190f35b34801561017b57600080fd5b50610190600160a060020a0360043516610611565b60408051918252519081900360200190f35b3480156101ae57600080fd5b506101906106c4565b3480156101c357600080fd5b506101e1600160a060020a036004358116906024351660443561074a565b604080519115158252519081900360200190f35b34801561020157600080fd5b5061020a610946565b60408051600160a060020a039092168252519081900360200190f35b34801561023257600080fd5b5061023e600435610955565b005b34801561024c57600080fd5b5061023e600435610c25565b34801561026457600080fd5b506101e1600435610cc6565b34801561027c57600080fd5b506101e1600160a060020a0360043581169060243581169060443516606435610f28565b3480156102ac57600080fd5b506101e1600160a060020a0360043581169060243516604435611286565b3480156102d657600080fd5b50610190600160a060020a0360043516611356565b3480156102f757600080fd5b5061020a6113df565b34801561030c57600080fd5b50604080516020600480358082013583810280860185019096528085526101e195369593946024949385019291829185019084908082843750506040805187358901803560208181028481018201909552818452989b9a9989019892975090820195509350839250850190849080828437509497506113ee9650505050505050565b34801561039a57600080fd5b506101e1600160a060020a0360043581169060243516604435611676565b3480156103c457600080fd5b5061020a6117cf565b3480156103d957600080fd5b506040805160206004803580820135838102808601850190965280855261023e95369593946024949385019291829185019084908082843750506040805187358901803560208181028481018201909552818452989b9a998901989297509082019550935083925085019084908082843750506040805187358901803560208181028481018201909552818452989b9a99890198929750908201955093508392508501908490808284375094975050509235600160a060020a031693506117de92505050565b3480156104ab57600080fd5b50610190600160a060020a0360043516602435611b2d565b3480156104cf57600080fd5b50610190611bef565b3480156104e457600080fd5b5061020a600435611bf5565b3480156104fc57600080fd5b50610190600160a060020a0360043581169060243516611c10565b34801561052357600080fd5b50610190611ca2565b34801561053857600080fd5b506101e1600160a060020a0360043581169060243516604435611ca8565b34801561056257600080fd5b506101e1600160a060020a0360043516611ed6565b34801561058357600080fd5b506040805160206004803580820135838102808601850190965280855261023e9536959394602494938501929182918501908490808284375094975050509235600160a060020a03169350611eeb92505050565b3480156105e357600080fd5b5061020a612183565b60086020526000908152604090208054600190910154600160a060020a039091169082565b6000600160a060020a038216151561062857600080fd5b610630612192565b6040805160208181018352600160a060020a0380871680845260008681526002845285902093518454921673ffffffffffffffffffffffffffffffffffffffff1990921691909117909255825184815233918101919091528083019190915290519192507fd76fc900a7e1a6fcf11d54b7ba943918df6c53a3128140658c389b3da1e997ba919081900360600190a1919050565b6000600460009054906101000a9004600160a060020a0316600160a060020a03166318160ddd6040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561071957600080fd5b505af115801561072d573d6000803e3d6000fd5b505050506040513d602081101561074357600080fd5b5051905090565b60035460009081908190600160a060020a0316331461076857600080fd5b600160a060020a038516151561077d57600080fd5b600480546040805160e060020a635c658165028152600160a060020a038a811694820194909452888416602482015290519290911691635c658165916044808201926020929091908290030181600087803b1580156107db57600080fd5b505af11580156107ef573d6000803e3d6000fd5b505050506040513d602081101561080557600080fd5b50519150508281018181101561081a57600080fd5b60048054604080517fda46098c000000000000000000000000000000000000000000000000000000008152600160a060020a038a8116948201949094528884166024820152604481018590529051929091169163da46098c9160648082019260009290919082900301818387803b15801561089457600080fd5b505af11580156108a8573d6000803e3d6000fd5b5050600354604080517f5687f2b8000000000000000000000000000000000000000000000000000000008152600160a060020a038b811660048301528a81166024830152604482018790529151919092169350635687f2b89250606480830192600092919082900301818387803b15801561092257600080fd5b505af1158015610936573d6000803e3d6000fd5b5060019998505050505050505050565b600154600160a060020a031681565b6001546000908190819081908190600160a060020a0316331461097757600080fd5b60008681526008602052604090208054909550600160a060020a031693508315156109a157600080fd5b6001808601546000888152600860209081526040808320805473ffffffffffffffffffffffffffffffffffffffff191681559094018290556004805485517f18160ddd0000000000000000000000000000000000000000000000000000000081529551949850600160a060020a0316946318160ddd948183019490929183900390910190829087803b158015610a3657600080fd5b505af1158015610a4a573d6000803e3d6000fd5b505050506040513d6020811015610a6057600080fd5b5051915050818101818110610c1d5760048054604080517ff7ea7a3d00000000000000000000000000000000000000000000000000000000815292830184905251600160a060020a039091169163f7ea7a3d91602480830192600092919082900301818387803b158015610ad357600080fd5b505af1158015610ae7573d6000803e3d6000fd5b5050600480546040805160e160020a6310f29c1d028152600160a060020a038a8116948201949094526024810189905290519290911693506321e5383a925060448082019260009290919082900301818387803b158015610b4757600080fd5b505af1158015610b5b573d6000803e3d6000fd5b505060408051898152600160a060020a038816602082015280820187905290517f1445852a2ef41b86fd81f90a02261a68635ceb02cdbc73f9c5f690af8288f3609350908190036060019150a16003546040805160e060020a6323de6651028152600060048201819052600160a060020a03888116602484015260448301889052925192909316926323de66519260648084019382900301818387803b158015610c0457600080fd5b505af1158015610c18573d6000803e3d6000fd5b505050505b505050505050565b600154600160a060020a03163314610c3c57600080fd5b610c45816121d3565b60018054600160a060020a0392831673ffffffffffffffffffffffffffffffffffffffff1991821617825560008481526002602090815260409182902080549093169092559154825185815293169083015280517f9a99272c0f6b7a30ef9e76e684a7cd408bfd4f11a72f36a8e276253c920e442d9281900390910190a150565b600480546040805160e060020a6327e235e30281523393810193909352516000928392600160a060020a0316916327e235e39160248082019260209290919082900301818787803b158015610d1a57600080fd5b505af1158015610d2e573d6000803e3d6000fd5b505050506040513d6020811015610d4457600080fd5b5051905080831115610d5557600080fd5b600480546040805160e260020a6338c110ef0281523393810193909352858403602484015251600160a060020a039091169163e30443bc91604480830192600092919082900301818387803b158015610dad57600080fd5b505af1158015610dc1573d6000803e3d6000fd5b505060048054604080517f18160ddd0000000000000000000000000000000000000000000000000000000081529051600160a060020a03909216945063f7ea7a3d9350879285926318160ddd9280830192602092918290030181600087803b158015610e2c57600080fd5b505af1158015610e40573d6000803e3d6000fd5b505050506040513d6020811015610e5657600080fd5b50516040805160e060020a63ffffffff861602815292909103600483015251602480830192600092919082900301818387803b158015610e9557600080fd5b505af1158015610ea9573d6000803e3d6000fd5b50506003546040805160e060020a6323de6651028152336004820152600060248201819052604482018990529151600160a060020a0390931694506323de665193506064808201939182900301818387803b158015610f0757600080fd5b505af1158015610f1b573d6000803e3d6000fd5b5060019695505050505050565b60035460009081908190600160a060020a03163314610f4657600080fd5b600160a060020a0385161515610f5b57600080fd5b600480546040805160e060020a6327e235e3028152600160a060020a038a811694820194909452905192909116916327e235e3916024808201926020929091908290030181600087803b158015610fb157600080fd5b505af1158015610fc5573d6000803e3d6000fd5b505050506040513d6020811015610fdb57600080fd5b5051915081841115610fec57600080fd5b600480546040805160e060020a635c658165028152600160a060020a038a8116948201949094528a8416602482015290519290911691635c658165916044808201926020929091908290030181600087803b15801561104a57600080fd5b505af115801561105e573d6000803e3d6000fd5b505050506040513d602081101561107457600080fd5b505190508084111561108557600080fd5b600480546040805160e260020a6338c110ef028152600160a060020a038a81169482019490945287860360248201529051929091169163e30443bc9160448082019260009290919082900301818387803b1580156110e257600080fd5b505af11580156110f6573d6000803e3d6000fd5b5050600480546040805160e160020a6310f29c1d028152600160a060020a038b811694820194909452602481018a905290519290911693506321e5383a925060448082019260009290919082900301818387803b15801561115657600080fd5b505af115801561116a573d6000803e3d6000fd5b505060048054604080517fda46098c000000000000000000000000000000000000000000000000000000008152600160a060020a038c8116948201949094528c841660248201528987036044820152905192909116935063da46098c925060648082019260009290919082900301818387803b1580156111e957600080fd5b505af11580156111fd573d6000803e3d6000fd5b50506003546040805160e060020a6323de6651028152600160a060020a038b811660048301528a81166024830152604482018a905291519190921693506323de66519250606480830192600092919082900301818387803b15801561126157600080fd5b505af1158015611275573d6000803e3d6000fd5b5060019a9950505050505050505050565b60035460009081908190600160a060020a031633146112a457600080fd5b600160a060020a03851615156112b957600080fd5b600480546040805160e060020a635c658165028152600160a060020a038a811694820194909452888416602482015290519290911691635c658165916044808201926020929091908290030181600087803b15801561131757600080fd5b505af115801561132b573d6000803e3d6000fd5b505050506040513d602081101561134157600080fd5b50519150508281038181111561081a57600080fd5b600480546040805160e060020a6327e235e3028152600160a060020a03858116948201949094529051600093909216916327e235e39160248082019260209290919082900301818787803b1580156113ad57600080fd5b505af11580156113c1573d6000803e3d6000fd5b505050506040513d60208110156113d757600080fd5b505192915050565b600354600160a060020a031681565b6000806000806000808651885114151561140757600080fd5b8751600480546040805160e060020a6327e235e3028152339381019390935251929750600160a060020a0316916327e235e3916024808201926020929091908290030181600087803b15801561145c57600080fd5b505af1158015611470573d6000803e3d6000fd5b505050506040513d602081101561148657600080fd5b50519350600092505b848310156115f95787838151811015156114a557fe5b602090810290910101519150600160a060020a03821615156114c657600080fd5b86838151811015156114d457fe5b602090810290910101519050808410156114ed57600080fd5b33600160a060020a0383161461157757600480546040805160e160020a6310f29c1d028152600160a060020a0386811694820194909452602481018590529051968490039692909116916321e5383a9160448082019260009290919082900301818387803b15801561155e57600080fd5b505af1158015611572573d6000803e3d6000fd5b505050505b6003546040805160e060020a6323de6651028152336004820152600160a060020a03858116602483015260448201859052915191909216916323de665191606480830192600092919082900301818387803b1580156115d557600080fd5b505af11580156115e9573d6000803e3d6000fd5b50506001909401935061148f9050565b600480546040805160e260020a6338c110ef02815233938101939093526024830187905251600160a060020a039091169163e30443bc91604480830192600092919082900301818387803b15801561165057600080fd5b505af1158015611664573d6000803e3d6000fd5b5060019b9a5050505050505050505050565b600354600090600160a060020a0316331461169057600080fd5b600160a060020a03831615156116a557600080fd5b60048054604080517fda46098c000000000000000000000000000000000000000000000000000000008152600160a060020a03888116948201949094528684166024820152604481018690529051929091169163da46098c9160648082019260009290919082900301818387803b15801561171f57600080fd5b505af1158015611733573d6000803e3d6000fd5b5050600354604080517f5687f2b8000000000000000000000000000000000000000000000000000000008152600160a060020a0389811660048301528881166024830152604482018890529151919092169350635687f2b89250606480830192600092919082900301818387803b1580156117ad57600080fd5b505af11580156117c1573d6000803e3d6000fd5b506001979650505050505050565b600554600160a060020a031681565b6005546000908190819081908190600160a060020a0316331461180057600080fd5b600160a060020a038616151561181557600080fd5b87518951148015611827575086518951145b151561183257600080fd5b8851945060009350600092505b84831015611aa45760016006548a8581518110151561185a57fe5b906020019060200201518a8681518110151561187257fe5b906020019060200201518a8781518110151561188a57fe5b60209081029091018101516040805160008082528185018084529790975260ff9095168582015260608501939093526080840152905160a0808401949293601f19830193908390039091019190865af11580156118eb573d6000803e3d6000fd5b5050604051601f190151925050600160a060020a03821615611a9957600160a060020a038083166000818152600760209081526040808320805460ff1916600117905560048054825160e060020a6327e235e302815291820195909552905193909416936327e235e39360248083019491928390030190829087803b15801561197357600080fd5b505af1158015611987573d6000803e3d6000fd5b505050506040513d602081101561199d57600080fd5b505190506000811115611a9957600480546040805160e260020a6338c110ef028152600160a060020a0386811694820194909452600060248201819052915197850197939092169263e30443bc92604480820193929182900301818387803b158015611a0857600080fd5b505af1158015611a1c573d6000803e3d6000fd5b50506003546040805160e060020a6323de6651028152600160a060020a0387811660048301528b811660248301526044820187905291519190921693506323de66519250606480830192600092919082900301818387803b158015611a8057600080fd5b505af1158015611a94573d6000803e3d6000fd5b505050505b82600101925061183f565b6000841115611b2257600480546040805160e160020a6310f29c1d028152600160a060020a038a81169482019490945260248101889052905192909116916321e5383a9160448082019260009290919082900301818387803b158015611b0957600080fd5b505af1158015611b1d573d6000803e3d6000fd5b505050505b505050505050505050565b6000600160a060020a0383161515611b4457600080fd5b611b4c612192565b604080518082018252600160a060020a03808716808352602080840188815260008781526008835286902094518554941673ffffffffffffffffffffffffffffffffffffffff19909416939093178455915160019093019290925582518481529081019190915280820185905290519192507f021724c3943709b5b29b9cdcfb21e18e7355b036e07d869b4a69bd8a13ec45e8919081900360600190a192915050565b60005481565b600260205260009081526040902054600160a060020a031681565b600480546040805160e060020a635c658165028152600160a060020a03868116948201949094528484166024820152905160009390921691635c6581659160448082019260209290919082900301818787803b158015611c6f57600080fd5b505af1158015611c83573d6000803e3d6000fd5b505050506040513d6020811015611c9957600080fd5b50519392505050565b60065481565b6003546000908190600160a060020a03163314611cc457600080fd5b600160a060020a0384161515611cd957600080fd5b600480546040805160e060020a6327e235e3028152600160a060020a0389811694820194909452905192909116916327e235e3916024808201926020929091908290030181600087803b158015611d2f57600080fd5b505af1158015611d43573d6000803e3d6000fd5b505050506040513d6020811015611d5957600080fd5b5051905080831115611d6a57600080fd5b600480546040805160e260020a6338c110ef028152600160a060020a038981169482019490945286850360248201529051929091169163e30443bc9160448082019260009290919082900301818387803b158015611dc757600080fd5b505af1158015611ddb573d6000803e3d6000fd5b5050600480546040805160e160020a6310f29c1d028152600160a060020a038a8116948201949094526024810189905290519290911693506321e5383a925060448082019260009290919082900301818387803b158015611e3b57600080fd5b505af1158015611e4f573d6000803e3d6000fd5b50506003546040805160e060020a6323de6651028152600160a060020a038a8116600483015289811660248301526044820189905291519190921693506323de66519250606480830192600092919082900301818387803b158015611eb357600080fd5b505af1158015611ec7573d6000803e3d6000fd5b50600198975050505050505050565b60076020526000908152604090205460ff1681565b6005546000908190819081908190600160a060020a03163314611f0d57600080fd5b600160a060020a0386161515611f2257600080fd5b8651945060009350600092505b848310156120fc578683815181101515611f4557fe5b6020908102909101810151600160a060020a0381166000908152600790925260409091205490925060ff16156120f157600480546040805160e060020a6327e235e3028152600160a060020a0386811694820194909452905192909116916327e235e3916024808201926020929091908290030181600087803b158015611fcb57600080fd5b505af1158015611fdf573d6000803e3d6000fd5b505050506040513d6020811015611ff557600080fd5b5051905060008111156120f157600480546040805160e260020a6338c110ef028152600160a060020a0386811694820194909452600060248201819052915197850197939092169263e30443bc92604480820193929182900301818387803b15801561206057600080fd5b505af1158015612074573d6000803e3d6000fd5b50506003546040805160e060020a6323de6651028152600160a060020a0387811660048301528b811660248301526044820187905291519190921693506323de66519250606480830192600092919082900301818387803b1580156120d857600080fd5b505af11580156120ec573d6000803e3d6000fd5b505050505b826001019250611f2f565b600084111561217a57600480546040805160e160020a6310f29c1d028152600160a060020a038a81169482019490945260248101889052905192909116916321e5383a9160448082019260009290919082900301818387803b15801561216157600080fd5b505af1158015612175573d6000803e3d6000fd5b505050505b50505050505050565b600454600160a060020a031681565b60008054600101908190556040805160001943014081526c010000000000000000000000003002602082015260348101929092525190819003605401902090565b60008181526002602052604081208054600160a060020a031615156121f757600080fd5b54600160a060020a0316929150505600a165627a7a723058206676b36cf0afd08e2e513a358032c47fccab787d3bdbb3689b0638a8058a46a40029`

// DeployERC20Impl deploys a new Ethereum contract, binding an instance of ERC20Impl to it.
func DeployERC20Impl(auth *bind.TransactOpts, backend bind.ContractBackend, _erc20Proxy common.Address, _erc20Store common.Address, _custodian common.Address, _sweeper common.Address) (common.Address, *types.Transaction, *ERC20Impl, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ImplABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20ImplBin), backend, _erc20Proxy, _erc20Store, _custodian, _sweeper)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Impl{ERC20ImplCaller: ERC20ImplCaller{contract: contract}, ERC20ImplTransactor: ERC20ImplTransactor{contract: contract}, ERC20ImplFilterer: ERC20ImplFilterer{contract: contract}}, nil
}

// ERC20Impl is an auto generated Go binding around an Ethereum contract.
type ERC20Impl struct {
	ERC20ImplCaller     // Read-only binding to the contract
	ERC20ImplTransactor // Write-only binding to the contract
	ERC20ImplFilterer   // Log filterer for contract events
}

// ERC20ImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20ImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20ImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20ImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20ImplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20ImplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20ImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20ImplSession struct {
	Contract     *ERC20Impl        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20ImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20ImplCallerSession struct {
	Contract *ERC20ImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ERC20ImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20ImplTransactorSession struct {
	Contract     *ERC20ImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ERC20ImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20ImplRaw struct {
	Contract *ERC20Impl // Generic contract binding to access the raw methods on
}

// ERC20ImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20ImplCallerRaw struct {
	Contract *ERC20ImplCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20ImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20ImplTransactorRaw struct {
	Contract *ERC20ImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Impl creates a new instance of ERC20Impl, bound to a specific deployed contract.
func NewERC20Impl(address common.Address, backend bind.ContractBackend) (*ERC20Impl, error) {
	contract, err := bindERC20Impl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Impl{ERC20ImplCaller: ERC20ImplCaller{contract: contract}, ERC20ImplTransactor: ERC20ImplTransactor{contract: contract}, ERC20ImplFilterer: ERC20ImplFilterer{contract: contract}}, nil
}

// NewERC20ImplCaller creates a new read-only instance of ERC20Impl, bound to a specific deployed contract.
func NewERC20ImplCaller(address common.Address, caller bind.ContractCaller) (*ERC20ImplCaller, error) {
	contract, err := bindERC20Impl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20ImplCaller{contract: contract}, nil
}

// NewERC20ImplTransactor creates a new write-only instance of ERC20Impl, bound to a specific deployed contract.
func NewERC20ImplTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20ImplTransactor, error) {
	contract, err := bindERC20Impl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20ImplTransactor{contract: contract}, nil
}

// NewERC20ImplFilterer creates a new log filterer instance of ERC20Impl, bound to a specific deployed contract.
func NewERC20ImplFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20ImplFilterer, error) {
	contract, err := bindERC20Impl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20ImplFilterer{contract: contract}, nil
}

// bindERC20Impl binds a generic wrapper to an already deployed contract.
func bindERC20Impl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ImplABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Impl *ERC20ImplRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Impl.Contract.ERC20ImplCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Impl *ERC20ImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Impl.Contract.ERC20ImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Impl *ERC20ImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Impl.Contract.ERC20ImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Impl *ERC20ImplCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Impl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Impl *ERC20ImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Impl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Impl *ERC20ImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Impl.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_ERC20Impl *ERC20ImplCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Impl.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_ERC20Impl *ERC20ImplSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ERC20Impl.Contract.Allowance(&_ERC20Impl.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_ERC20Impl *ERC20ImplCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ERC20Impl.Contract.Allowance(&_ERC20Impl.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ERC20Impl *ERC20ImplCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Impl.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ERC20Impl *ERC20ImplSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC20Impl.Contract.BalanceOf(&_ERC20Impl.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ERC20Impl *ERC20ImplCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC20Impl.Contract.BalanceOf(&_ERC20Impl.CallOpts, _owner)
}

// Custodian is a free data retrieval call binding the contract method 0x375b74c3.
//
// Solidity: function custodian() constant returns(address)
func (_ERC20Impl *ERC20ImplCaller) Custodian(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC20Impl.contract.Call(opts, out, "custodian")
	return *ret0, err
}

// Custodian is a free data retrieval call binding the contract method 0x375b74c3.
//
// Solidity: function custodian() constant returns(address)
func (_ERC20Impl *ERC20ImplSession) Custodian() (common.Address, error) {
	return _ERC20Impl.Contract.Custodian(&_ERC20Impl.CallOpts)
}

// Custodian is a free data retrieval call binding the contract method 0x375b74c3.
//
// Solidity: function custodian() constant returns(address)
func (_ERC20Impl *ERC20ImplCallerSession) Custodian() (common.Address, error) {
	return _ERC20Impl.Contract.Custodian(&_ERC20Impl.CallOpts)
}

// CustodianChangeReqs is a free data retrieval call binding the contract method 0xcf6e4488.
//
// Solidity: function custodianChangeReqs( bytes32) constant returns(proposedNew address)
func (_ERC20Impl *ERC20ImplCaller) CustodianChangeReqs(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC20Impl.contract.Call(opts, out, "custodianChangeReqs", arg0)
	return *ret0, err
}

// CustodianChangeReqs is a free data retrieval call binding the contract method 0xcf6e4488.
//
// Solidity: function custodianChangeReqs( bytes32) constant returns(proposedNew address)
func (_ERC20Impl *ERC20ImplSession) CustodianChangeReqs(arg0 [32]byte) (common.Address, error) {
	return _ERC20Impl.Contract.CustodianChangeReqs(&_ERC20Impl.CallOpts, arg0)
}

// CustodianChangeReqs is a free data retrieval call binding the contract method 0xcf6e4488.
//
// Solidity: function custodianChangeReqs( bytes32) constant returns(proposedNew address)
func (_ERC20Impl *ERC20ImplCallerSession) CustodianChangeReqs(arg0 [32]byte) (common.Address, error) {
	return _ERC20Impl.Contract.CustodianChangeReqs(&_ERC20Impl.CallOpts, arg0)
}

// Erc20Proxy is a free data retrieval call binding the contract method 0x7f555b03.
//
// Solidity: function erc20Proxy() constant returns(address)
func (_ERC20Impl *ERC20ImplCaller) Erc20Proxy(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC20Impl.contract.Call(opts, out, "erc20Proxy")
	return *ret0, err
}

// Erc20Proxy is a free data retrieval call binding the contract method 0x7f555b03.
//
// Solidity: function erc20Proxy() constant returns(address)
func (_ERC20Impl *ERC20ImplSession) Erc20Proxy() (common.Address, error) {
	return _ERC20Impl.Contract.Erc20Proxy(&_ERC20Impl.CallOpts)
}

// Erc20Proxy is a free data retrieval call binding the contract method 0x7f555b03.
//
// Solidity: function erc20Proxy() constant returns(address)
func (_ERC20Impl *ERC20ImplCallerSession) Erc20Proxy() (common.Address, error) {
	return _ERC20Impl.Contract.Erc20Proxy(&_ERC20Impl.CallOpts)
}

// Erc20Store is a free data retrieval call binding the contract method 0xf602c312.
//
// Solidity: function erc20Store() constant returns(address)
func (_ERC20Impl *ERC20ImplCaller) Erc20Store(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC20Impl.contract.Call(opts, out, "erc20Store")
	return *ret0, err
}

// Erc20Store is a free data retrieval call binding the contract method 0xf602c312.
//
// Solidity: function erc20Store() constant returns(address)
func (_ERC20Impl *ERC20ImplSession) Erc20Store() (common.Address, error) {
	return _ERC20Impl.Contract.Erc20Store(&_ERC20Impl.CallOpts)
}

// Erc20Store is a free data retrieval call binding the contract method 0xf602c312.
//
// Solidity: function erc20Store() constant returns(address)
func (_ERC20Impl *ERC20ImplCallerSession) Erc20Store() (common.Address, error) {
	return _ERC20Impl.Contract.Erc20Store(&_ERC20Impl.CallOpts)
}

// LockRequestCount is a free data retrieval call binding the contract method 0xcb81fecf.
//
// Solidity: function lockRequestCount() constant returns(uint256)
func (_ERC20Impl *ERC20ImplCaller) LockRequestCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Impl.contract.Call(opts, out, "lockRequestCount")
	return *ret0, err
}

// LockRequestCount is a free data retrieval call binding the contract method 0xcb81fecf.
//
// Solidity: function lockRequestCount() constant returns(uint256)
func (_ERC20Impl *ERC20ImplSession) LockRequestCount() (*big.Int, error) {
	return _ERC20Impl.Contract.LockRequestCount(&_ERC20Impl.CallOpts)
}

// LockRequestCount is a free data retrieval call binding the contract method 0xcb81fecf.
//
// Solidity: function lockRequestCount() constant returns(uint256)
func (_ERC20Impl *ERC20ImplCallerSession) LockRequestCount() (*big.Int, error) {
	return _ERC20Impl.Contract.LockRequestCount(&_ERC20Impl.CallOpts)
}

// PendingPrintMap is a free data retrieval call binding the contract method 0x04ff7d3f.
//
// Solidity: function pendingPrintMap( bytes32) constant returns(receiver address, value uint256)
func (_ERC20Impl *ERC20ImplCaller) PendingPrintMap(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Receiver common.Address
	Value    *big.Int
}, error) {
	ret := new(struct {
		Receiver common.Address
		Value    *big.Int
	})
	out := ret
	err := _ERC20Impl.contract.Call(opts, out, "pendingPrintMap", arg0)
	return *ret, err
}

// PendingPrintMap is a free data retrieval call binding the contract method 0x04ff7d3f.
//
// Solidity: function pendingPrintMap( bytes32) constant returns(receiver address, value uint256)
func (_ERC20Impl *ERC20ImplSession) PendingPrintMap(arg0 [32]byte) (struct {
	Receiver common.Address
	Value    *big.Int
}, error) {
	return _ERC20Impl.Contract.PendingPrintMap(&_ERC20Impl.CallOpts, arg0)
}

// PendingPrintMap is a free data retrieval call binding the contract method 0x04ff7d3f.
//
// Solidity: function pendingPrintMap( bytes32) constant returns(receiver address, value uint256)
func (_ERC20Impl *ERC20ImplCallerSession) PendingPrintMap(arg0 [32]byte) (struct {
	Receiver common.Address
	Value    *big.Int
}, error) {
	return _ERC20Impl.Contract.PendingPrintMap(&_ERC20Impl.CallOpts, arg0)
}

// SweepMsg is a free data retrieval call binding the contract method 0xde5007ff.
//
// Solidity: function sweepMsg() constant returns(bytes32)
func (_ERC20Impl *ERC20ImplCaller) SweepMsg(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ERC20Impl.contract.Call(opts, out, "sweepMsg")
	return *ret0, err
}

// SweepMsg is a free data retrieval call binding the contract method 0xde5007ff.
//
// Solidity: function sweepMsg() constant returns(bytes32)
func (_ERC20Impl *ERC20ImplSession) SweepMsg() ([32]byte, error) {
	return _ERC20Impl.Contract.SweepMsg(&_ERC20Impl.CallOpts)
}

// SweepMsg is a free data retrieval call binding the contract method 0xde5007ff.
//
// Solidity: function sweepMsg() constant returns(bytes32)
func (_ERC20Impl *ERC20ImplCallerSession) SweepMsg() ([32]byte, error) {
	return _ERC20Impl.Contract.SweepMsg(&_ERC20Impl.CallOpts)
}

// Sweeper is a free data retrieval call binding the contract method 0x9189a59e.
//
// Solidity: function sweeper() constant returns(address)
func (_ERC20Impl *ERC20ImplCaller) Sweeper(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC20Impl.contract.Call(opts, out, "sweeper")
	return *ret0, err
}

// Sweeper is a free data retrieval call binding the contract method 0x9189a59e.
//
// Solidity: function sweeper() constant returns(address)
func (_ERC20Impl *ERC20ImplSession) Sweeper() (common.Address, error) {
	return _ERC20Impl.Contract.Sweeper(&_ERC20Impl.CallOpts)
}

// Sweeper is a free data retrieval call binding the contract method 0x9189a59e.
//
// Solidity: function sweeper() constant returns(address)
func (_ERC20Impl *ERC20ImplCallerSession) Sweeper() (common.Address, error) {
	return _ERC20Impl.Contract.Sweeper(&_ERC20Impl.CallOpts)
}

// SweptSet is a free data retrieval call binding the contract method 0xea5fbfd5.
//
// Solidity: function sweptSet( address) constant returns(bool)
func (_ERC20Impl *ERC20ImplCaller) SweptSet(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC20Impl.contract.Call(opts, out, "sweptSet", arg0)
	return *ret0, err
}

// SweptSet is a free data retrieval call binding the contract method 0xea5fbfd5.
//
// Solidity: function sweptSet( address) constant returns(bool)
func (_ERC20Impl *ERC20ImplSession) SweptSet(arg0 common.Address) (bool, error) {
	return _ERC20Impl.Contract.SweptSet(&_ERC20Impl.CallOpts, arg0)
}

// SweptSet is a free data retrieval call binding the contract method 0xea5fbfd5.
//
// Solidity: function sweptSet( address) constant returns(bool)
func (_ERC20Impl *ERC20ImplCallerSession) SweptSet(arg0 common.Address) (bool, error) {
	return _ERC20Impl.Contract.SweptSet(&_ERC20Impl.CallOpts, arg0)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Impl *ERC20ImplCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Impl.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Impl *ERC20ImplSession) TotalSupply() (*big.Int, error) {
	return _ERC20Impl.Contract.TotalSupply(&_ERC20Impl.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Impl *ERC20ImplCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Impl.Contract.TotalSupply(&_ERC20Impl.CallOpts)
}

// ApproveWithSender is a paid mutator transaction binding the contract method 0x89064fd2.
//
// Solidity: function approveWithSender(_sender address, _spender address, _value uint256) returns(success bool)
func (_ERC20Impl *ERC20ImplTransactor) ApproveWithSender(opts *bind.TransactOpts, _sender common.Address, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Impl.contract.Transact(opts, "approveWithSender", _sender, _spender, _value)
}

// ApproveWithSender is a paid mutator transaction binding the contract method 0x89064fd2.
//
// Solidity: function approveWithSender(_sender address, _spender address, _value uint256) returns(success bool)
func (_ERC20Impl *ERC20ImplSession) ApproveWithSender(_sender common.Address, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Impl.Contract.ApproveWithSender(&_ERC20Impl.TransactOpts, _sender, _spender, _value)
}

// ApproveWithSender is a paid mutator transaction binding the contract method 0x89064fd2.
//
// Solidity: function approveWithSender(_sender address, _spender address, _value uint256) returns(success bool)
func (_ERC20Impl *ERC20ImplTransactorSession) ApproveWithSender(_sender common.Address, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Impl.Contract.ApproveWithSender(&_ERC20Impl.TransactOpts, _sender, _spender, _value)
}

// BatchTransfer is a paid mutator transaction binding the contract method 0x88d695b2.
//
// Solidity: function batchTransfer(_tos address[], _values uint256[]) returns(success bool)
func (_ERC20Impl *ERC20ImplTransactor) BatchTransfer(opts *bind.TransactOpts, _tos []common.Address, _values []*big.Int) (*types.Transaction, error) {
	return _ERC20Impl.contract.Transact(opts, "batchTransfer", _tos, _values)
}

// BatchTransfer is a paid mutator transaction binding the contract method 0x88d695b2.
//
// Solidity: function batchTransfer(_tos address[], _values uint256[]) returns(success bool)
func (_ERC20Impl *ERC20ImplSession) BatchTransfer(_tos []common.Address, _values []*big.Int) (*types.Transaction, error) {
	return _ERC20Impl.Contract.BatchTransfer(&_ERC20Impl.TransactOpts, _tos, _values)
}

// BatchTransfer is a paid mutator transaction binding the contract method 0x88d695b2.
//
// Solidity: function batchTransfer(_tos address[], _values uint256[]) returns(success bool)
func (_ERC20Impl *ERC20ImplTransactorSession) BatchTransfer(_tos []common.Address, _values []*big.Int) (*types.Transaction, error) {
	return _ERC20Impl.Contract.BatchTransfer(&_ERC20Impl.TransactOpts, _tos, _values)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns(success bool)
func (_ERC20Impl *ERC20ImplTransactor) Burn(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Impl.contract.Transact(opts, "burn", _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns(success bool)
func (_ERC20Impl *ERC20ImplSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _ERC20Impl.Contract.Burn(&_ERC20Impl.TransactOpts, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns(success bool)
func (_ERC20Impl *ERC20ImplTransactorSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _ERC20Impl.Contract.Burn(&_ERC20Impl.TransactOpts, _value)
}

// ConfirmCustodianChange is a paid mutator transaction binding the contract method 0x3a8343ee.
//
// Solidity: function confirmCustodianChange(_lockId bytes32) returns()
func (_ERC20Impl *ERC20ImplTransactor) ConfirmCustodianChange(opts *bind.TransactOpts, _lockId [32]byte) (*types.Transaction, error) {
	return _ERC20Impl.contract.Transact(opts, "confirmCustodianChange", _lockId)
}

// ConfirmCustodianChange is a paid mutator transaction binding the contract method 0x3a8343ee.
//
// Solidity: function confirmCustodianChange(_lockId bytes32) returns()
func (_ERC20Impl *ERC20ImplSession) ConfirmCustodianChange(_lockId [32]byte) (*types.Transaction, error) {
	return _ERC20Impl.Contract.ConfirmCustodianChange(&_ERC20Impl.TransactOpts, _lockId)
}

// ConfirmCustodianChange is a paid mutator transaction binding the contract method 0x3a8343ee.
//
// Solidity: function confirmCustodianChange(_lockId bytes32) returns()
func (_ERC20Impl *ERC20ImplTransactorSession) ConfirmCustodianChange(_lockId [32]byte) (*types.Transaction, error) {
	return _ERC20Impl.Contract.ConfirmCustodianChange(&_ERC20Impl.TransactOpts, _lockId)
}

// ConfirmPrint is a paid mutator transaction binding the contract method 0x380ba30c.
//
// Solidity: function confirmPrint(_lockId bytes32) returns()
func (_ERC20Impl *ERC20ImplTransactor) ConfirmPrint(opts *bind.TransactOpts, _lockId [32]byte) (*types.Transaction, error) {
	return _ERC20Impl.contract.Transact(opts, "confirmPrint", _lockId)
}

// ConfirmPrint is a paid mutator transaction binding the contract method 0x380ba30c.
//
// Solidity: function confirmPrint(_lockId bytes32) returns()
func (_ERC20Impl *ERC20ImplSession) ConfirmPrint(_lockId [32]byte) (*types.Transaction, error) {
	return _ERC20Impl.Contract.ConfirmPrint(&_ERC20Impl.TransactOpts, _lockId)
}

// ConfirmPrint is a paid mutator transaction binding the contract method 0x380ba30c.
//
// Solidity: function confirmPrint(_lockId bytes32) returns()
func (_ERC20Impl *ERC20ImplTransactorSession) ConfirmPrint(_lockId [32]byte) (*types.Transaction, error) {
	return _ERC20Impl.Contract.ConfirmPrint(&_ERC20Impl.TransactOpts, _lockId)
}

// DecreaseApprovalWithSender is a paid mutator transaction binding the contract method 0x61e1077d.
//
// Solidity: function decreaseApprovalWithSender(_sender address, _spender address, _subtractedValue uint256) returns(success bool)
func (_ERC20Impl *ERC20ImplTransactor) DecreaseApprovalWithSender(opts *bind.TransactOpts, _sender common.Address, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Impl.contract.Transact(opts, "decreaseApprovalWithSender", _sender, _spender, _subtractedValue)
}

// DecreaseApprovalWithSender is a paid mutator transaction binding the contract method 0x61e1077d.
//
// Solidity: function decreaseApprovalWithSender(_sender address, _spender address, _subtractedValue uint256) returns(success bool)
func (_ERC20Impl *ERC20ImplSession) DecreaseApprovalWithSender(_sender common.Address, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Impl.Contract.DecreaseApprovalWithSender(&_ERC20Impl.TransactOpts, _sender, _spender, _subtractedValue)
}

// DecreaseApprovalWithSender is a paid mutator transaction binding the contract method 0x61e1077d.
//
// Solidity: function decreaseApprovalWithSender(_sender address, _spender address, _subtractedValue uint256) returns(success bool)
func (_ERC20Impl *ERC20ImplTransactorSession) DecreaseApprovalWithSender(_sender common.Address, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Impl.Contract.DecreaseApprovalWithSender(&_ERC20Impl.TransactOpts, _sender, _spender, _subtractedValue)
}

// EnableSweep is a paid mutator transaction binding the contract method 0xb7279ca6.
//
// Solidity: function enableSweep(_vs uint8[], _rs bytes32[], _ss bytes32[], _to address) returns()
func (_ERC20Impl *ERC20ImplTransactor) EnableSweep(opts *bind.TransactOpts, _vs []uint8, _rs [][32]byte, _ss [][32]byte, _to common.Address) (*types.Transaction, error) {
	return _ERC20Impl.contract.Transact(opts, "enableSweep", _vs, _rs, _ss, _to)
}

// EnableSweep is a paid mutator transaction binding the contract method 0xb7279ca6.
//
// Solidity: function enableSweep(_vs uint8[], _rs bytes32[], _ss bytes32[], _to address) returns()
func (_ERC20Impl *ERC20ImplSession) EnableSweep(_vs []uint8, _rs [][32]byte, _ss [][32]byte, _to common.Address) (*types.Transaction, error) {
	return _ERC20Impl.Contract.EnableSweep(&_ERC20Impl.TransactOpts, _vs, _rs, _ss, _to)
}

// EnableSweep is a paid mutator transaction binding the contract method 0xb7279ca6.
//
// Solidity: function enableSweep(_vs uint8[], _rs bytes32[], _ss bytes32[], _to address) returns()
func (_ERC20Impl *ERC20ImplTransactorSession) EnableSweep(_vs []uint8, _rs [][32]byte, _ss [][32]byte, _to common.Address) (*types.Transaction, error) {
	return _ERC20Impl.Contract.EnableSweep(&_ERC20Impl.TransactOpts, _vs, _rs, _ss, _to)
}

// IncreaseApprovalWithSender is a paid mutator transaction binding the contract method 0x2e0179b5.
//
// Solidity: function increaseApprovalWithSender(_sender address, _spender address, _addedValue uint256) returns(success bool)
func (_ERC20Impl *ERC20ImplTransactor) IncreaseApprovalWithSender(opts *bind.TransactOpts, _sender common.Address, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Impl.contract.Transact(opts, "increaseApprovalWithSender", _sender, _spender, _addedValue)
}

// IncreaseApprovalWithSender is a paid mutator transaction binding the contract method 0x2e0179b5.
//
// Solidity: function increaseApprovalWithSender(_sender address, _spender address, _addedValue uint256) returns(success bool)
func (_ERC20Impl *ERC20ImplSession) IncreaseApprovalWithSender(_sender common.Address, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Impl.Contract.IncreaseApprovalWithSender(&_ERC20Impl.TransactOpts, _sender, _spender, _addedValue)
}

// IncreaseApprovalWithSender is a paid mutator transaction binding the contract method 0x2e0179b5.
//
// Solidity: function increaseApprovalWithSender(_sender address, _spender address, _addedValue uint256) returns(success bool)
func (_ERC20Impl *ERC20ImplTransactorSession) IncreaseApprovalWithSender(_sender common.Address, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Impl.Contract.IncreaseApprovalWithSender(&_ERC20Impl.TransactOpts, _sender, _spender, _addedValue)
}

// ReplaySweep is a paid mutator transaction binding the contract method 0xeb55b2a3.
//
// Solidity: function replaySweep(_froms address[], _to address) returns()
func (_ERC20Impl *ERC20ImplTransactor) ReplaySweep(opts *bind.TransactOpts, _froms []common.Address, _to common.Address) (*types.Transaction, error) {
	return _ERC20Impl.contract.Transact(opts, "replaySweep", _froms, _to)
}

// ReplaySweep is a paid mutator transaction binding the contract method 0xeb55b2a3.
//
// Solidity: function replaySweep(_froms address[], _to address) returns()
func (_ERC20Impl *ERC20ImplSession) ReplaySweep(_froms []common.Address, _to common.Address) (*types.Transaction, error) {
	return _ERC20Impl.Contract.ReplaySweep(&_ERC20Impl.TransactOpts, _froms, _to)
}

// ReplaySweep is a paid mutator transaction binding the contract method 0xeb55b2a3.
//
// Solidity: function replaySweep(_froms address[], _to address) returns()
func (_ERC20Impl *ERC20ImplTransactorSession) ReplaySweep(_froms []common.Address, _to common.Address) (*types.Transaction, error) {
	return _ERC20Impl.Contract.ReplaySweep(&_ERC20Impl.TransactOpts, _froms, _to)
}

// RequestCustodianChange is a paid mutator transaction binding the contract method 0x15b21082.
//
// Solidity: function requestCustodianChange(_proposedCustodian address) returns(lockId bytes32)
func (_ERC20Impl *ERC20ImplTransactor) RequestCustodianChange(opts *bind.TransactOpts, _proposedCustodian common.Address) (*types.Transaction, error) {
	return _ERC20Impl.contract.Transact(opts, "requestCustodianChange", _proposedCustodian)
}

// RequestCustodianChange is a paid mutator transaction binding the contract method 0x15b21082.
//
// Solidity: function requestCustodianChange(_proposedCustodian address) returns(lockId bytes32)
func (_ERC20Impl *ERC20ImplSession) RequestCustodianChange(_proposedCustodian common.Address) (*types.Transaction, error) {
	return _ERC20Impl.Contract.RequestCustodianChange(&_ERC20Impl.TransactOpts, _proposedCustodian)
}

// RequestCustodianChange is a paid mutator transaction binding the contract method 0x15b21082.
//
// Solidity: function requestCustodianChange(_proposedCustodian address) returns(lockId bytes32)
func (_ERC20Impl *ERC20ImplTransactorSession) RequestCustodianChange(_proposedCustodian common.Address) (*types.Transaction, error) {
	return _ERC20Impl.Contract.RequestCustodianChange(&_ERC20Impl.TransactOpts, _proposedCustodian)
}

// RequestPrint is a paid mutator transaction binding the contract method 0xbe23d291.
//
// Solidity: function requestPrint(_receiver address, _value uint256) returns(lockId bytes32)
func (_ERC20Impl *ERC20ImplTransactor) RequestPrint(opts *bind.TransactOpts, _receiver common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Impl.contract.Transact(opts, "requestPrint", _receiver, _value)
}

// RequestPrint is a paid mutator transaction binding the contract method 0xbe23d291.
//
// Solidity: function requestPrint(_receiver address, _value uint256) returns(lockId bytes32)
func (_ERC20Impl *ERC20ImplSession) RequestPrint(_receiver common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Impl.Contract.RequestPrint(&_ERC20Impl.TransactOpts, _receiver, _value)
}

// RequestPrint is a paid mutator transaction binding the contract method 0xbe23d291.
//
// Solidity: function requestPrint(_receiver address, _value uint256) returns(lockId bytes32)
func (_ERC20Impl *ERC20ImplTransactorSession) RequestPrint(_receiver common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Impl.Contract.RequestPrint(&_ERC20Impl.TransactOpts, _receiver, _value)
}

// TransferFromWithSender is a paid mutator transaction binding the contract method 0x5d5e22cd.
//
// Solidity: function transferFromWithSender(_sender address, _from address, _to address, _value uint256) returns(success bool)
func (_ERC20Impl *ERC20ImplTransactor) TransferFromWithSender(opts *bind.TransactOpts, _sender common.Address, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Impl.contract.Transact(opts, "transferFromWithSender", _sender, _from, _to, _value)
}

// TransferFromWithSender is a paid mutator transaction binding the contract method 0x5d5e22cd.
//
// Solidity: function transferFromWithSender(_sender address, _from address, _to address, _value uint256) returns(success bool)
func (_ERC20Impl *ERC20ImplSession) TransferFromWithSender(_sender common.Address, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Impl.Contract.TransferFromWithSender(&_ERC20Impl.TransactOpts, _sender, _from, _to, _value)
}

// TransferFromWithSender is a paid mutator transaction binding the contract method 0x5d5e22cd.
//
// Solidity: function transferFromWithSender(_sender address, _from address, _to address, _value uint256) returns(success bool)
func (_ERC20Impl *ERC20ImplTransactorSession) TransferFromWithSender(_sender common.Address, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Impl.Contract.TransferFromWithSender(&_ERC20Impl.TransactOpts, _sender, _from, _to, _value)
}

// TransferWithSender is a paid mutator transaction binding the contract method 0xdfe0f0ca.
//
// Solidity: function transferWithSender(_sender address, _to address, _value uint256) returns(success bool)
func (_ERC20Impl *ERC20ImplTransactor) TransferWithSender(opts *bind.TransactOpts, _sender common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Impl.contract.Transact(opts, "transferWithSender", _sender, _to, _value)
}

// TransferWithSender is a paid mutator transaction binding the contract method 0xdfe0f0ca.
//
// Solidity: function transferWithSender(_sender address, _to address, _value uint256) returns(success bool)
func (_ERC20Impl *ERC20ImplSession) TransferWithSender(_sender common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Impl.Contract.TransferWithSender(&_ERC20Impl.TransactOpts, _sender, _to, _value)
}

// TransferWithSender is a paid mutator transaction binding the contract method 0xdfe0f0ca.
//
// Solidity: function transferWithSender(_sender address, _to address, _value uint256) returns(success bool)
func (_ERC20Impl *ERC20ImplTransactorSession) TransferWithSender(_sender common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Impl.Contract.TransferWithSender(&_ERC20Impl.TransactOpts, _sender, _to, _value)
}

// ERC20ImplCustodianChangeConfirmedIterator is returned from FilterCustodianChangeConfirmed and is used to iterate over the raw logs and unpacked data for CustodianChangeConfirmed events raised by the ERC20Impl contract.
type ERC20ImplCustodianChangeConfirmedIterator struct {
	Event *ERC20ImplCustodianChangeConfirmed // Event containing the contract specifics and raw log

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
func (it *ERC20ImplCustodianChangeConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20ImplCustodianChangeConfirmed)
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
		it.Event = new(ERC20ImplCustodianChangeConfirmed)
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
func (it *ERC20ImplCustodianChangeConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ImplCustodianChangeConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20ImplCustodianChangeConfirmed represents a CustodianChangeConfirmed event raised by the ERC20Impl contract.
type ERC20ImplCustodianChangeConfirmed struct {
	LockId       [32]byte
	NewCustodian common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCustodianChangeConfirmed is a free log retrieval operation binding the contract event 0x9a99272c0f6b7a30ef9e76e684a7cd408bfd4f11a72f36a8e276253c920e442d.
//
// Solidity: e CustodianChangeConfirmed(_lockId bytes32, _newCustodian address)
func (_ERC20Impl *ERC20ImplFilterer) FilterCustodianChangeConfirmed(opts *bind.FilterOpts) (*ERC20ImplCustodianChangeConfirmedIterator, error) {

	logs, sub, err := _ERC20Impl.contract.FilterLogs(opts, "CustodianChangeConfirmed")
	if err != nil {
		return nil, err
	}
	return &ERC20ImplCustodianChangeConfirmedIterator{contract: _ERC20Impl.contract, event: "CustodianChangeConfirmed", logs: logs, sub: sub}, nil
}

// WatchCustodianChangeConfirmed is a free log subscription operation binding the contract event 0x9a99272c0f6b7a30ef9e76e684a7cd408bfd4f11a72f36a8e276253c920e442d.
//
// Solidity: e CustodianChangeConfirmed(_lockId bytes32, _newCustodian address)
func (_ERC20Impl *ERC20ImplFilterer) WatchCustodianChangeConfirmed(opts *bind.WatchOpts, sink chan<- *ERC20ImplCustodianChangeConfirmed) (event.Subscription, error) {

	logs, sub, err := _ERC20Impl.contract.WatchLogs(opts, "CustodianChangeConfirmed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20ImplCustodianChangeConfirmed)
				if err := _ERC20Impl.contract.UnpackLog(event, "CustodianChangeConfirmed", log); err != nil {
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

// ERC20ImplCustodianChangeRequestedIterator is returned from FilterCustodianChangeRequested and is used to iterate over the raw logs and unpacked data for CustodianChangeRequested events raised by the ERC20Impl contract.
type ERC20ImplCustodianChangeRequestedIterator struct {
	Event *ERC20ImplCustodianChangeRequested // Event containing the contract specifics and raw log

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
func (it *ERC20ImplCustodianChangeRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20ImplCustodianChangeRequested)
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
		it.Event = new(ERC20ImplCustodianChangeRequested)
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
func (it *ERC20ImplCustodianChangeRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ImplCustodianChangeRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20ImplCustodianChangeRequested represents a CustodianChangeRequested event raised by the ERC20Impl contract.
type ERC20ImplCustodianChangeRequested struct {
	LockId            [32]byte
	MsgSender         common.Address
	ProposedCustodian common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCustodianChangeRequested is a free log retrieval operation binding the contract event 0xd76fc900a7e1a6fcf11d54b7ba943918df6c53a3128140658c389b3da1e997ba.
//
// Solidity: e CustodianChangeRequested(_lockId bytes32, _msgSender address, _proposedCustodian address)
func (_ERC20Impl *ERC20ImplFilterer) FilterCustodianChangeRequested(opts *bind.FilterOpts) (*ERC20ImplCustodianChangeRequestedIterator, error) {

	logs, sub, err := _ERC20Impl.contract.FilterLogs(opts, "CustodianChangeRequested")
	if err != nil {
		return nil, err
	}
	return &ERC20ImplCustodianChangeRequestedIterator{contract: _ERC20Impl.contract, event: "CustodianChangeRequested", logs: logs, sub: sub}, nil
}

// WatchCustodianChangeRequested is a free log subscription operation binding the contract event 0xd76fc900a7e1a6fcf11d54b7ba943918df6c53a3128140658c389b3da1e997ba.
//
// Solidity: e CustodianChangeRequested(_lockId bytes32, _msgSender address, _proposedCustodian address)
func (_ERC20Impl *ERC20ImplFilterer) WatchCustodianChangeRequested(opts *bind.WatchOpts, sink chan<- *ERC20ImplCustodianChangeRequested) (event.Subscription, error) {

	logs, sub, err := _ERC20Impl.contract.WatchLogs(opts, "CustodianChangeRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20ImplCustodianChangeRequested)
				if err := _ERC20Impl.contract.UnpackLog(event, "CustodianChangeRequested", log); err != nil {
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

// ERC20ImplPrintingConfirmedIterator is returned from FilterPrintingConfirmed and is used to iterate over the raw logs and unpacked data for PrintingConfirmed events raised by the ERC20Impl contract.
type ERC20ImplPrintingConfirmedIterator struct {
	Event *ERC20ImplPrintingConfirmed // Event containing the contract specifics and raw log

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
func (it *ERC20ImplPrintingConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20ImplPrintingConfirmed)
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
		it.Event = new(ERC20ImplPrintingConfirmed)
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
func (it *ERC20ImplPrintingConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ImplPrintingConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20ImplPrintingConfirmed represents a PrintingConfirmed event raised by the ERC20Impl contract.
type ERC20ImplPrintingConfirmed struct {
	LockId   [32]byte
	Receiver common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPrintingConfirmed is a free log retrieval operation binding the contract event 0x1445852a2ef41b86fd81f90a02261a68635ceb02cdbc73f9c5f690af8288f360.
//
// Solidity: e PrintingConfirmed(_lockId bytes32, _receiver address, _value uint256)
func (_ERC20Impl *ERC20ImplFilterer) FilterPrintingConfirmed(opts *bind.FilterOpts) (*ERC20ImplPrintingConfirmedIterator, error) {

	logs, sub, err := _ERC20Impl.contract.FilterLogs(opts, "PrintingConfirmed")
	if err != nil {
		return nil, err
	}
	return &ERC20ImplPrintingConfirmedIterator{contract: _ERC20Impl.contract, event: "PrintingConfirmed", logs: logs, sub: sub}, nil
}

// WatchPrintingConfirmed is a free log subscription operation binding the contract event 0x1445852a2ef41b86fd81f90a02261a68635ceb02cdbc73f9c5f690af8288f360.
//
// Solidity: e PrintingConfirmed(_lockId bytes32, _receiver address, _value uint256)
func (_ERC20Impl *ERC20ImplFilterer) WatchPrintingConfirmed(opts *bind.WatchOpts, sink chan<- *ERC20ImplPrintingConfirmed) (event.Subscription, error) {

	logs, sub, err := _ERC20Impl.contract.WatchLogs(opts, "PrintingConfirmed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20ImplPrintingConfirmed)
				if err := _ERC20Impl.contract.UnpackLog(event, "PrintingConfirmed", log); err != nil {
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

// ERC20ImplPrintingLockedIterator is returned from FilterPrintingLocked and is used to iterate over the raw logs and unpacked data for PrintingLocked events raised by the ERC20Impl contract.
type ERC20ImplPrintingLockedIterator struct {
	Event *ERC20ImplPrintingLocked // Event containing the contract specifics and raw log

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
func (it *ERC20ImplPrintingLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20ImplPrintingLocked)
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
		it.Event = new(ERC20ImplPrintingLocked)
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
func (it *ERC20ImplPrintingLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ImplPrintingLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20ImplPrintingLocked represents a PrintingLocked event raised by the ERC20Impl contract.
type ERC20ImplPrintingLocked struct {
	LockId   [32]byte
	Receiver common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPrintingLocked is a free log retrieval operation binding the contract event 0x021724c3943709b5b29b9cdcfb21e18e7355b036e07d869b4a69bd8a13ec45e8.
//
// Solidity: e PrintingLocked(_lockId bytes32, _receiver address, _value uint256)
func (_ERC20Impl *ERC20ImplFilterer) FilterPrintingLocked(opts *bind.FilterOpts) (*ERC20ImplPrintingLockedIterator, error) {

	logs, sub, err := _ERC20Impl.contract.FilterLogs(opts, "PrintingLocked")
	if err != nil {
		return nil, err
	}
	return &ERC20ImplPrintingLockedIterator{contract: _ERC20Impl.contract, event: "PrintingLocked", logs: logs, sub: sub}, nil
}

// WatchPrintingLocked is a free log subscription operation binding the contract event 0x021724c3943709b5b29b9cdcfb21e18e7355b036e07d869b4a69bd8a13ec45e8.
//
// Solidity: e PrintingLocked(_lockId bytes32, _receiver address, _value uint256)
func (_ERC20Impl *ERC20ImplFilterer) WatchPrintingLocked(opts *bind.WatchOpts, sink chan<- *ERC20ImplPrintingLocked) (event.Subscription, error) {

	logs, sub, err := _ERC20Impl.contract.WatchLogs(opts, "PrintingLocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20ImplPrintingLocked)
				if err := _ERC20Impl.contract.UnpackLog(event, "PrintingLocked", log); err != nil {
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

// ERC20ImplUpgradeableABI is the input ABI used to generate the binding from.
const ERC20ImplUpgradeableABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_proposedCustodian\",\"type\":\"address\"}],\"name\":\"requestCustodianChange\",\"outputs\":[{\"name\":\"lockId\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"custodian\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_lockId\",\"type\":\"bytes32\"}],\"name\":\"confirmCustodianChange\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"erc20Impl\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proposedImpl\",\"type\":\"address\"}],\"name\":\"requestImplChange\",\"outputs\":[{\"name\":\"lockId\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_lockId\",\"type\":\"bytes32\"}],\"name\":\"confirmImplChange\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"implChangeReqs\",\"outputs\":[{\"name\":\"proposedNew\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lockRequestCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"custodianChangeReqs\",\"outputs\":[{\"name\":\"proposedNew\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_custodian\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_lockId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_msgSender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_proposedImpl\",\"type\":\"address\"}],\"name\":\"ImplChangeRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_lockId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_newImpl\",\"type\":\"address\"}],\"name\":\"ImplChangeConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_lockId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_msgSender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_proposedCustodian\",\"type\":\"address\"}],\"name\":\"CustodianChangeRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_lockId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_newCustodian\",\"type\":\"address\"}],\"name\":\"CustodianChangeConfirmed\",\"type\":\"event\"}]"

// ERC20ImplUpgradeableBin is the compiled bytecode used for deploying new contracts.
const ERC20ImplUpgradeableBin = `0x608060405234801561001057600080fd5b506040516020806105d183398101604052516000805560018054600160a060020a03909216600160a060020a03199283161790556003805490911690556105758061005c6000396000f3006080604052600436106100985763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166315b21082811461009d578063375b74c3146100d05780633a8343ee146101015780633c389cc41461011b57806348f9e246146101305780638181b02914610151578063b508069b14610169578063cb81fecf14610181578063cf6e448814610196575b600080fd5b3480156100a957600080fd5b506100be600160a060020a03600435166101ae565b60408051918252519081900360200190f35b3480156100dc57600080fd5b506100e5610261565b60408051600160a060020a039092168252519081900360200190f35b34801561010d57600080fd5b50610119600435610270565b005b34801561012757600080fd5b506100e5610311565b34801561013c57600080fd5b506100be600160a060020a0360043516610320565b34801561015d57600080fd5b506101196004356103d3565b34801561017557600080fd5b506100e5600435610474565b34801561018d57600080fd5b506100be61048f565b3480156101a257600080fd5b506100e5600435610495565b6000600160a060020a03821615156101c557600080fd5b6101cd6104b0565b6040805160208181018352600160a060020a0380871680845260008681526002845285902093518454921673ffffffffffffffffffffffffffffffffffffffff1990921691909117909255825184815233918101919091528083019190915290519192507fd76fc900a7e1a6fcf11d54b7ba943918df6c53a3128140658c389b3da1e997ba919081900360600190a1919050565b600154600160a060020a031681565b600154600160a060020a0316331461028757600080fd5b610290816104f1565b60018054600160a060020a0392831673ffffffffffffffffffffffffffffffffffffffff1991821617825560008481526002602090815260409182902080549093169092559154825185815293169083015280517f9a99272c0f6b7a30ef9e76e684a7cd408bfd4f11a72f36a8e276253c920e442d9281900390910190a150565b600354600160a060020a031681565b6000600160a060020a038216151561033757600080fd5b61033f6104b0565b6040805160208181018352600160a060020a0380871680845260008681526004845285902093518454921673ffffffffffffffffffffffffffffffffffffffff1990921691909117909255825184815233918101919091528083019190915290519192507f5df12834436b8dc248df3f7f1796a3e39f851d610be49cdcd92514fa821b9f97919081900360600190a1919050565b600154600160a060020a031633146103ea57600080fd5b6103f381610525565b60038054600160a060020a0392831673ffffffffffffffffffffffffffffffffffffffff1991821617825560008481526004602090815260409182902080549093169092559154825185815293169083015280517f9d55b0349a0a4c5b511f72228170bb91d45c9ac78dba8ab5b4175d3ed42f06b39281900390910190a150565b600460205260009081526040902054600160a060020a031681565b60005481565b600260205260009081526040902054600160a060020a031681565b60008054600101908190556040805160001943014081526c010000000000000000000000003002602082015260348101929092525190819003605401902090565b60008181526002602052604081208054600160a060020a0316151561051557600080fd5b54600160a060020a031692915050565b60008181526004602052604081208054600160a060020a0316151561051557600080fd00a165627a7a723058207d87e96151e0e7cc5a037b72ff0505e1822cc6814d8559a868b21c95eadb80d20029`

// DeployERC20ImplUpgradeable deploys a new Ethereum contract, binding an instance of ERC20ImplUpgradeable to it.
func DeployERC20ImplUpgradeable(auth *bind.TransactOpts, backend bind.ContractBackend, _custodian common.Address) (common.Address, *types.Transaction, *ERC20ImplUpgradeable, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ImplUpgradeableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20ImplUpgradeableBin), backend, _custodian)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20ImplUpgradeable{ERC20ImplUpgradeableCaller: ERC20ImplUpgradeableCaller{contract: contract}, ERC20ImplUpgradeableTransactor: ERC20ImplUpgradeableTransactor{contract: contract}, ERC20ImplUpgradeableFilterer: ERC20ImplUpgradeableFilterer{contract: contract}}, nil
}

// ERC20ImplUpgradeable is an auto generated Go binding around an Ethereum contract.
type ERC20ImplUpgradeable struct {
	ERC20ImplUpgradeableCaller     // Read-only binding to the contract
	ERC20ImplUpgradeableTransactor // Write-only binding to the contract
	ERC20ImplUpgradeableFilterer   // Log filterer for contract events
}

// ERC20ImplUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20ImplUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20ImplUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20ImplUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20ImplUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20ImplUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20ImplUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20ImplUpgradeableSession struct {
	Contract     *ERC20ImplUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ERC20ImplUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20ImplUpgradeableCallerSession struct {
	Contract *ERC20ImplUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// ERC20ImplUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20ImplUpgradeableTransactorSession struct {
	Contract     *ERC20ImplUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ERC20ImplUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20ImplUpgradeableRaw struct {
	Contract *ERC20ImplUpgradeable // Generic contract binding to access the raw methods on
}

// ERC20ImplUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20ImplUpgradeableCallerRaw struct {
	Contract *ERC20ImplUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20ImplUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20ImplUpgradeableTransactorRaw struct {
	Contract *ERC20ImplUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20ImplUpgradeable creates a new instance of ERC20ImplUpgradeable, bound to a specific deployed contract.
func NewERC20ImplUpgradeable(address common.Address, backend bind.ContractBackend) (*ERC20ImplUpgradeable, error) {
	contract, err := bindERC20ImplUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20ImplUpgradeable{ERC20ImplUpgradeableCaller: ERC20ImplUpgradeableCaller{contract: contract}, ERC20ImplUpgradeableTransactor: ERC20ImplUpgradeableTransactor{contract: contract}, ERC20ImplUpgradeableFilterer: ERC20ImplUpgradeableFilterer{contract: contract}}, nil
}

// NewERC20ImplUpgradeableCaller creates a new read-only instance of ERC20ImplUpgradeable, bound to a specific deployed contract.
func NewERC20ImplUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*ERC20ImplUpgradeableCaller, error) {
	contract, err := bindERC20ImplUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20ImplUpgradeableCaller{contract: contract}, nil
}

// NewERC20ImplUpgradeableTransactor creates a new write-only instance of ERC20ImplUpgradeable, bound to a specific deployed contract.
func NewERC20ImplUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20ImplUpgradeableTransactor, error) {
	contract, err := bindERC20ImplUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20ImplUpgradeableTransactor{contract: contract}, nil
}

// NewERC20ImplUpgradeableFilterer creates a new log filterer instance of ERC20ImplUpgradeable, bound to a specific deployed contract.
func NewERC20ImplUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20ImplUpgradeableFilterer, error) {
	contract, err := bindERC20ImplUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20ImplUpgradeableFilterer{contract: contract}, nil
}

// bindERC20ImplUpgradeable binds a generic wrapper to an already deployed contract.
func bindERC20ImplUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ImplUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20ImplUpgradeable *ERC20ImplUpgradeableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20ImplUpgradeable.Contract.ERC20ImplUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20ImplUpgradeable *ERC20ImplUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20ImplUpgradeable.Contract.ERC20ImplUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20ImplUpgradeable *ERC20ImplUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20ImplUpgradeable.Contract.ERC20ImplUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20ImplUpgradeable *ERC20ImplUpgradeableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20ImplUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20ImplUpgradeable *ERC20ImplUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20ImplUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20ImplUpgradeable *ERC20ImplUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20ImplUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// Custodian is a free data retrieval call binding the contract method 0x375b74c3.
//
// Solidity: function custodian() constant returns(address)
func (_ERC20ImplUpgradeable *ERC20ImplUpgradeableCaller) Custodian(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC20ImplUpgradeable.contract.Call(opts, out, "custodian")
	return *ret0, err
}

// Custodian is a free data retrieval call binding the contract method 0x375b74c3.
//
// Solidity: function custodian() constant returns(address)
func (_ERC20ImplUpgradeable *ERC20ImplUpgradeableSession) Custodian() (common.Address, error) {
	return _ERC20ImplUpgradeable.Contract.Custodian(&_ERC20ImplUpgradeable.CallOpts)
}

// Custodian is a free data retrieval call binding the contract method 0x375b74c3.
//
// Solidity: function custodian() constant returns(address)
func (_ERC20ImplUpgradeable *ERC20ImplUpgradeableCallerSession) Custodian() (common.Address, error) {
	return _ERC20ImplUpgradeable.Contract.Custodian(&_ERC20ImplUpgradeable.CallOpts)
}

// CustodianChangeReqs is a free data retrieval call binding the contract method 0xcf6e4488.
//
// Solidity: function custodianChangeReqs( bytes32) constant returns(proposedNew address)
func (_ERC20ImplUpgradeable *ERC20ImplUpgradeableCaller) CustodianChangeReqs(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC20ImplUpgradeable.contract.Call(opts, out, "custodianChangeReqs", arg0)
	return *ret0, err
}

// CustodianChangeReqs is a free data retrieval call binding the contract method 0xcf6e4488.
//
// Solidity: function custodianChangeReqs( bytes32) constant returns(proposedNew address)
func (_ERC20ImplUpgradeable *ERC20ImplUpgradeableSession) CustodianChangeReqs(arg0 [32]byte) (common.Address, error) {
	return _ERC20ImplUpgradeable.Contract.CustodianChangeReqs(&_ERC20ImplUpgradeable.CallOpts, arg0)
}

// CustodianChangeReqs is a free data retrieval call binding the contract method 0xcf6e4488.
//
// Solidity: function custodianChangeReqs( bytes32) constant returns(proposedNew address)
func (_ERC20ImplUpgradeable *ERC20ImplUpgradeableCallerSession) CustodianChangeReqs(arg0 [32]byte) (common.Address, error) {
	return _ERC20ImplUpgradeable.Contract.CustodianChangeReqs(&_ERC20ImplUpgradeable.CallOpts, arg0)
}

// Erc20Impl is a free data retrieval call binding the contract method 0x3c389cc4.
//
// Solidity: function erc20Impl() constant returns(address)
func (_ERC20ImplUpgradeable *ERC20ImplUpgradeableCaller) Erc20Impl(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC20ImplUpgradeable.contract.Call(opts, out, "erc20Impl")
	return *ret0, err
}

// Erc20Impl is a free data retrieval call binding the contract method 0x3c389cc4.
//
// Solidity: function erc20Impl() constant returns(address)
func (_ERC20ImplUpgradeable *ERC20ImplUpgradeableSession) Erc20Impl() (common.Address, error) {
	return _ERC20ImplUpgradeable.Contract.Erc20Impl(&_ERC20ImplUpgradeable.CallOpts)
}

// Erc20Impl is a free data retrieval call binding the contract method 0x3c389cc4.
//
// Solidity: function erc20Impl() constant returns(address)
func (_ERC20ImplUpgradeable *ERC20ImplUpgradeableCallerSession) Erc20Impl() (common.Address, error) {
	return _ERC20ImplUpgradeable.Contract.Erc20Impl(&_ERC20ImplUpgradeable.CallOpts)
}

// ImplChangeReqs is a free data retrieval call binding the contract method 0xb508069b.
//
// Solidity: function implChangeReqs( bytes32) constant returns(proposedNew address)
func (_ERC20ImplUpgradeable *ERC20ImplUpgradeableCaller) ImplChangeReqs(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC20ImplUpgradeable.contract.Call(opts, out, "implChangeReqs", arg0)
	return *ret0, err
}

// ImplChangeReqs is a free data retrieval call binding the contract method 0xb508069b.
//
// Solidity: function implChangeReqs( bytes32) constant returns(proposedNew address)
func (_ERC20ImplUpgradeable *ERC20ImplUpgradeableSession) ImplChangeReqs(arg0 [32]byte) (common.Address, error) {
	return _ERC20ImplUpgradeable.Contract.ImplChangeReqs(&_ERC20ImplUpgradeable.CallOpts, arg0)
}

// ImplChangeReqs is a free data retrieval call binding the contract method 0xb508069b.
//
// Solidity: function implChangeReqs( bytes32) constant returns(proposedNew address)
func (_ERC20ImplUpgradeable *ERC20ImplUpgradeableCallerSession) ImplChangeReqs(arg0 [32]byte) (common.Address, error) {
	return _ERC20ImplUpgradeable.Contract.ImplChangeReqs(&_ERC20ImplUpgradeable.CallOpts, arg0)
}

// LockRequestCount is a free data retrieval call binding the contract method 0xcb81fecf.
//
// Solidity: function lockRequestCount() constant returns(uint256)
func (_ERC20ImplUpgradeable *ERC20ImplUpgradeableCaller) LockRequestCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20ImplUpgradeable.contract.Call(opts, out, "lockRequestCount")
	return *ret0, err
}

// LockRequestCount is a free data retrieval call binding the contract method 0xcb81fecf.
//
// Solidity: function lockRequestCount() constant returns(uint256)
func (_ERC20ImplUpgradeable *ERC20ImplUpgradeableSession) LockRequestCount() (*big.Int, error) {
	return _ERC20ImplUpgradeable.Contract.LockRequestCount(&_ERC20ImplUpgradeable.CallOpts)
}

// LockRequestCount is a free data retrieval call binding the contract method 0xcb81fecf.
//
// Solidity: function lockRequestCount() constant returns(uint256)
func (_ERC20ImplUpgradeable *ERC20ImplUpgradeableCallerSession) LockRequestCount() (*big.Int, error) {
	return _ERC20ImplUpgradeable.Contract.LockRequestCount(&_ERC20ImplUpgradeable.CallOpts)
}

// ConfirmCustodianChange is a paid mutator transaction binding the contract method 0x3a8343ee.
//
// Solidity: function confirmCustodianChange(_lockId bytes32) returns()
func (_ERC20ImplUpgradeable *ERC20ImplUpgradeableTransactor) ConfirmCustodianChange(opts *bind.TransactOpts, _lockId [32]byte) (*types.Transaction, error) {
	return _ERC20ImplUpgradeable.contract.Transact(opts, "confirmCustodianChange", _lockId)
}

// ConfirmCustodianChange is a paid mutator transaction binding the contract method 0x3a8343ee.
//
// Solidity: function confirmCustodianChange(_lockId bytes32) returns()
func (_ERC20ImplUpgradeable *ERC20ImplUpgradeableSession) ConfirmCustodianChange(_lockId [32]byte) (*types.Transaction, error) {
	return _ERC20ImplUpgradeable.Contract.ConfirmCustodianChange(&_ERC20ImplUpgradeable.TransactOpts, _lockId)
}

// ConfirmCustodianChange is a paid mutator transaction binding the contract method 0x3a8343ee.
//
// Solidity: function confirmCustodianChange(_lockId bytes32) returns()
func (_ERC20ImplUpgradeable *ERC20ImplUpgradeableTransactorSession) ConfirmCustodianChange(_lockId [32]byte) (*types.Transaction, error) {
	return _ERC20ImplUpgradeable.Contract.ConfirmCustodianChange(&_ERC20ImplUpgradeable.TransactOpts, _lockId)
}

// ConfirmImplChange is a paid mutator transaction binding the contract method 0x8181b029.
//
// Solidity: function confirmImplChange(_lockId bytes32) returns()
func (_ERC20ImplUpgradeable *ERC20ImplUpgradeableTransactor) ConfirmImplChange(opts *bind.TransactOpts, _lockId [32]byte) (*types.Transaction, error) {
	return _ERC20ImplUpgradeable.contract.Transact(opts, "confirmImplChange", _lockId)
}

// ConfirmImplChange is a paid mutator transaction binding the contract method 0x8181b029.
//
// Solidity: function confirmImplChange(_lockId bytes32) returns()
func (_ERC20ImplUpgradeable *ERC20ImplUpgradeableSession) ConfirmImplChange(_lockId [32]byte) (*types.Transaction, error) {
	return _ERC20ImplUpgradeable.Contract.ConfirmImplChange(&_ERC20ImplUpgradeable.TransactOpts, _lockId)
}

// ConfirmImplChange is a paid mutator transaction binding the contract method 0x8181b029.
//
// Solidity: function confirmImplChange(_lockId bytes32) returns()
func (_ERC20ImplUpgradeable *ERC20ImplUpgradeableTransactorSession) ConfirmImplChange(_lockId [32]byte) (*types.Transaction, error) {
	return _ERC20ImplUpgradeable.Contract.ConfirmImplChange(&_ERC20ImplUpgradeable.TransactOpts, _lockId)
}

// RequestCustodianChange is a paid mutator transaction binding the contract method 0x15b21082.
//
// Solidity: function requestCustodianChange(_proposedCustodian address) returns(lockId bytes32)
func (_ERC20ImplUpgradeable *ERC20ImplUpgradeableTransactor) RequestCustodianChange(opts *bind.TransactOpts, _proposedCustodian common.Address) (*types.Transaction, error) {
	return _ERC20ImplUpgradeable.contract.Transact(opts, "requestCustodianChange", _proposedCustodian)
}

// RequestCustodianChange is a paid mutator transaction binding the contract method 0x15b21082.
//
// Solidity: function requestCustodianChange(_proposedCustodian address) returns(lockId bytes32)
func (_ERC20ImplUpgradeable *ERC20ImplUpgradeableSession) RequestCustodianChange(_proposedCustodian common.Address) (*types.Transaction, error) {
	return _ERC20ImplUpgradeable.Contract.RequestCustodianChange(&_ERC20ImplUpgradeable.TransactOpts, _proposedCustodian)
}

// RequestCustodianChange is a paid mutator transaction binding the contract method 0x15b21082.
//
// Solidity: function requestCustodianChange(_proposedCustodian address) returns(lockId bytes32)
func (_ERC20ImplUpgradeable *ERC20ImplUpgradeableTransactorSession) RequestCustodianChange(_proposedCustodian common.Address) (*types.Transaction, error) {
	return _ERC20ImplUpgradeable.Contract.RequestCustodianChange(&_ERC20ImplUpgradeable.TransactOpts, _proposedCustodian)
}

// RequestImplChange is a paid mutator transaction binding the contract method 0x48f9e246.
//
// Solidity: function requestImplChange(_proposedImpl address) returns(lockId bytes32)
func (_ERC20ImplUpgradeable *ERC20ImplUpgradeableTransactor) RequestImplChange(opts *bind.TransactOpts, _proposedImpl common.Address) (*types.Transaction, error) {
	return _ERC20ImplUpgradeable.contract.Transact(opts, "requestImplChange", _proposedImpl)
}

// RequestImplChange is a paid mutator transaction binding the contract method 0x48f9e246.
//
// Solidity: function requestImplChange(_proposedImpl address) returns(lockId bytes32)
func (_ERC20ImplUpgradeable *ERC20ImplUpgradeableSession) RequestImplChange(_proposedImpl common.Address) (*types.Transaction, error) {
	return _ERC20ImplUpgradeable.Contract.RequestImplChange(&_ERC20ImplUpgradeable.TransactOpts, _proposedImpl)
}

// RequestImplChange is a paid mutator transaction binding the contract method 0x48f9e246.
//
// Solidity: function requestImplChange(_proposedImpl address) returns(lockId bytes32)
func (_ERC20ImplUpgradeable *ERC20ImplUpgradeableTransactorSession) RequestImplChange(_proposedImpl common.Address) (*types.Transaction, error) {
	return _ERC20ImplUpgradeable.Contract.RequestImplChange(&_ERC20ImplUpgradeable.TransactOpts, _proposedImpl)
}

// ERC20ImplUpgradeableCustodianChangeConfirmedIterator is returned from FilterCustodianChangeConfirmed and is used to iterate over the raw logs and unpacked data for CustodianChangeConfirmed events raised by the ERC20ImplUpgradeable contract.
type ERC20ImplUpgradeableCustodianChangeConfirmedIterator struct {
	Event *ERC20ImplUpgradeableCustodianChangeConfirmed // Event containing the contract specifics and raw log

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
func (it *ERC20ImplUpgradeableCustodianChangeConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20ImplUpgradeableCustodianChangeConfirmed)
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
		it.Event = new(ERC20ImplUpgradeableCustodianChangeConfirmed)
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
func (it *ERC20ImplUpgradeableCustodianChangeConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ImplUpgradeableCustodianChangeConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20ImplUpgradeableCustodianChangeConfirmed represents a CustodianChangeConfirmed event raised by the ERC20ImplUpgradeable contract.
type ERC20ImplUpgradeableCustodianChangeConfirmed struct {
	LockId       [32]byte
	NewCustodian common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCustodianChangeConfirmed is a free log retrieval operation binding the contract event 0x9a99272c0f6b7a30ef9e76e684a7cd408bfd4f11a72f36a8e276253c920e442d.
//
// Solidity: e CustodianChangeConfirmed(_lockId bytes32, _newCustodian address)
func (_ERC20ImplUpgradeable *ERC20ImplUpgradeableFilterer) FilterCustodianChangeConfirmed(opts *bind.FilterOpts) (*ERC20ImplUpgradeableCustodianChangeConfirmedIterator, error) {

	logs, sub, err := _ERC20ImplUpgradeable.contract.FilterLogs(opts, "CustodianChangeConfirmed")
	if err != nil {
		return nil, err
	}
	return &ERC20ImplUpgradeableCustodianChangeConfirmedIterator{contract: _ERC20ImplUpgradeable.contract, event: "CustodianChangeConfirmed", logs: logs, sub: sub}, nil
}

// WatchCustodianChangeConfirmed is a free log subscription operation binding the contract event 0x9a99272c0f6b7a30ef9e76e684a7cd408bfd4f11a72f36a8e276253c920e442d.
//
// Solidity: e CustodianChangeConfirmed(_lockId bytes32, _newCustodian address)
func (_ERC20ImplUpgradeable *ERC20ImplUpgradeableFilterer) WatchCustodianChangeConfirmed(opts *bind.WatchOpts, sink chan<- *ERC20ImplUpgradeableCustodianChangeConfirmed) (event.Subscription, error) {

	logs, sub, err := _ERC20ImplUpgradeable.contract.WatchLogs(opts, "CustodianChangeConfirmed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20ImplUpgradeableCustodianChangeConfirmed)
				if err := _ERC20ImplUpgradeable.contract.UnpackLog(event, "CustodianChangeConfirmed", log); err != nil {
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

// ERC20ImplUpgradeableCustodianChangeRequestedIterator is returned from FilterCustodianChangeRequested and is used to iterate over the raw logs and unpacked data for CustodianChangeRequested events raised by the ERC20ImplUpgradeable contract.
type ERC20ImplUpgradeableCustodianChangeRequestedIterator struct {
	Event *ERC20ImplUpgradeableCustodianChangeRequested // Event containing the contract specifics and raw log

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
func (it *ERC20ImplUpgradeableCustodianChangeRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20ImplUpgradeableCustodianChangeRequested)
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
		it.Event = new(ERC20ImplUpgradeableCustodianChangeRequested)
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
func (it *ERC20ImplUpgradeableCustodianChangeRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ImplUpgradeableCustodianChangeRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20ImplUpgradeableCustodianChangeRequested represents a CustodianChangeRequested event raised by the ERC20ImplUpgradeable contract.
type ERC20ImplUpgradeableCustodianChangeRequested struct {
	LockId            [32]byte
	MsgSender         common.Address
	ProposedCustodian common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCustodianChangeRequested is a free log retrieval operation binding the contract event 0xd76fc900a7e1a6fcf11d54b7ba943918df6c53a3128140658c389b3da1e997ba.
//
// Solidity: e CustodianChangeRequested(_lockId bytes32, _msgSender address, _proposedCustodian address)
func (_ERC20ImplUpgradeable *ERC20ImplUpgradeableFilterer) FilterCustodianChangeRequested(opts *bind.FilterOpts) (*ERC20ImplUpgradeableCustodianChangeRequestedIterator, error) {

	logs, sub, err := _ERC20ImplUpgradeable.contract.FilterLogs(opts, "CustodianChangeRequested")
	if err != nil {
		return nil, err
	}
	return &ERC20ImplUpgradeableCustodianChangeRequestedIterator{contract: _ERC20ImplUpgradeable.contract, event: "CustodianChangeRequested", logs: logs, sub: sub}, nil
}

// WatchCustodianChangeRequested is a free log subscription operation binding the contract event 0xd76fc900a7e1a6fcf11d54b7ba943918df6c53a3128140658c389b3da1e997ba.
//
// Solidity: e CustodianChangeRequested(_lockId bytes32, _msgSender address, _proposedCustodian address)
func (_ERC20ImplUpgradeable *ERC20ImplUpgradeableFilterer) WatchCustodianChangeRequested(opts *bind.WatchOpts, sink chan<- *ERC20ImplUpgradeableCustodianChangeRequested) (event.Subscription, error) {

	logs, sub, err := _ERC20ImplUpgradeable.contract.WatchLogs(opts, "CustodianChangeRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20ImplUpgradeableCustodianChangeRequested)
				if err := _ERC20ImplUpgradeable.contract.UnpackLog(event, "CustodianChangeRequested", log); err != nil {
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

// ERC20ImplUpgradeableImplChangeConfirmedIterator is returned from FilterImplChangeConfirmed and is used to iterate over the raw logs and unpacked data for ImplChangeConfirmed events raised by the ERC20ImplUpgradeable contract.
type ERC20ImplUpgradeableImplChangeConfirmedIterator struct {
	Event *ERC20ImplUpgradeableImplChangeConfirmed // Event containing the contract specifics and raw log

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
func (it *ERC20ImplUpgradeableImplChangeConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20ImplUpgradeableImplChangeConfirmed)
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
		it.Event = new(ERC20ImplUpgradeableImplChangeConfirmed)
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
func (it *ERC20ImplUpgradeableImplChangeConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ImplUpgradeableImplChangeConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20ImplUpgradeableImplChangeConfirmed represents a ImplChangeConfirmed event raised by the ERC20ImplUpgradeable contract.
type ERC20ImplUpgradeableImplChangeConfirmed struct {
	LockId  [32]byte
	NewImpl common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterImplChangeConfirmed is a free log retrieval operation binding the contract event 0x9d55b0349a0a4c5b511f72228170bb91d45c9ac78dba8ab5b4175d3ed42f06b3.
//
// Solidity: e ImplChangeConfirmed(_lockId bytes32, _newImpl address)
func (_ERC20ImplUpgradeable *ERC20ImplUpgradeableFilterer) FilterImplChangeConfirmed(opts *bind.FilterOpts) (*ERC20ImplUpgradeableImplChangeConfirmedIterator, error) {

	logs, sub, err := _ERC20ImplUpgradeable.contract.FilterLogs(opts, "ImplChangeConfirmed")
	if err != nil {
		return nil, err
	}
	return &ERC20ImplUpgradeableImplChangeConfirmedIterator{contract: _ERC20ImplUpgradeable.contract, event: "ImplChangeConfirmed", logs: logs, sub: sub}, nil
}

// WatchImplChangeConfirmed is a free log subscription operation binding the contract event 0x9d55b0349a0a4c5b511f72228170bb91d45c9ac78dba8ab5b4175d3ed42f06b3.
//
// Solidity: e ImplChangeConfirmed(_lockId bytes32, _newImpl address)
func (_ERC20ImplUpgradeable *ERC20ImplUpgradeableFilterer) WatchImplChangeConfirmed(opts *bind.WatchOpts, sink chan<- *ERC20ImplUpgradeableImplChangeConfirmed) (event.Subscription, error) {

	logs, sub, err := _ERC20ImplUpgradeable.contract.WatchLogs(opts, "ImplChangeConfirmed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20ImplUpgradeableImplChangeConfirmed)
				if err := _ERC20ImplUpgradeable.contract.UnpackLog(event, "ImplChangeConfirmed", log); err != nil {
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

// ERC20ImplUpgradeableImplChangeRequestedIterator is returned from FilterImplChangeRequested and is used to iterate over the raw logs and unpacked data for ImplChangeRequested events raised by the ERC20ImplUpgradeable contract.
type ERC20ImplUpgradeableImplChangeRequestedIterator struct {
	Event *ERC20ImplUpgradeableImplChangeRequested // Event containing the contract specifics and raw log

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
func (it *ERC20ImplUpgradeableImplChangeRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20ImplUpgradeableImplChangeRequested)
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
		it.Event = new(ERC20ImplUpgradeableImplChangeRequested)
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
func (it *ERC20ImplUpgradeableImplChangeRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ImplUpgradeableImplChangeRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20ImplUpgradeableImplChangeRequested represents a ImplChangeRequested event raised by the ERC20ImplUpgradeable contract.
type ERC20ImplUpgradeableImplChangeRequested struct {
	LockId       [32]byte
	MsgSender    common.Address
	ProposedImpl common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterImplChangeRequested is a free log retrieval operation binding the contract event 0x5df12834436b8dc248df3f7f1796a3e39f851d610be49cdcd92514fa821b9f97.
//
// Solidity: e ImplChangeRequested(_lockId bytes32, _msgSender address, _proposedImpl address)
func (_ERC20ImplUpgradeable *ERC20ImplUpgradeableFilterer) FilterImplChangeRequested(opts *bind.FilterOpts) (*ERC20ImplUpgradeableImplChangeRequestedIterator, error) {

	logs, sub, err := _ERC20ImplUpgradeable.contract.FilterLogs(opts, "ImplChangeRequested")
	if err != nil {
		return nil, err
	}
	return &ERC20ImplUpgradeableImplChangeRequestedIterator{contract: _ERC20ImplUpgradeable.contract, event: "ImplChangeRequested", logs: logs, sub: sub}, nil
}

// WatchImplChangeRequested is a free log subscription operation binding the contract event 0x5df12834436b8dc248df3f7f1796a3e39f851d610be49cdcd92514fa821b9f97.
//
// Solidity: e ImplChangeRequested(_lockId bytes32, _msgSender address, _proposedImpl address)
func (_ERC20ImplUpgradeable *ERC20ImplUpgradeableFilterer) WatchImplChangeRequested(opts *bind.WatchOpts, sink chan<- *ERC20ImplUpgradeableImplChangeRequested) (event.Subscription, error) {

	logs, sub, err := _ERC20ImplUpgradeable.contract.WatchLogs(opts, "ImplChangeRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20ImplUpgradeableImplChangeRequested)
				if err := _ERC20ImplUpgradeable.contract.UnpackLog(event, "ImplChangeRequested", log); err != nil {
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

// ERC20InterfaceABI is the input ABI used to generate the binding from.
const ERC20InterfaceABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// ERC20InterfaceBin is the compiled bytecode used for deploying new contracts.
const ERC20InterfaceBin = `0x`

// DeployERC20Interface deploys a new Ethereum contract, binding an instance of ERC20Interface to it.
func DeployERC20Interface(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20Interface, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20InterfaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20InterfaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Interface{ERC20InterfaceCaller: ERC20InterfaceCaller{contract: contract}, ERC20InterfaceTransactor: ERC20InterfaceTransactor{contract: contract}, ERC20InterfaceFilterer: ERC20InterfaceFilterer{contract: contract}}, nil
}

// ERC20Interface is an auto generated Go binding around an Ethereum contract.
type ERC20Interface struct {
	ERC20InterfaceCaller     // Read-only binding to the contract
	ERC20InterfaceTransactor // Write-only binding to the contract
	ERC20InterfaceFilterer   // Log filterer for contract events
}

// ERC20InterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20InterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20InterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20InterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20InterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20InterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20InterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20InterfaceSession struct {
	Contract     *ERC20Interface   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20InterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20InterfaceCallerSession struct {
	Contract *ERC20InterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ERC20InterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20InterfaceTransactorSession struct {
	Contract     *ERC20InterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ERC20InterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20InterfaceRaw struct {
	Contract *ERC20Interface // Generic contract binding to access the raw methods on
}

// ERC20InterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20InterfaceCallerRaw struct {
	Contract *ERC20InterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20InterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20InterfaceTransactorRaw struct {
	Contract *ERC20InterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Interface creates a new instance of ERC20Interface, bound to a specific deployed contract.
func NewERC20Interface(address common.Address, backend bind.ContractBackend) (*ERC20Interface, error) {
	contract, err := bindERC20Interface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Interface{ERC20InterfaceCaller: ERC20InterfaceCaller{contract: contract}, ERC20InterfaceTransactor: ERC20InterfaceTransactor{contract: contract}, ERC20InterfaceFilterer: ERC20InterfaceFilterer{contract: contract}}, nil
}

// NewERC20InterfaceCaller creates a new read-only instance of ERC20Interface, bound to a specific deployed contract.
func NewERC20InterfaceCaller(address common.Address, caller bind.ContractCaller) (*ERC20InterfaceCaller, error) {
	contract, err := bindERC20Interface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20InterfaceCaller{contract: contract}, nil
}

// NewERC20InterfaceTransactor creates a new write-only instance of ERC20Interface, bound to a specific deployed contract.
func NewERC20InterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20InterfaceTransactor, error) {
	contract, err := bindERC20Interface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20InterfaceTransactor{contract: contract}, nil
}

// NewERC20InterfaceFilterer creates a new log filterer instance of ERC20Interface, bound to a specific deployed contract.
func NewERC20InterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20InterfaceFilterer, error) {
	contract, err := bindERC20Interface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20InterfaceFilterer{contract: contract}, nil
}

// bindERC20Interface binds a generic wrapper to an already deployed contract.
func bindERC20Interface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20InterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Interface *ERC20InterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Interface.Contract.ERC20InterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Interface *ERC20InterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Interface.Contract.ERC20InterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Interface *ERC20InterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Interface.Contract.ERC20InterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Interface *ERC20InterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Interface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Interface *ERC20InterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Interface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Interface *ERC20InterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Interface.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_ERC20Interface *ERC20InterfaceCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Interface.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_ERC20Interface *ERC20InterfaceSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ERC20Interface.Contract.Allowance(&_ERC20Interface.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_ERC20Interface *ERC20InterfaceCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ERC20Interface.Contract.Allowance(&_ERC20Interface.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ERC20Interface *ERC20InterfaceCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Interface.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ERC20Interface *ERC20InterfaceSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC20Interface.Contract.BalanceOf(&_ERC20Interface.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ERC20Interface *ERC20InterfaceCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC20Interface.Contract.BalanceOf(&_ERC20Interface.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Interface *ERC20InterfaceCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Interface.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Interface *ERC20InterfaceSession) TotalSupply() (*big.Int, error) {
	return _ERC20Interface.Contract.TotalSupply(&_ERC20Interface.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Interface *ERC20InterfaceCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Interface.Contract.TotalSupply(&_ERC20Interface.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_ERC20Interface *ERC20InterfaceTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Interface.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_ERC20Interface *ERC20InterfaceSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Interface.Contract.Approve(&_ERC20Interface.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_ERC20Interface *ERC20InterfaceTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Interface.Contract.Approve(&_ERC20Interface.TransactOpts, _spender, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_ERC20Interface *ERC20InterfaceTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Interface.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_ERC20Interface *ERC20InterfaceSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Interface.Contract.Transfer(&_ERC20Interface.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_ERC20Interface *ERC20InterfaceTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Interface.Contract.Transfer(&_ERC20Interface.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_ERC20Interface *ERC20InterfaceTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Interface.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_ERC20Interface *ERC20InterfaceSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Interface.Contract.TransferFrom(&_ERC20Interface.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_ERC20Interface *ERC20InterfaceTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Interface.Contract.TransferFrom(&_ERC20Interface.TransactOpts, _from, _to, _value)
}

// ERC20InterfaceApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20Interface contract.
type ERC20InterfaceApprovalIterator struct {
	Event *ERC20InterfaceApproval // Event containing the contract specifics and raw log

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
func (it *ERC20InterfaceApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20InterfaceApproval)
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
		it.Event = new(ERC20InterfaceApproval)
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
func (it *ERC20InterfaceApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20InterfaceApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20InterfaceApproval represents a Approval event raised by the ERC20Interface contract.
type ERC20InterfaceApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_ERC20Interface *ERC20InterfaceFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*ERC20InterfaceApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _ERC20Interface.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20InterfaceApprovalIterator{contract: _ERC20Interface.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_ERC20Interface *ERC20InterfaceFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20InterfaceApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _ERC20Interface.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20InterfaceApproval)
				if err := _ERC20Interface.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC20InterfaceTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Interface contract.
type ERC20InterfaceTransferIterator struct {
	Event *ERC20InterfaceTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20InterfaceTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20InterfaceTransfer)
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
		it.Event = new(ERC20InterfaceTransfer)
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
func (it *ERC20InterfaceTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20InterfaceTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20InterfaceTransfer represents a Transfer event raised by the ERC20Interface contract.
type ERC20InterfaceTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _value uint256)
func (_ERC20Interface *ERC20InterfaceFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*ERC20InterfaceTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ERC20Interface.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20InterfaceTransferIterator{contract: _ERC20Interface.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _value uint256)
func (_ERC20Interface *ERC20InterfaceFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20InterfaceTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ERC20Interface.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20InterfaceTransfer)
				if err := _ERC20Interface.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ERC20ProxyABI is the input ABI used to generate the binding from.
const ERC20ProxyABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proposedCustodian\",\"type\":\"address\"}],\"name\":\"requestCustodianChange\",\"outputs\":[{\"name\":\"lockId\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"emitTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"custodian\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_lockId\",\"type\":\"bytes32\"}],\"name\":\"confirmCustodianChange\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"erc20Impl\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proposedImpl\",\"type\":\"address\"}],\"name\":\"requestImplChange\",\"outputs\":[{\"name\":\"lockId\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"emitApproval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_lockId\",\"type\":\"bytes32\"}],\"name\":\"confirmImplChange\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"implChangeReqs\",\"outputs\":[{\"name\":\"proposedNew\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lockRequestCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"custodianChangeReqs\",\"outputs\":[{\"name\":\"proposedNew\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_decimals\",\"type\":\"uint8\"},{\"name\":\"_custodian\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_lockId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_msgSender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_proposedImpl\",\"type\":\"address\"}],\"name\":\"ImplChangeRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_lockId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_newImpl\",\"type\":\"address\"}],\"name\":\"ImplChangeConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_lockId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_msgSender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_proposedCustodian\",\"type\":\"address\"}],\"name\":\"CustodianChangeRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_lockId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_newCustodian\",\"type\":\"address\"}],\"name\":\"CustodianChangeConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// ERC20ProxyBin is the compiled bytecode used for deploying new contracts.
const ERC20ProxyBin = `0x608060405234801561001057600080fd5b5060405162000fd338038062000fd383398101604090815281516020808401519284015160608501516000805560018054600160a060020a038316600160a060020a031991821617909155600380549091169055928501805190959490940193909291610082916005918701906100b4565b5082516100969060069060208601906100b4565b50506007805460ff191660ff929092169190911790555061014f9050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100f557805160ff1916838001178555610122565b82800160010185558215610122579182015b82811115610122578251825591602001919060010190610107565b5061012e929150610132565b5090565b61014c91905b8082111561012e5760008155600101610138565b90565b610e74806200015f6000396000f3006080604052600436106101275763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde03811461012c578063095ea7b3146101b657806315b21082146101ee57806318160ddd1461022157806323b872dd1461023657806323de665114610260578063313ce5671461028c578063375b74c3146102b75780633a8343ee146102e85780633c389cc41461030057806348f9e246146103155780635687f2b814610336578063661884631461036057806370a08231146103845780638181b029146103a557806395d89b41146103bd578063a9059cbb146103d2578063b508069b146103f6578063cb81fecf1461040e578063cf6e448814610423578063d73dd6231461043b578063dd62ed3e1461045f575b600080fd5b34801561013857600080fd5b50610141610486565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561017b578181015183820152602001610163565b50505050905090810190601f1680156101a85780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101c257600080fd5b506101da600160a060020a0360043516602435610514565b604080519115158252519081900360200190f35b3480156101fa57600080fd5b5061020f600160a060020a03600435166105bf565b60408051918252519081900360200190f35b34801561022d57600080fd5b5061020f610672565b34801561024257600080fd5b506101da600160a060020a0360043581169060243516604435610702565b34801561026c57600080fd5b5061028a600160a060020a03600435811690602435166044356107b6565b005b34801561029857600080fd5b506102a161081d565b6040805160ff9092168252519081900360200190f35b3480156102c357600080fd5b506102cc610826565b60408051600160a060020a039092168252519081900360200190f35b3480156102f457600080fd5b5061028a600435610835565b34801561030c57600080fd5b506102cc6108d6565b34801561032157600080fd5b5061020f600160a060020a03600435166108e5565b34801561034257600080fd5b5061028a600160a060020a0360043581169060243516604435610998565b34801561036c57600080fd5b506101da600160a060020a03600435166024356109ff565b34801561039057600080fd5b5061020f600160a060020a0360043516610a77565b3480156103b157600080fd5b5061028a600435610b14565b3480156103c957600080fd5b50610141610bb5565b3480156103de57600080fd5b506101da600160a060020a0360043516602435610c10565b34801561040257600080fd5b506102cc600435610c88565b34801561041a57600080fd5b5061020f610ca3565b34801561042f57600080fd5b506102cc600435610ca9565b34801561044757600080fd5b506101da600160a060020a0360043516602435610cc4565b34801561046b57600080fd5b5061020f600160a060020a0360043581169060243516610d3c565b6005805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561050c5780601f106104e15761010080835404028352916020019161050c565b820191906000526020600020905b8154815290600101906020018083116104ef57829003601f168201915b505050505081565b600354604080517f89064fd2000000000000000000000000000000000000000000000000000000008152336004820152600160a060020a03858116602483015260448201859052915160009392909216916389064fd29160648082019260209290919082900301818787803b15801561058c57600080fd5b505af11580156105a0573d6000803e3d6000fd5b505050506040513d60208110156105b657600080fd5b50519392505050565b6000600160a060020a03821615156105d657600080fd5b6105de610daf565b6040805160208181018352600160a060020a0380871680845260008681526002845285902093518454921673ffffffffffffffffffffffffffffffffffffffff1990921691909117909255825184815233918101919091528083019190915290519192507fd76fc900a7e1a6fcf11d54b7ba943918df6c53a3128140658c389b3da1e997ba919081900360600190a1919050565b600354604080517f18160ddd0000000000000000000000000000000000000000000000000000000081529051600092600160a060020a0316916318160ddd91600480830192602092919082900301818787803b1580156106d157600080fd5b505af11580156106e5573d6000803e3d6000fd5b505050506040513d60208110156106fb57600080fd5b5051905090565b600354604080517f5d5e22cd000000000000000000000000000000000000000000000000000000008152336004820152600160a060020a03868116602483015285811660448301526064820185905291516000939290921691635d5e22cd9160848082019260209290919082900301818787803b15801561078257600080fd5b505af1158015610796573d6000803e3d6000fd5b505050506040513d60208110156107ac57600080fd5b5051949350505050565b600354600160a060020a031633146107cd57600080fd5b81600160a060020a031683600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a3505050565b60075460ff1681565b600154600160a060020a031681565b600154600160a060020a0316331461084c57600080fd5b61085581610df0565b60018054600160a060020a0392831673ffffffffffffffffffffffffffffffffffffffff1991821617825560008481526002602090815260409182902080549093169092559154825185815293169083015280517f9a99272c0f6b7a30ef9e76e684a7cd408bfd4f11a72f36a8e276253c920e442d9281900390910190a150565b600354600160a060020a031681565b6000600160a060020a03821615156108fc57600080fd5b610904610daf565b6040805160208181018352600160a060020a0380871680845260008681526004845285902093518454921673ffffffffffffffffffffffffffffffffffffffff1990921691909117909255825184815233918101919091528083019190915290519192507f5df12834436b8dc248df3f7f1796a3e39f851d610be49cdcd92514fa821b9f97919081900360600190a1919050565b600354600160a060020a031633146109af57600080fd5b81600160a060020a031683600160a060020a03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040518082815260200191505060405180910390a3505050565b600354604080517f61e1077d000000000000000000000000000000000000000000000000000000008152336004820152600160a060020a03858116602483015260448201859052915160009392909216916361e1077d9160648082019260209290919082900301818787803b15801561058c57600080fd5b600354604080517f70a08231000000000000000000000000000000000000000000000000000000008152600160a060020a038481166004830152915160009392909216916370a082319160248082019260209290919082900301818787803b158015610ae257600080fd5b505af1158015610af6573d6000803e3d6000fd5b505050506040513d6020811015610b0c57600080fd5b505192915050565b600154600160a060020a03163314610b2b57600080fd5b610b3481610e24565b60038054600160a060020a0392831673ffffffffffffffffffffffffffffffffffffffff1991821617825560008481526004602090815260409182902080549093169092559154825185815293169083015280517f9d55b0349a0a4c5b511f72228170bb91d45c9ac78dba8ab5b4175d3ed42f06b39281900390910190a150565b6006805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561050c5780601f106104e15761010080835404028352916020019161050c565b600354604080517fdfe0f0ca000000000000000000000000000000000000000000000000000000008152336004820152600160a060020a038581166024830152604482018590529151600093929092169163dfe0f0ca9160648082019260209290919082900301818787803b15801561058c57600080fd5b600460205260009081526040902054600160a060020a031681565b60005481565b600260205260009081526040902054600160a060020a031681565b600354604080517f2e0179b5000000000000000000000000000000000000000000000000000000008152336004820152600160a060020a0385811660248301526044820185905291516000939290921691632e0179b59160648082019260209290919082900301818787803b15801561058c57600080fd5b600354604080517fdd62ed3e000000000000000000000000000000000000000000000000000000008152600160a060020a03858116600483015284811660248301529151600093929092169163dd62ed3e9160448082019260209290919082900301818787803b15801561058c57600080fd5b60008054600101908190556040805160001943014081526c010000000000000000000000003002602082015260348101929092525190819003605401902090565b60008181526002602052604081208054600160a060020a03161515610e1457600080fd5b54600160a060020a031692915050565b60008181526004602052604081208054600160a060020a03161515610e1457600080fd00a165627a7a7230582039489682c6b741429145d5932a75cc31b7fabe1f6ac7f95a9deb93d079e0da8c0029`

// DeployERC20Proxy deploys a new Ethereum contract, binding an instance of ERC20Proxy to it.
func DeployERC20Proxy(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _symbol string, _decimals uint8, _custodian common.Address) (common.Address, *types.Transaction, *ERC20Proxy, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20ProxyBin), backend, _name, _symbol, _decimals, _custodian)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Proxy{ERC20ProxyCaller: ERC20ProxyCaller{contract: contract}, ERC20ProxyTransactor: ERC20ProxyTransactor{contract: contract}, ERC20ProxyFilterer: ERC20ProxyFilterer{contract: contract}}, nil
}

// ERC20Proxy is an auto generated Go binding around an Ethereum contract.
type ERC20Proxy struct {
	ERC20ProxyCaller     // Read-only binding to the contract
	ERC20ProxyTransactor // Write-only binding to the contract
	ERC20ProxyFilterer   // Log filterer for contract events
}

// ERC20ProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20ProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20ProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20ProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20ProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20ProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20ProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20ProxySession struct {
	Contract     *ERC20Proxy       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20ProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20ProxyCallerSession struct {
	Contract *ERC20ProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ERC20ProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20ProxyTransactorSession struct {
	Contract     *ERC20ProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ERC20ProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20ProxyRaw struct {
	Contract *ERC20Proxy // Generic contract binding to access the raw methods on
}

// ERC20ProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20ProxyCallerRaw struct {
	Contract *ERC20ProxyCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20ProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20ProxyTransactorRaw struct {
	Contract *ERC20ProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Proxy creates a new instance of ERC20Proxy, bound to a specific deployed contract.
func NewERC20Proxy(address common.Address, backend bind.ContractBackend) (*ERC20Proxy, error) {
	contract, err := bindERC20Proxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Proxy{ERC20ProxyCaller: ERC20ProxyCaller{contract: contract}, ERC20ProxyTransactor: ERC20ProxyTransactor{contract: contract}, ERC20ProxyFilterer: ERC20ProxyFilterer{contract: contract}}, nil
}

// NewERC20ProxyCaller creates a new read-only instance of ERC20Proxy, bound to a specific deployed contract.
func NewERC20ProxyCaller(address common.Address, caller bind.ContractCaller) (*ERC20ProxyCaller, error) {
	contract, err := bindERC20Proxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20ProxyCaller{contract: contract}, nil
}

// NewERC20ProxyTransactor creates a new write-only instance of ERC20Proxy, bound to a specific deployed contract.
func NewERC20ProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20ProxyTransactor, error) {
	contract, err := bindERC20Proxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20ProxyTransactor{contract: contract}, nil
}

// NewERC20ProxyFilterer creates a new log filterer instance of ERC20Proxy, bound to a specific deployed contract.
func NewERC20ProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20ProxyFilterer, error) {
	contract, err := bindERC20Proxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20ProxyFilterer{contract: contract}, nil
}

// bindERC20Proxy binds a generic wrapper to an already deployed contract.
func bindERC20Proxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Proxy *ERC20ProxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Proxy.Contract.ERC20ProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Proxy *ERC20ProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Proxy.Contract.ERC20ProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Proxy *ERC20ProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Proxy.Contract.ERC20ProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Proxy *ERC20ProxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Proxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Proxy *ERC20ProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Proxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Proxy *ERC20ProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Proxy.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_ERC20Proxy *ERC20ProxyCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Proxy.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_ERC20Proxy *ERC20ProxySession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ERC20Proxy.Contract.Allowance(&_ERC20Proxy.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_ERC20Proxy *ERC20ProxyCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ERC20Proxy.Contract.Allowance(&_ERC20Proxy.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ERC20Proxy *ERC20ProxyCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Proxy.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ERC20Proxy *ERC20ProxySession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC20Proxy.Contract.BalanceOf(&_ERC20Proxy.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ERC20Proxy *ERC20ProxyCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC20Proxy.Contract.BalanceOf(&_ERC20Proxy.CallOpts, _owner)
}

// Custodian is a free data retrieval call binding the contract method 0x375b74c3.
//
// Solidity: function custodian() constant returns(address)
func (_ERC20Proxy *ERC20ProxyCaller) Custodian(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC20Proxy.contract.Call(opts, out, "custodian")
	return *ret0, err
}

// Custodian is a free data retrieval call binding the contract method 0x375b74c3.
//
// Solidity: function custodian() constant returns(address)
func (_ERC20Proxy *ERC20ProxySession) Custodian() (common.Address, error) {
	return _ERC20Proxy.Contract.Custodian(&_ERC20Proxy.CallOpts)
}

// Custodian is a free data retrieval call binding the contract method 0x375b74c3.
//
// Solidity: function custodian() constant returns(address)
func (_ERC20Proxy *ERC20ProxyCallerSession) Custodian() (common.Address, error) {
	return _ERC20Proxy.Contract.Custodian(&_ERC20Proxy.CallOpts)
}

// CustodianChangeReqs is a free data retrieval call binding the contract method 0xcf6e4488.
//
// Solidity: function custodianChangeReqs( bytes32) constant returns(proposedNew address)
func (_ERC20Proxy *ERC20ProxyCaller) CustodianChangeReqs(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC20Proxy.contract.Call(opts, out, "custodianChangeReqs", arg0)
	return *ret0, err
}

// CustodianChangeReqs is a free data retrieval call binding the contract method 0xcf6e4488.
//
// Solidity: function custodianChangeReqs( bytes32) constant returns(proposedNew address)
func (_ERC20Proxy *ERC20ProxySession) CustodianChangeReqs(arg0 [32]byte) (common.Address, error) {
	return _ERC20Proxy.Contract.CustodianChangeReqs(&_ERC20Proxy.CallOpts, arg0)
}

// CustodianChangeReqs is a free data retrieval call binding the contract method 0xcf6e4488.
//
// Solidity: function custodianChangeReqs( bytes32) constant returns(proposedNew address)
func (_ERC20Proxy *ERC20ProxyCallerSession) CustodianChangeReqs(arg0 [32]byte) (common.Address, error) {
	return _ERC20Proxy.Contract.CustodianChangeReqs(&_ERC20Proxy.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ERC20Proxy *ERC20ProxyCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ERC20Proxy.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ERC20Proxy *ERC20ProxySession) Decimals() (uint8, error) {
	return _ERC20Proxy.Contract.Decimals(&_ERC20Proxy.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ERC20Proxy *ERC20ProxyCallerSession) Decimals() (uint8, error) {
	return _ERC20Proxy.Contract.Decimals(&_ERC20Proxy.CallOpts)
}

// Erc20Impl is a free data retrieval call binding the contract method 0x3c389cc4.
//
// Solidity: function erc20Impl() constant returns(address)
func (_ERC20Proxy *ERC20ProxyCaller) Erc20Impl(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC20Proxy.contract.Call(opts, out, "erc20Impl")
	return *ret0, err
}

// Erc20Impl is a free data retrieval call binding the contract method 0x3c389cc4.
//
// Solidity: function erc20Impl() constant returns(address)
func (_ERC20Proxy *ERC20ProxySession) Erc20Impl() (common.Address, error) {
	return _ERC20Proxy.Contract.Erc20Impl(&_ERC20Proxy.CallOpts)
}

// Erc20Impl is a free data retrieval call binding the contract method 0x3c389cc4.
//
// Solidity: function erc20Impl() constant returns(address)
func (_ERC20Proxy *ERC20ProxyCallerSession) Erc20Impl() (common.Address, error) {
	return _ERC20Proxy.Contract.Erc20Impl(&_ERC20Proxy.CallOpts)
}

// ImplChangeReqs is a free data retrieval call binding the contract method 0xb508069b.
//
// Solidity: function implChangeReqs( bytes32) constant returns(proposedNew address)
func (_ERC20Proxy *ERC20ProxyCaller) ImplChangeReqs(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC20Proxy.contract.Call(opts, out, "implChangeReqs", arg0)
	return *ret0, err
}

// ImplChangeReqs is a free data retrieval call binding the contract method 0xb508069b.
//
// Solidity: function implChangeReqs( bytes32) constant returns(proposedNew address)
func (_ERC20Proxy *ERC20ProxySession) ImplChangeReqs(arg0 [32]byte) (common.Address, error) {
	return _ERC20Proxy.Contract.ImplChangeReqs(&_ERC20Proxy.CallOpts, arg0)
}

// ImplChangeReqs is a free data retrieval call binding the contract method 0xb508069b.
//
// Solidity: function implChangeReqs( bytes32) constant returns(proposedNew address)
func (_ERC20Proxy *ERC20ProxyCallerSession) ImplChangeReqs(arg0 [32]byte) (common.Address, error) {
	return _ERC20Proxy.Contract.ImplChangeReqs(&_ERC20Proxy.CallOpts, arg0)
}

// LockRequestCount is a free data retrieval call binding the contract method 0xcb81fecf.
//
// Solidity: function lockRequestCount() constant returns(uint256)
func (_ERC20Proxy *ERC20ProxyCaller) LockRequestCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Proxy.contract.Call(opts, out, "lockRequestCount")
	return *ret0, err
}

// LockRequestCount is a free data retrieval call binding the contract method 0xcb81fecf.
//
// Solidity: function lockRequestCount() constant returns(uint256)
func (_ERC20Proxy *ERC20ProxySession) LockRequestCount() (*big.Int, error) {
	return _ERC20Proxy.Contract.LockRequestCount(&_ERC20Proxy.CallOpts)
}

// LockRequestCount is a free data retrieval call binding the contract method 0xcb81fecf.
//
// Solidity: function lockRequestCount() constant returns(uint256)
func (_ERC20Proxy *ERC20ProxyCallerSession) LockRequestCount() (*big.Int, error) {
	return _ERC20Proxy.Contract.LockRequestCount(&_ERC20Proxy.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ERC20Proxy *ERC20ProxyCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC20Proxy.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ERC20Proxy *ERC20ProxySession) Name() (string, error) {
	return _ERC20Proxy.Contract.Name(&_ERC20Proxy.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ERC20Proxy *ERC20ProxyCallerSession) Name() (string, error) {
	return _ERC20Proxy.Contract.Name(&_ERC20Proxy.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ERC20Proxy *ERC20ProxyCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC20Proxy.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ERC20Proxy *ERC20ProxySession) Symbol() (string, error) {
	return _ERC20Proxy.Contract.Symbol(&_ERC20Proxy.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ERC20Proxy *ERC20ProxyCallerSession) Symbol() (string, error) {
	return _ERC20Proxy.Contract.Symbol(&_ERC20Proxy.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Proxy *ERC20ProxyCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Proxy.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Proxy *ERC20ProxySession) TotalSupply() (*big.Int, error) {
	return _ERC20Proxy.Contract.TotalSupply(&_ERC20Proxy.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Proxy *ERC20ProxyCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Proxy.Contract.TotalSupply(&_ERC20Proxy.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_ERC20Proxy *ERC20ProxyTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Proxy.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_ERC20Proxy *ERC20ProxySession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Proxy.Contract.Approve(&_ERC20Proxy.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_ERC20Proxy *ERC20ProxyTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Proxy.Contract.Approve(&_ERC20Proxy.TransactOpts, _spender, _value)
}

// ConfirmCustodianChange is a paid mutator transaction binding the contract method 0x3a8343ee.
//
// Solidity: function confirmCustodianChange(_lockId bytes32) returns()
func (_ERC20Proxy *ERC20ProxyTransactor) ConfirmCustodianChange(opts *bind.TransactOpts, _lockId [32]byte) (*types.Transaction, error) {
	return _ERC20Proxy.contract.Transact(opts, "confirmCustodianChange", _lockId)
}

// ConfirmCustodianChange is a paid mutator transaction binding the contract method 0x3a8343ee.
//
// Solidity: function confirmCustodianChange(_lockId bytes32) returns()
func (_ERC20Proxy *ERC20ProxySession) ConfirmCustodianChange(_lockId [32]byte) (*types.Transaction, error) {
	return _ERC20Proxy.Contract.ConfirmCustodianChange(&_ERC20Proxy.TransactOpts, _lockId)
}

// ConfirmCustodianChange is a paid mutator transaction binding the contract method 0x3a8343ee.
//
// Solidity: function confirmCustodianChange(_lockId bytes32) returns()
func (_ERC20Proxy *ERC20ProxyTransactorSession) ConfirmCustodianChange(_lockId [32]byte) (*types.Transaction, error) {
	return _ERC20Proxy.Contract.ConfirmCustodianChange(&_ERC20Proxy.TransactOpts, _lockId)
}

// ConfirmImplChange is a paid mutator transaction binding the contract method 0x8181b029.
//
// Solidity: function confirmImplChange(_lockId bytes32) returns()
func (_ERC20Proxy *ERC20ProxyTransactor) ConfirmImplChange(opts *bind.TransactOpts, _lockId [32]byte) (*types.Transaction, error) {
	return _ERC20Proxy.contract.Transact(opts, "confirmImplChange", _lockId)
}

// ConfirmImplChange is a paid mutator transaction binding the contract method 0x8181b029.
//
// Solidity: function confirmImplChange(_lockId bytes32) returns()
func (_ERC20Proxy *ERC20ProxySession) ConfirmImplChange(_lockId [32]byte) (*types.Transaction, error) {
	return _ERC20Proxy.Contract.ConfirmImplChange(&_ERC20Proxy.TransactOpts, _lockId)
}

// ConfirmImplChange is a paid mutator transaction binding the contract method 0x8181b029.
//
// Solidity: function confirmImplChange(_lockId bytes32) returns()
func (_ERC20Proxy *ERC20ProxyTransactorSession) ConfirmImplChange(_lockId [32]byte) (*types.Transaction, error) {
	return _ERC20Proxy.Contract.ConfirmImplChange(&_ERC20Proxy.TransactOpts, _lockId)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_ERC20Proxy *ERC20ProxyTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Proxy.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_ERC20Proxy *ERC20ProxySession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Proxy.Contract.DecreaseApproval(&_ERC20Proxy.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_ERC20Proxy *ERC20ProxyTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Proxy.Contract.DecreaseApproval(&_ERC20Proxy.TransactOpts, _spender, _subtractedValue)
}

// EmitApproval is a paid mutator transaction binding the contract method 0x5687f2b8.
//
// Solidity: function emitApproval(_owner address, _spender address, _value uint256) returns()
func (_ERC20Proxy *ERC20ProxyTransactor) EmitApproval(opts *bind.TransactOpts, _owner common.Address, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Proxy.contract.Transact(opts, "emitApproval", _owner, _spender, _value)
}

// EmitApproval is a paid mutator transaction binding the contract method 0x5687f2b8.
//
// Solidity: function emitApproval(_owner address, _spender address, _value uint256) returns()
func (_ERC20Proxy *ERC20ProxySession) EmitApproval(_owner common.Address, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Proxy.Contract.EmitApproval(&_ERC20Proxy.TransactOpts, _owner, _spender, _value)
}

// EmitApproval is a paid mutator transaction binding the contract method 0x5687f2b8.
//
// Solidity: function emitApproval(_owner address, _spender address, _value uint256) returns()
func (_ERC20Proxy *ERC20ProxyTransactorSession) EmitApproval(_owner common.Address, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Proxy.Contract.EmitApproval(&_ERC20Proxy.TransactOpts, _owner, _spender, _value)
}

// EmitTransfer is a paid mutator transaction binding the contract method 0x23de6651.
//
// Solidity: function emitTransfer(_from address, _to address, _value uint256) returns()
func (_ERC20Proxy *ERC20ProxyTransactor) EmitTransfer(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Proxy.contract.Transact(opts, "emitTransfer", _from, _to, _value)
}

// EmitTransfer is a paid mutator transaction binding the contract method 0x23de6651.
//
// Solidity: function emitTransfer(_from address, _to address, _value uint256) returns()
func (_ERC20Proxy *ERC20ProxySession) EmitTransfer(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Proxy.Contract.EmitTransfer(&_ERC20Proxy.TransactOpts, _from, _to, _value)
}

// EmitTransfer is a paid mutator transaction binding the contract method 0x23de6651.
//
// Solidity: function emitTransfer(_from address, _to address, _value uint256) returns()
func (_ERC20Proxy *ERC20ProxyTransactorSession) EmitTransfer(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Proxy.Contract.EmitTransfer(&_ERC20Proxy.TransactOpts, _from, _to, _value)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_ERC20Proxy *ERC20ProxyTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Proxy.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_ERC20Proxy *ERC20ProxySession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Proxy.Contract.IncreaseApproval(&_ERC20Proxy.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_ERC20Proxy *ERC20ProxyTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Proxy.Contract.IncreaseApproval(&_ERC20Proxy.TransactOpts, _spender, _addedValue)
}

// RequestCustodianChange is a paid mutator transaction binding the contract method 0x15b21082.
//
// Solidity: function requestCustodianChange(_proposedCustodian address) returns(lockId bytes32)
func (_ERC20Proxy *ERC20ProxyTransactor) RequestCustodianChange(opts *bind.TransactOpts, _proposedCustodian common.Address) (*types.Transaction, error) {
	return _ERC20Proxy.contract.Transact(opts, "requestCustodianChange", _proposedCustodian)
}

// RequestCustodianChange is a paid mutator transaction binding the contract method 0x15b21082.
//
// Solidity: function requestCustodianChange(_proposedCustodian address) returns(lockId bytes32)
func (_ERC20Proxy *ERC20ProxySession) RequestCustodianChange(_proposedCustodian common.Address) (*types.Transaction, error) {
	return _ERC20Proxy.Contract.RequestCustodianChange(&_ERC20Proxy.TransactOpts, _proposedCustodian)
}

// RequestCustodianChange is a paid mutator transaction binding the contract method 0x15b21082.
//
// Solidity: function requestCustodianChange(_proposedCustodian address) returns(lockId bytes32)
func (_ERC20Proxy *ERC20ProxyTransactorSession) RequestCustodianChange(_proposedCustodian common.Address) (*types.Transaction, error) {
	return _ERC20Proxy.Contract.RequestCustodianChange(&_ERC20Proxy.TransactOpts, _proposedCustodian)
}

// RequestImplChange is a paid mutator transaction binding the contract method 0x48f9e246.
//
// Solidity: function requestImplChange(_proposedImpl address) returns(lockId bytes32)
func (_ERC20Proxy *ERC20ProxyTransactor) RequestImplChange(opts *bind.TransactOpts, _proposedImpl common.Address) (*types.Transaction, error) {
	return _ERC20Proxy.contract.Transact(opts, "requestImplChange", _proposedImpl)
}

// RequestImplChange is a paid mutator transaction binding the contract method 0x48f9e246.
//
// Solidity: function requestImplChange(_proposedImpl address) returns(lockId bytes32)
func (_ERC20Proxy *ERC20ProxySession) RequestImplChange(_proposedImpl common.Address) (*types.Transaction, error) {
	return _ERC20Proxy.Contract.RequestImplChange(&_ERC20Proxy.TransactOpts, _proposedImpl)
}

// RequestImplChange is a paid mutator transaction binding the contract method 0x48f9e246.
//
// Solidity: function requestImplChange(_proposedImpl address) returns(lockId bytes32)
func (_ERC20Proxy *ERC20ProxyTransactorSession) RequestImplChange(_proposedImpl common.Address) (*types.Transaction, error) {
	return _ERC20Proxy.Contract.RequestImplChange(&_ERC20Proxy.TransactOpts, _proposedImpl)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_ERC20Proxy *ERC20ProxyTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Proxy.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_ERC20Proxy *ERC20ProxySession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Proxy.Contract.Transfer(&_ERC20Proxy.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_ERC20Proxy *ERC20ProxyTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Proxy.Contract.Transfer(&_ERC20Proxy.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_ERC20Proxy *ERC20ProxyTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Proxy.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_ERC20Proxy *ERC20ProxySession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Proxy.Contract.TransferFrom(&_ERC20Proxy.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_ERC20Proxy *ERC20ProxyTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Proxy.Contract.TransferFrom(&_ERC20Proxy.TransactOpts, _from, _to, _value)
}

// ERC20ProxyApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20Proxy contract.
type ERC20ProxyApprovalIterator struct {
	Event *ERC20ProxyApproval // Event containing the contract specifics and raw log

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
func (it *ERC20ProxyApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20ProxyApproval)
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
		it.Event = new(ERC20ProxyApproval)
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
func (it *ERC20ProxyApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ProxyApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20ProxyApproval represents a Approval event raised by the ERC20Proxy contract.
type ERC20ProxyApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_ERC20Proxy *ERC20ProxyFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*ERC20ProxyApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _ERC20Proxy.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ProxyApprovalIterator{contract: _ERC20Proxy.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_ERC20Proxy *ERC20ProxyFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20ProxyApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _ERC20Proxy.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20ProxyApproval)
				if err := _ERC20Proxy.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC20ProxyCustodianChangeConfirmedIterator is returned from FilterCustodianChangeConfirmed and is used to iterate over the raw logs and unpacked data for CustodianChangeConfirmed events raised by the ERC20Proxy contract.
type ERC20ProxyCustodianChangeConfirmedIterator struct {
	Event *ERC20ProxyCustodianChangeConfirmed // Event containing the contract specifics and raw log

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
func (it *ERC20ProxyCustodianChangeConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20ProxyCustodianChangeConfirmed)
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
		it.Event = new(ERC20ProxyCustodianChangeConfirmed)
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
func (it *ERC20ProxyCustodianChangeConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ProxyCustodianChangeConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20ProxyCustodianChangeConfirmed represents a CustodianChangeConfirmed event raised by the ERC20Proxy contract.
type ERC20ProxyCustodianChangeConfirmed struct {
	LockId       [32]byte
	NewCustodian common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCustodianChangeConfirmed is a free log retrieval operation binding the contract event 0x9a99272c0f6b7a30ef9e76e684a7cd408bfd4f11a72f36a8e276253c920e442d.
//
// Solidity: e CustodianChangeConfirmed(_lockId bytes32, _newCustodian address)
func (_ERC20Proxy *ERC20ProxyFilterer) FilterCustodianChangeConfirmed(opts *bind.FilterOpts) (*ERC20ProxyCustodianChangeConfirmedIterator, error) {

	logs, sub, err := _ERC20Proxy.contract.FilterLogs(opts, "CustodianChangeConfirmed")
	if err != nil {
		return nil, err
	}
	return &ERC20ProxyCustodianChangeConfirmedIterator{contract: _ERC20Proxy.contract, event: "CustodianChangeConfirmed", logs: logs, sub: sub}, nil
}

// WatchCustodianChangeConfirmed is a free log subscription operation binding the contract event 0x9a99272c0f6b7a30ef9e76e684a7cd408bfd4f11a72f36a8e276253c920e442d.
//
// Solidity: e CustodianChangeConfirmed(_lockId bytes32, _newCustodian address)
func (_ERC20Proxy *ERC20ProxyFilterer) WatchCustodianChangeConfirmed(opts *bind.WatchOpts, sink chan<- *ERC20ProxyCustodianChangeConfirmed) (event.Subscription, error) {

	logs, sub, err := _ERC20Proxy.contract.WatchLogs(opts, "CustodianChangeConfirmed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20ProxyCustodianChangeConfirmed)
				if err := _ERC20Proxy.contract.UnpackLog(event, "CustodianChangeConfirmed", log); err != nil {
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

// ERC20ProxyCustodianChangeRequestedIterator is returned from FilterCustodianChangeRequested and is used to iterate over the raw logs and unpacked data for CustodianChangeRequested events raised by the ERC20Proxy contract.
type ERC20ProxyCustodianChangeRequestedIterator struct {
	Event *ERC20ProxyCustodianChangeRequested // Event containing the contract specifics and raw log

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
func (it *ERC20ProxyCustodianChangeRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20ProxyCustodianChangeRequested)
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
		it.Event = new(ERC20ProxyCustodianChangeRequested)
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
func (it *ERC20ProxyCustodianChangeRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ProxyCustodianChangeRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20ProxyCustodianChangeRequested represents a CustodianChangeRequested event raised by the ERC20Proxy contract.
type ERC20ProxyCustodianChangeRequested struct {
	LockId            [32]byte
	MsgSender         common.Address
	ProposedCustodian common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCustodianChangeRequested is a free log retrieval operation binding the contract event 0xd76fc900a7e1a6fcf11d54b7ba943918df6c53a3128140658c389b3da1e997ba.
//
// Solidity: e CustodianChangeRequested(_lockId bytes32, _msgSender address, _proposedCustodian address)
func (_ERC20Proxy *ERC20ProxyFilterer) FilterCustodianChangeRequested(opts *bind.FilterOpts) (*ERC20ProxyCustodianChangeRequestedIterator, error) {

	logs, sub, err := _ERC20Proxy.contract.FilterLogs(opts, "CustodianChangeRequested")
	if err != nil {
		return nil, err
	}
	return &ERC20ProxyCustodianChangeRequestedIterator{contract: _ERC20Proxy.contract, event: "CustodianChangeRequested", logs: logs, sub: sub}, nil
}

// WatchCustodianChangeRequested is a free log subscription operation binding the contract event 0xd76fc900a7e1a6fcf11d54b7ba943918df6c53a3128140658c389b3da1e997ba.
//
// Solidity: e CustodianChangeRequested(_lockId bytes32, _msgSender address, _proposedCustodian address)
func (_ERC20Proxy *ERC20ProxyFilterer) WatchCustodianChangeRequested(opts *bind.WatchOpts, sink chan<- *ERC20ProxyCustodianChangeRequested) (event.Subscription, error) {

	logs, sub, err := _ERC20Proxy.contract.WatchLogs(opts, "CustodianChangeRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20ProxyCustodianChangeRequested)
				if err := _ERC20Proxy.contract.UnpackLog(event, "CustodianChangeRequested", log); err != nil {
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

// ERC20ProxyImplChangeConfirmedIterator is returned from FilterImplChangeConfirmed and is used to iterate over the raw logs and unpacked data for ImplChangeConfirmed events raised by the ERC20Proxy contract.
type ERC20ProxyImplChangeConfirmedIterator struct {
	Event *ERC20ProxyImplChangeConfirmed // Event containing the contract specifics and raw log

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
func (it *ERC20ProxyImplChangeConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20ProxyImplChangeConfirmed)
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
		it.Event = new(ERC20ProxyImplChangeConfirmed)
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
func (it *ERC20ProxyImplChangeConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ProxyImplChangeConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20ProxyImplChangeConfirmed represents a ImplChangeConfirmed event raised by the ERC20Proxy contract.
type ERC20ProxyImplChangeConfirmed struct {
	LockId  [32]byte
	NewImpl common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterImplChangeConfirmed is a free log retrieval operation binding the contract event 0x9d55b0349a0a4c5b511f72228170bb91d45c9ac78dba8ab5b4175d3ed42f06b3.
//
// Solidity: e ImplChangeConfirmed(_lockId bytes32, _newImpl address)
func (_ERC20Proxy *ERC20ProxyFilterer) FilterImplChangeConfirmed(opts *bind.FilterOpts) (*ERC20ProxyImplChangeConfirmedIterator, error) {

	logs, sub, err := _ERC20Proxy.contract.FilterLogs(opts, "ImplChangeConfirmed")
	if err != nil {
		return nil, err
	}
	return &ERC20ProxyImplChangeConfirmedIterator{contract: _ERC20Proxy.contract, event: "ImplChangeConfirmed", logs: logs, sub: sub}, nil
}

// WatchImplChangeConfirmed is a free log subscription operation binding the contract event 0x9d55b0349a0a4c5b511f72228170bb91d45c9ac78dba8ab5b4175d3ed42f06b3.
//
// Solidity: e ImplChangeConfirmed(_lockId bytes32, _newImpl address)
func (_ERC20Proxy *ERC20ProxyFilterer) WatchImplChangeConfirmed(opts *bind.WatchOpts, sink chan<- *ERC20ProxyImplChangeConfirmed) (event.Subscription, error) {

	logs, sub, err := _ERC20Proxy.contract.WatchLogs(opts, "ImplChangeConfirmed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20ProxyImplChangeConfirmed)
				if err := _ERC20Proxy.contract.UnpackLog(event, "ImplChangeConfirmed", log); err != nil {
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

// ERC20ProxyImplChangeRequestedIterator is returned from FilterImplChangeRequested and is used to iterate over the raw logs and unpacked data for ImplChangeRequested events raised by the ERC20Proxy contract.
type ERC20ProxyImplChangeRequestedIterator struct {
	Event *ERC20ProxyImplChangeRequested // Event containing the contract specifics and raw log

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
func (it *ERC20ProxyImplChangeRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20ProxyImplChangeRequested)
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
		it.Event = new(ERC20ProxyImplChangeRequested)
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
func (it *ERC20ProxyImplChangeRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ProxyImplChangeRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20ProxyImplChangeRequested represents a ImplChangeRequested event raised by the ERC20Proxy contract.
type ERC20ProxyImplChangeRequested struct {
	LockId       [32]byte
	MsgSender    common.Address
	ProposedImpl common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterImplChangeRequested is a free log retrieval operation binding the contract event 0x5df12834436b8dc248df3f7f1796a3e39f851d610be49cdcd92514fa821b9f97.
//
// Solidity: e ImplChangeRequested(_lockId bytes32, _msgSender address, _proposedImpl address)
func (_ERC20Proxy *ERC20ProxyFilterer) FilterImplChangeRequested(opts *bind.FilterOpts) (*ERC20ProxyImplChangeRequestedIterator, error) {

	logs, sub, err := _ERC20Proxy.contract.FilterLogs(opts, "ImplChangeRequested")
	if err != nil {
		return nil, err
	}
	return &ERC20ProxyImplChangeRequestedIterator{contract: _ERC20Proxy.contract, event: "ImplChangeRequested", logs: logs, sub: sub}, nil
}

// WatchImplChangeRequested is a free log subscription operation binding the contract event 0x5df12834436b8dc248df3f7f1796a3e39f851d610be49cdcd92514fa821b9f97.
//
// Solidity: e ImplChangeRequested(_lockId bytes32, _msgSender address, _proposedImpl address)
func (_ERC20Proxy *ERC20ProxyFilterer) WatchImplChangeRequested(opts *bind.WatchOpts, sink chan<- *ERC20ProxyImplChangeRequested) (event.Subscription, error) {

	logs, sub, err := _ERC20Proxy.contract.WatchLogs(opts, "ImplChangeRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20ProxyImplChangeRequested)
				if err := _ERC20Proxy.contract.UnpackLog(event, "ImplChangeRequested", log); err != nil {
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

// ERC20ProxyTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Proxy contract.
type ERC20ProxyTransferIterator struct {
	Event *ERC20ProxyTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20ProxyTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20ProxyTransfer)
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
		it.Event = new(ERC20ProxyTransfer)
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
func (it *ERC20ProxyTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ProxyTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20ProxyTransfer represents a Transfer event raised by the ERC20Proxy contract.
type ERC20ProxyTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _value uint256)
func (_ERC20Proxy *ERC20ProxyFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*ERC20ProxyTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ERC20Proxy.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ProxyTransferIterator{contract: _ERC20Proxy.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _value uint256)
func (_ERC20Proxy *ERC20ProxyFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20ProxyTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ERC20Proxy.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20ProxyTransfer)
				if err := _ERC20Proxy.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ERC20StoreABI is the input ABI used to generate the binding from.
const ERC20StoreABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_proposedCustodian\",\"type\":\"address\"}],\"name\":\"requestCustodianChange\",\"outputs\":[{\"name\":\"lockId\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_balanceIncrease\",\"type\":\"uint256\"}],\"name\":\"addBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"custodian\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_lockId\",\"type\":\"bytes32\"}],\"name\":\"confirmCustodianChange\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"erc20Impl\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proposedImpl\",\"type\":\"address\"}],\"name\":\"requestImplChange\",\"outputs\":[{\"name\":\"lockId\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_lockId\",\"type\":\"bytes32\"}],\"name\":\"confirmImplChange\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"implChangeReqs\",\"outputs\":[{\"name\":\"proposedNew\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lockRequestCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"custodianChangeReqs\",\"outputs\":[{\"name\":\"proposedNew\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setAllowance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_newBalance\",\"type\":\"uint256\"}],\"name\":\"setBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newTotalSupply\",\"type\":\"uint256\"}],\"name\":\"setTotalSupply\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_custodian\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_lockId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_msgSender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_proposedImpl\",\"type\":\"address\"}],\"name\":\"ImplChangeRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_lockId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_newImpl\",\"type\":\"address\"}],\"name\":\"ImplChangeConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_lockId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_msgSender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_proposedCustodian\",\"type\":\"address\"}],\"name\":\"CustodianChangeRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_lockId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_newCustodian\",\"type\":\"address\"}],\"name\":\"CustodianChangeConfirmed\",\"type\":\"event\"}]"

// ERC20StoreBin is the compiled bytecode used for deploying new contracts.
const ERC20StoreBin = `0x608060405234801561001057600080fd5b5060405160208061080a8339810160405251600080805560018054600160a060020a03909316600160a060020a0319938416179055600380549092169091556005556107a9806100616000396000f3006080604052600436106100e55763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166315b2108281146100ea57806318160ddd1461011d57806321e5383a1461013257806327e235e314610158578063375b74c3146101795780633a8343ee146101aa5780633c389cc4146101c257806348f9e246146101d75780635c658165146101f85780638181b0291461021f578063b508069b14610237578063cb81fecf1461024f578063cf6e448814610264578063da46098c1461027c578063e30443bc146102a6578063f7ea7a3d146102ca575b600080fd5b3480156100f657600080fd5b5061010b600160a060020a03600435166102e2565b60408051918252519081900360200190f35b34801561012957600080fd5b5061010b610395565b34801561013e57600080fd5b50610156600160a060020a036004351660243561039b565b005b34801561016457600080fd5b5061010b600160a060020a03600435166103d4565b34801561018557600080fd5b5061018e6103e6565b60408051600160a060020a039092168252519081900360200190f35b3480156101b657600080fd5b506101566004356103f5565b3480156101ce57600080fd5b5061018e610496565b3480156101e357600080fd5b5061010b600160a060020a03600435166104a5565b34801561020457600080fd5b5061010b600160a060020a0360043581169060243516610558565b34801561022b57600080fd5b50610156600435610575565b34801561024357600080fd5b5061018e600435610616565b34801561025b57600080fd5b5061010b610631565b34801561027057600080fd5b5061018e600435610637565b34801561028857600080fd5b50610156600160a060020a0360043581169060243516604435610652565b3480156102b257600080fd5b50610156600160a060020a0360043516602435610695565b3480156102d657600080fd5b506101566004356106c8565b6000600160a060020a03821615156102f957600080fd5b6103016106e4565b6040805160208181018352600160a060020a0380871680845260008681526002845285902093518454921673ffffffffffffffffffffffffffffffffffffffff1990921691909117909255825184815233918101919091528083019190915290519192507fd76fc900a7e1a6fcf11d54b7ba943918df6c53a3128140658c389b3da1e997ba919081900360600190a1919050565b60055481565b600354600160a060020a031633146103b257600080fd5b600160a060020a03909116600090815260066020526040902080549091019055565b60066020526000908152604090205481565b600154600160a060020a031681565b600154600160a060020a0316331461040c57600080fd5b61041581610725565b60018054600160a060020a0392831673ffffffffffffffffffffffffffffffffffffffff1991821617825560008481526002602090815260409182902080549093169092559154825185815293169083015280517f9a99272c0f6b7a30ef9e76e684a7cd408bfd4f11a72f36a8e276253c920e442d9281900390910190a150565b600354600160a060020a031681565b6000600160a060020a03821615156104bc57600080fd5b6104c46106e4565b6040805160208181018352600160a060020a0380871680845260008681526004845285902093518454921673ffffffffffffffffffffffffffffffffffffffff1990921691909117909255825184815233918101919091528083019190915290519192507f5df12834436b8dc248df3f7f1796a3e39f851d610be49cdcd92514fa821b9f97919081900360600190a1919050565b600760209081526000928352604080842090915290825290205481565b600154600160a060020a0316331461058c57600080fd5b61059581610759565b60038054600160a060020a0392831673ffffffffffffffffffffffffffffffffffffffff1991821617825560008481526004602090815260409182902080549093169092559154825185815293169083015280517f9d55b0349a0a4c5b511f72228170bb91d45c9ac78dba8ab5b4175d3ed42f06b39281900390910190a150565b600460205260009081526040902054600160a060020a031681565b60005481565b600260205260009081526040902054600160a060020a031681565b600354600160a060020a0316331461066957600080fd5b600160a060020a0392831660009081526007602090815260408083209490951682529290925291902055565b600354600160a060020a031633146106ac57600080fd5b600160a060020a03909116600090815260066020526040902055565b600354600160a060020a031633146106df57600080fd5b600555565b60008054600101908190556040805160001943014081526c010000000000000000000000003002602082015260348101929092525190819003605401902090565b60008181526002602052604081208054600160a060020a0316151561074957600080fd5b54600160a060020a031692915050565b60008181526004602052604081208054600160a060020a0316151561074957600080fd00a165627a7a72305820b87c68156dc2678e2898f9752f05314fca62dcd26a6d15b4aca4780a573522060029`

// DeployERC20Store deploys a new Ethereum contract, binding an instance of ERC20Store to it.
func DeployERC20Store(auth *bind.TransactOpts, backend bind.ContractBackend, _custodian common.Address) (common.Address, *types.Transaction, *ERC20Store, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20StoreABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20StoreBin), backend, _custodian)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Store{ERC20StoreCaller: ERC20StoreCaller{contract: contract}, ERC20StoreTransactor: ERC20StoreTransactor{contract: contract}, ERC20StoreFilterer: ERC20StoreFilterer{contract: contract}}, nil
}

// ERC20Store is an auto generated Go binding around an Ethereum contract.
type ERC20Store struct {
	ERC20StoreCaller     // Read-only binding to the contract
	ERC20StoreTransactor // Write-only binding to the contract
	ERC20StoreFilterer   // Log filterer for contract events
}

// ERC20StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20StoreSession struct {
	Contract     *ERC20Store       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20StoreCallerSession struct {
	Contract *ERC20StoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ERC20StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20StoreTransactorSession struct {
	Contract     *ERC20StoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ERC20StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20StoreRaw struct {
	Contract *ERC20Store // Generic contract binding to access the raw methods on
}

// ERC20StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20StoreCallerRaw struct {
	Contract *ERC20StoreCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20StoreTransactorRaw struct {
	Contract *ERC20StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Store creates a new instance of ERC20Store, bound to a specific deployed contract.
func NewERC20Store(address common.Address, backend bind.ContractBackend) (*ERC20Store, error) {
	contract, err := bindERC20Store(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Store{ERC20StoreCaller: ERC20StoreCaller{contract: contract}, ERC20StoreTransactor: ERC20StoreTransactor{contract: contract}, ERC20StoreFilterer: ERC20StoreFilterer{contract: contract}}, nil
}

// NewERC20StoreCaller creates a new read-only instance of ERC20Store, bound to a specific deployed contract.
func NewERC20StoreCaller(address common.Address, caller bind.ContractCaller) (*ERC20StoreCaller, error) {
	contract, err := bindERC20Store(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20StoreCaller{contract: contract}, nil
}

// NewERC20StoreTransactor creates a new write-only instance of ERC20Store, bound to a specific deployed contract.
func NewERC20StoreTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20StoreTransactor, error) {
	contract, err := bindERC20Store(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20StoreTransactor{contract: contract}, nil
}

// NewERC20StoreFilterer creates a new log filterer instance of ERC20Store, bound to a specific deployed contract.
func NewERC20StoreFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20StoreFilterer, error) {
	contract, err := bindERC20Store(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20StoreFilterer{contract: contract}, nil
}

// bindERC20Store binds a generic wrapper to an already deployed contract.
func bindERC20Store(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20StoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Store *ERC20StoreRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Store.Contract.ERC20StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Store *ERC20StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Store.Contract.ERC20StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Store *ERC20StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Store.Contract.ERC20StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Store *ERC20StoreCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Store *ERC20StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Store *ERC20StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Store.Contract.contract.Transact(opts, method, params...)
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed( address,  address) constant returns(uint256)
func (_ERC20Store *ERC20StoreCaller) Allowed(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Store.contract.Call(opts, out, "allowed", arg0, arg1)
	return *ret0, err
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed( address,  address) constant returns(uint256)
func (_ERC20Store *ERC20StoreSession) Allowed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _ERC20Store.Contract.Allowed(&_ERC20Store.CallOpts, arg0, arg1)
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed( address,  address) constant returns(uint256)
func (_ERC20Store *ERC20StoreCallerSession) Allowed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _ERC20Store.Contract.Allowed(&_ERC20Store.CallOpts, arg0, arg1)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_ERC20Store *ERC20StoreCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Store.contract.Call(opts, out, "balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_ERC20Store *ERC20StoreSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _ERC20Store.Contract.Balances(&_ERC20Store.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_ERC20Store *ERC20StoreCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _ERC20Store.Contract.Balances(&_ERC20Store.CallOpts, arg0)
}

// Custodian is a free data retrieval call binding the contract method 0x375b74c3.
//
// Solidity: function custodian() constant returns(address)
func (_ERC20Store *ERC20StoreCaller) Custodian(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC20Store.contract.Call(opts, out, "custodian")
	return *ret0, err
}

// Custodian is a free data retrieval call binding the contract method 0x375b74c3.
//
// Solidity: function custodian() constant returns(address)
func (_ERC20Store *ERC20StoreSession) Custodian() (common.Address, error) {
	return _ERC20Store.Contract.Custodian(&_ERC20Store.CallOpts)
}

// Custodian is a free data retrieval call binding the contract method 0x375b74c3.
//
// Solidity: function custodian() constant returns(address)
func (_ERC20Store *ERC20StoreCallerSession) Custodian() (common.Address, error) {
	return _ERC20Store.Contract.Custodian(&_ERC20Store.CallOpts)
}

// CustodianChangeReqs is a free data retrieval call binding the contract method 0xcf6e4488.
//
// Solidity: function custodianChangeReqs( bytes32) constant returns(proposedNew address)
func (_ERC20Store *ERC20StoreCaller) CustodianChangeReqs(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC20Store.contract.Call(opts, out, "custodianChangeReqs", arg0)
	return *ret0, err
}

// CustodianChangeReqs is a free data retrieval call binding the contract method 0xcf6e4488.
//
// Solidity: function custodianChangeReqs( bytes32) constant returns(proposedNew address)
func (_ERC20Store *ERC20StoreSession) CustodianChangeReqs(arg0 [32]byte) (common.Address, error) {
	return _ERC20Store.Contract.CustodianChangeReqs(&_ERC20Store.CallOpts, arg0)
}

// CustodianChangeReqs is a free data retrieval call binding the contract method 0xcf6e4488.
//
// Solidity: function custodianChangeReqs( bytes32) constant returns(proposedNew address)
func (_ERC20Store *ERC20StoreCallerSession) CustodianChangeReqs(arg0 [32]byte) (common.Address, error) {
	return _ERC20Store.Contract.CustodianChangeReqs(&_ERC20Store.CallOpts, arg0)
}

// Erc20Impl is a free data retrieval call binding the contract method 0x3c389cc4.
//
// Solidity: function erc20Impl() constant returns(address)
func (_ERC20Store *ERC20StoreCaller) Erc20Impl(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC20Store.contract.Call(opts, out, "erc20Impl")
	return *ret0, err
}

// Erc20Impl is a free data retrieval call binding the contract method 0x3c389cc4.
//
// Solidity: function erc20Impl() constant returns(address)
func (_ERC20Store *ERC20StoreSession) Erc20Impl() (common.Address, error) {
	return _ERC20Store.Contract.Erc20Impl(&_ERC20Store.CallOpts)
}

// Erc20Impl is a free data retrieval call binding the contract method 0x3c389cc4.
//
// Solidity: function erc20Impl() constant returns(address)
func (_ERC20Store *ERC20StoreCallerSession) Erc20Impl() (common.Address, error) {
	return _ERC20Store.Contract.Erc20Impl(&_ERC20Store.CallOpts)
}

// ImplChangeReqs is a free data retrieval call binding the contract method 0xb508069b.
//
// Solidity: function implChangeReqs( bytes32) constant returns(proposedNew address)
func (_ERC20Store *ERC20StoreCaller) ImplChangeReqs(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC20Store.contract.Call(opts, out, "implChangeReqs", arg0)
	return *ret0, err
}

// ImplChangeReqs is a free data retrieval call binding the contract method 0xb508069b.
//
// Solidity: function implChangeReqs( bytes32) constant returns(proposedNew address)
func (_ERC20Store *ERC20StoreSession) ImplChangeReqs(arg0 [32]byte) (common.Address, error) {
	return _ERC20Store.Contract.ImplChangeReqs(&_ERC20Store.CallOpts, arg0)
}

// ImplChangeReqs is a free data retrieval call binding the contract method 0xb508069b.
//
// Solidity: function implChangeReqs( bytes32) constant returns(proposedNew address)
func (_ERC20Store *ERC20StoreCallerSession) ImplChangeReqs(arg0 [32]byte) (common.Address, error) {
	return _ERC20Store.Contract.ImplChangeReqs(&_ERC20Store.CallOpts, arg0)
}

// LockRequestCount is a free data retrieval call binding the contract method 0xcb81fecf.
//
// Solidity: function lockRequestCount() constant returns(uint256)
func (_ERC20Store *ERC20StoreCaller) LockRequestCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Store.contract.Call(opts, out, "lockRequestCount")
	return *ret0, err
}

// LockRequestCount is a free data retrieval call binding the contract method 0xcb81fecf.
//
// Solidity: function lockRequestCount() constant returns(uint256)
func (_ERC20Store *ERC20StoreSession) LockRequestCount() (*big.Int, error) {
	return _ERC20Store.Contract.LockRequestCount(&_ERC20Store.CallOpts)
}

// LockRequestCount is a free data retrieval call binding the contract method 0xcb81fecf.
//
// Solidity: function lockRequestCount() constant returns(uint256)
func (_ERC20Store *ERC20StoreCallerSession) LockRequestCount() (*big.Int, error) {
	return _ERC20Store.Contract.LockRequestCount(&_ERC20Store.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Store *ERC20StoreCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Store.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Store *ERC20StoreSession) TotalSupply() (*big.Int, error) {
	return _ERC20Store.Contract.TotalSupply(&_ERC20Store.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Store *ERC20StoreCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Store.Contract.TotalSupply(&_ERC20Store.CallOpts)
}

// AddBalance is a paid mutator transaction binding the contract method 0x21e5383a.
//
// Solidity: function addBalance(_owner address, _balanceIncrease uint256) returns()
func (_ERC20Store *ERC20StoreTransactor) AddBalance(opts *bind.TransactOpts, _owner common.Address, _balanceIncrease *big.Int) (*types.Transaction, error) {
	return _ERC20Store.contract.Transact(opts, "addBalance", _owner, _balanceIncrease)
}

// AddBalance is a paid mutator transaction binding the contract method 0x21e5383a.
//
// Solidity: function addBalance(_owner address, _balanceIncrease uint256) returns()
func (_ERC20Store *ERC20StoreSession) AddBalance(_owner common.Address, _balanceIncrease *big.Int) (*types.Transaction, error) {
	return _ERC20Store.Contract.AddBalance(&_ERC20Store.TransactOpts, _owner, _balanceIncrease)
}

// AddBalance is a paid mutator transaction binding the contract method 0x21e5383a.
//
// Solidity: function addBalance(_owner address, _balanceIncrease uint256) returns()
func (_ERC20Store *ERC20StoreTransactorSession) AddBalance(_owner common.Address, _balanceIncrease *big.Int) (*types.Transaction, error) {
	return _ERC20Store.Contract.AddBalance(&_ERC20Store.TransactOpts, _owner, _balanceIncrease)
}

// ConfirmCustodianChange is a paid mutator transaction binding the contract method 0x3a8343ee.
//
// Solidity: function confirmCustodianChange(_lockId bytes32) returns()
func (_ERC20Store *ERC20StoreTransactor) ConfirmCustodianChange(opts *bind.TransactOpts, _lockId [32]byte) (*types.Transaction, error) {
	return _ERC20Store.contract.Transact(opts, "confirmCustodianChange", _lockId)
}

// ConfirmCustodianChange is a paid mutator transaction binding the contract method 0x3a8343ee.
//
// Solidity: function confirmCustodianChange(_lockId bytes32) returns()
func (_ERC20Store *ERC20StoreSession) ConfirmCustodianChange(_lockId [32]byte) (*types.Transaction, error) {
	return _ERC20Store.Contract.ConfirmCustodianChange(&_ERC20Store.TransactOpts, _lockId)
}

// ConfirmCustodianChange is a paid mutator transaction binding the contract method 0x3a8343ee.
//
// Solidity: function confirmCustodianChange(_lockId bytes32) returns()
func (_ERC20Store *ERC20StoreTransactorSession) ConfirmCustodianChange(_lockId [32]byte) (*types.Transaction, error) {
	return _ERC20Store.Contract.ConfirmCustodianChange(&_ERC20Store.TransactOpts, _lockId)
}

// ConfirmImplChange is a paid mutator transaction binding the contract method 0x8181b029.
//
// Solidity: function confirmImplChange(_lockId bytes32) returns()
func (_ERC20Store *ERC20StoreTransactor) ConfirmImplChange(opts *bind.TransactOpts, _lockId [32]byte) (*types.Transaction, error) {
	return _ERC20Store.contract.Transact(opts, "confirmImplChange", _lockId)
}

// ConfirmImplChange is a paid mutator transaction binding the contract method 0x8181b029.
//
// Solidity: function confirmImplChange(_lockId bytes32) returns()
func (_ERC20Store *ERC20StoreSession) ConfirmImplChange(_lockId [32]byte) (*types.Transaction, error) {
	return _ERC20Store.Contract.ConfirmImplChange(&_ERC20Store.TransactOpts, _lockId)
}

// ConfirmImplChange is a paid mutator transaction binding the contract method 0x8181b029.
//
// Solidity: function confirmImplChange(_lockId bytes32) returns()
func (_ERC20Store *ERC20StoreTransactorSession) ConfirmImplChange(_lockId [32]byte) (*types.Transaction, error) {
	return _ERC20Store.Contract.ConfirmImplChange(&_ERC20Store.TransactOpts, _lockId)
}

// RequestCustodianChange is a paid mutator transaction binding the contract method 0x15b21082.
//
// Solidity: function requestCustodianChange(_proposedCustodian address) returns(lockId bytes32)
func (_ERC20Store *ERC20StoreTransactor) RequestCustodianChange(opts *bind.TransactOpts, _proposedCustodian common.Address) (*types.Transaction, error) {
	return _ERC20Store.contract.Transact(opts, "requestCustodianChange", _proposedCustodian)
}

// RequestCustodianChange is a paid mutator transaction binding the contract method 0x15b21082.
//
// Solidity: function requestCustodianChange(_proposedCustodian address) returns(lockId bytes32)
func (_ERC20Store *ERC20StoreSession) RequestCustodianChange(_proposedCustodian common.Address) (*types.Transaction, error) {
	return _ERC20Store.Contract.RequestCustodianChange(&_ERC20Store.TransactOpts, _proposedCustodian)
}

// RequestCustodianChange is a paid mutator transaction binding the contract method 0x15b21082.
//
// Solidity: function requestCustodianChange(_proposedCustodian address) returns(lockId bytes32)
func (_ERC20Store *ERC20StoreTransactorSession) RequestCustodianChange(_proposedCustodian common.Address) (*types.Transaction, error) {
	return _ERC20Store.Contract.RequestCustodianChange(&_ERC20Store.TransactOpts, _proposedCustodian)
}

// RequestImplChange is a paid mutator transaction binding the contract method 0x48f9e246.
//
// Solidity: function requestImplChange(_proposedImpl address) returns(lockId bytes32)
func (_ERC20Store *ERC20StoreTransactor) RequestImplChange(opts *bind.TransactOpts, _proposedImpl common.Address) (*types.Transaction, error) {
	return _ERC20Store.contract.Transact(opts, "requestImplChange", _proposedImpl)
}

// RequestImplChange is a paid mutator transaction binding the contract method 0x48f9e246.
//
// Solidity: function requestImplChange(_proposedImpl address) returns(lockId bytes32)
func (_ERC20Store *ERC20StoreSession) RequestImplChange(_proposedImpl common.Address) (*types.Transaction, error) {
	return _ERC20Store.Contract.RequestImplChange(&_ERC20Store.TransactOpts, _proposedImpl)
}

// RequestImplChange is a paid mutator transaction binding the contract method 0x48f9e246.
//
// Solidity: function requestImplChange(_proposedImpl address) returns(lockId bytes32)
func (_ERC20Store *ERC20StoreTransactorSession) RequestImplChange(_proposedImpl common.Address) (*types.Transaction, error) {
	return _ERC20Store.Contract.RequestImplChange(&_ERC20Store.TransactOpts, _proposedImpl)
}

// SetAllowance is a paid mutator transaction binding the contract method 0xda46098c.
//
// Solidity: function setAllowance(_owner address, _spender address, _value uint256) returns()
func (_ERC20Store *ERC20StoreTransactor) SetAllowance(opts *bind.TransactOpts, _owner common.Address, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Store.contract.Transact(opts, "setAllowance", _owner, _spender, _value)
}

// SetAllowance is a paid mutator transaction binding the contract method 0xda46098c.
//
// Solidity: function setAllowance(_owner address, _spender address, _value uint256) returns()
func (_ERC20Store *ERC20StoreSession) SetAllowance(_owner common.Address, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Store.Contract.SetAllowance(&_ERC20Store.TransactOpts, _owner, _spender, _value)
}

// SetAllowance is a paid mutator transaction binding the contract method 0xda46098c.
//
// Solidity: function setAllowance(_owner address, _spender address, _value uint256) returns()
func (_ERC20Store *ERC20StoreTransactorSession) SetAllowance(_owner common.Address, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Store.Contract.SetAllowance(&_ERC20Store.TransactOpts, _owner, _spender, _value)
}

// SetBalance is a paid mutator transaction binding the contract method 0xe30443bc.
//
// Solidity: function setBalance(_owner address, _newBalance uint256) returns()
func (_ERC20Store *ERC20StoreTransactor) SetBalance(opts *bind.TransactOpts, _owner common.Address, _newBalance *big.Int) (*types.Transaction, error) {
	return _ERC20Store.contract.Transact(opts, "setBalance", _owner, _newBalance)
}

// SetBalance is a paid mutator transaction binding the contract method 0xe30443bc.
//
// Solidity: function setBalance(_owner address, _newBalance uint256) returns()
func (_ERC20Store *ERC20StoreSession) SetBalance(_owner common.Address, _newBalance *big.Int) (*types.Transaction, error) {
	return _ERC20Store.Contract.SetBalance(&_ERC20Store.TransactOpts, _owner, _newBalance)
}

// SetBalance is a paid mutator transaction binding the contract method 0xe30443bc.
//
// Solidity: function setBalance(_owner address, _newBalance uint256) returns()
func (_ERC20Store *ERC20StoreTransactorSession) SetBalance(_owner common.Address, _newBalance *big.Int) (*types.Transaction, error) {
	return _ERC20Store.Contract.SetBalance(&_ERC20Store.TransactOpts, _owner, _newBalance)
}

// SetTotalSupply is a paid mutator transaction binding the contract method 0xf7ea7a3d.
//
// Solidity: function setTotalSupply(_newTotalSupply uint256) returns()
func (_ERC20Store *ERC20StoreTransactor) SetTotalSupply(opts *bind.TransactOpts, _newTotalSupply *big.Int) (*types.Transaction, error) {
	return _ERC20Store.contract.Transact(opts, "setTotalSupply", _newTotalSupply)
}

// SetTotalSupply is a paid mutator transaction binding the contract method 0xf7ea7a3d.
//
// Solidity: function setTotalSupply(_newTotalSupply uint256) returns()
func (_ERC20Store *ERC20StoreSession) SetTotalSupply(_newTotalSupply *big.Int) (*types.Transaction, error) {
	return _ERC20Store.Contract.SetTotalSupply(&_ERC20Store.TransactOpts, _newTotalSupply)
}

// SetTotalSupply is a paid mutator transaction binding the contract method 0xf7ea7a3d.
//
// Solidity: function setTotalSupply(_newTotalSupply uint256) returns()
func (_ERC20Store *ERC20StoreTransactorSession) SetTotalSupply(_newTotalSupply *big.Int) (*types.Transaction, error) {
	return _ERC20Store.Contract.SetTotalSupply(&_ERC20Store.TransactOpts, _newTotalSupply)
}

// ERC20StoreCustodianChangeConfirmedIterator is returned from FilterCustodianChangeConfirmed and is used to iterate over the raw logs and unpacked data for CustodianChangeConfirmed events raised by the ERC20Store contract.
type ERC20StoreCustodianChangeConfirmedIterator struct {
	Event *ERC20StoreCustodianChangeConfirmed // Event containing the contract specifics and raw log

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
func (it *ERC20StoreCustodianChangeConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20StoreCustodianChangeConfirmed)
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
		it.Event = new(ERC20StoreCustodianChangeConfirmed)
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
func (it *ERC20StoreCustodianChangeConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20StoreCustodianChangeConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20StoreCustodianChangeConfirmed represents a CustodianChangeConfirmed event raised by the ERC20Store contract.
type ERC20StoreCustodianChangeConfirmed struct {
	LockId       [32]byte
	NewCustodian common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCustodianChangeConfirmed is a free log retrieval operation binding the contract event 0x9a99272c0f6b7a30ef9e76e684a7cd408bfd4f11a72f36a8e276253c920e442d.
//
// Solidity: e CustodianChangeConfirmed(_lockId bytes32, _newCustodian address)
func (_ERC20Store *ERC20StoreFilterer) FilterCustodianChangeConfirmed(opts *bind.FilterOpts) (*ERC20StoreCustodianChangeConfirmedIterator, error) {

	logs, sub, err := _ERC20Store.contract.FilterLogs(opts, "CustodianChangeConfirmed")
	if err != nil {
		return nil, err
	}
	return &ERC20StoreCustodianChangeConfirmedIterator{contract: _ERC20Store.contract, event: "CustodianChangeConfirmed", logs: logs, sub: sub}, nil
}

// WatchCustodianChangeConfirmed is a free log subscription operation binding the contract event 0x9a99272c0f6b7a30ef9e76e684a7cd408bfd4f11a72f36a8e276253c920e442d.
//
// Solidity: e CustodianChangeConfirmed(_lockId bytes32, _newCustodian address)
func (_ERC20Store *ERC20StoreFilterer) WatchCustodianChangeConfirmed(opts *bind.WatchOpts, sink chan<- *ERC20StoreCustodianChangeConfirmed) (event.Subscription, error) {

	logs, sub, err := _ERC20Store.contract.WatchLogs(opts, "CustodianChangeConfirmed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20StoreCustodianChangeConfirmed)
				if err := _ERC20Store.contract.UnpackLog(event, "CustodianChangeConfirmed", log); err != nil {
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

// ERC20StoreCustodianChangeRequestedIterator is returned from FilterCustodianChangeRequested and is used to iterate over the raw logs and unpacked data for CustodianChangeRequested events raised by the ERC20Store contract.
type ERC20StoreCustodianChangeRequestedIterator struct {
	Event *ERC20StoreCustodianChangeRequested // Event containing the contract specifics and raw log

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
func (it *ERC20StoreCustodianChangeRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20StoreCustodianChangeRequested)
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
		it.Event = new(ERC20StoreCustodianChangeRequested)
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
func (it *ERC20StoreCustodianChangeRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20StoreCustodianChangeRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20StoreCustodianChangeRequested represents a CustodianChangeRequested event raised by the ERC20Store contract.
type ERC20StoreCustodianChangeRequested struct {
	LockId            [32]byte
	MsgSender         common.Address
	ProposedCustodian common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCustodianChangeRequested is a free log retrieval operation binding the contract event 0xd76fc900a7e1a6fcf11d54b7ba943918df6c53a3128140658c389b3da1e997ba.
//
// Solidity: e CustodianChangeRequested(_lockId bytes32, _msgSender address, _proposedCustodian address)
func (_ERC20Store *ERC20StoreFilterer) FilterCustodianChangeRequested(opts *bind.FilterOpts) (*ERC20StoreCustodianChangeRequestedIterator, error) {

	logs, sub, err := _ERC20Store.contract.FilterLogs(opts, "CustodianChangeRequested")
	if err != nil {
		return nil, err
	}
	return &ERC20StoreCustodianChangeRequestedIterator{contract: _ERC20Store.contract, event: "CustodianChangeRequested", logs: logs, sub: sub}, nil
}

// WatchCustodianChangeRequested is a free log subscription operation binding the contract event 0xd76fc900a7e1a6fcf11d54b7ba943918df6c53a3128140658c389b3da1e997ba.
//
// Solidity: e CustodianChangeRequested(_lockId bytes32, _msgSender address, _proposedCustodian address)
func (_ERC20Store *ERC20StoreFilterer) WatchCustodianChangeRequested(opts *bind.WatchOpts, sink chan<- *ERC20StoreCustodianChangeRequested) (event.Subscription, error) {

	logs, sub, err := _ERC20Store.contract.WatchLogs(opts, "CustodianChangeRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20StoreCustodianChangeRequested)
				if err := _ERC20Store.contract.UnpackLog(event, "CustodianChangeRequested", log); err != nil {
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

// ERC20StoreImplChangeConfirmedIterator is returned from FilterImplChangeConfirmed and is used to iterate over the raw logs and unpacked data for ImplChangeConfirmed events raised by the ERC20Store contract.
type ERC20StoreImplChangeConfirmedIterator struct {
	Event *ERC20StoreImplChangeConfirmed // Event containing the contract specifics and raw log

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
func (it *ERC20StoreImplChangeConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20StoreImplChangeConfirmed)
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
		it.Event = new(ERC20StoreImplChangeConfirmed)
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
func (it *ERC20StoreImplChangeConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20StoreImplChangeConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20StoreImplChangeConfirmed represents a ImplChangeConfirmed event raised by the ERC20Store contract.
type ERC20StoreImplChangeConfirmed struct {
	LockId  [32]byte
	NewImpl common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterImplChangeConfirmed is a free log retrieval operation binding the contract event 0x9d55b0349a0a4c5b511f72228170bb91d45c9ac78dba8ab5b4175d3ed42f06b3.
//
// Solidity: e ImplChangeConfirmed(_lockId bytes32, _newImpl address)
func (_ERC20Store *ERC20StoreFilterer) FilterImplChangeConfirmed(opts *bind.FilterOpts) (*ERC20StoreImplChangeConfirmedIterator, error) {

	logs, sub, err := _ERC20Store.contract.FilterLogs(opts, "ImplChangeConfirmed")
	if err != nil {
		return nil, err
	}
	return &ERC20StoreImplChangeConfirmedIterator{contract: _ERC20Store.contract, event: "ImplChangeConfirmed", logs: logs, sub: sub}, nil
}

// WatchImplChangeConfirmed is a free log subscription operation binding the contract event 0x9d55b0349a0a4c5b511f72228170bb91d45c9ac78dba8ab5b4175d3ed42f06b3.
//
// Solidity: e ImplChangeConfirmed(_lockId bytes32, _newImpl address)
func (_ERC20Store *ERC20StoreFilterer) WatchImplChangeConfirmed(opts *bind.WatchOpts, sink chan<- *ERC20StoreImplChangeConfirmed) (event.Subscription, error) {

	logs, sub, err := _ERC20Store.contract.WatchLogs(opts, "ImplChangeConfirmed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20StoreImplChangeConfirmed)
				if err := _ERC20Store.contract.UnpackLog(event, "ImplChangeConfirmed", log); err != nil {
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

// ERC20StoreImplChangeRequestedIterator is returned from FilterImplChangeRequested and is used to iterate over the raw logs and unpacked data for ImplChangeRequested events raised by the ERC20Store contract.
type ERC20StoreImplChangeRequestedIterator struct {
	Event *ERC20StoreImplChangeRequested // Event containing the contract specifics and raw log

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
func (it *ERC20StoreImplChangeRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20StoreImplChangeRequested)
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
		it.Event = new(ERC20StoreImplChangeRequested)
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
func (it *ERC20StoreImplChangeRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20StoreImplChangeRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20StoreImplChangeRequested represents a ImplChangeRequested event raised by the ERC20Store contract.
type ERC20StoreImplChangeRequested struct {
	LockId       [32]byte
	MsgSender    common.Address
	ProposedImpl common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterImplChangeRequested is a free log retrieval operation binding the contract event 0x5df12834436b8dc248df3f7f1796a3e39f851d610be49cdcd92514fa821b9f97.
//
// Solidity: e ImplChangeRequested(_lockId bytes32, _msgSender address, _proposedImpl address)
func (_ERC20Store *ERC20StoreFilterer) FilterImplChangeRequested(opts *bind.FilterOpts) (*ERC20StoreImplChangeRequestedIterator, error) {

	logs, sub, err := _ERC20Store.contract.FilterLogs(opts, "ImplChangeRequested")
	if err != nil {
		return nil, err
	}
	return &ERC20StoreImplChangeRequestedIterator{contract: _ERC20Store.contract, event: "ImplChangeRequested", logs: logs, sub: sub}, nil
}

// WatchImplChangeRequested is a free log subscription operation binding the contract event 0x5df12834436b8dc248df3f7f1796a3e39f851d610be49cdcd92514fa821b9f97.
//
// Solidity: e ImplChangeRequested(_lockId bytes32, _msgSender address, _proposedImpl address)
func (_ERC20Store *ERC20StoreFilterer) WatchImplChangeRequested(opts *bind.WatchOpts, sink chan<- *ERC20StoreImplChangeRequested) (event.Subscription, error) {

	logs, sub, err := _ERC20Store.contract.WatchLogs(opts, "ImplChangeRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20StoreImplChangeRequested)
				if err := _ERC20Store.contract.UnpackLog(event, "ImplChangeRequested", log); err != nil {
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

// LockRequestableABI is the input ABI used to generate the binding from.
const LockRequestableABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"lockRequestCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// LockRequestableBin is the compiled bytecode used for deploying new contracts.
const LockRequestableBin = `0x6080604052348015600f57600080fd5b50600080556099806100226000396000f300608060405260043610603e5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663cb81fecf81146043575b600080fd5b348015604e57600080fd5b5060556067565b60408051918252519081900360200190f35b600054815600a165627a7a72305820b12546a16fe74bec63d8f6781ba2e113cb60185f95e47ee88648fba81a23c1b00029`

// DeployLockRequestable deploys a new Ethereum contract, binding an instance of LockRequestable to it.
func DeployLockRequestable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LockRequestable, error) {
	parsed, err := abi.JSON(strings.NewReader(LockRequestableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LockRequestableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LockRequestable{LockRequestableCaller: LockRequestableCaller{contract: contract}, LockRequestableTransactor: LockRequestableTransactor{contract: contract}, LockRequestableFilterer: LockRequestableFilterer{contract: contract}}, nil
}

// LockRequestable is an auto generated Go binding around an Ethereum contract.
type LockRequestable struct {
	LockRequestableCaller     // Read-only binding to the contract
	LockRequestableTransactor // Write-only binding to the contract
	LockRequestableFilterer   // Log filterer for contract events
}

// LockRequestableCaller is an auto generated read-only Go binding around an Ethereum contract.
type LockRequestableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockRequestableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LockRequestableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockRequestableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LockRequestableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockRequestableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LockRequestableSession struct {
	Contract     *LockRequestable  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LockRequestableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LockRequestableCallerSession struct {
	Contract *LockRequestableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// LockRequestableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LockRequestableTransactorSession struct {
	Contract     *LockRequestableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// LockRequestableRaw is an auto generated low-level Go binding around an Ethereum contract.
type LockRequestableRaw struct {
	Contract *LockRequestable // Generic contract binding to access the raw methods on
}

// LockRequestableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LockRequestableCallerRaw struct {
	Contract *LockRequestableCaller // Generic read-only contract binding to access the raw methods on
}

// LockRequestableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LockRequestableTransactorRaw struct {
	Contract *LockRequestableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLockRequestable creates a new instance of LockRequestable, bound to a specific deployed contract.
func NewLockRequestable(address common.Address, backend bind.ContractBackend) (*LockRequestable, error) {
	contract, err := bindLockRequestable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LockRequestable{LockRequestableCaller: LockRequestableCaller{contract: contract}, LockRequestableTransactor: LockRequestableTransactor{contract: contract}, LockRequestableFilterer: LockRequestableFilterer{contract: contract}}, nil
}

// NewLockRequestableCaller creates a new read-only instance of LockRequestable, bound to a specific deployed contract.
func NewLockRequestableCaller(address common.Address, caller bind.ContractCaller) (*LockRequestableCaller, error) {
	contract, err := bindLockRequestable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LockRequestableCaller{contract: contract}, nil
}

// NewLockRequestableTransactor creates a new write-only instance of LockRequestable, bound to a specific deployed contract.
func NewLockRequestableTransactor(address common.Address, transactor bind.ContractTransactor) (*LockRequestableTransactor, error) {
	contract, err := bindLockRequestable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LockRequestableTransactor{contract: contract}, nil
}

// NewLockRequestableFilterer creates a new log filterer instance of LockRequestable, bound to a specific deployed contract.
func NewLockRequestableFilterer(address common.Address, filterer bind.ContractFilterer) (*LockRequestableFilterer, error) {
	contract, err := bindLockRequestable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LockRequestableFilterer{contract: contract}, nil
}

// bindLockRequestable binds a generic wrapper to an already deployed contract.
func bindLockRequestable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LockRequestableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LockRequestable *LockRequestableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LockRequestable.Contract.LockRequestableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LockRequestable *LockRequestableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockRequestable.Contract.LockRequestableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LockRequestable *LockRequestableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LockRequestable.Contract.LockRequestableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LockRequestable *LockRequestableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LockRequestable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LockRequestable *LockRequestableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockRequestable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LockRequestable *LockRequestableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LockRequestable.Contract.contract.Transact(opts, method, params...)
}

// LockRequestCount is a free data retrieval call binding the contract method 0xcb81fecf.
//
// Solidity: function lockRequestCount() constant returns(uint256)
func (_LockRequestable *LockRequestableCaller) LockRequestCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LockRequestable.contract.Call(opts, out, "lockRequestCount")
	return *ret0, err
}

// LockRequestCount is a free data retrieval call binding the contract method 0xcb81fecf.
//
// Solidity: function lockRequestCount() constant returns(uint256)
func (_LockRequestable *LockRequestableSession) LockRequestCount() (*big.Int, error) {
	return _LockRequestable.Contract.LockRequestCount(&_LockRequestable.CallOpts)
}

// LockRequestCount is a free data retrieval call binding the contract method 0xcb81fecf.
//
// Solidity: function lockRequestCount() constant returns(uint256)
func (_LockRequestable *LockRequestableCallerSession) LockRequestCount() (*big.Int, error) {
	return _LockRequestable.Contract.LockRequestCount(&_LockRequestable.CallOpts)
}

// PrintLimiterABI is the input ABI used to generate the binding from.
const PrintLimiterABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_lockId\",\"type\":\"bytes32\"}],\"name\":\"confirmCustodianChangeProxy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_lockId\",\"type\":\"bytes32\"}],\"name\":\"confirmCeilingRaise\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"pendingRaiseMap\",\"outputs\":[{\"name\":\"raiseBy\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_lowerBy\",\"type\":\"uint256\"}],\"name\":\"lowerCeiling\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"custodian\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"erc20Impl\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"limitedPrinter\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_lockId\",\"type\":\"bytes32\"}],\"name\":\"confirmPrintProxy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupplyCeiling\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"limitedPrint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_raiseBy\",\"type\":\"uint256\"}],\"name\":\"requestCeilingRaise\",\"outputs\":[{\"name\":\"lockId\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lockRequestCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_erc20Impl\",\"type\":\"address\"},{\"name\":\"_custodian\",\"type\":\"address\"},{\"name\":\"_limitedPrinter\",\"type\":\"address\"},{\"name\":\"_initialCeiling\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_lockId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_raiseBy\",\"type\":\"uint256\"}],\"name\":\"CeilingRaiseLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_lockId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_raiseBy\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_newCeiling\",\"type\":\"uint256\"}],\"name\":\"CeilingRaiseConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_lowerBy\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_newCeiling\",\"type\":\"uint256\"}],\"name\":\"CeilingLowered\",\"type\":\"event\"}]"

// PrintLimiterBin is the compiled bytecode used for deploying new contracts.
const PrintLimiterBin = `0x608060405234801561001057600080fd5b5060405160808061079283398101604090815281516020830151918301516060909301516000805560018054600160a060020a03938416600160a060020a0319918216179091556002805494841694821694909417909355600380549290941691909216179091556004556107088061008a6000396000f3006080604052600436106100a05763ffffffff60e060020a600035041663051d995581146100a557806312460fdd146100bf578063264da19e146100d75780633351373914610101578063375b74c3146101195780633c389cc41461014a57806342569ff31461015f5780638cbf41451461017457806395a512331461018c578063a94c7c65146101a1578063c9bc5dbd146101c5578063cb81fecf146101dd575b600080fd5b3480156100b157600080fd5b506100bd6004356101f2565b005b3480156100cb57600080fd5b506100bd60043561028a565b3480156100e357600080fd5b506100ef600435610332565b60408051918252519081900360200190f35b34801561010d57600080fd5b506100bd600435610344565b34801561012557600080fd5b5061012e6103b6565b60408051600160a060020a039092168252519081900360200190f35b34801561015657600080fd5b5061012e6103c5565b34801561016b57600080fd5b5061012e6103d4565b34801561018057600080fd5b506100bd6004356103e3565b34801561019857600080fd5b506100ef610460565b3480156101ad57600080fd5b506100bd600160a060020a0360043516602435610466565b3480156101d157600080fd5b506100ef600435610620565b3480156101e957600080fd5b506100ef610695565b600254600160a060020a0316331461020957600080fd5b600154604080517f3a8343ee000000000000000000000000000000000000000000000000000000008152600481018490529051600160a060020a0390921691633a8343ee9160248082019260009290919082900301818387803b15801561026f57600080fd5b505af1158015610283573d6000803e3d6000fd5b5050505050565b60025460009081908190600160a060020a031633146102a857600080fd5b6000848152600560205260409020805490935091508115156102c957600080fd5b5060008381526005602052604081205560045480820190811061032c576004819055604080518581526020810184905280820183905290517f3779579f908c94480802e9c026789736b5c6a0382577c15f94feb7e23eafcfd29181900360600190a15b50505050565b60056020526000908152604090205481565b600354600090600160a060020a0316331461035e57600080fd5b506004548181039081111561037257600080fd5b6004819055604080518381526020810183905281517f76f283ae36a7f1bcca5e9475a544804430c59279a8b72325d56fb5ade52ab632929181900390910190a15050565b600254600160a060020a031681565b600154600160a060020a031681565b600354600160a060020a031681565b600254600160a060020a031633146103fa57600080fd5b600154604080517f380ba30c000000000000000000000000000000000000000000000000000000008152600481018490529051600160a060020a039092169163380ba30c9160248082019260009290919082900301818387803b15801561026f57600080fd5b60045481565b6003546000908190600160a060020a0316331461048257600080fd5b600160009054906101000a9004600160a060020a0316600160a060020a03166318160ddd6040518163ffffffff1660e060020a028152600401602060405180830381600087803b1580156104d557600080fd5b505af11580156104e9573d6000803e3d6000fd5b505050506040513d60208110156104ff57600080fd5b50519150508181018181101561051457600080fd5b60045481111561052357600080fd5b600154604080517fbe23d291000000000000000000000000000000000000000000000000000000008152600160a060020a038781166004830152602482018790529151919092169163380ba30c91839163be23d2919160448083019260209291908290030181600087803b15801561059a57600080fd5b505af11580156105ae573d6000803e3d6000fd5b505050506040513d60208110156105c457600080fd5b50516040805160e060020a63ffffffff8516028152600481019290925251602480830192600092919082900301818387803b15801561060257600080fd5b505af1158015610616573d6000803e3d6000fd5b5050505050505050565b600081151561062e57600080fd5b61063661069b565b60408051602081810183528582526000848152600582528390209151909155815183815290810185905281519293507f938f54351a24c466787f903e29b0590fe701da10a63db9b78c2a10ad8a8e1f67929081900390910190a1919050565b60005481565b60008054600101908190556040805160001943014081526c0100000000000000000000000030026020820152603481019290925251908190036054019020905600a165627a7a723058204984a630886cd024aa37ab4cb078507718a1e791f0faa493adcd6080ab8040010029`

// DeployPrintLimiter deploys a new Ethereum contract, binding an instance of PrintLimiter to it.
func DeployPrintLimiter(auth *bind.TransactOpts, backend bind.ContractBackend, _erc20Impl common.Address, _custodian common.Address, _limitedPrinter common.Address, _initialCeiling *big.Int) (common.Address, *types.Transaction, *PrintLimiter, error) {
	parsed, err := abi.JSON(strings.NewReader(PrintLimiterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PrintLimiterBin), backend, _erc20Impl, _custodian, _limitedPrinter, _initialCeiling)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PrintLimiter{PrintLimiterCaller: PrintLimiterCaller{contract: contract}, PrintLimiterTransactor: PrintLimiterTransactor{contract: contract}, PrintLimiterFilterer: PrintLimiterFilterer{contract: contract}}, nil
}

// PrintLimiter is an auto generated Go binding around an Ethereum contract.
type PrintLimiter struct {
	PrintLimiterCaller     // Read-only binding to the contract
	PrintLimiterTransactor // Write-only binding to the contract
	PrintLimiterFilterer   // Log filterer for contract events
}

// PrintLimiterCaller is an auto generated read-only Go binding around an Ethereum contract.
type PrintLimiterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrintLimiterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PrintLimiterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrintLimiterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PrintLimiterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrintLimiterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PrintLimiterSession struct {
	Contract     *PrintLimiter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PrintLimiterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PrintLimiterCallerSession struct {
	Contract *PrintLimiterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PrintLimiterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PrintLimiterTransactorSession struct {
	Contract     *PrintLimiterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PrintLimiterRaw is an auto generated low-level Go binding around an Ethereum contract.
type PrintLimiterRaw struct {
	Contract *PrintLimiter // Generic contract binding to access the raw methods on
}

// PrintLimiterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PrintLimiterCallerRaw struct {
	Contract *PrintLimiterCaller // Generic read-only contract binding to access the raw methods on
}

// PrintLimiterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PrintLimiterTransactorRaw struct {
	Contract *PrintLimiterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPrintLimiter creates a new instance of PrintLimiter, bound to a specific deployed contract.
func NewPrintLimiter(address common.Address, backend bind.ContractBackend) (*PrintLimiter, error) {
	contract, err := bindPrintLimiter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PrintLimiter{PrintLimiterCaller: PrintLimiterCaller{contract: contract}, PrintLimiterTransactor: PrintLimiterTransactor{contract: contract}, PrintLimiterFilterer: PrintLimiterFilterer{contract: contract}}, nil
}

// NewPrintLimiterCaller creates a new read-only instance of PrintLimiter, bound to a specific deployed contract.
func NewPrintLimiterCaller(address common.Address, caller bind.ContractCaller) (*PrintLimiterCaller, error) {
	contract, err := bindPrintLimiter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PrintLimiterCaller{contract: contract}, nil
}

// NewPrintLimiterTransactor creates a new write-only instance of PrintLimiter, bound to a specific deployed contract.
func NewPrintLimiterTransactor(address common.Address, transactor bind.ContractTransactor) (*PrintLimiterTransactor, error) {
	contract, err := bindPrintLimiter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PrintLimiterTransactor{contract: contract}, nil
}

// NewPrintLimiterFilterer creates a new log filterer instance of PrintLimiter, bound to a specific deployed contract.
func NewPrintLimiterFilterer(address common.Address, filterer bind.ContractFilterer) (*PrintLimiterFilterer, error) {
	contract, err := bindPrintLimiter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PrintLimiterFilterer{contract: contract}, nil
}

// bindPrintLimiter binds a generic wrapper to an already deployed contract.
func bindPrintLimiter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PrintLimiterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PrintLimiter *PrintLimiterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PrintLimiter.Contract.PrintLimiterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PrintLimiter *PrintLimiterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrintLimiter.Contract.PrintLimiterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PrintLimiter *PrintLimiterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PrintLimiter.Contract.PrintLimiterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PrintLimiter *PrintLimiterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PrintLimiter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PrintLimiter *PrintLimiterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrintLimiter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PrintLimiter *PrintLimiterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PrintLimiter.Contract.contract.Transact(opts, method, params...)
}

// Custodian is a free data retrieval call binding the contract method 0x375b74c3.
//
// Solidity: function custodian() constant returns(address)
func (_PrintLimiter *PrintLimiterCaller) Custodian(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PrintLimiter.contract.Call(opts, out, "custodian")
	return *ret0, err
}

// Custodian is a free data retrieval call binding the contract method 0x375b74c3.
//
// Solidity: function custodian() constant returns(address)
func (_PrintLimiter *PrintLimiterSession) Custodian() (common.Address, error) {
	return _PrintLimiter.Contract.Custodian(&_PrintLimiter.CallOpts)
}

// Custodian is a free data retrieval call binding the contract method 0x375b74c3.
//
// Solidity: function custodian() constant returns(address)
func (_PrintLimiter *PrintLimiterCallerSession) Custodian() (common.Address, error) {
	return _PrintLimiter.Contract.Custodian(&_PrintLimiter.CallOpts)
}

// Erc20Impl is a free data retrieval call binding the contract method 0x3c389cc4.
//
// Solidity: function erc20Impl() constant returns(address)
func (_PrintLimiter *PrintLimiterCaller) Erc20Impl(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PrintLimiter.contract.Call(opts, out, "erc20Impl")
	return *ret0, err
}

// Erc20Impl is a free data retrieval call binding the contract method 0x3c389cc4.
//
// Solidity: function erc20Impl() constant returns(address)
func (_PrintLimiter *PrintLimiterSession) Erc20Impl() (common.Address, error) {
	return _PrintLimiter.Contract.Erc20Impl(&_PrintLimiter.CallOpts)
}

// Erc20Impl is a free data retrieval call binding the contract method 0x3c389cc4.
//
// Solidity: function erc20Impl() constant returns(address)
func (_PrintLimiter *PrintLimiterCallerSession) Erc20Impl() (common.Address, error) {
	return _PrintLimiter.Contract.Erc20Impl(&_PrintLimiter.CallOpts)
}

// LimitedPrinter is a free data retrieval call binding the contract method 0x42569ff3.
//
// Solidity: function limitedPrinter() constant returns(address)
func (_PrintLimiter *PrintLimiterCaller) LimitedPrinter(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PrintLimiter.contract.Call(opts, out, "limitedPrinter")
	return *ret0, err
}

// LimitedPrinter is a free data retrieval call binding the contract method 0x42569ff3.
//
// Solidity: function limitedPrinter() constant returns(address)
func (_PrintLimiter *PrintLimiterSession) LimitedPrinter() (common.Address, error) {
	return _PrintLimiter.Contract.LimitedPrinter(&_PrintLimiter.CallOpts)
}

// LimitedPrinter is a free data retrieval call binding the contract method 0x42569ff3.
//
// Solidity: function limitedPrinter() constant returns(address)
func (_PrintLimiter *PrintLimiterCallerSession) LimitedPrinter() (common.Address, error) {
	return _PrintLimiter.Contract.LimitedPrinter(&_PrintLimiter.CallOpts)
}

// LockRequestCount is a free data retrieval call binding the contract method 0xcb81fecf.
//
// Solidity: function lockRequestCount() constant returns(uint256)
func (_PrintLimiter *PrintLimiterCaller) LockRequestCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PrintLimiter.contract.Call(opts, out, "lockRequestCount")
	return *ret0, err
}

// LockRequestCount is a free data retrieval call binding the contract method 0xcb81fecf.
//
// Solidity: function lockRequestCount() constant returns(uint256)
func (_PrintLimiter *PrintLimiterSession) LockRequestCount() (*big.Int, error) {
	return _PrintLimiter.Contract.LockRequestCount(&_PrintLimiter.CallOpts)
}

// LockRequestCount is a free data retrieval call binding the contract method 0xcb81fecf.
//
// Solidity: function lockRequestCount() constant returns(uint256)
func (_PrintLimiter *PrintLimiterCallerSession) LockRequestCount() (*big.Int, error) {
	return _PrintLimiter.Contract.LockRequestCount(&_PrintLimiter.CallOpts)
}

// PendingRaiseMap is a free data retrieval call binding the contract method 0x264da19e.
//
// Solidity: function pendingRaiseMap( bytes32) constant returns(raiseBy uint256)
func (_PrintLimiter *PrintLimiterCaller) PendingRaiseMap(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PrintLimiter.contract.Call(opts, out, "pendingRaiseMap", arg0)
	return *ret0, err
}

// PendingRaiseMap is a free data retrieval call binding the contract method 0x264da19e.
//
// Solidity: function pendingRaiseMap( bytes32) constant returns(raiseBy uint256)
func (_PrintLimiter *PrintLimiterSession) PendingRaiseMap(arg0 [32]byte) (*big.Int, error) {
	return _PrintLimiter.Contract.PendingRaiseMap(&_PrintLimiter.CallOpts, arg0)
}

// PendingRaiseMap is a free data retrieval call binding the contract method 0x264da19e.
//
// Solidity: function pendingRaiseMap( bytes32) constant returns(raiseBy uint256)
func (_PrintLimiter *PrintLimiterCallerSession) PendingRaiseMap(arg0 [32]byte) (*big.Int, error) {
	return _PrintLimiter.Contract.PendingRaiseMap(&_PrintLimiter.CallOpts, arg0)
}

// TotalSupplyCeiling is a free data retrieval call binding the contract method 0x95a51233.
//
// Solidity: function totalSupplyCeiling() constant returns(uint256)
func (_PrintLimiter *PrintLimiterCaller) TotalSupplyCeiling(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PrintLimiter.contract.Call(opts, out, "totalSupplyCeiling")
	return *ret0, err
}

// TotalSupplyCeiling is a free data retrieval call binding the contract method 0x95a51233.
//
// Solidity: function totalSupplyCeiling() constant returns(uint256)
func (_PrintLimiter *PrintLimiterSession) TotalSupplyCeiling() (*big.Int, error) {
	return _PrintLimiter.Contract.TotalSupplyCeiling(&_PrintLimiter.CallOpts)
}

// TotalSupplyCeiling is a free data retrieval call binding the contract method 0x95a51233.
//
// Solidity: function totalSupplyCeiling() constant returns(uint256)
func (_PrintLimiter *PrintLimiterCallerSession) TotalSupplyCeiling() (*big.Int, error) {
	return _PrintLimiter.Contract.TotalSupplyCeiling(&_PrintLimiter.CallOpts)
}

// ConfirmCeilingRaise is a paid mutator transaction binding the contract method 0x12460fdd.
//
// Solidity: function confirmCeilingRaise(_lockId bytes32) returns()
func (_PrintLimiter *PrintLimiterTransactor) ConfirmCeilingRaise(opts *bind.TransactOpts, _lockId [32]byte) (*types.Transaction, error) {
	return _PrintLimiter.contract.Transact(opts, "confirmCeilingRaise", _lockId)
}

// ConfirmCeilingRaise is a paid mutator transaction binding the contract method 0x12460fdd.
//
// Solidity: function confirmCeilingRaise(_lockId bytes32) returns()
func (_PrintLimiter *PrintLimiterSession) ConfirmCeilingRaise(_lockId [32]byte) (*types.Transaction, error) {
	return _PrintLimiter.Contract.ConfirmCeilingRaise(&_PrintLimiter.TransactOpts, _lockId)
}

// ConfirmCeilingRaise is a paid mutator transaction binding the contract method 0x12460fdd.
//
// Solidity: function confirmCeilingRaise(_lockId bytes32) returns()
func (_PrintLimiter *PrintLimiterTransactorSession) ConfirmCeilingRaise(_lockId [32]byte) (*types.Transaction, error) {
	return _PrintLimiter.Contract.ConfirmCeilingRaise(&_PrintLimiter.TransactOpts, _lockId)
}

// ConfirmCustodianChangeProxy is a paid mutator transaction binding the contract method 0x051d9955.
//
// Solidity: function confirmCustodianChangeProxy(_lockId bytes32) returns()
func (_PrintLimiter *PrintLimiterTransactor) ConfirmCustodianChangeProxy(opts *bind.TransactOpts, _lockId [32]byte) (*types.Transaction, error) {
	return _PrintLimiter.contract.Transact(opts, "confirmCustodianChangeProxy", _lockId)
}

// ConfirmCustodianChangeProxy is a paid mutator transaction binding the contract method 0x051d9955.
//
// Solidity: function confirmCustodianChangeProxy(_lockId bytes32) returns()
func (_PrintLimiter *PrintLimiterSession) ConfirmCustodianChangeProxy(_lockId [32]byte) (*types.Transaction, error) {
	return _PrintLimiter.Contract.ConfirmCustodianChangeProxy(&_PrintLimiter.TransactOpts, _lockId)
}

// ConfirmCustodianChangeProxy is a paid mutator transaction binding the contract method 0x051d9955.
//
// Solidity: function confirmCustodianChangeProxy(_lockId bytes32) returns()
func (_PrintLimiter *PrintLimiterTransactorSession) ConfirmCustodianChangeProxy(_lockId [32]byte) (*types.Transaction, error) {
	return _PrintLimiter.Contract.ConfirmCustodianChangeProxy(&_PrintLimiter.TransactOpts, _lockId)
}

// ConfirmPrintProxy is a paid mutator transaction binding the contract method 0x8cbf4145.
//
// Solidity: function confirmPrintProxy(_lockId bytes32) returns()
func (_PrintLimiter *PrintLimiterTransactor) ConfirmPrintProxy(opts *bind.TransactOpts, _lockId [32]byte) (*types.Transaction, error) {
	return _PrintLimiter.contract.Transact(opts, "confirmPrintProxy", _lockId)
}

// ConfirmPrintProxy is a paid mutator transaction binding the contract method 0x8cbf4145.
//
// Solidity: function confirmPrintProxy(_lockId bytes32) returns()
func (_PrintLimiter *PrintLimiterSession) ConfirmPrintProxy(_lockId [32]byte) (*types.Transaction, error) {
	return _PrintLimiter.Contract.ConfirmPrintProxy(&_PrintLimiter.TransactOpts, _lockId)
}

// ConfirmPrintProxy is a paid mutator transaction binding the contract method 0x8cbf4145.
//
// Solidity: function confirmPrintProxy(_lockId bytes32) returns()
func (_PrintLimiter *PrintLimiterTransactorSession) ConfirmPrintProxy(_lockId [32]byte) (*types.Transaction, error) {
	return _PrintLimiter.Contract.ConfirmPrintProxy(&_PrintLimiter.TransactOpts, _lockId)
}

// LimitedPrint is a paid mutator transaction binding the contract method 0xa94c7c65.
//
// Solidity: function limitedPrint(_receiver address, _value uint256) returns()
func (_PrintLimiter *PrintLimiterTransactor) LimitedPrint(opts *bind.TransactOpts, _receiver common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PrintLimiter.contract.Transact(opts, "limitedPrint", _receiver, _value)
}

// LimitedPrint is a paid mutator transaction binding the contract method 0xa94c7c65.
//
// Solidity: function limitedPrint(_receiver address, _value uint256) returns()
func (_PrintLimiter *PrintLimiterSession) LimitedPrint(_receiver common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PrintLimiter.Contract.LimitedPrint(&_PrintLimiter.TransactOpts, _receiver, _value)
}

// LimitedPrint is a paid mutator transaction binding the contract method 0xa94c7c65.
//
// Solidity: function limitedPrint(_receiver address, _value uint256) returns()
func (_PrintLimiter *PrintLimiterTransactorSession) LimitedPrint(_receiver common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PrintLimiter.Contract.LimitedPrint(&_PrintLimiter.TransactOpts, _receiver, _value)
}

// LowerCeiling is a paid mutator transaction binding the contract method 0x33513739.
//
// Solidity: function lowerCeiling(_lowerBy uint256) returns()
func (_PrintLimiter *PrintLimiterTransactor) LowerCeiling(opts *bind.TransactOpts, _lowerBy *big.Int) (*types.Transaction, error) {
	return _PrintLimiter.contract.Transact(opts, "lowerCeiling", _lowerBy)
}

// LowerCeiling is a paid mutator transaction binding the contract method 0x33513739.
//
// Solidity: function lowerCeiling(_lowerBy uint256) returns()
func (_PrintLimiter *PrintLimiterSession) LowerCeiling(_lowerBy *big.Int) (*types.Transaction, error) {
	return _PrintLimiter.Contract.LowerCeiling(&_PrintLimiter.TransactOpts, _lowerBy)
}

// LowerCeiling is a paid mutator transaction binding the contract method 0x33513739.
//
// Solidity: function lowerCeiling(_lowerBy uint256) returns()
func (_PrintLimiter *PrintLimiterTransactorSession) LowerCeiling(_lowerBy *big.Int) (*types.Transaction, error) {
	return _PrintLimiter.Contract.LowerCeiling(&_PrintLimiter.TransactOpts, _lowerBy)
}

// RequestCeilingRaise is a paid mutator transaction binding the contract method 0xc9bc5dbd.
//
// Solidity: function requestCeilingRaise(_raiseBy uint256) returns(lockId bytes32)
func (_PrintLimiter *PrintLimiterTransactor) RequestCeilingRaise(opts *bind.TransactOpts, _raiseBy *big.Int) (*types.Transaction, error) {
	return _PrintLimiter.contract.Transact(opts, "requestCeilingRaise", _raiseBy)
}

// RequestCeilingRaise is a paid mutator transaction binding the contract method 0xc9bc5dbd.
//
// Solidity: function requestCeilingRaise(_raiseBy uint256) returns(lockId bytes32)
func (_PrintLimiter *PrintLimiterSession) RequestCeilingRaise(_raiseBy *big.Int) (*types.Transaction, error) {
	return _PrintLimiter.Contract.RequestCeilingRaise(&_PrintLimiter.TransactOpts, _raiseBy)
}

// RequestCeilingRaise is a paid mutator transaction binding the contract method 0xc9bc5dbd.
//
// Solidity: function requestCeilingRaise(_raiseBy uint256) returns(lockId bytes32)
func (_PrintLimiter *PrintLimiterTransactorSession) RequestCeilingRaise(_raiseBy *big.Int) (*types.Transaction, error) {
	return _PrintLimiter.Contract.RequestCeilingRaise(&_PrintLimiter.TransactOpts, _raiseBy)
}

// PrintLimiterCeilingLoweredIterator is returned from FilterCeilingLowered and is used to iterate over the raw logs and unpacked data for CeilingLowered events raised by the PrintLimiter contract.
type PrintLimiterCeilingLoweredIterator struct {
	Event *PrintLimiterCeilingLowered // Event containing the contract specifics and raw log

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
func (it *PrintLimiterCeilingLoweredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrintLimiterCeilingLowered)
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
		it.Event = new(PrintLimiterCeilingLowered)
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
func (it *PrintLimiterCeilingLoweredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrintLimiterCeilingLoweredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrintLimiterCeilingLowered represents a CeilingLowered event raised by the PrintLimiter contract.
type PrintLimiterCeilingLowered struct {
	LowerBy    *big.Int
	NewCeiling *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCeilingLowered is a free log retrieval operation binding the contract event 0x76f283ae36a7f1bcca5e9475a544804430c59279a8b72325d56fb5ade52ab632.
//
// Solidity: e CeilingLowered(_lowerBy uint256, _newCeiling uint256)
func (_PrintLimiter *PrintLimiterFilterer) FilterCeilingLowered(opts *bind.FilterOpts) (*PrintLimiterCeilingLoweredIterator, error) {

	logs, sub, err := _PrintLimiter.contract.FilterLogs(opts, "CeilingLowered")
	if err != nil {
		return nil, err
	}
	return &PrintLimiterCeilingLoweredIterator{contract: _PrintLimiter.contract, event: "CeilingLowered", logs: logs, sub: sub}, nil
}

// WatchCeilingLowered is a free log subscription operation binding the contract event 0x76f283ae36a7f1bcca5e9475a544804430c59279a8b72325d56fb5ade52ab632.
//
// Solidity: e CeilingLowered(_lowerBy uint256, _newCeiling uint256)
func (_PrintLimiter *PrintLimiterFilterer) WatchCeilingLowered(opts *bind.WatchOpts, sink chan<- *PrintLimiterCeilingLowered) (event.Subscription, error) {

	logs, sub, err := _PrintLimiter.contract.WatchLogs(opts, "CeilingLowered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrintLimiterCeilingLowered)
				if err := _PrintLimiter.contract.UnpackLog(event, "CeilingLowered", log); err != nil {
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

// PrintLimiterCeilingRaiseConfirmedIterator is returned from FilterCeilingRaiseConfirmed and is used to iterate over the raw logs and unpacked data for CeilingRaiseConfirmed events raised by the PrintLimiter contract.
type PrintLimiterCeilingRaiseConfirmedIterator struct {
	Event *PrintLimiterCeilingRaiseConfirmed // Event containing the contract specifics and raw log

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
func (it *PrintLimiterCeilingRaiseConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrintLimiterCeilingRaiseConfirmed)
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
		it.Event = new(PrintLimiterCeilingRaiseConfirmed)
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
func (it *PrintLimiterCeilingRaiseConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrintLimiterCeilingRaiseConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrintLimiterCeilingRaiseConfirmed represents a CeilingRaiseConfirmed event raised by the PrintLimiter contract.
type PrintLimiterCeilingRaiseConfirmed struct {
	LockId     [32]byte
	RaiseBy    *big.Int
	NewCeiling *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCeilingRaiseConfirmed is a free log retrieval operation binding the contract event 0x3779579f908c94480802e9c026789736b5c6a0382577c15f94feb7e23eafcfd2.
//
// Solidity: e CeilingRaiseConfirmed(_lockId bytes32, _raiseBy uint256, _newCeiling uint256)
func (_PrintLimiter *PrintLimiterFilterer) FilterCeilingRaiseConfirmed(opts *bind.FilterOpts) (*PrintLimiterCeilingRaiseConfirmedIterator, error) {

	logs, sub, err := _PrintLimiter.contract.FilterLogs(opts, "CeilingRaiseConfirmed")
	if err != nil {
		return nil, err
	}
	return &PrintLimiterCeilingRaiseConfirmedIterator{contract: _PrintLimiter.contract, event: "CeilingRaiseConfirmed", logs: logs, sub: sub}, nil
}

// WatchCeilingRaiseConfirmed is a free log subscription operation binding the contract event 0x3779579f908c94480802e9c026789736b5c6a0382577c15f94feb7e23eafcfd2.
//
// Solidity: e CeilingRaiseConfirmed(_lockId bytes32, _raiseBy uint256, _newCeiling uint256)
func (_PrintLimiter *PrintLimiterFilterer) WatchCeilingRaiseConfirmed(opts *bind.WatchOpts, sink chan<- *PrintLimiterCeilingRaiseConfirmed) (event.Subscription, error) {

	logs, sub, err := _PrintLimiter.contract.WatchLogs(opts, "CeilingRaiseConfirmed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrintLimiterCeilingRaiseConfirmed)
				if err := _PrintLimiter.contract.UnpackLog(event, "CeilingRaiseConfirmed", log); err != nil {
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

// PrintLimiterCeilingRaiseLockedIterator is returned from FilterCeilingRaiseLocked and is used to iterate over the raw logs and unpacked data for CeilingRaiseLocked events raised by the PrintLimiter contract.
type PrintLimiterCeilingRaiseLockedIterator struct {
	Event *PrintLimiterCeilingRaiseLocked // Event containing the contract specifics and raw log

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
func (it *PrintLimiterCeilingRaiseLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrintLimiterCeilingRaiseLocked)
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
		it.Event = new(PrintLimiterCeilingRaiseLocked)
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
func (it *PrintLimiterCeilingRaiseLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrintLimiterCeilingRaiseLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrintLimiterCeilingRaiseLocked represents a CeilingRaiseLocked event raised by the PrintLimiter contract.
type PrintLimiterCeilingRaiseLocked struct {
	LockId  [32]byte
	RaiseBy *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCeilingRaiseLocked is a free log retrieval operation binding the contract event 0x938f54351a24c466787f903e29b0590fe701da10a63db9b78c2a10ad8a8e1f67.
//
// Solidity: e CeilingRaiseLocked(_lockId bytes32, _raiseBy uint256)
func (_PrintLimiter *PrintLimiterFilterer) FilterCeilingRaiseLocked(opts *bind.FilterOpts) (*PrintLimiterCeilingRaiseLockedIterator, error) {

	logs, sub, err := _PrintLimiter.contract.FilterLogs(opts, "CeilingRaiseLocked")
	if err != nil {
		return nil, err
	}
	return &PrintLimiterCeilingRaiseLockedIterator{contract: _PrintLimiter.contract, event: "CeilingRaiseLocked", logs: logs, sub: sub}, nil
}

// WatchCeilingRaiseLocked is a free log subscription operation binding the contract event 0x938f54351a24c466787f903e29b0590fe701da10a63db9b78c2a10ad8a8e1f67.
//
// Solidity: e CeilingRaiseLocked(_lockId bytes32, _raiseBy uint256)
func (_PrintLimiter *PrintLimiterFilterer) WatchCeilingRaiseLocked(opts *bind.WatchOpts, sink chan<- *PrintLimiterCeilingRaiseLocked) (event.Subscription, error) {

	logs, sub, err := _PrintLimiter.contract.WatchLogs(opts, "CeilingRaiseLocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrintLimiterCeilingRaiseLocked)
				if err := _PrintLimiter.contract.UnpackLog(event, "CeilingRaiseLocked", log); err != nil {
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
