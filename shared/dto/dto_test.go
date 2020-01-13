// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package dto

import (
	"testing"

	"github.com/zerjioang/go-hpc/util/str"

	"github.com/stretchr/testify/require"
)

func TestNewApiError(t *testing.T) {
	t.Run("instantiate", func(t *testing.T) {
		apiErr := NewApiError(200, "test-stack")
		require.NotNil(t, apiErr)
		require.Equal(t, string(apiErr.Err), "test-stack")
	})
}

func TestNewApiResponse(t *testing.T) {
	t.Run("instantiate", func(t *testing.T) {
		msg := NewApiResponse(str.UnsafeBytes("success"), []byte("foo-bar"))
		require.NotNil(t, msg)
		require.Equal(t, string(msg.Message), "success")
		require.Equal(t, msg.Data, []byte("foo-bar"))
	})
}
