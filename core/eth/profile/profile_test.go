// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package profile

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	/*
		example token
	*/
	testToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
)

func TestCreateConnectionProfileToken(t *testing.T) {
	t.Run("empty-connection-profile", func(t *testing.T) {
		NewConnectionProfile()
	})
	t.Run("connection-profile", func(t *testing.T) {
		_ = NewDefaultConnectionProfile()
	})
	t.Run("create-token", func(t *testing.T) {
		p := NewDefaultConnectionProfile()
		token, err := CreateConnectionProfileToken(p)
		assert.Nil(t, err)
		assert.NotNil(t, token)
		t.Log(token)
	})
	t.Run("parse-token", func(t *testing.T) {
		t.Run("parse-empty", func(t *testing.T) {
			_, err := ParseConnectionProfileToken("")
			assert.NotNil(t, err)
		})
	})
}
