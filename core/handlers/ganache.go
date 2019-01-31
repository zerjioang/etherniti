// Copyright gaethway
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/zerjioang/gaethway/core/eth"
)

// ganache controller
type GanacheController struct {
	// in memory based wallet manager
	walletManager eth.WalletManager
}

// constructor like function
func NewGanacheController(manager eth.WalletManager) GanacheController {
	ctl := GanacheController{}
	ctl.walletManager = manager
	return ctl
}

func (ctl GanacheController) getAccounts(c echo.Context) error {
	return nil
}

func (ctl GanacheController) getBlocks(c echo.Context) error {
	return nil
}

func (ctl GanacheController) getTransactions(c echo.Context) error {
	return nil
}

// implemented method from interface RouterRegistrable
func (ctl GanacheController) RegisterRouters(router *echo.Echo) {
	log.Info("exposing ganache controller methods")
	//http://localhost:8080/eth/create
	router.GET("/v1/ganache/accounts", ctl.getAccounts)
	router.GET("/v1/ganache/blocks", ctl.getBlocks)
	router.GET("/v1/ganache/tx", ctl.getTransactions)
}
