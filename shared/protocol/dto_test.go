// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package protocol

import (
	"testing"

	"github.com/zerjioang/etherniti/core/util/str"

	"github.com/stretchr/testify/require"
)

func TestNewApiError(t *testing.T) {
	t.Run("instantiate", func(t *testing.T) {
		apiErr := NewApiError(200, str.UnsafeBytes("test-stack"))
		require.NotNil(t, apiErr)
		require.Equal(t, apiErr.Data, "test-stack")
	})
}

func TestNewApiResponse(t *testing.T) {
	t.Run("instantiate", func(t *testing.T) {
		msg := NewApiResponse(str.UnsafeBytes("success"), nil)
		require.NotNil(t, msg)
		require.Equal(t, msg.Message, "success")
		require.Equal(t, msg.Data, 12345)
	})
}
