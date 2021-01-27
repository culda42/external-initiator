// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// ContractABI is the input ABI used to generate the binding from.
const ContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"b32\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"u256\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"b322\",\"type\":\"bytes32\"}],\"name\":\"Output\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"SetBytes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"SetBytes32\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SetUint256\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBytes32\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUint256\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_value\",\"type\":\"bytes\"}],\"name\":\"requestedBytes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_value\",\"type\":\"bytes32\"}],\"name\":\"requestedBytes32\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"requestedUint256\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_value\",\"type\":\"bytes\"}],\"name\":\"setBytes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_value\",\"type\":\"bytes32\"}],\"name\":\"setBytes32\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setUint256\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ContractBin is the compiled bytecode used for deploying new contracts.
var ContractBin = "0x60806040523480156100115760006000fd5b50610017565b61090a806100266000396000f3fe60806040523480156100115760006000fd5b50600436106100a25760003560e01c80636889597911610067578063688959791461013a578063c2b12a7314610158578063d2282dc514610174578063da359dc814610190578063ed53e511146101ac576100a2565b80626d6cae146100a85780630bcd3b33146100c65780631f903037146100e45780633345b4d01461010257806346ddd1ff1461011e576100a2565b60006000fd5b6100b06101c8565b6040516100bd919061072c565b60405180910390f35b6100ce6101d1565b6040516100db919061076d565b60405180910390f35b6100ec610262565b6040516100f9919061072c565b60405180910390f35b61011c600480360381019061011791906105f1565b61026b565b005b61013860048036038101906101339190610595565b61028f565b005b6101426102b5565b60405161014f9190610790565b60405180910390f35b610172600480360381019061016d919061052b565b6102be565b005b61018e60048036038101906101899190610679565b61031a565b005b6101aa60048036038101906101a59190610630565b61036e565b005b6101c660048036038101906101c19190610556565b6103d8565b005b60026000505481565b600360005080546101e190610825565b80601f016020809104026020016040519081016040528092919081815260200182805461020d90610825565b801561025a5780601f1061022f5761010080835404028352916020019161025a565b820191906000526020600020905b81548152906001019060200180831161023d57829003601f168201915b505050505081565b60006000505481565b8160026000508190906000191690555061028a8161031a63ffffffff16565b5b5050565b826002600050819090600019169055506102af828261036e63ffffffff16565b5b505050565b60016000505481565b8060006000508190906000191690555080600019163373ffffffffffffffffffffffffffffffffffffffff167fdc73ee99832252105ed74a404690c2f10ad1b294cbbeb0ff5cded48ef2aa437d60405160405180910390a35b50565b806001600050819090905550803373ffffffffffffffffffffffffffffffffffffffff167fd943f063acdb1c6f206cf6a3f6d1ba39687bcc07feb7f44019bdbd4773c9c28d60405160405180910390a35b50565b8181600360005091906103829291906103fc565b503373ffffffffffffffffffffffffffffffffffffffff167ff22a519d38e59bc517532f666f8da532fdd5356e68d617191e82a8fdcc8abdcf83836040516103cb929190610748565b60405180910390a25b5050565b816002600050819090600019169055506103f7816102be63ffffffff16565b5b5050565b82805461040890610825565b90600052602060002090601f01602090048101928261042a5760008555610476565b82601f1061044357803560ff1916838001178555610476565b82800160010185558215610476579182015b828111156104755782358260005090905591602001919060010190610455565b5b5090506104839190610487565b5090565b61048c565b808211156104a6576000818150600090555060010161048c565b5090566108d3565b6000813590506104bd8161089d565b5b92915050565b6000600083601f84011215156104da5760006000fd5b8235905067ffffffffffffffff8111156104f45760006000fd5b60208301915083600182028301111561050d5760006000fd5b5b9250929050565b600081359050610524816108b8565b5b92915050565b60006020828403121561053e5760006000fd5b600061054c848285016104ae565b9150505b92915050565b600060006040838503121561056b5760006000fd5b6000610579858286016104ae565b925050602061058a858286016104ae565b9150505b9250929050565b600060006000604084860312156105ac5760006000fd5b60006105ba868287016104ae565b935050602084013567ffffffffffffffff8111156105d85760006000fd5b6105e4868287016104c4565b92509250505b9250925092565b60006000604083850312156106065760006000fd5b6000610614858286016104ae565b925050602061062585828601610515565b9150505b9250929050565b60006000602083850312156106455760006000fd5b600083013567ffffffffffffffff8111156106605760006000fd5b61066c858286016104c4565b92509250505b9250929050565b60006020828403121561068c5760006000fd5b600061069a84828501610515565b9150505b92915050565b6106ad816107ca565b82525b5050565b60006106c083856107b8565b93506106cd8385846107e0565b6106d68361088b565b840190505b9392505050565b60006106ed826107ac565b6106f781856107b8565b93506107078185602086016107f0565b6107108161088b565b84019150505b92915050565b610725816107d5565b82525b5050565b600060208201905061074160008301846106a4565b5b92915050565b600060208201905081810360008301526107638184866106b4565b90505b9392505050565b6000602082019050818103600083015261078781846106e2565b90505b92915050565b60006020820190506107a5600083018461071c565b5b92915050565b6000815190505b919050565b60008282526020820190505b92915050565b60008190505b919050565b60008190505b919050565b828183376000838301525b505050565b60005b8381101561080f5780820151818401525b6020810190506107f3565b8381111561081e576000848401525b505b505050565b60006002820490506001821680151561083f57607f821691505b602082108114156108535761085261085a565b5b505b919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b565b6000601f19601f83011690505b919050565b6108a6816107ca565b811415156108b45760006000fd5b5b50565b6108c1816107d5565b811415156108cf5760006000fd5b5b50565bfea2646970667358221220ed30311acdf687f9515e20648a7e6bb6bcafe40d07c077e55942f1f6715a852964736f6c63430008000033"

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// GetBytes is a free data retrieval call binding the contract method 0x0bcd3b33.
//
// Solidity: function getBytes() view returns(bytes)
func (_Contract *ContractCaller) GetBytes(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getBytes")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetBytes is a free data retrieval call binding the contract method 0x0bcd3b33.
//
// Solidity: function getBytes() view returns(bytes)
func (_Contract *ContractSession) GetBytes() ([]byte, error) {
	return _Contract.Contract.GetBytes(&_Contract.CallOpts)
}

// GetBytes is a free data retrieval call binding the contract method 0x0bcd3b33.
//
// Solidity: function getBytes() view returns(bytes)
func (_Contract *ContractCallerSession) GetBytes() ([]byte, error) {
	return _Contract.Contract.GetBytes(&_Contract.CallOpts)
}

// GetBytes32 is a free data retrieval call binding the contract method 0x1f903037.
//
// Solidity: function getBytes32() view returns(bytes32)
func (_Contract *ContractCaller) GetBytes32(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getBytes32")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBytes32 is a free data retrieval call binding the contract method 0x1f903037.
//
// Solidity: function getBytes32() view returns(bytes32)
func (_Contract *ContractSession) GetBytes32() ([32]byte, error) {
	return _Contract.Contract.GetBytes32(&_Contract.CallOpts)
}

// GetBytes32 is a free data retrieval call binding the contract method 0x1f903037.
//
// Solidity: function getBytes32() view returns(bytes32)
func (_Contract *ContractCallerSession) GetBytes32() ([32]byte, error) {
	return _Contract.Contract.GetBytes32(&_Contract.CallOpts)
}

// GetUint256 is a free data retrieval call binding the contract method 0x68895979.
//
// Solidity: function getUint256() view returns(uint256)
func (_Contract *ContractCaller) GetUint256(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getUint256")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUint256 is a free data retrieval call binding the contract method 0x68895979.
//
// Solidity: function getUint256() view returns(uint256)
func (_Contract *ContractSession) GetUint256() (*big.Int, error) {
	return _Contract.Contract.GetUint256(&_Contract.CallOpts)
}

// GetUint256 is a free data retrieval call binding the contract method 0x68895979.
//
// Solidity: function getUint256() view returns(uint256)
func (_Contract *ContractCallerSession) GetUint256() (*big.Int, error) {
	return _Contract.Contract.GetUint256(&_Contract.CallOpts)
}

// RequestId is a free data retrieval call binding the contract method 0x006d6cae.
//
// Solidity: function requestId() view returns(bytes32)
func (_Contract *ContractCaller) RequestId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "requestId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RequestId is a free data retrieval call binding the contract method 0x006d6cae.
//
// Solidity: function requestId() view returns(bytes32)
func (_Contract *ContractSession) RequestId() ([32]byte, error) {
	return _Contract.Contract.RequestId(&_Contract.CallOpts)
}

// RequestId is a free data retrieval call binding the contract method 0x006d6cae.
//
// Solidity: function requestId() view returns(bytes32)
func (_Contract *ContractCallerSession) RequestId() ([32]byte, error) {
	return _Contract.Contract.RequestId(&_Contract.CallOpts)
}

// RequestedBytes is a paid mutator transaction binding the contract method 0x46ddd1ff.
//
// Solidity: function requestedBytes(bytes32 _requestId, bytes _value) returns()
func (_Contract *ContractTransactor) RequestedBytes(opts *bind.TransactOpts, _requestId [32]byte, _value []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "requestedBytes", _requestId, _value)
}

// RequestedBytes is a paid mutator transaction binding the contract method 0x46ddd1ff.
//
// Solidity: function requestedBytes(bytes32 _requestId, bytes _value) returns()
func (_Contract *ContractSession) RequestedBytes(_requestId [32]byte, _value []byte) (*types.Transaction, error) {
	return _Contract.Contract.RequestedBytes(&_Contract.TransactOpts, _requestId, _value)
}

// RequestedBytes is a paid mutator transaction binding the contract method 0x46ddd1ff.
//
// Solidity: function requestedBytes(bytes32 _requestId, bytes _value) returns()
func (_Contract *ContractTransactorSession) RequestedBytes(_requestId [32]byte, _value []byte) (*types.Transaction, error) {
	return _Contract.Contract.RequestedBytes(&_Contract.TransactOpts, _requestId, _value)
}

// RequestedBytes32 is a paid mutator transaction binding the contract method 0xed53e511.
//
// Solidity: function requestedBytes32(bytes32 _requestId, bytes32 _value) returns()
func (_Contract *ContractTransactor) RequestedBytes32(opts *bind.TransactOpts, _requestId [32]byte, _value [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "requestedBytes32", _requestId, _value)
}

// RequestedBytes32 is a paid mutator transaction binding the contract method 0xed53e511.
//
// Solidity: function requestedBytes32(bytes32 _requestId, bytes32 _value) returns()
func (_Contract *ContractSession) RequestedBytes32(_requestId [32]byte, _value [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.RequestedBytes32(&_Contract.TransactOpts, _requestId, _value)
}

// RequestedBytes32 is a paid mutator transaction binding the contract method 0xed53e511.
//
// Solidity: function requestedBytes32(bytes32 _requestId, bytes32 _value) returns()
func (_Contract *ContractTransactorSession) RequestedBytes32(_requestId [32]byte, _value [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.RequestedBytes32(&_Contract.TransactOpts, _requestId, _value)
}

// RequestedUint256 is a paid mutator transaction binding the contract method 0x3345b4d0.
//
// Solidity: function requestedUint256(bytes32 _requestId, uint256 _value) returns()
func (_Contract *ContractTransactor) RequestedUint256(opts *bind.TransactOpts, _requestId [32]byte, _value *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "requestedUint256", _requestId, _value)
}

// RequestedUint256 is a paid mutator transaction binding the contract method 0x3345b4d0.
//
// Solidity: function requestedUint256(bytes32 _requestId, uint256 _value) returns()
func (_Contract *ContractSession) RequestedUint256(_requestId [32]byte, _value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RequestedUint256(&_Contract.TransactOpts, _requestId, _value)
}

// RequestedUint256 is a paid mutator transaction binding the contract method 0x3345b4d0.
//
// Solidity: function requestedUint256(bytes32 _requestId, uint256 _value) returns()
func (_Contract *ContractTransactorSession) RequestedUint256(_requestId [32]byte, _value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RequestedUint256(&_Contract.TransactOpts, _requestId, _value)
}

// SetBytes is a paid mutator transaction binding the contract method 0xda359dc8.
//
// Solidity: function setBytes(bytes _value) returns()
func (_Contract *ContractTransactor) SetBytes(opts *bind.TransactOpts, _value []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setBytes", _value)
}

// SetBytes is a paid mutator transaction binding the contract method 0xda359dc8.
//
// Solidity: function setBytes(bytes _value) returns()
func (_Contract *ContractSession) SetBytes(_value []byte) (*types.Transaction, error) {
	return _Contract.Contract.SetBytes(&_Contract.TransactOpts, _value)
}

// SetBytes is a paid mutator transaction binding the contract method 0xda359dc8.
//
// Solidity: function setBytes(bytes _value) returns()
func (_Contract *ContractTransactorSession) SetBytes(_value []byte) (*types.Transaction, error) {
	return _Contract.Contract.SetBytes(&_Contract.TransactOpts, _value)
}

// SetBytes32 is a paid mutator transaction binding the contract method 0xc2b12a73.
//
// Solidity: function setBytes32(bytes32 _value) returns()
func (_Contract *ContractTransactor) SetBytes32(opts *bind.TransactOpts, _value [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setBytes32", _value)
}

// SetBytes32 is a paid mutator transaction binding the contract method 0xc2b12a73.
//
// Solidity: function setBytes32(bytes32 _value) returns()
func (_Contract *ContractSession) SetBytes32(_value [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.SetBytes32(&_Contract.TransactOpts, _value)
}

// SetBytes32 is a paid mutator transaction binding the contract method 0xc2b12a73.
//
// Solidity: function setBytes32(bytes32 _value) returns()
func (_Contract *ContractTransactorSession) SetBytes32(_value [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.SetBytes32(&_Contract.TransactOpts, _value)
}

// SetUint256 is a paid mutator transaction binding the contract method 0xd2282dc5.
//
// Solidity: function setUint256(uint256 _value) returns()
func (_Contract *ContractTransactor) SetUint256(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setUint256", _value)
}

// SetUint256 is a paid mutator transaction binding the contract method 0xd2282dc5.
//
// Solidity: function setUint256(uint256 _value) returns()
func (_Contract *ContractSession) SetUint256(_value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetUint256(&_Contract.TransactOpts, _value)
}

// SetUint256 is a paid mutator transaction binding the contract method 0xd2282dc5.
//
// Solidity: function setUint256(uint256 _value) returns()
func (_Contract *ContractTransactorSession) SetUint256(_value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetUint256(&_Contract.TransactOpts, _value)
}

// ContractOutputIterator is returned from FilterOutput and is used to iterate over the raw logs and unpacked data for Output events raised by the Contract contract.
type ContractOutputIterator struct {
	Event *ContractOutput // Event containing the contract specifics and raw log

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
func (it *ContractOutputIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOutput)
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
		it.Event = new(ContractOutput)
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
func (it *ContractOutputIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOutputIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOutput represents a Output event raised by the Contract contract.
type ContractOutput struct {
	B32  [32]byte
	U256 *big.Int
	B322 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOutput is a free log retrieval operation binding the contract event 0x98353d8a928d8ef0ed244d763914357f71b8a4f66644fe62ea0aaf01c113a355.
//
// Solidity: event Output(bytes32 b32, uint256 u256, bytes32 b322)
func (_Contract *ContractFilterer) FilterOutput(opts *bind.FilterOpts) (*ContractOutputIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Output")
	if err != nil {
		return nil, err
	}
	return &ContractOutputIterator{contract: _Contract.contract, event: "Output", logs: logs, sub: sub}, nil
}

// WatchOutput is a free log subscription operation binding the contract event 0x98353d8a928d8ef0ed244d763914357f71b8a4f66644fe62ea0aaf01c113a355.
//
// Solidity: event Output(bytes32 b32, uint256 u256, bytes32 b322)
func (_Contract *ContractFilterer) WatchOutput(opts *bind.WatchOpts, sink chan<- *ContractOutput) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Output")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOutput)
				if err := _Contract.contract.UnpackLog(event, "Output", log); err != nil {
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

// ParseOutput is a log parse operation binding the contract event 0x98353d8a928d8ef0ed244d763914357f71b8a4f66644fe62ea0aaf01c113a355.
//
// Solidity: event Output(bytes32 b32, uint256 u256, bytes32 b322)
func (_Contract *ContractFilterer) ParseOutput(log types.Log) (*ContractOutput, error) {
	event := new(ContractOutput)
	if err := _Contract.contract.UnpackLog(event, "Output", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSetBytesIterator is returned from FilterSetBytes and is used to iterate over the raw logs and unpacked data for SetBytes events raised by the Contract contract.
type ContractSetBytesIterator struct {
	Event *ContractSetBytes // Event containing the contract specifics and raw log

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
func (it *ContractSetBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSetBytes)
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
		it.Event = new(ContractSetBytes)
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
func (it *ContractSetBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSetBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSetBytes represents a SetBytes event raised by the Contract contract.
type ContractSetBytes struct {
	From  common.Address
	Value []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetBytes is a free log retrieval operation binding the contract event 0xf22a519d38e59bc517532f666f8da532fdd5356e68d617191e82a8fdcc8abdcf.
//
// Solidity: event SetBytes(address indexed from, bytes value)
func (_Contract *ContractFilterer) FilterSetBytes(opts *bind.FilterOpts, from []common.Address) (*ContractSetBytesIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SetBytes", fromRule)
	if err != nil {
		return nil, err
	}
	return &ContractSetBytesIterator{contract: _Contract.contract, event: "SetBytes", logs: logs, sub: sub}, nil
}

// WatchSetBytes is a free log subscription operation binding the contract event 0xf22a519d38e59bc517532f666f8da532fdd5356e68d617191e82a8fdcc8abdcf.
//
// Solidity: event SetBytes(address indexed from, bytes value)
func (_Contract *ContractFilterer) WatchSetBytes(opts *bind.WatchOpts, sink chan<- *ContractSetBytes, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SetBytes", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSetBytes)
				if err := _Contract.contract.UnpackLog(event, "SetBytes", log); err != nil {
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

// ParseSetBytes is a log parse operation binding the contract event 0xf22a519d38e59bc517532f666f8da532fdd5356e68d617191e82a8fdcc8abdcf.
//
// Solidity: event SetBytes(address indexed from, bytes value)
func (_Contract *ContractFilterer) ParseSetBytes(log types.Log) (*ContractSetBytes, error) {
	event := new(ContractSetBytes)
	if err := _Contract.contract.UnpackLog(event, "SetBytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSetBytes32Iterator is returned from FilterSetBytes32 and is used to iterate over the raw logs and unpacked data for SetBytes32 events raised by the Contract contract.
type ContractSetBytes32Iterator struct {
	Event *ContractSetBytes32 // Event containing the contract specifics and raw log

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
func (it *ContractSetBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSetBytes32)
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
		it.Event = new(ContractSetBytes32)
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
func (it *ContractSetBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSetBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSetBytes32 represents a SetBytes32 event raised by the Contract contract.
type ContractSetBytes32 struct {
	From  common.Address
	Value [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetBytes32 is a free log retrieval operation binding the contract event 0xdc73ee99832252105ed74a404690c2f10ad1b294cbbeb0ff5cded48ef2aa437d.
//
// Solidity: event SetBytes32(address indexed from, bytes32 indexed value)
func (_Contract *ContractFilterer) FilterSetBytes32(opts *bind.FilterOpts, from []common.Address, value [][32]byte) (*ContractSetBytes32Iterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SetBytes32", fromRule, valueRule)
	if err != nil {
		return nil, err
	}
	return &ContractSetBytes32Iterator{contract: _Contract.contract, event: "SetBytes32", logs: logs, sub: sub}, nil
}

// WatchSetBytes32 is a free log subscription operation binding the contract event 0xdc73ee99832252105ed74a404690c2f10ad1b294cbbeb0ff5cded48ef2aa437d.
//
// Solidity: event SetBytes32(address indexed from, bytes32 indexed value)
func (_Contract *ContractFilterer) WatchSetBytes32(opts *bind.WatchOpts, sink chan<- *ContractSetBytes32, from []common.Address, value [][32]byte) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SetBytes32", fromRule, valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSetBytes32)
				if err := _Contract.contract.UnpackLog(event, "SetBytes32", log); err != nil {
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

// ParseSetBytes32 is a log parse operation binding the contract event 0xdc73ee99832252105ed74a404690c2f10ad1b294cbbeb0ff5cded48ef2aa437d.
//
// Solidity: event SetBytes32(address indexed from, bytes32 indexed value)
func (_Contract *ContractFilterer) ParseSetBytes32(log types.Log) (*ContractSetBytes32, error) {
	event := new(ContractSetBytes32)
	if err := _Contract.contract.UnpackLog(event, "SetBytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSetUint256Iterator is returned from FilterSetUint256 and is used to iterate over the raw logs and unpacked data for SetUint256 events raised by the Contract contract.
type ContractSetUint256Iterator struct {
	Event *ContractSetUint256 // Event containing the contract specifics and raw log

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
func (it *ContractSetUint256Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSetUint256)
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
		it.Event = new(ContractSetUint256)
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
func (it *ContractSetUint256Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSetUint256Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSetUint256 represents a SetUint256 event raised by the Contract contract.
type ContractSetUint256 struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetUint256 is a free log retrieval operation binding the contract event 0xd943f063acdb1c6f206cf6a3f6d1ba39687bcc07feb7f44019bdbd4773c9c28d.
//
// Solidity: event SetUint256(address indexed from, uint256 indexed value)
func (_Contract *ContractFilterer) FilterSetUint256(opts *bind.FilterOpts, from []common.Address, value []*big.Int) (*ContractSetUint256Iterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SetUint256", fromRule, valueRule)
	if err != nil {
		return nil, err
	}
	return &ContractSetUint256Iterator{contract: _Contract.contract, event: "SetUint256", logs: logs, sub: sub}, nil
}

// WatchSetUint256 is a free log subscription operation binding the contract event 0xd943f063acdb1c6f206cf6a3f6d1ba39687bcc07feb7f44019bdbd4773c9c28d.
//
// Solidity: event SetUint256(address indexed from, uint256 indexed value)
func (_Contract *ContractFilterer) WatchSetUint256(opts *bind.WatchOpts, sink chan<- *ContractSetUint256, from []common.Address, value []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SetUint256", fromRule, valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSetUint256)
				if err := _Contract.contract.UnpackLog(event, "SetUint256", log); err != nil {
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

// ParseSetUint256 is a log parse operation binding the contract event 0xd943f063acdb1c6f206cf6a3f6d1ba39687bcc07feb7f44019bdbd4773c9c28d.
//
// Solidity: event SetUint256(address indexed from, uint256 indexed value)
func (_Contract *ContractFilterer) ParseSetUint256(log types.Log) (*ContractSetUint256, error) {
	event := new(ContractSetUint256)
	if err := _Contract.contract.UnpackLog(event, "SetUint256", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
