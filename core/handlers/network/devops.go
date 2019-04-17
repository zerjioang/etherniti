// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package network

import (
	"github.com/zerjioang/etherniti/core/eth/rpc"
	"github.com/zerjioang/etherniti/core/util/str"
	"net/http"

	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/shared/protocol"

	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

// eth devops controller
type DevOpsController struct {
	network *NetworkController
}

// constructor like function
func NewDevOpsController(network *NetworkController) DevOpsController {
	ctl := DevOpsController{}
	ctl.network = network
	return ctl
}

func (ctl *DevOpsController) deployContract(c echo.Context) error {
	//return ctl.sendTransaction(c)
	return errNotImplemented
}

func (ctl *DevOpsController) callContract(c echo.Context) error {
	contractAddress := c.Param("contract")
	//input data validation
	if contractAddress == "" {
		return api.ErrorStr(c, "invalid contract address provided")
	}
	methodName := c.Param("method")
	if methodName == "" {
		return api.ErrorStr(c, "invalid contract method name provided")
	}
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)
	logger.Info("devops controller request using context id: ", ctl.network.networkName)
	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	var params string
	if c.Request().Method == "post"{
		//post sent. params should be sent in body too

	}

	raw, err := client.ContractCall(contractAddress, methodName, params, ethrpc.LatestBlockNumber, "", "")
	if err != nil {
		// send invalid generation message
		return c.JSONBlob(http.StatusBadRequest,
			str.GetJsonBytes(
				protocol.NewApiError(http.StatusBadRequest, err.Error()),
			),
		)
	} else {
		return api.SendSuccess(c, "allowance", raw)
	}
}

// END of ERC20 functions

// implemented method from interface RouterRegistrable
func (ctl DevOpsController) RegisterRouters(router *echo.Group) {
	router.POST("/devops/deploy", ctl.deployContract)
	router.GET("/devops/call/:contract/:method", ctl.callContract)
	router.POST("/devops/call/:contract/:method", ctl.callContract)
}
