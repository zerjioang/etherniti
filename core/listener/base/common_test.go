// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package base

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/labstack/echo"
)

func TestNewDefaultServer(t *testing.T) {
	tests := []struct {
		name string
		want *echo.Echo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDefaultServer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDefaultServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewServer(t *testing.T) {
	type args struct {
		configurator func(e *echo.Echo)
	}
	tests := []struct {
		name string
		args args
		want *echo.Echo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewServer(tt.args.configurator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewContext(t *testing.T) {
	type args struct {
		e *echo.Echo
	}
	tests := []struct {
		name string
		args args
		want echo.Context
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewContext(tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewContextFromSocket(t *testing.T) {
	type args struct {
		e    *echo.Echo
		data []byte
	}
	tests := []struct {
		name  string
		args  args
		want  *http.Request
		want1 *httptest.ResponseRecorder
		want2 echo.Context
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := NewContextFromSocket(tt.args.e, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewContextFromSocket() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NewContextFromSocket() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("NewContextFromSocket() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
