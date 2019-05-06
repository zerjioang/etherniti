// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package hashset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtomicHashSet(t *testing.T) {
	t.Run("instantiate", func(t *testing.T) {
		assert.NotNil(t, NewAtomicHashSet())
	})
	t.Run("add", func(t *testing.T) {
		set := NewAtomicHashSet()
		assert.NotNil(t, set)

		set.Add("India")
		set.Add("Australia")
		set.Add("South Africa")
		set.Add("India") // adding duplicate elements
	})
	t.Run("count", func(t *testing.T) {
		t.Run("count-0", func(t *testing.T) {
			set := NewAtomicHashSet()
			assert.NotNil(t, set)

			assert.Equal(t, set.Size(), 0)
		})

		t.Run("count-1", func(t *testing.T) {
			set := NewAtomicHashSet()
			assert.NotNil(t, set)

			set.Add("India")

			assert.Equal(t, set.Size(), 1)
		})

		t.Run("count-0", func(t *testing.T) {
			set := NewAtomicHashSet()
			assert.NotNil(t, set)

			set.Add("India")
			set.Add("Australia")
			set.Add("South Africa")
			set.Add("India") // adding duplicate elements

			assert.Equal(t, set.Size(), 3)
		})
	})

	t.Run("double-clear", func(t *testing.T) {
		set := NewAtomicHashSet()
		assert.NotNil(t, set)

		set.Add("India")
		set.Add("Australia")
		set.Add("South Africa")
		set.Add("India") // adding duplicate elements

		assert.Equal(t, set.Size(), 3)

		set.Clear()
		assert.Equal(t, set.Size(), 0)

		set.Add("India") // adding duplicate elements
		assert.Equal(t, set.Size(), 1)

		set.Clear()
		assert.Equal(t, set.Size(), 0)
	})
}
