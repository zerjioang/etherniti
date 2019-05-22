package ws

import (
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

var (
	hub = NewHub()
)

func init(){
	go hub.run()
}

func WebsocketEntrypoint(c *echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		logger.Error("failed to upgrade to websocket: ", err)
		return err
	}
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in new goroutines.
	go client.read()
	// uncomment below to enable broadcast messages
	// go client.write()
	return nil
}