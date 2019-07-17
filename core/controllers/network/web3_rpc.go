// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package network

import (
	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/thirdparty/echo"
	"strings"
)

// eth web3 rpc controller
type Web3RpcController struct {
	network *NetworkController
}

// constructor like function
func NewWeb3RpcController(network *NetworkController) Web3RpcController {
	ctl := Web3RpcController{}
	ctl.network = network
	return ctl
}

// proxy pass client rpc request to appropiate target node
func (ctl *Web3RpcController) rpc(c *echo.Context) error {
	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)
	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	// once we have the rpc client, proxy pass the message
	response, reqErr := client.Proxy(c.Body())
	if reqErr.Occur() {
		// if error is a connection refused error or contains http://
		// lets hide full error content since it can disclose sensitive information such as peer node IP
		errStr := reqErr.Error()
		if strings.Contains(errStr, "connection refused"){
			return api.ErrorStr(c, "connection with the peer was refused")
		} else if strings.Contains(errStr, "http://"){
			return api.ErrorStr(c, "http connection with the peer was unsuccessful")
		} else {
			return api.StackError(c, reqErr)
		}
	}
	//send back to the client received response
	return api.SendRawSuccess(c, response)
}

// implemented method from interface RouterRegistrable
func (ctl Web3RpcController) RegisterRouters(router *echo.Group) {
	logger.Debug("adding controller raw JSON-RPC call supports")
	logger.Debug("exposing POST /rpc")
	router.POST("/rpc", ctl.rpc)
}
