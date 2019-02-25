// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"errors"
	"net/http"

	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/eth"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/util"

	"github.com/labstack/echo"
)

const (
	invalidAddress   = `{"message": "please, provide a valid ethereum or quorum address"}`
	accountKeyGenErr = `{"message": "failed to generate ecdsa private key"}`
)

var (
	noConnErrMsg           = "invalid connection profile key provided in the request header. Please, make sure you have created a connection profile indicating your peer node IP address or domain name."
	errNoConnectionProfile = errors.New(noConnErrMsg)
	accountKeyGenErrBytes  = util.Bytes(accountKeyGenErr)
	invalidAddressBytes    = util.Bytes(invalidAddress)
)

type EthController struct {
}

func NewEthController() EthController {
	ctl := EthController{}
	return ctl
}

// generates an ethereum new account (address+key)
func (ctl EthController) generateAddress(c echo.Context) error {

	// Create an account
	private, err := eth.GenerateNewKey()

	if err != nil {
		logger.Error("failed to generate ethereum account key", err)
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

// implemented method from interface RouterRegistrable
func (ctl EthController) RegisterRouters(router *echo.Group) {
	logger.Info("exposing eth controller methods")
	router.GET("/eth/create", ctl.generateAddress)
	router.GET("/eth/verify/:address", ctl.isValidAddress)

}
