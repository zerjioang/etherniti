// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package http

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"github.com/zerjioang/etherniti/core/listener/middleware"
	"github.com/zerjioang/etherniti/core/listener/swagger"

	"github.com/zerjioang/etherniti/core/listener/common"
	"github.com/zerjioang/etherniti/shared/def/listener"

	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

type HttpListener struct{}

//fetch specific server configuration
func (l HttpListener) ServerConfig() *http.Server {
	return common.DefaultHttpServerConfig
}

func (l HttpListener) Listen(notifier chan error) {
	logger.Info("loading Etherniti Proxy, a High Performance Web3 REST Proxy")
	logger.Info("loading http listener")
	//deploy http server only
	e := common.NewServer(middleware.ConfigureServerRoutes)
	logger.Info("starting http server...")
	logger.Info("interface: ", common.ListenInterface)
	logger.Info("endpoint: ", common.ListenAddr)
	swagger.ConfigureFromTemplate()
	// Start http server
	go func() {
		logger.Info("server listening")
		err := e.StartServer(l.ServerConfig())
		if err != nil {
			notifier <- err
			logger.Info("shutting down http server: ", err)
		}
	}()
	//enable graceful shutdown of http server
	l.ShutdownListener("http", e, notifier)
}

func (l HttpListener) ShutdownListener(listenerName string, instance *echo.Echo, notifier chan error) {
	// The make built-in returns a value of type T (not *T), and it's memory is
	// initialized.
	quit := make(chan os.Signal)

	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Info("graceful shutdown of the service requested")
	ctx, cancel := context.WithTimeout(context.Background(), common.ShutdownTimeout)
	logger.Info("shutting down " + listenerName + " server listener service...")
	if err := instance.Shutdown(ctx); err != nil {
		logger.Error(err)
		notifier <- err
	}
	cancel()
	logger.Info("graceful shutdown executed for " + listenerName + " listener")
	logger.Info("exiting...")
	notifier <- nil
}

// create new http listener instance
func NewHttpListenerCustom() HttpListener {
	d := HttpListener{}
	return d
}

// create new http listener instance
func NewHttpListener() listener.ListenerInterface {
	return NewHttpListenerCustom()
}
