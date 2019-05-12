// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package profile

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	/*
		{
		  "endpoint": "http://127.0.0.1:8545",
		  "address": "0x0",
		  "key": "0x0",
		  "version": "",
		  "exp": 1557669973,
		  "jti": "ff4c4c4b-fb0c-45be-b277-9da1cf9aa5de",
		  "iat": 1557669373,
		  "iss": "etherniti",
		  "nbf": 1557669373,
		  "validity": false
		}
	*/
	testToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbmRwb2ludCI6Imh0dHA6Ly8xMjcuMC4wLjE6ODU0NSIsImFkZHJlc3MiOiIweDAiLCJrZXkiOiIweDAiLCJ2ZXJzaW9uIjoiIiwiZXhwIjoxNTU3NjY5OTczLCJqdGkiOiJmZjRjNGM0Yi1mYjBjLTQ1YmUtYjI3Ny05ZGExY2Y5YWE1ZGUiLCJpYXQiOjE1NTc2NjkzNzMsImlzcyI6ImV0aGVybml0aSIsIm5iZiI6MTU1NzY2OTM3MywidmFsaWRpdHkiOmZhbHNlfQ.ZxNG4ejAyJJ6Ipab8rVaLI1_texIl646lxovggplqpI"
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
		t.Run("parse-token", func(t *testing.T) {
			profile, err := ParseConnectionProfileToken(testToken)
			assert.Nil(t, err)
			assert.NotNil(t, profile)
			t.Log(profile)
		})
	})
}
