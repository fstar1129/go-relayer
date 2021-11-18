// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package manager

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

// ManagerABI is the input ABI used to generate the binding from.
const ManagerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_relayerAddr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_msgSender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_swapID\",\"type\":\"bytes32\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_senderAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_recipientAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_swapID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_randomNumberHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_timestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes20\",\"name\":\"_refundAddr\",\"type\":\"bytes20\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_expireHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_outAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_laAmount\",\"type\":\"uint256\"}],\"name\":\"CreateUnbind\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_msgSender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_swapID\",\"type\":\"bytes32\"}],\"name\":\"Created\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_senderAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_recipientAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_swapID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_randomNumberHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_timestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes20\",\"name\":\"_refundAddr\",\"type\":\"bytes20\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_expireHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_outAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_laAmount\",\"type\":\"uint256\"}],\"name\":\"HTLT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"laContractAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txhash\",\"type\":\"bytes32\"}],\"name\":\"Spended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"Swapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_senderAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_recipientAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_swapID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_randomNumberHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"_timestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes20\",\"name\":\"_refundAddr\",\"type\":\"bytes20\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_expireHeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_outAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_laAmount\",\"type\":\"uint256\"}],\"name\":\"Unbinded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"bep2Symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"failedReason\",\"type\":\"uint32\"}],\"name\":\"bindFailure\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MINIMAL_NUM_CLAIMS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYERHUB_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ethTx\",\"type\":\"bytes32\"}],\"name\":\"approveSwap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"bindPackageRequests\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"packageType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"laTokenSymbol\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"peggyAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"expireTime\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ethTx\",\"type\":\"bytes32\"}],\"name\":\"createSwapRequest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"filledETHTx\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_randomNumberHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_heightSpan\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipientAddr\",\"type\":\"address\"},{\"internalType\":\"bytes20\",\"name\":\"_refundAddr\",\"type\":\"bytes20\"},{\"internalType\":\"uint256\",\"name\":\"_outAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_laAmount\",\"type\":\"uint256\"}],\"name\":\"htlt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mirrorPendingRecord\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ethTx\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_erc20Addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_laTokenAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_toAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"spendBind\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ethTx\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_erc20Addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_laTokenAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_toAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"spendUnBind\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"swapETH2LA\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"swapMappingETH2LA\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"swapMappingLA2ETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"swapRelayersRequest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"swapRequest\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_randomNumberHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_heightSpan\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_laTokenAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipientAddr\",\"type\":\"address\"},{\"internalType\":\"bytes20\",\"name\":\"_refundAddr\",\"type\":\"bytes20\"},{\"internalType\":\"uint256\",\"name\":\"_outAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_laAmount\",\"type\":\"uint256\"}],\"name\":\"unBind\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// Manager is an auto generated Go binding around an Ethereum contract.
type Manager struct {
	ManagerCaller     // Read-only binding to the contract
	ManagerTransactor // Write-only binding to the contract
	ManagerFilterer   // Log filterer for contract events
}

// ManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ManagerSession struct {
	Contract     *Manager          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ManagerCallerSession struct {
	Contract *ManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ManagerTransactorSession struct {
	Contract     *ManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ManagerRaw struct {
	Contract *Manager // Generic contract binding to access the raw methods on
}

// ManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ManagerCallerRaw struct {
	Contract *ManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ManagerTransactorRaw struct {
	Contract *ManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewManager creates a new instance of Manager, bound to a specific deployed contract.
func NewManager(address common.Address, backend bind.ContractBackend) (*Manager, error) {
	contract, err := bindManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Manager{ManagerCaller: ManagerCaller{contract: contract}, ManagerTransactor: ManagerTransactor{contract: contract}, ManagerFilterer: ManagerFilterer{contract: contract}}, nil
}

// NewManagerCaller creates a new read-only instance of Manager, bound to a specific deployed contract.
func NewManagerCaller(address common.Address, caller bind.ContractCaller) (*ManagerCaller, error) {
	contract, err := bindManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ManagerCaller{contract: contract}, nil
}

// NewManagerTransactor creates a new write-only instance of Manager, bound to a specific deployed contract.
func NewManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ManagerTransactor, error) {
	contract, err := bindManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ManagerTransactor{contract: contract}, nil
}

// NewManagerFilterer creates a new log filterer instance of Manager, bound to a specific deployed contract.
func NewManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ManagerFilterer, error) {
	contract, err := bindManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ManagerFilterer{contract: contract}, nil
}

// bindManager binds a generic wrapper to an already deployed contract.
func bindManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Manager *ManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Manager.Contract.ManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Manager *ManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Manager.Contract.ManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Manager *ManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Manager.Contract.ManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Manager *ManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Manager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Manager *ManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Manager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Manager *ManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Manager.Contract.contract.Transact(opts, method, params...)
}

// MINIMALNUMCLAIMS is a free data retrieval call binding the contract method 0x95787f7f.
//
// Solidity: function MINIMAL_NUM_CLAIMS() view returns(uint8)
func (_Manager *ManagerCaller) MINIMALNUMCLAIMS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Manager.contract.Call(opts, &out, "MINIMAL_NUM_CLAIMS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MINIMALNUMCLAIMS is a free data retrieval call binding the contract method 0x95787f7f.
//
// Solidity: function MINIMAL_NUM_CLAIMS() view returns(uint8)
func (_Manager *ManagerSession) MINIMALNUMCLAIMS() (uint8, error) {
	return _Manager.Contract.MINIMALNUMCLAIMS(&_Manager.CallOpts)
}

// MINIMALNUMCLAIMS is a free data retrieval call binding the contract method 0x95787f7f.
//
// Solidity: function MINIMAL_NUM_CLAIMS() view returns(uint8)
func (_Manager *ManagerCallerSession) MINIMALNUMCLAIMS() (uint8, error) {
	return _Manager.Contract.MINIMALNUMCLAIMS(&_Manager.CallOpts)
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() view returns(address)
func (_Manager *ManagerCaller) RELAYERHUBCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Manager.contract.Call(opts, &out, "RELAYERHUB_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() view returns(address)
func (_Manager *ManagerSession) RELAYERHUBCONTRACTADDR() (common.Address, error) {
	return _Manager.Contract.RELAYERHUBCONTRACTADDR(&_Manager.CallOpts)
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() view returns(address)
func (_Manager *ManagerCallerSession) RELAYERHUBCONTRACTADDR() (common.Address, error) {
	return _Manager.Contract.RELAYERHUBCONTRACTADDR(&_Manager.CallOpts)
}

// BindPackageRequests is a free data retrieval call binding the contract method 0x9768431f.
//
// Solidity: function bindPackageRequests(string ) view returns(uint8 packageType, bytes32 laTokenSymbol, address contractAddr, uint256 totalSupply, uint256 peggyAmount, uint8 decimals, uint64 expireTime)
func (_Manager *ManagerCaller) BindPackageRequests(opts *bind.CallOpts, arg0 string) (struct {
	PackageType   uint8
	LaTokenSymbol [32]byte
	ContractAddr  common.Address
	TotalSupply   *big.Int
	PeggyAmount   *big.Int
	Decimals      uint8
	ExpireTime    uint64
}, error) {
	var out []interface{}
	err := _Manager.contract.Call(opts, &out, "bindPackageRequests", arg0)

	outstruct := new(struct {
		PackageType   uint8
		LaTokenSymbol [32]byte
		ContractAddr  common.Address
		TotalSupply   *big.Int
		PeggyAmount   *big.Int
		Decimals      uint8
		ExpireTime    uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PackageType = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.LaTokenSymbol = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.ContractAddr = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.TotalSupply = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.PeggyAmount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Decimals = *abi.ConvertType(out[5], new(uint8)).(*uint8)
	outstruct.ExpireTime = *abi.ConvertType(out[6], new(uint64)).(*uint64)

	return *outstruct, err

}

// BindPackageRequests is a free data retrieval call binding the contract method 0x9768431f.
//
// Solidity: function bindPackageRequests(string ) view returns(uint8 packageType, bytes32 laTokenSymbol, address contractAddr, uint256 totalSupply, uint256 peggyAmount, uint8 decimals, uint64 expireTime)
func (_Manager *ManagerSession) BindPackageRequests(arg0 string) (struct {
	PackageType   uint8
	LaTokenSymbol [32]byte
	ContractAddr  common.Address
	TotalSupply   *big.Int
	PeggyAmount   *big.Int
	Decimals      uint8
	ExpireTime    uint64
}, error) {
	return _Manager.Contract.BindPackageRequests(&_Manager.CallOpts, arg0)
}

// BindPackageRequests is a free data retrieval call binding the contract method 0x9768431f.
//
// Solidity: function bindPackageRequests(string ) view returns(uint8 packageType, bytes32 laTokenSymbol, address contractAddr, uint256 totalSupply, uint256 peggyAmount, uint8 decimals, uint64 expireTime)
func (_Manager *ManagerCallerSession) BindPackageRequests(arg0 string) (struct {
	PackageType   uint8
	LaTokenSymbol [32]byte
	ContractAddr  common.Address
	TotalSupply   *big.Int
	PeggyAmount   *big.Int
	Decimals      uint8
	ExpireTime    uint64
}, error) {
	return _Manager.Contract.BindPackageRequests(&_Manager.CallOpts, arg0)
}

// FilledETHTx is a free data retrieval call binding the contract method 0x4e2dc7f1.
//
// Solidity: function filledETHTx(bytes32 ) view returns(bool)
func (_Manager *ManagerCaller) FilledETHTx(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Manager.contract.Call(opts, &out, "filledETHTx", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FilledETHTx is a free data retrieval call binding the contract method 0x4e2dc7f1.
//
// Solidity: function filledETHTx(bytes32 ) view returns(bool)
func (_Manager *ManagerSession) FilledETHTx(arg0 [32]byte) (bool, error) {
	return _Manager.Contract.FilledETHTx(&_Manager.CallOpts, arg0)
}

// FilledETHTx is a free data retrieval call binding the contract method 0x4e2dc7f1.
//
// Solidity: function filledETHTx(bytes32 ) view returns(bool)
func (_Manager *ManagerCallerSession) FilledETHTx(arg0 [32]byte) (bool, error) {
	return _Manager.Contract.FilledETHTx(&_Manager.CallOpts, arg0)
}

// MirrorPendingRecord is a free data retrieval call binding the contract method 0x37e6ecda.
//
// Solidity: function mirrorPendingRecord(address ) view returns(bool)
func (_Manager *ManagerCaller) MirrorPendingRecord(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Manager.contract.Call(opts, &out, "mirrorPendingRecord", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MirrorPendingRecord is a free data retrieval call binding the contract method 0x37e6ecda.
//
// Solidity: function mirrorPendingRecord(address ) view returns(bool)
func (_Manager *ManagerSession) MirrorPendingRecord(arg0 common.Address) (bool, error) {
	return _Manager.Contract.MirrorPendingRecord(&_Manager.CallOpts, arg0)
}

// MirrorPendingRecord is a free data retrieval call binding the contract method 0x37e6ecda.
//
// Solidity: function mirrorPendingRecord(address ) view returns(bool)
func (_Manager *ManagerCallerSession) MirrorPendingRecord(arg0 common.Address) (bool, error) {
	return _Manager.Contract.MirrorPendingRecord(&_Manager.CallOpts, arg0)
}

// SwapMappingETH2LA is a free data retrieval call binding the contract method 0xc17fca1b.
//
// Solidity: function swapMappingETH2LA(address ) view returns(address)
func (_Manager *ManagerCaller) SwapMappingETH2LA(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Manager.contract.Call(opts, &out, "swapMappingETH2LA", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwapMappingETH2LA is a free data retrieval call binding the contract method 0xc17fca1b.
//
// Solidity: function swapMappingETH2LA(address ) view returns(address)
func (_Manager *ManagerSession) SwapMappingETH2LA(arg0 common.Address) (common.Address, error) {
	return _Manager.Contract.SwapMappingETH2LA(&_Manager.CallOpts, arg0)
}

// SwapMappingETH2LA is a free data retrieval call binding the contract method 0xc17fca1b.
//
// Solidity: function swapMappingETH2LA(address ) view returns(address)
func (_Manager *ManagerCallerSession) SwapMappingETH2LA(arg0 common.Address) (common.Address, error) {
	return _Manager.Contract.SwapMappingETH2LA(&_Manager.CallOpts, arg0)
}

// SwapMappingLA2ETH is a free data retrieval call binding the contract method 0xef1dcab4.
//
// Solidity: function swapMappingLA2ETH(address ) view returns(address)
func (_Manager *ManagerCaller) SwapMappingLA2ETH(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Manager.contract.Call(opts, &out, "swapMappingLA2ETH", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwapMappingLA2ETH is a free data retrieval call binding the contract method 0xef1dcab4.
//
// Solidity: function swapMappingLA2ETH(address ) view returns(address)
func (_Manager *ManagerSession) SwapMappingLA2ETH(arg0 common.Address) (common.Address, error) {
	return _Manager.Contract.SwapMappingLA2ETH(&_Manager.CallOpts, arg0)
}

// SwapMappingLA2ETH is a free data retrieval call binding the contract method 0xef1dcab4.
//
// Solidity: function swapMappingLA2ETH(address ) view returns(address)
func (_Manager *ManagerCallerSession) SwapMappingLA2ETH(arg0 common.Address) (common.Address, error) {
	return _Manager.Contract.SwapMappingLA2ETH(&_Manager.CallOpts, arg0)
}

// SwapRelayersRequest is a free data retrieval call binding the contract method 0xa769df13.
//
// Solidity: function swapRelayersRequest(bytes32 , address ) view returns(bool)
func (_Manager *ManagerCaller) SwapRelayersRequest(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Manager.contract.Call(opts, &out, "swapRelayersRequest", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SwapRelayersRequest is a free data retrieval call binding the contract method 0xa769df13.
//
// Solidity: function swapRelayersRequest(bytes32 , address ) view returns(bool)
func (_Manager *ManagerSession) SwapRelayersRequest(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _Manager.Contract.SwapRelayersRequest(&_Manager.CallOpts, arg0, arg1)
}

// SwapRelayersRequest is a free data retrieval call binding the contract method 0xa769df13.
//
// Solidity: function swapRelayersRequest(bytes32 , address ) view returns(bool)
func (_Manager *ManagerCallerSession) SwapRelayersRequest(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _Manager.Contract.SwapRelayersRequest(&_Manager.CallOpts, arg0, arg1)
}

// SwapRequest is a free data retrieval call binding the contract method 0x85c2b1fd.
//
// Solidity: function swapRequest(bytes32 ) view returns(uint8)
func (_Manager *ManagerCaller) SwapRequest(opts *bind.CallOpts, arg0 [32]byte) (uint8, error) {
	var out []interface{}
	err := _Manager.contract.Call(opts, &out, "swapRequest", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SwapRequest is a free data retrieval call binding the contract method 0x85c2b1fd.
//
// Solidity: function swapRequest(bytes32 ) view returns(uint8)
func (_Manager *ManagerSession) SwapRequest(arg0 [32]byte) (uint8, error) {
	return _Manager.Contract.SwapRequest(&_Manager.CallOpts, arg0)
}

// SwapRequest is a free data retrieval call binding the contract method 0x85c2b1fd.
//
// Solidity: function swapRequest(bytes32 ) view returns(uint8)
func (_Manager *ManagerCallerSession) SwapRequest(arg0 [32]byte) (uint8, error) {
	return _Manager.Contract.SwapRequest(&_Manager.CallOpts, arg0)
}

// ApproveSwap is a paid mutator transaction binding the contract method 0xe35468cb.
//
// Solidity: function approveSwap(bytes32 _ethTx) payable returns(bool)
func (_Manager *ManagerTransactor) ApproveSwap(opts *bind.TransactOpts, _ethTx [32]byte) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "approveSwap", _ethTx)
}

// ApproveSwap is a paid mutator transaction binding the contract method 0xe35468cb.
//
// Solidity: function approveSwap(bytes32 _ethTx) payable returns(bool)
func (_Manager *ManagerSession) ApproveSwap(_ethTx [32]byte) (*types.Transaction, error) {
	return _Manager.Contract.ApproveSwap(&_Manager.TransactOpts, _ethTx)
}

// ApproveSwap is a paid mutator transaction binding the contract method 0xe35468cb.
//
// Solidity: function approveSwap(bytes32 _ethTx) payable returns(bool)
func (_Manager *ManagerTransactorSession) ApproveSwap(_ethTx [32]byte) (*types.Transaction, error) {
	return _Manager.Contract.ApproveSwap(&_Manager.TransactOpts, _ethTx)
}

// CreateSwapRequest is a paid mutator transaction binding the contract method 0xa1da1095.
//
// Solidity: function createSwapRequest(bytes32 _ethTx) payable returns(bool)
func (_Manager *ManagerTransactor) CreateSwapRequest(opts *bind.TransactOpts, _ethTx [32]byte) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "createSwapRequest", _ethTx)
}

// CreateSwapRequest is a paid mutator transaction binding the contract method 0xa1da1095.
//
// Solidity: function createSwapRequest(bytes32 _ethTx) payable returns(bool)
func (_Manager *ManagerSession) CreateSwapRequest(_ethTx [32]byte) (*types.Transaction, error) {
	return _Manager.Contract.CreateSwapRequest(&_Manager.TransactOpts, _ethTx)
}

// CreateSwapRequest is a paid mutator transaction binding the contract method 0xa1da1095.
//
// Solidity: function createSwapRequest(bytes32 _ethTx) payable returns(bool)
func (_Manager *ManagerTransactorSession) CreateSwapRequest(_ethTx [32]byte) (*types.Transaction, error) {
	return _Manager.Contract.CreateSwapRequest(&_Manager.TransactOpts, _ethTx)
}

// Htlt is a paid mutator transaction binding the contract method 0xc4711430.
//
// Solidity: function htlt(bytes32 _randomNumberHash, uint64 _timestamp, uint256 _heightSpan, address _recipientAddr, bytes20 _refundAddr, uint256 _outAmount, uint256 _laAmount) returns(bool)
func (_Manager *ManagerTransactor) Htlt(opts *bind.TransactOpts, _randomNumberHash [32]byte, _timestamp uint64, _heightSpan *big.Int, _recipientAddr common.Address, _refundAddr [20]byte, _outAmount *big.Int, _laAmount *big.Int) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "htlt", _randomNumberHash, _timestamp, _heightSpan, _recipientAddr, _refundAddr, _outAmount, _laAmount)
}

// Htlt is a paid mutator transaction binding the contract method 0xc4711430.
//
// Solidity: function htlt(bytes32 _randomNumberHash, uint64 _timestamp, uint256 _heightSpan, address _recipientAddr, bytes20 _refundAddr, uint256 _outAmount, uint256 _laAmount) returns(bool)
func (_Manager *ManagerSession) Htlt(_randomNumberHash [32]byte, _timestamp uint64, _heightSpan *big.Int, _recipientAddr common.Address, _refundAddr [20]byte, _outAmount *big.Int, _laAmount *big.Int) (*types.Transaction, error) {
	return _Manager.Contract.Htlt(&_Manager.TransactOpts, _randomNumberHash, _timestamp, _heightSpan, _recipientAddr, _refundAddr, _outAmount, _laAmount)
}

// Htlt is a paid mutator transaction binding the contract method 0xc4711430.
//
// Solidity: function htlt(bytes32 _randomNumberHash, uint64 _timestamp, uint256 _heightSpan, address _recipientAddr, bytes20 _refundAddr, uint256 _outAmount, uint256 _laAmount) returns(bool)
func (_Manager *ManagerTransactorSession) Htlt(_randomNumberHash [32]byte, _timestamp uint64, _heightSpan *big.Int, _recipientAddr common.Address, _refundAddr [20]byte, _outAmount *big.Int, _laAmount *big.Int) (*types.Transaction, error) {
	return _Manager.Contract.Htlt(&_Manager.TransactOpts, _randomNumberHash, _timestamp, _heightSpan, _recipientAddr, _refundAddr, _outAmount, _laAmount)
}

// SpendBind is a paid mutator transaction binding the contract method 0x8f831b32.
//
// Solidity: function spendBind(bytes32 _ethTx, address _erc20Addr, address _laTokenAddr, address _toAddress, uint256 _amount) payable returns(bool)
func (_Manager *ManagerTransactor) SpendBind(opts *bind.TransactOpts, _ethTx [32]byte, _erc20Addr common.Address, _laTokenAddr common.Address, _toAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "spendBind", _ethTx, _erc20Addr, _laTokenAddr, _toAddress, _amount)
}

// SpendBind is a paid mutator transaction binding the contract method 0x8f831b32.
//
// Solidity: function spendBind(bytes32 _ethTx, address _erc20Addr, address _laTokenAddr, address _toAddress, uint256 _amount) payable returns(bool)
func (_Manager *ManagerSession) SpendBind(_ethTx [32]byte, _erc20Addr common.Address, _laTokenAddr common.Address, _toAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Manager.Contract.SpendBind(&_Manager.TransactOpts, _ethTx, _erc20Addr, _laTokenAddr, _toAddress, _amount)
}

// SpendBind is a paid mutator transaction binding the contract method 0x8f831b32.
//
// Solidity: function spendBind(bytes32 _ethTx, address _erc20Addr, address _laTokenAddr, address _toAddress, uint256 _amount) payable returns(bool)
func (_Manager *ManagerTransactorSession) SpendBind(_ethTx [32]byte, _erc20Addr common.Address, _laTokenAddr common.Address, _toAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Manager.Contract.SpendBind(&_Manager.TransactOpts, _ethTx, _erc20Addr, _laTokenAddr, _toAddress, _amount)
}

// SpendUnBind is a paid mutator transaction binding the contract method 0xf8c8444d.
//
// Solidity: function spendUnBind(bytes32 _ethTx, address _erc20Addr, address _laTokenAddr, address _toAddress, uint256 _amount) payable returns(bool)
func (_Manager *ManagerTransactor) SpendUnBind(opts *bind.TransactOpts, _ethTx [32]byte, _erc20Addr common.Address, _laTokenAddr common.Address, _toAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "spendUnBind", _ethTx, _erc20Addr, _laTokenAddr, _toAddress, _amount)
}

// SpendUnBind is a paid mutator transaction binding the contract method 0xf8c8444d.
//
// Solidity: function spendUnBind(bytes32 _ethTx, address _erc20Addr, address _laTokenAddr, address _toAddress, uint256 _amount) payable returns(bool)
func (_Manager *ManagerSession) SpendUnBind(_ethTx [32]byte, _erc20Addr common.Address, _laTokenAddr common.Address, _toAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Manager.Contract.SpendUnBind(&_Manager.TransactOpts, _ethTx, _erc20Addr, _laTokenAddr, _toAddress, _amount)
}

// SpendUnBind is a paid mutator transaction binding the contract method 0xf8c8444d.
//
// Solidity: function spendUnBind(bytes32 _ethTx, address _erc20Addr, address _laTokenAddr, address _toAddress, uint256 _amount) payable returns(bool)
func (_Manager *ManagerTransactorSession) SpendUnBind(_ethTx [32]byte, _erc20Addr common.Address, _laTokenAddr common.Address, _toAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Manager.Contract.SpendUnBind(&_Manager.TransactOpts, _ethTx, _erc20Addr, _laTokenAddr, _toAddress, _amount)
}

// SwapETH2LA is a paid mutator transaction binding the contract method 0xd9a93ce9.
//
// Solidity: function swapETH2LA(address erc20Addr, uint256 amount) payable returns(bool)
func (_Manager *ManagerTransactor) SwapETH2LA(opts *bind.TransactOpts, erc20Addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "swapETH2LA", erc20Addr, amount)
}

// SwapETH2LA is a paid mutator transaction binding the contract method 0xd9a93ce9.
//
// Solidity: function swapETH2LA(address erc20Addr, uint256 amount) payable returns(bool)
func (_Manager *ManagerSession) SwapETH2LA(erc20Addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Manager.Contract.SwapETH2LA(&_Manager.TransactOpts, erc20Addr, amount)
}

// SwapETH2LA is a paid mutator transaction binding the contract method 0xd9a93ce9.
//
// Solidity: function swapETH2LA(address erc20Addr, uint256 amount) payable returns(bool)
func (_Manager *ManagerTransactorSession) SwapETH2LA(erc20Addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Manager.Contract.SwapETH2LA(&_Manager.TransactOpts, erc20Addr, amount)
}

// UnBind is a paid mutator transaction binding the contract method 0xbc174c8c.
//
// Solidity: function unBind(bytes32 _randomNumberHash, uint64 _timestamp, uint256 _heightSpan, address _laTokenAddr, address _recipientAddr, bytes20 _refundAddr, uint256 _outAmount, uint256 _laAmount) payable returns(bool)
func (_Manager *ManagerTransactor) UnBind(opts *bind.TransactOpts, _randomNumberHash [32]byte, _timestamp uint64, _heightSpan *big.Int, _laTokenAddr common.Address, _recipientAddr common.Address, _refundAddr [20]byte, _outAmount *big.Int, _laAmount *big.Int) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "unBind", _randomNumberHash, _timestamp, _heightSpan, _laTokenAddr, _recipientAddr, _refundAddr, _outAmount, _laAmount)
}

// UnBind is a paid mutator transaction binding the contract method 0xbc174c8c.
//
// Solidity: function unBind(bytes32 _randomNumberHash, uint64 _timestamp, uint256 _heightSpan, address _laTokenAddr, address _recipientAddr, bytes20 _refundAddr, uint256 _outAmount, uint256 _laAmount) payable returns(bool)
func (_Manager *ManagerSession) UnBind(_randomNumberHash [32]byte, _timestamp uint64, _heightSpan *big.Int, _laTokenAddr common.Address, _recipientAddr common.Address, _refundAddr [20]byte, _outAmount *big.Int, _laAmount *big.Int) (*types.Transaction, error) {
	return _Manager.Contract.UnBind(&_Manager.TransactOpts, _randomNumberHash, _timestamp, _heightSpan, _laTokenAddr, _recipientAddr, _refundAddr, _outAmount, _laAmount)
}

// UnBind is a paid mutator transaction binding the contract method 0xbc174c8c.
//
// Solidity: function unBind(bytes32 _randomNumberHash, uint64 _timestamp, uint256 _heightSpan, address _laTokenAddr, address _recipientAddr, bytes20 _refundAddr, uint256 _outAmount, uint256 _laAmount) payable returns(bool)
func (_Manager *ManagerTransactorSession) UnBind(_randomNumberHash [32]byte, _timestamp uint64, _heightSpan *big.Int, _laTokenAddr common.Address, _recipientAddr common.Address, _refundAddr [20]byte, _outAmount *big.Int, _laAmount *big.Int) (*types.Transaction, error) {
	return _Manager.Contract.UnBind(&_Manager.TransactOpts, _randomNumberHash, _timestamp, _heightSpan, _laTokenAddr, _recipientAddr, _refundAddr, _outAmount, _laAmount)
}

// ManagerClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the Manager contract.
type ManagerClaimedIterator struct {
	Event *ManagerClaimed // Event containing the contract specifics and raw log

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
func (it *ManagerClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManagerClaimed)
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
		it.Event = new(ManagerClaimed)
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
func (it *ManagerClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManagerClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManagerClaimed represents a Claimed event raised by the Manager contract.
type ManagerClaimed struct {
	MsgSender common.Address
	SwapID    [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0x91d697238e9aa9f3172d17522c9be529b94a892481554e1ea619369b5b12f39a.
//
// Solidity: event Claimed(address indexed _msgSender, bytes32 indexed _swapID)
func (_Manager *ManagerFilterer) FilterClaimed(opts *bind.FilterOpts, _msgSender []common.Address, _swapID [][32]byte) (*ManagerClaimedIterator, error) {

	var _msgSenderRule []interface{}
	for _, _msgSenderItem := range _msgSender {
		_msgSenderRule = append(_msgSenderRule, _msgSenderItem)
	}
	var _swapIDRule []interface{}
	for _, _swapIDItem := range _swapID {
		_swapIDRule = append(_swapIDRule, _swapIDItem)
	}

	logs, sub, err := _Manager.contract.FilterLogs(opts, "Claimed", _msgSenderRule, _swapIDRule)
	if err != nil {
		return nil, err
	}
	return &ManagerClaimedIterator{contract: _Manager.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0x91d697238e9aa9f3172d17522c9be529b94a892481554e1ea619369b5b12f39a.
//
// Solidity: event Claimed(address indexed _msgSender, bytes32 indexed _swapID)
func (_Manager *ManagerFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *ManagerClaimed, _msgSender []common.Address, _swapID [][32]byte) (event.Subscription, error) {

	var _msgSenderRule []interface{}
	for _, _msgSenderItem := range _msgSender {
		_msgSenderRule = append(_msgSenderRule, _msgSenderItem)
	}
	var _swapIDRule []interface{}
	for _, _swapIDItem := range _swapID {
		_swapIDRule = append(_swapIDRule, _swapIDItem)
	}

	logs, sub, err := _Manager.contract.WatchLogs(opts, "Claimed", _msgSenderRule, _swapIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManagerClaimed)
				if err := _Manager.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0x91d697238e9aa9f3172d17522c9be529b94a892481554e1ea619369b5b12f39a.
//
// Solidity: event Claimed(address indexed _msgSender, bytes32 indexed _swapID)
func (_Manager *ManagerFilterer) ParseClaimed(log types.Log) (*ManagerClaimed, error) {
	event := new(ManagerClaimed)
	if err := _Manager.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ManagerCreateUnbindIterator is returned from FilterCreateUnbind and is used to iterate over the raw logs and unpacked data for CreateUnbind events raised by the Manager contract.
type ManagerCreateUnbindIterator struct {
	Event *ManagerCreateUnbind // Event containing the contract specifics and raw log

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
func (it *ManagerCreateUnbindIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManagerCreateUnbind)
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
		it.Event = new(ManagerCreateUnbind)
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
func (it *ManagerCreateUnbindIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManagerCreateUnbindIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManagerCreateUnbind represents a CreateUnbind event raised by the Manager contract.
type ManagerCreateUnbind struct {
	SenderAddr       common.Address
	RecipientAddr    common.Address
	SwapID           [32]byte
	RandomNumberHash [32]byte
	Timestamp        uint64
	RefundAddr       [20]byte
	ExpireHeight     *big.Int
	OutAmount        *big.Int
	LaAmount         *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCreateUnbind is a free log retrieval operation binding the contract event 0xf3f48b99c0f6ef1f16092aed9ebd7d552e8b7647e0332990550d7626d7857ba5.
//
// Solidity: event CreateUnbind(address indexed _senderAddr, address indexed _recipientAddr, bytes32 indexed _swapID, bytes32 _randomNumberHash, uint64 _timestamp, bytes20 _refundAddr, uint256 _expireHeight, uint256 _outAmount, uint256 _laAmount)
func (_Manager *ManagerFilterer) FilterCreateUnbind(opts *bind.FilterOpts, _senderAddr []common.Address, _recipientAddr []common.Address, _swapID [][32]byte) (*ManagerCreateUnbindIterator, error) {

	var _senderAddrRule []interface{}
	for _, _senderAddrItem := range _senderAddr {
		_senderAddrRule = append(_senderAddrRule, _senderAddrItem)
	}
	var _recipientAddrRule []interface{}
	for _, _recipientAddrItem := range _recipientAddr {
		_recipientAddrRule = append(_recipientAddrRule, _recipientAddrItem)
	}
	var _swapIDRule []interface{}
	for _, _swapIDItem := range _swapID {
		_swapIDRule = append(_swapIDRule, _swapIDItem)
	}

	logs, sub, err := _Manager.contract.FilterLogs(opts, "CreateUnbind", _senderAddrRule, _recipientAddrRule, _swapIDRule)
	if err != nil {
		return nil, err
	}
	return &ManagerCreateUnbindIterator{contract: _Manager.contract, event: "CreateUnbind", logs: logs, sub: sub}, nil
}

// WatchCreateUnbind is a free log subscription operation binding the contract event 0xf3f48b99c0f6ef1f16092aed9ebd7d552e8b7647e0332990550d7626d7857ba5.
//
// Solidity: event CreateUnbind(address indexed _senderAddr, address indexed _recipientAddr, bytes32 indexed _swapID, bytes32 _randomNumberHash, uint64 _timestamp, bytes20 _refundAddr, uint256 _expireHeight, uint256 _outAmount, uint256 _laAmount)
func (_Manager *ManagerFilterer) WatchCreateUnbind(opts *bind.WatchOpts, sink chan<- *ManagerCreateUnbind, _senderAddr []common.Address, _recipientAddr []common.Address, _swapID [][32]byte) (event.Subscription, error) {

	var _senderAddrRule []interface{}
	for _, _senderAddrItem := range _senderAddr {
		_senderAddrRule = append(_senderAddrRule, _senderAddrItem)
	}
	var _recipientAddrRule []interface{}
	for _, _recipientAddrItem := range _recipientAddr {
		_recipientAddrRule = append(_recipientAddrRule, _recipientAddrItem)
	}
	var _swapIDRule []interface{}
	for _, _swapIDItem := range _swapID {
		_swapIDRule = append(_swapIDRule, _swapIDItem)
	}

	logs, sub, err := _Manager.contract.WatchLogs(opts, "CreateUnbind", _senderAddrRule, _recipientAddrRule, _swapIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManagerCreateUnbind)
				if err := _Manager.contract.UnpackLog(event, "CreateUnbind", log); err != nil {
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

// ParseCreateUnbind is a log parse operation binding the contract event 0xf3f48b99c0f6ef1f16092aed9ebd7d552e8b7647e0332990550d7626d7857ba5.
//
// Solidity: event CreateUnbind(address indexed _senderAddr, address indexed _recipientAddr, bytes32 indexed _swapID, bytes32 _randomNumberHash, uint64 _timestamp, bytes20 _refundAddr, uint256 _expireHeight, uint256 _outAmount, uint256 _laAmount)
func (_Manager *ManagerFilterer) ParseCreateUnbind(log types.Log) (*ManagerCreateUnbind, error) {
	event := new(ManagerCreateUnbind)
	if err := _Manager.contract.UnpackLog(event, "CreateUnbind", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ManagerCreatedIterator is returned from FilterCreated and is used to iterate over the raw logs and unpacked data for Created events raised by the Manager contract.
type ManagerCreatedIterator struct {
	Event *ManagerCreated // Event containing the contract specifics and raw log

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
func (it *ManagerCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManagerCreated)
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
		it.Event = new(ManagerCreated)
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
func (it *ManagerCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManagerCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManagerCreated represents a Created event raised by the Manager contract.
type ManagerCreated struct {
	MsgSender common.Address
	SwapID    [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCreated is a free log retrieval operation binding the contract event 0xd24d824860e7f36c08df207d22094623f901c3a69631edd487e915e1fac5d41a.
//
// Solidity: event Created(address indexed _msgSender, bytes32 indexed _swapID)
func (_Manager *ManagerFilterer) FilterCreated(opts *bind.FilterOpts, _msgSender []common.Address, _swapID [][32]byte) (*ManagerCreatedIterator, error) {

	var _msgSenderRule []interface{}
	for _, _msgSenderItem := range _msgSender {
		_msgSenderRule = append(_msgSenderRule, _msgSenderItem)
	}
	var _swapIDRule []interface{}
	for _, _swapIDItem := range _swapID {
		_swapIDRule = append(_swapIDRule, _swapIDItem)
	}

	logs, sub, err := _Manager.contract.FilterLogs(opts, "Created", _msgSenderRule, _swapIDRule)
	if err != nil {
		return nil, err
	}
	return &ManagerCreatedIterator{contract: _Manager.contract, event: "Created", logs: logs, sub: sub}, nil
}

// WatchCreated is a free log subscription operation binding the contract event 0xd24d824860e7f36c08df207d22094623f901c3a69631edd487e915e1fac5d41a.
//
// Solidity: event Created(address indexed _msgSender, bytes32 indexed _swapID)
func (_Manager *ManagerFilterer) WatchCreated(opts *bind.WatchOpts, sink chan<- *ManagerCreated, _msgSender []common.Address, _swapID [][32]byte) (event.Subscription, error) {

	var _msgSenderRule []interface{}
	for _, _msgSenderItem := range _msgSender {
		_msgSenderRule = append(_msgSenderRule, _msgSenderItem)
	}
	var _swapIDRule []interface{}
	for _, _swapIDItem := range _swapID {
		_swapIDRule = append(_swapIDRule, _swapIDItem)
	}

	logs, sub, err := _Manager.contract.WatchLogs(opts, "Created", _msgSenderRule, _swapIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManagerCreated)
				if err := _Manager.contract.UnpackLog(event, "Created", log); err != nil {
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

// ParseCreated is a log parse operation binding the contract event 0xd24d824860e7f36c08df207d22094623f901c3a69631edd487e915e1fac5d41a.
//
// Solidity: event Created(address indexed _msgSender, bytes32 indexed _swapID)
func (_Manager *ManagerFilterer) ParseCreated(log types.Log) (*ManagerCreated, error) {
	event := new(ManagerCreated)
	if err := _Manager.contract.UnpackLog(event, "Created", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ManagerHTLTIterator is returned from FilterHTLT and is used to iterate over the raw logs and unpacked data for HTLT events raised by the Manager contract.
type ManagerHTLTIterator struct {
	Event *ManagerHTLT // Event containing the contract specifics and raw log

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
func (it *ManagerHTLTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManagerHTLT)
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
		it.Event = new(ManagerHTLT)
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
func (it *ManagerHTLTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManagerHTLTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManagerHTLT represents a HTLT event raised by the Manager contract.
type ManagerHTLT struct {
	SenderAddr       common.Address
	RecipientAddr    common.Address
	SwapID           [32]byte
	RandomNumberHash [32]byte
	Timestamp        uint64
	RefundAddr       [20]byte
	ExpireHeight     *big.Int
	OutAmount        *big.Int
	LaAmount         *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterHTLT is a free log retrieval operation binding the contract event 0xb3e26d98380491276a8dce9d38fd1049e89070230ff5f36ebb55ead64500ade1.
//
// Solidity: event HTLT(address indexed _senderAddr, address indexed _recipientAddr, bytes32 indexed _swapID, bytes32 _randomNumberHash, uint64 _timestamp, bytes20 _refundAddr, uint256 _expireHeight, uint256 _outAmount, uint256 _laAmount)
func (_Manager *ManagerFilterer) FilterHTLT(opts *bind.FilterOpts, _senderAddr []common.Address, _recipientAddr []common.Address, _swapID [][32]byte) (*ManagerHTLTIterator, error) {

	var _senderAddrRule []interface{}
	for _, _senderAddrItem := range _senderAddr {
		_senderAddrRule = append(_senderAddrRule, _senderAddrItem)
	}
	var _recipientAddrRule []interface{}
	for _, _recipientAddrItem := range _recipientAddr {
		_recipientAddrRule = append(_recipientAddrRule, _recipientAddrItem)
	}
	var _swapIDRule []interface{}
	for _, _swapIDItem := range _swapID {
		_swapIDRule = append(_swapIDRule, _swapIDItem)
	}

	logs, sub, err := _Manager.contract.FilterLogs(opts, "HTLT", _senderAddrRule, _recipientAddrRule, _swapIDRule)
	if err != nil {
		return nil, err
	}
	return &ManagerHTLTIterator{contract: _Manager.contract, event: "HTLT", logs: logs, sub: sub}, nil
}

// WatchHTLT is a free log subscription operation binding the contract event 0xb3e26d98380491276a8dce9d38fd1049e89070230ff5f36ebb55ead64500ade1.
//
// Solidity: event HTLT(address indexed _senderAddr, address indexed _recipientAddr, bytes32 indexed _swapID, bytes32 _randomNumberHash, uint64 _timestamp, bytes20 _refundAddr, uint256 _expireHeight, uint256 _outAmount, uint256 _laAmount)
func (_Manager *ManagerFilterer) WatchHTLT(opts *bind.WatchOpts, sink chan<- *ManagerHTLT, _senderAddr []common.Address, _recipientAddr []common.Address, _swapID [][32]byte) (event.Subscription, error) {

	var _senderAddrRule []interface{}
	for _, _senderAddrItem := range _senderAddr {
		_senderAddrRule = append(_senderAddrRule, _senderAddrItem)
	}
	var _recipientAddrRule []interface{}
	for _, _recipientAddrItem := range _recipientAddr {
		_recipientAddrRule = append(_recipientAddrRule, _recipientAddrItem)
	}
	var _swapIDRule []interface{}
	for _, _swapIDItem := range _swapID {
		_swapIDRule = append(_swapIDRule, _swapIDItem)
	}

	logs, sub, err := _Manager.contract.WatchLogs(opts, "HTLT", _senderAddrRule, _recipientAddrRule, _swapIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManagerHTLT)
				if err := _Manager.contract.UnpackLog(event, "HTLT", log); err != nil {
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

// ParseHTLT is a log parse operation binding the contract event 0xb3e26d98380491276a8dce9d38fd1049e89070230ff5f36ebb55ead64500ade1.
//
// Solidity: event HTLT(address indexed _senderAddr, address indexed _recipientAddr, bytes32 indexed _swapID, bytes32 _randomNumberHash, uint64 _timestamp, bytes20 _refundAddr, uint256 _expireHeight, uint256 _outAmount, uint256 _laAmount)
func (_Manager *ManagerFilterer) ParseHTLT(log types.Log) (*ManagerHTLT, error) {
	event := new(ManagerHTLT)
	if err := _Manager.contract.UnpackLog(event, "HTLT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ManagerSpendedIterator is returned from FilterSpended and is used to iterate over the raw logs and unpacked data for Spended events raised by the Manager contract.
type ManagerSpendedIterator struct {
	Event *ManagerSpended // Event containing the contract specifics and raw log

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
func (it *ManagerSpendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManagerSpended)
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
		it.Event = new(ManagerSpended)
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
func (it *ManagerSpendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManagerSpendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManagerSpended represents a Spended event raised by the Manager contract.
type ManagerSpended struct {
	Erc20Addr      common.Address
	LaContractAddr common.Address
	ToAddress      common.Address
	Amount         *big.Int
	Txhash         [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSpended is a free log retrieval operation binding the contract event 0x94f942d6b1c0be8a7cc7305822711516127d4abb5059dfa382d1a0ba52bed688.
//
// Solidity: event Spended(address indexed erc20Addr, address indexed laContractAddr, address indexed toAddress, uint256 amount, bytes32 txhash)
func (_Manager *ManagerFilterer) FilterSpended(opts *bind.FilterOpts, erc20Addr []common.Address, laContractAddr []common.Address, toAddress []common.Address) (*ManagerSpendedIterator, error) {

	var erc20AddrRule []interface{}
	for _, erc20AddrItem := range erc20Addr {
		erc20AddrRule = append(erc20AddrRule, erc20AddrItem)
	}
	var laContractAddrRule []interface{}
	for _, laContractAddrItem := range laContractAddr {
		laContractAddrRule = append(laContractAddrRule, laContractAddrItem)
	}
	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _Manager.contract.FilterLogs(opts, "Spended", erc20AddrRule, laContractAddrRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return &ManagerSpendedIterator{contract: _Manager.contract, event: "Spended", logs: logs, sub: sub}, nil
}

// WatchSpended is a free log subscription operation binding the contract event 0x94f942d6b1c0be8a7cc7305822711516127d4abb5059dfa382d1a0ba52bed688.
//
// Solidity: event Spended(address indexed erc20Addr, address indexed laContractAddr, address indexed toAddress, uint256 amount, bytes32 txhash)
func (_Manager *ManagerFilterer) WatchSpended(opts *bind.WatchOpts, sink chan<- *ManagerSpended, erc20Addr []common.Address, laContractAddr []common.Address, toAddress []common.Address) (event.Subscription, error) {

	var erc20AddrRule []interface{}
	for _, erc20AddrItem := range erc20Addr {
		erc20AddrRule = append(erc20AddrRule, erc20AddrItem)
	}
	var laContractAddrRule []interface{}
	for _, laContractAddrItem := range laContractAddr {
		laContractAddrRule = append(laContractAddrRule, laContractAddrItem)
	}
	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _Manager.contract.WatchLogs(opts, "Spended", erc20AddrRule, laContractAddrRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManagerSpended)
				if err := _Manager.contract.UnpackLog(event, "Spended", log); err != nil {
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

// ParseSpended is a log parse operation binding the contract event 0x94f942d6b1c0be8a7cc7305822711516127d4abb5059dfa382d1a0ba52bed688.
//
// Solidity: event Spended(address indexed erc20Addr, address indexed laContractAddr, address indexed toAddress, uint256 amount, bytes32 txhash)
func (_Manager *ManagerFilterer) ParseSpended(log types.Log) (*ManagerSpended, error) {
	event := new(ManagerSpended)
	if err := _Manager.contract.UnpackLog(event, "Spended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ManagerSwappedIterator is returned from FilterSwapped and is used to iterate over the raw logs and unpacked data for Swapped events raised by the Manager contract.
type ManagerSwappedIterator struct {
	Event *ManagerSwapped // Event containing the contract specifics and raw log

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
func (it *ManagerSwappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManagerSwapped)
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
		it.Event = new(ManagerSwapped)
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
func (it *ManagerSwappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManagerSwappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManagerSwapped represents a Swapped event raised by the Manager contract.
type ManagerSwapped struct {
	Erc20Addr common.Address
	FromAddr  common.Address
	Amount    *big.Int
	FeeAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSwapped is a free log retrieval operation binding the contract event 0xa078c4190abe07940190effc1846be0ccf03ad6007bc9e93f9697d0b460befbb.
//
// Solidity: event Swapped(address indexed erc20Addr, address indexed fromAddr, uint256 amount, uint256 feeAmount)
func (_Manager *ManagerFilterer) FilterSwapped(opts *bind.FilterOpts, erc20Addr []common.Address, fromAddr []common.Address) (*ManagerSwappedIterator, error) {

	var erc20AddrRule []interface{}
	for _, erc20AddrItem := range erc20Addr {
		erc20AddrRule = append(erc20AddrRule, erc20AddrItem)
	}
	var fromAddrRule []interface{}
	for _, fromAddrItem := range fromAddr {
		fromAddrRule = append(fromAddrRule, fromAddrItem)
	}

	logs, sub, err := _Manager.contract.FilterLogs(opts, "Swapped", erc20AddrRule, fromAddrRule)
	if err != nil {
		return nil, err
	}
	return &ManagerSwappedIterator{contract: _Manager.contract, event: "Swapped", logs: logs, sub: sub}, nil
}

// WatchSwapped is a free log subscription operation binding the contract event 0xa078c4190abe07940190effc1846be0ccf03ad6007bc9e93f9697d0b460befbb.
//
// Solidity: event Swapped(address indexed erc20Addr, address indexed fromAddr, uint256 amount, uint256 feeAmount)
func (_Manager *ManagerFilterer) WatchSwapped(opts *bind.WatchOpts, sink chan<- *ManagerSwapped, erc20Addr []common.Address, fromAddr []common.Address) (event.Subscription, error) {

	var erc20AddrRule []interface{}
	for _, erc20AddrItem := range erc20Addr {
		erc20AddrRule = append(erc20AddrRule, erc20AddrItem)
	}
	var fromAddrRule []interface{}
	for _, fromAddrItem := range fromAddr {
		fromAddrRule = append(fromAddrRule, fromAddrItem)
	}

	logs, sub, err := _Manager.contract.WatchLogs(opts, "Swapped", erc20AddrRule, fromAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManagerSwapped)
				if err := _Manager.contract.UnpackLog(event, "Swapped", log); err != nil {
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

// ParseSwapped is a log parse operation binding the contract event 0xa078c4190abe07940190effc1846be0ccf03ad6007bc9e93f9697d0b460befbb.
//
// Solidity: event Swapped(address indexed erc20Addr, address indexed fromAddr, uint256 amount, uint256 feeAmount)
func (_Manager *ManagerFilterer) ParseSwapped(log types.Log) (*ManagerSwapped, error) {
	event := new(ManagerSwapped)
	if err := _Manager.contract.UnpackLog(event, "Swapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ManagerUnbindedIterator is returned from FilterUnbinded and is used to iterate over the raw logs and unpacked data for Unbinded events raised by the Manager contract.
type ManagerUnbindedIterator struct {
	Event *ManagerUnbinded // Event containing the contract specifics and raw log

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
func (it *ManagerUnbindedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManagerUnbinded)
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
		it.Event = new(ManagerUnbinded)
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
func (it *ManagerUnbindedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManagerUnbindedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManagerUnbinded represents a Unbinded event raised by the Manager contract.
type ManagerUnbinded struct {
	SenderAddr       common.Address
	RecipientAddr    common.Address
	SwapID           [32]byte
	RandomNumberHash [32]byte
	Timestamp        uint64
	RefundAddr       [20]byte
	ExpireHeight     *big.Int
	OutAmount        *big.Int
	LaAmount         *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUnbinded is a free log retrieval operation binding the contract event 0x6508ef685db702729b794640245ca23228b5c5fab476f9e4e45e688f1fcee4d1.
//
// Solidity: event Unbinded(address indexed _senderAddr, address indexed _recipientAddr, bytes32 indexed _swapID, bytes32 _randomNumberHash, uint64 _timestamp, bytes20 _refundAddr, uint256 _expireHeight, uint256 _outAmount, uint256 _laAmount)
func (_Manager *ManagerFilterer) FilterUnbinded(opts *bind.FilterOpts, _senderAddr []common.Address, _recipientAddr []common.Address, _swapID [][32]byte) (*ManagerUnbindedIterator, error) {

	var _senderAddrRule []interface{}
	for _, _senderAddrItem := range _senderAddr {
		_senderAddrRule = append(_senderAddrRule, _senderAddrItem)
	}
	var _recipientAddrRule []interface{}
	for _, _recipientAddrItem := range _recipientAddr {
		_recipientAddrRule = append(_recipientAddrRule, _recipientAddrItem)
	}
	var _swapIDRule []interface{}
	for _, _swapIDItem := range _swapID {
		_swapIDRule = append(_swapIDRule, _swapIDItem)
	}

	logs, sub, err := _Manager.contract.FilterLogs(opts, "Unbinded", _senderAddrRule, _recipientAddrRule, _swapIDRule)
	if err != nil {
		return nil, err
	}
	return &ManagerUnbindedIterator{contract: _Manager.contract, event: "Unbinded", logs: logs, sub: sub}, nil
}

// WatchUnbinded is a free log subscription operation binding the contract event 0x6508ef685db702729b794640245ca23228b5c5fab476f9e4e45e688f1fcee4d1.
//
// Solidity: event Unbinded(address indexed _senderAddr, address indexed _recipientAddr, bytes32 indexed _swapID, bytes32 _randomNumberHash, uint64 _timestamp, bytes20 _refundAddr, uint256 _expireHeight, uint256 _outAmount, uint256 _laAmount)
func (_Manager *ManagerFilterer) WatchUnbinded(opts *bind.WatchOpts, sink chan<- *ManagerUnbinded, _senderAddr []common.Address, _recipientAddr []common.Address, _swapID [][32]byte) (event.Subscription, error) {

	var _senderAddrRule []interface{}
	for _, _senderAddrItem := range _senderAddr {
		_senderAddrRule = append(_senderAddrRule, _senderAddrItem)
	}
	var _recipientAddrRule []interface{}
	for _, _recipientAddrItem := range _recipientAddr {
		_recipientAddrRule = append(_recipientAddrRule, _recipientAddrItem)
	}
	var _swapIDRule []interface{}
	for _, _swapIDItem := range _swapID {
		_swapIDRule = append(_swapIDRule, _swapIDItem)
	}

	logs, sub, err := _Manager.contract.WatchLogs(opts, "Unbinded", _senderAddrRule, _recipientAddrRule, _swapIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManagerUnbinded)
				if err := _Manager.contract.UnpackLog(event, "Unbinded", log); err != nil {
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

// ParseUnbinded is a log parse operation binding the contract event 0x6508ef685db702729b794640245ca23228b5c5fab476f9e4e45e688f1fcee4d1.
//
// Solidity: event Unbinded(address indexed _senderAddr, address indexed _recipientAddr, bytes32 indexed _swapID, bytes32 _randomNumberHash, uint64 _timestamp, bytes20 _refundAddr, uint256 _expireHeight, uint256 _outAmount, uint256 _laAmount)
func (_Manager *ManagerFilterer) ParseUnbinded(log types.Log) (*ManagerUnbinded, error) {
	event := new(ManagerUnbinded)
	if err := _Manager.contract.UnpackLog(event, "Unbinded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ManagerBindFailureIterator is returned from FilterBindFailure and is used to iterate over the raw logs and unpacked data for BindFailure events raised by the Manager contract.
type ManagerBindFailureIterator struct {
	Event *ManagerBindFailure // Event containing the contract specifics and raw log

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
func (it *ManagerBindFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManagerBindFailure)
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
		it.Event = new(ManagerBindFailure)
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
func (it *ManagerBindFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManagerBindFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManagerBindFailure represents a BindFailure event raised by the Manager contract.
type ManagerBindFailure struct {
	ContractAddr common.Address
	Bep2Symbol   string
	FailedReason uint32
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBindFailure is a free log retrieval operation binding the contract event 0x831c0ef4d93bda3bce08b69ae3f29ef1a6e052b833200988554158494405a107.
//
// Solidity: event bindFailure(address indexed contractAddr, string bep2Symbol, uint32 failedReason)
func (_Manager *ManagerFilterer) FilterBindFailure(opts *bind.FilterOpts, contractAddr []common.Address) (*ManagerBindFailureIterator, error) {

	var contractAddrRule []interface{}
	for _, contractAddrItem := range contractAddr {
		contractAddrRule = append(contractAddrRule, contractAddrItem)
	}

	logs, sub, err := _Manager.contract.FilterLogs(opts, "bindFailure", contractAddrRule)
	if err != nil {
		return nil, err
	}
	return &ManagerBindFailureIterator{contract: _Manager.contract, event: "bindFailure", logs: logs, sub: sub}, nil
}

// WatchBindFailure is a free log subscription operation binding the contract event 0x831c0ef4d93bda3bce08b69ae3f29ef1a6e052b833200988554158494405a107.
//
// Solidity: event bindFailure(address indexed contractAddr, string bep2Symbol, uint32 failedReason)
func (_Manager *ManagerFilterer) WatchBindFailure(opts *bind.WatchOpts, sink chan<- *ManagerBindFailure, contractAddr []common.Address) (event.Subscription, error) {

	var contractAddrRule []interface{}
	for _, contractAddrItem := range contractAddr {
		contractAddrRule = append(contractAddrRule, contractAddrItem)
	}

	logs, sub, err := _Manager.contract.WatchLogs(opts, "bindFailure", contractAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManagerBindFailure)
				if err := _Manager.contract.UnpackLog(event, "bindFailure", log); err != nil {
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

// ParseBindFailure is a log parse operation binding the contract event 0x831c0ef4d93bda3bce08b69ae3f29ef1a6e052b833200988554158494405a107.
//
// Solidity: event bindFailure(address indexed contractAddr, string bep2Symbol, uint32 failedReason)
func (_Manager *ManagerFilterer) ParseBindFailure(log types.Log) (*ManagerBindFailure, error) {
	event := new(ManagerBindFailure)
	if err := _Manager.contract.UnpackLog(event, "bindFailure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
