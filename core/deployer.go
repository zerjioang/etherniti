// Copyright MethW
// SPDX-License-Identifier: Apache License 2.0

package core

import (
	"net/http"
	"time"

	"github.com/zerjioang/methw/core/keystore/memory"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/zerjioang/methw/core/handlers"
)

type Deployer struct {
	// in memory storage of created wallets
	wallet *memory.InMemoryKeyStorage
}

func (deployer Deployer) Run() error {
	log.Info("loading Ethereum Multitenant Webapi (MethW)")
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
	e.HTTPErrorHandler = deployer.customHTTPErrorHandler

	// remove trailing slash for better usage
	e.Pre(middleware.RemoveTrailingSlash())

	// log all single request
	// configure logging level
	e.Logger.SetLevel(log.WARN)
	e.Use(middleware.Logger())

	// add CORS support
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// add gzip support if client requests it
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

func (deployer Deployer) customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	c.Logger().Error(err)
	errData := struct {
		Code  int
		Error string
	}{}
	errData.Code = code
	errData.Error = err.Error()
	_ = c.JSON(code, errData)
}

// register in echo server, allowed routes
func (deployer Deployer) register(server *echo.Echo) *echo.Echo {
	log.Info("registering routes")
	handlers.NewIndexController().RegisterRouters(server)
	handlers.NewProfileController().RegisterRouters(server)
	return server
}

func NewDeployer() Deployer {
	d := Deployer{}
	d.wallet = memory.NewInMemoryKeyStorage()
	return d
}
