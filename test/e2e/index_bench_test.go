// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package e2e

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zerjioang/etherniti/core/controllers"
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
			_ = controllers.NewIndexController()
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
			req := httptest.NewRequest(http.MethodGet, "/v1/", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			_ = controllers.Index(c)
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
			req := httptest.NewRequest(http.MethodGet, "/v1/metrics", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			ctl := controllers.NewIndexController()
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
			req := httptest.NewRequest(http.MethodGet, "/v1/integrity", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			ctl := controllers.NewIndexController()
			_ = ctl.Integrity(c)
		}
	})
}
