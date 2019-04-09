// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package socket

import (
	"context"
	"io"
	"net"
	"net/http"
	"syscall"
	"time"

	"github.com/zerjioang/etherniti/core/util/banner"

	"github.com/zerjioang/etherniti/core/listener/base"
	"github.com/zerjioang/etherniti/shared/def/listener"

	"github.com/zerjioang/etherniti/thirdparty/echo"

	"github.com/zerjioang/etherniti/core/logger"

	"github.com/zerjioang/etherniti/thirdparty/gommon/log"
)

// UNIX domain sockets are a method by which processes on the same host can communicate
// Communication is bidirectional with stream sockets and unidirectional with datagram sockets.
type UnixSocketListener struct {
	e    *echo.Echo
	path string
	mode bool
}

func (l UnixSocketListener) RunMode(socketPath string, background bool) {
	l.path = socketPath
	l.mode = background
}

func (l UnixSocketListener) Listen() error {
	logger.Info("loading Etherniti Proxy, an Ethereum Multitenant WebAPI via unix sockets")
	l.e = base.NewDefaultServer()
	if l.mode {
		l.background()
	} else {
		l.foreground()
	}
	return nil
}

// create new socket listener instance
func NewSocketListener() listener.ListenerInterface {
	d := UnixSocketListener{}
	return d
}

// socket reader function

func reader(r io.Reader) {
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf[:])
		if err != nil {
			return
		}
		log.Debug("Client got:", string(buf[0:n]))
	}
}

// new socket client listener
func socketClient() {
	c, err := net.Dial("unix", "/tmp/go.sock")
	if err != nil {
		log.Fatal("Dial error", err)
	}
	defer c.Close()

	go reader(c)
	for {
		msg := "hi"
		_, err := c.Write([]byte(msg))
		if err != nil {
			log.Fatal("Write error:", err)
			break
		}
		log.Debug("Client sent:", msg)
		time.Sleep(1e9)
	}
}

// new http format socket client
func socketHttpClient(socketPath string) http.Client {
	httpc := http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", socketPath)
			},
		},
	}
	return httpc
}

// new socket server listener. echo server
func (l UnixSocketListener) unixServerListener(c net.Conn) {
	for {
		buf := make([]byte, 512)
		nr, err := c.Read(buf)
		if err != nil {
			return
		}
		// input data via socket
		data := buf[0:nr]
		log.Debug("server got:", string(data))

		//build request
		req, rec, _ := base.NewContextFromSocket(l.e, data)
		// execute appropiare handler for current request
		l.e.ServeHTTP(rec, req)
		responseBytes := rec.Body.Bytes()

		log.Debug("unix socket server response is: ", string(responseBytes))
		// write data back to the socket
		_, err = c.Write(data)
		if err != nil {
			log.Error("writing to client error: ", err)
		}
	}
}

// run unix socket server instance in background
func (l UnixSocketListener) background() {
	go l.foreground()
}

// run unix socket server instance in foreground
func (l UnixSocketListener) foreground() error {
	log.Debug("starting unix socket server")
	// Instead of identifying a server by an IP address and port,
	// a UNIX domain socket is known by a pathname.
	// Obviously the client and server have to agree on the pathname
	// for them to find each other.
	// The server binds the pathname to the socket:

	// Note that, once created,
	// this socket file will continue to exist,
	// even after the server exits.
	// If the server subsequently restarts,
	// the file prevents re-binding:
	unErr := syscall.Unlink(l.path)
	if unErr != nil {
		log.Warn("failed to unlink unix socket: ", unErr)
	}
	println(banner.WelcomeBanner())
	ln, err := net.Listen("unix", l.path)
	if err != nil {
		log.Error("unix socket listen error: ", err)
		return err
	}

	l.e.Listener = ln
	return l.e.Start("")

	/*sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)
	go func(ln net.Listener, c chan os.Signal) {
		sig := <-c
		log.Printf("caught signal %s: shutting down.", sig)
		closeErr := ln.Close()
		if closeErr != nil {
			log.Error("unix socket close error: ", closeErr)
		}
	}(ln, sigc)

	for {
		fd, err := ln.Accept()
		if err != nil {
			log.Error("unix socket accept error: ", err)
			return err
		}
		go l.unixServerListener(fd)
	}*/
}
