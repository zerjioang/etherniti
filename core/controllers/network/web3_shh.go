// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package network

import (
	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/data"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

// eth web3 controller
type Web3ShhController struct {
	network *NetworkController
}

// constructor like function
func NewWeb3ShhController(network *NetworkController) Web3ShhController {
	ctl := Web3ShhController{}
	ctl.network = network
	return ctl
}

// BEGIN of web3 shh functions

// ShhVersion calls shh protocol shh_version json-rpc call
func (ctl *Web3ShhController) shhVersion(c *echo.Context) error {
	logger.Info("web3 shh controller request using context id: ", ctl.network.networkName)

	// get our client context
	client, cliErr := ctl.network.getRpcClient(c)
	if cliErr != nil {
		return api.Error(c, cliErr)
	}

	response, err := client.ShhVersion()
	if err != nil {
		// send invalid response message
		return api.Error(c, err)
	} else {
		ctl.network.cache.Set(data.ShhVersion, response)
		return api.SendSuccess(c, data.ShhVersion, response)
	}
}

// ShhPost calls shh protocol shh_post json-rpc call
func (ctl *Web3ShhController) shhPost(c *echo.Context) error {
	return errNotImplemented
}

// ShhNewIdentity calls shh protocol shh_newidentity json-rpc call
func (ctl *Web3ShhController) shhNewIdentity(c *echo.Context) error {
	return errNotImplemented
}

// ShhHasIdentity calls shh protocol shh_hasidentity json-rpc call
func (ctl *Web3ShhController) shhHasIdentity(c *echo.Context) error {
	return errNotImplemented
}

// ShhNewGroup calls shh protocol shh_newgroup json-rpc call
func (ctl *Web3ShhController) shhNewGroup(c *echo.Context) error {
	return errNotImplemented
}

// ShhAddToGroup calls shh protocol shh_addtogroup json-rpc call
func (ctl *Web3ShhController) shhAddToGroup(c *echo.Context) error {
	return errNotImplemented
}

// ShhNewFilter calls shh protocol shh_newfilter json-rpc call
func (ctl *Web3ShhController) shhNewFilter(c *echo.Context) error {
	return nil
}

// ShhUninstallFilter calls shh protocol shh_uninstallfilter json-rpc call
func (ctl *Web3ShhController) shhUninstallFilter(c *echo.Context) error {
	return errNotImplemented
}

// ShhGetFilterChanges calls shh protocol shh_getfilterchanges json-rpc call
func (ctl *Web3ShhController) shhGetFilterChanges(c *echo.Context) error {
	return errNotImplemented
}

// ShhGetMessages calls shh protocol shh_getmessages json-rpc call
func (ctl *Web3ShhController) shhGetMessages(c *echo.Context) error {
	return errNotImplemented
}

// END of web3 shh functions

// implemented method from interface RouterRegistrable
func (ctl Web3ShhController) RegisterRouters(router *echo.Group) {

	logger.Debug("registering eth_shh methods")

	router.GET("/shh/version", ctl.shhVersion)

	router.POST("/shh/message", ctl.shhPost)

	router.POST("/shh/identity", ctl.shhNewIdentity)
	router.GET("/shh/identity/:address", ctl.shhHasIdentity)

	router.POST("/shh/group", ctl.shhNewGroup)
	router.POST("/shh/group/add", ctl.shhAddToGroup)

	router.POST("/shh/filter", ctl.shhNewFilter)
	router.DELETE("/shh/filter", ctl.shhUninstallFilter)
	router.GET("/shh/filter/changes", ctl.shhGetFilterChanges)

	router.GET("/shh/messages", ctl.shhGetMessages)
}
