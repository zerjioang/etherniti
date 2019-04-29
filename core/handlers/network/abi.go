// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package network

import (
	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/data"
	"github.com/zerjioang/etherniti/core/eth/fixtures/abi"
	"github.com/zerjioang/etherniti/core/handlers/clientcache"
	"github.com/zerjioang/etherniti/core/modules/concurrentmap"
	"github.com/zerjioang/etherniti/shared/constants"

	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

// eth abi controller
type AbiController struct {
	abiData concurrentmap.ConcurrentMap
}



// constructor like function
func NewAbiController() AbiController {
	ctl := AbiController{}
	ctl.abiData = concurrentmap.New()
	return ctl
}

// profile abi data getter
func (ctl *AbiController) getAbi(c echo.ContextInterface) error {
	var code int
	code, c = clientcache.Cached(c, true, 10) // cache policy: 10 seconds

	contractAddress := c.Param("contract")
	if contractAddress == "" {
		return api.ErrorStr(c, data.ProvideContractName)
	} else {
		d, found := ctl.abiData.Get(contractAddress)
		if found {
			return c.JSON(code, d)
		} else {
			return api.ErrorStr(c, data.NoResults)
		}
	}
}

// profile abi data setter
func (ctl *AbiController) setAbi(c echo.ContextInterface) error {
	_, c = clientcache.Cached(c, true, 10) // cache policy: 10 seconds

	contractAddress := c.Param("contract")
	if contractAddress == "" {
		return api.ErrorStr(c, data.ProvideContractAddress)
	} else {
		// read body abi data, if exists
		req := abi.ABI{}
		if err := c.Bind(&req); err != nil {
			// return a binding error
			logger.Error("failed to bind request data to model: ", err)
			return api.ErrorStr(c, constants.BindErr)
		}
		if req.Methods != nil && len(req.Methods) > 0 {
			ctl.abiData.Set(contractAddress, req)
			return api.Success(c, data.LinkSuccess, nil)
		} else {
			return api.ErrorStr(c, data.InvalidAbi)
		}
	}
}

// implemented method from interface RouterRegistrable
func (ctl AbiController) RegisterRouters(router *echo.Group) {
	router.GET("/abi/:contract", ctl.getAbi)
	router.POST("/abi/:contract", ctl.setAbi)
}
