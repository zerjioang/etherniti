// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/zerjioang/etherniti/core/api/protocol"
	"github.com/zerjioang/etherniti/core/eth/fixtures"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/server"
	"github.com/zerjioang/etherniti/core/util"

	"github.com/labstack/echo"
	"github.com/zerjioang/etherniti/core/eth"
)

var (
	ctx = context.Background()
)

// eth web3 controller
type Web3Controller struct {
	//ethereum interaction cache
	cache *cache.Cache
}

// constructor like function
func NewWeb3Controller() Web3Controller {
	ctl := Web3Controller{}
	ctl.cache = cache.New(5*time.Minute, 10*time.Minute)
	return ctl
}

// check if an ethereum address is a contract address
func (ctl Web3Controller) getBalance(c echo.Context) error {
	// cast to our context
	cc, ok := c.(*server.EthernitiContext)
	if !ok {
		return protocol.ErrorStr(c, "failed to execute requested operation")
	}

	clientInstance, err := cc.ClientInstance()
	if err != nil {
		// there was an trycatch recovering client instance
		apiErr := protocol.NewApiError(http.StatusBadRequest, err.Error())
		apiErrRaw := util.GetJsonBytes(apiErr)
		return c.JSONBlob(http.StatusBadRequest, apiErrRaw)
	}
	targetAddr := c.Param("address")
	// check if not empty
	if targetAddr != "" {
		result, err := clientInstance.EthGetBalance(targetAddr, "latest")
		if err != nil {
			//some trycatch happen, return trycatch to client
			apiErr := protocol.NewApiError(http.StatusBadRequest, err.Error())
			apiErrRaw := util.GetJsonBytes(apiErr)
			return c.JSONBlob(http.StatusBadRequest, apiErrRaw)
		}
		return c.JSONBlob(http.StatusOK, util.GetJsonBytes(result))
	}
	// send invalid address message
	return c.JSONBlob(http.StatusBadRequest, invalidAddressBytes)
}

// check if an ethereum address is a contract address
func (ctl Web3Controller) getBalanceAtBlock(c echo.Context) error {
	// cast to our context
	cc, ok := c.(*server.EthernitiContext)
	if !ok {
		return protocol.ErrorStr(c, "failed to execute requested operation")
	}

	clientInstance, err := cc.ClientInstance()
	if err != nil {
		// there was an trycatch recovering client instance
		apiErr := protocol.NewApiError(http.StatusBadRequest, err.Error())
		apiErrRaw := util.GetJsonBytes(apiErr)
		return c.JSONBlob(http.StatusBadRequest, apiErrRaw)
	}
	targetAddr := c.Param("address")
	block := c.Param("block")
	// check if not empty
	if targetAddr != "" {
		result, err := clientInstance.EthGetBalance(targetAddr, block)
		if err != nil {
			//some trycatch happen, return trycatch to client
			apiErr := protocol.NewApiError(http.StatusBadRequest, err.Error())
			apiErrRaw := util.GetJsonBytes(apiErr)
			return c.JSONBlob(http.StatusBadRequest, apiErrRaw)
		}
		return c.JSONBlob(http.StatusOK, util.GetJsonBytes(result))
	}
	// send invalid address message
	return c.JSONBlob(http.StatusBadRequest, invalidAddressBytes)
}

// get node information
func (ctl Web3Controller) getNodeInfo(c echo.Context) error {
	// cast to our context
	cc, ok := c.(*server.EthernitiContext)
	if !ok {
		return protocol.ErrorStr(c, "failed to execute requested operation")
	}

	clientInstance, err := cc.ClientInstance()
	if err != nil {
		// there was an trycatch recovering client instance
		apiErr := protocol.NewApiError(http.StatusBadRequest, err.Error())
		apiErrRaw := util.GetJsonBytes(apiErr)
		return c.JSONBlob(http.StatusBadRequest, apiErrRaw)
	}
	data, err := clientInstance.EthNodeInfo()
	if err != nil {
		// send invalid address message
		return c.JSONBlob(http.StatusBadRequest, invalidAddressBytes)
	} else {
		return protocol.Success(c, "eth_info", data)
	}
}

func (ctl Web3Controller) getAccounts(c echo.Context) error {
	// cast to our context
	cc, ok := c.(*server.EthernitiContext)
	if !ok {
		return protocol.ErrorStr(c, "failed to execute requested operation")
	}

	client, cliErr := cc.ClientInstance()
	if cliErr != nil {
		return protocol.Error(c, cliErr)
	}
	list, err := client.EthAccounts()
	if err != nil {
		// send invalid generation message
		return c.JSONBlob(http.StatusInternalServerError,
			util.GetJsonBytes(
				protocol.NewApiError(http.StatusInternalServerError, err.Error()),
			),
		)
	} else {
		return c.JSONBlob(
			http.StatusOK,
			util.GetJsonBytes(
				protocol.NewApiResponse("ethereum accounts readed", list),
			),
		)
	}
}

/*
{
  "jsonrpc": "2.0",
  "method": "eth_getBalance",
  "params": ["0x0ADfCCa4B2a1132F82488546AcA086D7E24EA324", "latest"],
  "id": 1
}
*/
func (ctl Web3Controller) getAccountsWithBalance(c echo.Context) error {

	// cast to our context
	cc, ok := c.(*server.EthernitiContext)
	if !ok {
		return protocol.ErrorStr(c, "failed to execute requested operation")
	}

	client, cliErr := cc.ClientInstance()
	if cliErr != nil {
		return protocol.Error(c, cliErr)
	}
	list, err := client.EthAccounts()

	type wrapper struct {
		Account string `json:"account"`
		Balance string `json:"balance"`
		Eth     string `json:"eth"`
		Key     string `json:"key"`
	}
	wrapperList := make([]wrapper, len(list))

	if err != nil {
		// send invalid generation message
		return c.JSONBlob(http.StatusInternalServerError,
			util.GetJsonBytes(
				protocol.NewApiError(http.StatusInternalServerError, err.Error()),
			),
		)
	} else {
		//iterate over account
		for i := 0; i < len(list); i++ {
			currentAccount := list[i]
			bigInt, err := client.EthGetBalance(currentAccount, "latest")
			if err != nil {
				logger.Error("failed to get account balance", currentAccount, err)
			} else {
				item := &wrapperList[i]
				item.Account = currentAccount
				item.Balance = bigInt.String()
				item.Eth = fixtures.ToEth(bigInt).String()
				item.Key = "secret"
			}
		}
		return c.JSONBlob(
			http.StatusOK,
			util.GetJsonBytes(
				protocol.NewApiResponse("ethereum accounts and their balance readed", wrapperList),
			),
		)
	}
}

func (ctl Web3Controller) getBlocks(c echo.Context) error {
	return nil
}

func (ctl Web3Controller) coinbase(c echo.Context) error {

	// cast to our context
	cc, ok := c.(*server.EthernitiContext)
	if !ok {
		return protocol.ErrorStr(c, "failed to execute requested operation")
	}

	raw, found := ctl.cache.Get("eth_coinbase")
	if found && raw != nil {
		//cache hit
		//return result to client
		return c.JSONBlob(
			http.StatusOK,
			util.GetJsonBytes(
				protocol.NewApiResponse("coinbase address", raw),
			),
		)
	} else {
		//cache miss
		client, cliErr := cc.ClientInstance()
		if cliErr != nil {
			return protocol.Error(c, cliErr)
		}
		result, err := client.Call("eth_coinbase")
		if err == nil {
			if result != nil {
				// add result to cache
				ctl.cache.Set("eth_coinbase", result, cache.DefaultExpiration)
				//return result to client
				return c.JSONBlob(
					http.StatusOK,
					util.GetJsonBytes(
						protocol.NewApiResponse("coinbase address", result),
					),
				)
			} else {
				return c.JSONBlob(http.StatusBadRequest,
					util.GetJsonBytes(
						protocol.NewApiError(http.StatusBadRequest, "empty response from server"),
					),
				)
			}
		} else {
			return c.JSONBlob(http.StatusBadRequest,
				util.GetJsonBytes(
					protocol.NewApiError(http.StatusBadRequest, "failed to get coinbase address: "+err.Error()),
				),
			)
		}
	}
}

// check if an ethereum address is a contract address
func (ctl Web3Controller) isContractAddress(c echo.Context) error {
	// cast to our context
	cc, ok := c.(*server.EthernitiContext)
	if !ok {
		return protocol.ErrorStr(c, "failed to execute requested operation")
	}
	clientInstance, err := cc.ClientInstance()
	if err != nil {
		// there was an error recovering client instance
		apiErr := protocol.NewApiError(http.StatusBadRequest, err.Error())
		apiErrRaw := util.GetJsonBytes(apiErr)
		return c.JSONBlob(http.StatusBadRequest, apiErrRaw)
	}
	targetAddr := c.Param("address")
	// check if not empty
	if targetAddr != "" {
		result, err := eth.IsSmartContractAddress(clientInstance, targetAddr)
		if err != nil {
			//some error happen, return error to client
			apiErr := protocol.NewApiError(http.StatusBadRequest, err.Error())
			apiErrRaw := util.GetJsonBytes(apiErr)
			return c.JSONBlob(http.StatusBadRequest, apiErrRaw)
		}
		return c.JSONBlob(http.StatusOK, util.GetJsonBytes(result))
	}
	// send invalid address message
	return c.JSONBlob(http.StatusBadRequest, invalidAddressBytes)
}

// implemented method from interface RouterRegistrable
func (ctl Web3Controller) RegisterRouters(router *echo.Group) {
	router.GET("/eth/node/info", ctl.getNodeInfo)

	router.GET("/eth/is/contract/:address", ctl.isContractAddress)

	router.GET("/eth/accounts", ctl.getAccounts)
	router.GET("/eth/accountsBalanced", ctl.getAccountsWithBalance)

	router.GET("/eth/blocks", ctl.getBlocks)

	router.GET("/eth/coinbase", ctl.coinbase)

	router.GET("/eth/getbalance/:address", ctl.getBalance)
	router.GET("/eth/getbalance/:address/block/:block", ctl.getBalanceAtBlock)
}
