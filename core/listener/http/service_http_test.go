// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package http

import (
	"crypto/tls"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/zerjioang/etherniti/core/server/mods/ratelimit"
	"github.com/zerjioang/etherniti/shared/def/listener"
	"github.com/zerjioang/etherniti/thirdparty/echo"
)

func TestHttpListener(t *testing.T) {
	t.Run("instantiation", func(t *testing.T) {
		NewHttpListener()
	})
	t.Run("run", func(t *testing.T) {
		s := NewHttpListener()
		err := s.Listen()
		if err != nil {
			t.Error(err)
		}
		time.Sleep(200000 * time.Second)
	})
	t.Run("request-status", func(t *testing.T) {
		s := NewHttpListener()
		// run the socket servre
		err := s.Listen()
		if err != nil {
			t.Error(err)
		}
		// wait one second to bootup
		time.Sleep(1 * time.Second)
	})
}

func Test_recoverFromPem(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recoverFromPem()
		})
	}
}

func TestHttpListener_GetLocalHostTLS(t *testing.T) {
	type fields struct {
		limiter ratelimit.RateLimitEngine
	}
	tests := []struct {
		name    string
		fields  fields
		want    tls.Certificate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := HttpListener{
				limiter: tt.fields.limiter,
			}
			got, err := l.GetLocalHostTLS()
			if (err != nil) != tt.wantErr {
				t.Errorf("HttpListener.GetLocalHostTLS() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HttpListener.GetLocalHostTLS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHttpListener_RunMode(t *testing.T) {
	type fields struct {
		limiter ratelimit.RateLimitEngine
	}
	type args struct {
		address    string
		background bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := HttpListener{
				limiter: tt.fields.limiter,
			}
			l.RunMode(tt.args.address, tt.args.background)
		})
	}
}

func TestHttpListener_Listen(t *testing.T) {
	type fields struct {
		limiter ratelimit.RateLimitEngine
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := HttpListener{
				limiter: tt.fields.limiter,
			}
			if err := l.Listen(); (err != nil) != tt.wantErr {
				t.Errorf("HttpListener.Listen() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHttpListener_shutdown(t *testing.T) {
	type fields struct {
		limiter ratelimit.RateLimitEngine
	}
	type args struct {
		httpInstance  *echo.Echo
		httpsInstance *echo.Echo
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := HttpListener{
				limiter: tt.fields.limiter,
			}
			l.shutdown(tt.args.httpInstance, tt.args.httpsInstance)
		})
	}
}

func TestHttpListener_buildSecureServerConfig(t *testing.T) {
	type fields struct {
		limiter ratelimit.RateLimitEngine
	}
	type args struct {
		e *echo.Echo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Server
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := HttpListener{
				limiter: tt.fields.limiter,
			}
			got, err := l.buildSecureServerConfig(tt.args.e)
			if (err != nil) != tt.wantErr {
				t.Errorf("HttpListener.buildSecureServerConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HttpListener.buildSecureServerConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHttpListener_buildInsecureServerConfig(t *testing.T) {
	type fields struct {
		limiter ratelimit.RateLimitEngine
	}
	tests := []struct {
		name    string
		fields  fields
		want    *http.Server
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := HttpListener{
				limiter: tt.fields.limiter,
			}
			got, err := l.buildInsecureServerConfig()
			if (err != nil) != tt.wantErr {
				t.Errorf("HttpListener.buildInsecureServerConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HttpListener.buildInsecureServerConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_configureSwaggerJson(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			configureSwaggerJson()
		})
	}
}

func Test_configureSwaggerJsonWithDir(t *testing.T) {
	type args struct {
		resources string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			configureSwaggerJsonWithDir(tt.args.resources)
		})
	}
}

func TestNewHttpListener(t *testing.T) {
	tests := []struct {
		name string
		want listener.ListenerInterface
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHttpListener(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHttpListener() = %v, want %v", got, tt.want)
			}
		})
	}
}
