// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetRedirectUrl(t *testing.T) {
	t.Run("redirect", func(t *testing.T) {
		redirectUrl := GetRedirectUrl("localhost", "/test")
		require.Equal(t, redirectUrl, "https://localhost/test")
	})
}
