// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package core

import (
	"context"
	"crypto/tls"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/zerjioang/etherniti/core/server/mods/ratelimit"
	"github.com/zerjioang/etherniti/core/server/mods/tor"

	"github.com/zerjioang/etherniti/core/eth"

	"github.com/zerjioang/etherniti/core/config"

	"github.com/zerjioang/etherniti/core/api"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/zerjioang/etherniti/core/handlers"
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
	gzipConfig = middleware.GzipConfig{
		Level: 5,
	}
	localhostCert tls.Certificate
	certEtr       error
)

func recoverName() {
	if r := recover(); r!= nil {
		log.Info("recovered from ", r)
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
		log.Error("failed to load SSL crypto data")
	}
}

type Deployer struct {
	manager eth.WalletManager
}

func (deployer Deployer) GetLocalHostTLS() (tls.Certificate, error) {
	return localhostCert, certEtr
}

func (deployer Deployer) Run() {
	log.Info("loading Ethereum Multitenant Webapi (etherniti)")

	if config.EnableHttpsRedirect {
		//build http server
		httpServerInstance := deployer.newServerInstance()
		// add redirects from http to https
		log.Info("[LAYER] http to https redirect")
		httpServerInstance.Pre(deployer.httpsRedirect)

		// Start http server
		go func() {
			s, err := deployer.buildInsecureServerConfig(httpServerInstance)
			if err != nil {
				log.Error("failed to build http server configuration", err)
			} else {
				log.Info("starting http server...")
				err := httpServerInstance.StartServer(s)
				if err != nil {
					log.Error("shutting down http the server", err)
				}
			}
		}()
		// Start https server
		httpsServerInstance := deployer.newServerInstance()
		go func() {
			s, err := deployer.buildSecureServerConfig(httpServerInstance)
			if err != nil {
				log.Error("failed to build https server configuration", err)
			} else {
				log.Info("starting https server...")
				err := httpServerInstance.StartServer(s)
				if err != nil {
					log.Error("shutting down https the server", err)
				}
			}
		}()
		//graceful shutdown of http and https server
		deployer.shutdown(httpServerInstance, httpsServerInstance)
	} else {
		//deploy http server only
		e := deployer.newServerInstance()
		s, err := deployer.buildInsecureServerConfig(e)
		if err != nil {
			log.Error("failed to build server configuration", err)
		} else {
			deployer.configureRoutes(e)
			// Start server
			go func() {
				log.Info("starting http server...")
				err := e.StartServer(s)
				if err != nil {
					e.Logger.Info("shutting down http server", err)
				}
			}()
			//graceful shutdown of http server
			deployer.shutdown(e, nil)
		}
	}
}

func (deployer Deployer) shutdown(httpInstance *echo.Echo, httpsInstance *echo.Echo) {
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log.Info("graceful shutdown of the service requested")
	if httpInstance != nil {
		log.Info("shutting down http server...")
		if err := httpInstance.Shutdown(ctx); err != nil {
			log.Error(err)
		}
	}
	if httpsInstance != nil {
		log.Info("shutting down https secure server...")
		if err := httpsInstance.Shutdown(ctx); err != nil {
			log.Error(err)
		}
	}
	log.Info("graceful shutdown executed")
	log.Info("exiting...")
}

func (deployer Deployer) buildSecureServerConfig(e *echo.Echo) (*http.Server, error) {
	cert, err := deployer.GetLocalHostTLS()
	if err != nil {
		log.Fatal("failed to setup TLS configuration due to error", err)
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
		h.Set("access-control-allow-credentials", "true")
		h.Set("x-xss-protection", "1; mode=block")
		h.Set("strict-transport-security", "max-age=31536000; includeSubDomains; preload")
		//public-key-pins: pin-sha256="t/OMbKSZLWdYUDmhOyUzS+ptUbrdVgb6Tv2R+EMLxJM="; pin-sha256="PvQGL6PvKOp6Nk3Y9B7npcpeL40twdPwZ4kA2IiixqA="; pin-sha256="ZyZ2XrPkTuoiLk/BR5FseiIV/diN3eWnSewbAIUMcn8="; pin-sha256="0kDINA/6eVxlkns5z2zWv2/vHhxGne/W0Sau/ypt3HY="; pin-sha256="ktYQT9vxVN4834AQmuFcGlSysT1ZJAxg+8N1NkNG/N8="; pin-sha256="rwsQi0+82AErp+MzGE7UliKxbmJ54lR/oPheQFZURy8="; max-age=600; report-uri="https://www.keycdn.com"
		h.Set("X-Content-Type-Options", "nosniff")
		h.Set("Content-Security-Policy", "default-src 'self' 'unsafe-inline' font-src fonts.googleapis.com fonts.gstatic.com")
		h.Set("Expect-CT", "enforce, max-age=30")
		h.Set("X-UA-Compatible", "IE=Edge,chrome=1")
		h.Set("x-frame-options", "SAMEORIGIN")
		h.Set("Referrer-Policy", "same-origin")
		h.Set("Feature-Policy", "microphone 'none'; payment 'none'; sync-xhr 'self'")
		h.Set("x-firefox-spdy", "h2")
		return next(c)
	}
}

// fakeServer middleware function.
func (deployer Deployer) fakeServer(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// add fake server header
		h := c.Response().Header()
		h.Set("server", "Apache/2.0.54")
		h.Set("x-powered-by", "PHP/5.1.6")
		return next(c)
	}
}

// fakeServer antiBots function.
func (deployer Deployer) antiBots(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// add antibots policy
		ua := c.Request().UserAgent()
		ua = strings.ToLower(ua)
		if ua == "" || deployer.isBotRequest(ua) {
			//drop the request
			log.Warn("drop request: User-Agent =", ua)
			return userAgentErr
		}
		return next(c)
	}
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

// check if user agent string contains bot strings similarities
func (deployer Deployer) isBotRequest(userAgent string) bool {
	var lock = false
	var size = len(api.BadBotsList)
	for i := 0; i < size && !lock; i++ {
		lock = strings.Contains(userAgent, api.BadBotsList[i])
	}
	return lock
}

// keepalive middleware function.
func (deployer Deployer) keepalive(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// add keep alive headers
		h := c.Response().Header()
		h.Set("Connection", "Keep-Alive")
		h.Set("Keep-Alive", "timeout=5, max=1000")
		return next(c)
	}
}

// custom http error handler
func (deployer Deployer) customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	// log error always? less performance in production
	c.Logger().Error(err)

	apiErr := api.NewApiError(code, err.Error())
	_ = c.JSON(code, apiErr)
}

// register in echo server, allowed routes
func (deployer Deployer) register(server *echo.Echo) *echo.Echo {
	log.Info("registering routes")
	handlers.NewIndexController().RegisterRouters(server)
	handlers.NewProfileController().RegisterRouters(server)
	handlers.NewTransactionController().RegisterRouters(server)
	handlers.NewEthController(deployer.manager).RegisterRouters(server)
	handlers.NewTokenController(deployer.manager).RegisterRouters(server)
	return server
}

// build new http server instance
func (deployer Deployer) newServerInstance() *echo.Echo {
	// build a the server
	e := echo.New()
	// enable debug mode
	e.Debug = false
	//hide the banner
	e.HideBanner = true
	return e
}

// configure deployer internal configuration
func (deployer Deployer) configureRoutes(e *echo.Echo) {
	// add a custom error handler
	log.Info("[LAYER] custom error handler")
	e.HTTPErrorHandler = deployer.customHTTPErrorHandler

	// log all single request
	// configure logging level
	log.Info("[LAYER] logger at warn level")
	if config.EnableLogging {
		e.Logger.SetLevel(config.LogLevel)
		e.Use(middleware.Logger())
	}

	// antibots, crawler middleware
	// avoid bots and crawlers
	log.Info("[LAYER] antibots")
	e.Pre(deployer.antiBots)

	// avoid bots and crawlers
	log.Info("[LAYER] hostname check")
	e.Pre(deployer.hostnameCheck)

	// remove trailing slash for better usage
	log.Info("[LAYER] trailing slash remover")
	e.Pre(middleware.RemoveTrailingSlash())

	// add CORS support
	log.Info("[LAYER] cors support")
	e.Use(middleware.CORSWithConfig(corsConfig))

	log.Info("[LAYER] server http headers hardening")
	// add server api request hardening using http headers
	e.Use(deployer.hardening)

	log.Info("[LAYER] fake server http header")
	// add fake server header
	e.Use(deployer.fakeServer)

	if config.EnableRateLimit {
		// add rate limit control
		log.Info("[LAYER] rest api rate limit middleware added")
		e.Use(ratelimit.RateLimit)
	}

	if config.BlockTorConnections {
		// add rate limit control
		log.Info("[LAYER] tor connections blocker middleware added")
		e.Use(tor.BlockTorConnections)
	}

	log.Info("[LAYER] unique request id")
	// Request ID middleware generates a unique id for a request.
	if config.UseUniqueRequestId {
		e.Use(middleware.RequestID())
	}

	// add gzip support if client requests it
	log.Info("[LAYER] gzip compression")
	e.Use(middleware.GzipWithConfig(gzipConfig))

	// avoid panics
	log.Info("[LAYER] panic recovery")
	e.Use(middleware.Recover())

	log.Info("[LAYER] / static files")
	//load root static folder
	e.Static("/", resources+"/root")
	e.Static("/phpinfo.php", resources+"/root/phpinfo.php")

	// load swagger ui files
	log.Info("[LAYER] /swagger files")
	e.Static("/swagger", resources+"/swagger")

	deployer.register(e)
}

// create new deployer instance
func NewDeployer() Deployer {
	d := Deployer{}
	d.manager = eth.NewWalletManager()
	return d
}
