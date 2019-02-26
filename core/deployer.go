// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package core

import (
	"context"
	"crypto/tls"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/zerjioang/etherniti/core/server/mods/ratelimit"
	"github.com/zerjioang/etherniti/core/util"

	"github.com/zerjioang/etherniti/core/eth/profile"
	"github.com/zerjioang/etherniti/core/handlers"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/release"
	"github.com/zerjioang/etherniti/core/server"

	"github.com/zerjioang/etherniti/core/server/mods/tor"

	"github.com/zerjioang/etherniti/core/eth"

	"github.com/zerjioang/etherniti/core/config"

	"github.com/zerjioang/etherniti/core/api"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

var (
	userAgentErr = errors.New("not authorized. security policy not satisfied")
	gopath       = os.Getenv("GOPATH")
	resources    = gopath + "/src/github.com/zerjioang/etherniti/resources"
	corsConfig   = middleware.CORSConfig{
		AllowOrigins: config.AllowedCorsOriginList,
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			"X-Language",
			config.HttpProfileHeaderkey,
		},
	}
	accessLogFormat = `{"time":"${time_unix}","id":"${id}","ip":"${remote_ip}",` +
		`"host":"${host}","method":"${method}","referer":"${referer}","uri":"${uri}","ua":"${user_agent}",` +
		`"status":${status},"err":"${trycatch}","latency":${latency},"latency_human":"${latency_human}"` +
		`,"in":${bytes_in},"out":${bytes_out}}` + "\n"
	gzipConfig = middleware.GzipConfig{
		Level: 5,
	}
	localhostCert tls.Certificate
	certEtr       error
)

func recoverName() {
	if r := recover(); r != nil {
		logger.Info("recovered from ", r)
	}
}

func init() {
	defer recoverName()
	certBytes := config.GetCertPem()
	keyBytes := config.GetKeyPem()
	if certBytes != nil && len(certBytes) > 0 &&
		keyBytes != nil && len(keyBytes) > 0 {
		localhostCert, certEtr = tls.X509KeyPair(
			certBytes,
			keyBytes,
		)
	} else {
		logger.Error("failed to load SSL crypto data")
	}
}

type Deployer struct {
	manager eth.WalletManager
}

func (deployer Deployer) GetLocalHostTLS() (tls.Certificate, error) {
	return localhostCert, certEtr
}

func (deployer Deployer) Run() {
	logger.Info("loading Etherniti Proxy, an Ethereum Multitenant WebAPI")
	if config.EnableHttpsRedirect {
		//build http server
		httpServerInstance := deployer.newServerInstance()
		// add redirects from http to https
		logger.Info("[LAYER] http to https redirect")
		httpServerInstance.Pre(deployer.httpsRedirect)

		// Start http server
		go func() {
			s, err := deployer.buildInsecureServerConfig(httpServerInstance)
			if err != nil {
				logger.Error("failed to build http server configuration", err)
			} else {
				logger.Info("starting http server...")
				err := httpServerInstance.StartServer(s)
				if err != nil {
					logger.Error("shutting down http the server", err)
				}
			}
		}()
		// Start https server
		secureServer := deployer.newServerInstance()
		go func() {
			s, err := deployer.buildSecureServerConfig(secureServer)
			if err != nil {
				logger.Error("failed to build https server configuration", err)
			} else {
				logger.Info("starting https server...")
				deployer.configureRoutes(secureServer)
				err := secureServer.StartServer(s)
				if err != nil {
					logger.Error("shutting down https the server", err)
				}
			}
		}()
		//graceful shutdown of http and https server
		deployer.shutdown(httpServerInstance, secureServer)
	} else {
		//deploy http server only
		e := deployer.newServerInstance()
		s, err := deployer.buildInsecureServerConfig(e)
		if err != nil {
			logger.Error("failed to build server configuration", err)
		} else {
			deployer.configureRoutes(e)
			// Start server
			go func() {
				logger.Info("starting http server...")
				err := e.StartServer(s)
				if err != nil {
					logger.Info("shutting down http server", err)
				}
			}()
			//graceful shutdown of http server
			deployer.shutdown(e, nil)
		}
	}
}

func (deployer Deployer) shutdown(httpInstance *echo.Echo, httpsInstance *echo.Echo) {
	// The make built-in returns a value of type T (not *T), and it's memory is
	// initialized.
	quit := make(chan os.Signal)

	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	logger.Info("graceful shutdown of the service requested")
	if httpInstance != nil {
		logger.Info("shutting down http server...")
		if err := httpInstance.Shutdown(ctx); err != nil {
			logger.Error(err)
		}
	}
	if httpsInstance != nil {
		logger.Info("shutting down https secure server...")
		if err := httpsInstance.Shutdown(ctx); err != nil {
			logger.Error(err)
		}
	}
	logger.Info("graceful shutdown executed")
	logger.Info("exiting...")
}

func (deployer Deployer) buildSecureServerConfig(e *echo.Echo) (*http.Server, error) {
	cert, err := deployer.GetLocalHostTLS()
	if err != nil {
		log.Fatal("failed to setup TLS configuration due to trycatch", err)
		return nil, err
	}

	//prepare tls configuration
	var tlsConf tls.Config
	tlsConf.Certificates = []tls.Certificate{cert}
	if !e.DisableHTTP2 {
		tlsConf.NextProtos = append(tlsConf.NextProtos, "h2")
	}

	//configure custom secure server
	return &http.Server{
		Addr:         config.HttpsAddress,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		TLSConfig:    &tlsConf,
	}, nil
}

func (deployer Deployer) buildInsecureServerConfig(e *echo.Echo) (*http.Server, error) {
	//configure custom secure server
	return &http.Server{
		Addr:         config.HttpAddress,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}, nil
}

// http to http redirect function
func (deployer Deployer) httpsRedirect(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
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
func (deployer Deployer) hardening(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
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
func (deployer Deployer) fakeServer(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// add fake server header
		h := c.Response().Header()
		h.Set("server", "Apache/2.0.54")
		h.Set("x-powered-by", "PHP/5.1.6")
		return next(c)
	}
}

// bots blacklist function.
func (deployer Deployer) antiBots(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// add antibots policy
		ua := c.Request().UserAgent()
		ua = util.ToLowerAscii(ua)
		if ua == "" {
			//drop the request
			logger.Warn("drop request: no user-agent provided")
			return userAgentErr
		} else if deployer.isBotRequest(ua) {
			//drop the request
			logger.Warn("drop request: provided user-agent is considered as a bot", ua)
			return userAgentErr
		}
		return next(c)
	}
}

// check if user agent string contains bot strings similarities
func (deployer Deployer) isBotRequest(userAgent string) bool {
	var lock = false
	var size = len(api.BadBotsList)
	for i := 0; i < size && !lock; i++ {
		lock = strings.Contains(userAgent, api.BadBotsList[i])
	}
	return lock
}

// check if http request host value is allowed or not
func (deployer Deployer) hostnameCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
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
func (deployer Deployer) keepalive(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// add keep alive headers in the response if requested by the client
		connectionMode := c.Request().Header.Get("Connection")
		connectionMode = util.ToLowerAscii(connectionMode)
		if connectionMode == "keep-alive" {
			// keep alive connection mode requested
			h := c.Response().Header()
			h.Set("Connection", "Keep-Alive")
			h.Set("Keep-Alive", "timeout=5, max=1000")
		}
		return next(c)
	}
}

// inject support for custom context on each request
func (deployer Deployer) injectCustomContext(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &server.EthernitiContext{Context: c}
		return h(cc)
	}
}

// jwt middleware function.
func (deployer Deployer) jwt(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// add fake server header
		h := c.Request().Header
		token := h.Get(config.HttpProfileHeaderkey)
		if token == "" {
			return handlers.ErrorStr(c, "please provide a connection profile token for this kind of call")
		}
		_, parseErr := profile.ParseConnectionProfileToken(token)
		if parseErr != nil {
			return handlers.Error(c, parseErr)
		}
		return next(c)
	}
}

// custom http error handler. returns error messages as json
func (deployer Deployer) customHTTPErrorHandler(err error, c echo.Context) {
	// use code snippet below to customize http return code
	/*
		code := http.StatusInternalServerError
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}
	*/
	_ = handlers.Error(c, err)
}

// build new http server instance
func (deployer Deployer) newServerInstance() *echo.Echo {
	// build a the server
	e := echo.New()
	// enable debug mode
	e.Debug = config.DebugServer
	e.HidePort = config.HideServerDataInConsole
	//hide the banner
	e.HideBanner = config.HideServerDataInConsole
	return e
}

// configure deployer internal configuration
func (deployer Deployer) configureRoutes(e *echo.Echo) {
	// add a custom trycatch handler
	logger.Info("[LAYER] custom trycatch handler")
	e.HTTPErrorHandler = deployer.customHTTPErrorHandler

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
	e.Pre(deployer.antiBots)

	// avoid bots and crawlers checking origin host value
	logger.Info("[LAYER] hostname check")
	e.Pre(deployer.hostnameCheck)

	// add CORS support
	if config.EnableCors {
		logger.Info("[LAYER] cors support")
		e.Use(middleware.CORSWithConfig(corsConfig))
	}

	logger.Info("[LAYER] server http headers hardening")
	// add server api request hardening using http headers
	e.Use(deployer.hardening)

	logger.Info("[LAYER] fake server http header")
	// add fake server header
	e.Use(deployer.fakeServer)

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

	// add custom context handler
	logger.Info("[LAYER] injecting custom context")
	e.Use(deployer.injectCustomContext)

	// avoid panics
	logger.Info("[LAYER] panic recovery")
	e.Use(middleware.Recover())

	// register version 1 api calls
	apiGroup := e.Group("/v1", deployer.apiV1)
	deployer.register(apiGroup)

	logger.Info("[LAYER] / static files")
	//load root static folder
	e.Static("/", resources+"/root")
	e.Static("/phpinfo.php", resources+"/root/phpinfo.php")

	// load swagger ui files
	logger.Info("[LAYER] /swagger files")
	e.Static("/swagger", resources+"/swagger")

	// register root calls
	e.GET("/", handlers.Index)
	e.GET("/v1", handlers.Index)
	e.GET("/v1/public", handlers.Index)

	//configure swagger json from template data
	configureSwaggerJson()
}

// create a group for all /api/v1 functions
func (deployer Deployer) apiV1(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}

// register in echo server, allowed routes
func (deployer Deployer) register(group *echo.Group) *echo.Group {
	logger.Info("registering context free routes")

	publicGroup := group.Group("/public", deployer.apiV1)

	handlers.NewIndexController().RegisterRouters(publicGroup)
	handlers.NewProfileController().RegisterRouters(publicGroup)
	handlers.NewSecurityController().RegisterRouters(publicGroup)
	handlers.NewWalletController().RegisterRouters(publicGroup)
	handlers.NewEthController().RegisterRouters(publicGroup)

	privateGroup := group.Group("/private", deployer.apiV1)
	privateGroup.Use(deployer.jwt)
	//add jwt middleware to private group
	handlers.NewWeb3Controller(deployer.manager).RegisterRouters(privateGroup)
	handlers.NewTokenController(deployer.manager).RegisterRouters(privateGroup)
	return group
}

func configureSwaggerJson() {
	//read template file
	log.Debug("reading swagger json file")
	raw, err := ioutil.ReadFile(resources + "/swagger/swagger-template.json")
	if err != nil {
		logger.Error("failed reading swagger template file", err)
		return
	}
	//replace hardcoded variables
	str := string(raw)
	str = strings.Replace(str, "$title", "Etherniti Proxy REST API", -1)
	str = strings.Replace(str, "$version", release.Version, -1)
	str = strings.Replace(str, "$host", config.SwaggerApiDomain, -1)
	str = strings.Replace(str, "$basepath", "/v1", -1)
	str = strings.Replace(str, "$header-auth-key", config.HttpProfileHeaderkey, -1)
	//write swagger.json file
	writeErr := ioutil.WriteFile(resources+"/swagger/swagger.json", []byte(str), os.ModePerm)
	if writeErr != nil {
		logger.Error("failed writing swagger.json file", writeErr)
		return
	}
}

// create new deployer instance
func NewDeployer() Deployer {
	d := Deployer{}
	d.manager = eth.NewWalletManager()
	return d
}
