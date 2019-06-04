// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package ws

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
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

type WebsocketListener struct {
}

var (
	// define http server config for listener service
	defaultHttpServerConfig = http.Server{
		Addr:         config.GetListeningAddressWithPort(),
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}
)

func (l WebsocketListener) RunMode(address string, background bool) {
}

func (l WebsocketListener) Listen(notifier chan error) {
	logger.Info("loading Etherniti Proxy, a High Performance Web3 REST Proxy")
	//deploy http server only
	e := common.NewServer(middleware.ConfigureServerRoutes)
	logger.Info("starting websocket server...")
	logger.Info("interface: ", config.GetHttpInterface())
	swagger.ConfigureFromTemplate()
	// Start server
	go func() {
		println(banner.WelcomeBanner())
		logger.Info("starting websocket server...")
		err := e.StartServer(&defaultHttpServerConfig)
		if err != nil {
			notifier <- err
			logger.Info("shutting down websocket server", err)
		}
	}()
	//enable graceful shutdown of http server
	l.shutdown(e, notifier)
}

func (l WebsocketListener) shutdown(httpInstance *echo.Echo, notifier chan error) {
	// The make built-in returns a value of type T (not *T), and it's memory is
	// initialized.
	quit := make(chan os.Signal)

	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	logger.Info("graceful shutdown of the service requested")
	logger.Info("shutting down websocket server...")
	if err := httpInstance.Shutdown(ctx); err != nil {
		logger.Error(err)
		notifier <- err
	}
	cancel()
	logger.Info("graceful shutdown executed for websocket listener")
	logger.Info("exiting...")
	notifier <- nil
}

// create new deployer instance
func NewWebsocketListener() listener.ListenerInterface {
	d := WebsocketListener{}
	return d
}
