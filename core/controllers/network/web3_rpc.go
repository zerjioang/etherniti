// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package network

import (
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

// eth web3 controller
type Web3DbController struct {
	network *NetworkController
}

// constructor like function
func NewWeb3DbController(network *NetworkController) Web3DbController {
	ctl := Web3DbController{}
	ctl.network = network
	return ctl
}

// BEGIN of web3 db functions

// dbPutString calls db protocol db_putString json-rpc call
func (ctl *Web3DbController) dbPutString(c *echo.Context) error {
	return errNotImplemented
}

// dbGetString calls db protocol db_getString json-rpc call
func (ctl *Web3DbController) dbGetString(c *echo.Context) error {
	return errNotImplemented
}

// dbPutHex calls db protocol db_putHex json-rpc call
func (ctl *Web3DbController) dbPutHex(c *echo.Context) error {
	return errNotImplemented
}

// dbGetHex calls db protocol db_getHex json-rpc call
func (ctl *Web3DbController) dbGetHex(c *echo.Context) error {
	return errNotImplemented
}

// END of web3 db functions

// implemented method from interface RouterRegistrable
func (ctl Web3DbController) RegisterRouters(router *echo.Group) {
	router.GET("/db", ctl.dbGetString)
	router.POST("/db", ctl.dbPutString)
	router.GET("/db/hex", ctl.dbGetHex)
	router.POST("/db/hex", ctl.dbPutHex)
}
