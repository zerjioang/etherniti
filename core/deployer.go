// Copyright gaethway
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

	"github.com/zerjioang/gaethway/core/eth"

	"github.com/zerjioang/gaethway/core/config"

	"github.com/zerjioang/gaethway/core/api"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/zerjioang/gaethway/core/handlers"
)

var (
	userAgentErr = errors.New("not authorized. security policy not satisfied")
	gopath       = os.Getenv("GOPATH")
	resources    = gopath + "/src/github.com/zerjioang/gaethway/resources"
	corsConfig   = middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}
	gzipConfig = middleware.GzipConfig{
		Level: 5,
	}
	localhostCert, certEtr = tls.X509KeyPair(
		[]byte(config.CertPem),
		[]byte(config.KeyPem),
	)
)

type Deployer struct {
	manager eth.WalletManager
}

func (deployer Deployer) GetLocalHostTLS() (tls.Certificate, error) {
	return localhostCert, certEtr
}

func (deployer Deployer) Run() {
	log.Info("loading Ethereum Multitenant Webapi (gaethway)")

	httpServerInstance := echo.New()
	httpServerInstance.HideBanner = true
	// add redirects from http to https
	log.Info("[LAYER] http to https redirect")
	httpServerInstance.Pre(deployer.httpsRedirect)

	// Start http server
	go func() {
		log.Info("starting http server...")
		err := httpServerInstance.Start(config.HttpAddress)
		if err != nil {
			log.Error("shutting down http the server")
		}
	}()

	// build a secure http server
	e := echo.New()

	// enable debug mode
	e.Debug = true

	cert, err := deployer.GetLocalHostTLS()
	if err != nil {
		log.Fatal("failed to setup TLS configuration due to error", err)
		return
	}

	//prepare tls configuration
	var tlsConf tls.Config
	tlsConf.Certificates = []tls.Certificate{cert}
	if !e.DisableHTTP2 {
		tlsConf.NextProtos = append(tlsConf.NextProtos, "h2")
	}

	//configure custom secure server
	s := &http.Server{
		Addr:         config.HttpsAddress,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		TLSConfig:    &tlsConf,
	}
	//hide the banner
	e.HideBanner = true

	// add a custom error handler
	log.Info("[LAYER] custom error handler")
	e.HTTPErrorHandler = deployer.customHTTPErrorHandler

	// antibots, crawler middleware
	// avoid bots and crawlers
	log.Info("[LAYER] antibots")
	e.Pre(deployer.antiBots)

	// remove trailing slash for better usage
	log.Info("[LAYER] trailing slash remover")
	e.Pre(middleware.RemoveTrailingSlash())

	// log all single request
	// configure logging level
	log.Info("[LAYER] logger at warn level")
	e.Logger.SetLevel(log.WARN)
	e.Use(middleware.Logger())

	// add CORS support
	log.Info("[LAYER] cors support")
	e.Use(middleware.CORSWithConfig(corsConfig))

	log.Info("[LAYER] server http headers hardening")
	// add server api request hardening using http headers
	e.Use(deployer.hardening)

	log.Info("[LAYER] fake server http header")
	// add fake server header
	e.Use(deployer.fakeServer)

	log.Info("[LAYER] unique request id")
	// Request ID middleware generates a unique id for a request.
	e.Use(middleware.RequestID())

	// add gzip support if client requests it
	log.Info("[LAYER] gzip compression")
	e.Use(middleware.GzipWithConfig(gzipConfig))

	// avoid panics
	log.Info("[LAYER] panic recovery")
	e.Use(middleware.Recover())

	log.Info("[LAYER] / static files")
	//load root static folder
	e.Static("/", resources+"/root")

	// load swagger ui files
	log.Info("[LAYER] /swagger files")
	e.Static("/swagger", resources+"/swagger")

	deployer.register(e)

	// Start secure server
	go func() {
		log.Info("starting https secure server...")
		err := e.StartServer(s)
		if err != nil {
			e.Logger.Info("shutting down https secure the server", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log.Info("graceful shutdown of the service requested")
	log.Info("shutting down http server...")
	if err := httpServerInstance.Shutdown(ctx); err != nil {
		log.Error(err)
	}
	log.Info("shutting down https secure server...")
	if err := e.Shutdown(ctx); err != nil {
		log.Error(err)
	}
	log.Info("graceful shutdown executed")
	log.Info("exiting...")
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
		h.Set("Content-Security-Policy", "default-src 'self' 'unsafe-inline'")
		h.Set("Expect-CT", "enforce, max-age=30")
		h.Set("X-UA-Compatible", "IE=Edge,chrome=1")
		h.Set("x-frame-options", "SAMEORIGIN")
		h.Set("Referrer-Policy", "same-origin")
		h.Set("Feature-Policy", "microphone 'none'; payment 'none'; sync-xhr 'self'")
		h.Set("x-firefox-spdy", "h2")
		h.Set("x-powered-by", "PHP/5.6.38")
		return next(c)
	}
}

// fakeServer middleware function.
func (deployer Deployer) fakeServer(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// add fake server header
		h := c.Response().Header()
		h.Set("server", "Apache")
		h.Set("x-powered-by", "PHP/5.6.38")
		return next(c)
	}
}

// fakeServer antiBots function.
func (deployer Deployer) antiBots(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// add antibots policy
		ua := c.Request().UserAgent()
		if ua == "" || deployer.isBotRequest(ua) {
			//drop the request
			return userAgentErr
		}
		return next(c)
	}
}

// check if user agent string contains bot strings similarities
func (deployer Deployer) isBotRequest(userAgent string) bool {
	var lock = false
	for i := 0; i < len(api.BadBotsList) && !lock; i++ {
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
	c.Logger().Error(err)
	_ = c.JSON(code, api.NewApiError(
		code, err.Error()),
	)
}

// register in echo server, allowed routes
func (deployer Deployer) register(server *echo.Echo) *echo.Echo {
	log.Info("registering routes")
	handlers.NewIndexController().RegisterRouters(server)
	handlers.NewProfileController().RegisterRouters(server)
	handlers.NewTransactionController().RegisterRouters(server)
	handlers.NewEthController(deployer.manager).RegisterRouters(server)
	handlers.NewTokenController(deployer.manager).RegisterRouters(server)
	handlers.NewGanacheController(deployer.manager).RegisterRouters(server)
	return server
}

func NewDeployer() Deployer {
	d := Deployer{}
	d.manager = eth.NewWalletManager()
	return d
}
