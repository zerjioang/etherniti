// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package socket

import (
	"context"
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/zerjioang/etherniti/util/banner"

	"github.com/zerjioang/etherniti/core/listener/common"
	"github.com/zerjioang/etherniti/shared/def/listener"

	"github.com/zerjioang/go-hpc/thirdparty/echo"

	"github.com/zerjioang/etherniti/core/logger"

	"github.com/zerjioang/go-hpc/thirdparty/gommon/log"
)

// UNIX domain sockets are a method by which processes on the same host can communicate
// Communication is bidirectional with stream sockets and unidirectional with datagram sockets.
type UnixSocketListener struct {
	e    *echo.Echo
	path string
	mode bool
}

func (l UnixSocketListener) Listen(notifier chan error) {
	logger.Info("loading Etherniti Proxy, a High Performance Web3 REST Proxy via unix sockets")
	l.e = common.NewDefaultServer()
	if l.mode {
		l.background(notifier)
		//graceful shutdown of http and https server
		l.shutdown(notifier)
	} else {
		l.foreground(notifier)
	}
}

// fetch specific server configuration
// this method will return nil for unix listener
func (l UnixSocketListener) ServerConfig() *http.Server {
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
	_ = c.Close()
}

// new http format socket client
func socketHttpClient(socketPath string) *http.Client {
	httpc := &http.Client{
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
		req, rec, _ := common.NewContextFromSocket(l.e, data)
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
func (l UnixSocketListener) background(notifier chan error) {
	go l.foreground(notifier)
}

// run unix socket server instance in foreground
func (l UnixSocketListener) foreground(notifier chan error) {
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
	logger.Info("starting unix socket server...")
	ln, err := net.Listen("unix", l.path)
	if err != nil {
		log.Error("unix socket listen error: ", err)
		notifier <- err
	}

	l.e.Listener = ln
	notifier <- l.e.Start("")

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

func (l UnixSocketListener) shutdown(notifier chan error) {
	// The make built-in returns a value of type T (not *T), and it's memory is
	// initialized.
	quit := make(chan os.Signal)

	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Info("graceful shutdown of the unix socket listener requested")
	if l.e.Listener != nil {
		logger.Info("shutting down unix socket listener...")
		if err := l.e.Listener.Close(); err != nil {
			logger.Error(err)
			notifier <- err
		}
	}
	logger.Info("graceful shutdown executed for unix socket listener")
	logger.Info("exiting...")
	notifier <- nil
}
