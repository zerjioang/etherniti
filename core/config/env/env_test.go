// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package env

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvironment(t *testing.T) {
	t.Run("get-env", func(t *testing.T) {
		cfg := New()
		assert.NotNil(t, cfg)
	})
	t.Run("get-env-goroutines", func(t *testing.T) {
		var g sync.WaitGroup
		total := 200
		g.Add(total)
		for i := 0; i < total; i++ {
			go func() {
				cfg := New()
				assert.NotNil(t, cfg)
				g.Done()
			}()
		}
		g.Wait()
	})
	t.Run("read-env-all", func(t *testing.T) {
		cfg := New()
		cfg.Load()
	})
	t.Run("read-env-all-goroutines", func(t *testing.T) {
		var g sync.WaitGroup
		total := 200
		g.Add(total)
		for i := 0; i < total; i++ {
			go func() {
				cfg := New()
				cfg.Load()
				assert.NotNil(t, cfg)
				g.Done()
			}()
		}
		g.Wait()
	})
	t.Run("read-env-all-shared-goroutines", func(t *testing.T) {
		var g sync.WaitGroup
		total := 200
		cfg := New()
		g.Add(total)
		for i := 0; i < total; i++ {
			go func() {
				cfg.Load()
				assert.NotNil(t, cfg)
				g.Done()
			}()
		}
		g.Wait()
	})
	t.Run("read-env-key", func(t *testing.T) {
		cfg := New()
		cfg.Load()
		v, found := cfg.Read("HOME")
		assert.NotNil(t, v)
		assert.True(t, found)
	})
	t.Run("read-env-key-goroutines", func(t *testing.T) {
		var g sync.WaitGroup
		total := 200
		g.Add(total)
		cfg := New()
		cfg.Load()
		for i := 0; i < total; i++ {
			go func() {
				v, found := cfg.Read("HOME")
				assert.NotNil(t, v)
				assert.True(t, found)
				g.Done()
			}()
		}
		g.Wait()
	})
	t.Run("string-env-key-goroutines", func(t *testing.T) {
		var g sync.WaitGroup
		total := 200
		g.Add(total)
		cfg := New()
		cfg.Load()
		for i := 0; i < total; i++ {
			go func() {
				v := cfg.String("HOME")
				assert.NotNil(t, v)
				g.Done()
			}()
		}
		g.Wait()
	})
	t.Run("int-env-key-goroutines", func(t *testing.T) {
		var g sync.WaitGroup
		total := 200
		g.Add(total)
		cfg := New()
		cfg.Load()
		for i := 0; i < total; i++ {
			go func() {
				v := cfg.Int("HOME", 0)
				assert.NotNil(t, v)
				g.Done()
			}()
		}
		g.Wait()
	})
}
