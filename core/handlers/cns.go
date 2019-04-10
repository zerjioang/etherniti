// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/modules/cns"
	"github.com/zerjioang/etherniti/thirdparty/echo"
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
	//new profile request
	req := cns.ContractInfo{}
	if err := c.Bind(&req); err != nil {
		// return a binding error
		logger.Error("failed to bind request data to model: ", err)
		return api.ErrorStr(c, bindErr)
	}
	e := req.Validate()
	if e.Occur() {
		logger.Error("failed to validate registration data: ", e.Error())
		return api.ErrorStr(c, e.Error())
	} else {
		// user entered data is valid. register it
		ctl.ns.Register(req)
		return api.SendSuccess(c, "contract successfully registered in naming service", req.Id())
	}
}

func (ctl *ContractNameSpaceController) unregister(c echo.Context) error {
	id := c.Param("id")
	ctl.ns.Unregister(id)
	return api.Success(c, "contract successfully unregistered from naming service", id)
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
	logger.Info("exposing cns controller methods")
	router.POST("/registry", ctl.register)
	router.DELETE("/registry/:id", ctl.unregister)
	router.GET("/registry/:id", ctl.resolve)
}
