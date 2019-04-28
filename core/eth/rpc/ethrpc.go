// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package ethrpc

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"

	"github.com/zerjioang/etherniti/core/modules/cache"

	"github.com/zerjioang/etherniti/core/eth/fixtures/abi"
	"github.com/zerjioang/etherniti/core/eth/rpc/model"

	"github.com/zerjioang/etherniti/core/util/str"

	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/modules/encoding/hex"

	"github.com/zerjioang/etherniti/core/eth/paramencoder"

	"github.com/zerjioang/etherniti/core/eth/fixtures"

	"github.com/zerjioang/etherniti/thirdparty/gommon/log"
)

// https://documenter.getpostman.com/view/4117254/ethereum-json-rpc/RVu7CT5J
type contractFunction func(string) (string, error)
type ParamsCallback func() string

var (
	instance         = new(EthRPC)
	summaryFunctions = []contractFunction{
		instance.Erc20Name,
		instance.Erc20Symbol,
		instance.Erc20Decimals,
		instance.Erc20TotalSupply,
	}
	summaryFunctionsNames = []string{
		"name", "symbol", "decimals", "totalsupply",
	}
)

var (
	oneEth      = big.NewInt(1000000000000000000)
	oneEthInt64 = oneEth.Int64()
)

// EthRPC - Ethereum rpc client
type EthRPC struct {
	//ethereum or quorum node endpoint
	url string
	//ethereum interaction cache
	cache *cache.MemoryCache
	// http client
	client http.Client
	// debug flag
	Debug bool
}

// New create new rpc client with given url
func NewDefaultRPCPtr(url string) *EthRPC {
	c := NewDefaultRPC(url)
	return &c
}

// New create new rpc client with given url
func NewDefaultRPC(url string) EthRPC {
	rpc := EthRPC{
		url:    url,
		cache:  cache.Instance(),
		client: http.Client{},
		Debug:  true,
	}
	return rpc
}

func (rpc *EthRPC) post(method string, target interface{}, params ParamsCallback) error {
	paramsStr := ""
	if params != nil {
		paramsStr = params()
	}
	result, err := rpc.makePostWithMethodParams(method, paramsStr)
	if err != nil {
		return err
	}

	if target == nil || result == nil {
		return nil
	}

	return json.Unmarshal(result, target)
}

// URL returns client url
func (rpc *EthRPC) URL() string {
	return rpc.url
}

// makePostWithMethodParams returns raw response of method post
/*

eth_call

Executes a new message post immediately without creating a transaction on the block chain.
Parameters

    Object - The transaction post object

    from: DATA, 20 Bytes - (optional) The address the transaction is sent from.
    to: DATA, 20 Bytes - The address the transaction is directed to.
    gas: QUANTITY - (optional) Integer of the gas provided for the transaction execution. eth_call consumes zero gas, but this parameter may be needed by some executions.
    gasPrice: QUANTITY - (optional) Integer of the gasPrice used for each paid gas
    value: QUANTITY - (optional) Integer of the value sent with this transaction
    data: DATA - (optional) Hash of the method signature and encoded parameters. For details see Ethereum Contract ABI

    QUANTITY|TAG - integer block number, or the string "latest", "earliest" or "pending", see the default block parameter

Returns

DATA - the return value of executed contract.
*/
func (rpc *EthRPC) makePostWithMethodParams(method string, params string) (json.RawMessage, error) {
	if params == "" {
		request := `{"id": 1,"jsonrpc": "2.0","method": "` + method + `"}`
		return rpc.makePostRaw(request)
	} else {
		request := `{"id": 1,"jsonrpc": "2.0","method": "` + method + `","params": ` + params + `}`
		return rpc.makePostRaw(request)
	}
}

func (rpc *EthRPC) makePostRaw(data string) (json.RawMessage, error) {

	log.Info("sending request: ", data)
	response, err := rpc.client.Post(rpc.url, "application/json", strings.NewReader(data))
	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	//responseData readed, close body
	_ = response.Body.Close()

	log.Info("response received", str.UnsafeString(responseData))

	resp := model.EthResponse{}
	unmErr := json.Unmarshal(responseData, &resp)
	if unmErr != nil {
		return nil, unmErr
	}

	if resp.Error != nil {
		return nil, resp.Errored()
	}

	return resp.Result, nil
}

// RawCall returns raw response of method post (Deprecated)
func (rpc *EthRPC) RawCall(method string, params string) (json.RawMessage, error) {
	return rpc.makePostWithMethodParams(method, params)
}

func (rpc *EthRPC) EthMethodNoParams(methodName string) (interface{}, error) {
	var response interface{}
	err := rpc.post(methodName, &response, nil)
	return response, err
}

// returns ethereum node information
func (rpc *EthRPC) EthNodeInfo() (string, error) {
	var response string

	err := rpc.post("eth_info", &response, nil)
	return response, err
}

// Web3ClientVersion returns the current client version.
func (rpc *EthRPC) Web3ClientVersion() (string, error) {
	var clientVersion string

	err := rpc.post("web3_clientVersion", &clientVersion, nil)
	return clientVersion, err
}

func (rpc *EthRPC) IsGanache() (bool, error) {
	data, err := rpc.Web3ClientVersion()
	if err != nil {
		return false, err
	} else {
		// check if response data is similar to ganache response
		isGanache := strings.Contains(data, "ethereum-js") || strings.Contains(data, "TestRPC")
		return isGanache, nil
	}
}

// Web3Sha3 returns Keccak-256 (not the standardized SHA3-256) of the given data.
func (rpc *EthRPC) Web3Sha3(data []byte) (string, error) {
	var hash string
	//prepare the params of the sha3 function
	params := func() string {
		return fixtures.Encode(data)
	}
	err := rpc.post("web3_sha3", &hash, params)
	return hash, err
}

// NetVersion returns the current network protocol version.
func (rpc *EthRPC) NetVersion() (string, error) {
	var version string

	err := rpc.post("net_version", &version, nil)
	return version, err
}

// NetListening returns true if client is actively listening for network connections.
func (rpc *EthRPC) NetListening() (bool, error) {
	var listening bool

	err := rpc.post("net_listening", &listening, nil)
	return listening, err
}

// NetPeerCount returns number of peers currently connected to the client.
func (rpc *EthRPC) NetPeerCount() (int, error) {
	var response int
	if err := rpc.post("net_peerCount", &response, nil); err != nil {
		return 0, err
	}

	return response, nil
}

// EthProtocolVersion returns the current ethereum protocol version.
func (rpc *EthRPC) EthProtocolVersion() (string, error) {
	var protocolVersion string

	err := rpc.post("eth_protocolVersion", &protocolVersion, nil)
	return protocolVersion, err
}

// EthSyncing returns an object with data about the sync status or false.
func (rpc *EthRPC) EthSyncing() (*Syncing, error) {
	result, err := rpc.makePostWithMethodParams("eth_syncing", "")
	if err != nil {
		return nil, err
	}
	syncing := new(Syncing)
	if bytes.Equal(result, []byte("false")) {
		return syncing, nil
	}
	err = json.Unmarshal(result, syncing)
	return syncing, err
}

// returns ethereum node information
func (rpc *EthRPC) EthNetVersion() (string, error) {
	var response string

	err := rpc.post("net_version", &response, nil)
	return response, err
}

// EthCoinbase returns the client coinbase address
func (rpc *EthRPC) EthCoinbase() (string, error) {
	var address string

	err := rpc.post("eth_coinbase", &address, nil)
	return address, err
}

// EthMining returns true if client is actively mining new blocks.
func (rpc *EthRPC) EthMining() (bool, error) {
	var mining bool

	err := rpc.post("eth_mining", &mining, nil)
	return mining, err
}

// EthHashrate returns the number of hashes per second that the node is mining with.
func (rpc *EthRPC) EthHashrate() (int, error) {
	var response string

	if err := rpc.post("eth_hashrate", &response, nil); err != nil {
		return 0, err
	}

	return ParseInt(response)
}

// EthGasPrice returns the current price per gas in wei.
func (rpc *EthRPC) EthGasPrice() (int64, error) {
	var response string
	if err := rpc.post("eth_gasPrice", &response, nil); err != nil {
		return 0, err
	}
	// example 0x4a817c800
	// fast check if string starts with 0x
	if len(response) > 2 && response[0] == 48 && response[1] == 120 {
		response = response[2:]
	}
	//return ParseBigInt(response)
	return ParseHexToInt(response)
}

// EthAccounts returns a list of addresses owned by client.
func (rpc *EthRPC) EthAccounts() ([]string, error) {
	var accounts []string
	err := rpc.post("eth_accounts", &accounts, nil)
	return accounts, err
}

// EthBlockNumber returns the number of most recent block.
func (rpc *EthRPC) EthBlockNumber() (int, error) {
	var response string
	if err := rpc.post("eth_blockNumber", &response, nil); err != nil {
		return 0, err
	}

	return ParseInt(response)
}

// EthGetBalance returns the balance of the account of given address in wei.
func (rpc *EthRPC) EthGetBalance(address string, block string) (*big.Int, error) {
	var response string
	//prepare the params of the get balance function
	params := func() string {
		return "[" + rpc.doubleQuote(address) + "," + rpc.doubleQuote(block) + "]"
	}
	if err := rpc.post("eth_getBalance", &response, params); err != nil {
		return new(big.Int), err
	}
	return ParseBigInt(response)
}

// EthGetStorageAt returns the value from a storage position at a given address.
func (rpc *EthRPC) EthGetStorageAt(data string, position int, tag string) (string, error) {
	var result string
	//prepare the params of the function
	params := func() string {
		return "[" + rpc.doubleQuote(data) + "," + rpc.doubleQuote(IntToHex(position)) + "," + rpc.doubleQuote(tag) + "]"
	}
	err := rpc.post("eth_getStorageAt", &result, params)
	return result, err
}

// EthGetTransactionCount returns the number of transactions sent from an address.
func (rpc *EthRPC) EthGetTransactionCount(address string, block string) (int, error) {
	var response string
	//prepare the params of the function
	params := func() string {
		return "[" + rpc.doubleQuote(address) + "," + rpc.doubleQuote(block) + "]"
	}
	if err := rpc.post("eth_getTransactionCount", &response, params); err != nil {
		return 0, err
	}
	return ParseInt(response)
}

// EthGetBlockTransactionCountByHash returns the number of transactions in a block from a block matching the given block hash.
func (rpc *EthRPC) EthGetBlockTransactionCountByHash(hash string) (int, error) {
	var response string
	//prepare the params of the function
	params := func() string {
		return "[" + rpc.doubleQuote(hash) + "]"
	}
	if err := rpc.post("eth_getBlockTransactionCountByHash", &response, params); err != nil {
		return 0, err
	}
	return ParseInt(response)
}

// EthGetBlockTransactionCountByNumber returns the number of transactions in a block from a block matching the given block
func (rpc *EthRPC) EthGetBlockTransactionCountByNumber(number int) (int, error) {
	var response string
	//prepare the params of the function
	params := func() string {
		return "[" + rpc.doubleQuote(IntToHex(number)) + "]"
	}
	if err := rpc.post("eth_getBlockTransactionCountByNumber", &response, params); err != nil {
		return 0, err
	}
	return ParseInt(response)
}

// EthGetUncleCountByBlockHash returns the number of uncles in a block from a block matching the given block hash.
func (rpc *EthRPC) EthGetUncleCountByBlockHash(hash string) (int, error) {
	var response string
	//prepare the params of the function
	params := func() string {
		return "[" + rpc.doubleQuote(hash) + "]"
	}
	if err := rpc.post("eth_getUncleCountByBlockHash", &response, params); err != nil {
		return 0, err
	}
	return ParseInt(response)
}

// EthGetUncleCountByBlockNumber returns the number of uncles in a block from a block matching the given block number.
func (rpc *EthRPC) EthGetUncleCountByBlockNumber(number int) (int, error) {
	var response string
	//prepare the params of the function
	params := func() string {
		return "[" + rpc.doubleQuote(IntToHex(number)) + "]"
	}
	if err := rpc.post("eth_getUncleCountByBlockNumber", &response, params); err != nil {
		return 0, err
	}
	return ParseInt(response)
}

// double quotes given string
func (rpc *EthRPC) doubleQuote(data string) string {
	return `"` + data + `"`
}

// EthGetCode returns code at a given address.
func (rpc *EthRPC) EthGetCode(address string, block string) (string, error) {
	var code string
	//prepare the params of the function
	params := func() string {
		return "[" + rpc.doubleQuote(address) + "," + rpc.doubleQuote(block) + "]"
	}
	err := rpc.post("eth_getCode", &code, params)
	return code, err
}

// EthSign signs data with a given address.
// Calculates an Ethereum specific signature
// with: sign(keccak256("\x19Ethereum Signed Message:\n" + len(message) + message)))
func (rpc *EthRPC) EthSign(address, data string) (string, error) {
	var signature string
	//prepare the params of the function
	params := func() string {
		return "[" + rpc.doubleQuote(address) + "," + rpc.doubleQuote(data) + "]"
	}
	err := rpc.post("eth_sign", &signature, params)
	return signature, err
}

// EthSendTransaction creates new message post transaction
// or a contract creation, if the data field contains code.
func (rpc *EthRPC) EthSendTransaction(transaction TransactionData) (string, error) {
	var hash string
	//prepare the params of the function
	params := func() string {
		raw, err := transaction.MarshalJSON()
		if err != nil {
			logger.Error("failed to marshal transaction data: ", err)
		}
		return string(raw)
	}
	err := rpc.post("eth_sendTransaction", &hash, params)
	return hash, err
}

// EthSendRawTransaction creates new message post transaction
// or a contract creation for signed transactions.
func (rpc *EthRPC) EthSendRawTransaction(data string) (string, error) {
	var hash string
	//prepare the params of the function
	params := func() string {
		return data
	}
	err := rpc.post("eth_sendRawTransaction", &hash, params)
	return hash, err
}

// EthCall executes a new message post immediately without
// creating a transaction on the block chain.
func (rpc *EthRPC) EthCall(transaction TransactionData, tag string) (string, error) {
	var data string
	//prepare the params of the function
	params := func() string {
		raw, err := transaction.MarshalJSON()
		if err != nil {
			logger.Error("failed to marshal transaction data: ", err)
		}
		return string(raw) + "," + tag
	}
	err := rpc.post("eth_call", &data, params)
	return data, err
}

// EthEstimateGas makes a post or transaction, which won't be
// added to the blockchain and returns the used gas, which can
// be used for estimating the used gas.
func (rpc *EthRPC) EthEstimateGas(transaction TransactionData) (int, error) {
	var response string
	//prepare the params of the function
	params := func() string {
		raw, err := transaction.MarshalJSON()
		if err != nil {
			logger.Error("failed to marshal transaction data: ", err)
		}
		return string(raw)
	}
	err := rpc.post("eth_estimateGas", &response, params)
	if err != nil {
		return 0, err
	}
	return ParseInt(response)
}

// getBlock gets current block information
func (rpc *EthRPC) getBlock(method string, withTransactions bool, params string) (*Block, error) {
	result, err := rpc.makePostWithMethodParams(method, params)
	if err != nil {
		return nil, err
	}
	if bytes.Equal(result, []byte("null")) {
		return nil, nil
	}

	var response proxyBlock
	if withTransactions {
		response = new(proxyBlockWithTransactions)
	} else {
		response = new(proxyBlockWithoutTransactions)
	}

	err = json.Unmarshal(result, response)
	if err != nil {
		return nil, err
	}

	block := response.toBlock()
	return &block, nil
}

// EthGetBlockByHash returns information about a block by hash.
func (rpc *EthRPC) EthGetBlockByHash(hash string, withTransactions bool) (*Block, error) {
	params := hash
	return rpc.getBlock("eth_getBlockByHash", withTransactions, params)
}

// EthGetBlockByNumber returns information about a block by block number.
func (rpc *EthRPC) EthGetBlockByNumber(number int, withTransactions bool) (*Block, error) {
	params := IntToHex(number)
	return rpc.getBlock("eth_getBlockByNumber", withTransactions, params)
}

func (rpc *EthRPC) getTransaction(method string, params ParamsCallback) (*Transaction, error) {
	transaction := new(Transaction)
	err := rpc.post(method, transaction, params)
	return transaction, err
}

// EthGetTransactionByHash returns the information about a transaction requested by transaction hash.
func (rpc *EthRPC) EthGetTransactionByHash(hash string) (*Transaction, error) {
	params := func() string {
		return "[" + rpc.doubleQuote(hash) + "]"
	}
	return rpc.getTransaction("eth_getTransactionByHash", params)
}

// EthGetTransactionByBlockHashAndIndex returns information about a transaction by block hash and transaction index position.
func (rpc *EthRPC) EthGetTransactionByBlockHashAndIndex(blockHash string, transactionIndex int) (*Transaction, error) {
	params := func() string {
		return blockHash + "," + IntToHex(transactionIndex)
	}
	return rpc.getTransaction("eth_getTransactionByBlockHashAndIndex", params)
}

// EthGetTransactionByBlockNumberAndIndex returns information about a transaction by block number and transaction index position.
func (rpc *EthRPC) EthGetTransactionByBlockNumberAndIndex(blockNumber, transactionIndex int) (*Transaction, error) {
	params := func() string {
		return IntToHex(blockNumber) + "," + IntToHex(transactionIndex)
	}
	return rpc.getTransaction("eth_getTransactionByBlockNumberAndIndex", params)
}

// EthGetTransactionReceipt returns the receipt of a transaction by transaction hash.
// Note That the receipt is not available for pending transactions.
func (rpc *EthRPC) EthGetTransactionReceipt(hash string) (*TransactionReceipt, error) {
	transactionReceipt := new(TransactionReceipt)
	params := func() string {
		return rpc.doubleQuote(hash)
	}
	err := rpc.post("eth_getTransactionReceipt", transactionReceipt, params)
	if err != nil {
		return nil, err
	}
	return transactionReceipt, nil
}

// TODO implement
// EthGetPendingTransactions returns the list of pending transactions
func (rpc *EthRPC) EthGetPendingTransactions(hash string) (*TransactionReceipt, error) {
	transactionReceipt := new(TransactionReceipt)
	params := func() string {
		return hash
	}
	err := rpc.post("eth_pendingTransactions", transactionReceipt, params)
	if err != nil {
		return nil, err
	}
	return transactionReceipt, nil
}

// EthGetCompilers returns a list of available compilers in the client.
// @deprecated
func (rpc *EthRPC) EthGetCompilers() ([]string, error) {
	var compilers []string
	params := func() string {
		return ""
	}
	err := rpc.post("eth_getCompilers", &compilers, params)
	return compilers, err
}

// TODO implement
// eth_compileSolidity
// @deprecated
func (rpc *EthRPC) EthCompileSolidity() ([]string, error) {
	var compilers []string
	params := func() string {
		return ""
	}
	err := rpc.post("eth_compileSolidity", &compilers, params)
	return compilers, err
}

// EthNewFilter creates a new filter object.
func (rpc *EthRPC) EthNewFilter(filter FilterParams) (string, error) {
	var filterID string
	params := func() string {
		return filter.String()
	}
	err := rpc.post("eth_newFilter", &filterID, params)
	return filterID, err
}

// EthNewBlockFilter creates a filter in the node, to notify when a new block arrives.
// To check if the state has changed, post EthGetFilterChanges.
func (rpc *EthRPC) EthNewBlockFilter() (string, error) {
	var filterID string
	params := func() string {
		return ""
	}
	err := rpc.post("eth_newBlockFilter", &filterID, params)
	return filterID, err
}

// EthNewPendingTransactionFilter creates a filter in the node, to notify when new pending transactions arrive.
// To check if the state has changed, post EthGetFilterChanges.
func (rpc *EthRPC) EthNewPendingTransactionFilter() (string, error) {
	var filterID string
	params := func() string {
		return ""
	}
	err := rpc.post("eth_newPendingTransactionFilter", &filterID, params)
	return filterID, err
}

// EthUninstallFilter uninstalls a filter with given id.
func (rpc *EthRPC) EthUninstallFilter(filterID string) (bool, error) {
	var res bool
	params := func() string {
		return filterID
	}
	err := rpc.post("eth_uninstallFilter", &res, params)
	return res, err
}

// EthGetFilterChanges polling method for a filter, which returns an array of logs which occurred since last poll.
func (rpc *EthRPC) EthGetFilterChanges(filterID string) ([]Log, error) {
	var logs []Log
	params := func() string {
		return filterID
	}
	err := rpc.post("eth_getFilterChanges", &logs, params)
	return logs, err
}

// EthGetFilterLogs returns an array of all logs matching filter with given id.
func (rpc *EthRPC) EthGetFilterLogs(filterID string) ([]Log, error) {
	var logs []Log
	params := func() string {
		return filterID
	}
	err := rpc.post("eth_getFilterLogs", &logs, params)
	return logs, err
}

// EthGetLogs returns an array of all logs matching a given filter object.
func (rpc *EthRPC) EthGetLogs(filter FilterParams) ([]Log, error) {
	var logs []Log
	params := func() string {
		return filter.String()
	}
	err := rpc.post("eth_getLogs", &logs, params)
	return logs, err
}

// TODO implement
// EthGetWork returns an array of all logs matching a given filter object.
func (rpc *EthRPC) EthGetWork(filter FilterParams) ([]Log, error) {
	var logs []Log
	params := func() string {
		return filter.String()
	}
	err := rpc.post("eth_getWork", &logs, params)
	return logs, err
}

// TODO implement
// EthSubmitWork
func (rpc *EthRPC) EthSubmitWork(filter FilterParams) ([]Log, error) {
	var logs []Log
	params := func() string {
		return filter.String()
	}
	err := rpc.post("eth_submitWork", &logs, params)
	return logs, err
}

// TODO implement
// EthSubmitHashrate
func (rpc *EthRPC) EthSubmitHashrate(filter FilterParams) ([]Log, error) {
	var logs []Log
	params := func() string {
		return filter.String()
	}
	err := rpc.post("eth_submitHashrate", &logs, params)
	return logs, err
}

// TODO implement
// EthGetProof
func (rpc *EthRPC) EthGetProof(filter FilterParams) ([]Log, error) {
	var logs []Log
	params := func() string {
		return filter.String()
	}
	err := rpc.post("eth_getProof", &logs, params)
	return logs, err
}

func (rpc *EthRPC) generateTransactionPayload(contract string, data string, block string, gas string, gasprice string, params *model.EthRequestParams) string {
	requestParams := map[string]interface{}{
		"to":   contract,
		"data": data,
		/*
			"gas":      "0xaae60", //700000,
			"gasPrice": "0x15f90", //90000,
		*/
	}
	if gas != "" {
		requestParams["gas"] = gas
	}
	if gasprice != "" {
		requestParams["gasPrice"] = gasprice
	}
	raw, _ := json.Marshal(requestParams)
	paramsStr := str.UnsafeString(raw)
	request := `{"id":1, "jsonrpc":"2.0","method":"eth_call","params":[` + paramsStr + `]}`
	return request
}

// todo add from field
func (rpc *EthRPC) generateCallPayload(contract string, data string, block string) string {
	if block != model.NoPeriod {
		request := `{"id": 1,"jsonrpc": "2.0","method": "eth_call",
"params":[{
"to": ` + rpc.doubleQuote(contract) + `,
"data": ` + rpc.doubleQuote(data) + `},
` + rpc.doubleQuote(block) + `]}`
		return request
	} else {
		request := `{"id": 1,"jsonrpc": "2.0","method": "eth_call",
"params":[{
"to": ` + rpc.doubleQuote(contract) + `,
"data": ` + rpc.doubleQuote(data) + `}]}`
		return request
	}
}

// this method converts standard contract params to abi encoded params given a
// contract address, method name and abi model
func (rpc *EthRPC) convertParamsToAbi(contract string, method string, args interface{}) ([]byte, error) {
	var abiModel abi.ABI
	//try to fetch the abi model linked to given contract address
	return abiModel.Pack(method, args)
}

// post ethereum network contract with no parameters
func (rpc *EthRPC) ContractCall(contract string, methodName string, params string, block string, gas string, gasprice string) (string, error) {
	abiparams, abiEncErr := rpc.convertParamsToAbi(contract, methodName, params)
	if abiEncErr != nil {
		//failed to encode post abi data
		logger.Error("failed to encode contract post abi parameters: ", abiEncErr)
		return "", abiEncErr
	} else {
		paramsStr := str.UnsafeString(abiparams)
		data := paramsStr + "," + "," + gas + "," + gasprice
		payload := rpc.generateCallPayload(contract, data, block)
		raw, err := rpc.makePostRaw(payload)
		if err == nil {
			var data string
			unErr := json.Unmarshal(raw, &data)
			return data, unErr
		}
		return "", err
	}
}

// post ethereum network contract with no parameters
func (rpc *EthRPC) contractCallAbiParams(contract string, data string, block string) (string, error) {
	payload := rpc.generateCallPayload(contract, data, block)
	raw, err := rpc.makePostRaw(payload)
	if err == nil {
		var data string
		unErr := json.Unmarshal(raw, &data)
		return data, unErr
	}
	return "", err
}

func (rpc *EthRPC) Erc20Summary(contract string) (map[string]string, error) {
	var response = map[string]string{
		summaryFunctionsNames[0]: "",
		summaryFunctionsNames[1]: "",
		summaryFunctionsNames[2]: "",
		summaryFunctionsNames[3]: "",
	}
	// copy target url to summary functions
	instance.url = rpc.url
	// execute erc20 summary functions
	for i := 0; i < len(summaryFunctions); i++ {
		raw, err := summaryFunctions[i](contract)
		if err == nil && raw != "" {
			response[summaryFunctionsNames[i]] = raw
		}
	}
	return response, nil
}

func (rpc *EthRPC) Erc20TotalSupply(contract string) (string, error) {
	return rpc.contractCallAbiParams(contract, paramencoder.TotalSupplyParams, model.LatestBlockNumber)
}

func (rpc *EthRPC) Erc20Symbol(contract string) (string, error) {
	return rpc.contractCallAbiParams(contract, paramencoder.SymbolParams, model.LatestBlockNumber)
}

func (rpc *EthRPC) Erc20Name(contract string) (string, error) {
	return rpc.contractCallAbiParams(contract, paramencoder.NameParams, model.LatestBlockNumber)
}

func (rpc *EthRPC) Erc20Decimals(contract string) (string, error) {
	return rpc.contractCallAbiParams(contract, paramencoder.DecimalsParams, model.LatestBlockNumber)
}

func (rpc *EthRPC) Erc20BalanceOf(contract string, tokenOwner string) (json.RawMessage, error) {
	tokenOwnerAddress, decodeErr := fromStringToAddress(tokenOwner)
	if decodeErr != nil {
		logger.Error("failed to read and decode provided Ethereum address", decodeErr)
		return nil, decodeErr
	}
	abiparams, encErr := paramencoder.LoadErc20Abi().Pack("balanceOf", tokenOwnerAddress)
	if encErr != nil {
		logger.Error("failed to encode ABI parameters for ERC20 balanceof method", encErr)
		return nil, encErr
	}
	// encode to hexadecimal abiparams
	dataContent := hex.ToEthHex(abiparams)
	req := rpc.generateCallPayload(contract, dataContent, model.LatestBlockNumber)
	return rpc.makePostRaw(req)
}

func (rpc *EthRPC) Erc20Allowance(contract string, tokenOwner string, spender string) (json.RawMessage, error) {
	tokenOwnerAddress, decodeErr := fromStringToAddress(tokenOwner)
	if decodeErr != nil {
		logger.Error("failed to read and decode provided Ethereum address", decodeErr)
		return nil, decodeErr
	}
	spenderAddress, decodeErr := fromStringToAddress(spender)
	if decodeErr != nil {
		logger.Error("failed to read and decode provided Ethereum address", decodeErr)
		return nil, decodeErr
	}
	abiparams, encErr := paramencoder.LoadErc20Abi().Pack("allowance", tokenOwnerAddress, spenderAddress)
	if encErr != nil {
		logger.Error("failed to encode ABI parameters for ERC20 allowance method", encErr)
		return nil, encErr
	}
	// encode to hexadecimal abiparams
	dataContent := hex.ToEthHex(abiparams)
	params := model.EthRequestParams{
		To:   contract,
		Data: dataContent,
		Tag:  model.LatestBlockNumber,
	}
	return rpc.makePostWithMethodParams("eth_call", params.String())
}

func (rpc *EthRPC) Erc20Transfer(contract string, address string, amount int) (json.RawMessage, error) {
	senderAddress, decodeErr := fromStringToAddress(address)
	if decodeErr != nil {
		logger.Error("failed to read and decode provided Ethereum address", decodeErr)
		return nil, decodeErr
	}
	abiparams, encErr := paramencoder.LoadErc20Abi().Pack("transfer", senderAddress, amount)
	if encErr != nil {
		logger.Error("failed to encode ABI parameters for ERC20 transfer method", encErr)
		return nil, encErr
	}
	// encode to hexadecimal abiparams
	dataContent := hex.ToEthHex(abiparams)
	params := model.EthRequestParams{
		To:   contract,
		Data: dataContent,
		Tag:  model.LatestBlockNumber,
	}
	return rpc.makePostWithMethodParams("eth_sendTransaction", params.String())
}

// curl localhost:8545 -X POST --data '{"jsonrpc":"2.0","method":"eth_sendTransaction","params":[{"from": "0x8aff0a12f3e8d55cc718d36f84e002c335df2f4a", "data": "606060405260728060106000396000f360606040526000357c0100000000000000000000000000000000000000000000000000000000900480636ffa1caa146037576035565b005b604b60048080359060200190919050506061565b6040518082815260200191505060405180910390f35b6000816002029050606d565b91905056"}],"id":1}
func (rpc *EthRPC) DeployContract(fromAddress string, bytecode string, gas string, gasPrice string) (json.RawMessage, error) {
	payload := `[{"from":` + rpc.doubleQuote(fromAddress) + `,
"data":` + rpc.doubleQuote(bytecode) + `,
"gasprice":` + rpc.doubleQuote(gasPrice) + `,
"gas":` + rpc.doubleQuote(gas) + `}]`
	return rpc.makePostWithMethodParams("eth_sendTransaction", payload)
}

func (rpc *EthRPC) IsSmartContractAddress(addr string) (bool, error) {
	bytecode, err := rpc.EthGetCode(addr, model.LatestBlockNumber)
	// if the address has valid bytecode, is a contract
	// if is not code addres 0x is returned
	return len(bytecode) > 2, err
}

// helper methods

func fromStringToAddress(addr string) (fixtures.Address, error) {
	var a fixtures.Address
	raw, decodeErr := hex.FromEthHex(addr)
	if decodeErr != nil {
		logger.Error("failed to read and decode provided Ethereum address", decodeErr)
		return a, decodeErr
	}
	a.SetBytes(raw)
	return a, nil
}

// @deprecated
// Eth1 returns 1 ethereum value (10^18 wei)
func (rpc *EthRPC) Eth1() *big.Int {
	return Eth1()
}

// Eth1 returns 1 ethereum value (10^18 wei)
func Eth1() *big.Int {
	return oneEth
}

// Eth1 returns 1 ethereum value (10^18 wei)
func Eth1Int64() int64 {
	return oneEthInt64
}
