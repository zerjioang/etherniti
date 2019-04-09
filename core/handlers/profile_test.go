// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/zerjioang/etherniti/thirdparty/echo"
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

func TestNewProfileController(t *testing.T) {
	tests := []struct {
		name string
		want ProfileController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProfileController(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProfileController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProfileController_create(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		ctl     ProfileController
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctl := ProfileController{}
			if err := ctl.create(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("ProfileController.create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProfileController_validate(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		ctl     ProfileController
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctl := ProfileController{}
			if err := ctl.validate(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("ProfileController.validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProfileController_count(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		ctl     ProfileController
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctl := ProfileController{}
			if err := ctl.count(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("ProfileController.count() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProfileController_RegisterRouters(t *testing.T) {
	type args struct {
		router *echo.Group
	}
	tests := []struct {
		name string
		ctl  ProfileController
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctl := ProfileController{}
			ctl.RegisterRouters(tt.args.router)
		})
	}
}
