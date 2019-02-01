package config

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetRedirectUrl(t *testing.T) {
	t.Run("redirect", func(t *testing.T) {
		redirectUrl := GetRedirectUrl("localhost", "/test")
		require.Equal(t, redirectUrl, "https://localhost/test")
	})
}
