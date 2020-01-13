// Copyright go-phc (https://github.com/zerjioang/go-hpc)
// SPDX-License-Identifier: Apache License 2.0

package mail

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zerjioang/go-hpc/lib/mailer/gmail"
	"github.com/zerjioang/go-hpc/lib/mailer/model"
	"github.com/zerjioang/go-hpc/lib/mailer/sendgrid"
)

const (
	testUserName = "test"
	testEmail    = "test@etherniti.org"
)

func TestEmails(t *testing.T) {
	t.Run("registration-via-sendgrid", func(t *testing.T) {
		err := SendActivationEmail(&model.AuthMailRequest{
			Username: testUserName,
			Email:    testEmail,
		}, "", sendgrid.SendGridMailDelivery)
		assert.Nil(t, err)
	})
	t.Run("registration-via-gmail", func(t *testing.T) {
		err := SendActivationEmail(&model.AuthMailRequest{
			Username: testUserName,
			Email:    testEmail,
		}, "", gmail.SendWithGmail)
		assert.Nil(t, err)
	})
	t.Run("recovery-request", func(t *testing.T) {
		err := SendRecoveryEmail(&model.AuthMailRequest{
			Email: testEmail,
		}, "", sendgrid.SendGridMailDelivery)
		assert.Nil(t, err)
	})
	t.Run("login-detect", func(t *testing.T) {
		err := SendLoginEmail(&model.AuthMailRequest{
			Email: testEmail,
		}, "", sendgrid.SendGridMailDelivery)
		assert.Nil(t, err)
	})
}
