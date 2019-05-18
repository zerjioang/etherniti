package project

import (
	"github.com/zerjioang/etherniti/core/controllers/common"
	"github.com/zerjioang/etherniti/core/data"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/modules/fastime"
	"github.com/zerjioang/etherniti/core/modules/stack"
	"github.com/zerjioang/etherniti/core/util/id"
	"github.com/zerjioang/etherniti/core/util/ip"
	"github.com/zerjioang/etherniti/core/util/str"
	"github.com/zerjioang/etherniti/shared/protocol"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

type Project struct {
	// implement interface to be a rest-db-crud able struct
	common.DatabaseObjectInterface

	// unique project indetifier used for database storage
	Uuid string `json:"uuid"`
	// project name
	Name string `json:"name"`
	// owner id (sha256 of owner user email)
	Owner string `json:"owner"`
	// internal project id assigned
	ProjectId string `json:"id"`
	// internal project secret id assigned
	ProjectSecret string `json:"secret"`
	// project creation date
	CreationDate int64 `json:"created"`
	// ip address who created the project
	Ip uint32 `json:"ip"`
}

// implementation of interface DatabaseObjectInterface
func (project Project) Key() []byte {
	return str.UnsafeBytes(project.Uuid)
}
func (project Project) Value() []byte {
	return str.GetJsonBytes(project)
}
func (project Project) New() common.DatabaseObjectInterface {
	return NewEmptyProject()
}

// custom validation logic for read operation
// return nil if everyone can read
func (project Project) CanRead(context *echo.Context, key string) error {
	// todo check if current project id belongs to current user
	return nil
}

// custom validation logic for update operation
// return nil if everyone can update
func (project Project) CanUpdate(context *echo.Context, key string) error {
	// todo check if current project id belongs to current user
	return nil
}

// custom validation logic for delete operation
// return nil if everyone can delete
func (project Project) CanDelete(context *echo.Context, key string) error {
	// todo check if current project id belongs to current user
	return nil
}

// custom validation logic for write operation
// return nil if everyone can write
func (project Project) CanWrite(context *echo.Context) error {
	return nil
}

// custom validation logic for list operation
// return nil if everyone can list
func (project Project) CanList(context *echo.Context) error {
	// todo check if current project id belongs to current user
	return nil
}

func (project Project) Bind(context *echo.Context) (common.DatabaseObjectInterface, stack.Error) {
	//new project creation request
	var req protocol.ProjectRequest
	if err := context.Bind(&req); err != nil {
		// return a binding error
		logger.Error("failed to bind request data to model: ", err)
		return nil, data.ErrBind
	}
	e := req.Validate()
	if e.Occur() {
		logger.Error("failed to validate new project data: ", e.Error())
		return nil, e
	} else {
		// get required data to build a new project
		intIP := ip.Ip2int(context.RealIP())
		// get user uuid
		projectOwner := context.UserUuid()

		if intIP == 0 || projectOwner == "" {
			logger.Error("failed to create new project: missing data")
			return nil, data.StackErrProject
		} else {
			p := NewProject(req.Name, projectOwner, intIP)
			return p, stack.Nil()
		}
	}
}

func NewEmptyProject() Project {
	return Project{}
}

func NewDBProject() common.DatabaseObjectInterface {
	return NewEmptyProject()
}

func NewProject(name string, owner string, ip uint32) *Project {
	p := new(Project)
	p.Uuid = id.GenerateIDString().String()
	p.Name = name
	p.Owner = owner
	p.CreationDate = fastime.Now().Unix()
	p.ProjectId = id.GenerateIDString().String()
	p.ProjectSecret = id.GenerateIDString().String()
	p.Ip = ip
	return p
}
