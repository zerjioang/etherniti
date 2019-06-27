package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponse(t *testing.T) {
	t.Run("instantiation", func(t *testing.T) {
		r := NewLoginResponse("foo-bar")
		assert.NotNil(t, r)
		assert.Equal(t, r.Token, "foo-bar")
	})
}
