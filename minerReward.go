// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MinerReward

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

// MinerRewardABI is the input ABI used to generate the binding from.
const MinerRewardABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isMining\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"CreateMiner\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Addresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"Miners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isMining\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"}],\"name\":\"createMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMiners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"}],\"name\":\"stopMining\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MinerRewardBin is the compiled bytecode used for deploying new contracts.
var MinerRewardBin = "0x608060405234801561001057600080fd5b50604080518082019091526005808252640302e312e360dc1b602090920191825261003d91600091610043565b50610117565b82805461004f906100dc565b90600052602060002090601f01602090048101928261007157600085556100b7565b82601f1061008a57805160ff19168380011785556100b7565b828001600101855582156100b7579182015b828111156100b757825182559160200191906001019061009c565b506100c39291506100c7565b5090565b5b808211156100c357600081556001016100c8565b600181811c908216806100f057607f821691505b6020821081141561011157634e487b7160e01b600052602260045260246000fd5b50919050565b6105ec806101266000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80637501ce2d1161005b5780637501ce2d146100bd578063b021f5671461011f578063cad58c271461014a578063f10708961461015f57600080fd5b80630d8e6e2c146100825780631633da6e146100a057806354fd4d50146100b5575b600080fd5b61008a610191565b60405161009791906104e7565b60405180910390f35b6100a8610223565b604051610097919061049a565b61008a610284565b6100fa6100cb366004610451565b600160208190526000918252604090912080549101546001600160a01b03821691600160a01b900460ff169083565b604080516001600160a01b039094168452911515602084015290820152606001610097565b61013261012d366004610481565b610312565b6040516001600160a01b039091168152602001610097565b61015d610158366004610451565b61033c565b005b61015d61016d366004610451565b6001600160a01b03166000908152600160205260409020805460ff60a01b19169055565b6060600080546101a09061053c565b80601f01602080910402602001604051908101604052809291908181526020018280546101cc9061053c565b80156102195780601f106101ee57610100808354040283529160200191610219565b820191906000526020600020905b8154815290600101906020018083116101fc57829003601f168201915b5050505050905090565b6060600280548060200260200160405190810160405280929190818152602001828054801561021957602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161025d575050505050905090565b600080546102919061053c565b80601f01602080910402602001604051908101604052809291908181526020018280546102bd9061053c565b801561030a5780601f106102df5761010080835404028352916020019161030a565b820191906000526020600020905b8154815290600101906020018083116102ed57829003601f168201915b505050505081565b6002818154811061032257600080fd5b6000918252602090912001546001600160a01b0316905081565b604080516060810182526001600160a01b0380841680835260016020808501828152438688019081526000948552918390529583208551815497511515600160a01b026001600160a81b03199098169516949094179590951783559351919093015590805b6002548110156103fa57836001600160a01b0316600282815481106103c8576103c86105a0565b6000918252602090912001546001600160a01b031614156103e857600191505b806103f281610577565b9150506103a1565b508061044c57600280546001810182556000919091527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace0180546001600160a01b0319166001600160a01b0385161790555b505050565b60006020828403121561046357600080fd5b81356001600160a01b038116811461047a57600080fd5b9392505050565b60006020828403121561049357600080fd5b5035919050565b6020808252825182820181905260009190848201906040850190845b818110156104db5783516001600160a01b0316835292840192918401916001016104b6565b50909695505050505050565b600060208083528351808285015260005b81811015610514578581018301518582016040015282016104f8565b81811115610526576000604083870101525b50601f01601f1916929092016040019392505050565b600181811c9082168061055057607f821691505b6020821081141561057157634e487b7160e01b600052602260045260246000fd5b50919050565b600060001982141561059957634e487b7160e01b600052601160045260246000fd5b5060010190565b634e487b7160e01b600052603260045260246000fdfea26469706673582212206b1a27ed492b9963f658b8dfcaa5c3450ad9191add61a3c944ad2e9b19c9bbf764736f6c63430008060033"

// DeployMinerReward deploys a new Ethereum contract, binding an instance of MinerReward to it.
func DeployMinerReward(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MinerReward, error) {
	parsed, err := abi.JSON(strings.NewReader(MinerRewardABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MinerRewardBin), backend)
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
