// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package e2e

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zerjioang/etherniti/shared"

	"github.com/zerjioang/etherniti/core/controllers/index"

	"github.com/zerjioang/etherniti/shared/constants"

	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/go-hpc/thirdparty/echo"
)

func BenchmarkIndexController(b *testing.B) {

	b.Run("instantiate", func(b *testing.B) {
		logger.Enabled(false)

		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = index.NewIndexController()
		}
	})

	b.Run("index", func(b *testing.B) {
		// Setup
		e := echo.New()
		ctl := index.NewIndexController()
		logger.Enabled(false)

		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			req := httptest.NewRequest(http.MethodGet, constants.ApiVersion+"/", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := shared.AdquireContext(e.NewContext(req, rec))
			_ = ctl.Index(ctx)
			shared.ReleaseContext(ctx)
		}
	})
	b.Run("status", func(b *testing.B) {
		// Setup
		e := echo.New()
		ctl := index.NewIndexController()
		logger.Enabled(false)

		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			req := httptest.NewRequest(http.MethodGet, "/v1/metrics", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := shared.AdquireContext(e.NewContext(req, rec))
			_ = ctl.Status(ctx)
			shared.ReleaseContext(ctx)
		}
	})
	b.Run("integrity", func(b *testing.B) {
		// Setup
		e := echo.New()
		ctl := index.NewIndexController()
		logger.Enabled(false)

		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			req := httptest.NewRequest(http.MethodGet, "/v1/integrity", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := shared.AdquireContext(e.NewContext(req, rec))
			_ = ctl.Integrity(ctx)
			shared.ReleaseContext(ctx)
		}
	})
}
