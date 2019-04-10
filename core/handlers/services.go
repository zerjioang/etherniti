// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/server"
	"github.com/zerjioang/etherniti/shared/constants"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

const (
	infuraJwtErrorMessage   = "please provide an Infura connection profile token including provided Infura endpoint URL (https://$NETWORK.infura.io/v3/$PROJECT_ID) for this kind of call."
	quiknodeJwtErrorMessage = "please provide a QuikNode connection profile token including provided full peer endpoint URL"
	jwtErrorMessage         = "please provide a connection profile token for this kind of call"
)

func infuraJwt(next echo.HandlerFunc) echo.HandlerFunc {
	return jwt(next, infuraJwtErrorMessage)
}

func quiknodeJwt(next echo.HandlerFunc) echo.HandlerFunc {
	return jwt(next, quiknodeJwtErrorMessage)
}

func privateJwt(next echo.HandlerFunc) echo.HandlerFunc {
	return jwt(next, jwtErrorMessage)
}

// jwt middleware function.
func jwt(next echo.HandlerFunc, errorMsg string) echo.HandlerFunc {
	return func(c echo.Context) error {
		// convert context in etherniti context
		cc := c.(*server.EthernitiContext)
		token := cc.ReadConnectionProfileToken()
		if token == "" {
			return api.ErrorStr(c, errorMsg)
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
	// /v1
	groupV1 := e.Group(constants.ApiVersion, next)
	logger.Info("registering context free routes")
	// /v1/
	publicGroup := groupV1.Group(constants.PublicApi, next)

	// /v1/...
	NewIndexController().RegisterRouters(publicGroup)
	NewProfileController().RegisterRouters(publicGroup)
	NewSecurityController().RegisterRouters(publicGroup)
	NewWalletController().RegisterRouters(publicGroup)
	NewSolcController().RegisterRouters(publicGroup)
	NewContractNameSpaceController().RegisterRouters(publicGroup)

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
	//NewTokenController(deployer.manager).RegisterRouters(privateGroup)
	return groupV1
}
