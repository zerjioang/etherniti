// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package protocol

import (
	"github.com/zerjioang/etherniti/core/util/str"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewApiError(t *testing.T) {
	t.Run("instantiate", func(t *testing.T) {
		apiErr := NewApiError(200, str.UnsafeBytes("test-trycatch"))
		require.NotNil(t, apiErr)
		require.Equal(t, apiErr.Code, 200)
		require.Equal(t, apiErr.Details, "test-trycatch")
	})
}

func TestNewApiResponse(t *testing.T) {
	t.Run("instantiate", func(t *testing.T) {
		msg := NewApiResponse(str.UnsafeBytes("success"), 12345)
		require.NotNil(t, msg)
		require.Equal(t, msg.Code, 200)
		require.Equal(t, msg.Message, "success")
		require.Equal(t, msg.Result, 12345)
	})
}
