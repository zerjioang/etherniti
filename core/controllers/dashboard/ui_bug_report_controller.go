// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package dashboard

import (
	"errors"

	"github.com/zerjioang/etherniti/core/controllers/wrap"

	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/shared"
	"github.com/zerjioang/go-hpc/thirdparty/echo"
)

var (
	ErrFeatureNotAvailable = errors.New("feature not available in your current etherniti proxy instance")
)

// user interface helper controller
type UIController struct {
}

// constructor like function
func NewUIController() UIController {
	ctl := UIController{}
	return ctl
}

// received information about specified buf from a client
func (ctl UIController) bugReport(c *shared.EthernitiContext) error {
	return api.Error(c, ErrFeatureNotAvailable)
}

// implemented method from interface RouterRegistrable
func (ctl UIController) RegisterRouters(router *echo.Group) {
	logger.Info("exposing external controller methods")
	router.POST("/bug", wrap.Call(ctl.bugReport))
}
