// Copyright gaethway
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/gommon/log"
	"github.com/zerjioang/gaethway/core/api"
	"github.com/zerjioang/gaethway/core/eth"
	"github.com/zerjioang/gaethway/core/keystore/memory"
	"github.com/zerjioang/gaethway/core/util"

	"github.com/labstack/echo"
	"github.com/patrickmn/go-cache"
)

const (
	invalidAddress = `{"message": "please, provide a valid ethereum or quorum address"}`
)

type EthController struct {
	// in memory storage of created wallets
	wallet *memory.InMemoryKeyStorage
	cache  *cache.Cache
}

func NewEthController() EthController {
	ctl := EthController{}
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	ctl.cache = cache.New(5*time.Minute, 10*time.Minute)
	// create an in memory wallet
	ctl.wallet = memory.NewInMemoryKeyStorage()
	return ctl
}

// check if an ethereum address is valid
func (ctl EthController) isValidAddress(c echo.Context) error {
	//read user entered address
	targetAddr := c.Param("address")
	// check if not empty
	if targetAddr != "" {
		result := eth.IsValidAddress(targetAddr)
		return c.JSONBlob(http.StatusOK, util.GetJsonBytes(result))
	}
	// send invalid address message
	return c.JSONBlob(http.StatusBadRequest, util.Bytes(invalidAddress))
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
	return c.JSONBlob(http.StatusBadRequest, util.Bytes(invalidAddress))
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
	return c.JSONBlob(http.StatusBadRequest, util.Bytes(invalidAddress))
}

// from incoming http request, it recovers the eth client linked to it
func (ctl EthController) getClientInstance(c echo.Context) (*ethclient.Client, error) {
	requestProfileKey := c.Request().Header.Get(api.HttpProfileHeaderkey)
	wallet, found := ctl.wallet.Get(requestProfileKey)
	if !found {
		return nil, errors.New("invalid profile key provided in the request header")
	}
	return wallet.Client(), nil
}

// implemented method from interface RouterRegistrable
func (ctl EthController) RegisterRouters(router *echo.Echo) {
	log.Info("exposing eth controller methods")
	//http://localhost:8080/eth/verify/0x71c7656ec7ab88b098defb751b7401b5f6d8976f
	router.GET("/eth/verify/:address", ctl.isValidAddress)
	//http://localhost:8080/eth/hascontract/0x71c7656ec7ab88b098defb751b7401b5f6d8976f
	router.GET("/eth/hascontract/:address", ctl.isContractAddress)
	//http://localhost:8080/eth/getbalance/0x71c7656ec7ab88b098defb751b7401b5f6d8976f
	router.GET("/eth/getbalance/:address", ctl.getBalance)
}
