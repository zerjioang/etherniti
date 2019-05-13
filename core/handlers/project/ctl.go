// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package project

import (
	"encoding/hex"

	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/data"
	"github.com/zerjioang/etherniti/core/db"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/util/ip"
	"github.com/zerjioang/etherniti/core/util/str"
	"github.com/zerjioang/etherniti/shared/protocol"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

type ProjectController struct {
	storage *db.Db
}

// constructor like function
func NewProjectController() ProjectController {
	pc := ProjectController{}
	err := pc.initStorage()
	if err != nil {
		logger.Error("failed to initialize project controller database module")
	}
	return pc
}

func (ctl *ProjectController) initStorage() error {
	var err error
	ctl.storage, err = db.NewCollection("project")
	if err != nil {
		logger.Error("failed to initialize projects db collection: ", err)
	}
	return err
}

func (ctl *ProjectController) Create(c *echo.Context) error {

	//new project creation request
	var req protocol.ProjectRequest
	if err := c.Bind(&req); err != nil {
		// return a binding error
		logger.Error("failed to bind request data to model: ", err)
		return api.ErrorStr(c, data.BindErr)
	}
	e := req.Validate()
	if e.Occur() {
		logger.Error("failed to validate new project data: ", e.Error())
		return api.ErrorStr(c, str.UnsafeBytes(e.Error()))
	} else {
		// get required data to build a new project
		//get user ip
		c.RealIP()
		intIP := ip.Ip2int(c.RealIP())
		// get user uuid
		projectOwner := c.UserUuid()

		if intIP == 0 || projectOwner == "" {
			logger.Error("failed to create new project: missing data")
			return api.ErrorStr(c, str.UnsafeBytes("missing project data"))
		}

		p := NewProject(req.Name, projectOwner, intIP)
		k, v := p.Storage()
		writeErr := ctl.storage.PutUniqueKeyValue(k, v)
		if writeErr != nil {
			return api.Error(c, writeErr)
		} else {
			return api.SendSuccess(c, []byte("project successfully created"), p.Uuid)
		}
	}
}

func (ctl *ProjectController) Read(c *echo.Context) error {
	projectId := c.Param("id")
	if projectId != "" {
		raw, err := hex.DecodeString(projectId)
		if err != nil {
			return api.Error(c, err)
		}
		// todo check if current project id belongs to current user
		projectData, readErr := ctl.storage.GetKeyValue(raw)
		if readErr != nil {
			return api.Error(c, readErr)
		}
		return api.SendSuccessBlob(c, projectData)
	} else {
		return api.ErrorStr(c, []byte("provide a project id"))
	}
}

func (ctl *ProjectController) Update(c *echo.Context) error {
	projectId := c.Param("id")
	if projectId != "" {
		raw, err := hex.DecodeString(projectId)
		if err != nil {
			return api.Error(c, err)
		}
		// todo check if current project id belongs to current user
		projectData, readErr := ctl.storage.GetKeyValue(raw)
		if readErr != nil {
			return api.Error(c, readErr)
		}
		return api.SendSuccessBlob(c, projectData)
	} else {
		return api.ErrorStr(c, []byte("provide a project id"))
	}
}

func (ctl *ProjectController) Delete(c *echo.Context) error {
	projectId := c.Param("id")
	if projectId != "" {
		raw, err := hex.DecodeString(projectId)
		if err != nil {
			return api.Error(c, err)
		}
		// todo check if current project id belongs to current user
		deleteErr := ctl.storage.Delete(raw)
		if deleteErr != nil {
			return api.Error(c, deleteErr)
		}
		return api.SendSuccessBlob(c, []byte("project successfully deleted"))
	} else {
		return api.ErrorStr(c, []byte("provide a project id"))
	}
}

func (ctl *ProjectController) List(c *echo.Context) error {
	results, err := ctl.storage.List("")
	if err != nil {
		return api.Error(c, err)
	} else {
		return api.SendSuccess(c, []byte("list"), results)
	}
}

// implemented method from interface RouterRegistrable
func (ctl ProjectController) RegisterRouters(router *echo.Group) {
	logger.Info("exposing project controller methods")
	router.GET("/projects", ctl.List)
	router.POST("/project", ctl.Create)
	router.GET("/project/:id", ctl.Read)
	router.PUT("/project/:id", ctl.Update)
	router.DELETE("/project/:id", ctl.Delete)
}
