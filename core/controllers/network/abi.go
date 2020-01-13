// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package network

import (
	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/controllers/wrap"
	"github.com/zerjioang/etherniti/core/data"
	"github.com/zerjioang/etherniti/shared"
	"github.com/zerjioang/go-hpc/lib/codes"
	"github.com/zerjioang/go-hpc/lib/concurrentmap"
	"github.com/zerjioang/go-hpc/lib/eth/fixtures/abi"

	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/go-hpc/thirdparty/echo"
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
func (ctl *AbiController) getAbi(c *shared.EthernitiContext) error {
	c.OnSuccessCachePolicy = 10
	contractAddress := c.Param("contract")
	if contractAddress == "" {
		return api.ErrorBytes(c, data.ProvideContractName)
	} else {
		d, found := ctl.abiData.Get(contractAddress)
		if found {
			return c.JSON(codes.StatusOK, d)
		} else {
			return api.ErrorBytes(c, data.NoResults)
		}
	}
}

// profile abi data setter
func (ctl *AbiController) setAbi(c *shared.EthernitiContext) error {
	c.OnSuccessCachePolicy = 10

	contractAddress := c.Param("contract")
	if contractAddress == "" {
		return api.ErrorBytes(c, data.ProvideContractAddress)
	} else {
		// read body abi data, if exists
		req := abi.ABI{}
		if err := c.Bind(&req); err != nil {
			// return a binding error
			logger.Error(data.FailedToBind, err)
			return api.ErrorBytes(c, data.BindErr)
		}
		if req.Methods != nil && len(req.Methods) > 0 {
			ctl.abiData.Set(contractAddress, req)
			return api.Success(c, data.LinkSuccess, nil)
		} else {
			return api.ErrorBytes(c, data.InvalidAbi)
		}
	}
}

// implemented method from interface RouterRegistrable
func (ctl AbiController) RegisterRouters(router *echo.Group) {
	router.GET("/abi/:contract", wrap.Call(ctl.getAbi))
	router.POST("/abi/:contract", wrap.Call(ctl.setAbi))
}
