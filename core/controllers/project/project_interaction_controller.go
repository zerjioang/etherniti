// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package project

import (
	"errors"

	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/controllers/wrap"
	"github.com/zerjioang/etherniti/core/data"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/shared"
	"github.com/zerjioang/go-hpc/lib/eth/rpc"
	"github.com/zerjioang/go-hpc/lib/eth/rpc/client"
	"github.com/zerjioang/go-hpc/thirdparty/echo"
)

var (
	errInvalidParams = errors.New("provide a valid project name or version")
)

type ProjectInteractionController struct {
	projects *ProjectController
	client   *client.EthClient
}

// constructor like function
func NewProjectInteractionControllerPtr(p *ProjectController, client *client.EthClient) *ProjectInteractionController {
	pc := new(ProjectInteractionController)
	pc.projects = p
	pc.client = client
	return pc
}

func (ctl *ProjectInteractionController) contractCall(c *shared.EthernitiContext) error {
	// read user authentication uuid
	uid := c.AuthUserUuid()
	if uid == "" {
		logger.Error("could not read user authorization from request content")
		return api.ErrorBytes(c, []byte("unknown user. operation denied"))
	}
	// read user project params: name and version/tag only
	name := c.Param("project")
	version := c.Param("version")
	methodName := c.Param("operation")
	// try to read requested project by name and user id
	projectData, err := ctl.projects.ReadProject(uid, name)
	if err != nil {
		logger.Error("failed to read project data: ", err)
		return api.Error(c, err)
	}
	//recover connection details from project data
	endpoint := projectData.Endpoint
	contractVersion, err := projectData.ResolveContract(version)
	if err != nil {
		logger.Error("failed to fetch requested contract version. Make sure it is exist and version data is correct: ", err)
		return api.Error(c, err)
	}

	params := ""
	gas := projectData.Gas
	gasprice := projectData.GasPrice
	block := projectData.Block

	isDebug := true

	// build an rpc client
	web3Client := ethrpc.NewDefaultRPCPtr(endpoint, isDebug, ctl.client)
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
		return api.Error(c, err)
	}
	return api.SendSuccess(c, []byte("contract_call"), []byte(result))
}

func (ctl *ProjectInteractionController) sendTransaction(c *shared.EthernitiContext) error {
	// read user authentication uuid
	uid := c.AuthUserUuid()
	if uid == "" {
		logger.Error("could not read user authorization from request content")
		return api.ErrorBytes(c, []byte("unknown user. operation denied"))
	}
	// read user project params: name and version/tag only
	name := c.Param("project")
	version := c.Param("version")
	methodName := c.Param("operation")
	if methodName == "" {
		logger.Error("failed to execute transaction request. operation is required")
		return api.ErrorBytes(c, []byte("an operation name is required"))
	}
	// try to read requested project by name and user id
	projectData, err := ctl.projects.ReadProject(uid, name)
	if err != nil {
		logger.Error("failed to read project data: ", err)
		return api.Error(c, err)
	}
	//recover connection details from project data
	endpoint := projectData.Endpoint
	contractVersion, err := projectData.ResolveContract(version)
	if err != nil {
		logger.Error("failed to fetch requested contract version. Make sure it is exist and version data is correct: ", err)
		return api.Error(c, err)
	}

	isDebug := true
	// build an rpc client
	web3Client := ethrpc.NewDefaultRPCPtr(endpoint, isDebug, ctl.client)
	// proxy pass user request
	result, err := web3Client.EthSendTransactionPtr(&ethrpc.TransactionData{
		From:        "",
		To:          contractVersion.Address,
		GasStr:      "",
		GasPriceStr: "",
		ValueStr:    "",
		Data:        "",
		NonceStr:    "",
	})
	if err != nil {
		logger.Error("failed to call contract: ", err)
		return api.Error(c, err)
	}
	return api.SendSuccess(c, []byte("contract_call"), []byte(result))
}

func (ctl *ProjectInteractionController) proxyPassProject(c *shared.EthernitiContext) error {
	switch c.Request().Method {
	case "GET":
		// contract call
		return ctl.contractCall(c)
	case "POST":
		// send transaction
		return ctl.sendTransaction(c)
	default:
		// operation not supported
		return api.StackError(c, data.ErrOperationNotSupported)
	}
}

// implemented method from interface RouterRegistrable
func (ctl ProjectInteractionController) RegisterRouters(router *echo.Group) {
	logger.Info("exposing custom projects interaction controller methods")
	router.GET(":project/:version/:operation", wrap.Call(ctl.proxyPassProject))
	router.POST(":project/:version/:operation", wrap.Call(ctl.proxyPassProject))
}
