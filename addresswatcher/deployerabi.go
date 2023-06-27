// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package addresswatcher

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

// AddresswatcherMetaData contains all meta data concerning the Addresswatcher contract.
var AddresswatcherMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"NFTAddress\",\"type\":\"address\"}],\"name\":\"TokenDeployed\",\"type\":\"event\"}]",
}

// AddresswatcherABI is the input ABI used to generate the binding from.
// Deprecated: Use AddresswatcherMetaData.ABI instead.
var AddresswatcherABI = AddresswatcherMetaData.ABI

// Addresswatcher is an auto generated Go binding around an Ethereum contract.
type Addresswatcher struct {
	AddresswatcherCaller     // Read-only binding to the contract
	AddresswatcherTransactor // Write-only binding to the contract
	AddresswatcherFilterer   // Log filterer for contract events
}

// AddresswatcherCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddresswatcherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddresswatcherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddresswatcherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddresswatcherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddresswatcherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddresswatcherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddresswatcherSession struct {
	Contract     *Addresswatcher   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddresswatcherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddresswatcherCallerSession struct {
	Contract *AddresswatcherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// AddresswatcherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddresswatcherTransactorSession struct {
	Contract     *AddresswatcherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AddresswatcherRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddresswatcherRaw struct {
	Contract *Addresswatcher // Generic contract binding to access the raw methods on
}

// AddresswatcherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddresswatcherCallerRaw struct {
	Contract *AddresswatcherCaller // Generic read-only contract binding to access the raw methods on
}

// AddresswatcherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddresswatcherTransactorRaw struct {
	Contract *AddresswatcherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddresswatcher creates a new instance of Addresswatcher, bound to a specific deployed contract.
func NewAddresswatcher(address common.Address, backend bind.ContractBackend) (*Addresswatcher, error) {
	contract, err := bindAddresswatcher(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Addresswatcher{AddresswatcherCaller: AddresswatcherCaller{contract: contract}, AddresswatcherTransactor: AddresswatcherTransactor{contract: contract}, AddresswatcherFilterer: AddresswatcherFilterer{contract: contract}}, nil
}

// NewAddresswatcherCaller creates a new read-only instance of Addresswatcher, bound to a specific deployed contract.
func NewAddresswatcherCaller(address common.Address, caller bind.ContractCaller) (*AddresswatcherCaller, error) {
	contract, err := bindAddresswatcher(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddresswatcherCaller{contract: contract}, nil
}

// NewAddresswatcherTransactor creates a new write-only instance of Addresswatcher, bound to a specific deployed contract.
func NewAddresswatcherTransactor(address common.Address, transactor bind.ContractTransactor) (*AddresswatcherTransactor, error) {
	contract, err := bindAddresswatcher(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddresswatcherTransactor{contract: contract}, nil
}

// NewAddresswatcherFilterer creates a new log filterer instance of Addresswatcher, bound to a specific deployed contract.
func NewAddresswatcherFilterer(address common.Address, filterer bind.ContractFilterer) (*AddresswatcherFilterer, error) {
	contract, err := bindAddresswatcher(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddresswatcherFilterer{contract: contract}, nil
}

// bindAddresswatcher binds a generic wrapper to an already deployed contract.
func bindAddresswatcher(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AddresswatcherMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Addresswatcher *AddresswatcherRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Addresswatcher.Contract.AddresswatcherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Addresswatcher *AddresswatcherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Addresswatcher.Contract.AddresswatcherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Addresswatcher *AddresswatcherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Addresswatcher.Contract.AddresswatcherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Addresswatcher *AddresswatcherCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Addresswatcher.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Addresswatcher *AddresswatcherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Addresswatcher.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Addresswatcher *AddresswatcherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Addresswatcher.Contract.contract.Transact(opts, method, params...)
}

// AddresswatcherTokenDeployedIterator is returned from FilterTokenDeployed and is used to iterate over the raw logs and unpacked data for TokenDeployed events raised by the Addresswatcher contract.
type AddresswatcherTokenDeployedIterator struct {
	Event *AddresswatcherTokenDeployed // Event containing the contract specifics and raw log

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
func (it *AddresswatcherTokenDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddresswatcherTokenDeployed)
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
		it.Event = new(AddresswatcherTokenDeployed)
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
func (it *AddresswatcherTokenDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddresswatcherTokenDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddresswatcherTokenDeployed represents a TokenDeployed event raised by the Addresswatcher contract.
type AddresswatcherTokenDeployed struct {
	NFTAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTokenDeployed is a free log retrieval operation binding the contract event 0x91d24864a084ab70b268a1f865e757ca12006cf298d763b6be697302ef86498c.
//
// Solidity: event TokenDeployed(address NFTAddress)
func (_Addresswatcher *AddresswatcherFilterer) FilterTokenDeployed(opts *bind.FilterOpts) (*AddresswatcherTokenDeployedIterator, error) {

	logs, sub, err := _Addresswatcher.contract.FilterLogs(opts, "TokenDeployed")
	if err != nil {
		return nil, err
	}
	return &AddresswatcherTokenDeployedIterator{contract: _Addresswatcher.contract, event: "TokenDeployed", logs: logs, sub: sub}, nil
}

// WatchTokenDeployed is a free log subscription operation binding the contract event 0x91d24864a084ab70b268a1f865e757ca12006cf298d763b6be697302ef86498c.
//
// Solidity: event TokenDeployed(address NFTAddress)
func (_Addresswatcher *AddresswatcherFilterer) WatchTokenDeployed(opts *bind.WatchOpts, sink chan<- *AddresswatcherTokenDeployed) (event.Subscription, error) {

	logs, sub, err := _Addresswatcher.contract.WatchLogs(opts, "TokenDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddresswatcherTokenDeployed)
				if err := _Addresswatcher.contract.UnpackLog(event, "TokenDeployed", log); err != nil {
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

// ParseTokenDeployed is a log parse operation binding the contract event 0x91d24864a084ab70b268a1f865e757ca12006cf298d763b6be697302ef86498c.
//
// Solidity: event TokenDeployed(address NFTAddress)
func (_Addresswatcher *AddresswatcherFilterer) ParseTokenDeployed(log types.Log) (*AddresswatcherTokenDeployed, error) {
	event := new(AddresswatcherTokenDeployed)
	if err := _Addresswatcher.contract.UnpackLog(event, "TokenDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
