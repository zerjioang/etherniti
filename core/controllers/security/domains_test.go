// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package security

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDomainBlacklist(t *testing.T) {
	t.Run("get-domain-blacklist", func(t *testing.T) {
		list := DomainBlacklist()
		assert.NotNil(t, list)
		assert.True(t, len(list) > 0)
	})
	t.Run("get-domain-blacklist-bytes", func(t *testing.T) {
		raw := DomainBlacklistBytesData()
		assert.NotNil(t, raw)
		assert.True(t, len(raw) > 0)
	})
}
