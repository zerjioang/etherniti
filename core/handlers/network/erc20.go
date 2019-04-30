// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package network

import (
	"math/big"
	"net/http"
	"strconv"

	"github.com/zerjioang/etherniti/core/data"
	"github.com/zerjioang/etherniti/core/util/str"

	"github.com/zerjioang/etherniti/core/eth/paramencoder"
	"github.com/zerjioang/etherniti/core/modules/encoding/hex"

	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

// eth erc20 controller
type Erc20Controller struct {
	network *NetworkController
}

// constructor like function
func NewErc20Controller(network *NetworkController) Erc20Controller {
	ctl := Erc20Controller{}
	ctl.network = network
	return ctl
}

// generic method that executes queries against erc20 contract
func (ctl *Erc20Controller) queryContract(c *echo.Context, methodName string, f func(contract string) (string, error), unpacked interface{}) error {
	contractAddress := c.Param("contract")
	//input data validation
	if contractAddress == "" {
		return api.ErrorStr(c, data.InvalidContractAddress)
	}
	raw, err := f(contractAddress)
	if err != nil {
		// send invalid generation message
		return api.ErrorCode(c, http.StatusBadRequest, err)
	} else {
		rawBytes, decodeErr := hex.FromEthHex(raw)
		if decodeErr != nil {
			return api.ErrorStr(c, str.UnsafeBytes("failed to hex decode network response: "+decodeErr.Error()))
		}
		err := paramencoder.LoadErc20Abi().Unpack(&unpacked, methodName, rawBytes)
		if err != nil {
			return api.ErrorStr(c, str.UnsafeBytes("failed to decode network response: "+err.Error()))
		} else {
			return api.SendSuccess(c, str.UnsafeBytes(methodName), unpacked)
		}
	}
}

// get the total supply of the contract at given target network
func (ctl *Erc20Controller) name(c *echo.Context) error {
	rpcClient, err := ctl.network.getRpcClient(c)
	if err == nil {
		return ctl.queryContract(c, "name", rpcClient.Erc20Name, new(string))
	}
	return nil
}

// get the total supply of the contract at given target network
func (ctl *Erc20Controller) symbol(c *echo.Context) error {
	rpcClient, err := ctl.network.getRpcClient(c)
	if err == nil {
		return ctl.queryContract(c, "symbol", rpcClient.Erc20Symbol, new(string))
	}
	return nil
}

// get the total supply of the contract at given target network
func (ctl *Erc20Controller) totalSupply(c *echo.Context) error {
	rpcClient, err := ctl.network.getRpcClient(c)
	if err == nil {
		var unpacked *big.Int
		return ctl.queryContract(c, "totalSupply", rpcClient.Erc20TotalSupply, &unpacked)
	}
	return nil
}

// get the total supply of the contract at given target network
func (ctl *Erc20Controller) decimals(c *echo.Context) error {
	rpcClient, err := ctl.network.getRpcClient(c)
	if err == nil {
		var unpacked *big.Int
		return ctl.queryContract(c, "decimals", rpcClient.Erc20Decimals, &unpacked)
	}
	return nil
}

// get the total supply of the contract at given target network
func (ctl *Erc20Controller) balanceof(c *echo.Context) error {
	contractAddress := c.Param("contract")
	//input data validation
	if contractAddress == "" {
		return api.ErrorStr(c, data.InvalidContractAddress)
	}
	address := c.Param("address")
	//input data validation
	if address == "" {
		return api.ErrorStr(c, data.InvalidAccountAddress)
	}
	client, cliErr := ctl.network.getRpcClient(c)
	if cliErr != nil {
		return nil
	} else {
		raw, err := client.Erc20BalanceOf(contractAddress, address)
		if err != nil {
			// send invalid generation message
			return api.ErrorCode(c, http.StatusBadRequest, err)
		} else {
			var unpacked *big.Int
			rawBytes, decodeErr := hex.FromEthHex(string(raw))
			if decodeErr != nil {
				return api.ErrorStr(c, str.UnsafeBytes("failed to hex decode network response: "+decodeErr.Error()))
			}
			err := paramencoder.LoadErc20Abi().Unpack(&unpacked, "decimals", rawBytes)
			if err != nil {
				return api.ErrorStr(c, str.UnsafeBytes("failed to decode network response: "+err.Error()))
			} else {
				return api.SendSuccess(c, data.BalanceOf, unpacked)
			}
		}
	}
}

// get the summary of information of given erc20 contract at given target network
func (ctl *Erc20Controller) summary(c *echo.Context) error {
	contractAddress := c.Param("contract")
	//input data validation
	if contractAddress == "" {
		return api.ErrorStr(c, data.InvalidContractAddress)
	}
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)
	logger.Info("erc20 controller request using context id: ", ctl.network.networkName)
	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	raw, err := client.Erc20Summary(contractAddress)
	if err != nil {
		// send invalid generation message
		return api.ErrorCode(c, http.StatusBadRequest, err)
	} else {
		if err != nil {
			return api.ErrorStr(c, str.UnsafeBytes("failed to decode network response: "+err.Error()))
		} else {
			return api.SendSuccess(c, data.Summary, raw)
		}
	}
}

// get the allowance status of the contract at given target network
func (ctl *Erc20Controller) allowance(c *echo.Context) error {
	contractAddress := c.Param("contract")
	//input data validation
	if contractAddress == "" {
		return api.ErrorStr(c, data.InvalidContractAddress)
	}
	ownerAddress := c.Param("owner")
	//input data validation
	if ownerAddress == "" {
		return api.ErrorStr(c, data.InvalidAccountOwner)
	}
	spenderAddress := c.Param("spender")
	//input data validation
	if spenderAddress == "" {
		return api.ErrorStr(c, data.InvalidAccountSpender)
	}
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)
	logger.Info("erc20 controller request using context id: ", ctl.network.networkName)
	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	raw, err := client.Erc20Allowance(contractAddress, ownerAddress, spenderAddress)
	if err != nil {
		// send invalid generation message
		return api.ErrorCode(c, http.StatusBadRequest, err)
	} else {
		return api.SendSuccess(c, data.Allowance, raw)
	}
}

// transfer(address to, uint tokens) public returns (bool success);
// ------------------------------------------------------------------------
// transfer the balance from token owner's account to `to` account
// - Owner's account must have sufficient balance to transfer
// - 0 value transfers are allowed
// ------------------------------------------------------------------------
func (ctl *Erc20Controller) transfer(c *echo.Context) error {
	contractAddress := c.Param("contract")
	//input data validation
	if contractAddress == "" {
		return api.ErrorStr(c, data.InvalidContractAddress)
	}
	receiverAddress := c.Param("address")
	//input data validation
	if receiverAddress == "" {
		return api.ErrorStr(c, data.InvalidReceiverAddress)
	}
	amount := c.Param("amount")
	tokenAmount, pErr := strconv.Atoi(amount)
	//input data validation
	if amount == "" || pErr != nil {
		return api.ErrorStr(c, data.InvalidTokenValue)
	}
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)
	logger.Info("erc20 controller request using context id: ", ctl.network.networkName)
	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	raw, err := client.Erc20Transfer(contractAddress, receiverAddress, tokenAmount)
	if err != nil {
		// send invalid generation message
		return api.ErrorCode(c, http.StatusBadRequest, err)
	} else {
		return api.SendSuccess(c, data.Allowance, raw)
	}
}

//approve(address spender, uint tokens) public returns (bool success);
// ------------------------------------------------------------------------
// Token owner can approve for `spender` to transferFrom(...) `tokens`
// from the token owner's account
//
// https://github.com/ethereum/EIPs/blob/master/EIPS/eip-20-token-standard.md
// recommends that there are no checks for the approval double-spend attack
// as this should be implemented in user interfaces
// ------------------------------------------------------------------------
func (ctl *Erc20Controller) Approve(c *echo.Context) error {
	return nil
}

//transferFrom(address from, address to, uint tokens) public returns (bool success);
// ------------------------------------------------------------------------
// transfer `tokens` from the `from` account to the `to` account
//
// The calling account must already have sufficient tokens approve(...)-d
// for spending from the `from` account and
// - From account must have sufficient balance to transfer
// - Spender must have sufficient allowance to transfer
// - 0 value transfers are allowed
// ------------------------------------------------------------------------
func (ctl *Erc20Controller) TransferFrom(c *echo.Context) error {
	return nil
}

// END of ERC20 functions

// implemented method from interface RouterRegistrable
func (ctl Erc20Controller) RegisterRouters(router *echo.Group) {
	router.GET("/erc20/:contract/summary", ctl.summary)
	router.GET("/erc20/:contract/name", ctl.name)
	router.GET("/erc20/:contract/symbol", ctl.symbol)
	router.GET("/erc20/:contract/totalsupply", ctl.totalSupply)
	router.GET("/erc20/:contract/decimals", ctl.decimals)
	router.GET("/erc20/:contract/balanceof/:address", ctl.balanceof)
	router.GET("/erc20/:contract/allowance/:owner/to/:spender", ctl.allowance)
	router.GET("/erc20/:contract/transfer/:address/:amount", ctl.transfer)
}
