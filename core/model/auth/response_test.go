package auth

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponse(t *testing.T) {
	t.Run("instantiation", func(t *testing.T) {
		r := NewLoginResponse("foo-bar")
		assert.NotNil(t, r)
		assert.Equal(t, r.Token, "foo-bar")
	})
	t.Run("json", func(t *testing.T) {
		r := NewLoginResponse("foo-bar")
		assert.NotNil(t, r)

		r.Json()
		assert.Equal(t, string(r.Json()), `{"token":"foo-bar"}`)
	})
	t.Run("writer", func(t *testing.T) {
		t.Run("nil", func(t *testing.T) {
			r := NewLoginResponse("foo-bar")
			assert.NotNil(t, r)

			err := r.Writer(nil)
			assert.Nil(t, err)
		})
		t.Run("buffer-buffer", func(t *testing.T) {
			var buf bytes.Buffer
			r := NewLoginResponse("foo-bar")
			assert.NotNil(t, r)

			err := r.Writer(&buf)
			assert.Nil(t, err)

			assert.Equal(t, buf.String(), `{"token":"foo-bar"}`)
		})
	})
}
