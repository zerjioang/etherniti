// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMemoryCache_Get(t *testing.T) {
	c := NewMemoryCache()
	c.Set("foo", "bar")
	v, ok := c.Get("foo")
	assert.Equal(t, v, "bar")
	assert.True(t, ok)
}

func TestMemoryCache_Set(t *testing.T) {
	c := NewMemoryCache()
	c.Set("foo", "bar")
}

func TestNewMemoryCache(t *testing.T) {
	c := NewMemoryCache()
	assert.NotNil(t, c)
}
