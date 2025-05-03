// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vrf

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

// DappLinkVRFManagerMetaData contains all meta data concerning the DappLinkVRFManager contract.
var DappLinkVRFManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"blsRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"dappLinkAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fulfillRandomWords\",\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_randomWords\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRequestStatus\",\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"fulfilled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"randomWords\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_dappLinkAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_blsRegistry\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lastRequestId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestIds\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requestMapping\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"fulfilled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requestRandomWords\",\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_numWords\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDappLink\",\"inputs\":[{\"name\":\"_dappLinkAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"FillRandomWords\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"randomWords\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RequestSent\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_numWords\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"current\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405234801561001057600080fd5b5061001961001e565b6100d0565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161561006e5760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100cd5780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b610a0e806100df6000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c806393e4b8801161008c578063d8a4676f11610066578063d8a4676f146101da578063f0c28a41146101fb578063f2fde38b1461020e578063fc2a88c31461022157600080fd5b806393e4b880146101a1578063996869d0146101b4578063c0c53b8b146101c757600080fd5b80631b739ef1146100d457806338ba4614146100e9578063715018a6146100fc57806382e215ab146101045780638796ba8c1461013c5780638da5cb5b1461015d575b600080fd5b6100e76100e23660046107c4565b61022a565b005b6100e76100f73660046107fc565b610304565b6100e76103e1565b6101276101123660046108c6565b60046020526000908152604090205460ff1681565b60405190151581526020015b60405180910390f35b61014f61014a3660046108c6565b6103f5565b604051908152602001610133565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b03165b6040516001600160a01b039091168152602001610133565b600354610189906001600160a01b031681565b6100e76101c23660046108fb565b610416565b6100e76101d536600461091d565b610440565b6101ed6101e83660046108c6565b610582565b60405161013392919061099c565b600254610189906001600160a01b031681565b6100e761021c3660046108fb565b6105f8565b61014f60015481565b610232610636565b604080518082018252600080825282518181526020808201855280840191825286835260048152939091208251815460ff191690151517815590518051929391926102839260018501920190610764565b505060008054600180820183559180527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5630184905583905550604080518381526020810183905230918101919091527fe697eb68c0228bd7d4e553246a2a86e8402d0895e45092ef8ae87b4cfd29f016906060015b60405180910390a15050565b6002546001600160a01b031633146103635760405162461bcd60e51b815260206004820152601860248201527f446170704c696e6b5652462e6f6e6c79446170704c696e6b000000000000000060448201526064015b60405180910390fd5b60408051808201825260018082526020808301858152600087815260048352949094208351815460ff1916901515178155935180519394936103ac938501929190910190610764565b509050507ff3cb4deb0441dd096356debf166f879d78cadc19e4b94053c8bea6d3940de93a82826040516102f89291906109bf565b6103e9610636565b6103f36000610691565b565b6000818154811061040557600080fd5b600091825260209091200154905081565b61041e610636565b600280546001600160a01b0319166001600160a01b0392909216919091179055565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff166000811580156104865750825b905060008267ffffffffffffffff1660011480156104a35750303b155b9050811580156104b1575080155b156104cf5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156104f957845460ff60401b1916600160401b1785555b61050288610702565b600380546001600160a01b038089166001600160a01b03199283161790925560028054928a1692909116919091179055831561057857845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b600081815260046020908152604080832080546001909101805483518186028101860190945280845260609460ff9093169391928391908301828280156105e857602002820191906000526020600020905b8154815260200190600101908083116105d4575b5050505050905091509150915091565b610600610636565b6001600160a01b03811661062a57604051631e4fbdf760e01b81526000600482015260240161035a565b61063381610691565b50565b336106687f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146103f35760405163118cdaa760e01b815233600482015260240161035a565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b61070a610713565b6106338161075c565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff166103f357604051631afcd79f60e31b815260040160405180910390fd5b610600610713565b82805482825590600052602060002090810192821561079f579160200282015b8281111561079f578251825591602001919060010190610784565b506107ab9291506107af565b5090565b5b808211156107ab57600081556001016107b0565b600080604083850312156107d757600080fd5b50508035926020909101359150565b634e487b7160e01b600052604160045260246000fd5b6000806040838503121561080f57600080fd5b8235915060208084013567ffffffffffffffff8082111561082f57600080fd5b818601915086601f83011261084357600080fd5b813581811115610855576108556107e6565b8060051b604051601f19603f8301168101818110858211171561087a5761087a6107e6565b60405291825284820192508381018501918983111561089857600080fd5b938501935b828510156108b65784358452938501939285019261089d565b8096505050505050509250929050565b6000602082840312156108d857600080fd5b5035919050565b80356001600160a01b03811681146108f657600080fd5b919050565b60006020828403121561090d57600080fd5b610916826108df565b9392505050565b60008060006060848603121561093257600080fd5b61093b846108df565b9250610949602085016108df565b9150610957604085016108df565b90509250925092565b60008151808452602080850194506020840160005b8381101561099157815187529582019590820190600101610975565b509495945050505050565b82151581526040602082015260006109b76040830184610960565b949350505050565b8281526040602082015260006109b7604083018461096056fea2646970667358221220518cb57a5b19618354180270f288446574d7b756efbe1e9ca7eabb7d00948ce564736f6c63430008160033",
}

// DappLinkVRFManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use DappLinkVRFManagerMetaData.ABI instead.
var DappLinkVRFManagerABI = DappLinkVRFManagerMetaData.ABI

// DappLinkVRFManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DappLinkVRFManagerMetaData.Bin instead.
var DappLinkVRFManagerBin = DappLinkVRFManagerMetaData.Bin

// DeployDappLinkVRFManager deploys a new Ethereum contract, binding an instance of DappLinkVRFManager to it.
func DeployDappLinkVRFManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DappLinkVRFManager, error) {
	parsed, err := DappLinkVRFManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DappLinkVRFManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DappLinkVRFManager{DappLinkVRFManagerCaller: DappLinkVRFManagerCaller{contract: contract}, DappLinkVRFManagerTransactor: DappLinkVRFManagerTransactor{contract: contract}, DappLinkVRFManagerFilterer: DappLinkVRFManagerFilterer{contract: contract}}, nil
}

// DappLinkVRFManager is an auto generated Go binding around an Ethereum contract.
type DappLinkVRFManager struct {
	DappLinkVRFManagerCaller     // Read-only binding to the contract
	DappLinkVRFManagerTransactor // Write-only binding to the contract
	DappLinkVRFManagerFilterer   // Log filterer for contract events
}

// DappLinkVRFManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type DappLinkVRFManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DappLinkVRFManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DappLinkVRFManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DappLinkVRFManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DappLinkVRFManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DappLinkVRFManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DappLinkVRFManagerSession struct {
	Contract     *DappLinkVRFManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DappLinkVRFManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DappLinkVRFManagerCallerSession struct {
	Contract *DappLinkVRFManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// DappLinkVRFManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DappLinkVRFManagerTransactorSession struct {
	Contract     *DappLinkVRFManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// DappLinkVRFManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type DappLinkVRFManagerRaw struct {
	Contract *DappLinkVRFManager // Generic contract binding to access the raw methods on
}

// DappLinkVRFManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DappLinkVRFManagerCallerRaw struct {
	Contract *DappLinkVRFManagerCaller // Generic read-only contract binding to access the raw methods on
}

// DappLinkVRFManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DappLinkVRFManagerTransactorRaw struct {
	Contract *DappLinkVRFManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDappLinkVRFManager creates a new instance of DappLinkVRFManager, bound to a specific deployed contract.
func NewDappLinkVRFManager(address common.Address, backend bind.ContractBackend) (*DappLinkVRFManager, error) {
	contract, err := bindDappLinkVRFManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DappLinkVRFManager{DappLinkVRFManagerCaller: DappLinkVRFManagerCaller{contract: contract}, DappLinkVRFManagerTransactor: DappLinkVRFManagerTransactor{contract: contract}, DappLinkVRFManagerFilterer: DappLinkVRFManagerFilterer{contract: contract}}, nil
}

// NewDappLinkVRFManagerCaller creates a new read-only instance of DappLinkVRFManager, bound to a specific deployed contract.
func NewDappLinkVRFManagerCaller(address common.Address, caller bind.ContractCaller) (*DappLinkVRFManagerCaller, error) {
	contract, err := bindDappLinkVRFManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DappLinkVRFManagerCaller{contract: contract}, nil
}

// NewDappLinkVRFManagerTransactor creates a new write-only instance of DappLinkVRFManager, bound to a specific deployed contract.
func NewDappLinkVRFManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*DappLinkVRFManagerTransactor, error) {
	contract, err := bindDappLinkVRFManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DappLinkVRFManagerTransactor{contract: contract}, nil
}

// NewDappLinkVRFManagerFilterer creates a new log filterer instance of DappLinkVRFManager, bound to a specific deployed contract.
func NewDappLinkVRFManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*DappLinkVRFManagerFilterer, error) {
	contract, err := bindDappLinkVRFManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DappLinkVRFManagerFilterer{contract: contract}, nil
}

// bindDappLinkVRFManager binds a generic wrapper to an already deployed contract.
func bindDappLinkVRFManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DappLinkVRFManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DappLinkVRFManager *DappLinkVRFManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DappLinkVRFManager.Contract.DappLinkVRFManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DappLinkVRFManager *DappLinkVRFManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappLinkVRFManager.Contract.DappLinkVRFManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DappLinkVRFManager *DappLinkVRFManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DappLinkVRFManager.Contract.DappLinkVRFManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DappLinkVRFManager *DappLinkVRFManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DappLinkVRFManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DappLinkVRFManager *DappLinkVRFManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappLinkVRFManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DappLinkVRFManager *DappLinkVRFManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DappLinkVRFManager.Contract.contract.Transact(opts, method, params...)
}

// BlsRegistry is a free data retrieval call binding the contract method 0x93e4b880.
//
// Solidity: function blsRegistry() view returns(address)
func (_DappLinkVRFManager *DappLinkVRFManagerCaller) BlsRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DappLinkVRFManager.contract.Call(opts, &out, "blsRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsRegistry is a free data retrieval call binding the contract method 0x93e4b880.
//
// Solidity: function blsRegistry() view returns(address)
func (_DappLinkVRFManager *DappLinkVRFManagerSession) BlsRegistry() (common.Address, error) {
	return _DappLinkVRFManager.Contract.BlsRegistry(&_DappLinkVRFManager.CallOpts)
}

// BlsRegistry is a free data retrieval call binding the contract method 0x93e4b880.
//
// Solidity: function blsRegistry() view returns(address)
func (_DappLinkVRFManager *DappLinkVRFManagerCallerSession) BlsRegistry() (common.Address, error) {
	return _DappLinkVRFManager.Contract.BlsRegistry(&_DappLinkVRFManager.CallOpts)
}

// DappLinkAddress is a free data retrieval call binding the contract method 0xf0c28a41.
//
// Solidity: function dappLinkAddress() view returns(address)
func (_DappLinkVRFManager *DappLinkVRFManagerCaller) DappLinkAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DappLinkVRFManager.contract.Call(opts, &out, "dappLinkAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DappLinkAddress is a free data retrieval call binding the contract method 0xf0c28a41.
//
// Solidity: function dappLinkAddress() view returns(address)
func (_DappLinkVRFManager *DappLinkVRFManagerSession) DappLinkAddress() (common.Address, error) {
	return _DappLinkVRFManager.Contract.DappLinkAddress(&_DappLinkVRFManager.CallOpts)
}

// DappLinkAddress is a free data retrieval call binding the contract method 0xf0c28a41.
//
// Solidity: function dappLinkAddress() view returns(address)
func (_DappLinkVRFManager *DappLinkVRFManagerCallerSession) DappLinkAddress() (common.Address, error) {
	return _DappLinkVRFManager.Contract.DappLinkAddress(&_DappLinkVRFManager.CallOpts)
}

// GetRequestStatus is a free data retrieval call binding the contract method 0xd8a4676f.
//
// Solidity: function getRequestStatus(uint256 _requestId) view returns(bool fulfilled, uint256[] randomWords)
func (_DappLinkVRFManager *DappLinkVRFManagerCaller) GetRequestStatus(opts *bind.CallOpts, _requestId *big.Int) (struct {
	Fulfilled   bool
	RandomWords []*big.Int
}, error) {
	var out []interface{}
	err := _DappLinkVRFManager.contract.Call(opts, &out, "getRequestStatus", _requestId)

	outstruct := new(struct {
		Fulfilled   bool
		RandomWords []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fulfilled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.RandomWords = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetRequestStatus is a free data retrieval call binding the contract method 0xd8a4676f.
//
// Solidity: function getRequestStatus(uint256 _requestId) view returns(bool fulfilled, uint256[] randomWords)
func (_DappLinkVRFManager *DappLinkVRFManagerSession) GetRequestStatus(_requestId *big.Int) (struct {
	Fulfilled   bool
	RandomWords []*big.Int
}, error) {
	return _DappLinkVRFManager.Contract.GetRequestStatus(&_DappLinkVRFManager.CallOpts, _requestId)
}

// GetRequestStatus is a free data retrieval call binding the contract method 0xd8a4676f.
//
// Solidity: function getRequestStatus(uint256 _requestId) view returns(bool fulfilled, uint256[] randomWords)
func (_DappLinkVRFManager *DappLinkVRFManagerCallerSession) GetRequestStatus(_requestId *big.Int) (struct {
	Fulfilled   bool
	RandomWords []*big.Int
}, error) {
	return _DappLinkVRFManager.Contract.GetRequestStatus(&_DappLinkVRFManager.CallOpts, _requestId)
}

// LastRequestId is a free data retrieval call binding the contract method 0xfc2a88c3.
//
// Solidity: function lastRequestId() view returns(uint256)
func (_DappLinkVRFManager *DappLinkVRFManagerCaller) LastRequestId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DappLinkVRFManager.contract.Call(opts, &out, "lastRequestId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRequestId is a free data retrieval call binding the contract method 0xfc2a88c3.
//
// Solidity: function lastRequestId() view returns(uint256)
func (_DappLinkVRFManager *DappLinkVRFManagerSession) LastRequestId() (*big.Int, error) {
	return _DappLinkVRFManager.Contract.LastRequestId(&_DappLinkVRFManager.CallOpts)
}

// LastRequestId is a free data retrieval call binding the contract method 0xfc2a88c3.
//
// Solidity: function lastRequestId() view returns(uint256)
func (_DappLinkVRFManager *DappLinkVRFManagerCallerSession) LastRequestId() (*big.Int, error) {
	return _DappLinkVRFManager.Contract.LastRequestId(&_DappLinkVRFManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DappLinkVRFManager *DappLinkVRFManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DappLinkVRFManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DappLinkVRFManager *DappLinkVRFManagerSession) Owner() (common.Address, error) {
	return _DappLinkVRFManager.Contract.Owner(&_DappLinkVRFManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DappLinkVRFManager *DappLinkVRFManagerCallerSession) Owner() (common.Address, error) {
	return _DappLinkVRFManager.Contract.Owner(&_DappLinkVRFManager.CallOpts)
}

// RequestIds is a free data retrieval call binding the contract method 0x8796ba8c.
//
// Solidity: function requestIds(uint256 ) view returns(uint256)
func (_DappLinkVRFManager *DappLinkVRFManagerCaller) RequestIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DappLinkVRFManager.contract.Call(opts, &out, "requestIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestIds is a free data retrieval call binding the contract method 0x8796ba8c.
//
// Solidity: function requestIds(uint256 ) view returns(uint256)
func (_DappLinkVRFManager *DappLinkVRFManagerSession) RequestIds(arg0 *big.Int) (*big.Int, error) {
	return _DappLinkVRFManager.Contract.RequestIds(&_DappLinkVRFManager.CallOpts, arg0)
}

// RequestIds is a free data retrieval call binding the contract method 0x8796ba8c.
//
// Solidity: function requestIds(uint256 ) view returns(uint256)
func (_DappLinkVRFManager *DappLinkVRFManagerCallerSession) RequestIds(arg0 *big.Int) (*big.Int, error) {
	return _DappLinkVRFManager.Contract.RequestIds(&_DappLinkVRFManager.CallOpts, arg0)
}

// RequestMapping is a free data retrieval call binding the contract method 0x82e215ab.
//
// Solidity: function requestMapping(uint256 ) view returns(bool fulfilled)
func (_DappLinkVRFManager *DappLinkVRFManagerCaller) RequestMapping(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _DappLinkVRFManager.contract.Call(opts, &out, "requestMapping", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RequestMapping is a free data retrieval call binding the contract method 0x82e215ab.
//
// Solidity: function requestMapping(uint256 ) view returns(bool fulfilled)
func (_DappLinkVRFManager *DappLinkVRFManagerSession) RequestMapping(arg0 *big.Int) (bool, error) {
	return _DappLinkVRFManager.Contract.RequestMapping(&_DappLinkVRFManager.CallOpts, arg0)
}

// RequestMapping is a free data retrieval call binding the contract method 0x82e215ab.
//
// Solidity: function requestMapping(uint256 ) view returns(bool fulfilled)
func (_DappLinkVRFManager *DappLinkVRFManagerCallerSession) RequestMapping(arg0 *big.Int) (bool, error) {
	return _DappLinkVRFManager.Contract.RequestMapping(&_DappLinkVRFManager.CallOpts, arg0)
}

// FulfillRandomWords is a paid mutator transaction binding the contract method 0x38ba4614.
//
// Solidity: function fulfillRandomWords(uint256 _requestId, uint256[] _randomWords) returns()
func (_DappLinkVRFManager *DappLinkVRFManagerTransactor) FulfillRandomWords(opts *bind.TransactOpts, _requestId *big.Int, _randomWords []*big.Int) (*types.Transaction, error) {
	return _DappLinkVRFManager.contract.Transact(opts, "fulfillRandomWords", _requestId, _randomWords)
}

// FulfillRandomWords is a paid mutator transaction binding the contract method 0x38ba4614.
//
// Solidity: function fulfillRandomWords(uint256 _requestId, uint256[] _randomWords) returns()
func (_DappLinkVRFManager *DappLinkVRFManagerSession) FulfillRandomWords(_requestId *big.Int, _randomWords []*big.Int) (*types.Transaction, error) {
	return _DappLinkVRFManager.Contract.FulfillRandomWords(&_DappLinkVRFManager.TransactOpts, _requestId, _randomWords)
}

// FulfillRandomWords is a paid mutator transaction binding the contract method 0x38ba4614.
//
// Solidity: function fulfillRandomWords(uint256 _requestId, uint256[] _randomWords) returns()
func (_DappLinkVRFManager *DappLinkVRFManagerTransactorSession) FulfillRandomWords(_requestId *big.Int, _randomWords []*big.Int) (*types.Transaction, error) {
	return _DappLinkVRFManager.Contract.FulfillRandomWords(&_DappLinkVRFManager.TransactOpts, _requestId, _randomWords)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address initialOwner, address _dappLinkAddress, address _blsRegistry) returns()
func (_DappLinkVRFManager *DappLinkVRFManagerTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, _dappLinkAddress common.Address, _blsRegistry common.Address) (*types.Transaction, error) {
	return _DappLinkVRFManager.contract.Transact(opts, "initialize", initialOwner, _dappLinkAddress, _blsRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address initialOwner, address _dappLinkAddress, address _blsRegistry) returns()
func (_DappLinkVRFManager *DappLinkVRFManagerSession) Initialize(initialOwner common.Address, _dappLinkAddress common.Address, _blsRegistry common.Address) (*types.Transaction, error) {
	return _DappLinkVRFManager.Contract.Initialize(&_DappLinkVRFManager.TransactOpts, initialOwner, _dappLinkAddress, _blsRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address initialOwner, address _dappLinkAddress, address _blsRegistry) returns()
func (_DappLinkVRFManager *DappLinkVRFManagerTransactorSession) Initialize(initialOwner common.Address, _dappLinkAddress common.Address, _blsRegistry common.Address) (*types.Transaction, error) {
	return _DappLinkVRFManager.Contract.Initialize(&_DappLinkVRFManager.TransactOpts, initialOwner, _dappLinkAddress, _blsRegistry)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DappLinkVRFManager *DappLinkVRFManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappLinkVRFManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DappLinkVRFManager *DappLinkVRFManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _DappLinkVRFManager.Contract.RenounceOwnership(&_DappLinkVRFManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DappLinkVRFManager *DappLinkVRFManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DappLinkVRFManager.Contract.RenounceOwnership(&_DappLinkVRFManager.TransactOpts)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0x1b739ef1.
//
// Solidity: function requestRandomWords(uint256 _requestId, uint256 _numWords) returns()
func (_DappLinkVRFManager *DappLinkVRFManagerTransactor) RequestRandomWords(opts *bind.TransactOpts, _requestId *big.Int, _numWords *big.Int) (*types.Transaction, error) {
	return _DappLinkVRFManager.contract.Transact(opts, "requestRandomWords", _requestId, _numWords)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0x1b739ef1.
//
// Solidity: function requestRandomWords(uint256 _requestId, uint256 _numWords) returns()
func (_DappLinkVRFManager *DappLinkVRFManagerSession) RequestRandomWords(_requestId *big.Int, _numWords *big.Int) (*types.Transaction, error) {
	return _DappLinkVRFManager.Contract.RequestRandomWords(&_DappLinkVRFManager.TransactOpts, _requestId, _numWords)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0x1b739ef1.
//
// Solidity: function requestRandomWords(uint256 _requestId, uint256 _numWords) returns()
func (_DappLinkVRFManager *DappLinkVRFManagerTransactorSession) RequestRandomWords(_requestId *big.Int, _numWords *big.Int) (*types.Transaction, error) {
	return _DappLinkVRFManager.Contract.RequestRandomWords(&_DappLinkVRFManager.TransactOpts, _requestId, _numWords)
}

// SetDappLink is a paid mutator transaction binding the contract method 0x996869d0.
//
// Solidity: function setDappLink(address _dappLinkAddress) returns()
func (_DappLinkVRFManager *DappLinkVRFManagerTransactor) SetDappLink(opts *bind.TransactOpts, _dappLinkAddress common.Address) (*types.Transaction, error) {
	return _DappLinkVRFManager.contract.Transact(opts, "setDappLink", _dappLinkAddress)
}

// SetDappLink is a paid mutator transaction binding the contract method 0x996869d0.
//
// Solidity: function setDappLink(address _dappLinkAddress) returns()
func (_DappLinkVRFManager *DappLinkVRFManagerSession) SetDappLink(_dappLinkAddress common.Address) (*types.Transaction, error) {
	return _DappLinkVRFManager.Contract.SetDappLink(&_DappLinkVRFManager.TransactOpts, _dappLinkAddress)
}

// SetDappLink is a paid mutator transaction binding the contract method 0x996869d0.
//
// Solidity: function setDappLink(address _dappLinkAddress) returns()
func (_DappLinkVRFManager *DappLinkVRFManagerTransactorSession) SetDappLink(_dappLinkAddress common.Address) (*types.Transaction, error) {
	return _DappLinkVRFManager.Contract.SetDappLink(&_DappLinkVRFManager.TransactOpts, _dappLinkAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DappLinkVRFManager *DappLinkVRFManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DappLinkVRFManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DappLinkVRFManager *DappLinkVRFManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DappLinkVRFManager.Contract.TransferOwnership(&_DappLinkVRFManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DappLinkVRFManager *DappLinkVRFManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DappLinkVRFManager.Contract.TransferOwnership(&_DappLinkVRFManager.TransactOpts, newOwner)
}

// DappLinkVRFManagerFillRandomWordsIterator is returned from FilterFillRandomWords and is used to iterate over the raw logs and unpacked data for FillRandomWords events raised by the DappLinkVRFManager contract.
type DappLinkVRFManagerFillRandomWordsIterator struct {
	Event *DappLinkVRFManagerFillRandomWords // Event containing the contract specifics and raw log

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
func (it *DappLinkVRFManagerFillRandomWordsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DappLinkVRFManagerFillRandomWords)
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
		it.Event = new(DappLinkVRFManagerFillRandomWords)
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
func (it *DappLinkVRFManagerFillRandomWordsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DappLinkVRFManagerFillRandomWordsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DappLinkVRFManagerFillRandomWords represents a FillRandomWords event raised by the DappLinkVRFManager contract.
type DappLinkVRFManagerFillRandomWords struct {
	RequestId   *big.Int
	RandomWords []*big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFillRandomWords is a free log retrieval operation binding the contract event 0xf3cb4deb0441dd096356debf166f879d78cadc19e4b94053c8bea6d3940de93a.
//
// Solidity: event FillRandomWords(uint256 requestId, uint256[] randomWords)
func (_DappLinkVRFManager *DappLinkVRFManagerFilterer) FilterFillRandomWords(opts *bind.FilterOpts) (*DappLinkVRFManagerFillRandomWordsIterator, error) {

	logs, sub, err := _DappLinkVRFManager.contract.FilterLogs(opts, "FillRandomWords")
	if err != nil {
		return nil, err
	}
	return &DappLinkVRFManagerFillRandomWordsIterator{contract: _DappLinkVRFManager.contract, event: "FillRandomWords", logs: logs, sub: sub}, nil
}

// WatchFillRandomWords is a free log subscription operation binding the contract event 0xf3cb4deb0441dd096356debf166f879d78cadc19e4b94053c8bea6d3940de93a.
//
// Solidity: event FillRandomWords(uint256 requestId, uint256[] randomWords)
func (_DappLinkVRFManager *DappLinkVRFManagerFilterer) WatchFillRandomWords(opts *bind.WatchOpts, sink chan<- *DappLinkVRFManagerFillRandomWords) (event.Subscription, error) {

	logs, sub, err := _DappLinkVRFManager.contract.WatchLogs(opts, "FillRandomWords")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DappLinkVRFManagerFillRandomWords)
				if err := _DappLinkVRFManager.contract.UnpackLog(event, "FillRandomWords", log); err != nil {
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

// ParseFillRandomWords is a log parse operation binding the contract event 0xf3cb4deb0441dd096356debf166f879d78cadc19e4b94053c8bea6d3940de93a.
//
// Solidity: event FillRandomWords(uint256 requestId, uint256[] randomWords)
func (_DappLinkVRFManager *DappLinkVRFManagerFilterer) ParseFillRandomWords(log types.Log) (*DappLinkVRFManagerFillRandomWords, error) {
	event := new(DappLinkVRFManagerFillRandomWords)
	if err := _DappLinkVRFManager.contract.UnpackLog(event, "FillRandomWords", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DappLinkVRFManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the DappLinkVRFManager contract.
type DappLinkVRFManagerInitializedIterator struct {
	Event *DappLinkVRFManagerInitialized // Event containing the contract specifics and raw log

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
func (it *DappLinkVRFManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DappLinkVRFManagerInitialized)
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
		it.Event = new(DappLinkVRFManagerInitialized)
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
func (it *DappLinkVRFManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DappLinkVRFManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DappLinkVRFManagerInitialized represents a Initialized event raised by the DappLinkVRFManager contract.
type DappLinkVRFManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_DappLinkVRFManager *DappLinkVRFManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*DappLinkVRFManagerInitializedIterator, error) {

	logs, sub, err := _DappLinkVRFManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DappLinkVRFManagerInitializedIterator{contract: _DappLinkVRFManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_DappLinkVRFManager *DappLinkVRFManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DappLinkVRFManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _DappLinkVRFManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DappLinkVRFManagerInitialized)
				if err := _DappLinkVRFManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_DappLinkVRFManager *DappLinkVRFManagerFilterer) ParseInitialized(log types.Log) (*DappLinkVRFManagerInitialized, error) {
	event := new(DappLinkVRFManagerInitialized)
	if err := _DappLinkVRFManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DappLinkVRFManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DappLinkVRFManager contract.
type DappLinkVRFManagerOwnershipTransferredIterator struct {
	Event *DappLinkVRFManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DappLinkVRFManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DappLinkVRFManagerOwnershipTransferred)
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
		it.Event = new(DappLinkVRFManagerOwnershipTransferred)
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
func (it *DappLinkVRFManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DappLinkVRFManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DappLinkVRFManagerOwnershipTransferred represents a OwnershipTransferred event raised by the DappLinkVRFManager contract.
type DappLinkVRFManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DappLinkVRFManager *DappLinkVRFManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DappLinkVRFManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DappLinkVRFManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DappLinkVRFManagerOwnershipTransferredIterator{contract: _DappLinkVRFManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DappLinkVRFManager *DappLinkVRFManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DappLinkVRFManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DappLinkVRFManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DappLinkVRFManagerOwnershipTransferred)
				if err := _DappLinkVRFManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_DappLinkVRFManager *DappLinkVRFManagerFilterer) ParseOwnershipTransferred(log types.Log) (*DappLinkVRFManagerOwnershipTransferred, error) {
	event := new(DappLinkVRFManagerOwnershipTransferred)
	if err := _DappLinkVRFManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DappLinkVRFManagerRequestSentIterator is returned from FilterRequestSent and is used to iterate over the raw logs and unpacked data for RequestSent events raised by the DappLinkVRFManager contract.
type DappLinkVRFManagerRequestSentIterator struct {
	Event *DappLinkVRFManagerRequestSent // Event containing the contract specifics and raw log

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
func (it *DappLinkVRFManagerRequestSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DappLinkVRFManagerRequestSent)
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
		it.Event = new(DappLinkVRFManagerRequestSent)
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
func (it *DappLinkVRFManagerRequestSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DappLinkVRFManagerRequestSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DappLinkVRFManagerRequestSent represents a RequestSent event raised by the DappLinkVRFManager contract.
type DappLinkVRFManagerRequestSent struct {
	RequestId *big.Int
	NumWords  *big.Int
	Current   common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRequestSent is a free log retrieval operation binding the contract event 0xe697eb68c0228bd7d4e553246a2a86e8402d0895e45092ef8ae87b4cfd29f016.
//
// Solidity: event RequestSent(uint256 requestId, uint256 _numWords, address current)
func (_DappLinkVRFManager *DappLinkVRFManagerFilterer) FilterRequestSent(opts *bind.FilterOpts) (*DappLinkVRFManagerRequestSentIterator, error) {

	logs, sub, err := _DappLinkVRFManager.contract.FilterLogs(opts, "RequestSent")
	if err != nil {
		return nil, err
	}
	return &DappLinkVRFManagerRequestSentIterator{contract: _DappLinkVRFManager.contract, event: "RequestSent", logs: logs, sub: sub}, nil
}

// WatchRequestSent is a free log subscription operation binding the contract event 0xe697eb68c0228bd7d4e553246a2a86e8402d0895e45092ef8ae87b4cfd29f016.
//
// Solidity: event RequestSent(uint256 requestId, uint256 _numWords, address current)
func (_DappLinkVRFManager *DappLinkVRFManagerFilterer) WatchRequestSent(opts *bind.WatchOpts, sink chan<- *DappLinkVRFManagerRequestSent) (event.Subscription, error) {

	logs, sub, err := _DappLinkVRFManager.contract.WatchLogs(opts, "RequestSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DappLinkVRFManagerRequestSent)
				if err := _DappLinkVRFManager.contract.UnpackLog(event, "RequestSent", log); err != nil {
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

// ParseRequestSent is a log parse operation binding the contract event 0xe697eb68c0228bd7d4e553246a2a86e8402d0895e45092ef8ae87b4cfd29f016.
//
// Solidity: event RequestSent(uint256 requestId, uint256 _numWords, address current)
func (_DappLinkVRFManager *DappLinkVRFManagerFilterer) ParseRequestSent(log types.Log) (*DappLinkVRFManagerRequestSent, error) {
	event := new(DappLinkVRFManagerRequestSent)
	if err := _DappLinkVRFManager.contract.UnpackLog(event, "RequestSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
