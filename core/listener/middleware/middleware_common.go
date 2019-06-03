// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package middleware

import (
	"errors"

	"github.com/zerjioang/etherniti/core/config/edition"
	"github.com/zerjioang/etherniti/core/controllers/ws"

	"github.com/zerjioang/etherniti/core/modules/metrics/prometheus_metrics"

	"github.com/zerjioang/etherniti/core/modules/cyber"
	"github.com/zerjioang/etherniti/core/server/ratelimit"

	middlewareLogger "github.com/zerjioang/etherniti/thirdparty/middleware/logger"

	"github.com/zerjioang/etherniti/core/controllers"
	"github.com/zerjioang/etherniti/shared/constants"
	"github.com/zerjioang/etherniti/thirdparty/echo/middleware"

	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

var (
	securityErr = errors.New("not authorized. security policy not satisfied")
	corsConfig  = middleware.CORSConfig{
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			"X-Language",
			"Authorization",
			constants.HttpProfileHeaderkey,
		},
	}
	accessLogFormat = `{"time":"${time_unix}","id":"${id}","ip":"${remote_ip}",` +
		`"host":"${host}","method":"${method}","referer":"${referer}","uri":"${uri}","ua":"${user_agent}",` +
		`"status":${status},"err":"${stack}","latency":${latency},"latency_human":"${latency_human}"` +
		`,"in":${bytes_in},"out":${bytes_out}}` + "\n"
	gzipConfig = middleware.GzipConfig{
		Level: 5,
	}
)

// custom http error handler. returns error messages as json
func customHTTPErrorHandler(err error, c *echo.Context) {
	// use code snippet below to customize http return code
	/*
		code := protocol.StatusInternalServerError
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}
	*/
	_ = api.Error(c, err)
}

// http to http redirect function
func HttpsRedirect(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		if c.IsHttp() {
			req := c.Request()
			return c.Redirect(301, config.GetRedirectUrl(req.Host, req.RequestURI))
		}
		return next(c)
	}
}

// configure deployer internal configuration
func ConfigureServerRoutes(e *echo.Echo) {
	// add a custom stack handler
	logger.Info("[LAYER] /=> custom error handler")
	e.HTTPErrorHandler = customHTTPErrorHandler

	// log all single request
	// configure logging level
	// always enable logging for opensource and for those who requested
	if edition.IsOpenSource() || config.EnableLogging() {
		logger.Info("[LAYER] /=> logger level")
		e.Logger.SetLevel(config.LogLevel())
		e.Pre(middlewareLogger.LoggerWithConfig(middlewareLogger.LoggerConfig{
			Format: accessLogFormat,
		}))
	}

	// only for enterprise version, add suport for metrics
	if edition.IsEnterprise() && config.EnableMetrics() {
		logger.Info("[LAYER] /=> metrics")
		e.Pre(prometheus_metrics.MetricsCollector)
	}

	if config.IsHttpMode() || config.IsHttpsMode() {
		// remove trailing slash for better usage
		logger.Info("[LAYER] /=> trailing slash remover")
		e.Pre(middleware.RemoveTrailingSlash())

		if config.EnableSecureMode() {
			// antibots, crawler middleware
			// avoid bots and crawlers
			logger.Info("[LAYER] /=> security")
			e.Pre(secure)
		}

		// add CORS support
		if config.EnableCors() {
			logger.Info("[LAYER] /=> CORS support")
			e.Use(middleware.CORSWithConfig(corsConfig))
		}
	}

	// always enable rate limit for opensource version and for those who requested
	if edition.IsOpenSource() || config.EnableRateLimit() {
		// add rate limit control
		logger.Info("[LAYER] /=> rest api rate limit middleware added")
		e.Use(ratelimit.RateLimit)
	}

	// enable analytics for pro version and for those who requested
	if edition.IsEnterprise() && config.EnableAnalytics() {
		logger.Info("[LAYER] /=> analytics")
		e.Use(cyber.Analytics)
	}

	// Request ID middleware generates a unique id for a request.
	if config.UseUniqueRequestId() {
		logger.Info("[LAYER] /=> unique request id")
		e.Use(middleware.RequestID())
	}

	// enable analytics for pro version and for those who requested
	if edition.IsEnterprise() {
		// add gzip support if client requests it
		logger.Info("[LAYER] /=> gzip compression")
		e.Use(middleware.GzipWithConfig(gzipConfig))
	}

	// avoid panics
	logger.Info("[LAYER] /=> panic recovery")
	e.Use(middleware.Recover())

	//http, https, unix socket
	// register services version 1 api calls
	controllers.RegisterServices(e)

	// start websocket handler if requested
	if edition.IsEnterprise() && config.IsWebSocketMode() {
		logger.Info("[LAYER] /=> websocket")
		e.GET("/ws", ws.WebsocketEntrypoint)
	}

	logger.Info("[LAYER] /=> static files")
	//load root static folder
	e.Static("/", config.ResourcesDirRoot)

	e.Static("/phpinfo.php", config.ResourcesDirPHP)

	// load swagger ui files
	logger.Info("[LAYER] /=> swagger files")
	e.Static("/swagger", config.ResourcesDirSwagger)

	// RegisterServices root calls
	RegisterRoot(e)
}

// RegisterServices in echo server, allowed routes
func RegisterRoot(e *echo.Echo) {
	e.GET("/v1", controllers.Index)
	e.GET("/v1/hi", controllers.Index)
}

func GetTestSetup() *echo.Echo {
	testServer := echo.New()
	ConfigureServerRoutes(testServer)
	return testServer
}
