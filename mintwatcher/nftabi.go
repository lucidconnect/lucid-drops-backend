// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mintwatcher

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

var LocalTestWatcher = &Mintwatcher{}

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
)

// MintwatcherMetaData contains all meta data concerning the Mintwatcher contract.
var MintwatcherMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_collectionName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_collectionSymbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_collectionOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"_batchURIs\",\"type\":\"string[]\"}],\"name\":\"addItems\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"batchBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"batchMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"customMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"itemIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"receiptientAddresses\",\"type\":\"address[]\"}],\"name\":\"sendNFTs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uris\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MintwatcherABI is the input ABI used to generate the binding from.
// Deprecated: Use MintwatcherMetaData.ABI instead.
var MintwatcherABI = MintwatcherMetaData.ABI

// Mintwatcher is an auto generated Go binding around an Ethereum contract.
type Mintwatcher struct {
	MintwatcherCaller     // Read-only binding to the contract
	MintwatcherTransactor // Write-only binding to the contract
	MintwatcherFilterer   // Log filterer for contract events
}

// MintwatcherCaller is an auto generated read-only Go binding around an Ethereum contract.
type MintwatcherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintwatcherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MintwatcherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintwatcherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MintwatcherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintwatcherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MintwatcherSession struct {
	Contract     *Mintwatcher      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MintwatcherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MintwatcherCallerSession struct {
	Contract *MintwatcherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MintwatcherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MintwatcherTransactorSession struct {
	Contract     *MintwatcherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MintwatcherRaw is an auto generated low-level Go binding around an Ethereum contract.
type MintwatcherRaw struct {
	Contract *Mintwatcher // Generic contract binding to access the raw methods on
}

// MintwatcherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MintwatcherCallerRaw struct {
	Contract *MintwatcherCaller // Generic read-only contract binding to access the raw methods on
}

// MintwatcherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MintwatcherTransactorRaw struct {
	Contract *MintwatcherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMintwatcher creates a new instance of Mintwatcher, bound to a specific deployed contract.
func NewMintwatcher(address common.Address, backend bind.ContractBackend) (*Mintwatcher, error) {
	contract, err := bindMintwatcher(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mintwatcher{MintwatcherCaller: MintwatcherCaller{contract: contract}, MintwatcherTransactor: MintwatcherTransactor{contract: contract}, MintwatcherFilterer: MintwatcherFilterer{contract: contract}}, nil
}

// NewMintwatcherCaller creates a new read-only instance of Mintwatcher, bound to a specific deployed contract.
func NewMintwatcherCaller(address common.Address, caller bind.ContractCaller) (*MintwatcherCaller, error) {
	contract, err := bindMintwatcher(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MintwatcherCaller{contract: contract}, nil
}

// NewMintwatcherTransactor creates a new write-only instance of Mintwatcher, bound to a specific deployed contract.
func NewMintwatcherTransactor(address common.Address, transactor bind.ContractTransactor) (*MintwatcherTransactor, error) {
	contract, err := bindMintwatcher(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MintwatcherTransactor{contract: contract}, nil
}

// NewMintwatcherFilterer creates a new log filterer instance of Mintwatcher, bound to a specific deployed contract.
func NewMintwatcherFilterer(address common.Address, filterer bind.ContractFilterer) (*MintwatcherFilterer, error) {
	contract, err := bindMintwatcher(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MintwatcherFilterer{contract: contract}, nil
}

// bindMintwatcher binds a generic wrapper to an already deployed contract.
func bindMintwatcher(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MintwatcherABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mintwatcher *MintwatcherRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mintwatcher.Contract.MintwatcherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mintwatcher *MintwatcherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mintwatcher.Contract.MintwatcherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mintwatcher *MintwatcherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mintwatcher.Contract.MintwatcherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mintwatcher *MintwatcherCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mintwatcher.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mintwatcher *MintwatcherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mintwatcher.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mintwatcher *MintwatcherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mintwatcher.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address , uint256 ) view returns(uint256)
func (_Mintwatcher *MintwatcherCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Mintwatcher.contract.Call(opts, &out, "balanceOf", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address , uint256 ) view returns(uint256)
func (_Mintwatcher *MintwatcherSession) BalanceOf(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Mintwatcher.Contract.BalanceOf(&_Mintwatcher.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address , uint256 ) view returns(uint256)
func (_Mintwatcher *MintwatcherCallerSession) BalanceOf(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Mintwatcher.Contract.BalanceOf(&_Mintwatcher.CallOpts, arg0, arg1)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] owners, uint256[] ids) view returns(uint256[] balances)
func (_Mintwatcher *MintwatcherCaller) BalanceOfBatch(opts *bind.CallOpts, owners []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Mintwatcher.contract.Call(opts, &out, "balanceOfBatch", owners, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] owners, uint256[] ids) view returns(uint256[] balances)
func (_Mintwatcher *MintwatcherSession) BalanceOfBatch(owners []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Mintwatcher.Contract.BalanceOfBatch(&_Mintwatcher.CallOpts, owners, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] owners, uint256[] ids) view returns(uint256[] balances)
func (_Mintwatcher *MintwatcherCallerSession) BalanceOfBatch(owners []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Mintwatcher.Contract.BalanceOfBatch(&_Mintwatcher.CallOpts, owners, ids)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_Mintwatcher *MintwatcherCaller) IsApprovedForAll(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Mintwatcher.contract.Call(opts, &out, "isApprovedForAll", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_Mintwatcher *MintwatcherSession) IsApprovedForAll(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Mintwatcher.Contract.IsApprovedForAll(&_Mintwatcher.CallOpts, arg0, arg1)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_Mintwatcher *MintwatcherCallerSession) IsApprovedForAll(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Mintwatcher.Contract.IsApprovedForAll(&_Mintwatcher.CallOpts, arg0, arg1)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mintwatcher *MintwatcherCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Mintwatcher.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mintwatcher *MintwatcherSession) Name() (string, error) {
	return _Mintwatcher.Contract.Name(&_Mintwatcher.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mintwatcher *MintwatcherCallerSession) Name() (string, error) {
	return _Mintwatcher.Contract.Name(&_Mintwatcher.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mintwatcher *MintwatcherCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mintwatcher.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mintwatcher *MintwatcherSession) Owner() (common.Address, error) {
	return _Mintwatcher.Contract.Owner(&_Mintwatcher.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mintwatcher *MintwatcherCallerSession) Owner() (common.Address, error) {
	return _Mintwatcher.Contract.Owner(&_Mintwatcher.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Mintwatcher *MintwatcherCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Mintwatcher.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Mintwatcher *MintwatcherSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Mintwatcher.Contract.SupportsInterface(&_Mintwatcher.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Mintwatcher *MintwatcherCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Mintwatcher.Contract.SupportsInterface(&_Mintwatcher.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mintwatcher *MintwatcherCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Mintwatcher.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mintwatcher *MintwatcherSession) Symbol() (string, error) {
	return _Mintwatcher.Contract.Symbol(&_Mintwatcher.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mintwatcher *MintwatcherCallerSession) Symbol() (string, error) {
	return _Mintwatcher.Contract.Symbol(&_Mintwatcher.CallOpts)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 id) view returns(string)
func (_Mintwatcher *MintwatcherCaller) Uri(opts *bind.CallOpts, id *big.Int) (string, error) {
	var out []interface{}
	err := _Mintwatcher.contract.Call(opts, &out, "uri", id)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 id) view returns(string)
func (_Mintwatcher *MintwatcherSession) Uri(id *big.Int) (string, error) {
	return _Mintwatcher.Contract.Uri(&_Mintwatcher.CallOpts, id)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 id) view returns(string)
func (_Mintwatcher *MintwatcherCallerSession) Uri(id *big.Int) (string, error) {
	return _Mintwatcher.Contract.Uri(&_Mintwatcher.CallOpts, id)
}

// Uris is a free data retrieval call binding the contract method 0x1253c546.
//
// Solidity: function uris(uint256 ) view returns(string)
func (_Mintwatcher *MintwatcherCaller) Uris(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Mintwatcher.contract.Call(opts, &out, "uris", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uris is a free data retrieval call binding the contract method 0x1253c546.
//
// Solidity: function uris(uint256 ) view returns(string)
func (_Mintwatcher *MintwatcherSession) Uris(arg0 *big.Int) (string, error) {
	return _Mintwatcher.Contract.Uris(&_Mintwatcher.CallOpts, arg0)
}

// Uris is a free data retrieval call binding the contract method 0x1253c546.
//
// Solidity: function uris(uint256 ) view returns(string)
func (_Mintwatcher *MintwatcherCallerSession) Uris(arg0 *big.Int) (string, error) {
	return _Mintwatcher.Contract.Uris(&_Mintwatcher.CallOpts, arg0)
}

// AddItems is a paid mutator transaction binding the contract method 0xffd95ba2.
//
// Solidity: function addItems(string[] _batchURIs) returns()
func (_Mintwatcher *MintwatcherTransactor) AddItems(opts *bind.TransactOpts, _batchURIs []string) (*types.Transaction, error) {
	return _Mintwatcher.contract.Transact(opts, "addItems", _batchURIs)
}

// AddItems is a paid mutator transaction binding the contract method 0xffd95ba2.
//
// Solidity: function addItems(string[] _batchURIs) returns()
func (_Mintwatcher *MintwatcherSession) AddItems(_batchURIs []string) (*types.Transaction, error) {
	return _Mintwatcher.Contract.AddItems(&_Mintwatcher.TransactOpts, _batchURIs)
}

// AddItems is a paid mutator transaction binding the contract method 0xffd95ba2.
//
// Solidity: function addItems(string[] _batchURIs) returns()
func (_Mintwatcher *MintwatcherTransactorSession) AddItems(_batchURIs []string) (*types.Transaction, error) {
	return _Mintwatcher.Contract.AddItems(&_Mintwatcher.TransactOpts, _batchURIs)
}

// BatchBurn is a paid mutator transaction binding the contract method 0xf6eb127a.
//
// Solidity: function batchBurn(address from, uint256[] ids, uint256[] amounts) returns()
func (_Mintwatcher *MintwatcherTransactor) BatchBurn(opts *bind.TransactOpts, from common.Address, ids []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _Mintwatcher.contract.Transact(opts, "batchBurn", from, ids, amounts)
}

// BatchBurn is a paid mutator transaction binding the contract method 0xf6eb127a.
//
// Solidity: function batchBurn(address from, uint256[] ids, uint256[] amounts) returns()
func (_Mintwatcher *MintwatcherSession) BatchBurn(from common.Address, ids []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _Mintwatcher.Contract.BatchBurn(&_Mintwatcher.TransactOpts, from, ids, amounts)
}

// BatchBurn is a paid mutator transaction binding the contract method 0xf6eb127a.
//
// Solidity: function batchBurn(address from, uint256[] ids, uint256[] amounts) returns()
func (_Mintwatcher *MintwatcherTransactorSession) BatchBurn(from common.Address, ids []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _Mintwatcher.Contract.BatchBurn(&_Mintwatcher.TransactOpts, from, ids, amounts)
}

// BatchMint is a paid mutator transaction binding the contract method 0xb48ab8b6.
//
// Solidity: function batchMint(address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Mintwatcher *MintwatcherTransactor) BatchMint(opts *bind.TransactOpts, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Mintwatcher.contract.Transact(opts, "batchMint", to, ids, amounts, data)
}

// BatchMint is a paid mutator transaction binding the contract method 0xb48ab8b6.
//
// Solidity: function batchMint(address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Mintwatcher *MintwatcherSession) BatchMint(to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Mintwatcher.Contract.BatchMint(&_Mintwatcher.TransactOpts, to, ids, amounts, data)
}

// BatchMint is a paid mutator transaction binding the contract method 0xb48ab8b6.
//
// Solidity: function batchMint(address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Mintwatcher *MintwatcherTransactorSession) BatchMint(to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Mintwatcher.Contract.BatchMint(&_Mintwatcher.TransactOpts, to, ids, amounts, data)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address from, uint256 id, uint256 amount) returns()
func (_Mintwatcher *MintwatcherTransactor) Burn(opts *bind.TransactOpts, from common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Mintwatcher.contract.Transact(opts, "burn", from, id, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address from, uint256 id, uint256 amount) returns()
func (_Mintwatcher *MintwatcherSession) Burn(from common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Mintwatcher.Contract.Burn(&_Mintwatcher.TransactOpts, from, id, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address from, uint256 id, uint256 amount) returns()
func (_Mintwatcher *MintwatcherTransactorSession) Burn(from common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Mintwatcher.Contract.Burn(&_Mintwatcher.TransactOpts, from, id, amount)
}

// CustomMint is a paid mutator transaction binding the contract method 0x90b3c27d.
//
// Solidity: function customMint(address to, uint256 id, uint256 amount, bytes data) returns()
func (_Mintwatcher *MintwatcherTransactor) CustomMint(opts *bind.TransactOpts, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Mintwatcher.contract.Transact(opts, "customMint", to, id, amount, data)
}

// CustomMint is a paid mutator transaction binding the contract method 0x90b3c27d.
//
// Solidity: function customMint(address to, uint256 id, uint256 amount, bytes data) returns()
func (_Mintwatcher *MintwatcherSession) CustomMint(to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Mintwatcher.Contract.CustomMint(&_Mintwatcher.TransactOpts, to, id, amount, data)
}

// CustomMint is a paid mutator transaction binding the contract method 0x90b3c27d.
//
// Solidity: function customMint(address to, uint256 id, uint256 amount, bytes data) returns()
func (_Mintwatcher *MintwatcherTransactorSession) CustomMint(to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Mintwatcher.Contract.CustomMint(&_Mintwatcher.TransactOpts, to, id, amount, data)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mintwatcher *MintwatcherTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mintwatcher.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mintwatcher *MintwatcherSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mintwatcher.Contract.RenounceOwnership(&_Mintwatcher.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mintwatcher *MintwatcherTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mintwatcher.Contract.RenounceOwnership(&_Mintwatcher.TransactOpts)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Mintwatcher *MintwatcherTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Mintwatcher.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Mintwatcher *MintwatcherSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Mintwatcher.Contract.SafeBatchTransferFrom(&_Mintwatcher.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Mintwatcher *MintwatcherTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Mintwatcher.Contract.SafeBatchTransferFrom(&_Mintwatcher.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Mintwatcher *MintwatcherTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Mintwatcher.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Mintwatcher *MintwatcherSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Mintwatcher.Contract.SafeTransferFrom(&_Mintwatcher.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Mintwatcher *MintwatcherTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Mintwatcher.Contract.SafeTransferFrom(&_Mintwatcher.TransactOpts, from, to, id, amount, data)
}

// SendNFTs is a paid mutator transaction binding the contract method 0xfc7f370b.
//
// Solidity: function sendNFTs(uint256[] itemIds, address[] receiptientAddresses) returns()
func (_Mintwatcher *MintwatcherTransactor) SendNFTs(opts *bind.TransactOpts, itemIds []*big.Int, receiptientAddresses []common.Address) (*types.Transaction, error) {
	return _Mintwatcher.contract.Transact(opts, "sendNFTs", itemIds, receiptientAddresses)
}

// SendNFTs is a paid mutator transaction binding the contract method 0xfc7f370b.
//
// Solidity: function sendNFTs(uint256[] itemIds, address[] receiptientAddresses) returns()
func (_Mintwatcher *MintwatcherSession) SendNFTs(itemIds []*big.Int, receiptientAddresses []common.Address) (*types.Transaction, error) {
	return _Mintwatcher.Contract.SendNFTs(&_Mintwatcher.TransactOpts, itemIds, receiptientAddresses)
}

// SendNFTs is a paid mutator transaction binding the contract method 0xfc7f370b.
//
// Solidity: function sendNFTs(uint256[] itemIds, address[] receiptientAddresses) returns()
func (_Mintwatcher *MintwatcherTransactorSession) SendNFTs(itemIds []*big.Int, receiptientAddresses []common.Address) (*types.Transaction, error) {
	return _Mintwatcher.Contract.SendNFTs(&_Mintwatcher.TransactOpts, itemIds, receiptientAddresses)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Mintwatcher *MintwatcherTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Mintwatcher.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Mintwatcher *MintwatcherSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Mintwatcher.Contract.SetApprovalForAll(&_Mintwatcher.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Mintwatcher *MintwatcherTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Mintwatcher.Contract.SetApprovalForAll(&_Mintwatcher.TransactOpts, operator, approved)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mintwatcher *MintwatcherTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Mintwatcher.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mintwatcher *MintwatcherSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mintwatcher.Contract.TransferOwnership(&_Mintwatcher.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mintwatcher *MintwatcherTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mintwatcher.Contract.TransferOwnership(&_Mintwatcher.TransactOpts, newOwner)
}

// MintwatcherApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Mintwatcher contract.
type MintwatcherApprovalForAllIterator struct {
	Event *MintwatcherApprovalForAll // Event containing the contract specifics and raw log

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
func (it *MintwatcherApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintwatcherApprovalForAll)
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
		it.Event = new(MintwatcherApprovalForAll)
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
func (it *MintwatcherApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintwatcherApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintwatcherApprovalForAll represents a ApprovalForAll event raised by the Mintwatcher contract.
type MintwatcherApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Mintwatcher *MintwatcherFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*MintwatcherApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Mintwatcher.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &MintwatcherApprovalForAllIterator{contract: _Mintwatcher.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Mintwatcher *MintwatcherFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *MintwatcherApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Mintwatcher.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintwatcherApprovalForAll)
				if err := _Mintwatcher.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Mintwatcher *MintwatcherFilterer) ParseApprovalForAll(log types.Log) (*MintwatcherApprovalForAll, error) {
	event := new(MintwatcherApprovalForAll)
	if err := _Mintwatcher.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintwatcherOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Mintwatcher contract.
type MintwatcherOwnershipTransferredIterator struct {
	Event *MintwatcherOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MintwatcherOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintwatcherOwnershipTransferred)
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
		it.Event = new(MintwatcherOwnershipTransferred)
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
func (it *MintwatcherOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintwatcherOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintwatcherOwnershipTransferred represents a OwnershipTransferred event raised by the Mintwatcher contract.
type MintwatcherOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mintwatcher *MintwatcherFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MintwatcherOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mintwatcher.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MintwatcherOwnershipTransferredIterator{contract: _Mintwatcher.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mintwatcher *MintwatcherFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MintwatcherOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mintwatcher.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintwatcherOwnershipTransferred)
				if err := _Mintwatcher.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mintwatcher *MintwatcherFilterer) ParseOwnershipTransferred(log types.Log) (*MintwatcherOwnershipTransferred, error) {
	event := new(MintwatcherOwnershipTransferred)
	if err := _Mintwatcher.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintwatcherTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the Mintwatcher contract.
type MintwatcherTransferBatchIterator struct {
	Event *MintwatcherTransferBatch // Event containing the contract specifics and raw log

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
func (it *MintwatcherTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintwatcherTransferBatch)
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
		it.Event = new(MintwatcherTransferBatch)
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
func (it *MintwatcherTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintwatcherTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintwatcherTransferBatch represents a TransferBatch event raised by the Mintwatcher contract.
type MintwatcherTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Amounts  []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] amounts)
func (_Mintwatcher *MintwatcherFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*MintwatcherTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Mintwatcher.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MintwatcherTransferBatchIterator{contract: _Mintwatcher.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] amounts)
func (_Mintwatcher *MintwatcherFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *MintwatcherTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Mintwatcher.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintwatcherTransferBatch)
				if err := _Mintwatcher.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] amounts)
func (_Mintwatcher *MintwatcherFilterer) ParseTransferBatch(log types.Log) (*MintwatcherTransferBatch, error) {
	event := new(MintwatcherTransferBatch)
	if err := _Mintwatcher.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintwatcherTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the Mintwatcher contract.
type MintwatcherTransferSingleIterator struct {
	Event *MintwatcherTransferSingle // Event containing the contract specifics and raw log

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
func (it *MintwatcherTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintwatcherTransferSingle)
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
		it.Event = new(MintwatcherTransferSingle)
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
func (it *MintwatcherTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintwatcherTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintwatcherTransferSingle represents a TransferSingle event raised by the Mintwatcher contract.
type MintwatcherTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 amount)
func (_Mintwatcher *MintwatcherFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*MintwatcherTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Mintwatcher.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MintwatcherTransferSingleIterator{contract: _Mintwatcher.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 amount)
func (_Mintwatcher *MintwatcherFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *MintwatcherTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Mintwatcher.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintwatcherTransferSingle)
				if err := _Mintwatcher.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 amount)
func (_Mintwatcher *MintwatcherFilterer) ParseTransferSingle(log types.Log) (*MintwatcherTransferSingle, error) {
	event := new(MintwatcherTransferSingle)
	if err := _Mintwatcher.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintwatcherURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the Mintwatcher contract.
type MintwatcherURIIterator struct {
	Event *MintwatcherURI // Event containing the contract specifics and raw log

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
func (it *MintwatcherURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintwatcherURI)
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
		it.Event = new(MintwatcherURI)
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
func (it *MintwatcherURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintwatcherURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintwatcherURI represents a URI event raised by the Mintwatcher contract.
type MintwatcherURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Mintwatcher *MintwatcherFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*MintwatcherURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Mintwatcher.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &MintwatcherURIIterator{contract: _Mintwatcher.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Mintwatcher *MintwatcherFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *MintwatcherURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Mintwatcher.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintwatcherURI)
				if err := _Mintwatcher.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Mintwatcher *MintwatcherFilterer) ParseURI(log types.Log) (*MintwatcherURI, error) {
	event := new(MintwatcherURI)
	if err := _Mintwatcher.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
