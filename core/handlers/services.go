package handlers

import (
	"github.com/labstack/echo"
	"github.com/zerjioang/etherniti/core/api/protocol"
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
			return protocol.ErrorStr(c, "please provide a connection profile token for this kind of call")
		}

		_, parseErr := cc.ConnectionProfileSetup()
		if parseErr != nil {
			return protocol.Error(c, parseErr)
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
func RegisterRoot(e *echo.Echo) {
	e.GET("/", Index)
	e.GET("/v1", Index)
	e.GET("/v1/public", Index)
}

// RegisterServices in echo server, allowed routes
func RegisterServices(e *echo.Echo) *echo.Group {
	group := e.Group("/v1", next)
	logger.Info("registering context free routes")

	publicGroup := group.Group("/public", next)

	NewIndexController().RegisterRouters(publicGroup)
	NewProfileController().RegisterRouters(publicGroup)
	NewSecurityController().RegisterRouters(publicGroup)
	NewWalletController().RegisterRouters(publicGroup)
	NewEthController().RegisterRouters(publicGroup)

	//register public ethereum network related services
	NewRopstenController().RegisterRouters(publicGroup)
	NewRinkebyController().RegisterRouters(publicGroup)
	NewKovanController().RegisterRouters(publicGroup)
	NewMainNetController().RegisterRouters(publicGroup)

	privateGroup := group.Group("/private", next)
	privateGroup.Use(jwt)
	//add jwt middleware to private group
	NewWeb3Controller().RegisterRouters(privateGroup)
	//NewTokenController(deployer.manager).RegisterRouters(privateGroup)
	return group
}
