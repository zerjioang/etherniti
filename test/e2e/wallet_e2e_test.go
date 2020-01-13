// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package e2e

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zerjioang/etherniti/shared"

	"github.com/zerjioang/etherniti/core/controllers/wallet"

	"github.com/stretchr/testify/assert"
	"github.com/zerjioang/go-hpc/lib/codes"
	"github.com/zerjioang/go-hpc/thirdparty/echo"
)

func TestWalletController(t *testing.T) {
	t.Run("new-entropy", func(t *testing.T) {
		// Setup
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/v1/wallet/entropy/20", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := shared.AdquireContext(e.NewContext(req, rec))
		ctl := wallet.NewWalletController()

		// Assertions
		if assert.NoError(t, ctl.Entropy(ctx)) {
			assert.Equal(t, codes.StatusBadRequest, rec.Code)
		}
		shared.ReleaseContext(ctx)
	})
}
