// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package middleware

import (
	"reflect"
	"testing"

	"github.com/zerjioang/etherniti/thirdparty/echo"
)

func Test_customHTTPErrorHandler(t *testing.T) {
	type args struct {
		err error
		c   echo.ContextInterface
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			customHTTPErrorHandler(tt.args.err, tt.args.c)
		})
	}
}

func TestHttpsRedirect(t *testing.T) {
	type args struct {
		next echo.HandlerFunc
	}
	tests := []struct {
		name string
		args args
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HttpsRedirect(tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HttpsRedirect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hardening(t *testing.T) {
	type args struct {
		next echo.HandlerFunc
	}
	tests := []struct {
		name string
		args args
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hardening(tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hardening() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fakeServer(t *testing.T) {
	type args struct {
		next echo.HandlerFunc
	}
	tests := []struct {
		name string
		args args
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fakeServer(tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fakeServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_antiBots(t *testing.T) {
	type args struct {
		next echo.HandlerFunc
	}
	tests := []struct {
		name string
		args args
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := antiBots(tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("antiBots() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isBotRequest(t *testing.T) {
	type args struct {
		userAgent string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBotRequest(tt.args.userAgent); got != tt.want {
				t.Errorf("isBotRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hostnameCheck(t *testing.T) {
	type args struct {
		next echo.HandlerFunc
	}
	tests := []struct {
		name string
		args args
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hostnameCheck(tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hostnameCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_keepalive(t *testing.T) {
	type args struct {
		next echo.HandlerFunc
	}
	tests := []struct {
		name string
		args args
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := keepalive(tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("keepalive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_customContext(t *testing.T) {
	type args struct {
		next echo.HandlerFunc
	}
	tests := []struct {
		name string
		args args
		want echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := customContext(tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("customContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfigureServerRoutes(t *testing.T) {
	type args struct {
		e *echo.Echo
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ConfigureServerRoutes(tt.args.e)
		})
	}
}

func TestRegisterRoot(t *testing.T) {
	type args struct {
		e *echo.Echo
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterRoot(tt.args.e)
		})
	}
}

func TestGetTestSetup(t *testing.T) {
	tests := []struct {
		name string
		want *echo.Echo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTestSetup(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTestSetup() = %v, want %v", got, tt.want)
			}
		})
	}
}
