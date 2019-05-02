// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package e2e

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zerjioang/etherniti/core/handlers"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

func BenchmarkIndexController(b *testing.B) {

	b.Run("instantiate", func(b *testing.B) {
		logger.Enabled(false)

		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = handlers.NewIndexController()
		}
	})

	b.Run("index", func(b *testing.B) {
		// Setup
		e := echo.New()

		logger.Enabled(false)

		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			req := httptest.NewRequest(http.MethodGet, "/v1/public/", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			_ = handlers.Index(c)
		}
	})
	b.Run("status", func(b *testing.B) {
		// Setup
		e := echo.New()
		logger.Enabled(false)

		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			req := httptest.NewRequest(http.MethodGet, "/v1/public/status", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			ctl := handlers.NewIndexController()
			_ = ctl.Status(c)
		}
	})
	b.Run("integrity", func(b *testing.B) {
		// Setup
		e := echo.New()

		logger.Enabled(false)

		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			req := httptest.NewRequest(http.MethodGet, "/v1/public/integrity", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			ctl := handlers.NewIndexController()
			_ = ctl.Integrity(c)
		}
	})
}
