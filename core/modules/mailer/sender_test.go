// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package mailer

import (
	"testing"

	"github.com/zerjioang/etherniti/core/model/auth"

	"github.com/zerjioang/etherniti/core/modules/mailer/gmail"
	"github.com/zerjioang/etherniti/core/modules/mailer/sendgrid"

	"github.com/stretchr/testify/assert"
)

const (
	testUserName = "test"
	testEmail    = "test@etherniti.org"
)

func TestEmails(t *testing.T) {
	t.Run("registration-via-sendgrid", func(t *testing.T) {
		err := SendActivationEmail(&auth.AuthRequest{
			Username: testUserName,
			Email:    testEmail,
		}, sendgrid.SendGridMailDelivery)
		assert.Nil(t, err)
	})
	t.Run("registration-via-gmail", func(t *testing.T) {
		err := SendActivationEmail(&auth.AuthRequest{
			Username: testUserName,
			Email:    testEmail,
		}, gmail.SendWithGmail)
		assert.Nil(t, err)
	})
	t.Run("recovery-request", func(t *testing.T) {
		err := SendRecoveryEmail(&auth.AuthRequest{
			Email: testEmail,
		}, sendgrid.SendGridMailDelivery)
		assert.Nil(t, err)
	})
	t.Run("login-detect", func(t *testing.T) {
		err := SendLoginEmail(&auth.AuthRequest{
			Email: testEmail,
		}, sendgrid.SendGridMailDelivery)
		assert.Nil(t, err)
	})
}
