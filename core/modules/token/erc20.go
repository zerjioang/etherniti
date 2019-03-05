// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package token

import (
	"errors"
	"math/big"

	"github.com/labstack/gommon/log"
	"github.com/zerjioang/etherniti/core/eth/rpc"
)

const (
	// TokenABI is the input ABI used to generate the binding from.
	defaultErc20TokenAbi = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenOwner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenOwner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"tokenOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"
)

type ERC20Token struct {
	address     string
	name        string
	symbol      string
	decimals    uint8
	totalSupply *big.Int
	abi         string
	cli         ethrpc.EthRPC
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken() ERC20Token {
	t := ERC20Token{}
	t.abi = defaultErc20TokenAbi
	return t
}

func (t *ERC20Token) Address(address string) {
	t.address = address
}

func (t *ERC20Token) Abi(abi string) {
	t.abi = abi
}

func (t *ERC20Token) Client(cli ethrpc.EthRPC) {
	t.cli = cli
}

func (t ERC20Token) getParams() interface{} {
	fromAddress := "" //DATA, 20 Bytes - (optional) The address the transaction is sent from.
	toAddress := ""   //DATA, 20 Bytes - The address the transaction is directed to.
	gas := 100        // QUANTITY - (optional) Integer of the gas provided for the transaction execution. eth_call consumes zero gas, but this parameter may be needed by some executions.
	gasPrice := ""    // QUANTITY - (optional) Integer of the gasPrice used for each paid gas
	value := ""       // QUANTITY - (optional) Integer of the value sent with this transaction
	data := ""        // DATA - (optional) Hash of the method signature and encoded parameters. For details see Ethereum Contract ABI
	block := "latest" //integer block number, or the string "latest", "earliest" or "pending", see the default block parameter

	params := []interface{}{fromAddress, toAddress, gas, gasPrice, value, data, block}
	return params
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (t *ERC20Token) Name() (string, error) {
	params := t.getParams()
	result, err := t.Call("name", params)
	if err != nil {
		return "", err
	}
	data, ok := result.(string)
	if ok {
		t.name = data
	}
	return data, nil
}

// Name is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function name() constant returns(string)
func (t *ERC20Token) Symbol() (string, error) {
	params := t.getParams()
	result, err := t.Call("symbol", params)
	if err != nil {
		return "", err
	}
	data, ok := result.(string)
	if ok {
		t.name = data
	}
	return data, nil
}

// Name is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function name() constant returns(string)
func (t *ERC20Token) Decimals() (uint8, error) {
	params := t.getParams()
	result, err := t.Call("decimals", params)
	if err != nil {
		return 0, err
	}
	data, ok := result.(uint8)
	if ok {
		t.decimals = data
	}
	return data, nil
}

// Name is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function name() constant returns(string)
func (t *ERC20Token) TotalSupply() (*big.Int, error) {
	var data = new(big.Int)
	var ok bool
	params := t.getParams()
	result, err := t.Call("totalSupply", params)
	if err != nil {
		return data, err
	}
	data, ok = result.(*big.Int)
	if ok {
		t.totalSupply = data
	}
	return data, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (t ERC20Token) Call(method string, params ...interface{}) (interface{}, error) {
	response, err := t.cli.Call(method, params)
	log.Debug(response)
	log.Debug(err)
	return nil, errors.New("not implemented")
}
