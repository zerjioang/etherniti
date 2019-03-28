// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"github.com/zerjioang/etherniti/core/handlers/clientcache"

	"github.com/labstack/echo"
	"github.com/zerjioang/etherniti/core/logger"
)

type DevOpsController struct {
}

func init() {
}

func NewDevOpsController() DevOpsController {
	ctl := DevOpsController{}
	return ctl
}

func (ctl *DevOpsController) deployContract(c echo.Context) error {
	var code int
	code, c = clientcache.Cached(c, true, 5) // 5 seconds cache directive
	return c.JSONBlob(code, []byte{})
}

// implemented method from interface RouterRegistrable
func (ctl DevOpsController) RegisterRouters(router *echo.Group) {
	logger.Info("exposing devops controller methods")
	router.POST("/devops/deploy", ctl.deployContract)
}
