// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MinerReward

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/fgeth/fgeth"
	"github.com/fgeth/fgeth/accounts/abi"
	"github.com/fgeth/fgeth/accounts/abi/bind"
	"github.com/fgeth/fgeth/common"
	"github.com/fgeth/fgeth/core/types"
	"github.com/fgeth/fgeth/event"
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

// MinerRewardMetaData contains all meta data concerning the MinerReward contract.
var MinerRewardMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isMining\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"CreateMiner\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Addresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"Miners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isMining\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"}],\"name\":\"createMiner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMiners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"}],\"name\":\"stopMining\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604080518082019091526005808252640302e312e360dc1b602090920191825261003d91600091610043565b506100de565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061008457805160ff19168380011785556100b1565b828001600101855582156100b1579182015b828111156100b1578251825591602001919060010190610096565b506100bd9291506100c1565b5090565b6100db91905b808211156100bd57600081556001016100c7565b90565b6105a1806100ed6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80637501ce2d1161005b5780637501ce2d1461015f578063b021f567146101ad578063cad58c27146101e6578063f10708961461020e5761007d565b80630d8e6e2c146100825780631633da6e146100ff57806354fd4d5014610157575b600080fd5b61008a610234565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100c45781810151838201526020016100ac565b50505050905090810190601f1680156100f15780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101076102ca565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561014357818101518382015260200161012b565b505050509050019250505060405180910390f35b61008a61032b565b6101856004803603602081101561017557600080fd5b50356001600160a01b03166103b9565b604080516001600160a01b039094168452911515602084015282820152519081900360600190f35b6101ca600480360360208110156101c357600080fd5b50356103e8565b604080516001600160a01b039092168252519081900360200190f35b61020c600480360360208110156101fc57600080fd5b50356001600160a01b031661040f565b005b61020c6004803603602081101561022457600080fd5b50356001600160a01b0316610528565b60008054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156102c05780601f10610295576101008083540402835291602001916102c0565b820191906000526020600020905b8154815290600101906020018083116102a357829003601f168201915b5050505050905090565b606060028054806020026020016040519081016040528092919081815260200182805480156102c057602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610304575050505050905090565b6000805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103b15780601f10610386576101008083540402835291602001916103b1565b820191906000526020600020905b81548152906001019060200180831161039457829003601f168201915b505050505081565b600160208190526000918252604090912080549101546001600160a01b03821691600160a01b900460ff169083565b600281815481106103f557fe5b6000918252602090912001546001600160a01b0316905081565b61041761054c565b50604080516060810182526001600160a01b0383811680835260016020808501828152438688019081526000948552918390529583208551815497516001600160a01b031990981695169490941760ff60a01b1916600160a01b961515969096029590951783559351919093015590805b6002548110156104d157836001600160a01b0316600282815481106104a957fe5b6000918252602090912001546001600160a01b031614156104c957600191505b600101610488565b508061052357600280546001810182556000919091527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace0180546001600160a01b0319166001600160a01b0385161790555b505050565b6001600160a01b03166000908152600160205260409020805460ff60a01b19169055565b60408051606081018252600080825260208201819052918101919091529056fea265627a7a72315820b40ad665d4b906a8188906fb55da0a15d1d65872bbe8e9e678aa7d8a32973e6f64736f6c63430005100032",
}

// MinerRewardABI is the input ABI used to generate the binding from.
// Deprecated: Use MinerRewardMetaData.ABI instead.
var MinerRewardABI = MinerRewardMetaData.ABI

// MinerRewardBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MinerRewardMetaData.Bin instead.
var MinerRewardBin = MinerRewardMetaData.Bin

// DeployMinerReward deploys a new Ethereum contract, binding an instance of MinerReward to it.
func DeployMinerReward(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MinerReward, error) {
	parsed, err := MinerRewardMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MinerRewardBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MinerReward{MinerRewardCaller: MinerRewardCaller{contract: contract}, MinerRewardTransactor: MinerRewardTransactor{contract: contract}, MinerRewardFilterer: MinerRewardFilterer{contract: contract}}, nil
}

// MinerReward is an auto generated Go binding around an Ethereum contract.
type MinerReward struct {
	MinerRewardCaller     // Read-only binding to the contract
	MinerRewardTransactor // Write-only binding to the contract
	MinerRewardFilterer   // Log filterer for contract events
}

// MinerRewardCaller is an auto generated read-only Go binding around an Ethereum contract.
type MinerRewardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinerRewardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MinerRewardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinerRewardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MinerRewardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinerRewardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MinerRewardSession struct {
	Contract     *MinerReward      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MinerRewardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MinerRewardCallerSession struct {
	Contract *MinerRewardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MinerRewardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MinerRewardTransactorSession struct {
	Contract     *MinerRewardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MinerRewardRaw is an auto generated low-level Go binding around an Ethereum contract.
type MinerRewardRaw struct {
	Contract *MinerReward // Generic contract binding to access the raw methods on
}

// MinerRewardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MinerRewardCallerRaw struct {
	Contract *MinerRewardCaller // Generic read-only contract binding to access the raw methods on
}

// MinerRewardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MinerRewardTransactorRaw struct {
	Contract *MinerRewardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMinerReward creates a new instance of MinerReward, bound to a specific deployed contract.
func NewMinerReward(address common.Address, backend bind.ContractBackend) (*MinerReward, error) {
	contract, err := bindMinerReward(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MinerReward{MinerRewardCaller: MinerRewardCaller{contract: contract}, MinerRewardTransactor: MinerRewardTransactor{contract: contract}, MinerRewardFilterer: MinerRewardFilterer{contract: contract}}, nil
}

// NewMinerRewardCaller creates a new read-only instance of MinerReward, bound to a specific deployed contract.
func NewMinerRewardCaller(address common.Address, caller bind.ContractCaller) (*MinerRewardCaller, error) {
	contract, err := bindMinerReward(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MinerRewardCaller{contract: contract}, nil
}

// NewMinerRewardTransactor creates a new write-only instance of MinerReward, bound to a specific deployed contract.
func NewMinerRewardTransactor(address common.Address, transactor bind.ContractTransactor) (*MinerRewardTransactor, error) {
	contract, err := bindMinerReward(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MinerRewardTransactor{contract: contract}, nil
}

// NewMinerRewardFilterer creates a new log filterer instance of MinerReward, bound to a specific deployed contract.
func NewMinerRewardFilterer(address common.Address, filterer bind.ContractFilterer) (*MinerRewardFilterer, error) {
	contract, err := bindMinerReward(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MinerRewardFilterer{contract: contract}, nil
}

// bindMinerReward binds a generic wrapper to an already deployed contract.
func bindMinerReward(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MinerRewardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MinerReward *MinerRewardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MinerReward.Contract.MinerRewardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MinerReward *MinerRewardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinerReward.Contract.MinerRewardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MinerReward *MinerRewardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MinerReward.Contract.MinerRewardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MinerReward *MinerRewardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MinerReward.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MinerReward *MinerRewardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinerReward.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MinerReward *MinerRewardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MinerReward.Contract.contract.Transact(opts, method, params...)
}

// Addresses is a free data retrieval call binding the contract method 0xb021f567.
//
// Solidity: function Addresses(uint256 ) view returns(address)
func (_MinerReward *MinerRewardCaller) Addresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MinerReward.contract.Call(opts, &out, "Addresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addresses is a free data retrieval call binding the contract method 0xb021f567.
//
// Solidity: function Addresses(uint256 ) view returns(address)
func (_MinerReward *MinerRewardSession) Addresses(arg0 *big.Int) (common.Address, error) {
	return _MinerReward.Contract.Addresses(&_MinerReward.CallOpts, arg0)
}

// Addresses is a free data retrieval call binding the contract method 0xb021f567.
//
// Solidity: function Addresses(uint256 ) view returns(address)
func (_MinerReward *MinerRewardCallerSession) Addresses(arg0 *big.Int) (common.Address, error) {
	return _MinerReward.Contract.Addresses(&_MinerReward.CallOpts, arg0)
}

// Miners is a free data retrieval call binding the contract method 0x7501ce2d.
//
// Solidity: function Miners(address ) view returns(address coinbase, bool isMining, uint256 blockNumber)
func (_MinerReward *MinerRewardCaller) Miners(opts *bind.CallOpts, arg0 common.Address) (struct {
	Coinbase    common.Address
	IsMining    bool
	BlockNumber *big.Int
}, error) {
	var out []interface{}
	err := _MinerReward.contract.Call(opts, &out, "Miners", arg0)

	outstruct := new(struct {
		Coinbase    common.Address
		IsMining    bool
		BlockNumber *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Coinbase = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.IsMining = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.BlockNumber = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Miners is a free data retrieval call binding the contract method 0x7501ce2d.
//
// Solidity: function Miners(address ) view returns(address coinbase, bool isMining, uint256 blockNumber)
func (_MinerReward *MinerRewardSession) Miners(arg0 common.Address) (struct {
	Coinbase    common.Address
	IsMining    bool
	BlockNumber *big.Int
}, error) {
	return _MinerReward.Contract.Miners(&_MinerReward.CallOpts, arg0)
}

// Miners is a free data retrieval call binding the contract method 0x7501ce2d.
//
// Solidity: function Miners(address ) view returns(address coinbase, bool isMining, uint256 blockNumber)
func (_MinerReward *MinerRewardCallerSession) Miners(arg0 common.Address) (struct {
	Coinbase    common.Address
	IsMining    bool
	BlockNumber *big.Int
}, error) {
	return _MinerReward.Contract.Miners(&_MinerReward.CallOpts, arg0)
}

// GetMiners is a free data retrieval call binding the contract method 0x1633da6e.
//
// Solidity: function getMiners() view returns(address[])
func (_MinerReward *MinerRewardCaller) GetMiners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _MinerReward.contract.Call(opts, &out, "getMiners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetMiners is a free data retrieval call binding the contract method 0x1633da6e.
//
// Solidity: function getMiners() view returns(address[])
func (_MinerReward *MinerRewardSession) GetMiners() ([]common.Address, error) {
	return _MinerReward.Contract.GetMiners(&_MinerReward.CallOpts)
}

// GetMiners is a free data retrieval call binding the contract method 0x1633da6e.
//
// Solidity: function getMiners() view returns(address[])
func (_MinerReward *MinerRewardCallerSession) GetMiners() ([]common.Address, error) {
	return _MinerReward.Contract.GetMiners(&_MinerReward.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() view returns(string)
func (_MinerReward *MinerRewardCaller) GetVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MinerReward.contract.Call(opts, &out, "getVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() view returns(string)
func (_MinerReward *MinerRewardSession) GetVersion() (string, error) {
	return _MinerReward.Contract.GetVersion(&_MinerReward.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() view returns(string)
func (_MinerReward *MinerRewardCallerSession) GetVersion() (string, error) {
	return _MinerReward.Contract.GetVersion(&_MinerReward.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_MinerReward *MinerRewardCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MinerReward.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_MinerReward *MinerRewardSession) Version() (string, error) {
	return _MinerReward.Contract.Version(&_MinerReward.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_MinerReward *MinerRewardCallerSession) Version() (string, error) {
	return _MinerReward.Contract.Version(&_MinerReward.CallOpts)
}

// CreateMiner is a paid mutator transaction binding the contract method 0xcad58c27.
//
// Solidity: function createMiner(address coinbase) returns()
func (_MinerReward *MinerRewardTransactor) CreateMiner(opts *bind.TransactOpts, coinbase common.Address) (*types.Transaction, error) {
	return _MinerReward.contract.Transact(opts, "createMiner", coinbase)
}

// CreateMiner is a paid mutator transaction binding the contract method 0xcad58c27.
//
// Solidity: function createMiner(address coinbase) returns()
func (_MinerReward *MinerRewardSession) CreateMiner(coinbase common.Address) (*types.Transaction, error) {
	return _MinerReward.Contract.CreateMiner(&_MinerReward.TransactOpts, coinbase)
}

// CreateMiner is a paid mutator transaction binding the contract method 0xcad58c27.
//
// Solidity: function createMiner(address coinbase) returns()
func (_MinerReward *MinerRewardTransactorSession) CreateMiner(coinbase common.Address) (*types.Transaction, error) {
	return _MinerReward.Contract.CreateMiner(&_MinerReward.TransactOpts, coinbase)
}

// StopMining is a paid mutator transaction binding the contract method 0xf1070896.
//
// Solidity: function stopMining(address coinbase) returns()
func (_MinerReward *MinerRewardTransactor) StopMining(opts *bind.TransactOpts, coinbase common.Address) (*types.Transaction, error) {
	return _MinerReward.contract.Transact(opts, "stopMining", coinbase)
}

// StopMining is a paid mutator transaction binding the contract method 0xf1070896.
//
// Solidity: function stopMining(address coinbase) returns()
func (_MinerReward *MinerRewardSession) StopMining(coinbase common.Address) (*types.Transaction, error) {
	return _MinerReward.Contract.StopMining(&_MinerReward.TransactOpts, coinbase)
}

// StopMining is a paid mutator transaction binding the contract method 0xf1070896.
//
// Solidity: function stopMining(address coinbase) returns()
func (_MinerReward *MinerRewardTransactorSession) StopMining(coinbase common.Address) (*types.Transaction, error) {
	return _MinerReward.Contract.StopMining(&_MinerReward.TransactOpts, coinbase)
}

// MinerRewardCreateMinerIterator is returned from FilterCreateMiner and is used to iterate over the raw logs and unpacked data for CreateMiner events raised by the MinerReward contract.
type MinerRewardCreateMinerIterator struct {
	Event *MinerRewardCreateMiner // Event containing the contract specifics and raw log

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
func (it *MinerRewardCreateMinerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinerRewardCreateMiner)
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
		it.Event = new(MinerRewardCreateMiner)
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
func (it *MinerRewardCreateMinerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinerRewardCreateMinerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinerRewardCreateMiner represents a CreateMiner event raised by the MinerReward contract.
type MinerRewardCreateMiner struct {
	Coinbase    common.Address
	IsMining    bool
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCreateMiner is a free log retrieval operation binding the contract event 0xeeeaa9f61382d128196985a2305fb0670378bf80cda75b30a82ad503d79aa76a.
//
// Solidity: event CreateMiner(address coinbase, bool isMining, uint256 blockNumber)
func (_MinerReward *MinerRewardFilterer) FilterCreateMiner(opts *bind.FilterOpts) (*MinerRewardCreateMinerIterator, error) {

	logs, sub, err := _MinerReward.contract.FilterLogs(opts, "CreateMiner")
	if err != nil {
		return nil, err
	}
	return &MinerRewardCreateMinerIterator{contract: _MinerReward.contract, event: "CreateMiner", logs: logs, sub: sub}, nil
}

// WatchCreateMiner is a free log subscription operation binding the contract event 0xeeeaa9f61382d128196985a2305fb0670378bf80cda75b30a82ad503d79aa76a.
//
// Solidity: event CreateMiner(address coinbase, bool isMining, uint256 blockNumber)
func (_MinerReward *MinerRewardFilterer) WatchCreateMiner(opts *bind.WatchOpts, sink chan<- *MinerRewardCreateMiner) (event.Subscription, error) {

	logs, sub, err := _MinerReward.contract.WatchLogs(opts, "CreateMiner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinerRewardCreateMiner)
				if err := _MinerReward.contract.UnpackLog(event, "CreateMiner", log); err != nil {
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

// ParseCreateMiner is a log parse operation binding the contract event 0xeeeaa9f61382d128196985a2305fb0670378bf80cda75b30a82ad503d79aa76a.
//
// Solidity: event CreateMiner(address coinbase, bool isMining, uint256 blockNumber)
func (_MinerReward *MinerRewardFilterer) ParseCreateMiner(log types.Log) (*MinerRewardCreateMiner, error) {
	event := new(MinerRewardCreateMiner)
	if err := _MinerReward.contract.UnpackLog(event, "CreateMiner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
