// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package project

import (
	"errors"

	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/data"
	"github.com/zerjioang/etherniti/core/eth/rpc"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/model/project"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

var (
	errInvalidParams = errors.New("provide a valid project name or version")
)

type ProjectInteractionController struct {
	projects ProjectController
}

// constructor like function
func NewProjectInteractionController(p ProjectController) ProjectInteractionController {
	pc := ProjectInteractionController{}
	pc.projects = p
	return pc
}

// constructor like function
func NewProjectInteractionControllerPtr(p *ProjectController) *ProjectInteractionController {
	pc := new(ProjectInteractionController)
	pc.projects = *p
	return pc
}

func (ctl ProjectInteractionController) getProjectData(uid string, name string, version string) (*project.Project, error) {
	// todo optimize validation: trim string and deep version check
	if name == "" || version == "" {
		return nil, errInvalidParams
	}
	return ctl.projects.ReadProject(name, version)
}

func (ctl *ProjectInteractionController) getProjectDataPtr(uid string, name string, version string) (*project.Project, error) {
	// todo optimize validation: trim string and deep version check
	if name == "" || version == "" {
		return nil, errInvalidParams
	}
	return ctl.projects.ReadProject(name, version)
}

func (ctl *ProjectInteractionController) contractCall(context *echo.Context) error {
	// read user authentication uuid
	uid := context.UserUuid()
	if uid == "" {
		logger.Error("could not read user authorization from request content")
		return api.ErrorStr(context, []byte("unknown user. operation denied"))
	}
	// read user project params: name and version/tag only
	name := context.Param("project")
	version := context.Param("version")
	methodName := context.Param("operation")
	// try to read requested project by name and user id
	projectData, err := ctl.projects.ReadProject(uid, name)
	if err != nil {
		logger.Error("failed to read project data: ", err)
		return api.Error(context, err)
	}
	//recover connection details from project data
	endpoint := projectData.Endpoint
	contractVersion, err := projectData.ResolveContract(version)
	if err != nil {
		logger.Error("failed to fetch requested contract version. Make sure it is exist and version data is correct: ", err)
		return api.Error(context, err)
	}

	params := ""
	gas := projectData.Gas
	gasprice := projectData.GasPrice
	block := projectData.Block

	isDebug := true

	// build an rpc client
	web3Client := ethrpc.NewDefaultRPCPtr(endpoint, isDebug)
	// proxy pass user request
	result, err := web3Client.ContractCall(
		contractVersion.Address,
		methodName,
		params,
		block,
		gas,
		gasprice,
	)
	if err != nil {
		logger.Error("failed to call contract: ", err)
		return api.Error(context, err)
	}
	return api.SendSuccess(context, []byte("contract_call"), []byte(result))
}

func (ctl *ProjectInteractionController) sendTransaction(context *echo.Context) error {
	// read user authentication uuid
	uid := context.UserUuid()
	if uid == "" {
		logger.Error("could not read user authorization from request content")
		return api.ErrorStr(context, []byte("unknown user. operation denied"))
	}
	// read user project params: name and version/tag only
	name := context.Param("project")
	version := context.Param("version")
	methodName := context.Param("operation")
	if methodName == "" {
		logger.Error("failed to execute transaction request. operation is required")
		return api.ErrorStr(context, []byte("an operation name is required"))
	}
	// try to read requested project by name and user id
	projectData, err := ctl.projects.ReadProject(uid, name)
	if err != nil {
		logger.Error("failed to read project data: ", err)
		return api.Error(context, err)
	}
	//recover connection details from project data
	endpoint := projectData.Endpoint
	contractVersion, err := projectData.ResolveContract(version)
	if err != nil {
		logger.Error("failed to fetch requested contract version. Make sure it is exist and version data is correct: ", err)
		return api.Error(context, err)
	}

	isDebug := true
	// build an rpc client
	web3Client := ethrpc.NewDefaultRPCPtr(endpoint, isDebug)
	// proxy pass user request
	result, err := web3Client.EthSendTransactionPtr(&ethrpc.TransactionData{
		From:     "",
		To:       contractVersion.Address,
		Gas:      0,
		GasPrice: nil,
		Value:    nil,
		Data:     "",
		Nonce:    0,
	})
	if err != nil {
		logger.Error("failed to call contract: ", err)
		return api.Error(context, err)
	}
	return api.SendSuccess(context, []byte("contract_call"), []byte(result))
}

func (ctl *ProjectInteractionController) proxyPassProject(context *echo.Context) error {
	switch context.Request().Method {
	case "GET":
		// contract call
		return ctl.contractCall(context)
	case "POST":
		// send transaction
		return ctl.sendTransaction(context)
	default:
		// operation not supported
		return api.StackError(context, data.ErrOperationNotSupported)
	}
}

// implemented method from interface RouterRegistrable
func (ctl ProjectInteractionController) RegisterRouters(router *echo.Group) {
	logger.Info("exposing custom projects interaction controller methods")
	router.GET(":project/:version/:operation", ctl.proxyPassProject)
	router.POST(":project/:version/:operation", ctl.proxyPassProject)
}
