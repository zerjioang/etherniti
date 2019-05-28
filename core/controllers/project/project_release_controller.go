package project

import (
	"github.com/zerjioang/etherniti/core/controllers/common"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/model/project"
)

type ProjectReleaseController struct {
	common.DatabaseController
}

// constructor like function
func NewProjectReleaseController() ProjectReleaseController {
	pc := ProjectReleaseController{}
	var err error
	pc.DatabaseController, err = common.NewDatabaseController("/projects", "release", project.NewDBProject)
	if err != nil {
		logger.Error("failed to create project controller ", err)
	}
	return pc
}

// constructor like function
func NewProjectReleaseControllerPtr() *ProjectReleaseController {
	pc := NewProjectReleaseController()
	return &pc
}
