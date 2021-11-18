// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridge

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

// BridgeProposal is an auto generated low-level Go binding around an user-defined struct.
type BridgeProposal struct {
	ResourceID    [32]byte
	DataHash      [32]byte
	YesVotes      []common.Address
	NoVotes       []common.Address
	Status        uint8
	ProposedBlock *big.Int
}

// BridgeABI is the input ABI used to generate the binding from.
const BridgeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes8\",\"name\":\"destinationChainID\",\"type\":\"bytes8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes8\",\"name\":\"originChainID\",\"type\":\"bytes8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumBridge.ProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"ProposalEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes8\",\"name\":\"originChainID\",\"type\":\"bytes8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumBridge.ProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"}],\"name\":\"ProposalVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"RelayerThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardCollected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"INIT_CHAINID\",\"outputs\":[{\"internalType\":\"bytes8\",\"name\":\"\",\"type\":\"bytes8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_EXPIRY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_RELAYERTHRESHOLD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_chainID\",\"outputs\":[{\"internalType\":\"bytes8\",\"name\":\"\",\"type\":\"bytes8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes8\",\"name\":\"\",\"type\":\"bytes8\"}],\"name\":\"_depositCounts\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bytes8\",\"name\":\"\",\"type\":\"bytes8\"}],\"name\":\"_depositRecords\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_expiry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_hasVotedOnProposal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_proposals\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"enumBridge.ProposalStatus\",\"name\":\"_status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_proposedBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_relayerHubAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_relayerThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToHandlerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_totalProposals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"adminChangeFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"adminChangeRelayerThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"adminCollectFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminPauseTransfers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"adminSetBurnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRelayerHub\",\"type\":\"address\"}],\"name\":\"adminSetRelayerHub\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"adminSetResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminUnpauseTransfers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOrTokenID\",\"type\":\"uint256\"}],\"name\":\"adminWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes8\",\"name\":\"chainID\",\"type\":\"bytes8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"}],\"name\":\"cancelProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes8\",\"name\":\"destinationChainID\",\"type\":\"bytes8\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes8\",\"name\":\"chainID\",\"type\":\"bytes8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes8\",\"name\":\"originChainID\",\"type\":\"bytes8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"getProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"_yesVotes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_noVotes\",\"type\":\"address[]\"},{\"internalType\":\"enumBridge.ProposalStatus\",\"name\":\"_status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_proposedBlock\",\"type\":\"uint256\"}],\"internalType\":\"structBridge.Proposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"relayerCollectReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"sub\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes8\",\"name\":\"chainID\",\"type\":\"bytes8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"voteProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// INITCHAINID is a free data retrieval call binding the contract method 0x45fb7494.
//
// Solidity: function INIT_CHAINID() view returns(bytes8)
func (_Bridge *BridgeCaller) INITCHAINID(opts *bind.CallOpts) ([8]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "INIT_CHAINID")

	if err != nil {
		return *new([8]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([8]byte)).(*[8]byte)

	return out0, err

}

// INITCHAINID is a free data retrieval call binding the contract method 0x45fb7494.
//
// Solidity: function INIT_CHAINID() view returns(bytes8)
func (_Bridge *BridgeSession) INITCHAINID() ([8]byte, error) {
	return _Bridge.Contract.INITCHAINID(&_Bridge.CallOpts)
}

// INITCHAINID is a free data retrieval call binding the contract method 0x45fb7494.
//
// Solidity: function INIT_CHAINID() view returns(bytes8)
func (_Bridge *BridgeCallerSession) INITCHAINID() ([8]byte, error) {
	return _Bridge.Contract.INITCHAINID(&_Bridge.CallOpts)
}

// INITEXPIRY is a free data retrieval call binding the contract method 0x0461ce72.
//
// Solidity: function INIT_EXPIRY() view returns(uint256)
func (_Bridge *BridgeCaller) INITEXPIRY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "INIT_EXPIRY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITEXPIRY is a free data retrieval call binding the contract method 0x0461ce72.
//
// Solidity: function INIT_EXPIRY() view returns(uint256)
func (_Bridge *BridgeSession) INITEXPIRY() (*big.Int, error) {
	return _Bridge.Contract.INITEXPIRY(&_Bridge.CallOpts)
}

// INITEXPIRY is a free data retrieval call binding the contract method 0x0461ce72.
//
// Solidity: function INIT_EXPIRY() view returns(uint256)
func (_Bridge *BridgeCallerSession) INITEXPIRY() (*big.Int, error) {
	return _Bridge.Contract.INITEXPIRY(&_Bridge.CallOpts)
}

// INITFEE is a free data retrieval call binding the contract method 0x82b4beec.
//
// Solidity: function INIT_FEE() view returns(uint256)
func (_Bridge *BridgeCaller) INITFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "INIT_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITFEE is a free data retrieval call binding the contract method 0x82b4beec.
//
// Solidity: function INIT_FEE() view returns(uint256)
func (_Bridge *BridgeSession) INITFEE() (*big.Int, error) {
	return _Bridge.Contract.INITFEE(&_Bridge.CallOpts)
}

// INITFEE is a free data retrieval call binding the contract method 0x82b4beec.
//
// Solidity: function INIT_FEE() view returns(uint256)
func (_Bridge *BridgeCallerSession) INITFEE() (*big.Int, error) {
	return _Bridge.Contract.INITFEE(&_Bridge.CallOpts)
}

// INITRELAYERTHRESHOLD is a free data retrieval call binding the contract method 0xcacc26ed.
//
// Solidity: function INIT_RELAYERTHRESHOLD() view returns(uint256)
func (_Bridge *BridgeCaller) INITRELAYERTHRESHOLD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "INIT_RELAYERTHRESHOLD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITRELAYERTHRESHOLD is a free data retrieval call binding the contract method 0xcacc26ed.
//
// Solidity: function INIT_RELAYERTHRESHOLD() view returns(uint256)
func (_Bridge *BridgeSession) INITRELAYERTHRESHOLD() (*big.Int, error) {
	return _Bridge.Contract.INITRELAYERTHRESHOLD(&_Bridge.CallOpts)
}

// INITRELAYERTHRESHOLD is a free data retrieval call binding the contract method 0xcacc26ed.
//
// Solidity: function INIT_RELAYERTHRESHOLD() view returns(uint256)
func (_Bridge *BridgeCallerSession) INITRELAYERTHRESHOLD() (*big.Int, error) {
	return _Bridge.Contract.INITRELAYERTHRESHOLD(&_Bridge.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xbeab7131.
//
// Solidity: function _chainID() view returns(bytes8)
func (_Bridge *BridgeCaller) ChainID(opts *bind.CallOpts) ([8]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_chainID")

	if err != nil {
		return *new([8]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([8]byte)).(*[8]byte)

	return out0, err

}

// ChainID is a free data retrieval call binding the contract method 0xbeab7131.
//
// Solidity: function _chainID() view returns(bytes8)
func (_Bridge *BridgeSession) ChainID() ([8]byte, error) {
	return _Bridge.Contract.ChainID(&_Bridge.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xbeab7131.
//
// Solidity: function _chainID() view returns(bytes8)
func (_Bridge *BridgeCallerSession) ChainID() ([8]byte, error) {
	return _Bridge.Contract.ChainID(&_Bridge.CallOpts)
}

// DepositCounts is a free data retrieval call binding the contract method 0x174e300b.
//
// Solidity: function _depositCounts(bytes8 ) view returns(uint64)
func (_Bridge *BridgeCaller) DepositCounts(opts *bind.CallOpts, arg0 [8]byte) (uint64, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_depositCounts", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// DepositCounts is a free data retrieval call binding the contract method 0x174e300b.
//
// Solidity: function _depositCounts(bytes8 ) view returns(uint64)
func (_Bridge *BridgeSession) DepositCounts(arg0 [8]byte) (uint64, error) {
	return _Bridge.Contract.DepositCounts(&_Bridge.CallOpts, arg0)
}

// DepositCounts is a free data retrieval call binding the contract method 0x174e300b.
//
// Solidity: function _depositCounts(bytes8 ) view returns(uint64)
func (_Bridge *BridgeCallerSession) DepositCounts(arg0 [8]byte) (uint64, error) {
	return _Bridge.Contract.DepositCounts(&_Bridge.CallOpts, arg0)
}

// DepositRecords is a free data retrieval call binding the contract method 0xa48ace38.
//
// Solidity: function _depositRecords(uint64 , bytes8 ) view returns(bytes)
func (_Bridge *BridgeCaller) DepositRecords(opts *bind.CallOpts, arg0 uint64, arg1 [8]byte) ([]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_depositRecords", arg0, arg1)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// DepositRecords is a free data retrieval call binding the contract method 0xa48ace38.
//
// Solidity: function _depositRecords(uint64 , bytes8 ) view returns(bytes)
func (_Bridge *BridgeSession) DepositRecords(arg0 uint64, arg1 [8]byte) ([]byte, error) {
	return _Bridge.Contract.DepositRecords(&_Bridge.CallOpts, arg0, arg1)
}

// DepositRecords is a free data retrieval call binding the contract method 0xa48ace38.
//
// Solidity: function _depositRecords(uint64 , bytes8 ) view returns(bytes)
func (_Bridge *BridgeCallerSession) DepositRecords(arg0 uint64, arg1 [8]byte) ([]byte, error) {
	return _Bridge.Contract.DepositRecords(&_Bridge.CallOpts, arg0, arg1)
}

// Expiry is a free data retrieval call binding the contract method 0xc5ec8970.
//
// Solidity: function _expiry() view returns(uint256)
func (_Bridge *BridgeCaller) Expiry(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_expiry")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Expiry is a free data retrieval call binding the contract method 0xc5ec8970.
//
// Solidity: function _expiry() view returns(uint256)
func (_Bridge *BridgeSession) Expiry() (*big.Int, error) {
	return _Bridge.Contract.Expiry(&_Bridge.CallOpts)
}

// Expiry is a free data retrieval call binding the contract method 0xc5ec8970.
//
// Solidity: function _expiry() view returns(uint256)
func (_Bridge *BridgeCallerSession) Expiry() (*big.Int, error) {
	return _Bridge.Contract.Expiry(&_Bridge.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xc5b37c22.
//
// Solidity: function _fee() view returns(uint256)
func (_Bridge *BridgeCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xc5b37c22.
//
// Solidity: function _fee() view returns(uint256)
func (_Bridge *BridgeSession) Fee() (*big.Int, error) {
	return _Bridge.Contract.Fee(&_Bridge.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xc5b37c22.
//
// Solidity: function _fee() view returns(uint256)
func (_Bridge *BridgeCallerSession) Fee() (*big.Int, error) {
	return _Bridge.Contract.Fee(&_Bridge.CallOpts)
}

// HasVotedOnProposal is a free data retrieval call binding the contract method 0x86c99f0e.
//
// Solidity: function _hasVotedOnProposal(bytes32 , bytes32 , address ) view returns(bool)
func (_Bridge *BridgeCaller) HasVotedOnProposal(opts *bind.CallOpts, arg0 [32]byte, arg1 [32]byte, arg2 common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_hasVotedOnProposal", arg0, arg1, arg2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasVotedOnProposal is a free data retrieval call binding the contract method 0x86c99f0e.
//
// Solidity: function _hasVotedOnProposal(bytes32 , bytes32 , address ) view returns(bool)
func (_Bridge *BridgeSession) HasVotedOnProposal(arg0 [32]byte, arg1 [32]byte, arg2 common.Address) (bool, error) {
	return _Bridge.Contract.HasVotedOnProposal(&_Bridge.CallOpts, arg0, arg1, arg2)
}

// HasVotedOnProposal is a free data retrieval call binding the contract method 0x86c99f0e.
//
// Solidity: function _hasVotedOnProposal(bytes32 , bytes32 , address ) view returns(bool)
func (_Bridge *BridgeCallerSession) HasVotedOnProposal(arg0 [32]byte, arg1 [32]byte, arg2 common.Address) (bool, error) {
	return _Bridge.Contract.HasVotedOnProposal(&_Bridge.CallOpts, arg0, arg1, arg2)
}

// Proposals is a free data retrieval call binding the contract method 0x7bda741c.
//
// Solidity: function _proposals(bytes32 , bytes32 ) view returns(bytes32 _resourceID, bytes32 _dataHash, uint8 _status, uint256 _proposedBlock)
func (_Bridge *BridgeCaller) Proposals(opts *bind.CallOpts, arg0 [32]byte, arg1 [32]byte) (struct {
	ResourceID    [32]byte
	DataHash      [32]byte
	Status        uint8
	ProposedBlock *big.Int
}, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_proposals", arg0, arg1)

	outstruct := new(struct {
		ResourceID    [32]byte
		DataHash      [32]byte
		Status        uint8
		ProposedBlock *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ResourceID = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.DataHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Status = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.ProposedBlock = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Proposals is a free data retrieval call binding the contract method 0x7bda741c.
//
// Solidity: function _proposals(bytes32 , bytes32 ) view returns(bytes32 _resourceID, bytes32 _dataHash, uint8 _status, uint256 _proposedBlock)
func (_Bridge *BridgeSession) Proposals(arg0 [32]byte, arg1 [32]byte) (struct {
	ResourceID    [32]byte
	DataHash      [32]byte
	Status        uint8
	ProposedBlock *big.Int
}, error) {
	return _Bridge.Contract.Proposals(&_Bridge.CallOpts, arg0, arg1)
}

// Proposals is a free data retrieval call binding the contract method 0x7bda741c.
//
// Solidity: function _proposals(bytes32 , bytes32 ) view returns(bytes32 _resourceID, bytes32 _dataHash, uint8 _status, uint256 _proposedBlock)
func (_Bridge *BridgeCallerSession) Proposals(arg0 [32]byte, arg1 [32]byte) (struct {
	ResourceID    [32]byte
	DataHash      [32]byte
	Status        uint8
	ProposedBlock *big.Int
}, error) {
	return _Bridge.Contract.Proposals(&_Bridge.CallOpts, arg0, arg1)
}

// RelayerHubAddress is a free data retrieval call binding the contract method 0xb6d97623.
//
// Solidity: function _relayerHubAddress() view returns(address)
func (_Bridge *BridgeCaller) RelayerHubAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_relayerHubAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RelayerHubAddress is a free data retrieval call binding the contract method 0xb6d97623.
//
// Solidity: function _relayerHubAddress() view returns(address)
func (_Bridge *BridgeSession) RelayerHubAddress() (common.Address, error) {
	return _Bridge.Contract.RelayerHubAddress(&_Bridge.CallOpts)
}

// RelayerHubAddress is a free data retrieval call binding the contract method 0xb6d97623.
//
// Solidity: function _relayerHubAddress() view returns(address)
func (_Bridge *BridgeCallerSession) RelayerHubAddress() (common.Address, error) {
	return _Bridge.Contract.RelayerHubAddress(&_Bridge.CallOpts)
}

// RelayerThreshold is a free data retrieval call binding the contract method 0xd7a9cd79.
//
// Solidity: function _relayerThreshold() view returns(uint256)
func (_Bridge *BridgeCaller) RelayerThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_relayerThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RelayerThreshold is a free data retrieval call binding the contract method 0xd7a9cd79.
//
// Solidity: function _relayerThreshold() view returns(uint256)
func (_Bridge *BridgeSession) RelayerThreshold() (*big.Int, error) {
	return _Bridge.Contract.RelayerThreshold(&_Bridge.CallOpts)
}

// RelayerThreshold is a free data retrieval call binding the contract method 0xd7a9cd79.
//
// Solidity: function _relayerThreshold() view returns(uint256)
func (_Bridge *BridgeCallerSession) RelayerThreshold() (*big.Int, error) {
	return _Bridge.Contract.RelayerThreshold(&_Bridge.CallOpts)
}

// ResourceIDToHandlerAddress is a free data retrieval call binding the contract method 0x84db809f.
//
// Solidity: function _resourceIDToHandlerAddress(bytes32 ) view returns(address)
func (_Bridge *BridgeCaller) ResourceIDToHandlerAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_resourceIDToHandlerAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResourceIDToHandlerAddress is a free data retrieval call binding the contract method 0x84db809f.
//
// Solidity: function _resourceIDToHandlerAddress(bytes32 ) view returns(address)
func (_Bridge *BridgeSession) ResourceIDToHandlerAddress(arg0 [32]byte) (common.Address, error) {
	return _Bridge.Contract.ResourceIDToHandlerAddress(&_Bridge.CallOpts, arg0)
}

// ResourceIDToHandlerAddress is a free data retrieval call binding the contract method 0x84db809f.
//
// Solidity: function _resourceIDToHandlerAddress(bytes32 ) view returns(address)
func (_Bridge *BridgeCallerSession) ResourceIDToHandlerAddress(arg0 [32]byte) (common.Address, error) {
	return _Bridge.Contract.ResourceIDToHandlerAddress(&_Bridge.CallOpts, arg0)
}

// TotalProposals is a free data retrieval call binding the contract method 0x9d5773e0.
//
// Solidity: function _totalProposals() view returns(uint256)
func (_Bridge *BridgeCaller) TotalProposals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_totalProposals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalProposals is a free data retrieval call binding the contract method 0x9d5773e0.
//
// Solidity: function _totalProposals() view returns(uint256)
func (_Bridge *BridgeSession) TotalProposals() (*big.Int, error) {
	return _Bridge.Contract.TotalProposals(&_Bridge.CallOpts)
}

// TotalProposals is a free data retrieval call binding the contract method 0x9d5773e0.
//
// Solidity: function _totalProposals() view returns(uint256)
func (_Bridge *BridgeCallerSession) TotalProposals() (*big.Int, error) {
	return _Bridge.Contract.TotalProposals(&_Bridge.CallOpts)
}

// GetProposal is a free data retrieval call binding the contract method 0x3c69171b.
//
// Solidity: function getProposal(bytes8 originChainID, uint64 depositNonce, bytes32 dataHash) view returns((bytes32,bytes32,address[],address[],uint8,uint256))
func (_Bridge *BridgeCaller) GetProposal(opts *bind.CallOpts, originChainID [8]byte, depositNonce uint64, dataHash [32]byte) (BridgeProposal, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getProposal", originChainID, depositNonce, dataHash)

	if err != nil {
		return *new(BridgeProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(BridgeProposal)).(*BridgeProposal)

	return out0, err

}

// GetProposal is a free data retrieval call binding the contract method 0x3c69171b.
//
// Solidity: function getProposal(bytes8 originChainID, uint64 depositNonce, bytes32 dataHash) view returns((bytes32,bytes32,address[],address[],uint8,uint256))
func (_Bridge *BridgeSession) GetProposal(originChainID [8]byte, depositNonce uint64, dataHash [32]byte) (BridgeProposal, error) {
	return _Bridge.Contract.GetProposal(&_Bridge.CallOpts, originChainID, depositNonce, dataHash)
}

// GetProposal is a free data retrieval call binding the contract method 0x3c69171b.
//
// Solidity: function getProposal(bytes8 originChainID, uint64 depositNonce, bytes32 dataHash) view returns((bytes32,bytes32,address[],address[],uint8,uint256))
func (_Bridge *BridgeCallerSession) GetProposal(originChainID [8]byte, depositNonce uint64, dataHash [32]byte) (BridgeProposal, error) {
	return _Bridge.Contract.GetProposal(&_Bridge.CallOpts, originChainID, depositNonce, dataHash)
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

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bridge *BridgeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bridge *BridgeSession) Paused() (bool, error) {
	return _Bridge.Contract.Paused(&_Bridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bridge *BridgeCallerSession) Paused() (bool, error) {
	return _Bridge.Contract.Paused(&_Bridge.CallOpts)
}

// Sub is a free data retrieval call binding the contract method 0xb67d77c5.
//
// Solidity: function sub(uint256 a, uint256 b) pure returns(uint256)
func (_Bridge *BridgeCaller) Sub(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "sub", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Sub is a free data retrieval call binding the contract method 0xb67d77c5.
//
// Solidity: function sub(uint256 a, uint256 b) pure returns(uint256)
func (_Bridge *BridgeSession) Sub(a *big.Int, b *big.Int) (*big.Int, error) {
	return _Bridge.Contract.Sub(&_Bridge.CallOpts, a, b)
}

// Sub is a free data retrieval call binding the contract method 0xb67d77c5.
//
// Solidity: function sub(uint256 a, uint256 b) pure returns(uint256)
func (_Bridge *BridgeCallerSession) Sub(a *big.Int, b *big.Int) (*big.Int, error) {
	return _Bridge.Contract.Sub(&_Bridge.CallOpts, a, b)
}

// AdminChangeFee is a paid mutator transaction binding the contract method 0x91c404ac.
//
// Solidity: function adminChangeFee(uint256 newFee) returns()
func (_Bridge *BridgeTransactor) AdminChangeFee(opts *bind.TransactOpts, newFee *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "adminChangeFee", newFee)
}

// AdminChangeFee is a paid mutator transaction binding the contract method 0x91c404ac.
//
// Solidity: function adminChangeFee(uint256 newFee) returns()
func (_Bridge *BridgeSession) AdminChangeFee(newFee *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.AdminChangeFee(&_Bridge.TransactOpts, newFee)
}

// AdminChangeFee is a paid mutator transaction binding the contract method 0x91c404ac.
//
// Solidity: function adminChangeFee(uint256 newFee) returns()
func (_Bridge *BridgeTransactorSession) AdminChangeFee(newFee *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.AdminChangeFee(&_Bridge.TransactOpts, newFee)
}

// AdminChangeRelayerThreshold is a paid mutator transaction binding the contract method 0x4e056005.
//
// Solidity: function adminChangeRelayerThreshold(uint256 newThreshold) returns()
func (_Bridge *BridgeTransactor) AdminChangeRelayerThreshold(opts *bind.TransactOpts, newThreshold *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "adminChangeRelayerThreshold", newThreshold)
}

// AdminChangeRelayerThreshold is a paid mutator transaction binding the contract method 0x4e056005.
//
// Solidity: function adminChangeRelayerThreshold(uint256 newThreshold) returns()
func (_Bridge *BridgeSession) AdminChangeRelayerThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.AdminChangeRelayerThreshold(&_Bridge.TransactOpts, newThreshold)
}

// AdminChangeRelayerThreshold is a paid mutator transaction binding the contract method 0x4e056005.
//
// Solidity: function adminChangeRelayerThreshold(uint256 newThreshold) returns()
func (_Bridge *BridgeTransactorSession) AdminChangeRelayerThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.AdminChangeRelayerThreshold(&_Bridge.TransactOpts, newThreshold)
}

// AdminCollectFees is a paid mutator transaction binding the contract method 0xe77b33fd.
//
// Solidity: function adminCollectFees(uint256 amount) returns()
func (_Bridge *BridgeTransactor) AdminCollectFees(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "adminCollectFees", amount)
}

// AdminCollectFees is a paid mutator transaction binding the contract method 0xe77b33fd.
//
// Solidity: function adminCollectFees(uint256 amount) returns()
func (_Bridge *BridgeSession) AdminCollectFees(amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.AdminCollectFees(&_Bridge.TransactOpts, amount)
}

// AdminCollectFees is a paid mutator transaction binding the contract method 0xe77b33fd.
//
// Solidity: function adminCollectFees(uint256 amount) returns()
func (_Bridge *BridgeTransactorSession) AdminCollectFees(amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.AdminCollectFees(&_Bridge.TransactOpts, amount)
}

// AdminPauseTransfers is a paid mutator transaction binding the contract method 0x80ae1c28.
//
// Solidity: function adminPauseTransfers() returns()
func (_Bridge *BridgeTransactor) AdminPauseTransfers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "adminPauseTransfers")
}

// AdminPauseTransfers is a paid mutator transaction binding the contract method 0x80ae1c28.
//
// Solidity: function adminPauseTransfers() returns()
func (_Bridge *BridgeSession) AdminPauseTransfers() (*types.Transaction, error) {
	return _Bridge.Contract.AdminPauseTransfers(&_Bridge.TransactOpts)
}

// AdminPauseTransfers is a paid mutator transaction binding the contract method 0x80ae1c28.
//
// Solidity: function adminPauseTransfers() returns()
func (_Bridge *BridgeTransactorSession) AdminPauseTransfers() (*types.Transaction, error) {
	return _Bridge.Contract.AdminPauseTransfers(&_Bridge.TransactOpts)
}

// AdminSetBurnable is a paid mutator transaction binding the contract method 0x8c0c2631.
//
// Solidity: function adminSetBurnable(address handlerAddress, address tokenAddress) returns()
func (_Bridge *BridgeTransactor) AdminSetBurnable(opts *bind.TransactOpts, handlerAddress common.Address, tokenAddress common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "adminSetBurnable", handlerAddress, tokenAddress)
}

// AdminSetBurnable is a paid mutator transaction binding the contract method 0x8c0c2631.
//
// Solidity: function adminSetBurnable(address handlerAddress, address tokenAddress) returns()
func (_Bridge *BridgeSession) AdminSetBurnable(handlerAddress common.Address, tokenAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AdminSetBurnable(&_Bridge.TransactOpts, handlerAddress, tokenAddress)
}

// AdminSetBurnable is a paid mutator transaction binding the contract method 0x8c0c2631.
//
// Solidity: function adminSetBurnable(address handlerAddress, address tokenAddress) returns()
func (_Bridge *BridgeTransactorSession) AdminSetBurnable(handlerAddress common.Address, tokenAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AdminSetBurnable(&_Bridge.TransactOpts, handlerAddress, tokenAddress)
}

// AdminSetRelayerHub is a paid mutator transaction binding the contract method 0x5d699c5d.
//
// Solidity: function adminSetRelayerHub(address newRelayerHub) returns()
func (_Bridge *BridgeTransactor) AdminSetRelayerHub(opts *bind.TransactOpts, newRelayerHub common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "adminSetRelayerHub", newRelayerHub)
}

// AdminSetRelayerHub is a paid mutator transaction binding the contract method 0x5d699c5d.
//
// Solidity: function adminSetRelayerHub(address newRelayerHub) returns()
func (_Bridge *BridgeSession) AdminSetRelayerHub(newRelayerHub common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AdminSetRelayerHub(&_Bridge.TransactOpts, newRelayerHub)
}

// AdminSetRelayerHub is a paid mutator transaction binding the contract method 0x5d699c5d.
//
// Solidity: function adminSetRelayerHub(address newRelayerHub) returns()
func (_Bridge *BridgeTransactorSession) AdminSetRelayerHub(newRelayerHub common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AdminSetRelayerHub(&_Bridge.TransactOpts, newRelayerHub)
}

// AdminSetResource is a paid mutator transaction binding the contract method 0xcb10f215.
//
// Solidity: function adminSetResource(address handlerAddress, bytes32 resourceID, address tokenAddress) returns()
func (_Bridge *BridgeTransactor) AdminSetResource(opts *bind.TransactOpts, handlerAddress common.Address, resourceID [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "adminSetResource", handlerAddress, resourceID, tokenAddress)
}

// AdminSetResource is a paid mutator transaction binding the contract method 0xcb10f215.
//
// Solidity: function adminSetResource(address handlerAddress, bytes32 resourceID, address tokenAddress) returns()
func (_Bridge *BridgeSession) AdminSetResource(handlerAddress common.Address, resourceID [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AdminSetResource(&_Bridge.TransactOpts, handlerAddress, resourceID, tokenAddress)
}

// AdminSetResource is a paid mutator transaction binding the contract method 0xcb10f215.
//
// Solidity: function adminSetResource(address handlerAddress, bytes32 resourceID, address tokenAddress) returns()
func (_Bridge *BridgeTransactorSession) AdminSetResource(handlerAddress common.Address, resourceID [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AdminSetResource(&_Bridge.TransactOpts, handlerAddress, resourceID, tokenAddress)
}

// AdminUnpauseTransfers is a paid mutator transaction binding the contract method 0xffaac0eb.
//
// Solidity: function adminUnpauseTransfers() returns()
func (_Bridge *BridgeTransactor) AdminUnpauseTransfers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "adminUnpauseTransfers")
}

// AdminUnpauseTransfers is a paid mutator transaction binding the contract method 0xffaac0eb.
//
// Solidity: function adminUnpauseTransfers() returns()
func (_Bridge *BridgeSession) AdminUnpauseTransfers() (*types.Transaction, error) {
	return _Bridge.Contract.AdminUnpauseTransfers(&_Bridge.TransactOpts)
}

// AdminUnpauseTransfers is a paid mutator transaction binding the contract method 0xffaac0eb.
//
// Solidity: function adminUnpauseTransfers() returns()
func (_Bridge *BridgeTransactorSession) AdminUnpauseTransfers() (*types.Transaction, error) {
	return _Bridge.Contract.AdminUnpauseTransfers(&_Bridge.TransactOpts)
}

// AdminWithdraw is a paid mutator transaction binding the contract method 0x780cf004.
//
// Solidity: function adminWithdraw(address handlerAddress, address tokenAddress, address recipient, uint256 amountOrTokenID) returns()
func (_Bridge *BridgeTransactor) AdminWithdraw(opts *bind.TransactOpts, handlerAddress common.Address, tokenAddress common.Address, recipient common.Address, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "adminWithdraw", handlerAddress, tokenAddress, recipient, amountOrTokenID)
}

// AdminWithdraw is a paid mutator transaction binding the contract method 0x780cf004.
//
// Solidity: function adminWithdraw(address handlerAddress, address tokenAddress, address recipient, uint256 amountOrTokenID) returns()
func (_Bridge *BridgeSession) AdminWithdraw(handlerAddress common.Address, tokenAddress common.Address, recipient common.Address, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.AdminWithdraw(&_Bridge.TransactOpts, handlerAddress, tokenAddress, recipient, amountOrTokenID)
}

// AdminWithdraw is a paid mutator transaction binding the contract method 0x780cf004.
//
// Solidity: function adminWithdraw(address handlerAddress, address tokenAddress, address recipient, uint256 amountOrTokenID) returns()
func (_Bridge *BridgeTransactorSession) AdminWithdraw(handlerAddress common.Address, tokenAddress common.Address, recipient common.Address, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.AdminWithdraw(&_Bridge.TransactOpts, handlerAddress, tokenAddress, recipient, amountOrTokenID)
}

// CancelProposal is a paid mutator transaction binding the contract method 0x68872644.
//
// Solidity: function cancelProposal(bytes8 chainID, uint64 depositNonce, address recipientAddress, uint256 amount, bytes32 resourceID) returns()
func (_Bridge *BridgeTransactor) CancelProposal(opts *bind.TransactOpts, chainID [8]byte, depositNonce uint64, recipientAddress common.Address, amount *big.Int, resourceID [32]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "cancelProposal", chainID, depositNonce, recipientAddress, amount, resourceID)
}

// CancelProposal is a paid mutator transaction binding the contract method 0x68872644.
//
// Solidity: function cancelProposal(bytes8 chainID, uint64 depositNonce, address recipientAddress, uint256 amount, bytes32 resourceID) returns()
func (_Bridge *BridgeSession) CancelProposal(chainID [8]byte, depositNonce uint64, recipientAddress common.Address, amount *big.Int, resourceID [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.CancelProposal(&_Bridge.TransactOpts, chainID, depositNonce, recipientAddress, amount, resourceID)
}

// CancelProposal is a paid mutator transaction binding the contract method 0x68872644.
//
// Solidity: function cancelProposal(bytes8 chainID, uint64 depositNonce, address recipientAddress, uint256 amount, bytes32 resourceID) returns()
func (_Bridge *BridgeTransactorSession) CancelProposal(chainID [8]byte, depositNonce uint64, recipientAddress common.Address, amount *big.Int, resourceID [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.CancelProposal(&_Bridge.TransactOpts, chainID, depositNonce, recipientAddress, amount, resourceID)
}

// Deposit is a paid mutator transaction binding the contract method 0x92b7064c.
//
// Solidity: function deposit(bytes8 destinationChainID, bytes32 resourceID, uint256 amount, address recipientAddress) payable returns()
func (_Bridge *BridgeTransactor) Deposit(opts *bind.TransactOpts, destinationChainID [8]byte, resourceID [32]byte, amount *big.Int, recipientAddress common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "deposit", destinationChainID, resourceID, amount, recipientAddress)
}

// Deposit is a paid mutator transaction binding the contract method 0x92b7064c.
//
// Solidity: function deposit(bytes8 destinationChainID, bytes32 resourceID, uint256 amount, address recipientAddress) payable returns()
func (_Bridge *BridgeSession) Deposit(destinationChainID [8]byte, resourceID [32]byte, amount *big.Int, recipientAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.Deposit(&_Bridge.TransactOpts, destinationChainID, resourceID, amount, recipientAddress)
}

// Deposit is a paid mutator transaction binding the contract method 0x92b7064c.
//
// Solidity: function deposit(bytes8 destinationChainID, bytes32 resourceID, uint256 amount, address recipientAddress) payable returns()
func (_Bridge *BridgeTransactorSession) Deposit(destinationChainID [8]byte, resourceID [32]byte, amount *big.Int, recipientAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.Deposit(&_Bridge.TransactOpts, destinationChainID, resourceID, amount, recipientAddress)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xb14ecafd.
//
// Solidity: function executeProposal(bytes8 chainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount) returns()
func (_Bridge *BridgeTransactor) ExecuteProposal(opts *bind.TransactOpts, chainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "executeProposal", chainID, depositNonce, resourceID, recipientAddress, amount)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xb14ecafd.
//
// Solidity: function executeProposal(bytes8 chainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount) returns()
func (_Bridge *BridgeSession) ExecuteProposal(chainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.ExecuteProposal(&_Bridge.TransactOpts, chainID, depositNonce, resourceID, recipientAddress, amount)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xb14ecafd.
//
// Solidity: function executeProposal(bytes8 chainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount) returns()
func (_Bridge *BridgeTransactorSession) ExecuteProposal(chainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.ExecuteProposal(&_Bridge.TransactOpts, chainID, depositNonce, resourceID, recipientAddress, amount)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Bridge *BridgeTransactor) Init(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "init")
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Bridge *BridgeSession) Init() (*types.Transaction, error) {
	return _Bridge.Contract.Init(&_Bridge.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Bridge *BridgeTransactorSession) Init() (*types.Transaction, error) {
	return _Bridge.Contract.Init(&_Bridge.TransactOpts)
}

// RelayerCollectReward is a paid mutator transaction binding the contract method 0xf408da91.
//
// Solidity: function relayerCollectReward() returns()
func (_Bridge *BridgeTransactor) RelayerCollectReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "relayerCollectReward")
}

// RelayerCollectReward is a paid mutator transaction binding the contract method 0xf408da91.
//
// Solidity: function relayerCollectReward() returns()
func (_Bridge *BridgeSession) RelayerCollectReward() (*types.Transaction, error) {
	return _Bridge.Contract.RelayerCollectReward(&_Bridge.TransactOpts)
}

// RelayerCollectReward is a paid mutator transaction binding the contract method 0xf408da91.
//
// Solidity: function relayerCollectReward() returns()
func (_Bridge *BridgeTransactorSession) RelayerCollectReward() (*types.Transaction, error) {
	return _Bridge.Contract.RelayerCollectReward(&_Bridge.TransactOpts)
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

// VoteProposal is a paid mutator transaction binding the contract method 0xb3fc7eae.
//
// Solidity: function voteProposal(bytes8 chainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount) returns()
func (_Bridge *BridgeTransactor) VoteProposal(opts *bind.TransactOpts, chainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "voteProposal", chainID, depositNonce, resourceID, recipientAddress, amount)
}

// VoteProposal is a paid mutator transaction binding the contract method 0xb3fc7eae.
//
// Solidity: function voteProposal(bytes8 chainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount) returns()
func (_Bridge *BridgeSession) VoteProposal(chainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.VoteProposal(&_Bridge.TransactOpts, chainID, depositNonce, resourceID, recipientAddress, amount)
}

// VoteProposal is a paid mutator transaction binding the contract method 0xb3fc7eae.
//
// Solidity: function voteProposal(bytes8 chainID, uint64 depositNonce, bytes32 resourceID, address recipientAddress, uint256 amount) returns()
func (_Bridge *BridgeTransactorSession) VoteProposal(chainID [8]byte, depositNonce uint64, resourceID [32]byte, recipientAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.VoteProposal(&_Bridge.TransactOpts, chainID, depositNonce, resourceID, recipientAddress, amount)
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
	DestinationChainID [8]byte
	ResourceID         [32]byte
	DepositNonce       uint64
	Depositor          common.Address
	RecipientAddress   common.Address
	TokenAddress       common.Address
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x6a8f7ffca1e3ab54554c709a9eb23e11fc5063f0b24e40087c40ab3e7027f189.
//
// Solidity: event Deposit(bytes8 destinationChainID, bytes32 resourceID, uint64 depositNonce, address depositor, address recipientAddress, address tokenAddress, uint256 amount)
func (_Bridge *BridgeFilterer) FilterDeposit(opts *bind.FilterOpts) (*BridgeDepositIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &BridgeDepositIterator{contract: _Bridge.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x6a8f7ffca1e3ab54554c709a9eb23e11fc5063f0b24e40087c40ab3e7027f189.
//
// Solidity: event Deposit(bytes8 destinationChainID, bytes32 resourceID, uint64 depositNonce, address depositor, address recipientAddress, address tokenAddress, uint256 amount)
func (_Bridge *BridgeFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *BridgeDeposit) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Deposit")
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

// ParseDeposit is a log parse operation binding the contract event 0x6a8f7ffca1e3ab54554c709a9eb23e11fc5063f0b24e40087c40ab3e7027f189.
//
// Solidity: event Deposit(bytes8 destinationChainID, bytes32 resourceID, uint64 depositNonce, address depositor, address recipientAddress, address tokenAddress, uint256 amount)
func (_Bridge *BridgeFilterer) ParseDeposit(log types.Log) (*BridgeDeposit, error) {
	event := new(BridgeDeposit)
	if err := _Bridge.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// BridgePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Bridge contract.
type BridgePausedIterator struct {
	Event *BridgePaused // Event containing the contract specifics and raw log

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
func (it *BridgePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgePaused)
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
		it.Event = new(BridgePaused)
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
func (it *BridgePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgePaused represents a Paused event raised by the Bridge contract.
type BridgePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bridge *BridgeFilterer) FilterPaused(opts *bind.FilterOpts) (*BridgePausedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BridgePausedIterator{contract: _Bridge.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bridge *BridgeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BridgePaused) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgePaused)
				if err := _Bridge.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bridge *BridgeFilterer) ParsePaused(log types.Log) (*BridgePaused, error) {
	event := new(BridgePaused)
	if err := _Bridge.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeProposalEventIterator is returned from FilterProposalEvent and is used to iterate over the raw logs and unpacked data for ProposalEvent events raised by the Bridge contract.
type BridgeProposalEventIterator struct {
	Event *BridgeProposalEvent // Event containing the contract specifics and raw log

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
func (it *BridgeProposalEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeProposalEvent)
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
		it.Event = new(BridgeProposalEvent)
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
func (it *BridgeProposalEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeProposalEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeProposalEvent represents a ProposalEvent event raised by the Bridge contract.
type BridgeProposalEvent struct {
	OriginChainID [8]byte
	DepositNonce  uint64
	Status        uint8
	ResourceID    [32]byte
	DataHash      [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterProposalEvent is a free log retrieval operation binding the contract event 0x8dc49847a011c3b316cd0f50cf982e0fd5b3ddb7fdf970fc81a25557f0923a73.
//
// Solidity: event ProposalEvent(bytes8 originChainID, uint64 depositNonce, uint8 status, bytes32 resourceID, bytes32 dataHash)
func (_Bridge *BridgeFilterer) FilterProposalEvent(opts *bind.FilterOpts) (*BridgeProposalEventIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "ProposalEvent")
	if err != nil {
		return nil, err
	}
	return &BridgeProposalEventIterator{contract: _Bridge.contract, event: "ProposalEvent", logs: logs, sub: sub}, nil
}

// WatchProposalEvent is a free log subscription operation binding the contract event 0x8dc49847a011c3b316cd0f50cf982e0fd5b3ddb7fdf970fc81a25557f0923a73.
//
// Solidity: event ProposalEvent(bytes8 originChainID, uint64 depositNonce, uint8 status, bytes32 resourceID, bytes32 dataHash)
func (_Bridge *BridgeFilterer) WatchProposalEvent(opts *bind.WatchOpts, sink chan<- *BridgeProposalEvent) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "ProposalEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeProposalEvent)
				if err := _Bridge.contract.UnpackLog(event, "ProposalEvent", log); err != nil {
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

// ParseProposalEvent is a log parse operation binding the contract event 0x8dc49847a011c3b316cd0f50cf982e0fd5b3ddb7fdf970fc81a25557f0923a73.
//
// Solidity: event ProposalEvent(bytes8 originChainID, uint64 depositNonce, uint8 status, bytes32 resourceID, bytes32 dataHash)
func (_Bridge *BridgeFilterer) ParseProposalEvent(log types.Log) (*BridgeProposalEvent, error) {
	event := new(BridgeProposalEvent)
	if err := _Bridge.contract.UnpackLog(event, "ProposalEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeProposalVoteIterator is returned from FilterProposalVote and is used to iterate over the raw logs and unpacked data for ProposalVote events raised by the Bridge contract.
type BridgeProposalVoteIterator struct {
	Event *BridgeProposalVote // Event containing the contract specifics and raw log

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
func (it *BridgeProposalVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeProposalVote)
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
		it.Event = new(BridgeProposalVote)
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
func (it *BridgeProposalVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeProposalVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeProposalVote represents a ProposalVote event raised by the Bridge contract.
type BridgeProposalVote struct {
	OriginChainID [8]byte
	DepositNonce  uint64
	Status        uint8
	ResourceID    [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterProposalVote is a free log retrieval operation binding the contract event 0x85f41114efc645854a10eef33ef4dec54341cb3ec3ab32386c92c881f3b1b505.
//
// Solidity: event ProposalVote(bytes8 originChainID, uint64 depositNonce, uint8 status, bytes32 resourceID)
func (_Bridge *BridgeFilterer) FilterProposalVote(opts *bind.FilterOpts) (*BridgeProposalVoteIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "ProposalVote")
	if err != nil {
		return nil, err
	}
	return &BridgeProposalVoteIterator{contract: _Bridge.contract, event: "ProposalVote", logs: logs, sub: sub}, nil
}

// WatchProposalVote is a free log subscription operation binding the contract event 0x85f41114efc645854a10eef33ef4dec54341cb3ec3ab32386c92c881f3b1b505.
//
// Solidity: event ProposalVote(bytes8 originChainID, uint64 depositNonce, uint8 status, bytes32 resourceID)
func (_Bridge *BridgeFilterer) WatchProposalVote(opts *bind.WatchOpts, sink chan<- *BridgeProposalVote) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "ProposalVote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeProposalVote)
				if err := _Bridge.contract.UnpackLog(event, "ProposalVote", log); err != nil {
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

// ParseProposalVote is a log parse operation binding the contract event 0x85f41114efc645854a10eef33ef4dec54341cb3ec3ab32386c92c881f3b1b505.
//
// Solidity: event ProposalVote(bytes8 originChainID, uint64 depositNonce, uint8 status, bytes32 resourceID)
func (_Bridge *BridgeFilterer) ParseProposalVote(log types.Log) (*BridgeProposalVote, error) {
	event := new(BridgeProposalVote)
	if err := _Bridge.contract.UnpackLog(event, "ProposalVote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRelayerThresholdChangedIterator is returned from FilterRelayerThresholdChanged and is used to iterate over the raw logs and unpacked data for RelayerThresholdChanged events raised by the Bridge contract.
type BridgeRelayerThresholdChangedIterator struct {
	Event *BridgeRelayerThresholdChanged // Event containing the contract specifics and raw log

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
func (it *BridgeRelayerThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRelayerThresholdChanged)
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
		it.Event = new(BridgeRelayerThresholdChanged)
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
func (it *BridgeRelayerThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRelayerThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRelayerThresholdChanged represents a RelayerThresholdChanged event raised by the Bridge contract.
type BridgeRelayerThresholdChanged struct {
	NewThreshold *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRelayerThresholdChanged is a free log retrieval operation binding the contract event 0xa20d6b84cd798a24038be305eff8a45ca82ef54a2aa2082005d8e14c0a4746c8.
//
// Solidity: event RelayerThresholdChanged(uint256 newThreshold)
func (_Bridge *BridgeFilterer) FilterRelayerThresholdChanged(opts *bind.FilterOpts) (*BridgeRelayerThresholdChangedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RelayerThresholdChanged")
	if err != nil {
		return nil, err
	}
	return &BridgeRelayerThresholdChangedIterator{contract: _Bridge.contract, event: "RelayerThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchRelayerThresholdChanged is a free log subscription operation binding the contract event 0xa20d6b84cd798a24038be305eff8a45ca82ef54a2aa2082005d8e14c0a4746c8.
//
// Solidity: event RelayerThresholdChanged(uint256 newThreshold)
func (_Bridge *BridgeFilterer) WatchRelayerThresholdChanged(opts *bind.WatchOpts, sink chan<- *BridgeRelayerThresholdChanged) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RelayerThresholdChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRelayerThresholdChanged)
				if err := _Bridge.contract.UnpackLog(event, "RelayerThresholdChanged", log); err != nil {
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

// ParseRelayerThresholdChanged is a log parse operation binding the contract event 0xa20d6b84cd798a24038be305eff8a45ca82ef54a2aa2082005d8e14c0a4746c8.
//
// Solidity: event RelayerThresholdChanged(uint256 newThreshold)
func (_Bridge *BridgeFilterer) ParseRelayerThresholdChanged(log types.Log) (*BridgeRelayerThresholdChanged, error) {
	event := new(BridgeRelayerThresholdChanged)
	if err := _Bridge.contract.UnpackLog(event, "RelayerThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRewardCollectedIterator is returned from FilterRewardCollected and is used to iterate over the raw logs and unpacked data for RewardCollected events raised by the Bridge contract.
type BridgeRewardCollectedIterator struct {
	Event *BridgeRewardCollected // Event containing the contract specifics and raw log

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
func (it *BridgeRewardCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRewardCollected)
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
		it.Event = new(BridgeRewardCollected)
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
func (it *BridgeRewardCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRewardCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRewardCollected represents a RewardCollected event raised by the Bridge contract.
type BridgeRewardCollected struct {
	Relayer common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRewardCollected is a free log retrieval operation binding the contract event 0xe8354b169cd993d5cdfad1036a9a3f1ea7ed77e430bccb279200fd088243f595.
//
// Solidity: event RewardCollected(address relayer, uint256 amount)
func (_Bridge *BridgeFilterer) FilterRewardCollected(opts *bind.FilterOpts) (*BridgeRewardCollectedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RewardCollected")
	if err != nil {
		return nil, err
	}
	return &BridgeRewardCollectedIterator{contract: _Bridge.contract, event: "RewardCollected", logs: logs, sub: sub}, nil
}

// WatchRewardCollected is a free log subscription operation binding the contract event 0xe8354b169cd993d5cdfad1036a9a3f1ea7ed77e430bccb279200fd088243f595.
//
// Solidity: event RewardCollected(address relayer, uint256 amount)
func (_Bridge *BridgeFilterer) WatchRewardCollected(opts *bind.WatchOpts, sink chan<- *BridgeRewardCollected) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RewardCollected")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRewardCollected)
				if err := _Bridge.contract.UnpackLog(event, "RewardCollected", log); err != nil {
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

// ParseRewardCollected is a log parse operation binding the contract event 0xe8354b169cd993d5cdfad1036a9a3f1ea7ed77e430bccb279200fd088243f595.
//
// Solidity: event RewardCollected(address relayer, uint256 amount)
func (_Bridge *BridgeFilterer) ParseRewardCollected(log types.Log) (*BridgeRewardCollected, error) {
	event := new(BridgeRewardCollected)
	if err := _Bridge.contract.UnpackLog(event, "RewardCollected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Bridge contract.
type BridgeUnpausedIterator struct {
	Event *BridgeUnpaused // Event containing the contract specifics and raw log

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
func (it *BridgeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeUnpaused)
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
		it.Event = new(BridgeUnpaused)
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
func (it *BridgeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeUnpaused represents a Unpaused event raised by the Bridge contract.
type BridgeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bridge *BridgeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BridgeUnpausedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BridgeUnpausedIterator{contract: _Bridge.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bridge *BridgeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BridgeUnpaused) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeUnpaused)
				if err := _Bridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bridge *BridgeFilterer) ParseUnpaused(log types.Log) (*BridgeUnpaused, error) {
	event := new(BridgeUnpaused)
	if err := _Bridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
