package fasthttp

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

/*var (
	socketRequestPool *sync.Pool
)

func init(){
	initPool()
}

// Func to init pool
func initPool() {
	socketRequestPool = &sync.Pool {
		New: func()interface{} {
			return new(socketRequest)
		},
	}
}*/

type socketRequest struct {
	client net.Conn
	raw    []byte
}

func Serve(address string) error {

	newConns := make(chan net.Conn, 128)
	deadConns := make(chan net.Conn, 128)
	parser := make(chan socketRequest, 128)

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

	listen := true
	for listen {
		select {
		// reads socket data and converts it to []byte
		case conn := <-newConns:
			go func() {
				buf := make([]byte, 1024)
				for {
					nr, err := conn.Read(buf)
					if err != nil {
						deadConns <- conn
						break
					} else {
						raw := buf[0:nr]
						parser <- socketRequest{client: conn, raw: raw}
					}
				}
			}()
		// monitors for dead connections in order to close them
		case deadConn := <-deadConns:
			_ = deadConn.Close()
		// process readed []byte data and writes a response back to the client
		case req := <-parser:
			// todo process http request
			// write response back to the client
			_, _ = req.client.Write(req.raw)
			// close the connection with that client
			_ = req.client.Close()
		}
	}
	return l.Close()
}
