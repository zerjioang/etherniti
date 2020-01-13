// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package network

import (
	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/controllers/wrap"
	"github.com/zerjioang/etherniti/core/data"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/shared"
	"github.com/zerjioang/etherniti/shared/constants"
	"github.com/zerjioang/etherniti/shared/dto"
	"github.com/zerjioang/go-hpc/thirdparty/echo"
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

// BEGIN of web3 db-badger functions

// dbPutString calls db-badger protocol db_putString json-rpc call
func (ctl *Web3DbController) dbPutString(c *shared.EthernitiContext) error {
	var req *dto.DbStorageRequest
	if err := c.Bind(&req); err != nil {
		// return a binding error
		logger.Error(data.FailedToBind, err)
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

// dbGetString calls db-badger protocol db_getString json-rpc call
func (ctl *Web3DbController) dbGetString(c *shared.EthernitiContext) error {
	db := c.Param("db-badger")
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

// dbPutHex calls db-badger protocol db_putHex json-rpc call
func (ctl *Web3DbController) dbPutHex(c *shared.EthernitiContext) error {
	var req *dto.DbStorageRequest
	if err := c.Bind(&req); err != nil {
		// return a binding error
		logger.Error(data.FailedToBind, err)
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

// dbGetHex calls db-badger protocol db_getHex json-rpc call
func (ctl *Web3DbController) dbGetHex(c *shared.EthernitiContext) error {
	db := c.Param("db-badger")
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

// END of web3 db-badger functions

// implemented method from interface RouterRegistrable
func (ctl Web3DbController) RegisterRouters(router *echo.Group) {
	logger.Debug("registerind eth_db methods")

	router.POST("/db-badger/string", wrap.Call(ctl.dbPutString))
	router.GET("/db-badger/string/:db-badger/:key", wrap.Call(ctl.dbGetString))

	router.POST("/db-badger/hex", wrap.Call(ctl.dbPutHex))
	router.GET("/db-badger/hex/:db-badger/:key", wrap.Call(ctl.dbGetHex))
}
