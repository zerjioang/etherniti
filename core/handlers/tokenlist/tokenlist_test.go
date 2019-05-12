// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package tokenlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenList(t *testing.T) {
	t.Run("get-token-address-by-name", func(t *testing.T) {
		address := GetTokenAddressByName("$IQN")
		assert.Equal(t, address, "0x0db8d8b76bc361bacbb72e2c491e06085a97ab31")
	})
	t.Run("get-token-symbol-by-address", func(t *testing.T) {
		address := GetTokenSymbol("0x0db8d8b76bc361bacbb72e2c491e06085a97ab31")
		assert.Equal(t, address, "$IQN")
	})
}
