// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"errors"
	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/modules/cns"

	"github.com/labstack/echo"
	"github.com/zerjioang/etherniti/core/logger"
)

// token controller
type ContractNameSpaceController struct {
	ns cns.ContractNameSystem
}

// constructor like function
func NewContractNameSpaceController() ContractNameSpaceController {
	ctl := ContractNameSpaceController{}
	ctl.ns = cns.NewContractNameSystem()
	return ctl
}

func (ctl *ContractNameSpaceController) register(c echo.Context) error {
	return errors.New("not implemented")
}

func (ctl *ContractNameSpaceController) unregister(c echo.Context) error {
	id := c.Param("id")
	ctl.ns.Unregister(id)
	return api.Success(c, "contract successfully unregistered from namespace", id)
}

func (ctl ContractNameSpaceController) resolve(c echo.Context) error {
	id := c.Param("id")
	data, found := ctl.ns.Resolve(id)
	if !found {
		return api.ErrorStr(c, "failed to resolve given contract id")
	} else {
		return api.SendSuccess(c, "contract information successfully resolved", data)
	}
}

// implemented method from interface RouterRegistrable
func (ctl ContractNameSpaceController) RegisterRouters(router *echo.Group) {
	logger.Info("exposing solc controller methods")
	router.POST("/cns", ctl.register)
	router.DELETE("/cns/:id", ctl.unregister)
	router.GET("/cns/:id", ctl.resolve)
}
