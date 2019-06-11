// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package common

import (
	"net/http"
	"net/http/httptest"

	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

var (
	cfg = config.GetDefaultOpts()
)

// create a mock  server for testing
func NewDefaultServer() *echo.Echo {
	// build a the server
	e := echo.New()
	// enable debug mode
	e.Debug = cfg.DebugServer()
	e.HidePort = cfg.HideServerData()
	//hide the banner
	e.HideBanner = cfg.HideServerData()
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
func NewContext(e *echo.Echo) *echo.Context {
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return c
}

func NewContextFromSocket(e *echo.Echo, data []byte) (*http.Request, *httptest.ResponseRecorder, *echo.Context) {
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return req, rec, c
}
