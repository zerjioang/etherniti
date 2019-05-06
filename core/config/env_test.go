// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package config

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvironment(t *testing.T) {
	t.Run("get-env", func(t *testing.T) {
		cfg := GetEnvironment()
		assert.NotNil(t, cfg)
	})
	t.Run("get-env-goroutines", func(t *testing.T) {
		var g sync.WaitGroup
		total := 200
		g.Add(total)
		for i := 0; i < total; i++ {
			go func() {
				cfg := GetEnvironment()
				assert.NotNil(t, cfg)
				g.Done()
			}()
		}
		g.Wait()
	})
}
