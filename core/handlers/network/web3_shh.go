// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package network

import (
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

// ShhVersion calls shh protocol shh_version json-rpc call
func (ctl *Web3ShhController) shhVersion(c echo.Context) error {
	return nil
}

// ShhPost calls shh protocol shh_post json-rpc call
func (ctl *Web3ShhController) shhPost(c echo.Context) error {
	return nil
}

// ShhNewIdentity calls shh protocol shh_newidentity json-rpc call
func (ctl *Web3ShhController) shhNewIdentity(c echo.Context) error {
	return nil
}

// ShhHasIdentity calls shh protocol shh_hasidentity json-rpc call
func (ctl *Web3ShhController) shhHasIdentity(c echo.Context) error {
	return nil
}

// ShhNewGroup calls shh protocol shh_newgroup json-rpc call
func (ctl *Web3ShhController) shhNewGroup(c echo.Context) error {
	return nil
}

// ShhAddToGroup calls shh protocol shh_addtogroup json-rpc call
func (ctl *Web3ShhController) shhAddToGroup(c echo.Context) error {
	return nil
}

// ShhNewFilter calls shh protocol shh_newfilter json-rpc call
func (ctl *Web3ShhController) shhNewFilter(c echo.Context) error {
	return nil
}

// ShhUninstallFilter calls shh protocol shh_uninstallfilter json-rpc call
func (ctl *Web3ShhController) shhUninstallFilter(c echo.Context) error {
	return nil
}

// ShhGetFilterChanges calls shh protocol shh_getfilterchanges json-rpc call
func (ctl *Web3ShhController) shhGetFilterChanges(c echo.Context) error {
	return nil
}

// ShhGetMessages calls shh protocol shh_getmessages json-rpc call
func (ctl *Web3ShhController) shhGetMessages(c echo.Context) error {
	return nil
}

// implemented method from interface RouterRegistrable
func (ctl Web3ShhController) RegisterRouters(router *echo.Group) {
	router.GET("/shh/version", ctl.shhVersion)
	router.GET("/shh/post", ctl.shhPost)
	router.GET("/shh/newidentity", ctl.shhNewIdentity)
	router.GET("/shh/hasidentity", ctl.shhHasIdentity)
	router.GET("/shh/newgroup", ctl.shhNewGroup)
	router.GET("/shh/addtogroup", ctl.shhAddToGroup)
	router.GET("/shh/newfilter", ctl.shhNewFilter)
	router.GET("/shh/uninstallfilter", ctl.shhUninstallFilter)
	router.GET("/shh/getfilterchanges", ctl.shhGetFilterChanges)
	router.GET("/shh/getmessages", ctl.shhGetMessages)
}
