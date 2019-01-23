// Copyright MethW
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"net/http"

	"github.com/zerjioang/methw/shared"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type IndexController struct {
	shared.AutoRouteable
}

const (
	indexWelcome = `{
  "name": "eth-wbapi",
  "description": "MethW: Ethereum Multitenant API",
  "cluster_name": "eth-wbapi",
  "version": "0.0.1",
  "env": "development",
  "tagline": "dapps everywhere"
}`
)

func NewIndexController() IndexController {
	dc := IndexController{}
	return dc
}

func (ctl IndexController) index(c echo.Context) error {
	return c.String(http.StatusOK, indexWelcome)
}

/*
implemented method from interface RouterRegistrable
*/
func (ctl IndexController) RegisterRouters(router *echo.Echo) {
	log.Debug("exposing GET /")
	router.GET("/", ctl.index)
}
