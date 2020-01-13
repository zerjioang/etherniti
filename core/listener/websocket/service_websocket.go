// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package ws

import (
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/zerjioang/go-hpc/lib/wsm"

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
	// create a pool to reduce the number of listening goroutines
	epoller *wsm.Epoll
)

func init() {
	// Start epoll
	logger.Info("starting websocket epoll")
	var err error
	epoller, err = wsm.MkEpoll()
	if err != nil {
		logger.Error("failed to start websocket epoll due to: ", err)
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade connection
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		return
	}
	if err := epoller.Add(conn); err != nil {
		logger.Error("failed to add connection: ", err)
		_ = conn.Close()
	}
}

func ListenWebsocket() {
	for {
		connections, err := epoller.Wait()
		if err != nil {
			logger.Error("failed to epoll wait: ", err)
			continue
		}
		for _, conn := range connections {
			if conn == nil {
				break
			}
			data, code, err := wsutil.ReadClientData(conn)
			if err != nil {
				if err := epoller.Remove(conn); err != nil {
					logger.Error("failed to remove: ", err)
				}
				_ = conn.Close()
			} else {
				logger.Debug("msg: ", code, string(data))
			}
		}
	}
}

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
