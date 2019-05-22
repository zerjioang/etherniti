package ws

import (
	"bytes"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/zerjioang/etherniti/core/logger"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:    1024,
		WriteBufferSize:   1024,
		HandshakeTimeout:  time.Second * 2,
		EnableCompression: true,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}

// read messages from the websocket connection to the hub.
//
// The application runs read in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
// todo handler errors properly
func (c *Client) read() {
	c.conn.SetReadLimit(maxMessageSize)
	_ = c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		_ = c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	run := true
	for run {
		msgType, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logger.Error("unexpected websocket close error: ", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		sendErr := c.conn.WriteMessage(msgType, []byte(strings.ToUpper(string(message))))
		if sendErr != nil {
			logger.Error("failed to write client response to websocket connection due to error: ", sendErr)
		}
		// uncomment this line to broadcast input message to all connected clients
		//c.hub.broadcast <- message
		//detect if client send a close message
		run = !(bytes.Compare(message, []byte("close")) == 0)
	}
	// close and unregister
	c.hub.unregister <- c
	_ = c.conn.Close()
}

// write messages from the hub to the websocket connection.
//
// A goroutine running write is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
// todo handler errors properly
func (c *Client) write() {
	ticker := time.NewTicker(pingPeriod)
	exec := true
	for exec {
		select {
		case message, ok := <-c.send:
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				_ = c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				exec = false
				break
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				exec = false
				break
			}
			_, _ = w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				_, _ = w.Write(newline)
				_, _ = w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				exec = false
				break
			}
		case <-ticker.C:
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				exec = false
				break
			}
		}
	}
	// stop and close
	ticker.Stop()
	_ = c.conn.Close()
}
