package project

import (
	"encoding/json"
	"errors"

	"github.com/zerjioang/etherniti/core/model"

	"github.com/zerjioang/etherniti/core/data"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/model/metadata"
	"github.com/zerjioang/etherniti/core/model/registry/version"
	"github.com/zerjioang/etherniti/core/modules/stack"
	"github.com/zerjioang/etherniti/core/util/id"
	"github.com/zerjioang/etherniti/core/util/str"
	"github.com/zerjioang/etherniti/shared/mixed"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

var (
	nilErr                    = stack.Nil()
	errContractNotFound       = errors.New("contract details not found")
	errNoContractNameProvided = stack.New("project name not provided in request")
)

type Project struct {
	// implement interface to be a rest-db-crud able struct
	mixed.DatabaseObjectInterface `json:"_,omitempty"`

	// internal project id assigned
	ProjectId string `json:"id,omitempty"`

	// project name
	Name string `json:"name,omitempty"`
	// project description
	Description string `json:"description,omitempty"`
	// project logo url
	ImageUrl string `json:"image,omitempty"`

	// internal project secret id assigned
	ProjectSecret string `json:"secret,omitempty"`

	//connection required data

	// peer endpoint url
	Endpoint string `json:"endpoint,omitempty"`
	// default gas value
	Gas string `json:"gas,omitempty"`
	// default gasprice value
	GasPrice string `json:"gasPrice,omitempty"`
	// default target block: latest by default
	Block string `json:"block,omitempty"`

	//list of linked contracts to this project
	// usually each entry in the array means a
	// deployed version of project's contract
	Contracts map[string]*version.ContractVersion `json:"contracts,omitempty"`

	// project metadata
	Metadata *metadata.Metadata `json:"metadata,omitempty"`
}

// implementation of interface DatabaseObjectInterface
func (project Project) Key() []byte {
	return str.UnsafeBytes(project.ProjectId)
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
	if err := context.Bind(&project); err != nil {
		// return a binding error
		logger.Error("failed to bind request data to model: ", err)
		return nil, data.ErrBind
	}
	// todo optimize this process
	// external clients will never be able to bind some fields, so delete them
	project.ProjectSecret = ""
	project.Metadata = nil
	project.ProjectId = id.GenerateIDString().String()

	e := project.Validate()
	if e.Occur() {
		logger.Error("failed to validate request project data: ", e.Error())
		return nil, e
	} else {
		if context.IntIp() == 0 || context.AuthenticatedUserUuid() == "" {
			logger.Error("failed to create new project: authentication data is incomplete")
			return nil, data.StackErrProject
		} else {
			project.Metadata = metadata.NewMetadata(context)
			return project, nilErr
		}
	}
}

// converts byte sequence to go project struct
func (project Project) Decode(data []byte) (mixed.DatabaseObjectInterface, stack.Error) {
	o := NewEmptyProject()
	err := json.Unmarshal(data, &o)
	return o, stack.Ret(err)
}

func (project Project) Validate() stack.Error {
	if project.Name == "" {
		return errNoContractNameProvided
	}
	return nilErr
}

func (project Project) ResolveContract(version string) (*version.ContractVersion, error) {
	version = str.ToLowerAscii(version)
	details, found := project.Contracts[version]
	if found {
		return details, nil
	}
	return nil, errContractNotFound
}

func (project Project) Update(o mixed.DatabaseObjectInterface) (mixed.DatabaseObjectInterface, stack.Error) {
	newPrj, ok := o.(*Project)
	if newPrj == nil || !ok {
		return nil, model.UnsupportedDataErr
	}
	//if new name is provided, update it
	project.Name = model.ConditionalStringUpdate(newPrj.Name, project.Name, "")
	project.Name = model.ConditionalStringUpdate(newPrj.Description, project.Description, "")
	project.ImageUrl = model.ConditionalStringUpdate(newPrj.ImageUrl, project.ImageUrl, "")
	project.Endpoint = model.ConditionalStringUpdate(newPrj.Endpoint, project.Endpoint, "")
	project.Gas = model.ConditionalStringUpdate(newPrj.Gas, project.Gas, "")
	project.GasPrice = model.ConditionalStringUpdate(newPrj.GasPrice, project.GasPrice, "")
	project.Block = model.ConditionalStringUpdate(newPrj.Block, project.Block, "")
	return project, nilErr
}

func NewEmptyProject() Project {
	return Project{}
}

func NewDBProject() mixed.DatabaseObjectInterface {
	return NewEmptyProject()
}

func NewProject(name string, mtdt *metadata.Metadata) *Project {
	p := new(Project)
	p.Name = name
	p.Metadata = mtdt
	p.ProjectId = id.GenerateIDString().String()
	p.ProjectSecret = id.GenerateIDString().String()
	return p
}
