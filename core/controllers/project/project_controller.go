// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package project

import (
	"github.com/zerjioang/etherniti/core/controllers/common"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/model/project"
	"github.com/zerjioang/go-hpc/thirdparty/echo"
)

type ProjectController struct {
	common.DatabaseController
}

// constructor like function
func NewProjectController() ProjectController {
	pc := ProjectController{}
	var err error
	pc.DatabaseController, err = common.NewDatabaseController("", "", "projects", project.NewDBProject)
	if err != nil {
		logger.Error("failed to create project controller ", err)
	}
	return pc
}

// constructor like function
func NewProjectControllerPtr() *ProjectController {
	pc := NewProjectController()
	return &pc
}

// helper methods
func (ctl ProjectController) getProjectData(uid string, name string, version string) (*project.Project, error) {
	// todo optimize validation: trim string and deep version check
	if name == "" || version == "" {
		return nil, errInvalidParams
	}
	return ctl.ReadProject(name, version)
}

func (ctl ProjectController) getProjectDataPtr(uid string, name string, version string) (*project.Project, error) {
	// todo optimize validation: trim string and deep version check
	if name == "" || version == "" {
		return nil, errInvalidParams
	}
	return ctl.ReadProject(name, version)
}

// implemented method from interface RouterRegistrable
func (ctl ProjectController) RegisterRouters(router *echo.Group) {
	logger.Info("exposing custom projects controller methods")
	ctl.DatabaseController.RegisterRouters(router)
	logger.Info("exposing custom projects controller methods")
}

func (ctl ProjectController) ReadProject(uid string, name string) (*project.Project, error) {
	return nil, nil
}
