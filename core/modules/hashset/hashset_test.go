// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package hashset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashSet(t *testing.T) {
	t.Run("instantiate", func(t *testing.T) {
		assert.NotNil(t, NewHashSet())
	})
	t.Run("add", func(t *testing.T) {
		set := NewHashSet()
		assert.NotNil(t, set)

		set.Add("India")
		set.Add("Australia")
		set.Add("South Africa")
		set.Add("India") // adding duplicate elements
	})
	t.Run("count", func(t *testing.T) {
		t.Run("count-0", func(t *testing.T) {
			set := NewHashSet()
			assert.NotNil(t, set)

			assert.Equal(t, set.Count(), 0)
		})

		t.Run("count-1", func(t *testing.T) {
			set := NewHashSet()
			assert.NotNil(t, set)

			set.Add("India")

			assert.Equal(t, set.Count(), 1)
		})

		t.Run("count-0", func(t *testing.T) {
			set := NewHashSet()
			assert.NotNil(t, set)

			set.Add("India")
			set.Add("Australia")
			set.Add("South Africa")
			set.Add("India") // adding duplicate elements

			assert.Equal(t, set.Count(), 3)
		})
	})

	t.Run("double-clear", func(t *testing.T) {
		set := NewHashSet()
		assert.NotNil(t, set)

		set.Add("India")
		set.Add("Australia")
		set.Add("South Africa")
		set.Add("India") // adding duplicate elements

		assert.Equal(t, set.Count(), 3)

		set.Clear()
		assert.Equal(t, set.Count(), 0)

		set.Add("India") // adding duplicate elements
		assert.Equal(t, set.Count(), 1)

		set.Clear()
		assert.Equal(t, set.Count(), 0)
	})
}
