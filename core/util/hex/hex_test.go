// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package hex

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	t.Run("default-encode", func(t *testing.T) {
		result := hex.EncodeToString([]byte("this-is-a-test"))
		assert.Equal(t, result, "746869732d69732d612d74657374")
	})
	t.Run("fast-encode", func(t *testing.T) {
		result := UnsafeEncodeToString([]byte("this-is-a-test"))
		assert.Equal(t, result, "746869732d69732d612d74657374")
	})
	t.Run("fast-encode-pooled", func(t *testing.T) {
		result := UnsafeEncodeToStringPooled([]byte("this-is-a-test"))
		assert.Equal(t, result, "746869732d69732d612d74657374")
	})
}

func TestDecode(t *testing.T) {
	t.Run("default-decode", func(t *testing.T) {
		result, err := hex.DecodeString("746869732d69732d612d74657374")
		assert.Nil(t, err)
		assert.Equal(t, result, []byte("this-is-a-test"))
	})
	t.Run("fast-decode", func(t *testing.T) {
		result, err := UnsafeDecodeString("746869732d69732d612d74657374")
		assert.Nil(t, err)
		assert.Equal(t, result, []byte("this-is-a-test"))
	})
}