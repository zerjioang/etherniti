// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package middleware

import (
	"errors"
	"github.com/zerjioang/etherniti/core/modules/bots"
	"strings"

	middlewareLogger "github.com/zerjioang/etherniti/thirdparty/middleware/logger"

	"github.com/zerjioang/etherniti/core/util/str"

	"github.com/zerjioang/etherniti/core/handlers"
	"github.com/zerjioang/etherniti/core/server/mods/ratelimit"
	"github.com/zerjioang/etherniti/core/server/mods/tor"
	"github.com/zerjioang/etherniti/shared/constants"
	"github.com/zerjioang/etherniti/thirdparty/echo/middleware"

	"github.com/zerjioang/etherniti/core/api"
	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/server"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

var (
	userAgentErr = errors.New("not authorized. security policy not satisfied")
	corsConfig   = middleware.CORSConfig{
		AllowOrigins: config.AllowedCorsOriginList,
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			"X-Language",
			constants.HttpProfileHeaderkey,
		},
	}
	accessLogFormat = `{"time":"${time_unix}","id":"${id}","ip":"${remote_ip}",` +
		`"host":"${host}","method":"${method}","referer":"${referer}","uri":"${uri}","ua":"${user_agent}",` +
		`"status":${status},"err":"${trycatch}","latency":${latency},"latency_human":"${latency_human}"` +
		`,"in":${bytes_in},"out":${bytes_out}}` + "\n"
	gzipConfig = middleware.GzipConfig{
		Level: 5,
	}
)

// custom http error handler. returns error messages as json
func customHTTPErrorHandler(err error, c echo.ContextInterface) {
	// use code snippet below to customize http return code
	/*
		code := http.StatusInternalServerError
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}
	*/
	_ = api.Error(c, err)
}

// http to http redirect function
func HttpsRedirect(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.ContextInterface) error {
		req := c.Request()
		scheme := c.Scheme()
		// host := req.Host
		if scheme == "http" {
			return c.Redirect(301, config.GetRedirectUrl(req.Host, req.RequestURI))
		}
		return next(c)
	}
}

// hardening middleware function.
func hardening(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.ContextInterface) error {
		// add security headers
		h := c.Response().Header()
		h.Set("server", "Apache")
		// h.Set("access-control-allow-credentials", "true")
		h.Set("x-xss-protection", "1; mode=block")
		h.Set("strict-transport-security", "max-age=31536000; includeSubDomains; preload")
		//public-key-pins: pin-sha256="t/OMbKSZLWdYUDmhOyUzS+ptUbrdVgb6Tv2R+EMLxJM="; pin-sha256="PvQGL6PvKOp6Nk3Y9B7npcpeL40twdPwZ4kA2IiixqA="; pin-sha256="ZyZ2XrPkTuoiLk/BR5FseiIV/diN3eWnSewbAIUMcn8="; pin-sha256="0kDINA/6eVxlkns5z2zWv2/vHhxGne/W0Sau/ypt3HY="; pin-sha256="ktYQT9vxVN4834AQmuFcGlSysT1ZJAxg+8N1NkNG/N8="; pin-sha256="rwsQi0+82AErp+MzGE7UliKxbmJ54lR/oPheQFZURy8="; max-age=600; report-uri="https://www.keycdn.com"
		h.Set("X-Content-Type-Options", "nosniff")
		//h.Set("Content-Security-Policy", "default-src 'self' 'unsafe-inline' font-src fonts.googleapis.com fonts.gstatic.com")
		h.Set("Expect-CT", "enforce, max-age=30")
		h.Set("X-UA-Compatible", "IE=Edge,chrome=1")
		h.Set("x-frame-options", "SAMEORIGIN")
		h.Set("Referrer-Policy", "same-origin")
		h.Set("Feature-Policy", "microphone 'none'; payment 'none'; sync-xhr 'self'")
		h.Set("x-firefox-spdy", "h2")
		return next(c)
	}
}

// fake server headers middleware function.
func fakeServer(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.ContextInterface) error {
		// add fake server header
		h := c.Response().Header()
		h.Set("server", "Apache/2.0.54")
		h.Set("x-powered-by", "PHP/5.1.6")
		return next(c)
	}
}

// bots blacklist function.
func antiBots(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.ContextInterface) error {
		// add antibots policy
		ua := c.Request().UserAgent()
		ua = str.ToLowerAscii(ua)
		if ua == "" {
			//drop the request
			logger.Warn("drop request: no user-agent provided")
			return userAgentErr
		} else if isBotRequest(ua) {
			//drop the request
			logger.Warn("drop request: provided user-agent is considered as a bot: ", ua)
			return userAgentErr
		}
		return next(c)
	}
}

// check if user agent string contains bot strings similarities
func isBotRequest(userAgent string) bool {
	return bots.GetBadBotsList().MatchAny(userAgent)
}

// check if http request host value is allowed or not
func hostnameCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.ContextInterface) error {
		// add Host policy
		h := c.Request().Host
		chunks := strings.Split(h, ":")
		var hostname = ""
		if len(chunks) == 1 {
			//no port defined in host header
			hostname = h
		} else if len(chunks) == 2 {
			//port defined in host header
			hostname = chunks[0]
		}
		var allowed = false
		var size = len(config.AllowedHostnames)
		for i := 0; i < size && !allowed; i++ {
			allowed = strings.Compare(hostname, config.AllowedHostnames[i]) == 0
		}
		if allowed {
			// fordward request to next middleware
			return next(c)
		} else {
			// drop the request
			return nil
		}
	}
}

// keepalive middleware function.
func keepalive(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.ContextInterface) error {
		// add keep alive headers in the response if requested by the client
		connectionMode := c.Request().Header.Get("Connection")
		connectionMode = str.ToLowerAscii(connectionMode)
		if connectionMode == "keep-alive" {
			// keep alive connection mode requested
			h := c.Response().Header()
			h.Set("Connection", "Keep-Alive")
			h.Set("Keep-Alive", "timeout=5, max=1000")
		}
		return next(c)
	}
}

// jwt middleware function.
func customContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.ContextInterface) error {
		// convert context in etherniti context
		cc := server.NewEthernitiContext(c)
		return next(cc)
	}
}

// configure deployer internal configuration
func ConfigureServerRoutes(e *echo.Echo) {
	// add a custom trycatch handler
	logger.Info("[LAYER] custom error handler")
	e.HTTPErrorHandler = customHTTPErrorHandler

	// log all single request
	// configure logging level
	logger.Info("[LAYER] logger level")
	if config.EnableLogging {
		e.Logger.SetLevel(config.LogLevel)
		e.Use(middlewareLogger.LoggerWithConfig(middlewareLogger.LoggerConfig{
			Format: accessLogFormat,
		}))
	}

	// custom context
	logger.Info("[LAYER] custom context")
	e.Use(customContext)

	if config.IsHttpMode() {
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
	handlers.RegisterServices(e)

	logger.Info("[LAYER] / static files")
	//load root static folder
	e.Static("/", config.ResourcesDirLanding)

	e.Static("/phpinfo.php", config.ResourcesDirPHP)

	// load swagger ui files
	logger.Info("[LAYER] /swagger files")
	e.Static("/swagger", config.ResourcesDirSwagger)

	// RegisterServices root calls
	RegisterRoot(e)
}

// RegisterServices in echo server, allowed routes
func RegisterRoot(e *echo.Echo) {
	e.GET("/v1", handlers.Index)
	e.GET("/v1/public", handlers.Index)
}

func GetTestSetup() *echo.Echo {
	testServer := echo.New()
	ConfigureServerRoutes(testServer)
	return testServer
}
