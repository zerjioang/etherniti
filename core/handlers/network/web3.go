// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package network

import (
	"math/big"
	"net/http"
	"strconv"
	"strings"

	"github.com/zerjioang/etherniti/core/eth"

	"github.com/zerjioang/etherniti/core/eth/rpc"

	"github.com/zerjioang/etherniti/core/util/str"

	"github.com/zerjioang/etherniti/core/eth/paramencoder"
	"github.com/zerjioang/etherniti/core/modules/encoding/hex"

	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/shared/protocol"

	"github.com/zerjioang/etherniti/core/eth/fixtures"
	"github.com/zerjioang/etherniti/core/handlers/clientcache"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/thirdparty/echo"
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
	network *NetworkController
}

// constructor like function
func NewWeb3Controller(network *NetworkController) Web3Controller {
	ctl := Web3Controller{}
	ctl.network = network
	return ctl
}

// get the balance of given ethereum address and target network
func (ctl *Web3Controller) getBalance(c echo.ContextInterface) error {
	targetAddr := c.Param("address")
	// check if not empty
	if targetAddr != "" {
		//try to get this information from the cache
		key := ctl.network.peer+"get_balance"+targetAddr
		result, found := ctl.network.cache.Get(key)
		if found && result != nil {
			//cache hit
			logger.Info(key, ": cache hit")
			return clientcache.CachedJsonBlob(c, true, clientcache.CacheInfinite, result.([]byte))
		} else {
			//cache miss
			logger.Info(key, ": cache miss")
			// get our client context
			client, cliErr := ctl.network.getRpcClient(c)
			logger.Info("web3 controller request using context id: ", ctl.network.networkName)
			if cliErr != nil {
				return api.Error(c, cliErr)
			}
			result, err := client.EthGetBalance(targetAddr, "latest")
			if err != nil {
				return api.Error(c, err)
			}
			// save result in the cache
			response := api.ToSuccess("balance", result)
			ctl.network.cache.Set(key, response)
			return clientcache.CachedJsonBlob(c, true, clientcache.CacheInfinite, response)
		}
	} else {
		// send invalid address message
		return api.ErrorStr(c, "please, provide a valid ethereum or quorum address")
	}
}

// check if an ethereum address is a contract address
func (ctl *Web3Controller) getBalanceAtBlock(c echo.ContextInterface) error {
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)
	logger.Info("web3 controller request using context id: ", ctl.network.networkName)
	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	targetAddr := c.Param("address")
	block := c.Param("block")
	// check if not empty
	if targetAddr != "" {
		result, err := client.EthGetBalance(targetAddr, block)
		if err != nil {
			return api.Error(c, err)
		}
		return api.SendSuccess(c, "balance_at_block", result)
	}
	// send invalid address message
	return api.ErrorStr(c, "please, provide a valid ethereum or quorum address")
}

// get node information
func (ctl *Web3Controller) getNodeInfo(c echo.ContextInterface) error {
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)
	logger.Info("web3 controller request using context id: ", ctl.network.networkName)
	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	data, err := client.EthNodeInfo()
	if err != nil {
		// send invalid response message
		return api.Error(c, err)
	} else {
		return api.Success(c, "eth_info", data)
	}
}

// get node information
func (ctl *Web3Controller) getNetworkVersion(c echo.ContextInterface) error {

	//try to get this information from the cache
	key := "net_version"
	result, found := ctl.network.cache.Get(key)
	if found && result != nil {
		//cache hit
		logger.Info(key, ": cache hit")
		response := api.ToSuccess(key, result.(string))
		return clientcache.CachedJsonBlob(c, true, clientcache.CacheInfinite, response)
	} else {
		//cache miss
		logger.Info(key, ": cache miss")
		// get our client context
		client, cliErr := ctl.network.getRpcClient(c)
		logger.Info("web3 controller request using context id: ", ctl.network.networkName)
		if cliErr != nil {
			return api.Error(c, cliErr)
		}
		data, err := client.EthNetVersion()
		if err != nil {
			// send invalid response message
			return api.Error(c, err)
		} else {
			// save result in the cache
			ctl.network.cache.Set(key, data)
			response := api.ToSuccess(key, data)
			return clientcache.CachedJsonBlob(c, true, clientcache.CacheInfinite, response)
		}
	}
}

// this functions detects is web3_clientVersion information is
// EthereumJS TestRPC/v2.5.5-beta.0/ethereum-js or similar
func (ctl *Web3Controller) isRunningGanache(c echo.ContextInterface) error {
	//try to get this information from the cache
	key := "is_ganache"
	result, found := ctl.network.cache.Get(key)
	if found && result != nil {
		//cache hit
		logger.Info(key, ": cache hit")
		response := api.ToSuccess(key, result)
		return clientcache.CachedJsonBlob(c, true, clientcache.CacheInfinite, response)
	} else {
		//cache miss
		logger.Info(key, ": cache miss")
		// get our client context
		client, cliErr := ctl.network.getRpcClient(c)
		logger.Info("web3 controller request using context id: ", ctl.network.networkName)
		if cliErr != nil {
			return api.Error(c, cliErr)
		}
		data, err := client.IsGanache()
		if err != nil {
			// send invalid response message
			return api.Error(c, err)
		} else {
			response := api.ToSuccess("is_ganache", data)
			return clientcache.CachedJsonBlob(c, true, clientcache.CacheInfinite, response)
		}
	}
}

func (ctl *Web3Controller) makeRpcCallNoParams(c echo.ContextInterface) error {
	//resolve method name from url
	methodName := c.Request().URL.Path
	//try to get this information from the cache
	// methodName example: /v1/public/ropsten/net/version
	chunks := strings.Split(methodName, "/")
	if len(chunks) < 4 {
		return api.ErrorStr(c, "invalid url or web3 method provided")
	}
	var key string
	if len(chunks) == 4 {
		key = chunks[3]
	} else if len(chunks) == 5 {
		key = chunks[3] + "_" + chunks[4]
	}
	//resolve method name from key value
	method := methodMap[key]
	cacheKey := ctl.network.peer + ":" + method
	result, found := ctl.network.cache.Get(cacheKey)
	if found && result != nil {
		//cache hit
		logger.Info(method, ": cache hit")
		response := api.ToSuccess(method, result)
		return clientcache.CachedJsonBlob(c, true, clientcache.CacheInfinite, response)
	} else {
		//cache miss
		logger.Info(method, ": cache miss")
		// get our client context
		client, cliErr := ctl.network.getRpcClient(c)
		logger.Info("web3 controller request using context id: ", ctl.network.networkName)
		if cliErr != nil {
			return api.Error(c, cliErr)
		}
		rpcResponse, err := client.EthMethodNoParams(method)
		if err != nil {
			// send invalid response message
			return api.Error(c, err)
		} else if rpcResponse == nil {
			// send invalid response message
			return api.ErrorStr(c, "the network peer did not return any response")
		} else {
			// save result in the cache
			ctl.network.cache.Set(cacheKey, rpcResponse)
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
func (ctl *Web3Controller) getAccountsWithBalance(c echo.ContextInterface) error {
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)
	logger.Info("web3 controller request using context id: ", ctl.network.networkName)
	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	// list all our accounts
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
			str.GetJsonBytes(
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
			str.GetJsonBytes(
				protocol.NewApiResponse("ethereum accounts and their balance readed", wrapperList),
			),
		)
	}
}

// check if an ethereum address is a contract address
func (ctl *Web3Controller) isContractAddress(c echo.ContextInterface) error {
	targetAddr := c.Param("address")
	// check if not empty
	if targetAddr != "" {
		// get our client context
		client, cliErr := ctl.network.getRpcClient(c)
		logger.Info("web3 controller request using context id: ", ctl.network.networkName)
		if cliErr != nil {
			return api.Error(c, cliErr)
		}
		result, err := client.IsSmartContractAddress(targetAddr)
		if err != nil {
			return api.Error(c, err)
		}
		response := api.ToSuccess("is_contract", result)
		return clientcache.CachedJsonBlob(c, true, clientcache.CacheInfinite, response)
	}
	// send invalid address message
	return api.ErrorStr(c, "please, provide a valid ethereum or quorum address")
}

// Start ERC20 functions

// get the total supply of the contract at given target network
func (ctl *Web3Controller) erc20Name(c echo.ContextInterface) error {
	contractAddress := c.Param("contract")
	//input data validation
	if contractAddress == "" {
		return api.ErrorStr(c, "invalid contract address provided")
	}
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)
	logger.Info("web3 controller request using context id: ", ctl.network.networkName)
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
func (ctl *Web3Controller) erc20Symbol(c echo.ContextInterface) error {
	contractAddress := c.Param("contract")
	//input data validation
	if contractAddress == "" {
		return api.ErrorStr(c, "invalid contract address provided")
	}
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)
	logger.Info("web3 controller request using context id: ", ctl.network.networkName)
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
func (ctl *Web3Controller) erc20totalSupply(c echo.ContextInterface) error {
	contractAddress := c.Param("contract")
	//input data validation
	if contractAddress == "" {
		return api.ErrorStr(c, "invalid contract address provided")
	}
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)
	logger.Info("web3 controller request using context id: ", ctl.network.networkName)
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
func (ctl *Web3Controller) erc20decimals(c echo.ContextInterface) error {
	contractAddress := c.Param("contract")
	//input data validation
	if contractAddress == "" {
		return api.ErrorStr(c, "invalid contract address provided")
	}
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)
	logger.Info("web3 controller request using context id: ", ctl.network.networkName)
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
func (ctl *Web3Controller) erc20Balanceof(c echo.ContextInterface) error {
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
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)
	logger.Info("web3 controller request using context id: ", ctl.network.networkName)
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

// get the allowance status of the contract at given target network
func (ctl *Web3Controller) erc20Allowance(c echo.ContextInterface) error {
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
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)
	logger.Info("web3 controller request using context id: ", ctl.network.networkName)
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
func (ctl *Web3Controller) erc20Transfer(c echo.ContextInterface) error {
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
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)
	logger.Info("web3 controller request using context id: ", ctl.network.networkName)
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
func (ctl *Web3Controller) erc20Approve(c echo.ContextInterface) error {
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
func (ctl *Web3Controller) erc20TransferFrom(c echo.ContextInterface) error {
	return nil
}

// eth.sendTransaction({from:sender, to:receiver, value: amount})
func (ctl *Web3Controller) sendTransaction(c echo.ContextInterface) error {
	to := c.Param("to")
	//input data validation
	if to == "" {
		return api.ErrorStr(c, "invalid destination address provided")
	}
	amount := c.Param("amount")
	tokenAmount, pErr := strconv.Atoi(amount)
	//input data validation
	if amount == "" || pErr != nil || tokenAmount <= 0 {
		return api.ErrorStr(c, "invalid ether amount value provided")
	}
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)
	logger.Info("web3 controller request using context id: ", ctl.network.networkName)
	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	//build our transaction
	var transaction ethrpc.TransactionData
	transaction.To = to
	transaction.Value = eth.ToWei(tokenAmount, 0)

	raw, err := client.EthSendTransaction(transaction)
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

// END of ERC20 functions

func (ctl *Web3Controller) getTransactionByHash(c echo.ContextInterface) error {
	txhash := c.Param("hash")
	// check if not empty
	if txhash != "" {

		// get our client context
		client, cliErr := ctl.network.getRpcClient(c)
		logger.Info("web3 controller request using context id: ", ctl.network.networkName)
		if cliErr != nil {
			return api.Error(c, cliErr)
		}

		result, err := client.EthGetTransactionReceipt(txhash)
		if err != nil {
			return api.Error(c, err)
		}
		return api.SendSuccess(c, "transaction_receipt", result)
	}
	// send invalid address message
	return api.ErrorStr(c, "please, provide a valid ethereum or quorum address")
}

// implemented method from interface RouterRegistrable
func (ctl Web3Controller) RegisterRouters(router *echo.Group) {

	router.GET("/is/ganache", ctl.isRunningGanache)

	router.GET("/client/version", ctl.makeRpcCallNoParams)
	router.GET("/net/version", ctl.makeRpcCallNoParams)
	router.GET("/net/peers", ctl.makeRpcCallNoParams)
	router.GET("/protocol/version", ctl.makeRpcCallNoParams)
	router.GET("/syncing", ctl.makeRpcCallNoParams)
	router.GET("/coinbase", ctl.makeRpcCallNoParams)
	router.GET("/mining", ctl.makeRpcCallNoParams)
	router.GET("/hashrate", ctl.makeRpcCallNoParams)
	router.GET("/gasprice", ctl.makeRpcCallNoParams)
	router.GET("/accounts", ctl.makeRpcCallNoParams)
	router.GET("/block/latest", ctl.makeRpcCallNoParams)
	router.GET("/block/current", ctl.makeRpcCallNoParams)
	router.GET("/compilers", ctl.makeRpcCallNoParams)

	router.GET("/tx/send", ctl.sendTransaction)
	router.GET("/tx/hash/:hash", ctl.getTransactionByHash)
	router.GET("/is/contract/:address", ctl.isContractAddress)

	router.GET("/accountsBalanced", ctl.getAccountsWithBalance)

	router.GET("/balance/:address", ctl.getBalance)
	router.GET("/balance/:address/block/:block", ctl.getBalanceAtBlock)

	router.GET("/erc20/:contract/name", ctl.erc20Name)
	router.GET("/erc20/:contract/symbol", ctl.erc20Symbol)
	router.GET("/erc20/:contract/totalsupply", ctl.erc20totalSupply)
	router.GET("/erc20/:contract/decimals", ctl.erc20decimals)
	router.GET("/erc20/:contract/balanceof/:address", ctl.erc20Balanceof)
	router.GET("/erc20/:contract/allowance/:owner/to/:spender", ctl.erc20Allowance)
	router.GET("/erc20/:contract/transfer/:address/:amount", ctl.erc20Transfer)
}
