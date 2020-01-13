// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package network

import (
	"math/big"
	"strconv"
	"strings"

	"github.com/zerjioang/etherniti/shared/dto"

	"github.com/zerjioang/etherniti/core/controllers/wrap"

	"github.com/zerjioang/go-hpc/lib/codes"
	"github.com/zerjioang/go-hpc/lib/eth/fixtures/common"

	"github.com/pkg/errors"

	"github.com/zerjioang/etherniti/shared"
	"github.com/zerjioang/etherniti/shared/constants"

	"github.com/zerjioang/go-hpc/lib/eth/paramencoder/erc20"

	"github.com/zerjioang/etherniti/core/data"

	"github.com/zerjioang/go-hpc/lib/eth"

	ethrpc "github.com/zerjioang/go-hpc/lib/eth/rpc"

	"github.com/zerjioang/go-hpc/util/str"

	"github.com/zerjioang/go-hpc/lib/encoding/hex"

	"github.com/zerjioang/etherniti/core/api"

	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/go-hpc/lib/eth/fixtures"
	"github.com/zerjioang/go-hpc/thirdparty/echo"
)

var (
	methodMap = map[string]string{
		"client_version":   "web3_clientVersion",
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
func (ctl *Web3Controller) getBalance(c *shared.EthernitiContext) error {
	targetAddr := c.Param("address")
	// check if not empty
	if targetAddr != "" {
		//try to get this information from the cache
		key := ctl.network.UniqueId() + "get_balance" + targetAddr
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
			c.OnSuccessCachePolicy = constants.CacheInfinite
			//todo add support for dynamic encoding instead of nil
			ctl.network.cache.Set(keyBytes, nil)
			return api.SendSuccess(c, data.Balance, BalanceResponse{Value: result, Raw: raw, Eth: fixtures.ToEth(result).String()})
		}
	} else {
		// send invalid address message
		return api.ErrorBytes(c, data.MissingAddress)
	}
}

// check if an ethereum address is a contract address
func (ctl *Web3Controller) getBalanceAtBlock(c *shared.EthernitiContext) error {
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
		return api.ErrorBytes(c, data.MissingAddress)
	}
	if !eth.IsValidBlockNumber(block) {
		//return error that block is invalid
		return api.Error(c, data.ErrInvalidBlockNumber)
	}
	result, raw, err := client.EthGetBalance(targetAddr, block)
	if err != nil {
		return api.Error(c, err)
	}
	return api.SendSuccess(c, data.BalanceAtBlock, BalanceResponse{Value: result, Raw: raw})
}

// get node information
func (ctl *Web3Controller) getNodeInfo(c *shared.EthernitiContext) error {
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
func (ctl *Web3Controller) getNetworkVersion(c *shared.EthernitiContext) error {

	//try to get this information from the cache
	key := data.NetVersion
	cachedValue, found := ctl.network.cache.Get(key)
	if found && cachedValue != nil {
		//cache hit
		logger.Info(key, ": cache hit")
		return api.SendSuccess(c, data.NetVersion, cachedValue)
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
			// save cachedValue in the cache
			ctl.network.cache.Set(key, response)
			c.OnSuccessCachePolicy = constants.CacheInfinite
			return api.SendSuccess(c, data.NetVersion, response)
		}
	}
}

// this functions detects is web3_clientVersion information is
// EthereumJS TestRPC/v2.5.5-beta.0/ethereum-js or similar
func (ctl *Web3Controller) isRunningGanache(c *shared.EthernitiContext) error {
	//try to get this information from the cache
	key := data.IsGanache
	cachedValue, found := ctl.network.cache.Get(key)
	if found && cachedValue != nil {
		//cache hit
		logger.Info(key, ": cache hit")
		return api.SendSuccess(c, data.IsGanache, cachedValue)
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
			//set cached value
			ctl.network.cache.Set(key, d)
			c.OnSuccessCachePolicy = constants.CacheInfinite
			return api.SendSuccess(c, data.IsGanache, d)
		}
	}
}

func (ctl *Web3Controller) sha3Node(c *shared.EthernitiContext) error {
	// 0 parse this http request body
	var req *dto.EthSha3Request
	if err := c.Bind(&req); err != nil {
		// return a binding error
		logger.Error(data.FailedToBind, err)
		return api.ErrorBytes(c, data.BindErr)
	}
	err := req.Validate()
	if err == nil {
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
		response, err := client.Web3Sha3(str.UnsafeBytes(req.Data))
		// 4 process ethereum response
		if err != nil {
			// send invalid response message
			return api.Error(c, err)
		} else {
			c.OnSuccessCachePolicy = constants.CacheInfinite
			return api.SendSuccess(c, data.Sha3, response)
		}
	} else {
		// error detected on input data
		return api.Error(c, err)
	}
}

func (ctl *Web3Controller) sha3BuiltIn(c *shared.EthernitiContext) error {
	// 1 parse this http request body
	var req *dto.EthSha3Request
	if err := c.Bind(&req); err != nil {
		// return a binding error
		logger.Error(data.FailedToBind, err)
		return api.ErrorBytes(c, data.BindErr)
	}
	// 2 validate input data
	err := req.Validate()
	if err == nil {
		// 3 generate keccak-256 hash withut connecting to remote node
		hash := fixtures.MessageHash(req.Data)
		if hash == nil || len(hash) == 0 {
			// send invalid response message
			return api.Error(c, errors.New("failed to create keccak-256 of provided input data"))
		} else {
			c.OnSuccessCachePolicy = constants.CacheInfinite
			return api.SendSuccess(c, data.Sha3, hex.ToEthHex(hash))
		}
	} else {
		// error during data validation
		return api.Error(c, err)
	}
}

func (ctl *Web3Controller) parseSignature(c *shared.EthernitiContext) error {
	// 0 parse this http request body
	var req *dto.EthSignatureParseRequest
	if err := c.Bind(&req); err != nil {
		// return a binding error
		logger.Error("failed to bind request data to req: ", err)
		return api.ErrorBytes(c, data.BindErr)
	}
	err := req.Validate()
	if err == nil {
		// generate keccak-256 hash withut connecting to remote node
		r, s, v := fixtures.ParseEthSignature(req.Signature)
		resp := dto.EthSignatureParseResponse{R: r, S: s, V: v}
		c.OnSuccessCachePolicy = constants.CacheInfinite
		return api.SendSuccess(c, data.EthSignatureParse, resp)
	} else {
		// error validating input data
		logger.Error("error validating input data: ", err)
		return api.Error(c, err)
	}
}

// TODO implement this function
func (ctl *Web3Controller) ecRecover(c *shared.EthernitiContext) error {
	return api.Error(c, errors.New("not implemented"))
}

func (ctl *Web3Controller) netVersion(c *shared.EthernitiContext) error {
	client, cliErr := ctl.network.getRpcClient(c)
	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	// make the call
	response, err := client.NetVersion()
	if err != nil {
		// send invalid response message
		return api.Error(c, err)
	} else {
		type netVersion struct {
			Version string `json:"version"`
			Name    string `json:"name"`
		}
		var wrapper netVersion
		wrapper.Version = response
		wrapper.Name = ethrpc.ResolveNetworkId(response)
		c.OnSuccessCachePolicy = constants.CacheInfinite
		return api.SendSuccess(c, data.NetVersion, wrapper)
	}
}

func (ctl *Web3Controller) makeRpcCallNoParams(c *shared.EthernitiContext) error {
	//resolve method name from url
	methodName := c.Request().URL.Path
	//try to get this information from the cache
	// methodName example: /v1/web3/private/net/version
	chunks := strings.Split(methodName, "/")
	if len(chunks) < 5 {
		return api.ErrorBytes(c, data.InvalidUrlWeb3)
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
	//TODO : in private context peer name is empty
	cacheKey := ctl.network.UniqueId() + ":" + method
	cacheKeyBytes := str.UnsafeBytes(cacheKey)
	result, found := ctl.network.cache.Get(cacheKeyBytes)
	if found && result != nil {
		//cache hit
		logger.Info(method, ": cache hit")
		c.OnSuccessCachePolicy = constants.CacheInfinite
		return api.SendSuccess(c, methodBytes, result)
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
			return api.ErrorBytes(c, data.NetworkNoResponse)
		} else {
			// save result in the cache
			ctl.network.cache.Set(cacheKeyBytes, rpcResponse)
			c.OnSuccessCachePolicy = constants.CacheInfinite
			return api.SendSuccess(c, methodBytes, rpcResponse)
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
func (ctl *Web3Controller) getAccountsWithBalance(c *shared.EthernitiContext) error {
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
		return api.Error(c, err)
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
func (ctl *Web3Controller) isContractAddress(c *shared.EthernitiContext) error {
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
		c.OnSuccessCachePolicy = constants.CacheInfinite
		return api.SendSuccess(c, data.IsContract, result)
	}
	// send invalid address message
	return api.ErrorBytes(c, data.MissingAddress)
}

// Start ERC20 functions

// get the total supply of the contract at given target network
func (ctl *Web3Controller) erc20Name(c *shared.EthernitiContext) error {
	contractAddress := c.Param("contract")
	//input data validation
	if contractAddress == "" {
		return api.ErrorBytes(c, data.InvalidContractAddress)
	}
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)

	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	raw, err := client.Erc20Name(contractAddress)
	if err != nil {
		// send invalid generation message
		return api.Error(c, err)
	} else {
		unpacked := ""
		rawBytes, decodeErr := hex.FromEthHex(string(raw))
		if decodeErr != nil {
			return api.ErrorBytes(c, str.UnsafeBytes("failed to hex decode network response: "+decodeErr.Error()))
		}
		err := erc20.LoadErc20Abi().Unpack(&unpacked, "name", rawBytes)
		if err != nil {
			return api.ErrorBytes(c, str.UnsafeBytes("failed to decode network response: "+err.Error()))
		} else {
			return api.SendSuccess(c, data.Name, unpacked)
		}
	}
}

// get the total supply of the contract at given target network
func (ctl *Web3Controller) erc20Symbol(c *shared.EthernitiContext) error {
	contractAddress := c.Param("contract")
	//input data validation
	if contractAddress == "" {
		return api.ErrorBytes(c, data.InvalidContractAddress)
	}
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)

	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	raw, err := client.Erc20Symbol(contractAddress)
	if err != nil {
		return api.Error(c, err)
	} else {
		unpacked := ""
		rawBytes, decodeErr := hex.FromEthHex(string(raw))
		if decodeErr != nil {
			return api.ErrorBytes(c, str.UnsafeBytes("failed to hex decode network response: "+decodeErr.Error()))
		}
		err := erc20.LoadErc20Abi().Unpack(&unpacked, "symbol", rawBytes)
		if err != nil {
			return api.ErrorBytes(c, str.UnsafeBytes("failed to decode network response: "+err.Error()))
		} else {
			return api.SendSuccess(c, data.Symbol, unpacked)
		}
	}
}

// get the total supply of the contract at given target network
func (ctl *Web3Controller) erc20totalSupply(c *shared.EthernitiContext) error {
	contractAddress := c.Param("contract")
	//input data validation
	if contractAddress == "" {
		return api.ErrorBytes(c, data.InvalidContractAddress)
	}
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)

	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	raw, err := client.Erc20TotalSupply(contractAddress)
	if err != nil {
		// send invalid generation message
		return api.ErrorCode(c, codes.StatusBadRequest, err)
	} else {
		var unpacked *big.Int
		rawBytes, decodeErr := hex.FromEthHex(string(raw))
		if decodeErr != nil {
			return api.ErrorBytes(c, str.UnsafeBytes("failed to hex decode network response: "+decodeErr.Error()))
		}
		err := erc20.LoadErc20Abi().Unpack(&unpacked, "totalSupply", rawBytes)
		if err != nil {
			return api.ErrorBytes(c, str.UnsafeBytes("failed to decode network response: "+err.Error()))
		} else {
			return api.SendSuccess(c, data.TotalSupply, unpacked)
		}
	}
}

// get the total supply of the contract at given target network
func (ctl *Web3Controller) erc20decimals(c *shared.EthernitiContext) error {
	contractAddress := c.Param("contract")
	//input data validation
	if contractAddress == "" {
		return api.ErrorBytes(c, data.InvalidContractAddress)
	}
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)

	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	raw, err := client.Erc20Decimals(contractAddress)
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
			return api.SendSuccess(c, data.Decimals, unpacked)
		}
	}
}

// get the total supply of the contract at given target network
func (ctl *Web3Controller) erc20Balanceof(c *shared.EthernitiContext) error {
	contractAddress := c.Param("contract")
	//input data validation
	if contractAddress == "" {
		return api.ErrorBytes(c, data.InvalidContractAddress)
	}
	address := c.Param("address")
	//input data validation
	if !eth.IsValidAddressLow(address) {
		return api.ErrorBytes(c, data.InvalidAccountAddress)
	}
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)

	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	raw, err := client.Erc20BalanceOf(contractAddress, address)
	if err != nil {
		// send invalid generation message
		return api.ErrorCode(c, codes.StatusBadRequest, err)
	} else {
		return api.SendSuccess(c, data.BalanceOf, raw)
	}
}

// get the allowance status of the contract at given target network
func (ctl *Web3Controller) erc20Allowance(c *shared.EthernitiContext) error {
	contractAddress := c.Param("contract")
	//input data validation
	if contractAddress == "" {
		return api.ErrorBytes(c, data.InvalidContractAddress)
	}
	ownerAddress := c.Param("owner")
	//input data validation
	if !eth.IsValidAddressLow(ownerAddress) {
		return api.ErrorBytes(c, data.InvalidAccountOwner)
	}
	spenderAddress := c.Param("spender")
	//input data validation
	if !eth.IsValidAddressLow(spenderAddress) {
		return api.ErrorBytes(c, data.InvalidAccountSpender)
	}
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)

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
func (ctl *Web3Controller) erc20Transfer(c *shared.EthernitiContext) error {
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

	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	raw, err := client.Erc20Transfer(contractAddress, receiverAddress, tokenAmount)
	if err != nil {
		// send invalid generation message
		return api.ErrorCode(c, codes.StatusBadRequest, err)
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
func (ctl *Web3Controller) erc20Approve(c *shared.EthernitiContext) error {
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
func (ctl *Web3Controller) erc20TransferFrom(c *shared.EthernitiContext) error {
	return nil
}

// eth.sendTransaction({from:sender, to:receiver, value: amount})
func (ctl *Web3Controller) sendTransaction(c *shared.EthernitiContext) error {
	//read input data
	var txData ethrpc.TransactionData
	if err := c.Bind(&txData); err != nil {
		// return a binding error
		logger.Error(data.FailedToBind, err)
		return api.ErrorBytes(c, data.BindErr)
	}

	if txData.To == "" {
		// to field can only be blank when requesting new contract deployments
		// so assume this tx is a contract deployment
		return ctl.sendContractDeploymentTransaction(c, &txData)
	} else {
		// assume this transaction is not a contract deployment
		//input data validation
		if !eth.IsValidAddressLow(txData.To) {
			return api.ErrorBytes(c, data.InvalidDstAddress)
		}

		//input data validation
		if !eth.IsValidAddressLow(txData.From) {
			return api.ErrorBytes(c, data.InvalidSrcAddress)
		}
		//input data validation
		if txData.ValueStr == "" {
			return api.ErrorBytes(c, data.InvalidEtherValue)
		}

		// get our client context
		client, cliErr := ctl.network.getRpcClient(c)
		if cliErr != nil {
			return api.Error(c, cliErr)
		}

		//send our transaction
		raw, err := client.EthSendTransactionPtr(&txData)
		if err != nil {
			// send invalid generation message
			return api.Error(c, err)
		} else {
			return api.SendSuccess(c, data.Allowance, raw)
		}
	}
}

func (ctl *Web3Controller) sendContractDeploymentTransaction(c *shared.EthernitiContext, txData *ethrpc.TransactionData) error {
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)
	if cliErr != nil {
		return api.Error(c, cliErr)
	}

	raw, err := client.EthSendTransactionPtr(txData)
	if err != nil {
		// send invalid generation message
		return api.Error(c, err)
	} else {
		return api.SendSuccess(c, data.Allowance, raw)
	}
}

// END of ERC20 functions

func (ctl *Web3Controller) getTransactionByHash(c *shared.EthernitiContext) error {
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
	return api.ErrorBytes(c, data.MissingAddress)
}

// Blockchain Access

// ChainId retrieves the current chain ID for transaction replay protection.
func (ctl *Web3Controller) chainId(c *shared.EthernitiContext) error {
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

func (ctl *Web3Controller) getUncleCountByBlockHash(c *shared.EthernitiContext) error {
	return ctl.network.Noop(c)
}

func (ctl *Web3Controller) getUncleCountByBlockNumber(c *shared.EthernitiContext) error {
	return ctl.network.Noop(c)
}

func (ctl *Web3Controller) getCode(c *shared.EthernitiContext) error {
	// read input parameters
	// 1 address
	address := c.Param("address")
	if !eth.IsValidAddressLow(address) {
		return api.Error(c, data.ErrInvalidAddress)
	}
	// 2 block
	block := c.Param("block")
	if block != "" && !eth.IsValidBlockNumber(block) {
		return api.Error(c, data.ErrInvalidBlockNumber)
	}
	// if user did not specify any block, set by default to latest
	if block == "" {
		block = "latest"
	}

	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)
	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	// Returns code at a given address.
	result, err := client.EthGetCode(address, block)
	if err != nil {
		return api.Error(c, err)
	}
	return api.SendSuccess(c, data.GetCode, result)
}

func (ctl *Web3Controller) signRemote(c *shared.EthernitiContext) error {
	// read input parameters
	var signReq *ethrpc.NodeSignRequest
	if err := c.Bind(&signReq); err != nil {
		// return a binding error
		logger.Error(data.FailedToBind, err)
		return api.ErrorBytes(c, data.BindErr)
	}
	//validate input params
	if !eth.IsValidAddressLow(signReq.Address) {
		return api.Error(c, data.ErrInvalidAddress)
	}
	if !common.IsOxPrefixedHex(signReq.Data) {
		return api.Error(c, data.ErrInvalidPayload)
	}

	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)
	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	// Returns code at a given address.
	result, err := client.EthSign(signReq.Address, signReq.Data)
	if err != nil {
		return api.Error(c, err)
	}
	return api.SendSuccess(c, data.EthSign, result)
}

func (ctl *Web3Controller) call(c *shared.EthernitiContext) error {
	return ctl.network.Noop(c)
}

func (ctl *Web3Controller) sendRawTransaction(c *shared.EthernitiContext) error {
	return ctl.network.Noop(c)
}

// EstimateGas tries to estimate the gas needed to execute a specific transaction based on
// the current pending state of the backend blockchain. There is no guarantee that this is
// the true gas limit requirement as other transactions may be added or removed by miners,
// but it should provide a basis for setting a reasonable default.
func (ctl *Web3Controller) estimateGas(c *shared.EthernitiContext) error {
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)
	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	var tx ethrpc.TransactionData
	if err := c.Bind(&tx); err != nil {
		// return a binding error
		logger.Error(data.FailedToBind, err)
		return api.ErrorBytes(c, data.BindErr)
	}
	//make the call with json body
	amount, err := client.EthEstimateGas(tx)
	if err != nil {
		return api.Error(c, err)
	}
	return api.SendSuccess(c, data.EstimateGas, amount)
}

func (ctl *Web3Controller) compileCode(c *shared.EthernitiContext, id []byte, compilerCall func(code string) ([]string, error)) error {
	var model map[string]string
	if err := c.Bind(&model); err != nil {
		// return a binding error
		logger.Error(data.FailedToBind, err)
		return api.ErrorBytes(c, data.BindErr)
	}
	contractStr, found := model["contract"]
	if found && len(contractStr) > 0 {
		// make the call
		response, err := compilerCall(contractStr)
		if err != nil {
			return api.Error(c, err)
		}
		return api.SendSuccess(c, id, response)
	}
	return api.Error(c, errors.New("invalid source contract content provided in the request"))
}

func (ctl *Web3Controller) compileLLL(c *shared.EthernitiContext) error {
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)
	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	// make the call
	return ctl.compileCode(c, data.CompileLLL, client.EthCompileLLL)
}

func (ctl *Web3Controller) compileSolidity(c *shared.EthernitiContext) error {
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)
	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	// make the call
	return ctl.compileCode(c, data.CompileSolidity, client.EthCompileSolidity)
}

func (ctl *Web3Controller) compileSerpent(c *shared.EthernitiContext) error {
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)
	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	// make the call
	return ctl.compileCode(c, data.CompileSerpent, client.EthCompileSerpent)
}

// TODO implement the method
// Returns the value from a storage position at a given address.
func (ctl *Web3Controller) ethGetStorageAt(c *shared.EthernitiContext) error {
	addr := c.Param("address")
	//input data validation
	if !eth.IsValidAddressLow(addr) {
		return api.ErrorBytes(c, data.InvalidAccountAddress)
	}
	//read block position
	block := c.Param("block")
	//validate block input position data
	if !eth.IsValidBlockNumber(block) {
		//return error that block is invalid
		return api.Error(c, data.ErrInvalidBlockNumber)
	}
	key := c.Param("key")
	if key == "" {
		//return invalid position index provided
		return api.Error(c, errors.New("invalid hexadecimal key provided"))
	}
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)
	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	// make the call
	response, err := client.EthGetStorageAt(addr, key, block)
	if err != nil {
		return api.Error(c, err)
	}
	return api.SendSuccess(c, data.TransactionCount, response)
}

// Returns the number of transactions sent from an address.
func (ctl *Web3Controller) getTransactionCount(c *shared.EthernitiContext) error {
	// read input parameters
	addr := c.Param("address")
	//input data validation
	if !eth.IsValidAddressLow(addr) {
		return api.ErrorBytes(c, data.InvalidAccountAddress)
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
func (ctl *Web3Controller) getBlockTransactionCountByHash(c *shared.EthernitiContext) error {
	// read input parameters
	hash := c.Param("hash")
	if !eth.IsValidBlockHash(hash) {
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
func (ctl *Web3Controller) getBlockTransactionCountByNumber(c *shared.EthernitiContext) error {
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

	router.GET("/is/ganache", wrap.Call(ctl.isRunningGanache))

	router.GET("/client/version", wrap.Call(ctl.makeRpcCallNoParams))
	router.POST("/sha3/remote", wrap.Call(ctl.sha3Node))
	router.POST("/sha3/local", wrap.Call(ctl.sha3BuiltIn))
	router.GET("/net/version", wrap.Call(ctl.netVersion))
	router.GET("/net/listening", wrap.Call(ctl.makeRpcCallNoParams))
	router.GET("/net/peers", wrap.Call(ctl.makeRpcCallNoParams))
	router.GET("/protocol/version", wrap.Call(ctl.makeRpcCallNoParams))
	router.GET("/syncing", wrap.Call(ctl.makeRpcCallNoParams))
	router.GET("/coinbase", wrap.Call(ctl.makeRpcCallNoParams))
	router.GET("/mining", wrap.Call(ctl.makeRpcCallNoParams))
	router.GET("/hashrate", wrap.Call(ctl.makeRpcCallNoParams))
	router.GET("/gasprice", wrap.Call(ctl.makeRpcCallNoParams))
	router.GET("/accounts", wrap.Call(ctl.makeRpcCallNoParams))
	router.GET("/accounts/balanced", wrap.Call(ctl.getAccountsWithBalance))
	router.GET("/block/latest", wrap.Call(ctl.makeRpcCallNoParams))
	router.GET("/balance/:address", wrap.Call(ctl.getBalance))
	router.GET("/balance/:address/block/:block", wrap.Call(ctl.getBalanceAtBlock))

	// https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_getstorageat
	router.GET("/storage/:address/:block/:position", wrap.Call(ctl.ethGetStorageAt))

	// https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_gettransactioncount
	router.GET("/tx/count/address/:address/:block", wrap.Call(ctl.getTransactionCount))

	// eth_getBlockTransactionCountByHash
	router.GET("/tx/count/hash/:hash", wrap.Call(ctl.getBlockTransactionCountByHash))
	//Returns the number of transactions in a block from a block matching the given block hash.

	//eth_getBlockTransactionCountByNumber
	router.GET("/tx/count/block/:number", wrap.Call(ctl.getBlockTransactionCountByNumber))
	//Returns the number of transactions in a block matching the given block number.

	// eth_getUncleCountByBlockHash
	router.GET("/uncle/count/hash/:hash", wrap.Call(ctl.getUncleCountByBlockHash))

	// eth_getUncleCountByBlockNumber
	// Returns the number of uncles in a block from a block matching the given block hash.
	router.GET("/uncle/count/block/:block", wrap.Call(ctl.getUncleCountByBlockNumber))

	// eth_getCode
	// Returns code at a given address.
	router.GET("/code/:address/:block", wrap.Call(ctl.getCode))

	// eth_sign
	// The sign method calculates an Ethereum specific signature with:
	// sign(keccak256("\x19Ethereum Signed Message:\n" + len(message) + message))).
	// By adding a prefix to the message makes the calculated signature recognisable
	// as an Ethereum specific signature. This prevents misuse where a malicious DApp
	// can sign arbitrary data (e.g. transaction) and use the signature to impersonate the victim.
	// Note the address to sign with must be unlocked.
	router.POST("/tx/sign-remote", wrap.Call(ctl.signRemote))
	router.POST("/tx/sign-local", wrap.Call(ctl.signTransactionLocal))

	router.POST("/signparse", wrap.Call(ctl.parseSignature))

	router.POST("/ecrecover", wrap.Call(ctl.ecRecover))

	// eth_sendTransaction
	router.POST("/tx/send", wrap.Call(ctl.sendTransaction))

	// eth_sendRawTransaction
	router.POST("/tx/send-raw", wrap.Call(ctl.sendRawTransaction))

	// eth_call
	router.POST("/call", wrap.Call(ctl.call))

	// eth_estimateGas
	router.POST("/estimategas", wrap.Call(ctl.estimateGas))

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
	router.GET("/compilers", wrap.Call(ctl.makeRpcCallNoParams))

	// eth_compileSolidity (DEPRECATED)
	router.GET("/compile/solidity", wrap.Call(ctl.compileSolidity))
	// eth_compileLLL (DEPRECATED)
	router.GET("/compile/lll", wrap.Call(ctl.compileLLL))
	// eth_compileSerpent (DEPRECATED)
	router.GET("/compile/serpent", wrap.Call(ctl.compileSerpent))

	router.GET("/chain/id", wrap.Call(ctl.chainId))

	router.GET("/tx/send", wrap.Call(ctl.sendTransaction))
	router.GET("/tx/receipt/:hash", wrap.Call(ctl.getTransactionByHash))
	router.GET("/is/contract/:address", wrap.Call(ctl.isContractAddress))

	router.GET("/erc20/:contract/name", wrap.Call(ctl.erc20Name))
	router.GET("/erc20/:contract/symbol", wrap.Call(ctl.erc20Symbol))
	router.GET("/erc20/:contract/totalsupply", wrap.Call(ctl.erc20totalSupply))
	router.GET("/erc20/:contract/decimals", wrap.Call(ctl.erc20decimals))
	router.GET("/erc20/:contract/balanceof/:address", wrap.Call(ctl.erc20Balanceof))
	router.GET("/erc20/:contract/allowance/:owner/to/:spender", wrap.Call(ctl.erc20Allowance))
	router.GET("/erc20/:contract/transfer/:address/:amount", wrap.Call(ctl.erc20Transfer))
}
