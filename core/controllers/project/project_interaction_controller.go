// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package project

import (
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

type ProjectInteractionController struct {
}

// constructor like function
func NewProjectInteractionController() ProjectInteractionController {
	pc := ProjectInteractionController{}
	return pc
}

// implemented method from interface RouterRegistrable
func (ctl ProjectInteractionController) RegisterRouters(router *echo.Group) {
	logger.Info("exposing custom projects interaction controller methods")
	router.GET(":project/:version/:operation", nil)
	router.POST(":project/:version/:operation", nil)
}
