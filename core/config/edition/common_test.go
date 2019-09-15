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
	t.Run("is-valid-edition-or", func(t *testing.T) {
		e := edition.IsOpenSource() || edition.IsEnterprise()
		assert.True(t, e)
	})
	t.Run("is-valid-edition-or-goroutines", func(t *testing.T) {
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
	t.Run("is-opensource", func(t *testing.T) {
		e := edition.IsOpenSource()
		if edition.Edition() == constants.OpenSource {
			assert.True(t, e)
		} else {
			assert.False(t, e)
		}
	})
	t.Run("is-opensource-goroutines", func(t *testing.T) {
		var g sync.WaitGroup
		total := 200
		g.Add(total)
		for i := 0; i < total; i++ {
			go func() {
				e := edition.IsOpenSource()
				if edition.Edition() == constants.OpenSource {
					assert.True(t, e)
				} else {
					assert.False(t, e)
				}
				g.Done()
			}()
		}
		g.Wait()
	})
	t.Run("is-enterprise", func(t *testing.T) {
		e := edition.IsEnterprise()
		if edition.Edition() == constants.Enterprise {
			assert.True(t, e)
		} else {
			assert.False(t, e)
		}
	})
	t.Run("is-enterprise-goroutines", func(t *testing.T) {
		var g sync.WaitGroup
		total := 200
		g.Add(total)
		for i := 0; i < total; i++ {
			go func() {
				e := edition.IsEnterprise()
				if edition.Edition() == constants.Enterprise {
					assert.True(t, e)
				} else {
					assert.False(t, e)
				}
				g.Done()
			}()
		}
		g.Wait()
	})
	t.Run("extra-setup", func(t *testing.T) {
		err := edition.ExtraSetup()
		assert.NoError(t, err)
	})
	t.Run("extra-setup-goroutines", func(t *testing.T) {
		var g sync.WaitGroup
		total := 200
		g.Add(total)
		for i := 0; i < total; i++ {
			go func() {
				err := edition.ExtraSetup()
				assert.NoError(t, err)
				g.Done()
			}()
		}
		g.Wait()
	})
}
