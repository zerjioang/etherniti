// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type IndexController struct {
}

const (
	indexWelcome = `{
  "name": "eth-wbapi",
  "description": "etherniti: Ethereum Multitenant API",
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

// implemented method from interface RouterRegistrable
func (ctl IndexController) RegisterRouters(router *echo.Echo) {
	log.Info("exposing index controller methods")
	router.GET("/v1", ctl.index)
	router.GET("/", ctl.index)
}
