// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package network

import (
	"strconv"

	"github.com/zerjioang/etherniti/core/data"

	"github.com/zerjioang/etherniti/core/eth"
	"github.com/zerjioang/etherniti/shared/protocol"

	"github.com/zerjioang/etherniti/core/api"
	ethrpc "github.com/zerjioang/etherniti/core/eth/rpc"
	"github.com/zerjioang/etherniti/core/eth/rpc/model"
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

func (ctl *DevOpsController) deployContract(c *echo.Context) error {

	//new deploy contract request
	req := protocol.DeployRequest{}
	if err := c.Bind(&req); err != nil {
		// return a binding error
		logger.Error("failed to bind request data to model: ", err)
		return api.ErrorBytes(c, data.BindErr)
	}

	// read from value
	from, callerErr := ctl.network.getCallerAddress(c)
	if callerErr != nil {
		return api.Error(c, callerErr)
	}

	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)

	if cliErr != nil {
		return api.Error(c, cliErr)
	}

	// detect if the target network is ganache
	ganacheDetected, _ := client.IsGanache()
	if ganacheDetected {
		// when using truffle
		// For each network, if unspecified, transaction options will default to the following values:
		// gas: Gas limit used for deploys. Default is 4712388.
		// gasPrice: Gas price used for deploys. Default is 100000000000 (100 Shannon).
		// from: From address used during migrations. Defaults to the first available account provided by your Ethereum client.
	}
	raw, err := client.DeployContract(from, req.Contract, "4712388", "100000000000")
	if err != nil {
		return api.Error(c, err)
	} else {
		return api.SendSuccess(c, data.Deploy, raw)
	}
}

// eth.sendTransaction({from:sender, to:receiver, value: amount})
func (ctl *DevOpsController) sendTransaction(c *echo.Context) error {
	to := c.Param("to")
	//input data validation
	if to == "" {
		return api.ErrorBytes(c, data.InvalidDstAddress)
	}
	amount := c.Param("amount")
	tokenAmount, pErr := strconv.Atoi(amount)
	//input data validation
	if amount == "" || pErr != nil || tokenAmount <= 0 {
		return api.ErrorBytes(c, data.InvalidEtherValue)
	}
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)

	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	//build our transaction
	var transaction ethrpc.TransactionData
	transaction.To = to
	transaction.Value = eth.ToWei(tokenAmount, 0)

	raw, err := client.EthSendTransaction(transaction)
	if err != nil {
		return api.Error(c, err)
	} else {
		return api.SendSuccess(c, data.Allowance, raw)
	}
}

func (ctl *DevOpsController) callContract(c *echo.Context) error {
	contractAddress := c.Param("contract")
	//input data validation
	if contractAddress == "" {
		return api.ErrorBytes(c, data.InvalidContractAddress)
	}
	methodName := c.Param("method")
	if methodName == "" {
		return api.ErrorBytes(c, data.InvalidMethodName)
	}
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)
	logger.Info("devops controller request using context id: ", ctl.network.networkName)
	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	var params string
	if c.Request().Method == "post" {
		//post sent. params should be sent in body too

	}

	raw, err := client.ContractCall(contractAddress, methodName, params, model.LatestBlockNumber, "", "")
	if err != nil {
		return api.Error(c, err)
	} else {
		return api.SendSuccess(c, data.Allowance, raw)
	}
}

// END of ERC20 functions

// implemented method from interface RouterRegistrable
func (ctl DevOpsController) RegisterRouters(router *echo.Group) {
	router.POST("/devops/deploy", ctl.deployContract)
	router.GET("/devops/call/:contract/:method", ctl.callContract)
	router.POST("/devops/call/:contract/:method", ctl.callContract)
}
