// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package directbtcminter

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

// DirectBTCMinterEvent is an auto generated low-level Go binding around an user-defined struct.
type DirectBTCMinterEvent struct {
	Recipient common.Address
	Amount    *big.Int
	State     uint8
}

// DirectbtcminterMetaData contains all meta data concerning the Directbtcminter contract.
var DirectbtcminterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Accepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allow\",\"type\":\"bool\"}],\"name\":\"RecipientChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Rejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"APPROVER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L1_MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_reqHash\",\"type\":\"bytes32\"}],\"name\":\"approveEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"directBTC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"eventIndexes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_defaultAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_directBTC\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_uniBTC\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextPendingEvent\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"enumDirectBTCMinter.EventState\",\"name\":\"state\",\"type\":\"uint8\"}],\"internalType\":\"structDirectBTCMinter.Event\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"processIdx\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_txHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"receiveEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"receivedEvents\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"enumDirectBTCMinter.EventState\",\"name\":\"state\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"recipients\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_reqHash\",\"type\":\"bytes32\"}],\"name\":\"rejectEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allow\",\"type\":\"bool\"}],\"name\":\"setRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniBTC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DirectbtcminterABI is the input ABI used to generate the binding from.
// Deprecated: Use DirectbtcminterMetaData.ABI instead.
var DirectbtcminterABI = DirectbtcminterMetaData.ABI

// Directbtcminter is an auto generated Go binding around an Ethereum contract.
type Directbtcminter struct {
	DirectbtcminterCaller     // Read-only binding to the contract
	DirectbtcminterTransactor // Write-only binding to the contract
	DirectbtcminterFilterer   // Log filterer for contract events
}

// DirectbtcminterCaller is an auto generated read-only Go binding around an Ethereum contract.
type DirectbtcminterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DirectbtcminterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DirectbtcminterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DirectbtcminterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DirectbtcminterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DirectbtcminterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DirectbtcminterSession struct {
	Contract     *Directbtcminter  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DirectbtcminterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DirectbtcminterCallerSession struct {
	Contract *DirectbtcminterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// DirectbtcminterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DirectbtcminterTransactorSession struct {
	Contract     *DirectbtcminterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// DirectbtcminterRaw is an auto generated low-level Go binding around an Ethereum contract.
type DirectbtcminterRaw struct {
	Contract *Directbtcminter // Generic contract binding to access the raw methods on
}

// DirectbtcminterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DirectbtcminterCallerRaw struct {
	Contract *DirectbtcminterCaller // Generic read-only contract binding to access the raw methods on
}

// DirectbtcminterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DirectbtcminterTransactorRaw struct {
	Contract *DirectbtcminterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDirectbtcminter creates a new instance of Directbtcminter, bound to a specific deployed contract.
func NewDirectbtcminter(address common.Address, backend bind.ContractBackend) (*Directbtcminter, error) {
	contract, err := bindDirectbtcminter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Directbtcminter{DirectbtcminterCaller: DirectbtcminterCaller{contract: contract}, DirectbtcminterTransactor: DirectbtcminterTransactor{contract: contract}, DirectbtcminterFilterer: DirectbtcminterFilterer{contract: contract}}, nil
}

// NewDirectbtcminterCaller creates a new read-only instance of Directbtcminter, bound to a specific deployed contract.
func NewDirectbtcminterCaller(address common.Address, caller bind.ContractCaller) (*DirectbtcminterCaller, error) {
	contract, err := bindDirectbtcminter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DirectbtcminterCaller{contract: contract}, nil
}

// NewDirectbtcminterTransactor creates a new write-only instance of Directbtcminter, bound to a specific deployed contract.
func NewDirectbtcminterTransactor(address common.Address, transactor bind.ContractTransactor) (*DirectbtcminterTransactor, error) {
	contract, err := bindDirectbtcminter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DirectbtcminterTransactor{contract: contract}, nil
}

// NewDirectbtcminterFilterer creates a new log filterer instance of Directbtcminter, bound to a specific deployed contract.
func NewDirectbtcminterFilterer(address common.Address, filterer bind.ContractFilterer) (*DirectbtcminterFilterer, error) {
	contract, err := bindDirectbtcminter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DirectbtcminterFilterer{contract: contract}, nil
}

// bindDirectbtcminter binds a generic wrapper to an already deployed contract.
func bindDirectbtcminter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DirectbtcminterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Directbtcminter *DirectbtcminterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Directbtcminter.Contract.DirectbtcminterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Directbtcminter *DirectbtcminterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Directbtcminter.Contract.DirectbtcminterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Directbtcminter *DirectbtcminterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Directbtcminter.Contract.DirectbtcminterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Directbtcminter *DirectbtcminterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Directbtcminter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Directbtcminter *DirectbtcminterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Directbtcminter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Directbtcminter *DirectbtcminterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Directbtcminter.Contract.contract.Transact(opts, method, params...)
}

// APPROVERROLE is a free data retrieval call binding the contract method 0x4245962b.
//
// Solidity: function APPROVER_ROLE() view returns(bytes32)
func (_Directbtcminter *DirectbtcminterCaller) APPROVERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Directbtcminter.contract.Call(opts, &out, "APPROVER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// APPROVERROLE is a free data retrieval call binding the contract method 0x4245962b.
//
// Solidity: function APPROVER_ROLE() view returns(bytes32)
func (_Directbtcminter *DirectbtcminterSession) APPROVERROLE() ([32]byte, error) {
	return _Directbtcminter.Contract.APPROVERROLE(&_Directbtcminter.CallOpts)
}

// APPROVERROLE is a free data retrieval call binding the contract method 0x4245962b.
//
// Solidity: function APPROVER_ROLE() view returns(bytes32)
func (_Directbtcminter *DirectbtcminterCallerSession) APPROVERROLE() ([32]byte, error) {
	return _Directbtcminter.Contract.APPROVERROLE(&_Directbtcminter.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Directbtcminter *DirectbtcminterCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Directbtcminter.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Directbtcminter *DirectbtcminterSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Directbtcminter.Contract.DEFAULTADMINROLE(&_Directbtcminter.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Directbtcminter *DirectbtcminterCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Directbtcminter.Contract.DEFAULTADMINROLE(&_Directbtcminter.CallOpts)
}

// L1MINTERROLE is a free data retrieval call binding the contract method 0x6aaf5c24.
//
// Solidity: function L1_MINTER_ROLE() view returns(bytes32)
func (_Directbtcminter *DirectbtcminterCaller) L1MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Directbtcminter.contract.Call(opts, &out, "L1_MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L1MINTERROLE is a free data retrieval call binding the contract method 0x6aaf5c24.
//
// Solidity: function L1_MINTER_ROLE() view returns(bytes32)
func (_Directbtcminter *DirectbtcminterSession) L1MINTERROLE() ([32]byte, error) {
	return _Directbtcminter.Contract.L1MINTERROLE(&_Directbtcminter.CallOpts)
}

// L1MINTERROLE is a free data retrieval call binding the contract method 0x6aaf5c24.
//
// Solidity: function L1_MINTER_ROLE() view returns(bytes32)
func (_Directbtcminter *DirectbtcminterCallerSession) L1MINTERROLE() ([32]byte, error) {
	return _Directbtcminter.Contract.L1MINTERROLE(&_Directbtcminter.CallOpts)
}

// DirectBTC is a free data retrieval call binding the contract method 0x38b1f052.
//
// Solidity: function directBTC() view returns(address)
func (_Directbtcminter *DirectbtcminterCaller) DirectBTC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Directbtcminter.contract.Call(opts, &out, "directBTC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DirectBTC is a free data retrieval call binding the contract method 0x38b1f052.
//
// Solidity: function directBTC() view returns(address)
func (_Directbtcminter *DirectbtcminterSession) DirectBTC() (common.Address, error) {
	return _Directbtcminter.Contract.DirectBTC(&_Directbtcminter.CallOpts)
}

// DirectBTC is a free data retrieval call binding the contract method 0x38b1f052.
//
// Solidity: function directBTC() view returns(address)
func (_Directbtcminter *DirectbtcminterCallerSession) DirectBTC() (common.Address, error) {
	return _Directbtcminter.Contract.DirectBTC(&_Directbtcminter.CallOpts)
}

// EventIndexes is a free data retrieval call binding the contract method 0xd01c1730.
//
// Solidity: function eventIndexes(uint256 ) view returns(bytes32)
func (_Directbtcminter *DirectbtcminterCaller) EventIndexes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Directbtcminter.contract.Call(opts, &out, "eventIndexes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EventIndexes is a free data retrieval call binding the contract method 0xd01c1730.
//
// Solidity: function eventIndexes(uint256 ) view returns(bytes32)
func (_Directbtcminter *DirectbtcminterSession) EventIndexes(arg0 *big.Int) ([32]byte, error) {
	return _Directbtcminter.Contract.EventIndexes(&_Directbtcminter.CallOpts, arg0)
}

// EventIndexes is a free data retrieval call binding the contract method 0xd01c1730.
//
// Solidity: function eventIndexes(uint256 ) view returns(bytes32)
func (_Directbtcminter *DirectbtcminterCallerSession) EventIndexes(arg0 *big.Int) ([32]byte, error) {
	return _Directbtcminter.Contract.EventIndexes(&_Directbtcminter.CallOpts, arg0)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Directbtcminter *DirectbtcminterCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Directbtcminter.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Directbtcminter *DirectbtcminterSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Directbtcminter.Contract.GetRoleAdmin(&_Directbtcminter.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Directbtcminter *DirectbtcminterCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Directbtcminter.Contract.GetRoleAdmin(&_Directbtcminter.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Directbtcminter *DirectbtcminterCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Directbtcminter.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Directbtcminter *DirectbtcminterSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Directbtcminter.Contract.HasRole(&_Directbtcminter.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Directbtcminter *DirectbtcminterCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Directbtcminter.Contract.HasRole(&_Directbtcminter.CallOpts, role, account)
}

// NextPendingEvent is a free data retrieval call binding the contract method 0x0a7fce82.
//
// Solidity: function nextPendingEvent() view returns(bytes32, (address,uint256,uint8))
func (_Directbtcminter *DirectbtcminterCaller) NextPendingEvent(opts *bind.CallOpts) ([32]byte, DirectBTCMinterEvent, error) {
	var out []interface{}
	err := _Directbtcminter.contract.Call(opts, &out, "nextPendingEvent")

	if err != nil {
		return *new([32]byte), *new(DirectBTCMinterEvent), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(DirectBTCMinterEvent)).(*DirectBTCMinterEvent)

	return out0, out1, err

}

// NextPendingEvent is a free data retrieval call binding the contract method 0x0a7fce82.
//
// Solidity: function nextPendingEvent() view returns(bytes32, (address,uint256,uint8))
func (_Directbtcminter *DirectbtcminterSession) NextPendingEvent() ([32]byte, DirectBTCMinterEvent, error) {
	return _Directbtcminter.Contract.NextPendingEvent(&_Directbtcminter.CallOpts)
}

// NextPendingEvent is a free data retrieval call binding the contract method 0x0a7fce82.
//
// Solidity: function nextPendingEvent() view returns(bytes32, (address,uint256,uint8))
func (_Directbtcminter *DirectbtcminterCallerSession) NextPendingEvent() ([32]byte, DirectBTCMinterEvent, error) {
	return _Directbtcminter.Contract.NextPendingEvent(&_Directbtcminter.CallOpts)
}

// ProcessIdx is a free data retrieval call binding the contract method 0xe459efe1.
//
// Solidity: function processIdx() view returns(uint64)
func (_Directbtcminter *DirectbtcminterCaller) ProcessIdx(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Directbtcminter.contract.Call(opts, &out, "processIdx")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ProcessIdx is a free data retrieval call binding the contract method 0xe459efe1.
//
// Solidity: function processIdx() view returns(uint64)
func (_Directbtcminter *DirectbtcminterSession) ProcessIdx() (uint64, error) {
	return _Directbtcminter.Contract.ProcessIdx(&_Directbtcminter.CallOpts)
}

// ProcessIdx is a free data retrieval call binding the contract method 0xe459efe1.
//
// Solidity: function processIdx() view returns(uint64)
func (_Directbtcminter *DirectbtcminterCallerSession) ProcessIdx() (uint64, error) {
	return _Directbtcminter.Contract.ProcessIdx(&_Directbtcminter.CallOpts)
}

// ReceivedEvents is a free data retrieval call binding the contract method 0x6333578b.
//
// Solidity: function receivedEvents(bytes32 ) view returns(address recipient, uint256 amount, uint8 state)
func (_Directbtcminter *DirectbtcminterCaller) ReceivedEvents(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Recipient common.Address
	Amount    *big.Int
	State     uint8
}, error) {
	var out []interface{}
	err := _Directbtcminter.contract.Call(opts, &out, "receivedEvents", arg0)

	outstruct := new(struct {
		Recipient common.Address
		Amount    *big.Int
		State     uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Recipient = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.State = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// ReceivedEvents is a free data retrieval call binding the contract method 0x6333578b.
//
// Solidity: function receivedEvents(bytes32 ) view returns(address recipient, uint256 amount, uint8 state)
func (_Directbtcminter *DirectbtcminterSession) ReceivedEvents(arg0 [32]byte) (struct {
	Recipient common.Address
	Amount    *big.Int
	State     uint8
}, error) {
	return _Directbtcminter.Contract.ReceivedEvents(&_Directbtcminter.CallOpts, arg0)
}

// ReceivedEvents is a free data retrieval call binding the contract method 0x6333578b.
//
// Solidity: function receivedEvents(bytes32 ) view returns(address recipient, uint256 amount, uint8 state)
func (_Directbtcminter *DirectbtcminterCallerSession) ReceivedEvents(arg0 [32]byte) (struct {
	Recipient common.Address
	Amount    *big.Int
	State     uint8
}, error) {
	return _Directbtcminter.Contract.ReceivedEvents(&_Directbtcminter.CallOpts, arg0)
}

// Recipients is a free data retrieval call binding the contract method 0xeb820312.
//
// Solidity: function recipients(address ) view returns(bool)
func (_Directbtcminter *DirectbtcminterCaller) Recipients(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Directbtcminter.contract.Call(opts, &out, "recipients", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Recipients is a free data retrieval call binding the contract method 0xeb820312.
//
// Solidity: function recipients(address ) view returns(bool)
func (_Directbtcminter *DirectbtcminterSession) Recipients(arg0 common.Address) (bool, error) {
	return _Directbtcminter.Contract.Recipients(&_Directbtcminter.CallOpts, arg0)
}

// Recipients is a free data retrieval call binding the contract method 0xeb820312.
//
// Solidity: function recipients(address ) view returns(bool)
func (_Directbtcminter *DirectbtcminterCallerSession) Recipients(arg0 common.Address) (bool, error) {
	return _Directbtcminter.Contract.Recipients(&_Directbtcminter.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Directbtcminter *DirectbtcminterCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Directbtcminter.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Directbtcminter *DirectbtcminterSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Directbtcminter.Contract.SupportsInterface(&_Directbtcminter.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Directbtcminter *DirectbtcminterCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Directbtcminter.Contract.SupportsInterface(&_Directbtcminter.CallOpts, interfaceId)
}

// UniBTC is a free data retrieval call binding the contract method 0x59f3d39b.
//
// Solidity: function uniBTC() view returns(address)
func (_Directbtcminter *DirectbtcminterCaller) UniBTC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Directbtcminter.contract.Call(opts, &out, "uniBTC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniBTC is a free data retrieval call binding the contract method 0x59f3d39b.
//
// Solidity: function uniBTC() view returns(address)
func (_Directbtcminter *DirectbtcminterSession) UniBTC() (common.Address, error) {
	return _Directbtcminter.Contract.UniBTC(&_Directbtcminter.CallOpts)
}

// UniBTC is a free data retrieval call binding the contract method 0x59f3d39b.
//
// Solidity: function uniBTC() view returns(address)
func (_Directbtcminter *DirectbtcminterCallerSession) UniBTC() (common.Address, error) {
	return _Directbtcminter.Contract.UniBTC(&_Directbtcminter.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Directbtcminter *DirectbtcminterCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Directbtcminter.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Directbtcminter *DirectbtcminterSession) Vault() (common.Address, error) {
	return _Directbtcminter.Contract.Vault(&_Directbtcminter.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Directbtcminter *DirectbtcminterCallerSession) Vault() (common.Address, error) {
	return _Directbtcminter.Contract.Vault(&_Directbtcminter.CallOpts)
}

// ApproveEvent is a paid mutator transaction binding the contract method 0x65a46399.
//
// Solidity: function approveEvent(bytes32 _reqHash) returns()
func (_Directbtcminter *DirectbtcminterTransactor) ApproveEvent(opts *bind.TransactOpts, _reqHash [32]byte) (*types.Transaction, error) {
	return _Directbtcminter.contract.Transact(opts, "approveEvent", _reqHash)
}

// ApproveEvent is a paid mutator transaction binding the contract method 0x65a46399.
//
// Solidity: function approveEvent(bytes32 _reqHash) returns()
func (_Directbtcminter *DirectbtcminterSession) ApproveEvent(_reqHash [32]byte) (*types.Transaction, error) {
	return _Directbtcminter.Contract.ApproveEvent(&_Directbtcminter.TransactOpts, _reqHash)
}

// ApproveEvent is a paid mutator transaction binding the contract method 0x65a46399.
//
// Solidity: function approveEvent(bytes32 _reqHash) returns()
func (_Directbtcminter *DirectbtcminterTransactorSession) ApproveEvent(_reqHash [32]byte) (*types.Transaction, error) {
	return _Directbtcminter.Contract.ApproveEvent(&_Directbtcminter.TransactOpts, _reqHash)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Directbtcminter *DirectbtcminterTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Directbtcminter.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Directbtcminter *DirectbtcminterSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Directbtcminter.Contract.GrantRole(&_Directbtcminter.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Directbtcminter *DirectbtcminterTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Directbtcminter.Contract.GrantRole(&_Directbtcminter.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _defaultAdmin, address _directBTC, address _vault, address _uniBTC) returns()
func (_Directbtcminter *DirectbtcminterTransactor) Initialize(opts *bind.TransactOpts, _defaultAdmin common.Address, _directBTC common.Address, _vault common.Address, _uniBTC common.Address) (*types.Transaction, error) {
	return _Directbtcminter.contract.Transact(opts, "initialize", _defaultAdmin, _directBTC, _vault, _uniBTC)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _defaultAdmin, address _directBTC, address _vault, address _uniBTC) returns()
func (_Directbtcminter *DirectbtcminterSession) Initialize(_defaultAdmin common.Address, _directBTC common.Address, _vault common.Address, _uniBTC common.Address) (*types.Transaction, error) {
	return _Directbtcminter.Contract.Initialize(&_Directbtcminter.TransactOpts, _defaultAdmin, _directBTC, _vault, _uniBTC)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _defaultAdmin, address _directBTC, address _vault, address _uniBTC) returns()
func (_Directbtcminter *DirectbtcminterTransactorSession) Initialize(_defaultAdmin common.Address, _directBTC common.Address, _vault common.Address, _uniBTC common.Address) (*types.Transaction, error) {
	return _Directbtcminter.Contract.Initialize(&_Directbtcminter.TransactOpts, _defaultAdmin, _directBTC, _vault, _uniBTC)
}

// ReceiveEvent is a paid mutator transaction binding the contract method 0x4b52f341.
//
// Solidity: function receiveEvent(address _recipient, bytes32 _txHash, uint256 _amount) returns()
func (_Directbtcminter *DirectbtcminterTransactor) ReceiveEvent(opts *bind.TransactOpts, _recipient common.Address, _txHash [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _Directbtcminter.contract.Transact(opts, "receiveEvent", _recipient, _txHash, _amount)
}

// ReceiveEvent is a paid mutator transaction binding the contract method 0x4b52f341.
//
// Solidity: function receiveEvent(address _recipient, bytes32 _txHash, uint256 _amount) returns()
func (_Directbtcminter *DirectbtcminterSession) ReceiveEvent(_recipient common.Address, _txHash [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _Directbtcminter.Contract.ReceiveEvent(&_Directbtcminter.TransactOpts, _recipient, _txHash, _amount)
}

// ReceiveEvent is a paid mutator transaction binding the contract method 0x4b52f341.
//
// Solidity: function receiveEvent(address _recipient, bytes32 _txHash, uint256 _amount) returns()
func (_Directbtcminter *DirectbtcminterTransactorSession) ReceiveEvent(_recipient common.Address, _txHash [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _Directbtcminter.Contract.ReceiveEvent(&_Directbtcminter.TransactOpts, _recipient, _txHash, _amount)
}

// RejectEvent is a paid mutator transaction binding the contract method 0xd80dcda4.
//
// Solidity: function rejectEvent(bytes32 _reqHash) returns()
func (_Directbtcminter *DirectbtcminterTransactor) RejectEvent(opts *bind.TransactOpts, _reqHash [32]byte) (*types.Transaction, error) {
	return _Directbtcminter.contract.Transact(opts, "rejectEvent", _reqHash)
}

// RejectEvent is a paid mutator transaction binding the contract method 0xd80dcda4.
//
// Solidity: function rejectEvent(bytes32 _reqHash) returns()
func (_Directbtcminter *DirectbtcminterSession) RejectEvent(_reqHash [32]byte) (*types.Transaction, error) {
	return _Directbtcminter.Contract.RejectEvent(&_Directbtcminter.TransactOpts, _reqHash)
}

// RejectEvent is a paid mutator transaction binding the contract method 0xd80dcda4.
//
// Solidity: function rejectEvent(bytes32 _reqHash) returns()
func (_Directbtcminter *DirectbtcminterTransactorSession) RejectEvent(_reqHash [32]byte) (*types.Transaction, error) {
	return _Directbtcminter.Contract.RejectEvent(&_Directbtcminter.TransactOpts, _reqHash)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Directbtcminter *DirectbtcminterTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Directbtcminter.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Directbtcminter *DirectbtcminterSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Directbtcminter.Contract.RenounceRole(&_Directbtcminter.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Directbtcminter *DirectbtcminterTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Directbtcminter.Contract.RenounceRole(&_Directbtcminter.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Directbtcminter *DirectbtcminterTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Directbtcminter.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Directbtcminter *DirectbtcminterSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Directbtcminter.Contract.RevokeRole(&_Directbtcminter.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Directbtcminter *DirectbtcminterTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Directbtcminter.Contract.RevokeRole(&_Directbtcminter.TransactOpts, role, account)
}

// SetRecipient is a paid mutator transaction binding the contract method 0xb33480ca.
//
// Solidity: function setRecipient(address _addr, bool allow) returns()
func (_Directbtcminter *DirectbtcminterTransactor) SetRecipient(opts *bind.TransactOpts, _addr common.Address, allow bool) (*types.Transaction, error) {
	return _Directbtcminter.contract.Transact(opts, "setRecipient", _addr, allow)
}

// SetRecipient is a paid mutator transaction binding the contract method 0xb33480ca.
//
// Solidity: function setRecipient(address _addr, bool allow) returns()
func (_Directbtcminter *DirectbtcminterSession) SetRecipient(_addr common.Address, allow bool) (*types.Transaction, error) {
	return _Directbtcminter.Contract.SetRecipient(&_Directbtcminter.TransactOpts, _addr, allow)
}

// SetRecipient is a paid mutator transaction binding the contract method 0xb33480ca.
//
// Solidity: function setRecipient(address _addr, bool allow) returns()
func (_Directbtcminter *DirectbtcminterTransactorSession) SetRecipient(_addr common.Address, allow bool) (*types.Transaction, error) {
	return _Directbtcminter.Contract.SetRecipient(&_Directbtcminter.TransactOpts, _addr, allow)
}

// DirectbtcminterAcceptedIterator is returned from FilterAccepted and is used to iterate over the raw logs and unpacked data for Accepted events raised by the Directbtcminter contract.
type DirectbtcminterAcceptedIterator struct {
	Event *DirectbtcminterAccepted // Event containing the contract specifics and raw log

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
func (it *DirectbtcminterAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DirectbtcminterAccepted)
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
		it.Event = new(DirectbtcminterAccepted)
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
func (it *DirectbtcminterAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DirectbtcminterAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DirectbtcminterAccepted represents a Accepted event raised by the Directbtcminter contract.
type DirectbtcminterAccepted struct {
	Recipient common.Address
	TxHash    [32]byte
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAccepted is a free log retrieval operation binding the contract event 0x4e3146457a9040b9825ace0a567115831772ecaf727b88bf9accb8eb876dd2ce.
//
// Solidity: event Accepted(address indexed recipient, bytes32 _txHash, uint256 _amount)
func (_Directbtcminter *DirectbtcminterFilterer) FilterAccepted(opts *bind.FilterOpts, recipient []common.Address) (*DirectbtcminterAcceptedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Directbtcminter.contract.FilterLogs(opts, "Accepted", recipientRule)
	if err != nil {
		return nil, err
	}
	return &DirectbtcminterAcceptedIterator{contract: _Directbtcminter.contract, event: "Accepted", logs: logs, sub: sub}, nil
}

// WatchAccepted is a free log subscription operation binding the contract event 0x4e3146457a9040b9825ace0a567115831772ecaf727b88bf9accb8eb876dd2ce.
//
// Solidity: event Accepted(address indexed recipient, bytes32 _txHash, uint256 _amount)
func (_Directbtcminter *DirectbtcminterFilterer) WatchAccepted(opts *bind.WatchOpts, sink chan<- *DirectbtcminterAccepted, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Directbtcminter.contract.WatchLogs(opts, "Accepted", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DirectbtcminterAccepted)
				if err := _Directbtcminter.contract.UnpackLog(event, "Accepted", log); err != nil {
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

// ParseAccepted is a log parse operation binding the contract event 0x4e3146457a9040b9825ace0a567115831772ecaf727b88bf9accb8eb876dd2ce.
//
// Solidity: event Accepted(address indexed recipient, bytes32 _txHash, uint256 _amount)
func (_Directbtcminter *DirectbtcminterFilterer) ParseAccepted(log types.Log) (*DirectbtcminterAccepted, error) {
	event := new(DirectbtcminterAccepted)
	if err := _Directbtcminter.contract.UnpackLog(event, "Accepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DirectbtcminterInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Directbtcminter contract.
type DirectbtcminterInitializedIterator struct {
	Event *DirectbtcminterInitialized // Event containing the contract specifics and raw log

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
func (it *DirectbtcminterInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DirectbtcminterInitialized)
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
		it.Event = new(DirectbtcminterInitialized)
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
func (it *DirectbtcminterInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DirectbtcminterInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DirectbtcminterInitialized represents a Initialized event raised by the Directbtcminter contract.
type DirectbtcminterInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Directbtcminter *DirectbtcminterFilterer) FilterInitialized(opts *bind.FilterOpts) (*DirectbtcminterInitializedIterator, error) {

	logs, sub, err := _Directbtcminter.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DirectbtcminterInitializedIterator{contract: _Directbtcminter.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Directbtcminter *DirectbtcminterFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DirectbtcminterInitialized) (event.Subscription, error) {

	logs, sub, err := _Directbtcminter.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DirectbtcminterInitialized)
				if err := _Directbtcminter.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Directbtcminter *DirectbtcminterFilterer) ParseInitialized(log types.Log) (*DirectbtcminterInitialized, error) {
	event := new(DirectbtcminterInitialized)
	if err := _Directbtcminter.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DirectbtcminterReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the Directbtcminter contract.
type DirectbtcminterReceivedIterator struct {
	Event *DirectbtcminterReceived // Event containing the contract specifics and raw log

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
func (it *DirectbtcminterReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DirectbtcminterReceived)
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
		it.Event = new(DirectbtcminterReceived)
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
func (it *DirectbtcminterReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DirectbtcminterReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DirectbtcminterReceived represents a Received event raised by the Directbtcminter contract.
type DirectbtcminterReceived struct {
	Recipient common.Address
	TxHash    [32]byte
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0x34b34fa9df7f5653d10fb7e1d3428378aba52f844a0c7c83b648b0a2e37ef7e8.
//
// Solidity: event Received(address indexed recipient, bytes32 _txHash, uint256 _amount)
func (_Directbtcminter *DirectbtcminterFilterer) FilterReceived(opts *bind.FilterOpts, recipient []common.Address) (*DirectbtcminterReceivedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Directbtcminter.contract.FilterLogs(opts, "Received", recipientRule)
	if err != nil {
		return nil, err
	}
	return &DirectbtcminterReceivedIterator{contract: _Directbtcminter.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0x34b34fa9df7f5653d10fb7e1d3428378aba52f844a0c7c83b648b0a2e37ef7e8.
//
// Solidity: event Received(address indexed recipient, bytes32 _txHash, uint256 _amount)
func (_Directbtcminter *DirectbtcminterFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *DirectbtcminterReceived, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Directbtcminter.contract.WatchLogs(opts, "Received", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DirectbtcminterReceived)
				if err := _Directbtcminter.contract.UnpackLog(event, "Received", log); err != nil {
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

// ParseReceived is a log parse operation binding the contract event 0x34b34fa9df7f5653d10fb7e1d3428378aba52f844a0c7c83b648b0a2e37ef7e8.
//
// Solidity: event Received(address indexed recipient, bytes32 _txHash, uint256 _amount)
func (_Directbtcminter *DirectbtcminterFilterer) ParseReceived(log types.Log) (*DirectbtcminterReceived, error) {
	event := new(DirectbtcminterReceived)
	if err := _Directbtcminter.contract.UnpackLog(event, "Received", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DirectbtcminterRecipientChangedIterator is returned from FilterRecipientChanged and is used to iterate over the raw logs and unpacked data for RecipientChanged events raised by the Directbtcminter contract.
type DirectbtcminterRecipientChangedIterator struct {
	Event *DirectbtcminterRecipientChanged // Event containing the contract specifics and raw log

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
func (it *DirectbtcminterRecipientChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DirectbtcminterRecipientChanged)
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
		it.Event = new(DirectbtcminterRecipientChanged)
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
func (it *DirectbtcminterRecipientChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DirectbtcminterRecipientChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DirectbtcminterRecipientChanged represents a RecipientChanged event raised by the Directbtcminter contract.
type DirectbtcminterRecipientChanged struct {
	Addr  common.Address
	Allow bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRecipientChanged is a free log retrieval operation binding the contract event 0xa18360beeed1d135b80e02eb97504ddfb82b97670b41d7a8007c21003eb2ec54.
//
// Solidity: event RecipientChanged(address indexed _addr, bool allow)
func (_Directbtcminter *DirectbtcminterFilterer) FilterRecipientChanged(opts *bind.FilterOpts, _addr []common.Address) (*DirectbtcminterRecipientChangedIterator, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}

	logs, sub, err := _Directbtcminter.contract.FilterLogs(opts, "RecipientChanged", _addrRule)
	if err != nil {
		return nil, err
	}
	return &DirectbtcminterRecipientChangedIterator{contract: _Directbtcminter.contract, event: "RecipientChanged", logs: logs, sub: sub}, nil
}

// WatchRecipientChanged is a free log subscription operation binding the contract event 0xa18360beeed1d135b80e02eb97504ddfb82b97670b41d7a8007c21003eb2ec54.
//
// Solidity: event RecipientChanged(address indexed _addr, bool allow)
func (_Directbtcminter *DirectbtcminterFilterer) WatchRecipientChanged(opts *bind.WatchOpts, sink chan<- *DirectbtcminterRecipientChanged, _addr []common.Address) (event.Subscription, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}

	logs, sub, err := _Directbtcminter.contract.WatchLogs(opts, "RecipientChanged", _addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DirectbtcminterRecipientChanged)
				if err := _Directbtcminter.contract.UnpackLog(event, "RecipientChanged", log); err != nil {
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

// ParseRecipientChanged is a log parse operation binding the contract event 0xa18360beeed1d135b80e02eb97504ddfb82b97670b41d7a8007c21003eb2ec54.
//
// Solidity: event RecipientChanged(address indexed _addr, bool allow)
func (_Directbtcminter *DirectbtcminterFilterer) ParseRecipientChanged(log types.Log) (*DirectbtcminterRecipientChanged, error) {
	event := new(DirectbtcminterRecipientChanged)
	if err := _Directbtcminter.contract.UnpackLog(event, "RecipientChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DirectbtcminterRejectedIterator is returned from FilterRejected and is used to iterate over the raw logs and unpacked data for Rejected events raised by the Directbtcminter contract.
type DirectbtcminterRejectedIterator struct {
	Event *DirectbtcminterRejected // Event containing the contract specifics and raw log

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
func (it *DirectbtcminterRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DirectbtcminterRejected)
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
		it.Event = new(DirectbtcminterRejected)
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
func (it *DirectbtcminterRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DirectbtcminterRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DirectbtcminterRejected represents a Rejected event raised by the Directbtcminter contract.
type DirectbtcminterRejected struct {
	Recipient common.Address
	TxHash    [32]byte
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRejected is a free log retrieval operation binding the contract event 0x63a4e409bcb8ca0d0cd26d68134b08efed23ca273e09d1d184eaa624c4c7ea7d.
//
// Solidity: event Rejected(address indexed recipient, bytes32 _txHash, uint256 _amount)
func (_Directbtcminter *DirectbtcminterFilterer) FilterRejected(opts *bind.FilterOpts, recipient []common.Address) (*DirectbtcminterRejectedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Directbtcminter.contract.FilterLogs(opts, "Rejected", recipientRule)
	if err != nil {
		return nil, err
	}
	return &DirectbtcminterRejectedIterator{contract: _Directbtcminter.contract, event: "Rejected", logs: logs, sub: sub}, nil
}

// WatchRejected is a free log subscription operation binding the contract event 0x63a4e409bcb8ca0d0cd26d68134b08efed23ca273e09d1d184eaa624c4c7ea7d.
//
// Solidity: event Rejected(address indexed recipient, bytes32 _txHash, uint256 _amount)
func (_Directbtcminter *DirectbtcminterFilterer) WatchRejected(opts *bind.WatchOpts, sink chan<- *DirectbtcminterRejected, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Directbtcminter.contract.WatchLogs(opts, "Rejected", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DirectbtcminterRejected)
				if err := _Directbtcminter.contract.UnpackLog(event, "Rejected", log); err != nil {
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

// ParseRejected is a log parse operation binding the contract event 0x63a4e409bcb8ca0d0cd26d68134b08efed23ca273e09d1d184eaa624c4c7ea7d.
//
// Solidity: event Rejected(address indexed recipient, bytes32 _txHash, uint256 _amount)
func (_Directbtcminter *DirectbtcminterFilterer) ParseRejected(log types.Log) (*DirectbtcminterRejected, error) {
	event := new(DirectbtcminterRejected)
	if err := _Directbtcminter.contract.UnpackLog(event, "Rejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DirectbtcminterRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Directbtcminter contract.
type DirectbtcminterRoleAdminChangedIterator struct {
	Event *DirectbtcminterRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *DirectbtcminterRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DirectbtcminterRoleAdminChanged)
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
		it.Event = new(DirectbtcminterRoleAdminChanged)
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
func (it *DirectbtcminterRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DirectbtcminterRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DirectbtcminterRoleAdminChanged represents a RoleAdminChanged event raised by the Directbtcminter contract.
type DirectbtcminterRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Directbtcminter *DirectbtcminterFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*DirectbtcminterRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Directbtcminter.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &DirectbtcminterRoleAdminChangedIterator{contract: _Directbtcminter.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Directbtcminter *DirectbtcminterFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *DirectbtcminterRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Directbtcminter.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DirectbtcminterRoleAdminChanged)
				if err := _Directbtcminter.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Directbtcminter *DirectbtcminterFilterer) ParseRoleAdminChanged(log types.Log) (*DirectbtcminterRoleAdminChanged, error) {
	event := new(DirectbtcminterRoleAdminChanged)
	if err := _Directbtcminter.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DirectbtcminterRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Directbtcminter contract.
type DirectbtcminterRoleGrantedIterator struct {
	Event *DirectbtcminterRoleGranted // Event containing the contract specifics and raw log

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
func (it *DirectbtcminterRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DirectbtcminterRoleGranted)
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
		it.Event = new(DirectbtcminterRoleGranted)
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
func (it *DirectbtcminterRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DirectbtcminterRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DirectbtcminterRoleGranted represents a RoleGranted event raised by the Directbtcminter contract.
type DirectbtcminterRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Directbtcminter *DirectbtcminterFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DirectbtcminterRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Directbtcminter.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DirectbtcminterRoleGrantedIterator{contract: _Directbtcminter.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Directbtcminter *DirectbtcminterFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *DirectbtcminterRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Directbtcminter.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DirectbtcminterRoleGranted)
				if err := _Directbtcminter.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Directbtcminter *DirectbtcminterFilterer) ParseRoleGranted(log types.Log) (*DirectbtcminterRoleGranted, error) {
	event := new(DirectbtcminterRoleGranted)
	if err := _Directbtcminter.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DirectbtcminterRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Directbtcminter contract.
type DirectbtcminterRoleRevokedIterator struct {
	Event *DirectbtcminterRoleRevoked // Event containing the contract specifics and raw log

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
func (it *DirectbtcminterRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DirectbtcminterRoleRevoked)
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
		it.Event = new(DirectbtcminterRoleRevoked)
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
func (it *DirectbtcminterRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DirectbtcminterRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DirectbtcminterRoleRevoked represents a RoleRevoked event raised by the Directbtcminter contract.
type DirectbtcminterRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Directbtcminter *DirectbtcminterFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DirectbtcminterRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Directbtcminter.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DirectbtcminterRoleRevokedIterator{contract: _Directbtcminter.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Directbtcminter *DirectbtcminterFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *DirectbtcminterRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Directbtcminter.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DirectbtcminterRoleRevoked)
				if err := _Directbtcminter.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Directbtcminter *DirectbtcminterFilterer) ParseRoleRevoked(log types.Log) (*DirectbtcminterRoleRevoked, error) {
	event := new(DirectbtcminterRoleRevoked)
	if err := _Directbtcminter.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
