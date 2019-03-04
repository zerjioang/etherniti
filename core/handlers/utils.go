// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/middleware"
	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/server/mods/ratelimit"
	"github.com/zerjioang/etherniti/core/server/mods/tor"

	"github.com/labstack/echo"
	"github.com/zerjioang/etherniti/core/server"
)

// for testing
func MockServer() *echo.Echo {
	// Setup
	testServer := echo.New()
	ConfigureRoutes(testServer)
	return testServer
}

func NewContext(e *echo.Echo) echo.Context {
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return server.NewEthernitiContext(c)
}

// configure deployer internal configuration
func ConfigureRoutes(e *echo.Echo) {
	// add a custom trycatch handler
	logger.Info("[LAYER] custom trycatch handler")
	e.HTTPErrorHandler = customHTTPErrorHandler

	// log all single request
	// configure logging level
	logger.Info("[LAYER] logger at warn level")
	if config.EnableLogging {
		e.Logger.SetLevel(config.LogLevel)
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: accessLogFormat,
		}))
	}

	// remove trailing slash for better usage
	logger.Info("[LAYER] trailing slash remover")
	e.Pre(middleware.RemoveTrailingSlash())

	// antibots, crawler middleware
	// avoid bots and crawlers
	logger.Info("[LAYER] antibots")
	e.Pre(antiBots)

	// avoid bots and crawlers checking origin host value
	logger.Info("[LAYER] hostname check")
	e.Pre(hostnameCheck)

	// custom context
	logger.Info("[LAYER] custom context")
	e.Use(customContext)

	// add CORS support
	if config.EnableCors {
		logger.Info("[LAYER] cors support")
		e.Use(middleware.CORSWithConfig(corsConfig))
	}

	logger.Info("[LAYER] server http headers hardening")
	// add server api request hardening using http headers
	e.Use(hardening)

	logger.Info("[LAYER] fake server http header")
	// add fake server header
	e.Use(fakeServer)

	if config.BlockTorConnections {
		// add rate limit control
		logger.Info("[LAYER] tor connections blocker middleware added")
		e.Use(tor.BlockTorConnections)
	}

	if config.EnableRateLimit {
		// add rate limit control
		logger.Info("[LAYER] rest api rate limit middleware added")
		e.Use(ratelimit.RateLimit)
	}

	// Request ID middleware generates a unique id for a request.
	if config.UseUniqueRequestId {
		logger.Info("[LAYER] unique request id")
		e.Use(middleware.RequestID())
	}

	// add gzip support if client requests it
	logger.Info("[LAYER] gzip compression")
	e.Use(middleware.GzipWithConfig(gzipConfig))

	// avoid panics
	logger.Info("[LAYER] panic recovery")
	e.Use(middleware.Recover())

	// RegisterServices version 1 api calls
	RegisterServices(e)

	logger.Info("[LAYER] / static files")
	//load root static folder
	e.Static("/", resources+"/root")
	e.Static("/phpinfo.php", resources+"/root/phpinfo.php")

	// load swagger ui files
	logger.Info("[LAYER] /swagger files")
	e.Static("/swagger", resources+"/swagger")

	// RegisterServices root calls
	RegisterRoot(e)

	//configure swagger json from template data
	configureSwaggerJson()
}

func GetTestSetup() *echo.Echo {
	testServer := echo.New()
	ConfigureRoutes(testServer)
	return testServer
}
