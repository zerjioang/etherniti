// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package mailer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zerjioang/etherniti/shared/protocol"
)

const (
	testUserName = "test"
	testEmail    = "test@etherniti.org"
)

func TestEmails(t *testing.T) {
	t.Run("registration", func(t *testing.T) {
		mailer := NewApiEmailer()
		err := mailer.SendActivationEmail(&protocol.RegisterRequest{
			Username: testUserName,
			Email:    testEmail,
		})
		assert.Nil(t, err)
	})
	t.Run("recovery-request", func(t *testing.T) {
		mailer := NewApiEmailer()
		err := mailer.SendRecoveryEmail(&protocol.RecoveryRequest{
			Email: testEmail,
		})
		assert.Nil(t, err)
	})
	t.Run("login-detect", func(t *testing.T) {
		mailer := NewApiEmailer()
		err := mailer.SendLoginEmail(&protocol.LoginRequest{
			Email: testEmail,
		})
		assert.Nil(t, err)
	})
}
