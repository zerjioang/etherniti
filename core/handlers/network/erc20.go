// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package network

import (
	"math/big"
	"net/http"
	"strconv"

	"github.com/zerjioang/etherniti/core/eth"
	"github.com/zerjioang/etherniti/core/util/str"

	"github.com/zerjioang/etherniti/core/eth/paramencoder"
	"github.com/zerjioang/etherniti/core/modules/encoding/hex"

	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/shared/protocol"

	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/server"

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

// get the total supply of the contract at given target network
func (ctl *Erc20Controller) name(c echo.Context) error {
	contractAddress := c.Param("contract")
	//input data validation
	if contractAddress == "" {
		return api.ErrorStr(c, "invalid contract address provided")
	}
	// cast to our context
	cc, ok := c.(*server.EthernitiContext)
	if !ok {
		return api.ErrorStr(c, "failed to execute requested operation")
	}
	// get our client context
	client, cId, cliErr := cc.RecoverEthClientFromTokenOrPeerUrl(ctl.network.peer)
	logger.Info("erc20 controller request using context id: ", cId)
	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	raw, err := client.Erc20Name(contractAddress)
	if err != nil {
		// send invalid generation message
		return c.JSONBlob(http.StatusBadRequest,
			str.GetJsonBytes(
				protocol.NewApiError(http.StatusBadRequest, err.Error()),
			),
		)
	} else {
		unpacked := ""
		rawBytes, decodeErr := hex.FromEthHex(string(raw))
		if decodeErr != nil {
			return api.ErrorStr(c, "failed to hex decode network response: "+decodeErr.Error())
		}
		err := paramencoder.LoadErc20Abi().Unpack(&unpacked, "name", rawBytes)
		if err != nil {
			return api.ErrorStr(c, "failed to decode network response: "+err.Error())
		} else {
			return api.SendSuccess(c, "name", unpacked)
		}
	}
}

// get the total supply of the contract at given target network
func (ctl *Erc20Controller) symbol(c echo.Context) error {
	contractAddress := c.Param("contract")
	//input data validation
	if contractAddress == "" {
		return api.ErrorStr(c, "invalid contract address provided")
	}
	// cast to our context
	cc, ok := c.(*server.EthernitiContext)
	if !ok {
		return api.ErrorStr(c, "failed to execute requested operation")
	}
	// get our client context
	client, cId, cliErr := cc.RecoverEthClientFromTokenOrPeerUrl(ctl.network.peer)
	logger.Info("erc20 controller request using context id: ", cId)
	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	raw, err := client.Erc20Symbol(contractAddress)
	if err != nil {
		// send invalid generation message
		return c.JSONBlob(http.StatusBadRequest,
			str.GetJsonBytes(
				protocol.NewApiError(http.StatusBadRequest, err.Error()),
			),
		)
	} else {
		unpacked := ""
		rawBytes, decodeErr := hex.FromEthHex(string(raw))
		if decodeErr != nil {
			return api.ErrorStr(c, "failed to hex decode network response: "+decodeErr.Error())
		}
		err := paramencoder.LoadErc20Abi().Unpack(&unpacked, "symbol", rawBytes)
		if err != nil {
			return api.ErrorStr(c, "failed to decode network response: "+err.Error())
		} else {
			return api.SendSuccess(c, "symbol", unpacked)
		}
	}
}

// get the total supply of the contract at given target network
func (ctl *Erc20Controller) totalSupply(c echo.Context) error {
	contractAddress := c.Param("contract")
	//input data validation
	if contractAddress == "" {
		return api.ErrorStr(c, "invalid contract address provided")
	}
	// cast to our context
	cc, ok := c.(*server.EthernitiContext)
	if !ok {
		return api.ErrorStr(c, "failed to execute requested operation")
	}
	// get our client context
	client, cId, cliErr := cc.RecoverEthClientFromTokenOrPeerUrl(ctl.network.peer)
	logger.Info("erc20 controller request using context id: ", cId)
	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	raw, err := client.Erc20TotalSupply(contractAddress)
	if err != nil {
		// send invalid generation message
		return c.JSONBlob(http.StatusBadRequest,
			str.GetJsonBytes(
				protocol.NewApiError(http.StatusBadRequest, err.Error()),
			),
		)
	} else {
		var unpacked *big.Int
		rawBytes, decodeErr := hex.FromEthHex(string(raw))
		if decodeErr != nil {
			return api.ErrorStr(c, "failed to hex decode network response: "+decodeErr.Error())
		}
		err := paramencoder.LoadErc20Abi().Unpack(&unpacked, "totalSupply", rawBytes)
		if err != nil {
			return api.ErrorStr(c, "failed to decode network response: "+err.Error())
		} else {
			return api.SendSuccess(c, "totalSupply", unpacked)
		}
	}
}

// get the total supply of the contract at given target network
func (ctl *Erc20Controller) decimals(c echo.Context) error {
	contractAddress := c.Param("contract")
	//input data validation
	if contractAddress == "" {
		return api.ErrorStr(c, "invalid contract address provided")
	}
	// cast to our context
	cc, ok := c.(*server.EthernitiContext)
	if !ok {
		return api.ErrorStr(c, "failed to execute requested operation")
	}
	// get our client context
	client, cId, cliErr := cc.RecoverEthClientFromTokenOrPeerUrl(ctl.network.peer)
	logger.Info("erc20 controller request using context id: ", cId)
	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	raw, err := client.Erc20Decimals(contractAddress)
	if err != nil {
		// send invalid generation message
		return c.JSONBlob(http.StatusBadRequest,
			str.GetJsonBytes(
				protocol.NewApiError(http.StatusBadRequest, err.Error()),
			),
		)
	} else {
		var unpacked *big.Int
		rawBytes, decodeErr := hex.FromEthHex(string(raw))
		if decodeErr != nil {
			return api.ErrorStr(c, "failed to hex decode network response: "+decodeErr.Error())
		}
		err := paramencoder.LoadErc20Abi().Unpack(&unpacked, "decimals", rawBytes)
		if err != nil {
			return api.ErrorStr(c, "failed to decode network response: "+err.Error())
		} else {
			return api.SendSuccess(c, "decimals", unpacked)
		}
	}
}

// get the total supply of the contract at given target network
func (ctl *Erc20Controller) balanceof(c echo.Context) error {
	contractAddress := c.Param("contract")
	//input data validation
	if contractAddress == "" {
		return api.ErrorStr(c, "invalid contract address provided")
	}
	address := c.Param("address")
	//input data validation
	if address == "" {
		return api.ErrorStr(c, "invalid account address provided")
	}
	// cast to our context
	cc, ok := c.(*server.EthernitiContext)
	if !ok {
		return api.ErrorStr(c, "failed to execute requested operation")
	}
	// get our client context
	client, cId, cliErr := cc.RecoverEthClientFromTokenOrPeerUrl(ctl.network.peer)
	logger.Info("erc20 controller request using context id: ", cId)
	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	raw, err := client.Erc20BalanceOf(contractAddress, address)
	if err != nil {
		// send invalid generation message
		return c.JSONBlob(http.StatusBadRequest,
			str.GetJsonBytes(
				protocol.NewApiError(http.StatusBadRequest, err.Error()),
			),
		)
	} else {
		return api.SendSuccess(c, "balanceof", raw)
	}
}

// get the summary of information of given erc20 contract
func (ctl *Erc20Controller) summary(c echo.Context) error {
	targetAddr := c.Param("address")
	// cast to our context
	cc, ok := c.(*server.EthernitiContext)
	if !ok {
		return api.ErrorStr(c, "failed to execute requested operation")
	}
	instance, err := eth.InstantiateToken(cc, targetAddr)
	if err == nil && instance != nil {
		//todo save token instance in memory

		//show token summary
		/*bal, err := instance.BalanceOf(&bind.CallOpts{}, ethAddr)
		if err != nil {
			log.Fatal(err)
		}

		name, err := instance.name(&bind.CallOpts{})
		if err != nil {
			log.Fatal(err)
		}

		symbol, err := instance.symbol(&bind.CallOpts{})
		if err != nil {
			log.Fatal(err)
		}

		decimals, err := instance.Decimals(&bind.CallOpts{})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("name: %s\n", name)         // "name: Golem Network"
		fmt.Printf("symbol: %s\n", symbol)     // "symbol: GNT"
		fmt.Printf("decimals: %v\n", decimals) // "decimals: 18"

		fmt.Printf("wei: %s\n", bal) // "wei: 74605500647408739782407023"

		fbal := new(big.Float)
		fbal.SetString(bal.String())
		value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))

		fmt.Printf("balance: %f", value) // "balance: 74605500.647409"
		*/
	}
	return nil
}

// get the allowance status of the contract at given target network
func (ctl *Erc20Controller) allowance(c echo.Context) error {
	contractAddress := c.Param("contract")
	//input data validation
	if contractAddress == "" {
		return api.ErrorStr(c, "invalid contract address provided")
	}
	ownerAddress := c.Param("owner")
	//input data validation
	if ownerAddress == "" {
		return api.ErrorStr(c, "invalid account owner address provided")
	}
	spenderAddress := c.Param("spender")
	//input data validation
	if spenderAddress == "" {
		return api.ErrorStr(c, "invalid account spender address provided")
	}
	// cast to our context
	cc, ok := c.(*server.EthernitiContext)
	if !ok {
		return api.ErrorStr(c, "failed to execute requested operation")
	}
	// get our client context
	client, cId, cliErr := cc.RecoverEthClientFromTokenOrPeerUrl(ctl.network.peer)
	logger.Info("erc20 controller request using context id: ", cId)
	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	raw, err := client.Erc20Allowance(contractAddress, ownerAddress, spenderAddress)
	if err != nil {
		// send invalid generation message
		return c.JSONBlob(http.StatusBadRequest,
			str.GetJsonBytes(
				protocol.NewApiError(http.StatusBadRequest, err.Error()),
			),
		)
	} else {
		return api.SendSuccess(c, "allowance", raw)
	}
}

// transfer(address to, uint tokens) public returns (bool success);
// ------------------------------------------------------------------------
// transfer the balance from token owner's account to `to` account
// - Owner's account must have sufficient balance to transfer
// - 0 value transfers are allowed
// ------------------------------------------------------------------------
func (ctl *Erc20Controller) transfer(c echo.Context) error {
	contractAddress := c.Param("contract")
	//input data validation
	if contractAddress == "" {
		return api.ErrorStr(c, "invalid contract address provided")
	}
	receiverAddress := c.Param("address")
	//input data validation
	if receiverAddress == "" {
		return api.ErrorStr(c, "invalid transfer receiver address provided")
	}
	amount := c.Param("amount")
	tokenAmount, pErr := strconv.Atoi(amount)
	//input data validation
	if amount == "" || pErr != nil {
		return api.ErrorStr(c, "invalid token amount value provided")
	}
	cc, ok := c.(*server.EthernitiContext)
	if !ok {
		return api.ErrorStr(c, "failed to execute requested operation")
	}
	// get our client context
	client, cId, cliErr := cc.RecoverEthClientFromTokenOrPeerUrl(ctl.network.peer)
	logger.Info("erc20 controller request using context id: ", cId)
	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	raw, err := client.Erc20Transfer(contractAddress, receiverAddress, tokenAmount)
	if err != nil {
		// send invalid generation message
		return c.JSONBlob(http.StatusBadRequest,
			str.GetJsonBytes(
				protocol.NewApiError(http.StatusBadRequest, err.Error()),
			),
		)
	} else {
		return api.SendSuccess(c, "allowance", raw)
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
func (ctl *Erc20Controller) Approve(c echo.Context) error {
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
func (ctl *Erc20Controller) TransferFrom(c echo.Context) error {
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
