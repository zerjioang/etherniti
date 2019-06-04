// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package controllers

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/controllers/project"
	"github.com/zerjioang/etherniti/core/controllers/registry"
	"github.com/zerjioang/etherniti/core/data"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/shared/constants"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

func infuraJwt(next echo.HandlerFunc) echo.HandlerFunc {
	return defaultJwt(next, data.InfuraJwtErrorMessage)
}

func quiknodeJwt(next echo.HandlerFunc) echo.HandlerFunc {
	return defaultJwt(next, data.QuiknodeJwtErrorMessage)
}

func privateJwt(next echo.HandlerFunc) echo.HandlerFunc {
	return defaultJwt(next, data.JwtErrorMessage)
}

func userJwt(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		tokenData := c.ReadToken("Authorization")
		decodedAuthData, err := ParseAuthenticationToken(tokenData)
		if err != nil {
			logger.Error("failed to process authentication token: ", err)
			return api.ErrorWithMessage(c, []byte("invalid authentication token"), err)
		} else {
			c.UserId = decodedAuthData.Uuid
		}
		return next(c)
	}
}

// jwt middleware function.
func defaultJwt(next echo.HandlerFunc, errorMsg []byte) echo.HandlerFunc {
	return func(c *echo.Context) error {
		_, parseErr := c.ConnectionProfileSetup()
		if parseErr != nil {
			return api.ErrorStr(c, errorMsg)
		}
		return next(c)
	}
}

// create a group for all /api/v1 functions
func next(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		return next(c)
	}
}

// RegisterServices in echo server, allowed routes
func RegisterServices(e *echo.Echo) *echo.Group {

	// /v1
	groupV1 := e.Group(constants.ApiVersion, next)
	if config.EnableMetrics() {
		logger.Info("registering prometheus_metrics metrics collector endpoint")
		e.GET("/metrics", echo.WrapHandler(
			promhttp.Handler(),
		),
		)
	}
	// /v1/hi
	indexCtl := NewIndexController()
	groupV1.GET("/hi", indexCtl.Index)

	// add authentication controller to root
	// /v1/auth/...
	// register ui rest
	NewUIAuthController().RegisterRouters(groupV1)

	// /v1/...
	indexCtl.RegisterRouters(groupV1)
	NewProfileController().RegisterRouters(groupV1)
	NewSecurityController().RegisterRouters(groupV1)
	NewWalletController().RegisterRouters(groupV1)
	NewSolcController().RegisterRouters(groupV1)

	//register external api calls
	// coin market cap: get eth price data
	NewExternalController().RegisterRouters(groupV1)

	//register web3 networks related services
	web3Group := groupV1.Group("/web3", next)

	// /v1/web3/ropsten
	ropstenGroup := web3Group.Group("/ropsten", next)
	NewRopstenController().RegisterRouters(ropstenGroup)

	// /v1/web3/rinkeby
	rinkebyGroup := web3Group.Group("/rinkeby", next)
	NewRinkebyController().RegisterRouters(rinkebyGroup)

	// /v1/web3/kovan
	kovanGroup := web3Group.Group("/kovan", next)
	NewKovanController().RegisterRouters(kovanGroup)

	// /v1/web3/mainnet
	mainnetGroup := web3Group.Group("/mainnet", next)
	NewMainNetController().RegisterRouters(mainnetGroup)

	// /v1/web3/infura
	infuraGroup := web3Group.Group("/infura", infuraJwt)
	NewInfuraController().RegisterRouters(infuraGroup)

	// /v1/web3/quiknode
	quiknodeGroup := web3Group.Group("/quiknode", quiknodeJwt)
	NewQuikNodeController().RegisterRouters(quiknodeGroup)

	// /v1/web3/private
	privateGroup := web3Group.Group(constants.PrivateApi, privateJwt)
	NewPrivateNetController().RegisterRouters(privateGroup)

	// register controllers related to user context
	// /v1/my
	userGroup := groupV1.Group("/my", userJwt)
	p := project.NewProjectControllerPtr()
	p.RegisterRouters(userGroup)
	registry.NewRegistryController().RegisterRouters(userGroup)

	// register project interaction controller
	// /v1/dapp
	dappGroup := groupV1.Group("/dapp", userJwt)
	project.NewProjectInteractionControllerPtr(p).RegisterRouters(dappGroup)
	return groupV1
}
