// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package config

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRedirectUrl(t *testing.T) {
	type args struct {
		host string
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"redirect",
			args{"127.0.0.1", "/test"},
			"https://127.0.0.1:8080/test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRedirectUrl(tt.args.host, tt.args.path); got != tt.want {
				t.Errorf("GetRedirectUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTLSCryptoData(t *testing.T) {
	t.Run("check-cert-pem", func(t *testing.T) {
		assert.NotNil(t, GetCertPem() != nil)
	})
	t.Run("check-cert-pem-goroutines", func(t *testing.T) {
		var g sync.WaitGroup
		total := 200
		g.Add(total)
		for i := 0; i < total; i++ {
			go func() {
				data := GetCertPem()
				assert.NotNil(t, data)
				g.Done()
			}()
		}
		g.Wait()
	})
	t.Run("check-key-pem", func(t *testing.T) {
		assert.NotNil(t, GetKeyPem() != nil)
	})
	t.Run("check-key-pem-goroutines", func(t *testing.T) {
		var g sync.WaitGroup
		total := 200
		g.Add(total)
		for i := 0; i < total; i++ {
			go func() {
				data := GetKeyPem()
				assert.NotNil(t, data)
				g.Done()
			}()
		}
		g.Wait()
	})
}

func TestCommonConfig(t *testing.T) {
	t.Run("is-http", func(t *testing.T) {
		assert.NotNil(t, IsHttpMode())
	})
	t.Run("is-http-goroutines", func(t *testing.T) {
		var g sync.WaitGroup
		total := 200
		g.Add(total)
		for i := 0; i < total; i++ {
			go func() {
				data := IsHttpMode()
				assert.NotNil(t, data)
				g.Done()
			}()
		}
		g.Wait()
	})
	t.Run("is-socket", func(t *testing.T) {
		assert.NotNil(t, IsUnixSocketMode())
	})
	t.Run("is-profiling-enabled", func(t *testing.T) {
		assert.NotNil(t, IsProfilingEnabled())
	})
	t.Run("is-service-listening-enabled", func(t *testing.T) {
		assert.NotNil(t, ServiceListeningMode())
	})
}
