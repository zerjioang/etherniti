// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package dashboard

import (
	"errors"

	"github.com/valyala/fasthttp"

	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

var (
	ErrFeatureNotAvailable = errors.New("feature not available in your current etherniti proxy instance")
)

// user interface helper controller
type UIController struct {
	// http client
	client *fasthttp.Client
}

// constructor like function
func NewUIController(client *fasthttp.Client) UIController {
	ctl := UIController{}
	ctl.client = client
	return ctl
}

// received information about specified buf from a client
func (ctl UIController) bugReport(c *echo.Context) error {
	return api.Error(c, ErrFeatureNotAvailable)
}

// implemented method from interface RouterRegistrable
func (ctl UIController) RegisterRouters(router *echo.Group) {
	logger.Info("exposing external controller methods")
	router.POST("/bug", ctl.bugReport)
}
