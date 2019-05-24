// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package project

import (
	"github.com/zerjioang/etherniti/core/controllers/common"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/model/project"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

type ProjectController struct {
	common.DatabaseController
}

// constructor like function
func NewProjectController() ProjectController {
	pc := ProjectController{}
	var err error
	pc.DatabaseController, err = common.NewDatabaseController("projects", project.NewDBProject)
	if err != nil {
		logger.Error("failed to create project controller ", err)
	}
	return pc
}

// implemented method from interface RouterRegistrable
func (ctl ProjectController) RegisterRouters(router *echo.Group) {
	logger.Info("exposing custom projects controller methods")
	ctl.DatabaseController.RegisterRouters(router)
}
