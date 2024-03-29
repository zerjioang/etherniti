// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package common

import (
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/go-hpc/thirdparty/echo"
)

var (
	//default etherniti proxy configuration
	cfg = config.GetDefaultOpts()
	//listening ip:port
	ListenAddr = cfg.GetListeningAddressWithPort()
	// secure listening ip:port
	ListenAddrSecure = cfg.GetListeningSecureAddressWithPort()
	// listening interface
	ListenInterface = cfg.GetHttpInterface()
	// define http server config for listener service
	DefaultHttpServerConfig = &http.Server{
		Addr:         ListenAddr,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}
	// define https server config for listener service
	DefaultHttpsServerConfig = &http.Server{
		Addr:         ListenAddrSecure,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}
	// listener shutdown timeout
	ShutdownTimeout = 10 * time.Second
)

// create a mock  server for testing
func NewDefaultServer() *echo.Echo {
	// build a the server
	e := echo.New()
	// enable debug mode
	e.Debug = cfg.DebugServer()
	e.HidePort = cfg.HideServerData()
	//hide the banner
	e.HideBanner = true
	return e
}
func NewServer(configurator func(e *echo.Echo)) *echo.Echo {
	// build a the server
	e := NewDefaultServer()
	if configurator != nil {
		configurator(e)
	}
	return e
}

// creates a new echo context
func NewContext(e *echo.Echo) echo.Context {
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Preload(req, rec)
	return c
}

func NewContextFromSocket(e *echo.Echo, data []byte) (*http.Request, *httptest.ResponseRecorder, echo.Context) {
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Preload(req, rec)
	return req, rec, c
}
