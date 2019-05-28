// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

// +build !pre
// +build !prod

package config

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	t.Run("IsProfilingEnabled", func(t *testing.T) {
		data := IsProfilingEnabled()
		assert.NotNil(t, data)
	})
	t.Run("IsProfilingEnabled-goroutines", func(t *testing.T) {
		var g sync.WaitGroup
		total := 200
		g.Add(total)
		for i := 0; i < total; i++ {
			go func() {
				data := IsProfilingEnabled()
				assert.NotNil(t, data)
				g.Done()
			}()
		}
		g.Wait()
	})
	t.Run("setup", func(t *testing.T) {
		err := Setup()
		assert.Nil(t, err)
	})
	t.Run("setup-goroutines", func(t *testing.T) {
		var g sync.WaitGroup
		total := 200
		g.Add(total)
		for i := 0; i < total; i++ {
			go func() {
				err := Setup()
				assert.Nil(t, err)
				g.Done()
			}()
		}
		g.Wait()
	})
}
