// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package middleware

import (
	"errors"
	"strings"

	"github.com/zerjioang/etherniti/core/modules/metrics/prometheus_metrics"

	"github.com/zerjioang/etherniti/core/data"
	"github.com/zerjioang/etherniti/core/modules/tor"

	"github.com/zerjioang/etherniti/core/modules/badips"
	"github.com/zerjioang/etherniti/core/modules/cyber"
	"github.com/zerjioang/etherniti/core/server/mods/ratelimit"

	"github.com/zerjioang/etherniti/core/modules/bots"
	middlewareLogger "github.com/zerjioang/etherniti/thirdparty/middleware/logger"

	"github.com/zerjioang/etherniti/core/util/str"

	"github.com/zerjioang/etherniti/core/handlers"
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
		req := c.Request()
		scheme := c.Scheme()
		// host := req.Host
		if scheme == "http" {
			return c.Redirect(301, config.GetRedirectUrl(req.Host, req.RequestURI))
		}
		return next(c)
	}
}

func secure(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		// add abuseIP policy
		ip := c.RealIP()
		if ip == "" {
			//drop the request
			logger.Warn("drop request: no IP provided")
			return securityErr
		} else if badips.IsBackListedIp(ip) {
			//drop the request
			logger.Warn("drop request: blacklisted IP detected: ", ip)
			return securityErr
		}

		request := c.Request()

		// add antibots policy
		ua := request.UserAgent()
		ua = str.ToLowerAscii(ua)
		if ua == "" {
			//drop the request
			logger.Warn("drop request: no user-agent provided")
			return securityErr
		} else if len(ua) < 4 || bots.GetBadBotsList().MatchAny(ua) {
			//drop the request
			logger.Warn("drop request: provided user-agent is considered as a bot: ", ua)
			return securityErr
		}

		// add hostname policy
		host := request.Host
		chunks := strings.Split(host, ":")
		var hostname = ""
		if len(chunks) == 1 {
			//no port defined in host header
			hostname = host
		} else if len(chunks) == 2 {
			//port defined in host header
			hostname = chunks[0]
		}
		allowed := config.AllowedHostnames.Contains(hostname)
		if !allowed {
			// drop the request
			logger.Warn("drop request: provided request does not specifies a valid host name in http headers")
			return securityErr
		}

		if config.BlockTorConnections() {
			// add rate limit control
			logger.Info("[LAYER] tor connections blocker middleware added")
			//get current request ip
			requestIp := request.RemoteAddr
			found := tor.TornodeSet.Contains(requestIp)
			if !found {
				//received request IP is not blacklisted
				return next(c)
			} else {
				// received request is done using on of the blacklisted tor nodes
				//return rate limit excedeed message
				logger.Warn("drop request: provided request is done using on of the blacklisted tor nodes")
				return c.FastBlob(200, echo.MIMEApplicationJSON, data.ErrBlockTorConnection)
			}
		}

		// add keep alive headers in the response if requested by the client
		h := request.Header
		connectionMode := h.Get("Connection")
		connectionMode = str.ToLowerAscii(connectionMode)
		/*
			Lista de parámetros separados por coma,
			cada uno consiste en un identificador y un valor separado por el signo igual ('=').
			Es posible establecer los siguientes identificadores:
			* timeout: indica la cantidad de  tiempo mínima  en la cual una conexión ociosa
			se debe mantener abierta (en segundos).
			Nótese que los timeouts mas largos que el timeout de TCP
			pueden ser ignorados si no se establece un mensaje de TCP
			keep-alive  en la capa de transporte.
			* max: indica el número máximo de peticiones que pueden ser
			enviadas en esta conexión antes de que sea cerrada. Si es  0,
			este valor es ignorado para las conexiones no segmentadas,
			ya que se enviara otra solicitud en la próxima respuesta.
			Una canalización de HTTP puede ser usada para limitar la división.
		*/
		response := c.Response()
		rh := response.Header()
		if strings.Contains(connectionMode, "keep-alive") {
			// keep alive connection mode requested
			rh.Set("Connection", "Keep-Alive")
			rh.Set("Keep-Alive", "timeout=5, max=1000")
		}

		// add security headers
		rh.Set("server", "Apache")
		// h.Set("access-control-allow-credentials", "true")
		rh.Set("x-xss-protection", "1; mode=block")
		rh.Set("strict-transport-security", "max-age=63072000; includeSubDomains; preload ") //2 years
		//public-key-pins: pin-sha256="t/OMbKSZLWdYUDmhOyUzS+ptUbrdVgb6Tv2R+EMLxJM="; pin-sha256="PvQGL6PvKOp6Nk3Y9B7npcpeL40twdPwZ4kA2IiixqA="; pin-sha256="ZyZ2XrPkTuoiLk/BR5FseiIV/diN3eWnSewbAIUMcn8="; pin-sha256="0kDINA/6eVxlkns5z2zWv2/vHhxGne/W0Sau/ypt3HY="; pin-sha256="ktYQT9vxVN4834AQmuFcGlSysT1ZJAxg+8N1NkNG/N8="; pin-sha256="rwsQi0+82AErp+MzGE7UliKxbmJ54lR/oPheQFZURy8="; max-age=600; report-uri="https://www.keycdn.com"
		rh.Set("X-Content-Type-Options", "nosniff")
		// report-uri http://reportcollector.example.com/collector.cgi
		if !config.IsDevelopment() {
			rh.Set("Content-Security-Policy", "default-src 'self' 'unsafe-inline' 'unsafe-eval' *.etherniti.org cdnjs.cloudflare.com fonts.googleapis.com fonts.gstatic.com")
		}
		rh.Set("Expect-CT", "enforce, max-age=30")
		rh.Set("X-UA-Compatible", "IE=Edge,chrome=1")
		rh.Set("x-frame-options", "SAMEORIGIN")
		rh.Set("Referrer-Policy", "same-origin")
		rh.Set("Feature-Policy", "microphone 'none'; payment 'none'; sync-xhr 'self'")
		rh.Set("x-firefox-spdy", "h2")

		// add fake server header
		rh.Set("server", "Apache/2.0.54")
		rh.Set("x-powered-by", "PHP/5.1.6")

		return next(c)
	}
}

// configure deployer internal configuration
func ConfigureServerRoutes(e *echo.Echo) {
	// add a custom trycatch handler
	logger.Info("[LAYER] custom error handler")
	e.HTTPErrorHandler = customHTTPErrorHandler

	// log all single request
	// configure logging level
	if config.EnableLogging() {
		logger.Info("[LAYER] logger level")
		e.Logger.SetLevel(config.LogLevel())
		e.Pre(middlewareLogger.LoggerWithConfig(middlewareLogger.LoggerConfig{
			Format: accessLogFormat,
		}))
	}

	if config.EnableMetrics() {
		logger.Info("[LAYER] metrics")
		e.Pre(prometheus_metrics.MetricsCollector)
	}

	if config.IsHttpMode() {
		// remove trailing slash for better usage
		logger.Info("[LAYER] trailing slash remover")
		e.Pre(middleware.RemoveTrailingSlash())

		if config.EnableSecureMode() {
			// antibots, crawler middleware
			// avoid bots and crawlers
			logger.Info("[LAYER] security")
			e.Pre(secure)
		}

		// add CORS support
		if config.EnableCors() {
			logger.Info("[LAYER] cors support")
			e.Use(middleware.CORSWithConfig(corsConfig))
		}
	}

	if config.EnableRateLimit() {
		// add rate limit control
		logger.Info("[LAYER] rest api rate limit middleware added")
		e.Use(ratelimit.RateLimit)
	}

	if config.EnableAnalytics() {
		logger.Info("[LAYER] analytics")
		e.Use(cyber.Analytics)
	}

	// Request ID middleware generates a unique id for a request.
	if config.UseUniqueRequestId() {
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
	e.Static("/", config.ResourcesDirRoot)

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
