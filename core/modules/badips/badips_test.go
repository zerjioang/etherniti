// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package badips

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBadIps(t *testing.T) {
	t.Run("get-list", func(t *testing.T) {
		l := GetBadIPList()
		assert.NotNil(t, l)
	})
	t.Run("contains-true", func(t *testing.T) {
		l := GetBadIPList()
		result := l.Contains("31.6.220.31")
		assert.True(t, result)
	})
	t.Run("contains-false", func(t *testing.T) {
		l := GetBadIPList()
		result := l.Contains("127.0.0.1")
		assert.False(t, result)
	})
}
