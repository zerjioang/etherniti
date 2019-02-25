// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
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
			ctl := NewProfileController()

			// Assertions
			if assert.NoError(t, ctl.create(c)) {
				assert.Equal(t, http.StatusBadRequest, rec.Code)
			}
		})
	})
	t.Run("read_profile", func(t *testing.T) {
		t.Run("empty-request", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			ctl := NewProfileController()

			// Assertions
			if assert.NoError(t, ctl.create(c)) {
				assert.Equal(t, http.StatusBadRequest, rec.Code)
			}
		})
	})
	t.Run("update_profile", func(t *testing.T) {
		t.Run("empty-request", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			ctl := NewProfileController()

			// Assertions
			if assert.NoError(t, ctl.create(c)) {
				assert.Equal(t, http.StatusBadRequest, rec.Code)
			}
		})
	})
}
