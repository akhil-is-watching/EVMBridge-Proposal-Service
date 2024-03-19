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
	_ = abi.ConvertType
)

// BridgeMetaData contains all meta data concerning the Bridge contract.
var BridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposer_\",\"type\":\"address\"},{\"internalType\":\"contractIToken\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"voteThreshold_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recepient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"DepositEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"proposalID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recepient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ProposalEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"proposalID\",\"type\":\"bytes32\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"proposalID\",\"type\":\"bytes32\"}],\"name\":\"ProposalVote\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"addChainID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"addValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recepient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proposalNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"getProposalID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proposalNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"propose\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposer_\",\"type\":\"address\"}],\"name\":\"setProposer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"proposalID\",\"type\":\"bytes32\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voteThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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
	parsed, err := BridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_Bridge *BridgeCaller) ChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "chainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_Bridge *BridgeSession) ChainId() (*big.Int, error) {
	return _Bridge.Contract.ChainId(&_Bridge.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_Bridge *BridgeCallerSession) ChainId() (*big.Int, error) {
	return _Bridge.Contract.ChainId(&_Bridge.CallOpts)
}

// GetProposalID is a free data retrieval call binding the contract method 0x5d38bfb2.
//
// Solidity: function getProposalID(bytes32 txHash, address receiver, uint256 amount, uint256 proposalNonce, uint256 chainID) pure returns(bytes32)
func (_Bridge *BridgeCaller) GetProposalID(opts *bind.CallOpts, txHash [32]byte, receiver common.Address, amount *big.Int, proposalNonce *big.Int, chainID *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getProposalID", txHash, receiver, amount, proposalNonce, chainID)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetProposalID is a free data retrieval call binding the contract method 0x5d38bfb2.
//
// Solidity: function getProposalID(bytes32 txHash, address receiver, uint256 amount, uint256 proposalNonce, uint256 chainID) pure returns(bytes32)
func (_Bridge *BridgeSession) GetProposalID(txHash [32]byte, receiver common.Address, amount *big.Int, proposalNonce *big.Int, chainID *big.Int) ([32]byte, error) {
	return _Bridge.Contract.GetProposalID(&_Bridge.CallOpts, txHash, receiver, amount, proposalNonce, chainID)
}

// GetProposalID is a free data retrieval call binding the contract method 0x5d38bfb2.
//
// Solidity: function getProposalID(bytes32 txHash, address receiver, uint256 amount, uint256 proposalNonce, uint256 chainID) pure returns(bytes32)
func (_Bridge *BridgeCallerSession) GetProposalID(txHash [32]byte, receiver common.Address, amount *big.Int, proposalNonce *big.Int, chainID *big.Int) ([32]byte, error) {
	return _Bridge.Contract.GetProposalID(&_Bridge.CallOpts, txHash, receiver, amount, proposalNonce, chainID)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address validator) view returns(bool)
func (_Bridge *BridgeCaller) IsValidator(opts *bind.CallOpts, validator common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "isValidator", validator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address validator) view returns(bool)
func (_Bridge *BridgeSession) IsValidator(validator common.Address) (bool, error) {
	return _Bridge.Contract.IsValidator(&_Bridge.CallOpts, validator)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address validator) view returns(bool)
func (_Bridge *BridgeCallerSession) IsValidator(validator common.Address) (bool, error) {
	return _Bridge.Contract.IsValidator(&_Bridge.CallOpts, validator)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Bridge *BridgeCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Bridge *BridgeSession) Nonce() (*big.Int, error) {
	return _Bridge.Contract.Nonce(&_Bridge.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Bridge *BridgeCallerSession) Nonce() (*big.Int, error) {
	return _Bridge.Contract.Nonce(&_Bridge.CallOpts)
}

// Proposer is a free data retrieval call binding the contract method 0xa8e4fb90.
//
// Solidity: function proposer() view returns(address)
func (_Bridge *BridgeCaller) Proposer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "proposer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Proposer is a free data retrieval call binding the contract method 0xa8e4fb90.
//
// Solidity: function proposer() view returns(address)
func (_Bridge *BridgeSession) Proposer() (common.Address, error) {
	return _Bridge.Contract.Proposer(&_Bridge.CallOpts)
}

// Proposer is a free data retrieval call binding the contract method 0xa8e4fb90.
//
// Solidity: function proposer() view returns(address)
func (_Bridge *BridgeCallerSession) Proposer() (common.Address, error) {
	return _Bridge.Contract.Proposer(&_Bridge.CallOpts)
}

// VoteThreshold is a free data retrieval call binding the contract method 0x4fe437d5.
//
// Solidity: function voteThreshold() view returns(uint256)
func (_Bridge *BridgeCaller) VoteThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "voteThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VoteThreshold is a free data retrieval call binding the contract method 0x4fe437d5.
//
// Solidity: function voteThreshold() view returns(uint256)
func (_Bridge *BridgeSession) VoteThreshold() (*big.Int, error) {
	return _Bridge.Contract.VoteThreshold(&_Bridge.CallOpts)
}

// VoteThreshold is a free data retrieval call binding the contract method 0x4fe437d5.
//
// Solidity: function voteThreshold() view returns(uint256)
func (_Bridge *BridgeCallerSession) VoteThreshold() (*big.Int, error) {
	return _Bridge.Contract.VoteThreshold(&_Bridge.CallOpts)
}

// AddChainID is a paid mutator transaction binding the contract method 0x20c29bca.
//
// Solidity: function addChainID(uint256 chainID) returns()
func (_Bridge *BridgeTransactor) AddChainID(opts *bind.TransactOpts, chainID *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "addChainID", chainID)
}

// AddChainID is a paid mutator transaction binding the contract method 0x20c29bca.
//
// Solidity: function addChainID(uint256 chainID) returns()
func (_Bridge *BridgeSession) AddChainID(chainID *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.AddChainID(&_Bridge.TransactOpts, chainID)
}

// AddChainID is a paid mutator transaction binding the contract method 0x20c29bca.
//
// Solidity: function addChainID(uint256 chainID) returns()
func (_Bridge *BridgeTransactorSession) AddChainID(chainID *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.AddChainID(&_Bridge.TransactOpts, chainID)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address validator) returns()
func (_Bridge *BridgeTransactor) AddValidator(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "addValidator", validator)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address validator) returns()
func (_Bridge *BridgeSession) AddValidator(validator common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AddValidator(&_Bridge.TransactOpts, validator)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address validator) returns()
func (_Bridge *BridgeTransactorSession) AddValidator(validator common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AddValidator(&_Bridge.TransactOpts, validator)
}

// Deposit is a paid mutator transaction binding the contract method 0x0efe6a8b.
//
// Solidity: function deposit(address recepient, uint256 amount, uint256 chainID) returns()
func (_Bridge *BridgeTransactor) Deposit(opts *bind.TransactOpts, recepient common.Address, amount *big.Int, chainID *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "deposit", recepient, amount, chainID)
}

// Deposit is a paid mutator transaction binding the contract method 0x0efe6a8b.
//
// Solidity: function deposit(address recepient, uint256 amount, uint256 chainID) returns()
func (_Bridge *BridgeSession) Deposit(recepient common.Address, amount *big.Int, chainID *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Deposit(&_Bridge.TransactOpts, recepient, amount, chainID)
}

// Deposit is a paid mutator transaction binding the contract method 0x0efe6a8b.
//
// Solidity: function deposit(address recepient, uint256 amount, uint256 chainID) returns()
func (_Bridge *BridgeTransactorSession) Deposit(recepient common.Address, amount *big.Int, chainID *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Deposit(&_Bridge.TransactOpts, recepient, amount, chainID)
}

// Propose is a paid mutator transaction binding the contract method 0xbf1727cc.
//
// Solidity: function propose(bytes32 txHash, address receiver, uint256 amount, uint256 proposalNonce, uint256 chainID) returns()
func (_Bridge *BridgeTransactor) Propose(opts *bind.TransactOpts, txHash [32]byte, receiver common.Address, amount *big.Int, proposalNonce *big.Int, chainID *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "propose", txHash, receiver, amount, proposalNonce, chainID)
}

// Propose is a paid mutator transaction binding the contract method 0xbf1727cc.
//
// Solidity: function propose(bytes32 txHash, address receiver, uint256 amount, uint256 proposalNonce, uint256 chainID) returns()
func (_Bridge *BridgeSession) Propose(txHash [32]byte, receiver common.Address, amount *big.Int, proposalNonce *big.Int, chainID *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Propose(&_Bridge.TransactOpts, txHash, receiver, amount, proposalNonce, chainID)
}

// Propose is a paid mutator transaction binding the contract method 0xbf1727cc.
//
// Solidity: function propose(bytes32 txHash, address receiver, uint256 amount, uint256 proposalNonce, uint256 chainID) returns()
func (_Bridge *BridgeTransactorSession) Propose(txHash [32]byte, receiver common.Address, amount *big.Int, proposalNonce *big.Int, chainID *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Propose(&_Bridge.TransactOpts, txHash, receiver, amount, proposalNonce, chainID)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address validator) returns()
func (_Bridge *BridgeTransactor) RemoveValidator(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "removeValidator", validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address validator) returns()
func (_Bridge *BridgeSession) RemoveValidator(validator common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RemoveValidator(&_Bridge.TransactOpts, validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address validator) returns()
func (_Bridge *BridgeTransactorSession) RemoveValidator(validator common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RemoveValidator(&_Bridge.TransactOpts, validator)
}

// SetProposer is a paid mutator transaction binding the contract method 0x1fb4a228.
//
// Solidity: function setProposer(address proposer_) returns()
func (_Bridge *BridgeTransactor) SetProposer(opts *bind.TransactOpts, proposer_ common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setProposer", proposer_)
}

// SetProposer is a paid mutator transaction binding the contract method 0x1fb4a228.
//
// Solidity: function setProposer(address proposer_) returns()
func (_Bridge *BridgeSession) SetProposer(proposer_ common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetProposer(&_Bridge.TransactOpts, proposer_)
}

// SetProposer is a paid mutator transaction binding the contract method 0x1fb4a228.
//
// Solidity: function setProposer(address proposer_) returns()
func (_Bridge *BridgeTransactorSession) SetProposer(proposer_ common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetProposer(&_Bridge.TransactOpts, proposer_)
}

// Vote is a paid mutator transaction binding the contract method 0xa69beaba.
//
// Solidity: function vote(bytes32 proposalID) returns()
func (_Bridge *BridgeTransactor) Vote(opts *bind.TransactOpts, proposalID [32]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "vote", proposalID)
}

// Vote is a paid mutator transaction binding the contract method 0xa69beaba.
//
// Solidity: function vote(bytes32 proposalID) returns()
func (_Bridge *BridgeSession) Vote(proposalID [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.Vote(&_Bridge.TransactOpts, proposalID)
}

// Vote is a paid mutator transaction binding the contract method 0xa69beaba.
//
// Solidity: function vote(bytes32 proposalID) returns()
func (_Bridge *BridgeTransactorSession) Vote(proposalID [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.Vote(&_Bridge.TransactOpts, proposalID)
}

// BridgeDepositEventIterator is returned from FilterDepositEvent and is used to iterate over the raw logs and unpacked data for DepositEvent events raised by the Bridge contract.
type BridgeDepositEventIterator struct {
	Event *BridgeDepositEvent // Event containing the contract specifics and raw log

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
func (it *BridgeDepositEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeDepositEvent)
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
		it.Event = new(BridgeDepositEvent)
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
func (it *BridgeDepositEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeDepositEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeDepositEvent represents a DepositEvent event raised by the Bridge contract.
type BridgeDepositEvent struct {
	Sender    common.Address
	Recepient common.Address
	Amount    *big.Int
	Nonce     *big.Int
	ChainID   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDepositEvent is a free log retrieval operation binding the contract event 0xf9e6944a20b1cd07f8a1cbb67e59d0db44c9eff947e6b113b0f571c70e6be673.
//
// Solidity: event DepositEvent(address sender, address recepient, uint256 amount, uint256 nonce, uint256 chainID)
func (_Bridge *BridgeFilterer) FilterDepositEvent(opts *bind.FilterOpts) (*BridgeDepositEventIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "DepositEvent")
	if err != nil {
		return nil, err
	}
	return &BridgeDepositEventIterator{contract: _Bridge.contract, event: "DepositEvent", logs: logs, sub: sub}, nil
}

// WatchDepositEvent is a free log subscription operation binding the contract event 0xf9e6944a20b1cd07f8a1cbb67e59d0db44c9eff947e6b113b0f571c70e6be673.
//
// Solidity: event DepositEvent(address sender, address recepient, uint256 amount, uint256 nonce, uint256 chainID)
func (_Bridge *BridgeFilterer) WatchDepositEvent(opts *bind.WatchOpts, sink chan<- *BridgeDepositEvent) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "DepositEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeDepositEvent)
				if err := _Bridge.contract.UnpackLog(event, "DepositEvent", log); err != nil {
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

// ParseDepositEvent is a log parse operation binding the contract event 0xf9e6944a20b1cd07f8a1cbb67e59d0db44c9eff947e6b113b0f571c70e6be673.
//
// Solidity: event DepositEvent(address sender, address recepient, uint256 amount, uint256 nonce, uint256 chainID)
func (_Bridge *BridgeFilterer) ParseDepositEvent(log types.Log) (*BridgeDepositEvent, error) {
	event := new(BridgeDepositEvent)
	if err := _Bridge.contract.UnpackLog(event, "DepositEvent", log); err != nil {
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
	ProposalID [32]byte
	Recepient  common.Address
	TxHash     [32]byte
	ChainID    *big.Int
	Nonce      *big.Int
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalEvent is a free log retrieval operation binding the contract event 0x62819ce06a00ca8a28dd683b302b56c5d85bfc98bafc6d2969d60f4ba6ff3157.
//
// Solidity: event ProposalEvent(bytes32 proposalID, address recepient, bytes32 txHash, uint256 chainID, uint256 nonce, uint256 amount)
func (_Bridge *BridgeFilterer) FilterProposalEvent(opts *bind.FilterOpts) (*BridgeProposalEventIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "ProposalEvent")
	if err != nil {
		return nil, err
	}
	return &BridgeProposalEventIterator{contract: _Bridge.contract, event: "ProposalEvent", logs: logs, sub: sub}, nil
}

// WatchProposalEvent is a free log subscription operation binding the contract event 0x62819ce06a00ca8a28dd683b302b56c5d85bfc98bafc6d2969d60f4ba6ff3157.
//
// Solidity: event ProposalEvent(bytes32 proposalID, address recepient, bytes32 txHash, uint256 chainID, uint256 nonce, uint256 amount)
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

// ParseProposalEvent is a log parse operation binding the contract event 0x62819ce06a00ca8a28dd683b302b56c5d85bfc98bafc6d2969d60f4ba6ff3157.
//
// Solidity: event ProposalEvent(bytes32 proposalID, address recepient, bytes32 txHash, uint256 chainID, uint256 nonce, uint256 amount)
func (_Bridge *BridgeFilterer) ParseProposalEvent(log types.Log) (*BridgeProposalEvent, error) {
	event := new(BridgeProposalEvent)
	if err := _Bridge.contract.UnpackLog(event, "ProposalEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the Bridge contract.
type BridgeProposalExecutedIterator struct {
	Event *BridgeProposalExecuted // Event containing the contract specifics and raw log

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
func (it *BridgeProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeProposalExecuted)
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
		it.Event = new(BridgeProposalExecuted)
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
func (it *BridgeProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeProposalExecuted represents a ProposalExecuted event raised by the Bridge contract.
type BridgeProposalExecuted struct {
	ProposalID [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x7b1bcf1ccf901a11589afff5504d59fd0a53780eed2a952adade0348985139e0.
//
// Solidity: event ProposalExecuted(bytes32 proposalID)
func (_Bridge *BridgeFilterer) FilterProposalExecuted(opts *bind.FilterOpts) (*BridgeProposalExecutedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return &BridgeProposalExecutedIterator{contract: _Bridge.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x7b1bcf1ccf901a11589afff5504d59fd0a53780eed2a952adade0348985139e0.
//
// Solidity: event ProposalExecuted(bytes32 proposalID)
func (_Bridge *BridgeFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *BridgeProposalExecuted) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeProposalExecuted)
				if err := _Bridge.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
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

// ParseProposalExecuted is a log parse operation binding the contract event 0x7b1bcf1ccf901a11589afff5504d59fd0a53780eed2a952adade0348985139e0.
//
// Solidity: event ProposalExecuted(bytes32 proposalID)
func (_Bridge *BridgeFilterer) ParseProposalExecuted(log types.Log) (*BridgeProposalExecuted, error) {
	event := new(BridgeProposalExecuted)
	if err := _Bridge.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
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
	ProposalID [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalVote is a free log retrieval operation binding the contract event 0xdc6efc31bac560ef02b2a63eef1102b7c6c02e43ac8fcf8d45cfd0de9d162be1.
//
// Solidity: event ProposalVote(bytes32 proposalID)
func (_Bridge *BridgeFilterer) FilterProposalVote(opts *bind.FilterOpts) (*BridgeProposalVoteIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "ProposalVote")
	if err != nil {
		return nil, err
	}
	return &BridgeProposalVoteIterator{contract: _Bridge.contract, event: "ProposalVote", logs: logs, sub: sub}, nil
}

// WatchProposalVote is a free log subscription operation binding the contract event 0xdc6efc31bac560ef02b2a63eef1102b7c6c02e43ac8fcf8d45cfd0de9d162be1.
//
// Solidity: event ProposalVote(bytes32 proposalID)
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

// ParseProposalVote is a log parse operation binding the contract event 0xdc6efc31bac560ef02b2a63eef1102b7c6c02e43ac8fcf8d45cfd0de9d162be1.
//
// Solidity: event ProposalVote(bytes32 proposalID)
func (_Bridge *BridgeFilterer) ParseProposalVote(log types.Log) (*BridgeProposalVote, error) {
	event := new(BridgeProposalVote)
	if err := _Bridge.contract.UnpackLog(event, "ProposalVote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
