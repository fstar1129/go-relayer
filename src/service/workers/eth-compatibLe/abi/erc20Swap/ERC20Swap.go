// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ERC20SwapABI is the input ABI used to generate the binding from.
const ERC20SwapABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc20Contract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_msgSender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_recipientAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_swapID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_randomNumberHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_randomNumber\",\"type\":\"bytes32\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_msgSender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_laRecipientAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_swapID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_randomNumberHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_timestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_ERC20TokenAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_LRC20TokenAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes20\",\"name\":\"_refundAddr\",\"type\":\"bytes20\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_expireHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_outAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_laAmount\",\"type\":\"uint256\"}],\"name\":\"HTLT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_msgSender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_recipientAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_swapID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_randomNumberHash\",\"type\":\"bytes32\"}],\"name\":\"Refunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_msgSender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_recipientAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_outAmount\",\"type\":\"uint256\"}],\"name\":\"Spended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_swapID\",\"type\":\"bytes32\"}],\"name\":\"SwapCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_swapID\",\"type\":\"bytes32\"}],\"name\":\"SwapFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_swapID\",\"type\":\"bytes32\"}],\"name\":\"SwapOpen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"paramChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_relayer\",\"type\":\"address\"}],\"name\":\"relayerRegister\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_relayer\",\"type\":\"address\"}],\"name\":\"relayerUnRegister\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ERC20Addr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_DUES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_REQUIRED_DEPOSIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYERHUB_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SYSTEM_REWARD_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"alreadyInit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_randomNumberHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_swapSender\",\"type\":\"address\"},{\"internalType\":\"bytes20\",\"name\":\"_LaSenderAddr\",\"type\":\"bytes20\"}],\"name\":\"calSwapID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"voters\",\"type\":\"uint256\"}],\"name\":\"checkQuorum\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_swapID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_randomNumber\",\"type\":\"bytes32\"}],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_swapID\",\"type\":\"bytes32\"}],\"name\":\"claimable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_swapID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_outAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_expireAt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_randomNumberHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_timestamp\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipientAddr\",\"type\":\"address\"}],\"name\":\"createOrClaim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createReq\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dues\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_randomNumberHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_heightSpan\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_laRecipientAddr\",\"type\":\"address\"},{\"internalType\":\"bytes20\",\"name\":\"_refundAddr\",\"type\":\"bytes20\"},{\"internalType\":\"address\",\"name\":\"_LRC20Addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_outAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_laAmount\",\"type\":\"uint256\"}],\"name\":\"htlt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"isRelayer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_swapID\",\"type\":\"bytes32\"}],\"name\":\"isSwapExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"onlyxRelayer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_swapID\",\"type\":\"bytes32\"}],\"name\":\"queryOpenSwap\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_randomNumberHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_expireHeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_outAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_swapID\",\"type\":\"bytes32\"}],\"name\":\"refund\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_swapID\",\"type\":\"bytes32\"}],\"name\":\"refundable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"relayersCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requiredDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reciever\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_outAmount\",\"type\":\"uint256\"}],\"name\":\"spend\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unregister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_swapID\",\"type\":\"bytes32\"},{\"internalType\":\"enumERC20AtomicSwapper.States\",\"name\":\"_state\",\"type\":\"uint8\"}],\"name\":\"voteState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20Swap is an auto generated Go binding around an Ethereum contract.
type ERC20Swap struct {
	ERC20SwapCaller     // Read-only binding to the contract
	ERC20SwapTransactor // Write-only binding to the contract
	ERC20SwapFilterer   // Log filterer for contract events
}

// ERC20SwapCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20SwapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20SwapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20SwapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20SwapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20SwapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20SwapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20SwapSession struct {
	Contract     *ERC20Swap        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20SwapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20SwapCallerSession struct {
	Contract *ERC20SwapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ERC20SwapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20SwapTransactorSession struct {
	Contract     *ERC20SwapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ERC20SwapRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20SwapRaw struct {
	Contract *ERC20Swap // Generic contract binding to access the raw methods on
}

// ERC20SwapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20SwapCallerRaw struct {
	Contract *ERC20SwapCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20SwapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20SwapTransactorRaw struct {
	Contract *ERC20SwapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Swap creates a new instance of ERC20Swap, bound to a specific deployed contract.
func NewERC20Swap(address common.Address, backend bind.ContractBackend) (*ERC20Swap, error) {
	contract, err := bindERC20Swap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Swap{ERC20SwapCaller: ERC20SwapCaller{contract: contract}, ERC20SwapTransactor: ERC20SwapTransactor{contract: contract}, ERC20SwapFilterer: ERC20SwapFilterer{contract: contract}}, nil
}

// NewERC20SwapCaller creates a new read-only instance of ERC20Swap, bound to a specific deployed contract.
func NewERC20SwapCaller(address common.Address, caller bind.ContractCaller) (*ERC20SwapCaller, error) {
	contract, err := bindERC20Swap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20SwapCaller{contract: contract}, nil
}

// NewERC20SwapTransactor creates a new write-only instance of ERC20Swap, bound to a specific deployed contract.
func NewERC20SwapTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20SwapTransactor, error) {
	contract, err := bindERC20Swap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20SwapTransactor{contract: contract}, nil
}

// NewERC20SwapFilterer creates a new log filterer instance of ERC20Swap, bound to a specific deployed contract.
func NewERC20SwapFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20SwapFilterer, error) {
	contract, err := bindERC20Swap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20SwapFilterer{contract: contract}, nil
}

// bindERC20Swap binds a generic wrapper to an already deployed contract.
func bindERC20Swap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20SwapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Swap *ERC20SwapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Swap.Contract.ERC20SwapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Swap *ERC20SwapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Swap.Contract.ERC20SwapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Swap *ERC20SwapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Swap.Contract.ERC20SwapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Swap *ERC20SwapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Swap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Swap *ERC20SwapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Swap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Swap *ERC20SwapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Swap.Contract.contract.Transact(opts, method, params...)
}

// ERC20Addr is a free data retrieval call binding the contract method 0x39fe6bdb.
//
// Solidity: function ERC20Addr() view returns(address)
func (_ERC20Swap *ERC20SwapCaller) ERC20Addr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Swap.contract.Call(opts, &out, "ERC20Addr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ERC20Addr is a free data retrieval call binding the contract method 0x39fe6bdb.
//
// Solidity: function ERC20Addr() view returns(address)
func (_ERC20Swap *ERC20SwapSession) ERC20Addr() (common.Address, error) {
	return _ERC20Swap.Contract.ERC20Addr(&_ERC20Swap.CallOpts)
}

// ERC20Addr is a free data retrieval call binding the contract method 0x39fe6bdb.
//
// Solidity: function ERC20Addr() view returns(address)
func (_ERC20Swap *ERC20SwapCallerSession) ERC20Addr() (common.Address, error) {
	return _ERC20Swap.Contract.ERC20Addr(&_ERC20Swap.CallOpts)
}

// INITDUES is a free data retrieval call binding the contract method 0x95468d26.
//
// Solidity: function INIT_DUES() view returns(uint256)
func (_ERC20Swap *ERC20SwapCaller) INITDUES(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Swap.contract.Call(opts, &out, "INIT_DUES")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITDUES is a free data retrieval call binding the contract method 0x95468d26.
//
// Solidity: function INIT_DUES() view returns(uint256)
func (_ERC20Swap *ERC20SwapSession) INITDUES() (*big.Int, error) {
	return _ERC20Swap.Contract.INITDUES(&_ERC20Swap.CallOpts)
}

// INITDUES is a free data retrieval call binding the contract method 0x95468d26.
//
// Solidity: function INIT_DUES() view returns(uint256)
func (_ERC20Swap *ERC20SwapCallerSession) INITDUES() (*big.Int, error) {
	return _ERC20Swap.Contract.INITDUES(&_ERC20Swap.CallOpts)
}

// INITREQUIREDDEPOSIT is a free data retrieval call binding the contract method 0x7ae23088.
//
// Solidity: function INIT_REQUIRED_DEPOSIT() view returns(uint256)
func (_ERC20Swap *ERC20SwapCaller) INITREQUIREDDEPOSIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Swap.contract.Call(opts, &out, "INIT_REQUIRED_DEPOSIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITREQUIREDDEPOSIT is a free data retrieval call binding the contract method 0x7ae23088.
//
// Solidity: function INIT_REQUIRED_DEPOSIT() view returns(uint256)
func (_ERC20Swap *ERC20SwapSession) INITREQUIREDDEPOSIT() (*big.Int, error) {
	return _ERC20Swap.Contract.INITREQUIREDDEPOSIT(&_ERC20Swap.CallOpts)
}

// INITREQUIREDDEPOSIT is a free data retrieval call binding the contract method 0x7ae23088.
//
// Solidity: function INIT_REQUIRED_DEPOSIT() view returns(uint256)
func (_ERC20Swap *ERC20SwapCallerSession) INITREQUIREDDEPOSIT() (*big.Int, error) {
	return _ERC20Swap.Contract.INITREQUIREDDEPOSIT(&_ERC20Swap.CallOpts)
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() view returns(address)
func (_ERC20Swap *ERC20SwapCaller) RELAYERHUBCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Swap.contract.Call(opts, &out, "RELAYERHUB_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() view returns(address)
func (_ERC20Swap *ERC20SwapSession) RELAYERHUBCONTRACTADDR() (common.Address, error) {
	return _ERC20Swap.Contract.RELAYERHUBCONTRACTADDR(&_ERC20Swap.CallOpts)
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() view returns(address)
func (_ERC20Swap *ERC20SwapCallerSession) RELAYERHUBCONTRACTADDR() (common.Address, error) {
	return _ERC20Swap.Contract.RELAYERHUBCONTRACTADDR(&_ERC20Swap.CallOpts)
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() view returns(address)
func (_ERC20Swap *ERC20SwapCaller) SYSTEMREWARDADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Swap.contract.Call(opts, &out, "SYSTEM_REWARD_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() view returns(address)
func (_ERC20Swap *ERC20SwapSession) SYSTEMREWARDADDR() (common.Address, error) {
	return _ERC20Swap.Contract.SYSTEMREWARDADDR(&_ERC20Swap.CallOpts)
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() view returns(address)
func (_ERC20Swap *ERC20SwapCallerSession) SYSTEMREWARDADDR() (common.Address, error) {
	return _ERC20Swap.Contract.SYSTEMREWARDADDR(&_ERC20Swap.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_ERC20Swap *ERC20SwapCaller) AlreadyInit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC20Swap.contract.Call(opts, &out, "alreadyInit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_ERC20Swap *ERC20SwapSession) AlreadyInit() (bool, error) {
	return _ERC20Swap.Contract.AlreadyInit(&_ERC20Swap.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_ERC20Swap *ERC20SwapCallerSession) AlreadyInit() (bool, error) {
	return _ERC20Swap.Contract.AlreadyInit(&_ERC20Swap.CallOpts)
}

// CalSwapID is a free data retrieval call binding the contract method 0x7ef3e92e.
//
// Solidity: function calSwapID(bytes32 _randomNumberHash, address _swapSender, bytes20 _LaSenderAddr) pure returns(bytes32)
func (_ERC20Swap *ERC20SwapCaller) CalSwapID(opts *bind.CallOpts, _randomNumberHash [32]byte, _swapSender common.Address, _LaSenderAddr [20]byte) ([32]byte, error) {
	var out []interface{}
	err := _ERC20Swap.contract.Call(opts, &out, "calSwapID", _randomNumberHash, _swapSender, _LaSenderAddr)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalSwapID is a free data retrieval call binding the contract method 0x7ef3e92e.
//
// Solidity: function calSwapID(bytes32 _randomNumberHash, address _swapSender, bytes20 _LaSenderAddr) pure returns(bytes32)
func (_ERC20Swap *ERC20SwapSession) CalSwapID(_randomNumberHash [32]byte, _swapSender common.Address, _LaSenderAddr [20]byte) ([32]byte, error) {
	return _ERC20Swap.Contract.CalSwapID(&_ERC20Swap.CallOpts, _randomNumberHash, _swapSender, _LaSenderAddr)
}

// CalSwapID is a free data retrieval call binding the contract method 0x7ef3e92e.
//
// Solidity: function calSwapID(bytes32 _randomNumberHash, address _swapSender, bytes20 _LaSenderAddr) pure returns(bytes32)
func (_ERC20Swap *ERC20SwapCallerSession) CalSwapID(_randomNumberHash [32]byte, _swapSender common.Address, _LaSenderAddr [20]byte) ([32]byte, error) {
	return _ERC20Swap.Contract.CalSwapID(&_ERC20Swap.CallOpts, _randomNumberHash, _swapSender, _LaSenderAddr)
}

// CheckQuorum is a free data retrieval call binding the contract method 0x9fc3b4e7.
//
// Solidity: function checkQuorum(uint256 voters) view returns(bool)
func (_ERC20Swap *ERC20SwapCaller) CheckQuorum(opts *bind.CallOpts, voters *big.Int) (bool, error) {
	var out []interface{}
	err := _ERC20Swap.contract.Call(opts, &out, "checkQuorum", voters)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckQuorum is a free data retrieval call binding the contract method 0x9fc3b4e7.
//
// Solidity: function checkQuorum(uint256 voters) view returns(bool)
func (_ERC20Swap *ERC20SwapSession) CheckQuorum(voters *big.Int) (bool, error) {
	return _ERC20Swap.Contract.CheckQuorum(&_ERC20Swap.CallOpts, voters)
}

// CheckQuorum is a free data retrieval call binding the contract method 0x9fc3b4e7.
//
// Solidity: function checkQuorum(uint256 voters) view returns(bool)
func (_ERC20Swap *ERC20SwapCallerSession) CheckQuorum(voters *big.Int) (bool, error) {
	return _ERC20Swap.Contract.CheckQuorum(&_ERC20Swap.CallOpts, voters)
}

// Claimable is a free data retrieval call binding the contract method 0x9b58e0a1.
//
// Solidity: function claimable(bytes32 _swapID) view returns(bool)
func (_ERC20Swap *ERC20SwapCaller) Claimable(opts *bind.CallOpts, _swapID [32]byte) (bool, error) {
	var out []interface{}
	err := _ERC20Swap.contract.Call(opts, &out, "claimable", _swapID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Claimable is a free data retrieval call binding the contract method 0x9b58e0a1.
//
// Solidity: function claimable(bytes32 _swapID) view returns(bool)
func (_ERC20Swap *ERC20SwapSession) Claimable(_swapID [32]byte) (bool, error) {
	return _ERC20Swap.Contract.Claimable(&_ERC20Swap.CallOpts, _swapID)
}

// Claimable is a free data retrieval call binding the contract method 0x9b58e0a1.
//
// Solidity: function claimable(bytes32 _swapID) view returns(bool)
func (_ERC20Swap *ERC20SwapCallerSession) Claimable(_swapID [32]byte) (bool, error) {
	return _ERC20Swap.Contract.Claimable(&_ERC20Swap.CallOpts, _swapID)
}

// CreateReq is a free data retrieval call binding the contract method 0x40f0f433.
//
// Solidity: function createReq() view returns(bool)
func (_ERC20Swap *ERC20SwapCaller) CreateReq(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC20Swap.contract.Call(opts, &out, "createReq")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CreateReq is a free data retrieval call binding the contract method 0x40f0f433.
//
// Solidity: function createReq() view returns(bool)
func (_ERC20Swap *ERC20SwapSession) CreateReq() (bool, error) {
	return _ERC20Swap.Contract.CreateReq(&_ERC20Swap.CallOpts)
}

// CreateReq is a free data retrieval call binding the contract method 0x40f0f433.
//
// Solidity: function createReq() view returns(bool)
func (_ERC20Swap *ERC20SwapCallerSession) CreateReq() (bool, error) {
	return _ERC20Swap.Contract.CreateReq(&_ERC20Swap.CallOpts)
}

// Dues is a free data retrieval call binding the contract method 0x6a87d780.
//
// Solidity: function dues() view returns(uint256)
func (_ERC20Swap *ERC20SwapCaller) Dues(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Swap.contract.Call(opts, &out, "dues")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dues is a free data retrieval call binding the contract method 0x6a87d780.
//
// Solidity: function dues() view returns(uint256)
func (_ERC20Swap *ERC20SwapSession) Dues() (*big.Int, error) {
	return _ERC20Swap.Contract.Dues(&_ERC20Swap.CallOpts)
}

// Dues is a free data retrieval call binding the contract method 0x6a87d780.
//
// Solidity: function dues() view returns(uint256)
func (_ERC20Swap *ERC20SwapCallerSession) Dues() (*big.Int, error) {
	return _ERC20Swap.Contract.Dues(&_ERC20Swap.CallOpts)
}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address sender) view returns(bool)
func (_ERC20Swap *ERC20SwapCaller) IsRelayer(opts *bind.CallOpts, sender common.Address) (bool, error) {
	var out []interface{}
	err := _ERC20Swap.contract.Call(opts, &out, "isRelayer", sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address sender) view returns(bool)
func (_ERC20Swap *ERC20SwapSession) IsRelayer(sender common.Address) (bool, error) {
	return _ERC20Swap.Contract.IsRelayer(&_ERC20Swap.CallOpts, sender)
}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address sender) view returns(bool)
func (_ERC20Swap *ERC20SwapCallerSession) IsRelayer(sender common.Address) (bool, error) {
	return _ERC20Swap.Contract.IsRelayer(&_ERC20Swap.CallOpts, sender)
}

// IsSwapExist is a free data retrieval call binding the contract method 0x50f7a03b.
//
// Solidity: function isSwapExist(bytes32 _swapID) view returns(bool)
func (_ERC20Swap *ERC20SwapCaller) IsSwapExist(opts *bind.CallOpts, _swapID [32]byte) (bool, error) {
	var out []interface{}
	err := _ERC20Swap.contract.Call(opts, &out, "isSwapExist", _swapID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSwapExist is a free data retrieval call binding the contract method 0x50f7a03b.
//
// Solidity: function isSwapExist(bytes32 _swapID) view returns(bool)
func (_ERC20Swap *ERC20SwapSession) IsSwapExist(_swapID [32]byte) (bool, error) {
	return _ERC20Swap.Contract.IsSwapExist(&_ERC20Swap.CallOpts, _swapID)
}

// IsSwapExist is a free data retrieval call binding the contract method 0x50f7a03b.
//
// Solidity: function isSwapExist(bytes32 _swapID) view returns(bool)
func (_ERC20Swap *ERC20SwapCallerSession) IsSwapExist(_swapID [32]byte) (bool, error) {
	return _ERC20Swap.Contract.IsSwapExist(&_ERC20Swap.CallOpts, _swapID)
}

// OnlyxRelayer is a free data retrieval call binding the contract method 0x9766030b.
//
// Solidity: function onlyxRelayer() view returns(bool)
func (_ERC20Swap *ERC20SwapCaller) OnlyxRelayer(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC20Swap.contract.Call(opts, &out, "onlyxRelayer")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OnlyxRelayer is a free data retrieval call binding the contract method 0x9766030b.
//
// Solidity: function onlyxRelayer() view returns(bool)
func (_ERC20Swap *ERC20SwapSession) OnlyxRelayer() (bool, error) {
	return _ERC20Swap.Contract.OnlyxRelayer(&_ERC20Swap.CallOpts)
}

// OnlyxRelayer is a free data retrieval call binding the contract method 0x9766030b.
//
// Solidity: function onlyxRelayer() view returns(bool)
func (_ERC20Swap *ERC20SwapCallerSession) OnlyxRelayer() (bool, error) {
	return _ERC20Swap.Contract.OnlyxRelayer(&_ERC20Swap.CallOpts)
}

// QueryOpenSwap is a free data retrieval call binding the contract method 0xb48017b1.
//
// Solidity: function queryOpenSwap(bytes32 _swapID) view returns(bytes32 _randomNumberHash, uint64 _timestamp, uint256 _expireHeight, uint256 _outAmount, address _sender, address _recipient)
func (_ERC20Swap *ERC20SwapCaller) QueryOpenSwap(opts *bind.CallOpts, _swapID [32]byte) (struct {
	RandomNumberHash [32]byte
	Timestamp        uint64
	ExpireHeight     *big.Int
	OutAmount        *big.Int
	Sender           common.Address
	Recipient        common.Address
}, error) {
	var out []interface{}
	err := _ERC20Swap.contract.Call(opts, &out, "queryOpenSwap", _swapID)

	outstruct := new(struct {
		RandomNumberHash [32]byte
		Timestamp        uint64
		ExpireHeight     *big.Int
		OutAmount        *big.Int
		Sender           common.Address
		Recipient        common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RandomNumberHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.ExpireHeight = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.OutAmount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Sender = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Recipient = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// QueryOpenSwap is a free data retrieval call binding the contract method 0xb48017b1.
//
// Solidity: function queryOpenSwap(bytes32 _swapID) view returns(bytes32 _randomNumberHash, uint64 _timestamp, uint256 _expireHeight, uint256 _outAmount, address _sender, address _recipient)
func (_ERC20Swap *ERC20SwapSession) QueryOpenSwap(_swapID [32]byte) (struct {
	RandomNumberHash [32]byte
	Timestamp        uint64
	ExpireHeight     *big.Int
	OutAmount        *big.Int
	Sender           common.Address
	Recipient        common.Address
}, error) {
	return _ERC20Swap.Contract.QueryOpenSwap(&_ERC20Swap.CallOpts, _swapID)
}

// QueryOpenSwap is a free data retrieval call binding the contract method 0xb48017b1.
//
// Solidity: function queryOpenSwap(bytes32 _swapID) view returns(bytes32 _randomNumberHash, uint64 _timestamp, uint256 _expireHeight, uint256 _outAmount, address _sender, address _recipient)
func (_ERC20Swap *ERC20SwapCallerSession) QueryOpenSwap(_swapID [32]byte) (struct {
	RandomNumberHash [32]byte
	Timestamp        uint64
	ExpireHeight     *big.Int
	OutAmount        *big.Int
	Sender           common.Address
	Recipient        common.Address
}, error) {
	return _ERC20Swap.Contract.QueryOpenSwap(&_ERC20Swap.CallOpts, _swapID)
}

// Refundable is a free data retrieval call binding the contract method 0x9fb31475.
//
// Solidity: function refundable(bytes32 _swapID) view returns(bool)
func (_ERC20Swap *ERC20SwapCaller) Refundable(opts *bind.CallOpts, _swapID [32]byte) (bool, error) {
	var out []interface{}
	err := _ERC20Swap.contract.Call(opts, &out, "refundable", _swapID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Refundable is a free data retrieval call binding the contract method 0x9fb31475.
//
// Solidity: function refundable(bytes32 _swapID) view returns(bool)
func (_ERC20Swap *ERC20SwapSession) Refundable(_swapID [32]byte) (bool, error) {
	return _ERC20Swap.Contract.Refundable(&_ERC20Swap.CallOpts, _swapID)
}

// Refundable is a free data retrieval call binding the contract method 0x9fb31475.
//
// Solidity: function refundable(bytes32 _swapID) view returns(bool)
func (_ERC20Swap *ERC20SwapCallerSession) Refundable(_swapID [32]byte) (bool, error) {
	return _ERC20Swap.Contract.Refundable(&_ERC20Swap.CallOpts, _swapID)
}

// RelayersCount is a free data retrieval call binding the contract method 0xaae8e9ed.
//
// Solidity: function relayersCount() view returns(uint256)
func (_ERC20Swap *ERC20SwapCaller) RelayersCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Swap.contract.Call(opts, &out, "relayersCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RelayersCount is a free data retrieval call binding the contract method 0xaae8e9ed.
//
// Solidity: function relayersCount() view returns(uint256)
func (_ERC20Swap *ERC20SwapSession) RelayersCount() (*big.Int, error) {
	return _ERC20Swap.Contract.RelayersCount(&_ERC20Swap.CallOpts)
}

// RelayersCount is a free data retrieval call binding the contract method 0xaae8e9ed.
//
// Solidity: function relayersCount() view returns(uint256)
func (_ERC20Swap *ERC20SwapCallerSession) RelayersCount() (*big.Int, error) {
	return _ERC20Swap.Contract.RelayersCount(&_ERC20Swap.CallOpts)
}

// RequiredDeposit is a free data retrieval call binding the contract method 0xfb7cfdd7.
//
// Solidity: function requiredDeposit() view returns(uint256)
func (_ERC20Swap *ERC20SwapCaller) RequiredDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Swap.contract.Call(opts, &out, "requiredDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequiredDeposit is a free data retrieval call binding the contract method 0xfb7cfdd7.
//
// Solidity: function requiredDeposit() view returns(uint256)
func (_ERC20Swap *ERC20SwapSession) RequiredDeposit() (*big.Int, error) {
	return _ERC20Swap.Contract.RequiredDeposit(&_ERC20Swap.CallOpts)
}

// RequiredDeposit is a free data retrieval call binding the contract method 0xfb7cfdd7.
//
// Solidity: function requiredDeposit() view returns(uint256)
func (_ERC20Swap *ERC20SwapCallerSession) RequiredDeposit() (*big.Int, error) {
	return _ERC20Swap.Contract.RequiredDeposit(&_ERC20Swap.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x84cc9dfb.
//
// Solidity: function claim(bytes32 _swapID, bytes32 _randomNumber) returns(bool)
func (_ERC20Swap *ERC20SwapTransactor) Claim(opts *bind.TransactOpts, _swapID [32]byte, _randomNumber [32]byte) (*types.Transaction, error) {
	return _ERC20Swap.contract.Transact(opts, "claim", _swapID, _randomNumber)
}

// Claim is a paid mutator transaction binding the contract method 0x84cc9dfb.
//
// Solidity: function claim(bytes32 _swapID, bytes32 _randomNumber) returns(bool)
func (_ERC20Swap *ERC20SwapSession) Claim(_swapID [32]byte, _randomNumber [32]byte) (*types.Transaction, error) {
	return _ERC20Swap.Contract.Claim(&_ERC20Swap.TransactOpts, _swapID, _randomNumber)
}

// Claim is a paid mutator transaction binding the contract method 0x84cc9dfb.
//
// Solidity: function claim(bytes32 _swapID, bytes32 _randomNumber) returns(bool)
func (_ERC20Swap *ERC20SwapTransactorSession) Claim(_swapID [32]byte, _randomNumber [32]byte) (*types.Transaction, error) {
	return _ERC20Swap.Contract.Claim(&_ERC20Swap.TransactOpts, _swapID, _randomNumber)
}

// CreateOrClaim is a paid mutator transaction binding the contract method 0x44844516.
//
// Solidity: function createOrClaim(bytes32 _swapID, uint256 _outAmount, uint256 _expireAt, bytes32 _randomNumberHash, uint64 _timestamp, address _token, address _recipientAddr) returns(bool)
func (_ERC20Swap *ERC20SwapTransactor) CreateOrClaim(opts *bind.TransactOpts, _swapID [32]byte, _outAmount *big.Int, _expireAt *big.Int, _randomNumberHash [32]byte, _timestamp uint64, _token common.Address, _recipientAddr common.Address) (*types.Transaction, error) {
	return _ERC20Swap.contract.Transact(opts, "createOrClaim", _swapID, _outAmount, _expireAt, _randomNumberHash, _timestamp, _token, _recipientAddr)
}

// CreateOrClaim is a paid mutator transaction binding the contract method 0x44844516.
//
// Solidity: function createOrClaim(bytes32 _swapID, uint256 _outAmount, uint256 _expireAt, bytes32 _randomNumberHash, uint64 _timestamp, address _token, address _recipientAddr) returns(bool)
func (_ERC20Swap *ERC20SwapSession) CreateOrClaim(_swapID [32]byte, _outAmount *big.Int, _expireAt *big.Int, _randomNumberHash [32]byte, _timestamp uint64, _token common.Address, _recipientAddr common.Address) (*types.Transaction, error) {
	return _ERC20Swap.Contract.CreateOrClaim(&_ERC20Swap.TransactOpts, _swapID, _outAmount, _expireAt, _randomNumberHash, _timestamp, _token, _recipientAddr)
}

// CreateOrClaim is a paid mutator transaction binding the contract method 0x44844516.
//
// Solidity: function createOrClaim(bytes32 _swapID, uint256 _outAmount, uint256 _expireAt, bytes32 _randomNumberHash, uint64 _timestamp, address _token, address _recipientAddr) returns(bool)
func (_ERC20Swap *ERC20SwapTransactorSession) CreateOrClaim(_swapID [32]byte, _outAmount *big.Int, _expireAt *big.Int, _randomNumberHash [32]byte, _timestamp uint64, _token common.Address, _recipientAddr common.Address) (*types.Transaction, error) {
	return _ERC20Swap.Contract.CreateOrClaim(&_ERC20Swap.TransactOpts, _swapID, _outAmount, _expireAt, _randomNumberHash, _timestamp, _token, _recipientAddr)
}

// Htlt is a paid mutator transaction binding the contract method 0xb6e6bd3e.
//
// Solidity: function htlt(bytes32 _randomNumberHash, uint64 _timestamp, uint256 _heightSpan, address _laRecipientAddr, bytes20 _refundAddr, address _LRC20Addr, uint256 _outAmount, uint256 _laAmount) returns(bool)
func (_ERC20Swap *ERC20SwapTransactor) Htlt(opts *bind.TransactOpts, _randomNumberHash [32]byte, _timestamp uint64, _heightSpan *big.Int, _laRecipientAddr common.Address, _refundAddr [20]byte, _LRC20Addr common.Address, _outAmount *big.Int, _laAmount *big.Int) (*types.Transaction, error) {
	return _ERC20Swap.contract.Transact(opts, "htlt", _randomNumberHash, _timestamp, _heightSpan, _laRecipientAddr, _refundAddr, _LRC20Addr, _outAmount, _laAmount)
}

// Htlt is a paid mutator transaction binding the contract method 0xb6e6bd3e.
//
// Solidity: function htlt(bytes32 _randomNumberHash, uint64 _timestamp, uint256 _heightSpan, address _laRecipientAddr, bytes20 _refundAddr, address _LRC20Addr, uint256 _outAmount, uint256 _laAmount) returns(bool)
func (_ERC20Swap *ERC20SwapSession) Htlt(_randomNumberHash [32]byte, _timestamp uint64, _heightSpan *big.Int, _laRecipientAddr common.Address, _refundAddr [20]byte, _LRC20Addr common.Address, _outAmount *big.Int, _laAmount *big.Int) (*types.Transaction, error) {
	return _ERC20Swap.Contract.Htlt(&_ERC20Swap.TransactOpts, _randomNumberHash, _timestamp, _heightSpan, _laRecipientAddr, _refundAddr, _LRC20Addr, _outAmount, _laAmount)
}

// Htlt is a paid mutator transaction binding the contract method 0xb6e6bd3e.
//
// Solidity: function htlt(bytes32 _randomNumberHash, uint64 _timestamp, uint256 _heightSpan, address _laRecipientAddr, bytes20 _refundAddr, address _LRC20Addr, uint256 _outAmount, uint256 _laAmount) returns(bool)
func (_ERC20Swap *ERC20SwapTransactorSession) Htlt(_randomNumberHash [32]byte, _timestamp uint64, _heightSpan *big.Int, _laRecipientAddr common.Address, _refundAddr [20]byte, _LRC20Addr common.Address, _outAmount *big.Int, _laAmount *big.Int) (*types.Transaction, error) {
	return _ERC20Swap.Contract.Htlt(&_ERC20Swap.TransactOpts, _randomNumberHash, _timestamp, _heightSpan, _laRecipientAddr, _refundAddr, _LRC20Addr, _outAmount, _laAmount)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_ERC20Swap *ERC20SwapTransactor) Init(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Swap.contract.Transact(opts, "init")
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_ERC20Swap *ERC20SwapSession) Init() (*types.Transaction, error) {
	return _ERC20Swap.Contract.Init(&_ERC20Swap.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_ERC20Swap *ERC20SwapTransactorSession) Init() (*types.Transaction, error) {
	return _ERC20Swap.Contract.Init(&_ERC20Swap.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0x7249fbb6.
//
// Solidity: function refund(bytes32 _swapID) returns(bool)
func (_ERC20Swap *ERC20SwapTransactor) Refund(opts *bind.TransactOpts, _swapID [32]byte) (*types.Transaction, error) {
	return _ERC20Swap.contract.Transact(opts, "refund", _swapID)
}

// Refund is a paid mutator transaction binding the contract method 0x7249fbb6.
//
// Solidity: function refund(bytes32 _swapID) returns(bool)
func (_ERC20Swap *ERC20SwapSession) Refund(_swapID [32]byte) (*types.Transaction, error) {
	return _ERC20Swap.Contract.Refund(&_ERC20Swap.TransactOpts, _swapID)
}

// Refund is a paid mutator transaction binding the contract method 0x7249fbb6.
//
// Solidity: function refund(bytes32 _swapID) returns(bool)
func (_ERC20Swap *ERC20SwapTransactorSession) Refund(_swapID [32]byte) (*types.Transaction, error) {
	return _ERC20Swap.Contract.Refund(&_ERC20Swap.TransactOpts, _swapID)
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() payable returns()
func (_ERC20Swap *ERC20SwapTransactor) Register(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Swap.contract.Transact(opts, "register")
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() payable returns()
func (_ERC20Swap *ERC20SwapSession) Register() (*types.Transaction, error) {
	return _ERC20Swap.Contract.Register(&_ERC20Swap.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() payable returns()
func (_ERC20Swap *ERC20SwapTransactorSession) Register() (*types.Transaction, error) {
	return _ERC20Swap.Contract.Register(&_ERC20Swap.TransactOpts)
}

// Spend is a paid mutator transaction binding the contract method 0xaf7d6ca3.
//
// Solidity: function spend(address _reciever, uint256 _outAmount) returns(bool)
func (_ERC20Swap *ERC20SwapTransactor) Spend(opts *bind.TransactOpts, _reciever common.Address, _outAmount *big.Int) (*types.Transaction, error) {
	return _ERC20Swap.contract.Transact(opts, "spend", _reciever, _outAmount)
}

// Spend is a paid mutator transaction binding the contract method 0xaf7d6ca3.
//
// Solidity: function spend(address _reciever, uint256 _outAmount) returns(bool)
func (_ERC20Swap *ERC20SwapSession) Spend(_reciever common.Address, _outAmount *big.Int) (*types.Transaction, error) {
	return _ERC20Swap.Contract.Spend(&_ERC20Swap.TransactOpts, _reciever, _outAmount)
}

// Spend is a paid mutator transaction binding the contract method 0xaf7d6ca3.
//
// Solidity: function spend(address _reciever, uint256 _outAmount) returns(bool)
func (_ERC20Swap *ERC20SwapTransactorSession) Spend(_reciever common.Address, _outAmount *big.Int) (*types.Transaction, error) {
	return _ERC20Swap.Contract.Spend(&_ERC20Swap.TransactOpts, _reciever, _outAmount)
}

// Unregister is a paid mutator transaction binding the contract method 0xe79a198f.
//
// Solidity: function unregister() returns()
func (_ERC20Swap *ERC20SwapTransactor) Unregister(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Swap.contract.Transact(opts, "unregister")
}

// Unregister is a paid mutator transaction binding the contract method 0xe79a198f.
//
// Solidity: function unregister() returns()
func (_ERC20Swap *ERC20SwapSession) Unregister() (*types.Transaction, error) {
	return _ERC20Swap.Contract.Unregister(&_ERC20Swap.TransactOpts)
}

// Unregister is a paid mutator transaction binding the contract method 0xe79a198f.
//
// Solidity: function unregister() returns()
func (_ERC20Swap *ERC20SwapTransactorSession) Unregister() (*types.Transaction, error) {
	return _ERC20Swap.Contract.Unregister(&_ERC20Swap.TransactOpts)
}

// VoteState is a paid mutator transaction binding the contract method 0x70d55c69.
//
// Solidity: function voteState(bytes32 _swapID, uint8 _state) returns()
func (_ERC20Swap *ERC20SwapTransactor) VoteState(opts *bind.TransactOpts, _swapID [32]byte, _state uint8) (*types.Transaction, error) {
	return _ERC20Swap.contract.Transact(opts, "voteState", _swapID, _state)
}

// VoteState is a paid mutator transaction binding the contract method 0x70d55c69.
//
// Solidity: function voteState(bytes32 _swapID, uint8 _state) returns()
func (_ERC20Swap *ERC20SwapSession) VoteState(_swapID [32]byte, _state uint8) (*types.Transaction, error) {
	return _ERC20Swap.Contract.VoteState(&_ERC20Swap.TransactOpts, _swapID, _state)
}

// VoteState is a paid mutator transaction binding the contract method 0x70d55c69.
//
// Solidity: function voteState(bytes32 _swapID, uint8 _state) returns()
func (_ERC20Swap *ERC20SwapTransactorSession) VoteState(_swapID [32]byte, _state uint8) (*types.Transaction, error) {
	return _ERC20Swap.Contract.VoteState(&_ERC20Swap.TransactOpts, _swapID, _state)
}

// ERC20SwapClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the ERC20Swap contract.
type ERC20SwapClaimedIterator struct {
	Event *ERC20SwapClaimed // Event containing the contract specifics and raw log

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
func (it *ERC20SwapClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20SwapClaimed)
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
		it.Event = new(ERC20SwapClaimed)
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
func (it *ERC20SwapClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20SwapClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20SwapClaimed represents a Claimed event raised by the ERC20Swap contract.
type ERC20SwapClaimed struct {
	MsgSender        common.Address
	RecipientAddr    common.Address
	SwapID           [32]byte
	RandomNumberHash [32]byte
	RandomNumber     [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0x9f46b1606087bdf4183ec7dfdbe68e4ab9129a6a37901c16a7b320ae11a96018.
//
// Solidity: event Claimed(address indexed _msgSender, address indexed _recipientAddr, bytes32 indexed _swapID, bytes32 _randomNumberHash, bytes32 _randomNumber)
func (_ERC20Swap *ERC20SwapFilterer) FilterClaimed(opts *bind.FilterOpts, _msgSender []common.Address, _recipientAddr []common.Address, _swapID [][32]byte) (*ERC20SwapClaimedIterator, error) {

	var _msgSenderRule []interface{}
	for _, _msgSenderItem := range _msgSender {
		_msgSenderRule = append(_msgSenderRule, _msgSenderItem)
	}
	var _recipientAddrRule []interface{}
	for _, _recipientAddrItem := range _recipientAddr {
		_recipientAddrRule = append(_recipientAddrRule, _recipientAddrItem)
	}
	var _swapIDRule []interface{}
	for _, _swapIDItem := range _swapID {
		_swapIDRule = append(_swapIDRule, _swapIDItem)
	}

	logs, sub, err := _ERC20Swap.contract.FilterLogs(opts, "Claimed", _msgSenderRule, _recipientAddrRule, _swapIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20SwapClaimedIterator{contract: _ERC20Swap.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0x9f46b1606087bdf4183ec7dfdbe68e4ab9129a6a37901c16a7b320ae11a96018.
//
// Solidity: event Claimed(address indexed _msgSender, address indexed _recipientAddr, bytes32 indexed _swapID, bytes32 _randomNumberHash, bytes32 _randomNumber)
func (_ERC20Swap *ERC20SwapFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *ERC20SwapClaimed, _msgSender []common.Address, _recipientAddr []common.Address, _swapID [][32]byte) (event.Subscription, error) {

	var _msgSenderRule []interface{}
	for _, _msgSenderItem := range _msgSender {
		_msgSenderRule = append(_msgSenderRule, _msgSenderItem)
	}
	var _recipientAddrRule []interface{}
	for _, _recipientAddrItem := range _recipientAddr {
		_recipientAddrRule = append(_recipientAddrRule, _recipientAddrItem)
	}
	var _swapIDRule []interface{}
	for _, _swapIDItem := range _swapID {
		_swapIDRule = append(_swapIDRule, _swapIDItem)
	}

	logs, sub, err := _ERC20Swap.contract.WatchLogs(opts, "Claimed", _msgSenderRule, _recipientAddrRule, _swapIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20SwapClaimed)
				if err := _ERC20Swap.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0x9f46b1606087bdf4183ec7dfdbe68e4ab9129a6a37901c16a7b320ae11a96018.
//
// Solidity: event Claimed(address indexed _msgSender, address indexed _recipientAddr, bytes32 indexed _swapID, bytes32 _randomNumberHash, bytes32 _randomNumber)
func (_ERC20Swap *ERC20SwapFilterer) ParseClaimed(log types.Log) (*ERC20SwapClaimed, error) {
	event := new(ERC20SwapClaimed)
	if err := _ERC20Swap.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20SwapHTLTIterator is returned from FilterHTLT and is used to iterate over the raw logs and unpacked data for HTLT events raised by the ERC20Swap contract.
type ERC20SwapHTLTIterator struct {
	Event *ERC20SwapHTLT // Event containing the contract specifics and raw log

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
func (it *ERC20SwapHTLTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20SwapHTLT)
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
		it.Event = new(ERC20SwapHTLT)
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
func (it *ERC20SwapHTLTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20SwapHTLTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20SwapHTLT represents a HTLT event raised by the ERC20Swap contract.
type ERC20SwapHTLT struct {
	MsgSender        common.Address
	LaRecipientAddr  common.Address
	SwapID           [32]byte
	RandomNumberHash [32]byte
	Timestamp        uint64
	ERC20TokenAddr   common.Address
	LRC20TokenAddr   common.Address
	RefundAddr       [20]byte
	ExpireHeight     *big.Int
	OutAmount        *big.Int
	LaAmount         *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterHTLT is a free log retrieval operation binding the contract event 0xa9d611af8ea77eff0d427e2ae8375a23513129a25c615378cff5594b03c36ce9.
//
// Solidity: event HTLT(address indexed _msgSender, address indexed _laRecipientAddr, bytes32 indexed _swapID, bytes32 _randomNumberHash, uint64 _timestamp, address _ERC20TokenAddr, address _LRC20TokenAddr, bytes20 _refundAddr, uint256 _expireHeight, uint256 _outAmount, uint256 _laAmount)
func (_ERC20Swap *ERC20SwapFilterer) FilterHTLT(opts *bind.FilterOpts, _msgSender []common.Address, _laRecipientAddr []common.Address, _swapID [][32]byte) (*ERC20SwapHTLTIterator, error) {

	var _msgSenderRule []interface{}
	for _, _msgSenderItem := range _msgSender {
		_msgSenderRule = append(_msgSenderRule, _msgSenderItem)
	}
	var _laRecipientAddrRule []interface{}
	for _, _laRecipientAddrItem := range _laRecipientAddr {
		_laRecipientAddrRule = append(_laRecipientAddrRule, _laRecipientAddrItem)
	}
	var _swapIDRule []interface{}
	for _, _swapIDItem := range _swapID {
		_swapIDRule = append(_swapIDRule, _swapIDItem)
	}

	logs, sub, err := _ERC20Swap.contract.FilterLogs(opts, "HTLT", _msgSenderRule, _laRecipientAddrRule, _swapIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20SwapHTLTIterator{contract: _ERC20Swap.contract, event: "HTLT", logs: logs, sub: sub}, nil
}

// WatchHTLT is a free log subscription operation binding the contract event 0xa9d611af8ea77eff0d427e2ae8375a23513129a25c615378cff5594b03c36ce9.
//
// Solidity: event HTLT(address indexed _msgSender, address indexed _laRecipientAddr, bytes32 indexed _swapID, bytes32 _randomNumberHash, uint64 _timestamp, address _ERC20TokenAddr, address _LRC20TokenAddr, bytes20 _refundAddr, uint256 _expireHeight, uint256 _outAmount, uint256 _laAmount)
func (_ERC20Swap *ERC20SwapFilterer) WatchHTLT(opts *bind.WatchOpts, sink chan<- *ERC20SwapHTLT, _msgSender []common.Address, _laRecipientAddr []common.Address, _swapID [][32]byte) (event.Subscription, error) {

	var _msgSenderRule []interface{}
	for _, _msgSenderItem := range _msgSender {
		_msgSenderRule = append(_msgSenderRule, _msgSenderItem)
	}
	var _laRecipientAddrRule []interface{}
	for _, _laRecipientAddrItem := range _laRecipientAddr {
		_laRecipientAddrRule = append(_laRecipientAddrRule, _laRecipientAddrItem)
	}
	var _swapIDRule []interface{}
	for _, _swapIDItem := range _swapID {
		_swapIDRule = append(_swapIDRule, _swapIDItem)
	}

	logs, sub, err := _ERC20Swap.contract.WatchLogs(opts, "HTLT", _msgSenderRule, _laRecipientAddrRule, _swapIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20SwapHTLT)
				if err := _ERC20Swap.contract.UnpackLog(event, "HTLT", log); err != nil {
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

// ParseHTLT is a log parse operation binding the contract event 0xa9d611af8ea77eff0d427e2ae8375a23513129a25c615378cff5594b03c36ce9.
//
// Solidity: event HTLT(address indexed _msgSender, address indexed _laRecipientAddr, bytes32 indexed _swapID, bytes32 _randomNumberHash, uint64 _timestamp, address _ERC20TokenAddr, address _LRC20TokenAddr, bytes20 _refundAddr, uint256 _expireHeight, uint256 _outAmount, uint256 _laAmount)
func (_ERC20Swap *ERC20SwapFilterer) ParseHTLT(log types.Log) (*ERC20SwapHTLT, error) {
	event := new(ERC20SwapHTLT)
	if err := _ERC20Swap.contract.UnpackLog(event, "HTLT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20SwapRefundedIterator is returned from FilterRefunded and is used to iterate over the raw logs and unpacked data for Refunded events raised by the ERC20Swap contract.
type ERC20SwapRefundedIterator struct {
	Event *ERC20SwapRefunded // Event containing the contract specifics and raw log

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
func (it *ERC20SwapRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20SwapRefunded)
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
		it.Event = new(ERC20SwapRefunded)
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
func (it *ERC20SwapRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20SwapRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20SwapRefunded represents a Refunded event raised by the ERC20Swap contract.
type ERC20SwapRefunded struct {
	MsgSender        common.Address
	RecipientAddr    common.Address
	SwapID           [32]byte
	RandomNumberHash [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRefunded is a free log retrieval operation binding the contract event 0x04eb8ae268f23cfe2f9d72fa12367b104af16959f6a93530a4cc0f50688124f9.
//
// Solidity: event Refunded(address indexed _msgSender, address indexed _recipientAddr, bytes32 indexed _swapID, bytes32 _randomNumberHash)
func (_ERC20Swap *ERC20SwapFilterer) FilterRefunded(opts *bind.FilterOpts, _msgSender []common.Address, _recipientAddr []common.Address, _swapID [][32]byte) (*ERC20SwapRefundedIterator, error) {

	var _msgSenderRule []interface{}
	for _, _msgSenderItem := range _msgSender {
		_msgSenderRule = append(_msgSenderRule, _msgSenderItem)
	}
	var _recipientAddrRule []interface{}
	for _, _recipientAddrItem := range _recipientAddr {
		_recipientAddrRule = append(_recipientAddrRule, _recipientAddrItem)
	}
	var _swapIDRule []interface{}
	for _, _swapIDItem := range _swapID {
		_swapIDRule = append(_swapIDRule, _swapIDItem)
	}

	logs, sub, err := _ERC20Swap.contract.FilterLogs(opts, "Refunded", _msgSenderRule, _recipientAddrRule, _swapIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20SwapRefundedIterator{contract: _ERC20Swap.contract, event: "Refunded", logs: logs, sub: sub}, nil
}

// WatchRefunded is a free log subscription operation binding the contract event 0x04eb8ae268f23cfe2f9d72fa12367b104af16959f6a93530a4cc0f50688124f9.
//
// Solidity: event Refunded(address indexed _msgSender, address indexed _recipientAddr, bytes32 indexed _swapID, bytes32 _randomNumberHash)
func (_ERC20Swap *ERC20SwapFilterer) WatchRefunded(opts *bind.WatchOpts, sink chan<- *ERC20SwapRefunded, _msgSender []common.Address, _recipientAddr []common.Address, _swapID [][32]byte) (event.Subscription, error) {

	var _msgSenderRule []interface{}
	for _, _msgSenderItem := range _msgSender {
		_msgSenderRule = append(_msgSenderRule, _msgSenderItem)
	}
	var _recipientAddrRule []interface{}
	for _, _recipientAddrItem := range _recipientAddr {
		_recipientAddrRule = append(_recipientAddrRule, _recipientAddrItem)
	}
	var _swapIDRule []interface{}
	for _, _swapIDItem := range _swapID {
		_swapIDRule = append(_swapIDRule, _swapIDItem)
	}

	logs, sub, err := _ERC20Swap.contract.WatchLogs(opts, "Refunded", _msgSenderRule, _recipientAddrRule, _swapIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20SwapRefunded)
				if err := _ERC20Swap.contract.UnpackLog(event, "Refunded", log); err != nil {
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

// ParseRefunded is a log parse operation binding the contract event 0x04eb8ae268f23cfe2f9d72fa12367b104af16959f6a93530a4cc0f50688124f9.
//
// Solidity: event Refunded(address indexed _msgSender, address indexed _recipientAddr, bytes32 indexed _swapID, bytes32 _randomNumberHash)
func (_ERC20Swap *ERC20SwapFilterer) ParseRefunded(log types.Log) (*ERC20SwapRefunded, error) {
	event := new(ERC20SwapRefunded)
	if err := _ERC20Swap.contract.UnpackLog(event, "Refunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20SwapSpendedIterator is returned from FilterSpended and is used to iterate over the raw logs and unpacked data for Spended events raised by the ERC20Swap contract.
type ERC20SwapSpendedIterator struct {
	Event *ERC20SwapSpended // Event containing the contract specifics and raw log

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
func (it *ERC20SwapSpendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20SwapSpended)
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
		it.Event = new(ERC20SwapSpended)
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
func (it *ERC20SwapSpendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20SwapSpendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20SwapSpended represents a Spended event raised by the ERC20Swap contract.
type ERC20SwapSpended struct {
	MsgSender     common.Address
	RecipientAddr common.Address
	OutAmount     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSpended is a free log retrieval operation binding the contract event 0x606e9255bd8aa582845da439bf33a9a65993f410151cdbddff6afcda94b75d56.
//
// Solidity: event Spended(address indexed _msgSender, address indexed _recipientAddr, uint256 _outAmount)
func (_ERC20Swap *ERC20SwapFilterer) FilterSpended(opts *bind.FilterOpts, _msgSender []common.Address, _recipientAddr []common.Address) (*ERC20SwapSpendedIterator, error) {

	var _msgSenderRule []interface{}
	for _, _msgSenderItem := range _msgSender {
		_msgSenderRule = append(_msgSenderRule, _msgSenderItem)
	}
	var _recipientAddrRule []interface{}
	for _, _recipientAddrItem := range _recipientAddr {
		_recipientAddrRule = append(_recipientAddrRule, _recipientAddrItem)
	}

	logs, sub, err := _ERC20Swap.contract.FilterLogs(opts, "Spended", _msgSenderRule, _recipientAddrRule)
	if err != nil {
		return nil, err
	}
	return &ERC20SwapSpendedIterator{contract: _ERC20Swap.contract, event: "Spended", logs: logs, sub: sub}, nil
}

// WatchSpended is a free log subscription operation binding the contract event 0x606e9255bd8aa582845da439bf33a9a65993f410151cdbddff6afcda94b75d56.
//
// Solidity: event Spended(address indexed _msgSender, address indexed _recipientAddr, uint256 _outAmount)
func (_ERC20Swap *ERC20SwapFilterer) WatchSpended(opts *bind.WatchOpts, sink chan<- *ERC20SwapSpended, _msgSender []common.Address, _recipientAddr []common.Address) (event.Subscription, error) {

	var _msgSenderRule []interface{}
	for _, _msgSenderItem := range _msgSender {
		_msgSenderRule = append(_msgSenderRule, _msgSenderItem)
	}
	var _recipientAddrRule []interface{}
	for _, _recipientAddrItem := range _recipientAddr {
		_recipientAddrRule = append(_recipientAddrRule, _recipientAddrItem)
	}

	logs, sub, err := _ERC20Swap.contract.WatchLogs(opts, "Spended", _msgSenderRule, _recipientAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20SwapSpended)
				if err := _ERC20Swap.contract.UnpackLog(event, "Spended", log); err != nil {
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

// ParseSpended is a log parse operation binding the contract event 0x606e9255bd8aa582845da439bf33a9a65993f410151cdbddff6afcda94b75d56.
//
// Solidity: event Spended(address indexed _msgSender, address indexed _recipientAddr, uint256 _outAmount)
func (_ERC20Swap *ERC20SwapFilterer) ParseSpended(log types.Log) (*ERC20SwapSpended, error) {
	event := new(ERC20SwapSpended)
	if err := _ERC20Swap.contract.UnpackLog(event, "Spended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20SwapSwapCompletedIterator is returned from FilterSwapCompleted and is used to iterate over the raw logs and unpacked data for SwapCompleted events raised by the ERC20Swap contract.
type ERC20SwapSwapCompletedIterator struct {
	Event *ERC20SwapSwapCompleted // Event containing the contract specifics and raw log

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
func (it *ERC20SwapSwapCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20SwapSwapCompleted)
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
		it.Event = new(ERC20SwapSwapCompleted)
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
func (it *ERC20SwapSwapCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20SwapSwapCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20SwapSwapCompleted represents a SwapCompleted event raised by the ERC20Swap contract.
type ERC20SwapSwapCompleted struct {
	SwapID [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSwapCompleted is a free log retrieval operation binding the contract event 0xcbd087f9a8a65f06ed193f7147c5f9da5d4c1b09a2de92b7f4196854a3e0e710.
//
// Solidity: event SwapCompleted(bytes32 indexed _swapID)
func (_ERC20Swap *ERC20SwapFilterer) FilterSwapCompleted(opts *bind.FilterOpts, _swapID [][32]byte) (*ERC20SwapSwapCompletedIterator, error) {

	var _swapIDRule []interface{}
	for _, _swapIDItem := range _swapID {
		_swapIDRule = append(_swapIDRule, _swapIDItem)
	}

	logs, sub, err := _ERC20Swap.contract.FilterLogs(opts, "SwapCompleted", _swapIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20SwapSwapCompletedIterator{contract: _ERC20Swap.contract, event: "SwapCompleted", logs: logs, sub: sub}, nil
}

// WatchSwapCompleted is a free log subscription operation binding the contract event 0xcbd087f9a8a65f06ed193f7147c5f9da5d4c1b09a2de92b7f4196854a3e0e710.
//
// Solidity: event SwapCompleted(bytes32 indexed _swapID)
func (_ERC20Swap *ERC20SwapFilterer) WatchSwapCompleted(opts *bind.WatchOpts, sink chan<- *ERC20SwapSwapCompleted, _swapID [][32]byte) (event.Subscription, error) {

	var _swapIDRule []interface{}
	for _, _swapIDItem := range _swapID {
		_swapIDRule = append(_swapIDRule, _swapIDItem)
	}

	logs, sub, err := _ERC20Swap.contract.WatchLogs(opts, "SwapCompleted", _swapIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20SwapSwapCompleted)
				if err := _ERC20Swap.contract.UnpackLog(event, "SwapCompleted", log); err != nil {
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

// ParseSwapCompleted is a log parse operation binding the contract event 0xcbd087f9a8a65f06ed193f7147c5f9da5d4c1b09a2de92b7f4196854a3e0e710.
//
// Solidity: event SwapCompleted(bytes32 indexed _swapID)
func (_ERC20Swap *ERC20SwapFilterer) ParseSwapCompleted(log types.Log) (*ERC20SwapSwapCompleted, error) {
	event := new(ERC20SwapSwapCompleted)
	if err := _ERC20Swap.contract.UnpackLog(event, "SwapCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20SwapSwapFailedIterator is returned from FilterSwapFailed and is used to iterate over the raw logs and unpacked data for SwapFailed events raised by the ERC20Swap contract.
type ERC20SwapSwapFailedIterator struct {
	Event *ERC20SwapSwapFailed // Event containing the contract specifics and raw log

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
func (it *ERC20SwapSwapFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20SwapSwapFailed)
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
		it.Event = new(ERC20SwapSwapFailed)
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
func (it *ERC20SwapSwapFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20SwapSwapFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20SwapSwapFailed represents a SwapFailed event raised by the ERC20Swap contract.
type ERC20SwapSwapFailed struct {
	SwapID [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSwapFailed is a free log retrieval operation binding the contract event 0xeb1a0bf05090d67053fc1b869a236ac67c310e155140a8e4e9c6bfebbfd17188.
//
// Solidity: event SwapFailed(bytes32 indexed _swapID)
func (_ERC20Swap *ERC20SwapFilterer) FilterSwapFailed(opts *bind.FilterOpts, _swapID [][32]byte) (*ERC20SwapSwapFailedIterator, error) {

	var _swapIDRule []interface{}
	for _, _swapIDItem := range _swapID {
		_swapIDRule = append(_swapIDRule, _swapIDItem)
	}

	logs, sub, err := _ERC20Swap.contract.FilterLogs(opts, "SwapFailed", _swapIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20SwapSwapFailedIterator{contract: _ERC20Swap.contract, event: "SwapFailed", logs: logs, sub: sub}, nil
}

// WatchSwapFailed is a free log subscription operation binding the contract event 0xeb1a0bf05090d67053fc1b869a236ac67c310e155140a8e4e9c6bfebbfd17188.
//
// Solidity: event SwapFailed(bytes32 indexed _swapID)
func (_ERC20Swap *ERC20SwapFilterer) WatchSwapFailed(opts *bind.WatchOpts, sink chan<- *ERC20SwapSwapFailed, _swapID [][32]byte) (event.Subscription, error) {

	var _swapIDRule []interface{}
	for _, _swapIDItem := range _swapID {
		_swapIDRule = append(_swapIDRule, _swapIDItem)
	}

	logs, sub, err := _ERC20Swap.contract.WatchLogs(opts, "SwapFailed", _swapIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20SwapSwapFailed)
				if err := _ERC20Swap.contract.UnpackLog(event, "SwapFailed", log); err != nil {
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

// ParseSwapFailed is a log parse operation binding the contract event 0xeb1a0bf05090d67053fc1b869a236ac67c310e155140a8e4e9c6bfebbfd17188.
//
// Solidity: event SwapFailed(bytes32 indexed _swapID)
func (_ERC20Swap *ERC20SwapFilterer) ParseSwapFailed(log types.Log) (*ERC20SwapSwapFailed, error) {
	event := new(ERC20SwapSwapFailed)
	if err := _ERC20Swap.contract.UnpackLog(event, "SwapFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20SwapSwapOpenIterator is returned from FilterSwapOpen and is used to iterate over the raw logs and unpacked data for SwapOpen events raised by the ERC20Swap contract.
type ERC20SwapSwapOpenIterator struct {
	Event *ERC20SwapSwapOpen // Event containing the contract specifics and raw log

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
func (it *ERC20SwapSwapOpenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20SwapSwapOpen)
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
		it.Event = new(ERC20SwapSwapOpen)
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
func (it *ERC20SwapSwapOpenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20SwapSwapOpenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20SwapSwapOpen represents a SwapOpen event raised by the ERC20Swap contract.
type ERC20SwapSwapOpen struct {
	SwapID [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSwapOpen is a free log retrieval operation binding the contract event 0x2da8cdd88ec56e298672640ec11e2f14f737c9414e8cfa5b681ee74636f67f74.
//
// Solidity: event SwapOpen(bytes32 indexed _swapID)
func (_ERC20Swap *ERC20SwapFilterer) FilterSwapOpen(opts *bind.FilterOpts, _swapID [][32]byte) (*ERC20SwapSwapOpenIterator, error) {

	var _swapIDRule []interface{}
	for _, _swapIDItem := range _swapID {
		_swapIDRule = append(_swapIDRule, _swapIDItem)
	}

	logs, sub, err := _ERC20Swap.contract.FilterLogs(opts, "SwapOpen", _swapIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20SwapSwapOpenIterator{contract: _ERC20Swap.contract, event: "SwapOpen", logs: logs, sub: sub}, nil
}

// WatchSwapOpen is a free log subscription operation binding the contract event 0x2da8cdd88ec56e298672640ec11e2f14f737c9414e8cfa5b681ee74636f67f74.
//
// Solidity: event SwapOpen(bytes32 indexed _swapID)
func (_ERC20Swap *ERC20SwapFilterer) WatchSwapOpen(opts *bind.WatchOpts, sink chan<- *ERC20SwapSwapOpen, _swapID [][32]byte) (event.Subscription, error) {

	var _swapIDRule []interface{}
	for _, _swapIDItem := range _swapID {
		_swapIDRule = append(_swapIDRule, _swapIDItem)
	}

	logs, sub, err := _ERC20Swap.contract.WatchLogs(opts, "SwapOpen", _swapIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20SwapSwapOpen)
				if err := _ERC20Swap.contract.UnpackLog(event, "SwapOpen", log); err != nil {
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

// ParseSwapOpen is a log parse operation binding the contract event 0x2da8cdd88ec56e298672640ec11e2f14f737c9414e8cfa5b681ee74636f67f74.
//
// Solidity: event SwapOpen(bytes32 indexed _swapID)
func (_ERC20Swap *ERC20SwapFilterer) ParseSwapOpen(log types.Log) (*ERC20SwapSwapOpen, error) {
	event := new(ERC20SwapSwapOpen)
	if err := _ERC20Swap.contract.UnpackLog(event, "SwapOpen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20SwapParamChangeIterator is returned from FilterParamChange and is used to iterate over the raw logs and unpacked data for ParamChange events raised by the ERC20Swap contract.
type ERC20SwapParamChangeIterator struct {
	Event *ERC20SwapParamChange // Event containing the contract specifics and raw log

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
func (it *ERC20SwapParamChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20SwapParamChange)
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
		it.Event = new(ERC20SwapParamChange)
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
func (it *ERC20SwapParamChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20SwapParamChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20SwapParamChange represents a ParamChange event raised by the ERC20Swap contract.
type ERC20SwapParamChange struct {
	Key   string
	Value []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterParamChange is a free log retrieval operation binding the contract event 0x6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a.
//
// Solidity: event paramChange(string key, bytes value)
func (_ERC20Swap *ERC20SwapFilterer) FilterParamChange(opts *bind.FilterOpts) (*ERC20SwapParamChangeIterator, error) {

	logs, sub, err := _ERC20Swap.contract.FilterLogs(opts, "paramChange")
	if err != nil {
		return nil, err
	}
	return &ERC20SwapParamChangeIterator{contract: _ERC20Swap.contract, event: "paramChange", logs: logs, sub: sub}, nil
}

// WatchParamChange is a free log subscription operation binding the contract event 0x6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a.
//
// Solidity: event paramChange(string key, bytes value)
func (_ERC20Swap *ERC20SwapFilterer) WatchParamChange(opts *bind.WatchOpts, sink chan<- *ERC20SwapParamChange) (event.Subscription, error) {

	logs, sub, err := _ERC20Swap.contract.WatchLogs(opts, "paramChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20SwapParamChange)
				if err := _ERC20Swap.contract.UnpackLog(event, "paramChange", log); err != nil {
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

// ParseParamChange is a log parse operation binding the contract event 0x6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a.
//
// Solidity: event paramChange(string key, bytes value)
func (_ERC20Swap *ERC20SwapFilterer) ParseParamChange(log types.Log) (*ERC20SwapParamChange, error) {
	event := new(ERC20SwapParamChange)
	if err := _ERC20Swap.contract.UnpackLog(event, "paramChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20SwapRelayerRegisterIterator is returned from FilterRelayerRegister and is used to iterate over the raw logs and unpacked data for RelayerRegister events raised by the ERC20Swap contract.
type ERC20SwapRelayerRegisterIterator struct {
	Event *ERC20SwapRelayerRegister // Event containing the contract specifics and raw log

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
func (it *ERC20SwapRelayerRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20SwapRelayerRegister)
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
		it.Event = new(ERC20SwapRelayerRegister)
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
func (it *ERC20SwapRelayerRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20SwapRelayerRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20SwapRelayerRegister represents a RelayerRegister event raised by the ERC20Swap contract.
type ERC20SwapRelayerRegister struct {
	Relayer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRelayerRegister is a free log retrieval operation binding the contract event 0xdb33a09d38b59a8fa8b7d92a1d82c8015e99f05f67ae9c9ae623157767959496.
//
// Solidity: event relayerRegister(address _relayer)
func (_ERC20Swap *ERC20SwapFilterer) FilterRelayerRegister(opts *bind.FilterOpts) (*ERC20SwapRelayerRegisterIterator, error) {

	logs, sub, err := _ERC20Swap.contract.FilterLogs(opts, "relayerRegister")
	if err != nil {
		return nil, err
	}
	return &ERC20SwapRelayerRegisterIterator{contract: _ERC20Swap.contract, event: "relayerRegister", logs: logs, sub: sub}, nil
}

// WatchRelayerRegister is a free log subscription operation binding the contract event 0xdb33a09d38b59a8fa8b7d92a1d82c8015e99f05f67ae9c9ae623157767959496.
//
// Solidity: event relayerRegister(address _relayer)
func (_ERC20Swap *ERC20SwapFilterer) WatchRelayerRegister(opts *bind.WatchOpts, sink chan<- *ERC20SwapRelayerRegister) (event.Subscription, error) {

	logs, sub, err := _ERC20Swap.contract.WatchLogs(opts, "relayerRegister")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20SwapRelayerRegister)
				if err := _ERC20Swap.contract.UnpackLog(event, "relayerRegister", log); err != nil {
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

// ParseRelayerRegister is a log parse operation binding the contract event 0xdb33a09d38b59a8fa8b7d92a1d82c8015e99f05f67ae9c9ae623157767959496.
//
// Solidity: event relayerRegister(address _relayer)
func (_ERC20Swap *ERC20SwapFilterer) ParseRelayerRegister(log types.Log) (*ERC20SwapRelayerRegister, error) {
	event := new(ERC20SwapRelayerRegister)
	if err := _ERC20Swap.contract.UnpackLog(event, "relayerRegister", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20SwapRelayerUnRegisterIterator is returned from FilterRelayerUnRegister and is used to iterate over the raw logs and unpacked data for RelayerUnRegister events raised by the ERC20Swap contract.
type ERC20SwapRelayerUnRegisterIterator struct {
	Event *ERC20SwapRelayerUnRegister // Event containing the contract specifics and raw log

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
func (it *ERC20SwapRelayerUnRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20SwapRelayerUnRegister)
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
		it.Event = new(ERC20SwapRelayerUnRegister)
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
func (it *ERC20SwapRelayerUnRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20SwapRelayerUnRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20SwapRelayerUnRegister represents a RelayerUnRegister event raised by the ERC20Swap contract.
type ERC20SwapRelayerUnRegister struct {
	Relayer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRelayerUnRegister is a free log retrieval operation binding the contract event 0xd17202129b83db7880d6b9f25df81c58ad46f7e0e2c92236b1aa10663a487667.
//
// Solidity: event relayerUnRegister(address _relayer)
func (_ERC20Swap *ERC20SwapFilterer) FilterRelayerUnRegister(opts *bind.FilterOpts) (*ERC20SwapRelayerUnRegisterIterator, error) {

	logs, sub, err := _ERC20Swap.contract.FilterLogs(opts, "relayerUnRegister")
	if err != nil {
		return nil, err
	}
	return &ERC20SwapRelayerUnRegisterIterator{contract: _ERC20Swap.contract, event: "relayerUnRegister", logs: logs, sub: sub}, nil
}

// WatchRelayerUnRegister is a free log subscription operation binding the contract event 0xd17202129b83db7880d6b9f25df81c58ad46f7e0e2c92236b1aa10663a487667.
//
// Solidity: event relayerUnRegister(address _relayer)
func (_ERC20Swap *ERC20SwapFilterer) WatchRelayerUnRegister(opts *bind.WatchOpts, sink chan<- *ERC20SwapRelayerUnRegister) (event.Subscription, error) {

	logs, sub, err := _ERC20Swap.contract.WatchLogs(opts, "relayerUnRegister")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20SwapRelayerUnRegister)
				if err := _ERC20Swap.contract.UnpackLog(event, "relayerUnRegister", log); err != nil {
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

// ParseRelayerUnRegister is a log parse operation binding the contract event 0xd17202129b83db7880d6b9f25df81c58ad46f7e0e2c92236b1aa10663a487667.
//
// Solidity: event relayerUnRegister(address _relayer)
func (_ERC20Swap *ERC20SwapFilterer) ParseRelayerUnRegister(log types.Log) (*ERC20SwapRelayerUnRegister, error) {
	event := new(ERC20SwapRelayerUnRegister)
	if err := _ERC20Swap.contract.UnpackLog(event, "relayerUnRegister", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
