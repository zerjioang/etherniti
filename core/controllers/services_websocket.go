package controllers

import (
	"bytes"
	"github.com/gorilla/websocket"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/thirdparty/echo"
	"net/http"
	"time"
)

var (
	upgrader = websocket.Upgrader{
		HandshakeTimeout: time.Second * 2,
		EnableCompression:true,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	clients = make(map[*websocket.Conn]bool)
)

func WebsocketEntrypoint(c *echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	run := true
	for run {
		// Write
		err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		if err != nil {
			c.Logger().Error(err)
		}

		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			logger.Error("websocket message error: ", err)
		}
		logger.Debug(msg)
		// check if user requested a disconnection
		run = bytes.Compare(msg, []byte("close")) == 0
	}
	return ws.Close()
}
