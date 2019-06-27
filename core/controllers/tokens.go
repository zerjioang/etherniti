// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package controllers

import (
	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/controllers/tokenlist"
	"github.com/zerjioang/etherniti/core/data"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/shared/constants"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

type TokenController struct {
}

// contructor like function
func NewTokenController() TokenController {
	dc := TokenController{}
	return dc
}

func (ctl TokenController) whoisAddress(c *echo.Context) error {
	return api.ErrorBytes(c, data.NotImplemented)
}

func (ctl TokenController) resolveContractAddress(c *echo.Context) error {
	//set cache policy
	c.OnSuccessCachePolicy = constants.CacheOneDay

	symbol := c.Param("symbol")
	address := tokenlist.GetTokenAddressByName(symbol)
	return api.SendSuccess(c, []byte("resolved contract address"), address)
}

func (ctl TokenController) resolveContractSymbol(c *echo.Context) error {
	//set cache policy
	c.OnSuccessCachePolicy = constants.CacheOneDay

	address := c.Param("address")
	symbol := tokenlist.GetTokenSymbol(address)
	return api.SendSuccess(c, []byte("resolved contract address"), symbol)
}

func (ctl TokenController) RegisterRouters(router *echo.Group) {
	logger.Debug("exposing token controller methods")
	router.GET("/token/whois/:address", ctl.whoisAddress)
	router.GET("/token/resolve/address/:symbol", ctl.resolveContractAddress)
	router.GET("/token/resolve/symbol/:address", ctl.resolveContractSymbol)
}
