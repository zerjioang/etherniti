// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package network

import (
	"github.com/pkg/errors"
	"math/big"
	"net/http"
	"strconv"
	"strings"

	"github.com/zerjioang/etherniti/shared/constants"

	"github.com/zerjioang/etherniti/core/eth/paramencoder/erc20"

	"github.com/zerjioang/etherniti/core/data"

	"github.com/zerjioang/etherniti/core/eth"

	"github.com/zerjioang/etherniti/core/eth/rpc"

	"github.com/zerjioang/etherniti/core/util/str"

	"github.com/zerjioang/etherniti/core/modules/encoding/hex"

	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/shared/protocol"

	"github.com/zerjioang/etherniti/core/eth/fixtures"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

var (
	methodMap = map[string]string{
		"client_version":   "web3_clientVersion",
		"net_version":      "net_version",
		"net_listening":    "net_listening",
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
		"block_current":    "eth_blockNumber",
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
func (ctl *Web3Controller) getBalance(c *echo.Context) error {
	targetAddr := c.Param("address")
	// check if not empty
	if targetAddr != "" {
		//try to get this information from the cache
		key := ctl.network.peer + "get_balance" + targetAddr
		keyBytes := str.UnsafeBytes(key)
		result, found := ctl.network.cache.Get(keyBytes)
		if found && result != nil {
			//cache hit
			logger.Info(key, ": cache hit")
			c.OnSuccessCachePolicy = constants.CacheInfinite
			return api.SendSuccessBlob(c, result.([]byte))
		} else {
			//cache miss
			logger.Info(key, ": cache miss")
			// get our client context
			client, cliErr := ctl.network.getRpcClient(c)

			if cliErr != nil {
				return api.Error(c, cliErr)
			}
			result, raw, err := client.EthGetBalance(targetAddr, "latest")
			if err != nil {
				return api.Error(c, err)
			}
			response := api.ToSuccess(data.Balance, BalanceResponse{Value: result, Raw: raw})
			ctl.network.cache.Set(keyBytes, response)
			c.OnSuccessCachePolicy = constants.CacheInfinite
			return api.SendSuccessBlob(c, response)
		}
	} else {
		// send invalid address message
		return api.ErrorStr(c, data.MissingAddress)
	}
}

// check if an ethereum address is a contract address
func (ctl *Web3Controller) getBalanceAtBlock(c *echo.Context) error {
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)
	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	targetAddr := c.Param("address")
	block := c.Param("block")
	//todo validate input parameters
	if !eth.IsValidAddressLow(targetAddr) {
		// send invalid address message
		return api.ErrorStr(c, data.MissingAddress)
	}
	if !eth.IsValidBlockNumber(block) {
		//return error that block is invalid
		return api.Error(c, data.ErrInvalidBlockNumber)
	}
	result, raw, err := client.EthGetBalance(targetAddr, block)
	if err != nil {
		return api.Error(c, err)
	}
	response := api.ToSuccess(data.Balance, BalanceResponse{Value: result, Raw: raw})
	return api.SendSuccess(c, data.BalanceAtBlock, response)
}

// get node information
func (ctl *Web3Controller) getNodeInfo(c *echo.Context) error {
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)

	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	d, err := client.EthNodeInfo()
	if err != nil {
		// send invalid response message
		return api.Error(c, err)
	} else {
		return api.Success(c, data.EthInfo, str.UnsafeBytes(d))
	}
}

// get node information
func (ctl *Web3Controller) getNetworkVersion(c *echo.Context) error {

	//try to get this information from the cache
	key := data.NetVersion
	result, found := ctl.network.cache.Get(key)
	if found && result != nil {
		//cache hit
		logger.Info(key, ": cache hit")
		response := api.ToSuccess(key, result.(string))
		c.OnSuccessCachePolicy = constants.CacheInfinite
		return api.SendSuccessBlob(c, response)
	} else {
		//cache miss
		logger.Info(key, ": cache miss")
		// get our client context
		client, cliErr := ctl.network.getRpcClient(c)

		if cliErr != nil {
			return api.Error(c, cliErr)
		}
		response, err := client.EthNetVersion()
		if err != nil {
			// send invalid response message
			return api.Error(c, err)
		} else {
			// save result in the cache
			ctl.network.cache.Set(key, response)
			c.OnSuccessCachePolicy = constants.CacheInfinite
			response := api.ToSuccess(key, response)
			return api.SendSuccessBlob(c, response)
		}
	}
}

// this functions detects is web3_clientVersion information is
// EthereumJS TestRPC/v2.5.5-beta.0/ethereum-js or similar
func (ctl *Web3Controller) isRunningGanache(c *echo.Context) error {
	//try to get this information from the cache
	key := data.IsGanache
	result, found := ctl.network.cache.Get(key)
	if found && result != nil {
		//cache hit
		logger.Info(key, ": cache hit")
		response := api.ToSuccess(key, result)
		c.OnSuccessCachePolicy = constants.CacheInfinite
		return api.SendSuccessBlob(c, response)
	} else {
		//cache miss
		logger.Info(key, ": cache miss")
		// get our client context
		client, cliErr := ctl.network.getRpcClient(c)

		if cliErr != nil {
			return api.Error(c, cliErr)
		}
		d, err := client.IsGanache()
		if err != nil {
			// send invalid response message
			return api.Error(c, err)
		} else {
			response := api.ToSuccess(data.IsGanache, d)
			c.OnSuccessCachePolicy = constants.CacheInfinite
			return api.SendSuccessBlob(c, response)
		}
	}
}

func (ctl *Web3Controller) sha3(c *echo.Context) error {
	// 0 parse this http rquest body
	var model map[string]interface{}
	if err := c.Bind(&model); err != nil {
		// return a binding error
		logger.Error("failed to bind request data to model: ", err)
		return api.ErrorStr(c, data.BindErr)
	}
	content, found := model["data"]
	if found && content != nil {
		//some data found in body content
		// try to convert data to string
		strData, ok := content.(string)
		if ok && len(strData) > 0 {
			//succesfully converted data to string
			// 1 create a cache key
			//ckey := "sha3:"+strData
			// 2 fetch our ethereum client instance
			// get our client context
			client, cliErr := ctl.network.getRpcClient(c)
			if cliErr != nil {
				return api.Error(c, cliErr)
			}
			// 3 make the call
			response, err := client.Web3Sha3(str.UnsafeBytes(strData))
			// 4 process ethereum response
			if err != nil {
				// send invalid response message
				return api.Error(c, err)
			} else {
				response := api.ToSuccess([]byte("sha3"), response)
				c.OnSuccessCachePolicy = constants.CacheInfinite
				return api.SendSuccessBlob(c, response)
			}
		} else {
			return api.Error(c, errors.New("no input data provided for sha3"))
		}
	} else {
		// no content send for key 'data'
		return api.Error(c, errors.New("no input data provided for sha3 at request body"))
	}
}

func (ctl *Web3Controller) makeRpcCallNoParams(c *echo.Context) error {
	//resolve method name from url
	methodName := c.Request().URL.Path
	//try to get this information from the cache
	// methodName example: /v1/web3/private/net/version
	chunks := strings.Split(methodName, "/")
	if len(chunks) < 5 {
		return api.ErrorStr(c, data.InvalidUrlWeb3)
	}
	var key string
	if len(chunks) == 6 {
		// example url: /v1/web3/private/net/version
		key = chunks[4] + "_" + chunks[5]
	} else if len(chunks) == 5 {
		// example url: /v1/web3/private/syncing
		key = chunks[4]
	}
	//resolve method name from key value
	method := methodMap[key]
	methodBytes := str.UnsafeBytes(method)
	//TODO : in private contenxt peer name is empty
	cacheKey := ctl.network.peer + ":" + method
	cacheKeyBytes := str.UnsafeBytes(cacheKey)
	result, found := ctl.network.cache.Get(cacheKeyBytes)
	if found && result != nil {
		//cache hit
		logger.Info(method, ": cache hit")
		response := api.ToSuccess(methodBytes, result)
		c.OnSuccessCachePolicy = constants.CacheInfinite
		return api.SendSuccessBlob(c, response)
	} else {
		//cache miss
		logger.Info(method, ": cache miss")
		// get our client context
		client, cliErr := ctl.network.getRpcClient(c)

		if cliErr != nil {
			return api.Error(c, cliErr)
		}
		rpcResponse, err := client.EthMethodNoParams(method)
		if err != nil {
			// send invalid response message
			return api.Error(c, err)
		} else if rpcResponse == nil {
			// send invalid response message
			return api.ErrorStr(c, data.NetworkNoResponse)
		} else {
			// save result in the cache
			ctl.network.cache.Set(cacheKeyBytes, rpcResponse)
			response := api.ToSuccess(methodBytes, rpcResponse)
			c.OnSuccessCachePolicy = constants.CacheInfinite
			return api.SendSuccessBlob(c, response)
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
func (ctl *Web3Controller) getAccountsWithBalance(c *echo.Context) error {
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)

	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	// list all our accounts
	list, err := client.EthAccounts()

	type wrapper struct {
		Account string `json:"account"`
		Balance string `json:"balance"`
		Eth     string `json:"eth"`
		Raw     string `json:"raw"`
		Key     string `json:"key"`
	}
	wrapperList := make([]wrapper, len(list))

	if err != nil {
		// send invalid generation message
		return c.JSONBlob(protocol.StatusInternalServerError,
			str.GetJsonBytes(
				protocol.NewApiError(protocol.StatusInternalServerError, str.UnsafeBytes(err.Error())),
			),
		)
	} else {
		//iterate over account
		for i := 0; i < len(list); i++ {
			currentAccount := list[i]
			bigInt, raw, err := client.EthGetBalance(currentAccount, "latest")
			if err != nil {
				logger.Error("failed to get account balance", currentAccount, err)
			} else {
				item := &wrapperList[i]
				item.Account = currentAccount
				item.Balance = bigInt.String()
				item.Eth = fixtures.ToEth(bigInt).String()
				item.Raw = raw
				item.Key = "secret"
			}
		}
		return api.SendSuccess(c, data.AccountsBalanced, wrapperList)
	}
}

// check if an ethereum address is a contract address
func (ctl *Web3Controller) isContractAddress(c *echo.Context) error {
	targetAddr := c.Param("address")
	// check if not empty
	if targetAddr != "" {
		// get our client context
		client, cliErr := ctl.network.getRpcClient(c)

		if cliErr != nil {
			return api.Error(c, cliErr)
		}
		result, err := client.IsSmartContractAddress(targetAddr)
		if err != nil {
			return api.Error(c, err)
		}
		response := api.ToSuccess(data.IsContract, result)
		c.OnSuccessCachePolicy = constants.CacheInfinite
		return api.SendSuccessBlob(c, response)
	}
	// send invalid address message
	return api.ErrorStr(c, data.MissingAddress)
}

// Start ERC20 functions

// get the total supply of the contract at given target network
func (ctl *Web3Controller) erc20Name(c *echo.Context) error {
	contractAddress := c.Param("contract")
	//input data validation
	if contractAddress == "" {
		return api.ErrorStr(c, data.InvalidContractAddress)
	}
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)

	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	raw, err := client.Erc20Name(contractAddress)
	if err != nil {
		// send invalid generation message
		return c.JSONBlob(protocol.StatusBadRequest,
			str.GetJsonBytes(
				protocol.NewApiError(protocol.StatusBadRequest, str.UnsafeBytes(err.Error())),
			),
		)
	} else {
		unpacked := ""
		rawBytes, decodeErr := hex.FromEthHex(string(raw))
		if decodeErr != nil {
			return api.ErrorStr(c, str.UnsafeBytes("failed to hex decode network response: "+decodeErr.Error()))
		}
		err := erc20.LoadErc20Abi().Unpack(&unpacked, "name", rawBytes)
		if err != nil {
			return api.ErrorStr(c, str.UnsafeBytes("failed to decode network response: "+err.Error()))
		} else {
			return api.SendSuccess(c, data.Name, unpacked)
		}
	}
}

// get the total supply of the contract at given target network
func (ctl *Web3Controller) erc20Symbol(c *echo.Context) error {
	contractAddress := c.Param("contract")
	//input data validation
	if contractAddress == "" {
		return api.ErrorStr(c, data.InvalidContractAddress)
	}
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)

	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	raw, err := client.Erc20Symbol(contractAddress)
	if err != nil {
		// send invalid generation message
		return c.JSONBlob(protocol.StatusBadRequest,
			str.GetJsonBytes(
				protocol.NewApiError(protocol.StatusBadRequest, str.UnsafeBytes(err.Error())),
			),
		)
	} else {
		unpacked := ""
		rawBytes, decodeErr := hex.FromEthHex(string(raw))
		if decodeErr != nil {
			return api.ErrorStr(c, str.UnsafeBytes("failed to hex decode network response: "+decodeErr.Error()))
		}
		err := erc20.LoadErc20Abi().Unpack(&unpacked, "symbol", rawBytes)
		if err != nil {
			return api.ErrorStr(c, str.UnsafeBytes("failed to decode network response: "+err.Error()))
		} else {
			return api.SendSuccess(c, data.Symbol, unpacked)
		}
	}
}

// get the total supply of the contract at given target network
func (ctl *Web3Controller) erc20totalSupply(c *echo.Context) error {
	contractAddress := c.Param("contract")
	//input data validation
	if contractAddress == "" {
		return api.ErrorStr(c, data.InvalidContractAddress)
	}
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)

	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	raw, err := client.Erc20TotalSupply(contractAddress)
	if err != nil {
		// send invalid generation message
		return c.JSONBlob(protocol.StatusBadRequest,
			str.GetJsonBytes(
				protocol.NewApiError(protocol.StatusBadRequest, str.UnsafeBytes(err.Error())),
			),
		)
	} else {
		var unpacked *big.Int
		rawBytes, decodeErr := hex.FromEthHex(string(raw))
		if decodeErr != nil {
			return api.ErrorStr(c, str.UnsafeBytes("failed to hex decode network response: "+decodeErr.Error()))
		}
		err := erc20.LoadErc20Abi().Unpack(&unpacked, "totalSupply", rawBytes)
		if err != nil {
			return api.ErrorStr(c, str.UnsafeBytes("failed to decode network response: "+err.Error()))
		} else {
			return api.SendSuccess(c, data.TotalSupply, unpacked)
		}
	}
}

// get the total supply of the contract at given target network
func (ctl *Web3Controller) erc20decimals(c *echo.Context) error {
	contractAddress := c.Param("contract")
	//input data validation
	if contractAddress == "" {
		return api.ErrorStr(c, data.InvalidContractAddress)
	}
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)

	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	raw, err := client.Erc20Decimals(contractAddress)
	if err != nil {
		// send invalid generation message
		return c.JSONBlob(http.StatusBadRequest,
			str.GetJsonBytes(
				protocol.NewApiError(http.StatusBadRequest, str.UnsafeBytes(err.Error())),
			),
		)
	} else {
		var unpacked *big.Int
		rawBytes, decodeErr := hex.FromEthHex(string(raw))
		if decodeErr != nil {
			return api.ErrorStr(c, str.UnsafeBytes("failed to hex decode network response: "+decodeErr.Error()))
		}
		err := erc20.LoadErc20Abi().Unpack(&unpacked, "decimals", rawBytes)
		if err != nil {
			return api.ErrorStr(c, str.UnsafeBytes("failed to decode network response: "+err.Error()))
		} else {
			return api.SendSuccess(c, data.Decimals, unpacked)
		}
	}
}

// get the total supply of the contract at given target network
func (ctl *Web3Controller) erc20Balanceof(c *echo.Context) error {
	contractAddress := c.Param("contract")
	//input data validation
	if contractAddress == "" {
		return api.ErrorStr(c, data.InvalidContractAddress)
	}
	address := c.Param("address")
	//input data validation
	if !eth.IsValidAddressLow(address) {
		return api.ErrorStr(c, data.InvalidAccountAddress)
	}
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)

	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	raw, err := client.Erc20BalanceOf(contractAddress, address)
	if err != nil {
		// send invalid generation message
		return c.JSONBlob(http.StatusBadRequest,
			str.GetJsonBytes(
				protocol.NewApiError(http.StatusBadRequest, str.UnsafeBytes(err.Error())),
			),
		)
	} else {
		return api.SendSuccess(c, data.BalanceOf, raw)
	}
}

// get the allowance status of the contract at given target network
func (ctl *Web3Controller) erc20Allowance(c *echo.Context) error {
	contractAddress := c.Param("contract")
	//input data validation
	if contractAddress == "" {
		return api.ErrorStr(c, data.InvalidContractAddress)
	}
	ownerAddress := c.Param("owner")
	//input data validation
	if !eth.IsValidAddressLow(ownerAddress) {
		return api.ErrorStr(c, data.InvalidAccountOwner)
	}
	spenderAddress := c.Param("spender")
	//input data validation
	if !eth.IsValidAddressLow(spenderAddress) {
		return api.ErrorStr(c, data.InvalidAccountSpender)
	}
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)

	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	raw, err := client.Erc20Allowance(contractAddress, ownerAddress, spenderAddress)
	if err != nil {
		// send invalid generation message
		return c.JSONBlob(http.StatusBadRequest,
			str.GetJsonBytes(
				protocol.NewApiError(http.StatusBadRequest, str.UnsafeBytes(err.Error())),
			),
		)
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
func (ctl *Web3Controller) erc20Transfer(c *echo.Context) error {
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

	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	raw, err := client.Erc20Transfer(contractAddress, receiverAddress, tokenAmount)
	if err != nil {
		// send invalid generation message
		return c.JSONBlob(http.StatusBadRequest,
			str.GetJsonBytes(
				protocol.NewApiError(http.StatusBadRequest, str.UnsafeBytes(err.Error())),
			),
		)
	} else {
		return api.SendSuccess(c, data.Transfer, raw)
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
func (ctl *Web3Controller) erc20Approve(c *echo.Context) error {
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
func (ctl *Web3Controller) erc20TransferFrom(c *echo.Context) error {
	return nil
}

// eth.sendTransaction({from:sender, to:receiver, value: amount})
func (ctl *Web3Controller) sendTransaction(c *echo.Context) error {
	to := c.Param("to")
	//input data validation
	if to == "" {
		return api.ErrorStr(c, data.InvalidDstAddress)
	}
	amount := c.Param("amount")
	tokenAmount, pErr := strconv.Atoi(amount)
	//input data validation
	if amount == "" || pErr != nil || tokenAmount <= 0 {
		return api.ErrorStr(c, data.InvalidEtherValue)
	}
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)

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
				protocol.NewApiError(http.StatusBadRequest, str.UnsafeBytes(err.Error())),
			),
		)
	} else {
		return api.SendSuccess(c, data.Allowance, raw)
	}
}

// END of ERC20 functions

func (ctl *Web3Controller) getTransactionByHash(c *echo.Context) error {
	txhash := c.Param("hash")
	// check if not empty
	if txhash != "" {

		// get our client context
		client, cliErr := ctl.network.getRpcClient(c)

		if cliErr != nil {
			return api.Error(c, cliErr)
		}

		result, err := client.EthGetTransactionReceipt(txhash)
		if err != nil {
			return api.Error(c, err)
		}
		return api.SendSuccess(c, data.TransactionReceipt, result)
	}
	// send invalid address message
	return api.ErrorStr(c, data.MissingAddress)
}

// Blockchain Access

// ChainId retrieves the current chain ID for transaction replay protection.
func (ctl *Web3Controller) chainId(c *echo.Context) error {
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)
	if cliErr != nil {
		return api.Error(c, cliErr)
	}

	result, err := client.ChainId()
	if err != nil {
		return api.Error(c, err)
	}
	return api.SendSuccess(c, data.ChainId, result)
}

func (ctl *Web3Controller) noop(c *echo.Context) error {
	return api.Error(c, errors.New("not implemented"))
}

// TODO implement the method
// Returns the value from a storage position at a given address.
func (ctl *Web3Controller) ethGetStorageAt(c *echo.Context) error {
	addr := c.Param("address")
	//input data validation
	if !eth.IsValidAddressLow(addr) {
		return api.ErrorStr(c, data.InvalidAccountAddress)
	}
	//read block position
	block := c.Param("block")
	//validate block input position data
	if !eth.IsValidBlockNumber(block) {
		//return error that block is invalid
		return api.Error(c, data.ErrInvalidBlockNumber)
	}
	pos := c.Param("position")
	if pos == "" {
		//return invalid position index provided
		return api.Error(c, errors.New("invalid storage position provided"))
	}
	return ctl.noop(c)
}

// Returns the number of transactions sent from an address.
func (ctl *Web3Controller) getTransactionCount(c *echo.Context) error {
	// read input parameters
	addr := c.Param("address")
	//input data validation
	if !eth.IsValidAddressLow(addr) {
		return api.ErrorStr(c, data.InvalidAccountAddress)
	}
	//read block position
	block := c.Param("block")
	//validate block input position data
	if !eth.IsValidBlockNumber(block) {
		//return error that block is invalid
		return api.Error(c, data.ErrInvalidBlockNumber)
	}
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)
	if cliErr != nil {
		return api.Error(c, cliErr)
	}

	result, err := client.EthGetTransactionCount(addr, block)
	if err != nil {
		return api.Error(c, err)
	}
	return api.SendSuccess(c, data.TransactionCount, result)
}

// Returns the number of transactions in a block from a block matching the given block hash.
func (ctl *Web3Controller) getBlockTransactionCountByHash(c *echo.Context) error {
	// read input parameters
	hash := c.Param("hash")
	//input data validation
	// TODO add strong block hash validation here
	if hash == "" {
		return api.Error(c, data.ErrInvalidBlockHash)
	}
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)
	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	// Returns: integer of the number of transactions in this block.
	result, err := client.EthGetBlockTransactionCountByHash(hash)
	if err != nil {
		return api.Error(c, err)
	}
	return api.SendSuccess(c, data.TransactionCountHash, result)
}

// Returns the number of transactions in a block matching the given block number.
func (ctl *Web3Controller) getBlockTransactionCountByNumber(c *echo.Context) error {
	// read input parameters
	number := c.Param("number")
	//input data validation
	if !eth.IsValidBlockNumber(number) {
		return api.Error(c, data.ErrInvalidBlockNumber)
	}
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)
	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	// Returns: integer of the number of transactions in this block.
	result, err := client.EthGetBlockTransactionCountByNumber(number)
	if err != nil {
		return api.Error(c, err)
	}
	return api.SendSuccess(c, data.TransactionCountBlockNumber, result)
}

// implemented method from interface RouterRegistrable
func (ctl Web3Controller) RegisterRouters(router *echo.Group) {

	router.GET("/is/ganache", ctl.isRunningGanache)

	router.GET("/client/version", ctl.makeRpcCallNoParams)
	router.POST("/sha3", ctl.sha3)
	router.GET("/net/version", ctl.makeRpcCallNoParams)
	router.GET("/net/listening", ctl.makeRpcCallNoParams)
	router.GET("/net/peers", ctl.makeRpcCallNoParams)
	router.GET("/protocol/version", ctl.makeRpcCallNoParams)
	router.GET("/syncing", ctl.makeRpcCallNoParams)
	router.GET("/coinbase", ctl.makeRpcCallNoParams)
	router.GET("/mining", ctl.makeRpcCallNoParams)
	router.GET("/hashrate", ctl.makeRpcCallNoParams)
	router.GET("/gasprice", ctl.makeRpcCallNoParams)
	router.GET("/accounts", ctl.makeRpcCallNoParams)
	router.GET("/accounts/balanced", ctl.getAccountsWithBalance)
	router.GET("/block/latest", ctl.makeRpcCallNoParams)
	router.GET("/balance/:address", ctl.getBalance)
	router.GET("/balance/:address/block/:block", ctl.getBalanceAtBlock)

	// https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_getstorageat
	router.GET("/storage/:address/:block/:position", ctl.ethGetStorageAt)

	// https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_gettransactioncount
	router.GET("/tx/count/address/{address}/{block}", ctl.getTransactionCount)

	// eth_getBlockTransactionCountByHash
	router.GET("/tx/count/hash/{hash}", ctl.getBlockTransactionCountByHash)
	//Returns the number of transactions in a block from a block matching the given block hash.

	//eth_getBlockTransactionCountByNumber
	router.GET("/tx/count/number/{number}", ctl.getBlockTransactionCountByNumber)
	//Returns the number of transactions in a block matching the given block number.

	// eth_getUncleCountByBlockHash
	// Returns the number of uncles in a block from a block matching the given block hash.

	// eth_getUncleCountByBlockNumber
	// Returns the number of uncles in a block from a block matching the given block number.

	// eth_getCode
	// Returns code at a given address.

	// eth_sign

	// eth_sendTransaction

	// eth_sendRawTransaction

	// eth_call

	// eth_estimateGas

	// eth_getBlockByHash

	//eth_getBlockByNumber

	//eth_getTransactionByHash

	//eth_getTransactionByBlockHashAndIndex

	//eth_getTransactionByBlockNumberAndIndex

	//eth_getTransactionReceipt

	//eth_pendingTransactions

	//eth_getUncleByBlockHashAndIndex

	//eth_getUncleByBlockNumberAndIndex

	//eth_getCompilers (DEPRECATED)
	//deprecated calls
	router.GET("/compilers", ctl.makeRpcCallNoParams)
	// eth_compileSolidity (DEPRECATED)
	router.GET("/compile/solidity", ctl.noop)
	// eth_compileLLL (DEPRECATED)
	router.GET("/compile/lll", ctl.noop)
	// eth_compileSerpent (DEPRECATED)
	router.GET("/compile/serpent", ctl.noop)

	router.GET("/chain/id", ctl.chainId)

	router.GET("/tx/send", ctl.sendTransaction)
	router.GET("/tx/hash/:hash", ctl.getTransactionByHash)
	router.GET("/is/contract/:address", ctl.isContractAddress)

	router.GET("/erc20/:contract/name", ctl.erc20Name)
	router.GET("/erc20/:contract/symbol", ctl.erc20Symbol)
	router.GET("/erc20/:contract/totalsupply", ctl.erc20totalSupply)
	router.GET("/erc20/:contract/decimals", ctl.erc20decimals)
	router.GET("/erc20/:contract/balanceof/:address", ctl.erc20Balanceof)
	router.GET("/erc20/:contract/allowance/:owner/to/:spender", ctl.erc20Allowance)
	router.GET("/erc20/:contract/transfer/:address/:amount", ctl.erc20Transfer)
}
