// Copyright gaethway
// SPDX-License-Identifier: Apache License 2.0

package core

import (
	"errors"
	"github.com/zerjioang/gaethway/core/api"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/zerjioang/gaethway/core/handlers"
)

var (
	userAgentErr = errors.New("not authorized. security policy not satisfied")
	gopath = os.Getenv("GOPATH")
	resources = gopath + "/src/github.com/zerjioang/gaethway/resources"
)

type Deployer struct {
}

func (deployer Deployer) Run() error {
	log.Info("loading Ethereum Multitenant Webapi (gaethway)")
	e := echo.New()

	// enable debug mode
	e.Debug = true

	//configure custom server
	s := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}
	//hide the banner
	e.HideBanner = true

	// add a custom error handler
	log.Info("[LAYER] custom error handler")
	e.HTTPErrorHandler = deployer.customHTTPErrorHandler


	//load root static folder
	e.Static("/", resources+"/root")

	// load swagger ui files
	e.Static("/swagger", resources+"/swagger")

	// antibots, crawler middleware
	// avoid bots and crawlers
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
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// add server api request hardening using http headers
	e.Use(deployer.hardening)

	// add fake server header
	e.Use(deployer.fakeServer)

	// add gzip support if client requests it
	log.Info("[LAYER] gzip compression")
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	// avoid panics
	e.Use(middleware.Recover())

	deployer.register(e)

	log.Info("starting http server...")
	err := e.StartServer(s)
	return err
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
		h.Set("X-Firefox-Spdy", "h2")
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

func (deployer Deployer) isBotRequest(userAgent string) bool {
	return false
}

// keepalive middleware function.
func (deployer Deployer) keepalive(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// add keep alive headers
		c.Response().Header().Set("Connection", "Keep-Alive")
		c.Response().Header().Set("Keep-Alive", "timeout=5, max=1000")
		return next(c)
	}
}

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
	handlers.NewEthController().RegisterRouters(server)
	return server
}

func NewDeployer() Deployer {
	d := Deployer{}
	return d
}
