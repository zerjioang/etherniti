// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package edition_test

import (
	"sync"
	"testing"

	"github.com/zerjioang/etherniti/core/config/edition"
	"github.com/zerjioang/etherniti/shared/constants"

	"github.com/stretchr/testify/assert"
)

func TestConfigExtra(t *testing.T) {
	t.Run("get-edition", func(t *testing.T) {
		e := edition.Edition()
		assert.NotNil(t, e)
		assert.True(t, e != constants.Unknown)
	})
	t.Run("get-edition-goroutines", func(t *testing.T) {
		var g sync.WaitGroup
		total := 200
		g.Add(total)
		for i := 0; i < total; i++ {
			go func() {
				e := edition.Edition()
				assert.NotNil(t, e)
				assert.True(t, e != constants.Unknown)
				g.Done()
			}()
		}
		g.Wait()
	})
	t.Run("is-valid-edition", func(t *testing.T) {
		e := edition.IsOpenSource() || edition.IsEnterprise()
		assert.True(t, e)
	})
	t.Run("is-valid-edition-goroutines", func(t *testing.T) {
		var g sync.WaitGroup
		total := 200
		g.Add(total)
		for i := 0; i < total; i++ {
			go func() {
				e := edition.IsOpenSource() || edition.IsEnterprise()
				assert.True(t, e)
				g.Done()
			}()
		}
		g.Wait()
	})
	t.Run("check-is-allowed-edition", func(t *testing.T) {
		e := edition.IsValidEdition()
		assert.True(t, e)
	})
	t.Run("is-valid-allowed-goroutines", func(t *testing.T) {
		var g sync.WaitGroup
		total := 200
		g.Add(total)
		for i := 0; i < total; i++ {
			go func() {
				e := edition.IsValidEdition()
				assert.True(t, e)
				g.Done()
			}()
		}
		g.Wait()
	})
}
