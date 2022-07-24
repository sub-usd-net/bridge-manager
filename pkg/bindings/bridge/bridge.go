// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridge

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
)

// BridgeDepositInfo is an auto generated low-level Go binding around an user-defined struct.
type BridgeDepositInfo struct {
	User   common.Address
	Amount *big.Int
}

// BridgeMetaData contains all meta data concerning the Bridge contract.
var BridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stableToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"avail\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"avail\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curr\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"MustBeSequential\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustNotBeZero\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CompleteTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"OwnerBorrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"OwnerReturn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"targetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"completeTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crossChainDepositId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"depositIdToBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"depositIdToDepositInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"depositId_\",\"type\":\"uint256\"}],\"name\":\"getDepositInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structBridge.DepositInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ownerBorrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ownerReturn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ownerWithdrawals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stableToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeMetaData.ABI instead.
var BridgeABI = BridgeMetaData.ABI

// Bridge is an auto generated Go binding around an Ethereum contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// CrossChainDepositId is a free data retrieval call binding the contract method 0xe02132ee.
//
// Solidity: function crossChainDepositId() view returns(uint256)
func (_Bridge *BridgeCaller) CrossChainDepositId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "crossChainDepositId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CrossChainDepositId is a free data retrieval call binding the contract method 0xe02132ee.
//
// Solidity: function crossChainDepositId() view returns(uint256)
func (_Bridge *BridgeSession) CrossChainDepositId() (*big.Int, error) {
	return _Bridge.Contract.CrossChainDepositId(&_Bridge.CallOpts)
}

// CrossChainDepositId is a free data retrieval call binding the contract method 0xe02132ee.
//
// Solidity: function crossChainDepositId() view returns(uint256)
func (_Bridge *BridgeCallerSession) CrossChainDepositId() (*big.Int, error) {
	return _Bridge.Contract.CrossChainDepositId(&_Bridge.CallOpts)
}

// CurrentIds is a free data retrieval call binding the contract method 0xe0a0f4fb.
//
// Solidity: function currentIds() view returns(uint256, uint256)
func (_Bridge *BridgeCaller) CurrentIds(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "currentIds")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// CurrentIds is a free data retrieval call binding the contract method 0xe0a0f4fb.
//
// Solidity: function currentIds() view returns(uint256, uint256)
func (_Bridge *BridgeSession) CurrentIds() (*big.Int, *big.Int, error) {
	return _Bridge.Contract.CurrentIds(&_Bridge.CallOpts)
}

// CurrentIds is a free data retrieval call binding the contract method 0xe0a0f4fb.
//
// Solidity: function currentIds() view returns(uint256, uint256)
func (_Bridge *BridgeCallerSession) CurrentIds() (*big.Int, *big.Int, error) {
	return _Bridge.Contract.CurrentIds(&_Bridge.CallOpts)
}

// DepositId is a free data retrieval call binding the contract method 0x9852099c.
//
// Solidity: function depositId() view returns(uint256)
func (_Bridge *BridgeCaller) DepositId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "depositId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositId is a free data retrieval call binding the contract method 0x9852099c.
//
// Solidity: function depositId() view returns(uint256)
func (_Bridge *BridgeSession) DepositId() (*big.Int, error) {
	return _Bridge.Contract.DepositId(&_Bridge.CallOpts)
}

// DepositId is a free data retrieval call binding the contract method 0x9852099c.
//
// Solidity: function depositId() view returns(uint256)
func (_Bridge *BridgeCallerSession) DepositId() (*big.Int, error) {
	return _Bridge.Contract.DepositId(&_Bridge.CallOpts)
}

// DepositIdToBlock is a free data retrieval call binding the contract method 0xb559a53a.
//
// Solidity: function depositIdToBlock(uint256 ) view returns(uint256)
func (_Bridge *BridgeCaller) DepositIdToBlock(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "depositIdToBlock", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositIdToBlock is a free data retrieval call binding the contract method 0xb559a53a.
//
// Solidity: function depositIdToBlock(uint256 ) view returns(uint256)
func (_Bridge *BridgeSession) DepositIdToBlock(arg0 *big.Int) (*big.Int, error) {
	return _Bridge.Contract.DepositIdToBlock(&_Bridge.CallOpts, arg0)
}

// DepositIdToBlock is a free data retrieval call binding the contract method 0xb559a53a.
//
// Solidity: function depositIdToBlock(uint256 ) view returns(uint256)
func (_Bridge *BridgeCallerSession) DepositIdToBlock(arg0 *big.Int) (*big.Int, error) {
	return _Bridge.Contract.DepositIdToBlock(&_Bridge.CallOpts, arg0)
}

// DepositIdToDepositInfo is a free data retrieval call binding the contract method 0x75f33e81.
//
// Solidity: function depositIdToDepositInfo(uint256 ) view returns(address user, uint256 amount)
func (_Bridge *BridgeCaller) DepositIdToDepositInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	User   common.Address
	Amount *big.Int
}, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "depositIdToDepositInfo", arg0)

	outstruct := new(struct {
		User   common.Address
		Amount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.User = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DepositIdToDepositInfo is a free data retrieval call binding the contract method 0x75f33e81.
//
// Solidity: function depositIdToDepositInfo(uint256 ) view returns(address user, uint256 amount)
func (_Bridge *BridgeSession) DepositIdToDepositInfo(arg0 *big.Int) (struct {
	User   common.Address
	Amount *big.Int
}, error) {
	return _Bridge.Contract.DepositIdToDepositInfo(&_Bridge.CallOpts, arg0)
}

// DepositIdToDepositInfo is a free data retrieval call binding the contract method 0x75f33e81.
//
// Solidity: function depositIdToDepositInfo(uint256 ) view returns(address user, uint256 amount)
func (_Bridge *BridgeCallerSession) DepositIdToDepositInfo(arg0 *big.Int) (struct {
	User   common.Address
	Amount *big.Int
}, error) {
	return _Bridge.Contract.DepositIdToDepositInfo(&_Bridge.CallOpts, arg0)
}

// GetDepositInfo is a free data retrieval call binding the contract method 0x5328c2bf.
//
// Solidity: function getDepositInfo(uint256 depositId_) view returns((address,uint256))
func (_Bridge *BridgeCaller) GetDepositInfo(opts *bind.CallOpts, depositId_ *big.Int) (BridgeDepositInfo, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getDepositInfo", depositId_)

	if err != nil {
		return *new(BridgeDepositInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(BridgeDepositInfo)).(*BridgeDepositInfo)

	return out0, err

}

// GetDepositInfo is a free data retrieval call binding the contract method 0x5328c2bf.
//
// Solidity: function getDepositInfo(uint256 depositId_) view returns((address,uint256))
func (_Bridge *BridgeSession) GetDepositInfo(depositId_ *big.Int) (BridgeDepositInfo, error) {
	return _Bridge.Contract.GetDepositInfo(&_Bridge.CallOpts, depositId_)
}

// GetDepositInfo is a free data retrieval call binding the contract method 0x5328c2bf.
//
// Solidity: function getDepositInfo(uint256 depositId_) view returns((address,uint256))
func (_Bridge *BridgeCallerSession) GetDepositInfo(depositId_ *big.Int) (BridgeDepositInfo, error) {
	return _Bridge.Contract.GetDepositInfo(&_Bridge.CallOpts, depositId_)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeSession) Owner() (common.Address, error) {
	return _Bridge.Contract.Owner(&_Bridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeCallerSession) Owner() (common.Address, error) {
	return _Bridge.Contract.Owner(&_Bridge.CallOpts)
}

// OwnerWithdrawals is a free data retrieval call binding the contract method 0x5c0b3d27.
//
// Solidity: function ownerWithdrawals(address ) view returns(uint256)
func (_Bridge *BridgeCaller) OwnerWithdrawals(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "ownerWithdrawals", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnerWithdrawals is a free data retrieval call binding the contract method 0x5c0b3d27.
//
// Solidity: function ownerWithdrawals(address ) view returns(uint256)
func (_Bridge *BridgeSession) OwnerWithdrawals(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.OwnerWithdrawals(&_Bridge.CallOpts, arg0)
}

// OwnerWithdrawals is a free data retrieval call binding the contract method 0x5c0b3d27.
//
// Solidity: function ownerWithdrawals(address ) view returns(uint256)
func (_Bridge *BridgeCallerSession) OwnerWithdrawals(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.OwnerWithdrawals(&_Bridge.CallOpts, arg0)
}

// StableToken is a free data retrieval call binding the contract method 0xa9d75b2b.
//
// Solidity: function stableToken() view returns(address)
func (_Bridge *BridgeCaller) StableToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "stableToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StableToken is a free data retrieval call binding the contract method 0xa9d75b2b.
//
// Solidity: function stableToken() view returns(address)
func (_Bridge *BridgeSession) StableToken() (common.Address, error) {
	return _Bridge.Contract.StableToken(&_Bridge.CallOpts)
}

// StableToken is a free data retrieval call binding the contract method 0xa9d75b2b.
//
// Solidity: function stableToken() view returns(address)
func (_Bridge *BridgeCallerSession) StableToken() (common.Address, error) {
	return _Bridge.Contract.StableToken(&_Bridge.CallOpts)
}

// CompleteTransfer is a paid mutator transaction binding the contract method 0x0901d6f7.
//
// Solidity: function completeTransfer(address beneficiary, uint256 targetId, uint256 amount) returns()
func (_Bridge *BridgeTransactor) CompleteTransfer(opts *bind.TransactOpts, beneficiary common.Address, targetId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "completeTransfer", beneficiary, targetId, amount)
}

// CompleteTransfer is a paid mutator transaction binding the contract method 0x0901d6f7.
//
// Solidity: function completeTransfer(address beneficiary, uint256 targetId, uint256 amount) returns()
func (_Bridge *BridgeSession) CompleteTransfer(beneficiary common.Address, targetId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.CompleteTransfer(&_Bridge.TransactOpts, beneficiary, targetId, amount)
}

// CompleteTransfer is a paid mutator transaction binding the contract method 0x0901d6f7.
//
// Solidity: function completeTransfer(address beneficiary, uint256 targetId, uint256 amount) returns()
func (_Bridge *BridgeTransactorSession) CompleteTransfer(beneficiary common.Address, targetId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.CompleteTransfer(&_Bridge.TransactOpts, beneficiary, targetId, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Bridge *BridgeTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Bridge *BridgeSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Deposit(&_Bridge.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Bridge *BridgeTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Deposit(&_Bridge.TransactOpts, amount)
}

// OwnerBorrow is a paid mutator transaction binding the contract method 0x08d9af12.
//
// Solidity: function ownerBorrow(uint256 amount) returns()
func (_Bridge *BridgeTransactor) OwnerBorrow(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "ownerBorrow", amount)
}

// OwnerBorrow is a paid mutator transaction binding the contract method 0x08d9af12.
//
// Solidity: function ownerBorrow(uint256 amount) returns()
func (_Bridge *BridgeSession) OwnerBorrow(amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.OwnerBorrow(&_Bridge.TransactOpts, amount)
}

// OwnerBorrow is a paid mutator transaction binding the contract method 0x08d9af12.
//
// Solidity: function ownerBorrow(uint256 amount) returns()
func (_Bridge *BridgeTransactorSession) OwnerBorrow(amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.OwnerBorrow(&_Bridge.TransactOpts, amount)
}

// OwnerReturn is a paid mutator transaction binding the contract method 0xba6059d9.
//
// Solidity: function ownerReturn(uint256 amount) returns()
func (_Bridge *BridgeTransactor) OwnerReturn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "ownerReturn", amount)
}

// OwnerReturn is a paid mutator transaction binding the contract method 0xba6059d9.
//
// Solidity: function ownerReturn(uint256 amount) returns()
func (_Bridge *BridgeSession) OwnerReturn(amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.OwnerReturn(&_Bridge.TransactOpts, amount)
}

// OwnerReturn is a paid mutator transaction binding the contract method 0xba6059d9.
//
// Solidity: function ownerReturn(uint256 amount) returns()
func (_Bridge *BridgeTransactorSession) OwnerReturn(amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.OwnerReturn(&_Bridge.TransactOpts, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bridge.Contract.RenounceOwnership(&_Bridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bridge.Contract.RenounceOwnership(&_Bridge.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.TransferOwnership(&_Bridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.TransferOwnership(&_Bridge.TransactOpts, newOwner)
}

// BridgeCompleteTransferIterator is returned from FilterCompleteTransfer and is used to iterate over the raw logs and unpacked data for CompleteTransfer events raised by the Bridge contract.
type BridgeCompleteTransferIterator struct {
	Event *BridgeCompleteTransfer // Event containing the contract specifics and raw log

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
func (it *BridgeCompleteTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeCompleteTransfer)
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
		it.Event = new(BridgeCompleteTransfer)
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
func (it *BridgeCompleteTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeCompleteTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeCompleteTransfer represents a CompleteTransfer event raised by the Bridge contract.
type BridgeCompleteTransfer struct {
	Beneficiary common.Address
	Id          *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCompleteTransfer is a free log retrieval operation binding the contract event 0x1a3467986fee07094e44ed6fbe328521301a92e34d0552fb39067655096163f3.
//
// Solidity: event CompleteTransfer(address indexed beneficiary, uint256 indexed id, uint256 amount)
func (_Bridge *BridgeFilterer) FilterCompleteTransfer(opts *bind.FilterOpts, beneficiary []common.Address, id []*big.Int) (*BridgeCompleteTransferIterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "CompleteTransfer", beneficiaryRule, idRule)
	if err != nil {
		return nil, err
	}
	return &BridgeCompleteTransferIterator{contract: _Bridge.contract, event: "CompleteTransfer", logs: logs, sub: sub}, nil
}

// WatchCompleteTransfer is a free log subscription operation binding the contract event 0x1a3467986fee07094e44ed6fbe328521301a92e34d0552fb39067655096163f3.
//
// Solidity: event CompleteTransfer(address indexed beneficiary, uint256 indexed id, uint256 amount)
func (_Bridge *BridgeFilterer) WatchCompleteTransfer(opts *bind.WatchOpts, sink chan<- *BridgeCompleteTransfer, beneficiary []common.Address, id []*big.Int) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "CompleteTransfer", beneficiaryRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeCompleteTransfer)
				if err := _Bridge.contract.UnpackLog(event, "CompleteTransfer", log); err != nil {
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

// ParseCompleteTransfer is a log parse operation binding the contract event 0x1a3467986fee07094e44ed6fbe328521301a92e34d0552fb39067655096163f3.
//
// Solidity: event CompleteTransfer(address indexed beneficiary, uint256 indexed id, uint256 amount)
func (_Bridge *BridgeFilterer) ParseCompleteTransfer(log types.Log) (*BridgeCompleteTransfer, error) {
	event := new(BridgeCompleteTransfer)
	if err := _Bridge.contract.UnpackLog(event, "CompleteTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Bridge contract.
type BridgeDepositIterator struct {
	Event *BridgeDeposit // Event containing the contract specifics and raw log

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
func (it *BridgeDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeDeposit)
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
		it.Event = new(BridgeDeposit)
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
func (it *BridgeDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeDeposit represents a Deposit event raised by the Bridge contract.
type BridgeDeposit struct {
	Depositor common.Address
	Id        *big.Int
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed depositor, uint256 indexed id, uint256 amount)
func (_Bridge *BridgeFilterer) FilterDeposit(opts *bind.FilterOpts, depositor []common.Address, id []*big.Int) (*BridgeDepositIterator, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Deposit", depositorRule, idRule)
	if err != nil {
		return nil, err
	}
	return &BridgeDepositIterator{contract: _Bridge.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed depositor, uint256 indexed id, uint256 amount)
func (_Bridge *BridgeFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *BridgeDeposit, depositor []common.Address, id []*big.Int) (event.Subscription, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Deposit", depositorRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeDeposit)
				if err := _Bridge.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed depositor, uint256 indexed id, uint256 amount)
func (_Bridge *BridgeFilterer) ParseDeposit(log types.Log) (*BridgeDeposit, error) {
	event := new(BridgeDeposit)
	if err := _Bridge.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeOwnerBorrowIterator is returned from FilterOwnerBorrow and is used to iterate over the raw logs and unpacked data for OwnerBorrow events raised by the Bridge contract.
type BridgeOwnerBorrowIterator struct {
	Event *BridgeOwnerBorrow // Event containing the contract specifics and raw log

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
func (it *BridgeOwnerBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeOwnerBorrow)
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
		it.Event = new(BridgeOwnerBorrow)
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
func (it *BridgeOwnerBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeOwnerBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeOwnerBorrow represents a OwnerBorrow event raised by the Bridge contract.
type BridgeOwnerBorrow struct {
	Owner  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOwnerBorrow is a free log retrieval operation binding the contract event 0x4fbfd3bb0fb7fbe460e3a494f1c890f838ee0f930ab79b6bf21f07b88372df0a.
//
// Solidity: event OwnerBorrow(address indexed owner, uint256 amount)
func (_Bridge *BridgeFilterer) FilterOwnerBorrow(opts *bind.FilterOpts, owner []common.Address) (*BridgeOwnerBorrowIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "OwnerBorrow", ownerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeOwnerBorrowIterator{contract: _Bridge.contract, event: "OwnerBorrow", logs: logs, sub: sub}, nil
}

// WatchOwnerBorrow is a free log subscription operation binding the contract event 0x4fbfd3bb0fb7fbe460e3a494f1c890f838ee0f930ab79b6bf21f07b88372df0a.
//
// Solidity: event OwnerBorrow(address indexed owner, uint256 amount)
func (_Bridge *BridgeFilterer) WatchOwnerBorrow(opts *bind.WatchOpts, sink chan<- *BridgeOwnerBorrow, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "OwnerBorrow", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeOwnerBorrow)
				if err := _Bridge.contract.UnpackLog(event, "OwnerBorrow", log); err != nil {
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

// ParseOwnerBorrow is a log parse operation binding the contract event 0x4fbfd3bb0fb7fbe460e3a494f1c890f838ee0f930ab79b6bf21f07b88372df0a.
//
// Solidity: event OwnerBorrow(address indexed owner, uint256 amount)
func (_Bridge *BridgeFilterer) ParseOwnerBorrow(log types.Log) (*BridgeOwnerBorrow, error) {
	event := new(BridgeOwnerBorrow)
	if err := _Bridge.contract.UnpackLog(event, "OwnerBorrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeOwnerReturnIterator is returned from FilterOwnerReturn and is used to iterate over the raw logs and unpacked data for OwnerReturn events raised by the Bridge contract.
type BridgeOwnerReturnIterator struct {
	Event *BridgeOwnerReturn // Event containing the contract specifics and raw log

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
func (it *BridgeOwnerReturnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeOwnerReturn)
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
		it.Event = new(BridgeOwnerReturn)
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
func (it *BridgeOwnerReturnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeOwnerReturnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeOwnerReturn represents a OwnerReturn event raised by the Bridge contract.
type BridgeOwnerReturn struct {
	Owner  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOwnerReturn is a free log retrieval operation binding the contract event 0x607fae14c29b49308dc94bcb474142bc68ea380c8790f229472c1f2a089647c7.
//
// Solidity: event OwnerReturn(address indexed owner, uint256 amount)
func (_Bridge *BridgeFilterer) FilterOwnerReturn(opts *bind.FilterOpts, owner []common.Address) (*BridgeOwnerReturnIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "OwnerReturn", ownerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeOwnerReturnIterator{contract: _Bridge.contract, event: "OwnerReturn", logs: logs, sub: sub}, nil
}

// WatchOwnerReturn is a free log subscription operation binding the contract event 0x607fae14c29b49308dc94bcb474142bc68ea380c8790f229472c1f2a089647c7.
//
// Solidity: event OwnerReturn(address indexed owner, uint256 amount)
func (_Bridge *BridgeFilterer) WatchOwnerReturn(opts *bind.WatchOpts, sink chan<- *BridgeOwnerReturn, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "OwnerReturn", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeOwnerReturn)
				if err := _Bridge.contract.UnpackLog(event, "OwnerReturn", log); err != nil {
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

// ParseOwnerReturn is a log parse operation binding the contract event 0x607fae14c29b49308dc94bcb474142bc68ea380c8790f229472c1f2a089647c7.
//
// Solidity: event OwnerReturn(address indexed owner, uint256 amount)
func (_Bridge *BridgeFilterer) ParseOwnerReturn(log types.Log) (*BridgeOwnerReturn, error) {
	event := new(BridgeOwnerReturn)
	if err := _Bridge.contract.UnpackLog(event, "OwnerReturn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bridge contract.
type BridgeOwnershipTransferredIterator struct {
	Event *BridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeOwnershipTransferred)
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
		it.Event = new(BridgeOwnershipTransferred)
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
func (it *BridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeOwnershipTransferred represents a OwnershipTransferred event raised by the Bridge contract.
type BridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridge *BridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeOwnershipTransferredIterator{contract: _Bridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridge *BridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeOwnershipTransferred)
				if err := _Bridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseOwnershipTransferred(log types.Log) (*BridgeOwnershipTransferred, error) {
	event := new(BridgeOwnershipTransferred)
	if err := _Bridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
