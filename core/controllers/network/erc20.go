// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package network

import (
	"math/big"
	"strconv"

	"github.com/zerjioang/etherniti/core/controllers/wrap"
	"github.com/zerjioang/etherniti/shared"

	"github.com/zerjioang/go-hpc/lib/codes"
	"github.com/zerjioang/go-hpc/lib/eth/paramencoder/erc20"

	"github.com/zerjioang/etherniti/core/data"
	"github.com/zerjioang/go-hpc/util/str"

	"github.com/zerjioang/go-hpc/lib/encoding/hex"

	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/go-hpc/thirdparty/echo"
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
func (ctl *Erc20Controller) queryContract(c *shared.EthernitiContext, methodName string, f func(contract string) (string, error), unpacked interface{}) error {
	contractAddress := c.Param("contract")
	//input data validation
	if contractAddress == "" {
		return api.ErrorBytes(c, data.InvalidContractAddress)
	}
	raw, err := f(contractAddress)
	if err != nil {
		// send invalid generation message
		return api.ErrorCode(c, codes.StatusBadRequest, err)
	} else {
		rawBytes, decodeErr := hex.FromEthHex(raw)
		if decodeErr != nil {
			return api.ErrorBytes(c, str.UnsafeBytes("failed to hex decode network response: "+decodeErr.Error()))
		}
		err := erc20.LoadErc20Abi().Unpack(&unpacked, methodName, rawBytes)
		if err != nil {
			return api.ErrorBytes(c, str.UnsafeBytes("failed to decode network response: "+err.Error()))
		} else {
			return api.SendSuccess(c, str.UnsafeBytes(methodName), unpacked)
		}
	}
}

// get the total supply of the contract at given target network
func (ctl *Erc20Controller) name(c *shared.EthernitiContext) error {
	rpcClient, err := ctl.network.getRpcClient(c)
	if err == nil {
		return ctl.queryContract(c, "name", rpcClient.Erc20Name, new(string))
	}
	return nil
}

// get the total supply of the contract at given target network
func (ctl *Erc20Controller) symbol(c *shared.EthernitiContext) error {
	rpcClient, err := ctl.network.getRpcClient(c)
	if err == nil {
		return ctl.queryContract(c, "symbol", rpcClient.Erc20Symbol, new(string))
	}
	return nil
}

// get the total supply of the contract at given target network
func (ctl *Erc20Controller) totalSupply(c *shared.EthernitiContext) error {
	rpcClient, err := ctl.network.getRpcClient(c)
	if err == nil {
		var unpacked *big.Int
		return ctl.queryContract(c, "totalSupply", rpcClient.Erc20TotalSupply, &unpacked)
	}
	return nil
}

// get the total supply of the contract at given target network
func (ctl *Erc20Controller) decimals(c *shared.EthernitiContext) error {
	rpcClient, err := ctl.network.getRpcClient(c)
	if err == nil {
		var unpacked *big.Int
		return ctl.queryContract(c, "decimals", rpcClient.Erc20Decimals, &unpacked)
	}
	return nil
}

// get the total supply of the contract at given target network
func (ctl *Erc20Controller) balanceof(c *shared.EthernitiContext) error {
	contractAddress := c.Param("contract")
	//input data validation
	if contractAddress == "" {
		return api.ErrorBytes(c, data.InvalidContractAddress)
	}
	address := c.Param("address")
	//input data validation
	if address == "" {
		return api.ErrorBytes(c, data.InvalidAccountAddress)
	}
	client, cliErr := ctl.network.getRpcClient(c)
	if cliErr != nil {
		return nil
	} else {
		raw, err := client.Erc20BalanceOf(contractAddress, address)
		if err != nil {
			// send invalid generation message
			return api.ErrorCode(c, codes.StatusBadRequest, err)
		} else {
			var unpacked *big.Int
			rawBytes, decodeErr := hex.FromEthHex(string(raw))
			if decodeErr != nil {
				return api.ErrorBytes(c, str.UnsafeBytes("failed to hex decode network response: "+decodeErr.Error()))
			}
			err := erc20.LoadErc20Abi().Unpack(&unpacked, "decimals", rawBytes)
			if err != nil {
				return api.ErrorBytes(c, str.UnsafeBytes("failed to decode network response: "+err.Error()))
			} else {
				return api.SendSuccess(c, data.BalanceOf, unpacked)
			}
		}
	}
}

// get the summary of information of given erc20 contract at given target network
func (ctl *Erc20Controller) summary(c *shared.EthernitiContext) error {
	contractAddress := c.Param("contract")
	//input data validation
	if contractAddress == "" {
		return api.ErrorBytes(c, data.InvalidContractAddress)
	}
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)
	logger.Info("erc20 controller request using context id: ", ctl.network.Name())
	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	raw, err := client.Erc20Summary(contractAddress)
	if err != nil {
		// send invalid generation message
		return api.ErrorCode(c, codes.StatusBadRequest, err)
	} else {
		if err != nil {
			return api.ErrorBytes(c, str.UnsafeBytes("failed to decode network response: "+err.Error()))
		} else {
			return api.SendSuccess(c, data.Summary, raw)
		}
	}
}

// get the allowance status of the contract at given target network
func (ctl *Erc20Controller) allowance(c *shared.EthernitiContext) error {
	contractAddress := c.Param("contract")
	//input data validation
	if contractAddress == "" {
		return api.ErrorBytes(c, data.InvalidContractAddress)
	}
	ownerAddress := c.Param("owner")
	//input data validation
	if ownerAddress == "" {
		return api.ErrorBytes(c, data.InvalidAccountOwner)
	}
	spenderAddress := c.Param("spender")
	//input data validation
	if spenderAddress == "" {
		return api.ErrorBytes(c, data.InvalidAccountSpender)
	}
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)
	logger.Info("erc20 controller request using context id: ", ctl.network.Name())
	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	raw, err := client.Erc20Allowance(contractAddress, ownerAddress, spenderAddress)
	if err != nil {
		// send invalid generation message
		return api.ErrorCode(c, codes.StatusBadRequest, err)
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
func (ctl *Erc20Controller) transfer(c *shared.EthernitiContext) error {
	contractAddress := c.Param("contract")
	//input data validation
	if contractAddress == "" {
		return api.ErrorBytes(c, data.InvalidContractAddress)
	}
	receiverAddress := c.Param("address")
	//input data validation
	if receiverAddress == "" {
		return api.ErrorBytes(c, data.InvalidReceiverAddress)
	}
	amount := c.Param("amount")
	tokenAmount, pErr := strconv.Atoi(amount)
	//input data validation
	if amount == "" || pErr != nil {
		return api.ErrorBytes(c, data.InvalidTokenValue)
	}
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)
	logger.Info("erc20 controller request using context id: ", ctl.network.Name())
	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	raw, err := client.Erc20Transfer(contractAddress, receiverAddress, tokenAmount)
	if err != nil {
		// send invalid generation message
		return api.ErrorCode(c, codes.StatusBadRequest, err)
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
func (ctl *Erc20Controller) Approve(c *shared.EthernitiContext) error {
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
func (ctl *Erc20Controller) TransferFrom(c *shared.EthernitiContext) error {
	return nil
}

// END of ERC20 functions

// implemented method from interface RouterRegistrable
func (ctl Erc20Controller) RegisterRouters(router *echo.Group) {
	router.GET("/erc20/:contract/summary", wrap.Call(ctl.summary))
	router.GET("/erc20/:contract/name", wrap.Call(ctl.name))
	router.GET("/erc20/:contract/symbol", wrap.Call(ctl.symbol))
	router.GET("/erc20/:contract/totalsupply", wrap.Call(ctl.totalSupply))
	router.GET("/erc20/:contract/decimals", wrap.Call(ctl.decimals))
	router.GET("/erc20/:contract/balanceof/:address", wrap.Call(ctl.balanceof))
	router.GET("/erc20/:contract/allowance/:owner/to/:spender", wrap.Call(ctl.allowance))
	router.GET("/erc20/:contract/transfer/:address/:amount", wrap.Call(ctl.transfer))
}
