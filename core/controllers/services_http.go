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
			return api.ErrorStr(c, []byte("invalid authentication token"))
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
	// add authentication controller to root
	// /v1/auth/...
	// register ui rest
	NewUIAuthController().RegisterRouters(groupV1)

	// /v1/public
	publicGroup := groupV1.Group(constants.PublicApi, next)

	// /v1/public/...
	NewIndexController().RegisterRouters(publicGroup)
	NewProfileController().RegisterRouters(publicGroup)
	NewSecurityController().RegisterRouters(publicGroup)
	NewWalletController().RegisterRouters(publicGroup)
	NewSolcController().RegisterRouters(publicGroup)

	//register external api calls
	// coin market cap: get eth price data
	NewExternalController().RegisterRouters(publicGroup)

	//register public ethereum network related services
	// /v1/ropsten
	ropstenGroup := groupV1.Group("/ropsten", next)
	NewRopstenController().RegisterRouters(ropstenGroup)

	rinkebyGroup := groupV1.Group("/rinkeby", next)
	NewRinkebyController().RegisterRouters(rinkebyGroup)

	kovanGroup := groupV1.Group("/kovan", next)
	NewKovanController().RegisterRouters(kovanGroup)

	mainnetGroup := groupV1.Group("/mainnet", next)
	NewMainNetController().RegisterRouters(mainnetGroup)

	infuraGroup := groupV1.Group("/infura", next)
	infuraGroup.Use(infuraJwt)
	NewInfuraController().RegisterRouters(infuraGroup)

	quiknodeGroup := groupV1.Group("/quiknode", next)
	quiknodeGroup.Use(quiknodeJwt)
	NewQuikNodeController().RegisterRouters(quiknodeGroup)

	privateGroup := groupV1.Group(constants.PrivateApi, next)
	privateGroup.Use(privateJwt)
	NewPrivateNetController().RegisterRouters(privateGroup)

	// register controllers related to user context
	userGroup := groupV1.Group("/my", next)
	userGroup.Use(userJwt)
	project.NewProjectController().RegisterRouters(userGroup)
	registry.NewRegistryController().RegisterRouters(publicGroup)
	return groupV1
}
