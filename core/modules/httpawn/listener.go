package httpawn

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

const (
	messageBufferSize         = 1024
	concurrentConnectionsSize = 128
	separator                 = "\n"
)

var (
	separatorRaw = []byte(separator)
)

func Serve(address string, r *Router) error {

	newConns := make(chan net.Conn, concurrentConnectionsSize)
	deadConns := make(chan net.Conn, concurrentConnectionsSize)
	parser := make(chan *socketRequest, concurrentConnectionsSize)

	println("serving tcp")
	l, err := net.Listen("tcp4", address)
	if err != nil {
		return err
	}
	// some randomness
	rand.Seed(time.Now().Unix())

	// use a goroutine to accept new clients connections up to 128
	go func() {
		for {
			conn, err := l.Accept()
			// customize our connection for performance
			if err != nil {
				fmt.Println("accept error:", err)
			} else {
				newConns <- conn
			}
		}
	}()

	for {
		select {
		// reads socket data and converts it to []byte
		case conn := <-newConns:
			go func() {
				buf := make([]byte, messageBufferSize)
				for {
					nr, err := conn.Read(buf)
					if err != nil {
						deadConns <- conn
						break
					} else {
						raw := buf[0:nr]
						parser <- sPool.Load(conn, raw)
					}
				}
			}()
		// monitors for dead connections in order to close them
		case deadConn := <-deadConns:
			_ = deadConn.Close()
		// process readed []byte data and writes a response back to the client
		case req := <-parser:
			header, body := processHttpRequest(r, req)
			// write response back to the client
			_, _ = req.client.Write(header)
			_, _ = req.client.Write(separatorRaw)
			_, _ = req.client.Write(body)
			// close the connection with that client
			_ = req.client.Close()
			// put request back in the pool
			sPool.Store(req)
		}
	}
	return l.Close()
}
