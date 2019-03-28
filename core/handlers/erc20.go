// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"errors"

	"github.com/labstack/echo"
	"github.com/zerjioang/etherniti/core/api"
)

// eth web3 controller
type Erc20Controller struct {
	NetworkController
}

// constructor like function
func NewErc20Controller() Erc20Controller {
	ctl := Erc20Controller{}
	ctl.NetworkController = NewNetworkController()
	return ctl
}

// implemented method from interface RouterRegistrable
func (ctl Erc20Controller) totalSupply(c echo.Context) error {
	contractAddress := c.Param("contract")
	if contractAddress == "" {
		return api.ErrorStr(c, "invalid contract address provided")
	}
	return errors.New("not implemented")
}

// implemented method from interface RouterRegistrable
func (ctl Erc20Controller) RegisterRouters(router *echo.Group) {
	router.GET("/erc20/:contract/totalSupply", ctl.totalSupply)
}
