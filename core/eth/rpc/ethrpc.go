// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package ethrpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"

	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/modules/encoding/hex"

	"github.com/zerjioang/etherniti/core/eth/paramencoder"

	"github.com/zerjioang/etherniti/core/eth/fixtures"

	"github.com/labstack/gommon/log"
)

// EthError - ethereum error
type EthError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err EthError) Error() string {
	return "an error occurred: " + err.Message + " with code " + strconv.Itoa(err.Code)
}

var (
	oneEth = big.NewInt(1000000000000000000)
)

type ethResponse struct {
	ID      int             `json:"id"`
	JSONRPC string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result"`
	Error   *EthError       `json:"error"`
}

func (response ethResponse) Errored() error {
	return errors.New(response.Error.Message + ". Error code: " + strconv.Itoa(response.Error.Code))
}

type ethRequest struct {
	ID      int           `json:"id"`
	JSONRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

// EthRPC - Ethereum rpc client
type EthRPC struct {
	//ethereum or quorum node endpoint
	url    string
	client http.Client
	// debug flag
	Debug bool
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

func (rpc EthRPC) call(method string, target interface{}, params ...interface{}) error {
	result, err := rpc.Call(method, params...)
	if err != nil {
		return err
	}

	if target == nil || result == nil {
		return nil
	}

	return json.Unmarshal(result, target)
}

// URL returns client url
func (rpc EthRPC) URL() string {
	return rpc.url
}

// Call returns raw response of method call
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
func (rpc EthRPC) Call(method string, params ...interface{}) (json.RawMessage, error) {
	request := ethRequest{
		ID:      1,
		JSONRPC: "2.0",
		Method:  method,
		Params:  params,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	log.Info("sending request: ", string(body))
	response, err := rpc.client.Post(rpc.url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	//responseData readed, close body
	_ = response.Body.Close()

	log.Info("response received", string(responseData))

	resp := ethResponse{}
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
func (rpc EthRPC) RawCall(method string, params ...interface{}) (json.RawMessage, error) {
	return rpc.Call(method, params...)
}

// Web3ClientVersion returns the current client version.
func (rpc EthRPC) Web3ClientVersion() (string, error) {
	var clientVersion string

	err := rpc.call("web3_clientVersion", &clientVersion)
	return clientVersion, err
}

// Web3Sha3 returns Keccak-256 (not the standardized SHA3-256) of the given data.
func (rpc EthRPC) Web3Sha3(data []byte) (string, error) {
	var hash string

	hashData := fixtures.Encode(data)
	err := rpc.call("web3_sha3", &hash, hashData)
	return hash, err
}

// NetVersion returns the current network protocol version.
func (rpc EthRPC) NetVersion() (string, error) {
	var version string

	err := rpc.call("net_version", &version)
	return version, err
}

// NetListening returns true if client is actively listening for network connections.
func (rpc EthRPC) NetListening() (bool, error) {
	var listening bool

	err := rpc.call("net_listening", &listening)
	return listening, err
}

// NetPeerCount returns number of peers currently connected to the client.
func (rpc EthRPC) NetPeerCount() (int, error) {
	var response int
	if err := rpc.call("net_peerCount", &response); err != nil {
		return 0, err
	}

	return response, nil
}

// EthProtocolVersion returns the current ethereum protocol version.
func (rpc EthRPC) EthProtocolVersion() (string, error) {
	var protocolVersion string

	err := rpc.call("eth_protocolVersion", &protocolVersion)
	return protocolVersion, err
}

// EthSyncing returns an object with data about the sync status or false.
func (rpc EthRPC) EthSyncing() (*Syncing, error) {
	result, err := rpc.RawCall("eth_syncing")
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
func (rpc EthRPC) EthNodeInfo() (string, error) {
	var response string

	err := rpc.call("eth_info", &response)
	return response, err
}

func (rpc EthRPC) EthMethodNoParams(methodName string) (interface{}, error) {
	var response interface{}
	err := rpc.call(methodName, &response)
	return response, err
}

// returns ethereum node information
func (rpc EthRPC) EthNetVersion() (string, error) {
	var response string

	err := rpc.call("net_version", &response)
	return response, err
}

// EthCoinbase returns the client coinbase address
func (rpc EthRPC) EthCoinbase() (string, error) {
	var address string

	err := rpc.call("eth_coinbase", &address)
	return address, err
}

// EthMining returns true if client is actively mining new blocks.
func (rpc EthRPC) EthMining() (bool, error) {
	var mining bool

	err := rpc.call("eth_mining", &mining)
	return mining, err
}

// EthHashrate returns the number of hashes per second that the node is mining with.
func (rpc EthRPC) EthHashrate() (int, error) {
	var response string

	if err := rpc.call("eth_hashrate", &response); err != nil {
		return 0, err
	}

	return ParseInt(response)
}

// EthGasPrice returns the current price per gas in wei.
func (rpc EthRPC) EthGasPrice() (int64, error) {
	var response string
	if err := rpc.call("eth_gasPrice", &response); err != nil {
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
func (rpc EthRPC) EthAccounts() ([]string, error) {
	accounts := []string{}

	err := rpc.call("eth_accounts", &accounts)
	return accounts, err
}

// EthBlockNumber returns the number of most recent block.
func (rpc EthRPC) EthBlockNumber() (int, error) {
	var response string
	if err := rpc.call("eth_blockNumber", &response); err != nil {
		return 0, err
	}

	return ParseInt(response)
}

// EthGetBalance returns the balance of the account of given address in wei.
func (rpc EthRPC) EthGetBalance(address, block string) (*big.Int, error) {
	var response string
	if err := rpc.call("eth_getBalance", &response, address, block); err != nil {
		return new(big.Int), err
	}

	return ParseBigInt(response)
}

// EthGetStorageAt returns the value from a storage position at a given address.
func (rpc EthRPC) EthGetStorageAt(data string, position int, tag string) (string, error) {
	var result string

	err := rpc.call("eth_getStorageAt", &result, data, IntToHex(position), tag)
	return result, err
}

// EthGetTransactionCount returns the number of transactions sent from an address.
func (rpc EthRPC) EthGetTransactionCount(address, block string) (int, error) {
	var response string

	if err := rpc.call("eth_getTransactionCount", &response, address, block); err != nil {
		return 0, err
	}

	return ParseInt(response)
}

// EthGetBlockTransactionCountByHash returns the number of transactions in a block from a block matching the given block hash.
func (rpc EthRPC) EthGetBlockTransactionCountByHash(hash string) (int, error) {
	var response string

	if err := rpc.call("eth_getBlockTransactionCountByHash", &response, hash); err != nil {
		return 0, err
	}

	return ParseInt(response)
}

// EthGetBlockTransactionCountByNumber returns the number of transactions in a block from a block matching the given block
func (rpc EthRPC) EthGetBlockTransactionCountByNumber(number int) (int, error) {
	var response string

	if err := rpc.call("eth_getBlockTransactionCountByNumber", &response, IntToHex(number)); err != nil {
		return 0, err
	}

	return ParseInt(response)
}

// EthGetUncleCountByBlockHash returns the number of uncles in a block from a block matching the given block hash.
func (rpc EthRPC) EthGetUncleCountByBlockHash(hash string) (int, error) {
	var response string

	if err := rpc.call("eth_getUncleCountByBlockHash", &response, hash); err != nil {
		return 0, err
	}

	return ParseInt(response)
}

// EthGetUncleCountByBlockNumber returns the number of uncles in a block from a block matching the given block number.
func (rpc EthRPC) EthGetUncleCountByBlockNumber(number int) (int, error) {
	var response string

	if err := rpc.call("eth_getUncleCountByBlockNumber", &response, IntToHex(number)); err != nil {
		return 0, err
	}

	return ParseInt(response)
}

// EthGetCode returns code at a given address.
func (rpc EthRPC) EthGetCode(address, block string) (string, error) {
	var code string

	err := rpc.call("eth_getCode", &code, address, block)
	return code, err
}

// EthSign signs data with a given address.
// Calculates an Ethereum specific signature with: sign(keccak256("\x19Ethereum Signed Message:\n" + len(message) + message)))
func (rpc EthRPC) EthSign(address, data string) (string, error) {
	var signature string

	err := rpc.call("eth_sign", &signature, address, data)
	return signature, err
}

// EthSendTransaction creates new message call transaction or a contract creation, if the data field contains code.
func (rpc EthRPC) EthSendTransaction(transaction T) (string, error) {
	var hash string

	err := rpc.call("eth_sendTransaction", &hash, transaction)
	return hash, err
}

// EthSendRawTransaction creates new message call transaction or a contract creation for signed transactions.
func (rpc EthRPC) EthSendRawTransaction(data string) (string, error) {
	var hash string

	err := rpc.call("eth_sendRawTransaction", &hash, data)
	return hash, err
}

// EthCall executes a new message call immediately without creating a transaction on the block chain.
func (rpc EthRPC) EthCall(transaction T, tag string) (string, error) {
	var data string

	err := rpc.call("eth_call", &data, transaction, tag)
	return data, err
}

// EthEstimateGas makes a call or transaction, which won't be added to the blockchain and returns the used gas, which can be used for estimating the used gas.
func (rpc EthRPC) EthEstimateGas(transaction T) (int, error) {
	var response string

	err := rpc.call("eth_estimateGas", &response, transaction)
	if err != nil {
		return 0, err
	}

	return ParseInt(response)
}

func (rpc EthRPC) getBlock(method string, withTransactions bool, params ...interface{}) (*Block, error) {
	result, err := rpc.RawCall(method, params...)
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
func (rpc EthRPC) EthGetBlockByHash(hash string, withTransactions bool) (*Block, error) {
	return rpc.getBlock("eth_getBlockByHash", withTransactions, hash, withTransactions)
}

// EthGetBlockByNumber returns information about a block by block number.
func (rpc EthRPC) EthGetBlockByNumber(number int, withTransactions bool) (*Block, error) {
	return rpc.getBlock("eth_getBlockByNumber", withTransactions, IntToHex(number), withTransactions)
}

func (rpc EthRPC) getTransaction(method string, params ...interface{}) (*Transaction, error) {
	transaction := new(Transaction)

	err := rpc.call(method, transaction, params...)
	return transaction, err
}

// EthGetTransactionByHash returns the information about a transaction requested by transaction hash.
func (rpc EthRPC) EthGetTransactionByHash(hash string) (*Transaction, error) {
	return rpc.getTransaction("eth_getTransactionByHash", hash)
}

// EthGetTransactionByBlockHashAndIndex returns information about a transaction by block hash and transaction index position.
func (rpc EthRPC) EthGetTransactionByBlockHashAndIndex(blockHash string, transactionIndex int) (*Transaction, error) {
	return rpc.getTransaction("eth_getTransactionByBlockHashAndIndex", blockHash, IntToHex(transactionIndex))
}

// EthGetTransactionByBlockNumberAndIndex returns information about a transaction by block number and transaction index position.
func (rpc EthRPC) EthGetTransactionByBlockNumberAndIndex(blockNumber, transactionIndex int) (*Transaction, error) {
	return rpc.getTransaction("eth_getTransactionByBlockNumberAndIndex", IntToHex(blockNumber), IntToHex(transactionIndex))
}

// EthGetTransactionReceipt returns the receipt of a transaction by transaction hash.
// Note That the receipt is not available for pending transactions.
func (rpc EthRPC) EthGetTransactionReceipt(hash string) (*TransactionReceipt, error) {
	transactionReceipt := new(TransactionReceipt)

	err := rpc.call("eth_getTransactionReceipt", transactionReceipt, hash)
	if err != nil {
		return nil, err
	}

	return transactionReceipt, nil
}

// EthGetCompilers returns a list of available compilers in the client.
func (rpc EthRPC) EthGetCompilers() ([]string, error) {
	compilers := []string{}

	err := rpc.call("eth_getCompilers", &compilers)
	return compilers, err
}

// EthNewFilter creates a new filter object.
func (rpc EthRPC) EthNewFilter(params FilterParams) (string, error) {
	var filterID string
	err := rpc.call("eth_newFilter", &filterID, params)
	return filterID, err
}

// EthNewBlockFilter creates a filter in the node, to notify when a new block arrives.
// To check if the state has changed, call EthGetFilterChanges.
func (rpc EthRPC) EthNewBlockFilter() (string, error) {
	var filterID string
	err := rpc.call("eth_newBlockFilter", &filterID)
	return filterID, err
}

// EthNewPendingTransactionFilter creates a filter in the node, to notify when new pending transactions arrive.
// To check if the state has changed, call EthGetFilterChanges.
func (rpc EthRPC) EthNewPendingTransactionFilter() (string, error) {
	var filterID string
	err := rpc.call("eth_newPendingTransactionFilter", &filterID)
	return filterID, err
}

// EthUninstallFilter uninstalls a filter with given id.
func (rpc EthRPC) EthUninstallFilter(filterID string) (bool, error) {
	var res bool
	err := rpc.call("eth_uninstallFilter", &res, filterID)
	return res, err
}

// EthGetFilterChanges polling method for a filter, which returns an array of logs which occurred since last poll.
func (rpc EthRPC) EthGetFilterChanges(filterID string) ([]Log, error) {
	var logs = []Log{}
	err := rpc.call("eth_getFilterChanges", &logs, filterID)
	return logs, err
}

// EthGetFilterLogs returns an array of all logs matching filter with given id.
func (rpc EthRPC) EthGetFilterLogs(filterID string) ([]Log, error) {
	var logs = []Log{}
	err := rpc.call("eth_getFilterLogs", &logs, filterID)
	return logs, err
}

// EthGetLogs returns an array of all logs matching a given filter object.
func (rpc EthRPC) EthGetLogs(params FilterParams) ([]Log, error) {
	var logs = []Log{}
	err := rpc.call("eth_getLogs", &logs, params)
	return logs, err
}

// Eth1 returns 1 ethereum value (10^18 wei)
func (rpc EthRPC) Eth1() *big.Int {
	return Eth1()
}

func (rpc EthRPC) Erc20TotalSupply(contract string) (json.RawMessage, error) {
	return rpc.Call("eth_call", map[string]interface{}{
		"to":   contract,
		"data": paramencoder.TotalSupplyParams,
		/*"gas":      "0xaae60", //700000,
		"gasPrice": "0x15f90", //90000,*/
	}, "latest")
}

func (rpc EthRPC) Erc20Decimals(contract string) (json.RawMessage, error) {
	return rpc.Call("eth_call", map[string]interface{}{
		"to":   contract,
		"data": paramencoder.DecimalsParams,
		/*"gas":      "0xaae60", //700000,
		"gasPrice": "0x15f90", //90000,*/
	}, "latest")
}

func (rpc EthRPC) Erc20BalanceOf(contract string, tokenOwner string) (json.RawMessage, error) {
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
	return rpc.Call("eth_call", map[string]interface{}{
		"to":   contract,
		"data": dataContent,
		/*"gas":      "0xaae60", //700000,
		"gasPrice": "0x15f90", //90000,*/
	}, "latest")
}

func (rpc EthRPC) Erc20Allowance(contract string, tokenOwner string, spender string) (json.RawMessage, error) {
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
	return rpc.Call("eth_call", map[string]interface{}{
		"to":   contract,
		"data": dataContent,
		/*"gas":      "0xaae60", //700000,
		"gasPrice": "0x15f90", //90000,*/
	}, "latest")
}

// curl localhost:8545 -X POST --data '{"jsonrpc":"2.0","method":"eth_sendTransaction","params":[{"from": "0x8aff0a12f3e8d55cc718d36f84e002c335df2f4a", "data": "606060405260728060106000396000f360606040526000357c0100000000000000000000000000000000000000000000000000000000900480636ffa1caa146037576035565b005b604b60048080359060200190919050506061565b6040518082815260200191505060405180910390f35b6000816002029050606d565b91905056"}],"id":1}
func (rpc EthRPC) DeployContract(fromAddress string, bytecode string) (json.RawMessage, error) {
	return rpc.Call("eth_sendTransaction", map[string]string{
		"from": fromAddress,
		"data": bytecode,
	})
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
func Eth1() *big.Int {
	return oneEth
}
