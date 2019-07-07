// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package network

import (
	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/data"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/shared/constants"
	"github.com/zerjioang/etherniti/shared/protocol"
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
	var req *protocol.DbStorageRequest
	if err := c.Bind(&req); err != nil {
		// return a binding error
		logger.Error("failed to bind request data to model: ", err)
		return api.ErrorBytes(c, data.BindErr)
	}
	err := req.Validate()
	if err == nil {
		client, cliErr := ctl.network.getRpcClient(c)
		if cliErr != nil {
			return api.Error(c, cliErr)
		}
		result, err := client.DbPutString(req.Database, req.Key, req.Value)
		if err != nil {
			// send invalid response message
			return api.Error(c, err)
		} else {
			c.OnSuccessCachePolicy = constants.CacheInfinite
			return api.SendSuccess(c, data.PutString, result)
		}
	} else {
		// error detected on input data
		return api.Error(c, err)
	}
}

// dbGetString calls db protocol db_getString json-rpc call
func (ctl *Web3DbController) dbGetString(c *echo.Context) error {
	db := c.Param("db")
	key := c.Param("key")
	client, cliErr := ctl.network.getRpcClient(c)
	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	result, err := client.DbGetHex(db, key)
	if err != nil {
		// send invalid response message
		return api.Error(c, err)
	} else {
		c.OnSuccessCachePolicy = constants.CacheInfinite
		return api.SendSuccess(c, data.GetString, result)
	}
}

// dbPutHex calls db protocol db_putHex json-rpc call
func (ctl *Web3DbController) dbPutHex(c *echo.Context) error {
	var req *protocol.DbStorageRequest
	if err := c.Bind(&req); err != nil {
		// return a binding error
		logger.Error("failed to bind request data to model: ", err)
		return api.ErrorBytes(c, data.BindErr)
	}
	err := req.Validate()
	if err == nil {
		client, cliErr := ctl.network.getRpcClient(c)
		if cliErr != nil {
			return api.Error(c, cliErr)
		}
		result, err := client.DbPutHex(req.Database, req.Key, req.Value)
		if err != nil {
			// send invalid response message
			return api.Error(c, err)
		} else {
			c.OnSuccessCachePolicy = constants.CacheInfinite
			return api.SendSuccess(c, data.PutHex, result)
		}
	} else {
		// error detected on input data
		return api.Error(c, err)
	}
}

// dbGetHex calls db protocol db_getHex json-rpc call
func (ctl *Web3DbController) dbGetHex(c *echo.Context) error {
	db := c.Param("db")
	key := c.Param("key")
	client, cliErr := ctl.network.getRpcClient(c)
	if cliErr != nil {
		return api.Error(c, cliErr)
	}
	result, err := client.DbGetHex(db, key)
	if err != nil {
		// send invalid response message
		return api.Error(c, err)
	} else {
		c.OnSuccessCachePolicy = constants.CacheInfinite
		return api.SendSuccess(c, data.GetHex, result)
	}
}

// END of web3 db functions

// implemented method from interface RouterRegistrable
func (ctl Web3DbController) RegisterRouters(router *echo.Group) {
	logger.Debug("registerind eth_db methods")

	router.POST("/db/string", ctl.dbPutString)
	router.GET("/db/string/:db/:key", ctl.dbGetString)

	router.POST("/db/hex", ctl.dbPutHex)
	router.GET("/db/hex/:db/:key", ctl.dbGetHex)
}
