// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package cache

import (
	"testing"

	"github.com/zerjioang/etherniti/core/util/str"

	"github.com/stretchr/testify/assert"
)

func TestMemoryCache(t *testing.T) {
	t.Run("instantiate", func(t *testing.T) {
		c := NewMemoryCache()
		assert.NotNil(t, c)
	})
	t.Run("get", func(t *testing.T) {
		c := NewMemoryCache()
		c.Set(str.UnsafeBytes("foo"), "bar")
		v, ok := c.Get(str.UnsafeBytes("foo"))
		assert.Equal(t, v, "bar")
		assert.True(t, ok)
	})
	t.Run("set", func(t *testing.T) {
		c := NewMemoryCache()
		c.Set(str.UnsafeBytes("foo"), "bar")
	})
}
