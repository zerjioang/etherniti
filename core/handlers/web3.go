// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/server"
	"github.com/zerjioang/etherniti/core/util"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/zerjioang/etherniti/core/eth"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

var (
	ctx = context.Background()
)

// eth web3 controller
type Web3Controller struct {
	// in memory based wallet manager
	walletManager eth.WalletManager
	//ethereum interaction cache
	cache *cache.Cache
}

// constructor like function
func NewWeb3Controller(manager eth.WalletManager) Web3Controller {
	ctl := Web3Controller{}
	ctl.walletManager = manager
	ctl.cache = cache.New(5*time.Minute, 10*time.Minute)
	return ctl
}

// check if an ethereum address is a contract address
func (ctl Web3Controller) getBalance(c echo.Context) error {
	// cast to our context
	cc, ok := c.(*server.EthernitiContext)
	if !ok {
		return ErrorStr(c, "failed to execute requested operation")
	}

	clientInstance, err := GetClientInstance(cc)
	if err != nil {
		// there was an trycatch recovering client instance
		apiErr := api.NewApiError(http.StatusBadRequest, err.Error())
		apiErrRaw := util.GetJsonBytes(apiErr)
		return c.JSONBlob(http.StatusBadRequest, apiErrRaw)
	}
	targetAddr := c.Param("address")
	// check if not empty
	if targetAddr != "" {
		result, err := clientInstance.EthGetBalance(targetAddr, "latest")
		if err != nil {
			//some trycatch happen, return trycatch to client
			apiErr := api.NewApiError(http.StatusBadRequest, err.Error())
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
		return ErrorStr(c, "failed to execute requested operation")
	}

	clientInstance, err := GetClientInstance(cc)
	if err != nil {
		// there was an trycatch recovering client instance
		apiErr := api.NewApiError(http.StatusBadRequest, err.Error())
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
			apiErr := api.NewApiError(http.StatusBadRequest, err.Error())
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
		return ErrorStr(c, "failed to execute requested operation")
	}

	clientInstance, err := GetClientInstance(cc)
	if err != nil {
		// there was an trycatch recovering client instance
		apiErr := api.NewApiError(http.StatusBadRequest, err.Error())
		apiErrRaw := util.GetJsonBytes(apiErr)
		return c.JSONBlob(http.StatusBadRequest, apiErrRaw)
	}
	data, err := clientInstance.EthNodeInfo()
	if err != nil {
		// send invalid address message
		return c.JSONBlob(http.StatusBadRequest, invalidAddressBytes)
	} else {
		return Success(c, "eth_info", data)
	}
}

func (ctl Web3Controller) getAccounts(c echo.Context) error {
	// cast to our context
	cc, ok := c.(*server.EthernitiContext)
	if !ok {
		return ErrorStr(c, "failed to execute requested operation")
	}

	client := GetClient(cc)
	list, err := client.EthAccounts()
	if err != nil {
		// send invalid generation message
		return c.JSONBlob(http.StatusInternalServerError,
			util.GetJsonBytes(
				api.NewApiError(http.StatusInternalServerError, err.Error()),
			),
		)
	} else {
		return c.JSONBlob(
			http.StatusOK,
			util.GetJsonBytes(
				api.NewApiResponse("ethereum accounts readed", list),
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
		return ErrorStr(c, "failed to execute requested operation")
	}

	client := GetClient(cc)
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
				api.NewApiError(http.StatusInternalServerError, err.Error()),
			),
		)
	} else {
		//iterate over account
		for i := 0; i < len(list); i++ {
			currentAccount := list[i]
			bigInt, err := client.EthGetBalance(currentAccount, "latest")
			if err != nil {
				log.Error("failed to get account balance", currentAccount, err)
			} else {
				item := &wrapperList[i]
				item.Account = currentAccount
				item.Balance = bigInt.String()
				item.Eth = eth.ToEth(bigInt).String()
				item.Key = "secret"
			}
		}
		return c.JSONBlob(
			http.StatusOK,
			util.GetJsonBytes(
				api.NewApiResponse("ethereum accounts and their balance readed", wrapperList),
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
		return ErrorStr(c, "failed to execute requested operation")
	}

	raw, found := ctl.cache.Get("eth_coinbase")
	if found && raw != nil {
		//cache hit
		//return result to client
		return c.JSONBlob(
			http.StatusOK,
			util.GetJsonBytes(
				api.NewApiResponse("coinbase address", raw),
			),
		)
	} else {
		//cache miss
		client := GetClient(cc)
		result, err := client.Call("eth_coinbase")
		if err == nil {
			if result != nil {
				// add result to cache
				ctl.cache.Set("eth_coinbase", result, cache.DefaultExpiration)
				//return result to client
				return c.JSONBlob(
					http.StatusOK,
					util.GetJsonBytes(
						api.NewApiResponse("coinbase address", result),
					),
				)
			} else {
				return c.JSONBlob(http.StatusBadRequest,
					util.GetJsonBytes(
						api.NewApiError(http.StatusBadRequest, "empty response from server"),
					),
				)
			}
		} else {
			return c.JSONBlob(http.StatusBadRequest,
				util.GetJsonBytes(
					api.NewApiError(http.StatusBadRequest, "failed to get coinbase address: "+err.Error()),
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
		return ErrorStr(c, "failed to execute requested operation")
	}
	clientInstance, err := GetClientInstance(cc)
	if err != nil {
		// there was an trycatch recovering client instance
		apiErr := api.NewApiError(http.StatusBadRequest, err.Error())
		apiErrRaw := util.GetJsonBytes(apiErr)
		return c.JSONBlob(http.StatusBadRequest, apiErrRaw)
	}
	targetAddr := c.Param("address")
	// check if not empty
	if targetAddr != "" {
		result, err := eth.IsSmartContractAddress(clientInstance, targetAddr)
		if err != nil {
			//some trycatch happen, return trycatch to client
			apiErr := api.NewApiError(http.StatusBadRequest, err.Error())
			apiErrRaw := util.GetJsonBytes(apiErr)
			return c.JSONBlob(http.StatusBadRequest, apiErrRaw)
		}
		return c.JSONBlob(http.StatusOK, util.GetJsonBytes(result))
	}
	// send invalid address message
	return c.JSONBlob(http.StatusBadRequest, invalidAddressBytes)
}

// https://medium.com/@akshay_111meher/creating-offline-raw-transactions-with-go-ethereum-8d6cc8174c5d
// sends new eth transaction using given configuration
func (ctl Web3Controller) DeployContract(to common.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte, compiledBytecode string, contractAbiStr string) (string, error) {
	// Construct the transaction
	d := time.Now().Add(1000 * time.Millisecond)
	cancellableCtx, cancel := context.WithDeadline(ctx, d)

	// connect the client
	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		log.Error(err)
		cancel()
		return "", err
	}

	// remove 0x
	bytecode := common.Hex2Bytes(compiledBytecode[2:])
	contractAbi, err := abi.JSON(strings.NewReader(contractAbiStr))
	input, err := contractAbi.Pack("")

	bytecode = append(bytecode, input...)

	// To create a raw transaction nonce is must.
	// Hence we get the account nonce using
	// The third parameter is nil because we want the latest account nonce
	// To get nonce at a particular block height
	nonce, err := client.NonceAt(
		cancellableCtx,
		common.HexToAddress("0x56724a9e4d2bb2dca01999acade2e88a92b11a9e"), nil)

	/*
			The arguments to this function in order are
		    1. nonce
		    2. to-address (we would need to convert this using common.HexToAddress(public address) )
		    3. balance to be sent (use big number)
		    4. gas limit
		    5. gas price
		    6. data (since this is not a contract transaction, we can pass nil )
	*/
	tx := types.NewContractCreation(nonce, big.NewInt(0), gasLimit, gasPrice, bytecode)
	// Define signer and chain id
	// chainID := big.NewInt(CHAIN_ID)
	// signer := types.NewEIP155Signer(chainID)
	// FrontierSigner
	// EIP155Signer
	signer := types.HomesteadSigner{}

	//create example account
	privateKey, err := eth.GenerateNewKey()
	// Sign the transaction signature with the private key
	signedTx, signatureErr := types.SignTx(tx, signer, privateKey)
	if signatureErr != nil {
		fmt.Println("signature create trycatch:")
		log.Error(signatureErr)
		cancel()
		return "", signatureErr
	}
	// Send the signed transaction to the network
	txErr := client.SendTransaction(cancellableCtx, signedTx)
	if txErr != nil {
		fmt.Println("send tx trycatch:")
		log.Error(txErr)
		cancel()
		return "", txErr
	} else {
		select {
		case <-time.After(1 - time.Millisecond):
			log.Info("tx send overslept")
		case <-cancellableCtx.Done():
			log.Info(cancellableCtx.Err())
		default:
			fmt.Printf("send success tx.hash=%s\n", signedTx.Hash().String())
			cancel()
			return signedTx.Hash().String(), nil
		}
	}
	cancel()
	return "", nil
}

// sends new eth transaction using given configuration
func (ctl Web3Controller) SendTransaction(to common.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte) (string, error) {
	// Construct the transaction
	d := time.Now().Add(1000 * time.Millisecond)
	cancellableCtx, cancel := context.WithDeadline(ctx, d)

	// connect the client
	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		log.Error(err)
		cancel()
		return "", err
	}

	// To create a raw transaction nonce is must.
	// Hence we get the account nonce using
	// The third parameter is nil because we want the latest account nonce
	// To get nonce at a particular block height
	nonce, err := client.NonceAt(
		cancellableCtx,
		common.HexToAddress("0x56724a9e4d2bb2dca01999acade2e88a92b11a9e"), nil)

	/*
			The arguments to this function in order are
		    1. nonce
		    2. to-address (we would need to convert this using common.HexToAddress(public address) )
		    3. balance to be sent (use big number)
		    4. gas limit
		    5. gas price
		    6. data (since this is not a contract transaction, we can pass nil )
	*/
	tx := types.NewTransaction(nonce, to, amount, gasLimit, gasPrice, data)
	// Define signer and chain id
	// chainID := big.NewInt(CHAIN_ID)
	// signer := types.NewEIP155Signer(chainID)
	// FrontierSigner
	// EIP155Signer
	signer := types.HomesteadSigner{}

	//create example account
	privateKey, err := eth.GenerateNewKey()
	// Sign the transaction signature with the private key
	signedTx, signatureErr := types.SignTx(tx, signer, privateKey)
	if signatureErr != nil {
		fmt.Println("signature create trycatch:")
		log.Error(signatureErr)
		cancel()
		return "", signatureErr
	}
	// Send the signed transaction to the network
	txErr := client.SendTransaction(cancellableCtx, signedTx)
	if txErr != nil {
		fmt.Println("send tx trycatch:")
		log.Error(txErr)
		cancel()
		return "", txErr
	} else {
		select {
		case <-time.After(1 - time.Millisecond):
			log.Info("tx send overslept")
		case <-cancellableCtx.Done():
			log.Info(cancellableCtx.Err())
		default:
			fmt.Printf("send success tx.hash=%s\n", signedTx.Hash().String())
			cancel()
			return signedTx.Hash().String(), nil
		}
	}
	cancel()
	return "", nil
}

// call deployed smart contract method
func (ctl Web3Controller) CallContract() {
	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		log.Error(err)
	}

	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Error(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Error("trycatch casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		log.Error(err)
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Error(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

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
