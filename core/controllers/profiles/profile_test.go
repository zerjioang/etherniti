// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package profiles

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zerjioang/etherniti/shared"

	"github.com/stretchr/testify/assert"
	"github.com/zerjioang/go-hpc/lib/codes"
	"github.com/zerjioang/go-hpc/thirdparty/echo"
)

// an end to end test from go
// it deploys a temporal http server for testing
func TestProfile(t *testing.T) {
	// Setup
	e := echo.New()

	t.Run("create_profile", func(t *testing.T) {
		t.Run("empty-request", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			ctx := shared.AdquireContext(c)
			ctl := NewProfileController()

			// Assertions
			if assert.NoError(t, ctl.create(ctx)) {
				assert.Equal(t, codes.StatusBadRequest, rec.Code)
			}
			shared.ReleaseContext(ctx)
		})
	})
	t.Run("read_profile", func(t *testing.T) {
		t.Run("empty-request", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			ctx := shared.AdquireContext(c)
			ctl := NewProfileController()

			// Assertions
			if assert.NoError(t, ctl.create(ctx)) {
				assert.Equal(t, codes.StatusBadRequest, rec.Code)
			}
			shared.ReleaseContext(ctx)
		})
	})
	t.Run("update_profile", func(t *testing.T) {
		t.Run("empty-request", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			ctx := shared.AdquireContext(c)
			ctl := NewProfileController()

			// Assertions
			if assert.NoError(t, ctl.create(ctx)) {
				assert.Equal(t, codes.StatusBadRequest, rec.Code)
			}
			shared.ReleaseContext(ctx)
		})
	})
}
