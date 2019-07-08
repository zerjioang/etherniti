// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package ws

import (
	"net/http"

	"github.com/zerjioang/etherniti/core/listener/https"
	"github.com/zerjioang/etherniti/core/listener/middleware"
	"github.com/zerjioang/etherniti/core/listener/swagger"

	"github.com/zerjioang/etherniti/core/listener/common"
	"github.com/zerjioang/etherniti/shared/def/listener"

	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/logger"
)

type WebsocketListener struct {
	https.HttpsListener
}

var (
	cfg = config.GetDefaultOpts()
)

func (l WebsocketListener) Listen(notifier chan error) {
	logger.Info("loading Etherniti Proxy, a High Performance Web3 REST Proxy")
	//deploy http server only
	e := common.NewServer(middleware.ConfigureServerRoutes)
	logger.Info("starting websocket server...")
	logger.Info("interface: ", cfg.GetHttpInterface())
	swagger.ConfigureFromTemplate()
	// Start server
	go func() {
		logger.Info("starting websocket server...")
		err := e.StartServer(l.ServerConfig())
		if err != nil {
			notifier <- err
			logger.Info("shutting down websocket server", err)
		}
	}()
	//enable graceful shutdown of http server
	l.ShutdownListener("websocket listener", e, notifier)
}

func (l WebsocketListener) ServerConfig() *http.Server {
	return common.DefaultHttpServerConfig
}

// create new deployer instance
func NewWebsocketListener() listener.ListenerInterface {
	d := WebsocketListener{}
	return d
}
