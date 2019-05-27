package project

import (
	"errors"

	"github.com/zerjioang/etherniti/core/data"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/model/metadata"
	"github.com/zerjioang/etherniti/core/model/registry/version"
	"github.com/zerjioang/etherniti/core/modules/stack"
	"github.com/zerjioang/etherniti/core/util/id"
	"github.com/zerjioang/etherniti/core/util/ip"
	"github.com/zerjioang/etherniti/core/util/str"
	"github.com/zerjioang/etherniti/shared/mixed"
	"github.com/zerjioang/etherniti/shared/protocol"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

type Project struct {
	// implement interface to be a rest-db-crud able struct
	mixed.DatabaseObjectInterface `json:"_,omitempty"`

	// unique project identifier used for database storage
	Uuid string `json:"sid"`

	// project name
	Name string `json:"name"`
	// project description
	Description string `json:"description"`
	// project logo url
	ImageUrl string `json:"image,omitempty"`

	// internal project id assigned
	ProjectId string `json:"id"`
	// internal project secret id assigned
	ProjectSecret string `json:"secret"`

	//connection required data

	// peer endpoint url
	Endpoint string
	// default gas value
	Gas string
	// default gasprice value
	GasPrice string
	// default target block: latest by default
	Block string

	//list of linked contracts to this project
	// usually each entry in the array means a
	// deployed version of project's contract
	Contracts map[string]*version.ContractVersion `json:"contracts"`

	// project metadata
	Metadata *metadata.Metadata `json:"metadata,omitempty"`
}

// implementation of interface DatabaseObjectInterface
func (project Project) Key() []byte {
	return str.UnsafeBytes(project.Uuid)
}
func (project Project) Value() []byte {
	return str.GetJsonBytes(project)
}
func (project Project) New() mixed.DatabaseObjectInterface {
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

func (project Project) Bind(context *echo.Context) (mixed.DatabaseObjectInterface, stack.Error) {
	//new project creation request
	var req protocol.ProjectRequest
	if err := context.Bind(&req); err != nil {
		// return a binding error
		logger.Error("failed to bind request data to model: ", err)
		return nil, data.ErrBind
	}
	// todo optimize this process
	// external clients will neber be able to bind some fields, so delete them
	project.Uuid = ""
	project.ProjectId = ""
	project.ProjectSecret = ""
	project.Metadata = nil

	e := req.Validate()
	if e.Occur() {
		logger.Error("failed to validate new project data: ", e.Error())
		return nil, e
	} else {
		// get required data to build a new project
		intIP := ip.Ip2int(context.RealIP())
		// get user uuid
		projectOwner := context.AuthenticatedUserUuid()

		if intIP == 0 || projectOwner == "" {
			logger.Error("failed to create new project: missing data")
			return nil, data.StackErrProject
		} else {
			project.Metadata = metadata.NewMetadata(context)
			return project, stack.Nil()
		}
	}
}

func (project Project) ResolveContract(version string) (*version.ContractVersion, error) {
	version = str.ToLowerAscii(version)
	details, found := project.Contracts[version]
	if found {
		return details, nil
	}
	return nil, errors.New("contract details not found")
}

func NewEmptyProject() Project {
	return Project{}
}

func NewDBProject() mixed.DatabaseObjectInterface {
	return NewEmptyProject()
}

func NewProject(name string, mtdt *metadata.Metadata) *Project {
	p := new(Project)
	p.Uuid = id.GenerateIDString().String()
	p.Name = name
	p.Metadata = mtdt
	p.ProjectId = id.GenerateIDString().String()
	p.ProjectSecret = id.GenerateIDString().String()
	return p
}
