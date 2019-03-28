// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"net/http"
	"strings"

	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/shared/protocol"

	"github.com/zerjioang/etherniti/core/eth/fixtures"
	"github.com/zerjioang/etherniti/core/handlers/clientcache"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/server"
	"github.com/zerjioang/etherniti/core/util"

	"github.com/labstack/echo"
	"github.com/zerjioang/etherniti/core/eth"
)

var (
	methodMap = map[string]string{
		"client_version":   "web3_clientVersion",
		"net_version":      "net_version",
		"net_peers":        "net_peerCount",
		"protocol_version": "eth_protocolVersion",
		"syncing":          "eth_syncing",
		"coinbase":         "eth_coinbase",
		"mining":           "eth_mining",
		"hashrate":         "eth_hashrate",
		"gasprice":         "eth_gasPrice",
		"accounts":         "eth_accounts",
		"block_latest":     "eth_blockNumber",
		"compilers":        "eth_getCompilers",
		"block_current":    "eth_getWork",
		"shh_version":      "shh_version",
		"shh_new":          "shh_newIdentity",
		"shh_group":        "shh_newGroup",
	}
)

// eth web3 controller
type Web3Controller struct {
	NetworkController
}

// constructor like function
func NewWeb3Controller() Web3Controller {
	ctl := Web3Controller{}
	ctl.NetworkController = NewNetworkController()
	return ctl
}

// check if an ethereum address is a contract address
func (ctl *Web3Controller) getBalance(c echo.Context) error {
	// cast to our context
	cc, ok := c.(*server.EthernitiContext)
	if !ok {
		return api.ErrorStr(c, "failed to execute requested operation")
	}

	clientInstance, cId, err := cc.RecoverEthClientFromTokenOrPeerUrl(ctl.peer)
	logger.Info("web3 request using context id: ", cId)
	if err != nil {
		return api.Error(c, err)
	}
	targetAddr := c.Param("address")
	// check if not empty
	if targetAddr != "" {
		result, err := clientInstance.EthGetBalance(targetAddr, "latest")
		if err != nil {
			return api.Error(c, err)
		}
		return c.JSONBlob(http.StatusOK, util.GetJsonBytes(result))
	}
	// send invalid address message
	return c.JSONBlob(http.StatusBadRequest, invalidAddressBytes)
}

// check if an ethereum address is a contract address
func (ctl *Web3Controller) getBalanceAtBlock(c echo.Context) error {
	// cast to our context
	cc, ok := c.(*server.EthernitiContext)
	if !ok {
		return api.ErrorStr(c, "failed to execute requested operation")
	}

	clientInstance, cId, err := cc.RecoverEthClientFromTokenOrPeerUrl(ctl.peer)
	logger.Info("web3 request using context id: ", cId)
	if err != nil {
		return api.Error(c, err)
	}
	targetAddr := c.Param("address")
	block := c.Param("block")
	// check if not empty
	if targetAddr != "" {
		result, err := clientInstance.EthGetBalance(targetAddr, block)
		if err != nil {
			return api.Error(c, err)
		}
		return c.JSONBlob(http.StatusOK, util.GetJsonBytes(result))
	}
	// send invalid address message
	return c.JSONBlob(http.StatusBadRequest, invalidAddressBytes)
}

// get node information
func (ctl *Web3Controller) getNodeInfo(c echo.Context) error {
	// cast to our context
	cc, ok := c.(*server.EthernitiContext)
	if !ok {
		return api.ErrorStr(c, "failed to execute requested operation")
	}

	clientInstance, cId, err := cc.RecoverEthClientFromTokenOrPeerUrl(ctl.peer)
	logger.Info("web3 request using context id: ", cId)
	if err != nil {
		// there was an error recovering client instance
		return api.Error(c, err)
	}
	data, err := clientInstance.EthNodeInfo()
	if err != nil {
		// send invalid response message
		return api.Error(c, err)
	} else {
		return api.Success(c, "eth_info", data)
	}
}

// get node information
func (ctl *Web3Controller) getNetworkVersion(c echo.Context) error {
	// cast to our context
	cc, ok := c.(*server.EthernitiContext)
	if !ok {
		return api.ErrorStr(c, "failed to execute requested operation")
	}

	//try to get this information from the cache
	key := "net_version"
	result, found := ctl.cache.Get(key)
	if found && result != nil {
		//cache hit
		logger.Info(key, ": cache hit")
		response := api.ToSuccess(key, result.(string))
		return clientcache.CachedJsonBlob(c, true, clientcache.CacheInfinite, response)
	} else {
		//cache miss
		logger.Info(key, ": cache miss")
		clientInstance, cId, err := cc.RecoverEthClientFromTokenOrPeerUrl(ctl.peer)
		logger.Info("web3 request using context id: ", cId)
		if err != nil {
			// there was an error recovering client instance
			return api.Error(c, err)
		}
		data, err := clientInstance.EthNetVersion()
		if err != nil {
			// send invalid response message
			return api.Error(c, err)
		} else {
			// save result in the cache
			ctl.cache.Set(key, data, clientcache.NoExpiration)
			response := api.ToSuccess(key, data)
			return clientcache.CachedJsonBlob(c, true, clientcache.CacheInfinite, response)
		}
	}
}

func (ctl *Web3Controller) makeRpcCallNoParams(c echo.Context) error {
	// cast to our context
	cc, ok := c.(*server.EthernitiContext)
	if !ok {
		return api.ErrorStr(c, "failed to execute requested operation")
	}

	clientInstance, cId, err := cc.RecoverEthClientFromTokenOrPeerUrl(ctl.peer)
	logger.Info("web3 request using context id: ", cId)
	if err != nil {
		// there was an error recovering client instance
		return api.Error(c, err)
	}

	//resolve method name from url
	methodName := c.Request().URL.Path
	//try to get this information from the cache
	// methodName example: /v1/public/ropsten/net/version
	chunks := strings.Split(methodName, "/")
	if len(chunks) < 5 {
		return api.ErrorStr(c, "invalid url or web3 method provided")
	}
	var key string
	if len(chunks) == 5 {
		key = chunks[4]
	} else if len(chunks) == 6 {
		key = chunks[4] + "_" + chunks[5]
	}
	//resolve method name from key value
	method := methodMap[key]
	cacheKey := cId + ":" + method
	result, found := ctl.cache.Get(cacheKey)
	if found && result != nil {
		//cache hit
		logger.Info(method, ": cache hit")
		response := api.ToSuccess(method, result)
		return clientcache.CachedJsonBlob(c, true, clientcache.CacheInfinite, response)
	} else {
		//cache miss
		logger.Info(method, ": cache miss")
		rpcResponse, err := clientInstance.EthMethodNoParams(method)
		if err != nil {
			// send invalid response message
			return api.Error(c, err)
		} else if rpcResponse == nil {
			// send invalid response message
			return api.ErrorStr(c, "the network peer did not return any response")
		} else {
			// save result in the cache
			ctl.cache.Set(cacheKey, rpcResponse, clientcache.NoExpiration)
			response := api.ToSuccess(method, rpcResponse)
			return clientcache.CachedJsonBlob(c, true, clientcache.CacheInfinite, response)
		}
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
func (ctl *Web3Controller) getAccountsWithBalance(c echo.Context) error {

	// cast to our context
	cc, ok := c.(*server.EthernitiContext)
	if !ok {
		return api.ErrorStr(c, "failed to execute requested operation")
	}

	client, cId, cliErr := cc.RecoverEthClientFromTokenOrPeerUrl(ctl.peer)
	logger.Info("web3 request using context id: ", cId)
	if cliErr != nil {
		return api.Error(c, cliErr)
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

// check if an ethereum address is a contract address
func (ctl *Web3Controller) isContractAddress(c echo.Context) error {
	// cast to our context
	cc, ok := c.(*server.EthernitiContext)
	if !ok {
		return api.ErrorStr(c, "failed to execute requested operation")
	}
	clientInstance, cId, err := cc.RecoverEthClientFromTokenOrPeerUrl(ctl.peer)
	logger.Info("web3 request using context id: ", cId)
	if err != nil {
		return api.Error(c, err)
	}
	targetAddr := c.Param("address")
	// check if not empty
	if targetAddr != "" {
		result, err := eth.IsSmartContractAddress(clientInstance, targetAddr)
		if err != nil {
			return api.Error(c, err)
		}
		return c.JSONBlob(http.StatusOK, util.GetJsonBytes(result))
	}
	// send invalid address message
	return c.JSONBlob(http.StatusBadRequest, invalidAddressBytes)
}

// implemented method from interface RouterRegistrable
func (ctl Web3Controller) RegisterRouters(router *echo.Group) {
	router.GET("/"+ctl.networkName+"/client/version", ctl.makeRpcCallNoParams)
	router.GET("/"+ctl.networkName+"/net/version", ctl.makeRpcCallNoParams)
	router.GET("/"+ctl.networkName+"/net/peers", ctl.makeRpcCallNoParams)
	router.GET("/"+ctl.networkName+"/protocol/version", ctl.makeRpcCallNoParams)
	router.GET("/"+ctl.networkName+"/syncing", ctl.makeRpcCallNoParams)
	router.GET("/"+ctl.networkName+"/coinbase", ctl.makeRpcCallNoParams)
	router.GET("/"+ctl.networkName+"/mining", ctl.makeRpcCallNoParams)
	router.GET("/"+ctl.networkName+"/hashrate", ctl.makeRpcCallNoParams)
	router.GET("/"+ctl.networkName+"/gasprice", ctl.makeRpcCallNoParams)
	router.GET("/"+ctl.networkName+"/accounts", ctl.makeRpcCallNoParams)
	router.GET("/"+ctl.networkName+"/block/latest", ctl.makeRpcCallNoParams)
	router.GET("/"+ctl.networkName+"/block/current", ctl.makeRpcCallNoParams)
	router.GET("/"+ctl.networkName+"/compilers", ctl.makeRpcCallNoParams)
	router.GET("/"+ctl.networkName+"/shh/version", ctl.makeRpcCallNoParams)
	router.GET("/"+ctl.networkName+"/shh/new", ctl.makeRpcCallNoParams)
	router.GET("/"+ctl.networkName+"/shh/group", ctl.makeRpcCallNoParams)

	router.GET("/"+ctl.networkName+"/is/contract/:address", ctl.isContractAddress)

	router.GET("/"+ctl.networkName+"/accountsBalanced", ctl.getAccountsWithBalance)

	router.GET("/"+ctl.networkName+"/balance/:address", ctl.getBalance)
	router.GET("/"+ctl.networkName+"/balance/:address/block/:block", ctl.getBalanceAtBlock)
}
