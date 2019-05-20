// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package security

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPhising(t *testing.T) {
	t.Run("blacklist-bytes", func(t *testing.T) {
		list := PhishingBlacklistRawBytes()
		assert.NotNil(t, list)
		assert.True(t, len(list) > 0)
	})
	t.Run("whitelist-bytes", func(t *testing.T) {
		raw := PhishingWhitelistRawBytes()
		assert.NotNil(t, raw)
		assert.True(t, len(raw) > 0)
	})
	t.Run("fuzzy-bytes", func(t *testing.T) {
		raw := FuzzyDataRawBytes()
		assert.NotNil(t, raw)
		assert.True(t, len(raw) > 0)
	})
	t.Run("contains", func(t *testing.T) {
		result := contains([]string{"a", "b"}, "a")
		assert.True(t, result)

		result2 := contains([]string{"a", "b"}, "c")
		assert.False(t, result2)
	})
	t.Run("is-dangerous", func(t *testing.T) {
		raw := isDangerous("")
		assert.NotNil(t, raw)
		assert.False(t, raw)
	})
}
