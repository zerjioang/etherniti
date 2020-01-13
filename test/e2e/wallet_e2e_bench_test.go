// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package e2e

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zerjioang/etherniti/shared"

	"github.com/zerjioang/etherniti/core/controllers/wallet"

	"github.com/zerjioang/etherniti/core/logger"

	"github.com/zerjioang/go-hpc/thirdparty/echo"
)

func BenchmarkWalletController(b *testing.B) {
	b.Run("new-entropy", func(b *testing.B) {
		// Setup
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/v1/wallet/entropy/20", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := shared.AdquireContext(e.NewContext(req, rec))
		ctl := wallet.NewWalletController()

		logger.Enabled(false)

		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = ctl.Entropy(ctx)
		}
		shared.ReleaseContext(ctx)
	})
}
