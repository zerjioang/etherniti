// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package ethrpc

import (
	"bytes"
	"encoding/json"
	"github.com/zerjioang/etherniti/core/eth/fixtures/abi"
	"github.com/zerjioang/etherniti/core/eth/rpc/model"
	"io/ioutil"
	"math/big"
	"net/http"

	"github.com/zerjioang/etherniti/core/util/str"

	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/modules/encoding/hex"

	"github.com/zerjioang/etherniti/core/eth/paramencoder"

	"github.com/zerjioang/etherniti/core/eth/fixtures"

	"github.com/zerjioang/etherniti/thirdparty/gommon/log"
)

// https://documenter.getpostman.com/view/4117254/ethereum-json-rpc/RVu7CT5J
type contractFunction func(string) (string, error)

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
	oneEth = big.NewInt(1000000000000000000)
)

// EthRPC - Ethereum rpc client
type EthRPC struct {
	//ethereum or quorum node endpoint
	url    string
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
		client: http.Client{},
		Debug:  true,
	}
	return rpc
}

func (rpc *EthRPC) call(method string, target interface{}, params *model.EthRequestParams) error {
	result, err := rpc.makePost(method, params, model.LatestBlockNumber)
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

// makePost returns raw response of method call
/*

eth_call

Executes a new message call immediately without creating a transaction on the block chain.
Parameters

    Object - The transaction call object

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
func (rpc *EthRPC) makePost(method string, params *model.EthRequestParams, period model.BlockPeriod) (json.RawMessage, error) {
	request := model.EthRequest{
		ID:      1,
		JSONRPC: "2.0",
		Method:  method,
		Params:  params,
	}
	if period != model.NoPeriod {
		request.SetBlockPeriod(period)
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	return rpc.makePostRaw(body)
}

func (rpc *EthRPC) makePostRaw(data []byte) (json.RawMessage, error) {

	log.Info("sending request: ", str.UnsafeString(data))
	response, err := rpc.client.Post(rpc.url, "application/json", bytes.NewBuffer(data))
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

// RawCall returns raw response of method call (Deprecated)
func (rpc *EthRPC) RawCall(method string, params *model.EthRequestParams) (json.RawMessage, error) {
	return rpc.makePost(method, params, model.LatestBlockNumber)
}

func (rpc *EthRPC) EthMethodNoParams(methodName string) (interface{}, error) {
	var response interface{}
	err := rpc.call(methodName, &response, nil)
	return response, err
}

// returns ethereum node information
func (rpc *EthRPC) EthNodeInfo() (string, error) {
	var response string

	err := rpc.call("eth_info", &response, nil)
	return response, err
}

// Web3ClientVersion returns the current client version.
func (rpc *EthRPC) Web3ClientVersion() (string, error) {
	var clientVersion string

	err := rpc.call("web3_clientVersion", &clientVersion, nil)
	return clientVersion, err
}

// Web3Sha3 returns Keccak-256 (not the standardized SHA3-256) of the given data.
func (rpc *EthRPC) Web3Sha3(data []byte) (string, error) {
	var hash string

	hashData := fixtures.Encode(data)
	err := rpc.call("web3_sha3", &hash, hashData)
	return hash, err
}

// NetVersion returns the current network protocol version.
func (rpc *EthRPC) NetVersion() (string, error) {
	var version string

	err := rpc.call("net_version", &version, nil)
	return version, err
}

// NetListening returns true if client is actively listening for network connections.
func (rpc *EthRPC) NetListening() (bool, error) {
	var listening bool

	err := rpc.call("net_listening", &listening, nil)
	return listening, err
}

// NetPeerCount returns number of peers currently connected to the client.
func (rpc *EthRPC) NetPeerCount() (int, error) {
	var response int
	if err := rpc.call("net_peerCount", &response, nil); err != nil {
		return 0, err
	}

	return response, nil
}

// EthProtocolVersion returns the current ethereum protocol version.
func (rpc *EthRPC) EthProtocolVersion() (string, error) {
	var protocolVersion string

	err := rpc.call("eth_protocolVersion", &protocolVersion, nil)
	return protocolVersion, err
}

// EthSyncing returns an object with data about the sync status or false.
func (rpc *EthRPC) EthSyncing() (*Syncing, error) {
	result, err := rpc.RawCall("eth_syncing" ,nil)
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

	err := rpc.call("net_version", &response, nil)
	return response, err
}

// EthCoinbase returns the client coinbase address
func (rpc *EthRPC) EthCoinbase() (string, error) {
	var address string

	err := rpc.call("eth_coinbase", &address, nil)
	return address, err
}

// EthMining returns true if client is actively mining new blocks.
func (rpc *EthRPC) EthMining() (bool, error) {
	var mining bool

	err := rpc.call("eth_mining", &mining, nil)
	return mining, err
}

// EthHashrate returns the number of hashes per second that the node is mining with.
func (rpc *EthRPC) EthHashrate() (int, error) {
	var response string

	if err := rpc.call("eth_hashrate", &response, nil); err != nil {
		return 0, err
	}

	return ParseInt(response)
}

// EthGasPrice returns the current price per gas in wei.
func (rpc *EthRPC) EthGasPrice() (int64, error) {
	var response string
	if err := rpc.call("eth_gasPrice", &response, nil); err != nil {
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
	err := rpc.call("eth_accounts", &accounts, nil)
	return accounts, err
}

// EthBlockNumber returns the number of most recent block.
func (rpc *EthRPC) EthBlockNumber() (int, error) {
	var response string
	if err := rpc.call("eth_blockNumber", &response, nil); err != nil {
		return 0, err
	}

	return ParseInt(response)
}

// EthGetBalance returns the balance of the account of given address in wei.
func (rpc *EthRPC) EthGetBalance(address, block model.BlockPeriod) (*big.Int, error) {
	var response string
	if err := rpc.call("eth_getBalance", &response, address, block); err != nil {
		return new(big.Int), err
	}
	return ParseBigInt(response)
}

// EthGetStorageAt returns the value from a storage position at a given address.
func (rpc *EthRPC) EthGetStorageAt(data string, position int, tag string) (string, error) {
	var result string
	err := rpc.call("eth_getStorageAt", &result, data, IntToHex(position), tag)
	return result, err
}

// EthGetTransactionCount returns the number of transactions sent from an address.
func (rpc *EthRPC) EthGetTransactionCount(address, block model.BlockPeriod) (int, error) {
	var response string
	if err := rpc.call("eth_getTransactionCount", &response, address, block); err != nil {
		return 0, err
	}
	return ParseInt(response)
}

// EthGetBlockTransactionCountByHash returns the number of transactions in a block from a block matching the given block hash.
func (rpc *EthRPC) EthGetBlockTransactionCountByHash(hash string) (int, error) {
	var response string
	if err := rpc.call("eth_getBlockTransactionCountByHash", &response, hash); err != nil {
		return 0, err
	}
	return ParseInt(response)
}

// EthGetBlockTransactionCountByNumber returns the number of transactions in a block from a block matching the given block
func (rpc *EthRPC) EthGetBlockTransactionCountByNumber(number int) (int, error) {
	var response string
	if err := rpc.call("eth_getBlockTransactionCountByNumber", &response, IntToHex(number)); err != nil {
		return 0, err
	}
	return ParseInt(response)
}

// EthGetUncleCountByBlockHash returns the number of uncles in a block from a block matching the given block hash.
func (rpc *EthRPC) EthGetUncleCountByBlockHash(hash string) (int, error) {
	var response string
	if err := rpc.call("eth_getUncleCountByBlockHash", &response, hash); err != nil {
		return 0, err
	}
	return ParseInt(response)
}

// EthGetUncleCountByBlockNumber returns the number of uncles in a block from a block matching the given block number.
func (rpc *EthRPC) EthGetUncleCountByBlockNumber(number int) (int, error) {
	var response string
	if err := rpc.call("eth_getUncleCountByBlockNumber", &response, IntToHex(number)); err != nil {
		return 0, err
	}
	return ParseInt(response)
}

// EthGetCode returns code at a given address.
func (rpc *EthRPC) EthGetCode(address string, block model.BlockPeriod) (string, error) {
	var code string
	err := rpc.call("eth_getCode", &code, address, block)
	return code, err
}

// EthSign signs data with a given address.
// Calculates an Ethereum specific signature
// with: sign(keccak256("\x19Ethereum Signed Message:\n" + len(message) + message)))
func (rpc *EthRPC) EthSign(address, data string) (string, error) {
	var signature string
	err := rpc.call("eth_sign", &signature, address, data)
	return signature, err
}

// EthSendTransaction creates new message call transaction
// or a contract creation, if the data field contains code.
func (rpc *EthRPC) EthSendTransaction(transaction TransactionData) (string, error) {
	var hash string
	err := rpc.call("eth_sendTransaction", &hash, transaction)
	return hash, err
}

// EthSendRawTransaction creates new message call transaction
// or a contract creation for signed transactions.
func (rpc *EthRPC) EthSendRawTransaction(data string) (string, error) {
	var hash string
	err := rpc.call("eth_sendRawTransaction", &hash, data)
	return hash, err
}

// EthCall executes a new message call immediately without
// creating a transaction on the block chain.
func (rpc *EthRPC) EthCall(transaction TransactionData, tag string) (string, error) {
	var data string
	err := rpc.call("eth_call", &data, transaction, tag)
	return data, err
}

// EthEstimateGas makes a call or transaction, which won't be
// added to the blockchain and returns the used gas, which can
// be used for estimating the used gas.
func (rpc *EthRPC) EthEstimateGas(transaction TransactionData) (int, error) {
	var response string
	err := rpc.call("eth_estimateGas", &response, transaction)
	if err != nil {
		return 0, err
	}
	return ParseInt(response)
}

// getBlock gets current block information
func (rpc *EthRPC) getBlock(method string, withTransactions bool, params *model.EthRequestParams) (*Block, error) {
	result, err := rpc.RawCall(method, params)
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
	return rpc.getBlock("eth_getBlockByHash", withTransactions, hash, withTransactions)
}

// EthGetBlockByNumber returns information about a block by block number.
func (rpc *EthRPC) EthGetBlockByNumber(number int, withTransactions bool) (*Block, error) {
	return rpc.getBlock("eth_getBlockByNumber", withTransactions, IntToHex(number), withTransactions)
}

func (rpc *EthRPC) getTransaction(method string, params *model.EthRequestParams) (*Transaction, error) {
	transaction := new(Transaction)

	err := rpc.call(method, transaction, params)
	return transaction, err
}

// EthGetTransactionByHash returns the information about a transaction requested by transaction hash.
func (rpc *EthRPC) EthGetTransactionByHash(hash string) (*Transaction, error) {
	return rpc.getTransaction("eth_getTransactionByHash", hash)
}

// EthGetTransactionByBlockHashAndIndex returns information about a transaction by block hash and transaction index position.
func (rpc *EthRPC) EthGetTransactionByBlockHashAndIndex(blockHash string, transactionIndex int) (*Transaction, error) {
	return rpc.getTransaction("eth_getTransactionByBlockHashAndIndex", blockHash, IntToHex(transactionIndex))
}

// EthGetTransactionByBlockNumberAndIndex returns information about a transaction by block number and transaction index position.
func (rpc *EthRPC) EthGetTransactionByBlockNumberAndIndex(blockNumber, transactionIndex int) (*Transaction, error) {
	return rpc.getTransaction("eth_getTransactionByBlockNumberAndIndex", IntToHex(blockNumber), IntToHex(transactionIndex))
}

// EthGetTransactionReceipt returns the receipt of a transaction by transaction hash.
// Note That the receipt is not available for pending transactions.
func (rpc *EthRPC) EthGetTransactionReceipt(hash string) (*TransactionReceipt, error) {
	transactionReceipt := new(TransactionReceipt)
	err := rpc.call("eth_getTransactionReceipt", transactionReceipt, hash)
	if err != nil {
		return nil, err
	}
	return transactionReceipt, nil
}

// TODO implement
// EthGetPendingTransactions returns the list of pending transactions
func (rpc *EthRPC) EthGetPendingTransactions(hash string) (*TransactionReceipt, error) {
	transactionReceipt := new(TransactionReceipt)
	err := rpc.call("eth_pendingTransactions", transactionReceipt, hash)
	if err != nil {
		return nil, err
	}
	return transactionReceipt, nil
}

// EthGetCompilers returns a list of available compilers in the client.
// @deprecated
func (rpc *EthRPC) EthGetCompilers() ([]string, error) {
	var compilers []string
	err := rpc.call("eth_getCompilers", &compilers)
	return compilers, err
}

// TODO implement
// eth_compileSolidity
// @deprecated
func (rpc *EthRPC) EthCompileSolidity() ([]string, error) {
	var compilers []string
	err := rpc.call("eth_compileSolidity", &compilers)
	return compilers, err
}

// EthNewFilter creates a new filter object.
func (rpc *EthRPC) EthNewFilter(params FilterParams) (string, error) {
	var filterID string
	err := rpc.call("eth_newFilter", &filterID, params)
	return filterID, err
}

// EthNewBlockFilter creates a filter in the node, to notify when a new block arrives.
// To check if the state has changed, call EthGetFilterChanges.
func (rpc *EthRPC) EthNewBlockFilter() (string, error) {
	var filterID string
	err := rpc.call("eth_newBlockFilter", &filterID)
	return filterID, err
}

// EthNewPendingTransactionFilter creates a filter in the node, to notify when new pending transactions arrive.
// To check if the state has changed, call EthGetFilterChanges.
func (rpc *EthRPC) EthNewPendingTransactionFilter() (string, error) {
	var filterID string
	err := rpc.call("eth_newPendingTransactionFilter", &filterID)
	return filterID, err
}

// EthUninstallFilter uninstalls a filter with given id.
func (rpc *EthRPC) EthUninstallFilter(filterID string) (bool, error) {
	var res bool
	err := rpc.call("eth_uninstallFilter", &res, filterID)
	return res, err
}

// EthGetFilterChanges polling method for a filter, which returns an array of logs which occurred since last poll.
func (rpc *EthRPC) EthGetFilterChanges(filterID string) ([]Log, error) {
	var logs []Log
	err := rpc.call("eth_getFilterChanges", &logs, filterID)
	return logs, err
}

// EthGetFilterLogs returns an array of all logs matching filter with given id.
func (rpc *EthRPC) EthGetFilterLogs(filterID string) ([]Log, error) {
	var logs []Log
	err := rpc.call("eth_getFilterLogs", &logs, filterID)
	return logs, err
}

// EthGetLogs returns an array of all logs matching a given filter object.
func (rpc *EthRPC) EthGetLogs(params FilterParams) ([]Log, error) {
	var logs []Log
	err := rpc.call("eth_getLogs", &logs, params)
	return logs, err
}

// TODO implement
// EthGetWork returns an array of all logs matching a given filter object.
func (rpc *EthRPC) EthGetWork(params FilterParams) ([]Log, error) {
	var logs []Log
	err := rpc.call("eth_getWork", &logs, params)
	return logs, err
}

// TODO implement
// EthSubmitWork
func (rpc *EthRPC) EthSubmitWork(params FilterParams) ([]Log, error) {
	var logs []Log
	err := rpc.call("eth_submitWork", &logs, params)
	return logs, err
}

// TODO implement
// EthSubmitHashrate
func (rpc *EthRPC) EthSubmitHashrate(params FilterParams) ([]Log, error) {
	var logs []Log
	err := rpc.call("eth_submitHashrate", &logs, params)
	return logs, err
}

// TODO implement
// EthGetProof
func (rpc *EthRPC) EthGetProof(params FilterParams) ([]Log, error) {
	var logs []Log
	err := rpc.call("eth_getProof", &logs, params)
	return logs, err
}

/*

db_putString

db_getString

db_putHex

db_getHex

shh_version

shh_post

shh_newIdentity

shh_hasIdentity

shh_newGroup

shh_addToGroup

shh_newFilter

shh_uninstallFilter

shh_getFilterChanges

shh_getMessages

*/

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

func (rpc *EthRPC) generateTransactionPayload(contract string, data string, block model.BlockPeriod, gas string, gasprice string, params *model.EthRequestParams) string {
	requestParams := map[string]interface{}{
		"to":   contract,
		"data": data,
		/*
		"gas":      "0xaae60", //700000,
		"gasPrice": "0x15f90", //90000,
		*/
	}
	if gas != "" {
		requestParams["gas"]=gas
	}
	if gasprice != "" {
		requestParams["gasPrice"]=gasprice
	}
	raw, _ := json.Marshal(requestParams)
	paramsStr := str.UnsafeString(raw)
	request := `{"id":1, "jsonrpc":"2.0","method":"eth_call","params":[`+paramsStr+`]}`
	return request
}

func (rpc *EthRPC) generateCallPayload(contract string, data string) []byte {
	request := model.EthRequest{
		ID:      1,
		JSONRPC: "2.0",
		Method:  "eth_call",
		Params:  &model.EthRequestParams{
			To:contract,
			Data:data,
		},
	}

	body, _ := json.Marshal(request)
	return body
}

// this method converts standard contract params to abi encoded params given a
// contract address, method name and abi model
func (rpc *EthRPC) convertParamsToAbi(contract string, method string, args interface{}) ([]byte, error) {
	var abiModel abi.ABI
	//try to fetch the abi model linked to given contract address
	return abiModel.Pack(method, args)
}

// call ethereum network contract with no parameters
func (rpc *EthRPC) ContractCall(contract string, methodName string, params string, block model.BlockPeriod, gas string, gasprice string) (string, error) {
	abiparams, abiEncErr := rpc.convertParamsToAbi(contract, methodName, params)
	if abiEncErr != nil {
		//failed to encode call abi data
		logger.Error("failed to encode contract call abi parameters: ", abiEncErr)
		return "", abiEncErr
	} else {
		paramsStr := str.UnsafeString(abiparams)
		payload := rpc.generateCallPayload(contract, "eth_call", paramsStr, block, gas, gasprice)
		raw, err := rpc.makePostRaw(payload)
		if err == nil {
			var data string
			unErr := json.Unmarshal(raw, &data)
			return data, unErr
		}
		return "", err
	}
}

// call ethereum network contract with no parameters
func (rpc *EthRPC) contractCallAbiParams(contract string, data string, block model.BlockPeriod) (string, error) {
	payload := rpc.generateCallPayload(contract, data, block)
	raw, err := rpc.makePostRaw(payload)
	if err == nil {
		var data string
		unErr := json.Unmarshal(raw, &data)
		return data, unErr
	}
	return "", err
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
	return rpc.makePost("eth_call", &model.EthRequestParams{
		To:   contract,
		Data: dataContent,
	}, model.LatestBlockNumber)
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
	return rpc.makePost("eth_call", &model.EthRequestParams{
		To:   contract,
		Data: dataContent,
	}, model.LatestBlockNumber)
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
	return rpc.makePost("eth_sendTransaction", &model.EthRequestParams{
		To:   contract,
		Data: dataContent,
	}, model.LatestBlockNumber)
}

// curl localhost:8545 -X POST --data '{"jsonrpc":"2.0","method":"eth_sendTransaction","params":[{"from": "0x8aff0a12f3e8d55cc718d36f84e002c335df2f4a", "data": "606060405260728060106000396000f360606040526000357c0100000000000000000000000000000000000000000000000000000000900480636ffa1caa146037576035565b005b604b60048080359060200190919050506061565b6040518082815260200191505060405180910390f35b6000816002029050606d565b91905056"}],"id":1}
func (rpc *EthRPC) DeployContract(fromAddress string, bytecode string) (json.RawMessage, error) {
	return rpc.makePost("eth_sendTransaction", &model.EthRequestParams{
		From: fromAddress,
		Data: bytecode,
	}, model.NoPeriod)
}

func (rpc *EthRPC) IsSmartContractAddress(addr string) (bool, error) {
	contractAddress, decodeErr := fromStringToAddress(addr)
	if decodeErr != nil {
		logger.Error("failed to read and decode provided Ethereum contract address", decodeErr)
		return false, decodeErr
	}

	bytecode, err := rpc.EthGetCode(contractAddress.Hex(), model.LatestBlockNumber)
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

// Eth1 returns 1 ethereum value (10^18 wei)
func (rpc *EthRPC) Eth1() *big.Int {
	return Eth1()
}

// Eth1 returns 1 ethereum value (10^18 wei)
func Eth1() *big.Int {
	return oneEth
}
