// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package e2e

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zerjioang/etherniti/core/controllers"
	"github.com/zerjioang/etherniti/shared/protocol"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

func TestWalletController(t *testing.T) {
	t.Run("new-entropy", func(t *testing.T) {
		// Setup
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/v1/wallet/entropy/20", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		ctl := controllers.NewWalletController()

		// Assertions
		if assert.NoError(t, ctl.Entropy(c)) {
			assert.Equal(t, protocol.StatusBadRequest, rec.Code)
		}
	})
}
