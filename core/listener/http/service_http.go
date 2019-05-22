// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package http

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/zerjioang/etherniti/core/listener/swagger"
	"github.com/zerjioang/etherniti/core/util/banner"

	"github.com/zerjioang/etherniti/core/listener/middleware"

	"github.com/zerjioang/etherniti/core/listener/common"
	"github.com/zerjioang/etherniti/shared/def/listener"

	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/server/ratelimit"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

type HttpListener struct {
	limiter ratelimit.RateLimitEngine
}

var (
	//listening ip:port
	listenAddr = config.GetListeningAddress()
	// define http server config for listener service
	defaultHttpServerConfig = http.Server{
		Addr:         listenAddr,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}
)

func (l HttpListener) RunMode(address string, background bool) {
}

func (l HttpListener) Listen(notifier chan error) {
	logger.Info("loading Etherniti Proxy, an Ethereum Multitenant WebAPI")
	//deploy http server only
	e := common.NewServer(middleware.ConfigureServerRoutes)
	logger.Info("starting http server...")
	logger.Info("interface: ", config.GetHttpInterface())
	logger.Info("endpoint: ", listenAddr)
	swagger.ConfigureFromTemplate()
	println(banner.WelcomeBanner())
	// Start server
	go func() {
		err := e.StartServer(&defaultHttpServerConfig)
		if err != nil {
			notifier <- err
			logger.Info("shutting down http server", err)
		}
	}()
	//enable graceful shutdown of http server
	l.shutdown(e, notifier)
}

func (l HttpListener) shutdown(httpInstance *echo.Echo, notifier chan error) {
	// The make built-in returns a value of type T (not *T), and it's memory is
	// initialized.
	quit := make(chan os.Signal)

	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	logger.Info("graceful shutdown of the service requested")
	logger.Info("shutting down http server...")
	if err := httpInstance.Shutdown(ctx); err != nil {
		logger.Error(err)
		notifier <- err
	}
	cancel()
	logger.Info("graceful shutdown executed for http listener")
	logger.Info("exiting...")
	notifier <- nil
}

// create new deployer instance
func NewHttpListener() listener.ListenerInterface {
	d := HttpListener{}
	d.limiter = ratelimit.NewRateLimitEngine()
	return d
}
