package project

import (
	"github.com/zerjioang/etherniti/core/modules/fastime"
	"github.com/zerjioang/etherniti/core/util/id"
	"github.com/zerjioang/etherniti/core/util/str"
)

type Project struct {
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

func (project Project) Storage() (key []byte, value []byte) {
	// set storage key, project uuid
	key = str.UnsafeBytes(project.Uuid)
	// set value data as project struct values as json
	value = str.GetJsonBytes(project)
	return
}

func NewEmptyProject() Project {
	return Project{}
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
