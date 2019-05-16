// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package project

import (
	"github.com/zerjioang/etherniti/core/controllers/common"
	"github.com/zerjioang/etherniti/core/logger"
)

type ProjectController struct {
	common.DatabaseController
}

// constructor like function
func NewProjectController() ProjectController {
	pc := ProjectController{}
	var err error
	pc.DatabaseController, err = common.NewDatabaseController("project", NewDBProject())
	if err != nil {
		logger.Error("failed to create database controller ", err)
	}
	return pc
}
