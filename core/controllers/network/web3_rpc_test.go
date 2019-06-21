// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeb3RpcController(t *testing.T) {
	t.Run("instantiate", func(t *testing.T) {
		rpc := NewWeb3RpcController(nil)
		assert.NotNil(t, rpc)
	})
}
