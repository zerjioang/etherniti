// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"errors"
	"math/big"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/eth/rpc"

	"github.com/labstack/gommon/log"
	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/eth"
	"github.com/zerjioang/etherniti/core/modules/ethfork/ethclient"
	"github.com/zerjioang/etherniti/core/util"

	"github.com/labstack/echo"
)

const (
	invalidAddress   = `{"message": "please, provide a valid ethereum or quorum address"}`
	accountKeyGenErr = `{"message": "failed to generate ecdsa private key"}`
)

var (
	noConnErrMsg           = "Invalid connection profile key provided in the request header. Please, make sure you have created a connection profile indicating your peer node IP address or domain name."
	errNoConnectionProfile = errors.New(noConnErrMsg)
	accountKeyGenErrBytes  = util.Bytes(accountKeyGenErr)
	invalidAddressBytes    = util.Bytes(invalidAddress)
)

type EthController struct {
	// in memory based wallet manager
	walletManager eth.WalletManager
	//ethereum interaction cache
	cache *cache.Cache
}

func NewEthController(manager eth.WalletManager) EthController {
	ctl := EthController{}
	ctl.walletManager = manager
	ctl.cache = cache.New(5*time.Minute, 10*time.Minute)
	return ctl
}

// generates an ethereum new account (address+key)
func (ctl EthController) generateAddress(c echo.Context) error {

	// Create an account
	private, err := eth.GenerateNewKey()

	if err != nil {
		log.Error("failed to generate ethereum account key", err)
		// send invalid generation message
		return c.JSONBlob(http.StatusInternalServerError, accountKeyGenErrBytes)
	}
	address := eth.GetAddressFromPrivateKey(private)
	privateKey := eth.GetPrivateKeyAsEthString(private)
	var response = map[string]string{
		"address": address.Hex(),
		"private": privateKey,
	}
	return c.JSONBlob(
		http.StatusOK,
		util.GetJsonBytes(
			api.NewApiResponse("ethereum account created", response),
		),
	)
}

// check if an ethereum address is valid
func (ctl EthController) isValidAddress(c echo.Context) error {
	//read user entered address
	targetAddr := c.Param("address")
	// check if not empty
	if targetAddr != "" {
		result := eth.IsValidAddress(targetAddr)
		return c.JSONBlob(http.StatusOK, util.GetJsonBytes(
			api.NewApiResponse("address validation checked", result),
		),
		)
	}
	// send invalid address message
	return c.JSONBlob(http.StatusBadRequest, invalidAddressBytes)
}

// check if an ethereum address is a contract address
func (ctl EthController) isContractAddress(c echo.Context) error {
	clientInstance, err := ctl.getClientInstance(c)
	if err != nil || clientInstance == nil {
		// there was an error recovering client instance
		apiErr := api.NewApiError(http.StatusBadRequest, err.Error())
		apiErrRaw := util.GetJsonBytes(apiErr)
		return c.JSONBlob(http.StatusBadRequest, apiErrRaw)
	}
	targetAddr := c.Param("address")
	// check if not empty
	if targetAddr != "" {
		result, err := eth.IsSmartContractAddress(clientInstance, targetAddr)
		if err != nil {
			//some error happen, return error to client
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
func (ctl EthController) getBalance(c echo.Context) error {
	clientInstance, err := ctl.getClientInstance(c)
	if err != nil || clientInstance == nil {
		// there was an error recovering client instance
		apiErr := api.NewApiError(http.StatusBadRequest, err.Error())
		apiErrRaw := util.GetJsonBytes(apiErr)
		return c.JSONBlob(http.StatusBadRequest, apiErrRaw)
	}
	targetAddr := c.Param("address")
	// check if not empty
	if targetAddr != "" {
		ethAddr := eth.ConvertAddress(targetAddr)
		result, err := eth.GetAccountBalance(clientInstance, ethAddr)
		if err != nil {
			//some error happen, return error to client
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
func (ctl EthController) getBalanceAtBlock(c echo.Context) error {
	clientInstance, err := ctl.getClientInstance(c)
	if err != nil || clientInstance == nil {
		// there was an error recovering client instance
		apiErr := api.NewApiError(http.StatusBadRequest, err.Error())
		apiErrRaw := util.GetJsonBytes(apiErr)
		return c.JSONBlob(http.StatusBadRequest, apiErrRaw)
	}
	targetAddr := c.Param("address")
	block := c.Param("block")
	// check if not empty
	if targetAddr != "" {
		ethAddr := eth.ConvertAddress(targetAddr)
		b := new(big.Int)
		b.SetString(block, 10)
		result, err := eth.GetAccountBalanceAtBlock(clientInstance, ethAddr, b)
		if err != nil {
			//some error happen, return error to client
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
func (ctl EthController) getNodeIndo(c echo.Context) error {
	clientInstance, err := ctl.getClientInstance(c)
	if err != nil || clientInstance == nil {
		// there was an error recovering client instance
		apiErr := api.NewApiError(http.StatusBadRequest, err.Error())
		apiErrRaw := util.GetJsonBytes(apiErr)
		return c.JSONBlob(http.StatusBadRequest, apiErrRaw)
	}
	// send invalid address message
	return c.JSONBlob(http.StatusBadRequest, invalidAddressBytes)
}

func (ctl EthController) getAccounts(c echo.Context) error {
	client := ctl.getClient(c)
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
func (ctl EthController) getAccountsWithBalance(c echo.Context) error {
	client := ctl.getClient(c)
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

func (ctl EthController) getBlocks(c echo.Context) error {
	return nil
}

func (ctl EthController) coinbase(c echo.Context) error {

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
		client := ctl.getClient(c)
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

// from incoming http request, it recovers the eth client linked to it
func (ctl EthController) getClientInstance(c echo.Context) (*ethclient.Client, error) {
	requestProfileKey := c.Request().Header.Get(config.HttpProfileHeaderkey)
	wallet, found := ctl.walletManager.Get(requestProfileKey)
	if !found {
		return nil, errNoConnectionProfile
	}
	return wallet.Client(), nil
}

func (ctl EthController) getClient(context echo.Context) ethrpc.EthRPC {
	client := ethrpc.NewDefaultRPC("http://127.0.0.1:8545")
	return client
}

// implemented method from interface RouterRegistrable
func (ctl EthController) RegisterRouters(router *echo.Echo) {
	log.Info("exposing eth controller methods")
	router.GET("/v1/eth/create", ctl.generateAddress)
	router.GET("/v1/eth/verify/:address", ctl.isValidAddress)
	router.GET("/v1/eth/hascontract/:address", ctl.isContractAddress)
	router.GET("/v1/eth/m/accounts", ctl.getAccounts)
	router.GET("/v1/eth/m/accountsBalanced", ctl.getAccountsWithBalance)
	router.GET("/v1/eth/m/blocks", ctl.getBlocks)
	router.GET("/v1/eth/m/coinbase", ctl.coinbase)

	router.GET("/v1/eth/m/getbalance/:address", ctl.getBalance)
	router.GET("/v1/eth/m/getbalance/:address/block/:block", ctl.getBalanceAtBlock)
}
