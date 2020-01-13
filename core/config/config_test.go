// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

// +build !pre
// +build !prod

package config_test

import (
	"sync"
	"testing"

	"github.com/zerjioang/etherniti/core/config"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	t.Run("IsProfilingEnabled", func(t *testing.T) {
		opts := config.GetDefaultOpts()
		data := config.IsProfilingEnabled(&opts)
		assert.NotNil(t, data)
	})
	t.Run("IsProfilingEnabled-goroutines", func(t *testing.T) {
		var g sync.WaitGroup
		total := 200
		g.Add(total)
		opts := config.GetDefaultOpts()
		for i := 0; i < total; i++ {
			go func() {
				data := config.IsProfilingEnabled(&opts)
				assert.NotNil(t, data)
				g.Done()
			}()
		}
		g.Wait()
	})
	t.Run("setup", func(t *testing.T) {
		opts := config.GetDefaultOpts()
		err := config.Setup(&opts)
		assert.Nil(t, err)
	})
	t.Run("setup-goroutines", func(t *testing.T) {
		var g sync.WaitGroup
		total := 200
		g.Add(total)
		opts := config.GetDefaultOpts()
		for i := 0; i < total; i++ {
			go func() {
				err := config.Setup(&opts)
				assert.Nil(t, err)
				g.Done()
			}()
		}
		g.Wait()
	})
}
