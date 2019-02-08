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
func TestIndex(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctl := NewIndexController()

	// Assertions
	if assert.NoError(t, ctl.index(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, indexWelcome, rec.Body.String())
	}
}
