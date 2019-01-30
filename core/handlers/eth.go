// Copyright gaethway
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"errors"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/gommon/log"
	"github.com/zerjioang/gaethway/core/api"
	"github.com/zerjioang/gaethway/core/eth"
	"github.com/zerjioang/gaethway/core/util"

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
}

func NewEthController(manager eth.WalletManager) EthController {
	ctl := EthController{}
	ctl.walletManager = manager
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

// from incoming http request, it recovers the eth client linked to it
func (ctl EthController) getClientInstance(c echo.Context) (*ethclient.Client, error) {
	requestProfileKey := c.Request().Header.Get(api.HttpProfileHeaderkey)
	wallet, found := ctl.walletManager.Get(requestProfileKey)
	if !found {
		return nil, errNoConnectionProfile
	}
	return wallet.Client(), nil
}

// implemented method from interface RouterRegistrable
func (ctl EthController) RegisterRouters(router *echo.Echo) {
	log.Info("exposing eth controller methods")
	//http://localhost:8080/eth/create
	router.GET("/v1/eth/create", ctl.generateAddress)
	//http://localhost:8080/eth/verify/0x71c7656ec7ab88b098defb751b7401b5f6d8976f
	router.GET("/v1/eth/verify/:address", ctl.isValidAddress)
	//http://localhost:8080/eth/hascontract/0x71c7656ec7ab88b098defb751b7401b5f6d8976f
	router.GET("/v1/eth/hascontract/:address", ctl.isContractAddress)
	//http://localhost:8080/eth/getbalance/0x71c7656ec7ab88b098defb751b7401b5f6d8976f
	router.GET("/v1/eth/getbalance/:address", ctl.getBalance)
	router.GET("/v1/eth/getbalance/:address/block/:block", ctl.getBalanceAtBlock)
}
