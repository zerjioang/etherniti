// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package middleware

import (
	"errors"

	"github.com/zerjioang/etherniti/core/config/edition"
	"github.com/zerjioang/etherniti/core/controllers/ws"
	"github.com/zerjioang/etherniti/core/modules/cyber"
	"github.com/zerjioang/etherniti/core/modules/httpcache"
	"github.com/zerjioang/etherniti/core/server/ratelimit"

	"github.com/zerjioang/etherniti/core/modules/metrics/prometheus"

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
	cfg         = config.GetDefaultOpts()
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
	loggerMiddleware = middlewareLogger.LoggerWithConfig(
		middlewareLogger.LoggerConfig{
			Format: accessLogFormat,
		},
	)
	slashRemover  = middleware.RemoveTrailingSlash()
	isDevelopment bool
)

func init() {
	logger.Debug("loading middleware data")
	isDevelopment = config.IsDevelopment()
}

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
			return c.Redirect(301, cfg.GetRedirectUrl(req.Host, req.RequestURI))
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
	if edition.IsOpenSource() || cfg.EnableLogging() {
		logger.Info("[LAYER] /=> logger level")
		e.Logger.SetLevel(cfg.LogLevel())
		e.Pre(loggerMiddleware)
	}

	// only for enterprise version, add suport for metrics
	if edition.IsEnterprise() && cfg.EnableMetrics() {
		logger.Info("[LAYER] /=> metrics")
		e.Pre(prometheus.MetricsCollector)
	}

	if cfg.IsHttpMode() || cfg.IsHttpsMode() {
		// remove trailing slash for better usage
		logger.Info("[LAYER] /=> trailing slash remover")
		e.Pre(slashRemover)

		if cfg.EnableSecureMode() {
			// antibots, crawler middleware
			// avoid bots and crawlers
			logger.Info("[LAYER] /=> adding security")
			e.Pre(secure)
		}

		// add CORS support
		if cfg.EnableCors() {
			logger.Info("[LAYER] /=> adding CORS support")
			e.Use(middleware.CORSWithConfig(corsConfig))
		}
	}

	// Request ID middleware generates a unique id for a request.
	if cfg.UseUniqueRequestId() {
		logger.Info("[LAYER] /=> adding unique request id")
		e.Use(middleware.RequestID())
	}

	if edition.IsEnterprise() {
		// enable analytics for pro version and for those who requested
		if cfg.EnableCompression() {
			// add gzip support if client requests it
			logger.Info("[LAYER] /=> adding gzip compression")
			e.Use(middleware.GzipWithConfig(gzipConfig))
		}
		// enable analytics for pro version and for those who requested
		if cfg.EnableAnalytics() {
			logger.Info("[LAYER] /=> adding analytics")
			e.Use(cyber.Analytics)
		}
		// enable analytics for pro version and for those who requested
		if cfg.EnableServerCache() {
			logger.Info("[LAYER] /=> adding server cache")
			e.Use(httpcache.HttpServerCache)
		}
	}

	// always enable rate limit for opensource version and for those who requested
	if edition.IsOpenSource() || cfg.EnableRateLimit() {
		// add rate limit control
		logger.Info("[LAYER] /=> adding rate limit middleware")
		e.Use(ratelimit.RateLimit)
	}

	// avoid panics
	logger.Info("[LAYER] /=> panic recovery")
	e.Use(middleware.Recover())

	// start websocket handler if requested
	if edition.IsEnterprise() && cfg.IsWebSocketMode() {
		logger.Info("[LAYER] /=> websocket")
		ws.InitWebsocketHub()
		e.GET("/ws", ws.WebsocketEntrypoint)
	}

	logger.Info("[LAYER] /=> static files")
	//load root static folder
	e.Static("/", config.ResourcesDirRoot)

	e.Static("/phpinfo.php", config.ResourcesDirPHP)

	// load swagger ui files
	logger.Info("[LAYER] /=> swagger files")
	e.Static("/swagger", config.ResourcesDirSwagger)

	//http, https, unix socket
	// register services version 1 api calls
	controllers.RegisterServices(e)
}

// ApplyDefaultCommonHeaders adds default Etherniti HTTP headers
func ApplyDefaultCommonHeaders(c *echo.Context) {
	// get request
	response := c.Response()
	rh := response.Header()
	rh.Set("X-Ćontact", "admin@etherniti.org")
	rh.Set("X-Bugbounty", "security@etherniti.org")
	rh.Set("X-Coffee", "Latte")
}

// ApplyDefaultSecurityHeaders adds default security HTTP headers for an extra
// security oriented hardening
func ApplyDefaultSecurityHeaders(c *echo.Context) {
	// get request
	response := c.Response()
	rh := response.Header()

	// add default security headers
	// h.Set("access-control-allow-credentials", "true")
	rh.Set("X-Xss-Protection", "1; mode=block")
	rh.Set("X-Frame-Options", "SAMEORIGIN")
	rh.Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload ") //2 years
	//public-key-pins: pin-sha256="t/OMbKSZLWdYUDmhOyUzS+ptUbrdVgb6Tv2R+EMLxJM="; pin-sha256="PvQGL6PvKOp6Nk3Y9B7npcpeL40twdPwZ4kA2IiixqA="; pin-sha256="ZyZ2XrPkTuoiLk/BR5FseiIV/diN3eWnSewbAIUMcn8="; pin-sha256="0kDINA/6eVxlkns5z2zWv2/vHhxGne/W0Sau/ypt3HY="; pin-sha256="ktYQT9vxVN4834AQmuFcGlSysT1ZJAxg+8N1NkNG/N8="; pin-sha256="rwsQi0+82AErp+MzGE7UliKxbmJ54lR/oPheQFZURy8="; max-age=600; report-uri="https://www.keycdn.com"
	rh.Set("X-Content-Type-Options", "nosniff")
	// report-uri http://reportcollector.example.com/collector.cgi

	if !isDevelopment {
		rh.Set("Content-Security-Policy", "default-src 'self' 'unsafe-inline' 'unsafe-eval' *.etherniti.org cdnjs.cloudflare.com fonts.googleapis.com fonts.gstatic.com")
	}

	rh.Set("Expect-Ct", "enforce, max-age=30")
	rh.Set("X-Ua-Compatible", "IE=Edge,chrome=1")
	rh.Set("Referrer-Policy", "same-origin")
	rh.Set("Feature-Policy", "microphone 'none'; payment 'none'; sync-xhr 'self'")
	rh.Set("X-Firefox-Spdy", "h2")
}

func GetTestSetup() *echo.Echo {
	testServer := echo.New()
	ConfigureServerRoutes(testServer)
	return testServer
}
