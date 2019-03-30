// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"github.com/labstack/echo"
	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/constants"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/server"
)

// jwt middleware function.
func jwt(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// convert context in etherniti context
		cc := c.(*server.EthernitiContext)
		token := cc.ReadConnectionProfileToken()
		if token == "" {
			return api.ErrorStr(c, "please provide a connection profile token for this kind of call")
		}

		_, parseErr := cc.ConnectionProfileSetup()
		if parseErr != nil {
			return api.Error(c, parseErr)
		}
		return next(cc)
	}
}

// create a group for all /api/v1 functions
func next(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}

// RegisterServices in echo server, allowed routes
func RegisterServices(e *echo.Echo) *echo.Group {
	group := e.Group(constants.ApiVersion, next)
	logger.Info("registering context free routes")

	publicGroup := group.Group(constants.PublicApi, next)

	NewIndexController().RegisterRouters(publicGroup)
	NewProfileController().RegisterRouters(publicGroup)
	NewSecurityController().RegisterRouters(publicGroup)
	NewWalletController().RegisterRouters(publicGroup)
	NewSolcController().RegisterRouters(publicGroup)
	NewContractNameSpaceController().RegisterRouters(publicGroup)

	//register public ethereum network related services
	NewRopstenController().RegisterRouters(publicGroup)
	NewRinkebyController().RegisterRouters(publicGroup)
	NewKovanController().RegisterRouters(publicGroup)
	NewMainNetController().RegisterRouters(publicGroup)

	privateGroup := group.Group(constants.PrivateApi, next)
	//add jwt middleware to private group
	privateGroup.Use(jwt)
	// add private or context dependant services
	NewPrivateNetController().RegisterRouters(privateGroup)
	NewDevOpsController().RegisterRouters(privateGroup)
	//NewTokenController(deployer.manager).RegisterRouters(privateGroup)
	return group
}
