// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package mailer

import (
	"net/mail"
	"strings"

	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/util/str"
	"github.com/zerjioang/etherniti/shared/protocol"
)

type ApiEmailer struct {
	maxLimit     int
	currentLimit int
	mailengine   *MailServerConfig
	from         *mail.Address
}

var (
	confirmEmailTemplate  string
	newLoginEmailTemplate string
	recoverEmailTemplate  string
)

func init() {
	logger.Debug("loading email templates...")
	confirmEmailTemplate = str.ReadFileAsString(config.ResourcesDirInternalEmail + "/confim.html")
	newLoginEmailTemplate = str.ReadFileAsString(config.ResourcesDirInternalEmail + "/confim.html")
	recoverEmailTemplate = str.ReadFileAsString(config.ResourcesDirInternalEmail + "/confim.html")
}

func NewApiEmailer() *ApiEmailer {
	em := new(ApiEmailer)
	// currently sent email amount
	em.currentLimit = 0
	//by default, google free email services only allow to send a maximum of 2000 emails per day
	em.maxLimit = 2000
	//default email sending address
	em.from = &mail.Address{Name: "Etherniti Project", Address: "noreply@etherniti.org"}
	//initialize default mail engine
	em.mailengine = GetMailServerConfigInstanceInit()
	return em
}

//todo add google markup
//guide: https://developers.google.com/gmail/markup/registering-with-google
//examples. https://developers.google.com/gmail/markup/reference/one-click-action
func (m *ApiEmailer) SendActivationEmail(registerRequest *protocol.RegisterRequest) error {

	//generate target email address
	targetUserEmailAddress := &mail.Address{Name: registerRequest.Username, Address: registerRequest.Email}

	//generate email body
	emailBody := confirmEmailTemplate

	//common for all emails
	emailBody = strings.Replace(emailBody, "{{appname}}", "Etherniti", -1)
	emailBody = strings.Replace(emailBody, "{{domain}}", "www.etherniti.org", -1)
	emailBody = strings.Replace(emailBody, "{{location}}", "Bilbao, Basque Country", -1)

	emailBody = strings.Replace(emailBody, "{{applogo}}", "https://avatars3.githubusercontent.com/u/47393730?s=30&v=4", -1)
	emailBody = strings.Replace(emailBody, "{{headerlogo}}", "https://user-images.githubusercontent.com/6706342/56867116-de523b80-69e1-11e9-86f5-1aa694aeed3f.jpg", -1)
	emailBody = strings.Replace(emailBody, "{{founderlogo}}", "https://avatars3.githubusercontent.com/u/6706342?s=60&v=4", -1)
	emailBody = strings.Replace(emailBody, "{{footnote}}", `"Cause real live is far away from movies...or not so much!"`, -1)

	emailBody = strings.Replace(emailBody, "{{title}}", "Etherniti - Account activation", -1)
	emailBody = strings.Replace(emailBody, "{{email_hash}}", registerRequest.Email, -1)

	emailBody = strings.Replace(emailBody, "{{unsubscribe_url}}", "https://cloud.etherniti.org/unsuscribe/{{email_hash}}", -1)
	emailBody = strings.Replace(emailBody, "{{confirm_url}}", "https://cloud.etherniti.org/auth/confirm/registration/123", -1)

	emailBody = strings.Replace(emailBody, "{{username}}", registerRequest.Username, -1)
	emailBody = strings.Replace(emailBody, "{{app_homepage}}", "https://dashboard.etherniti.org", -1)

	//security simple
	emailBody = strings.Replace(emailBody, "{{pgp}}", "406193B0B0639ECD19C42E92BE3FCCB5AD67564C1A371C34A69D11F380E7FB6B", -1)

	//send email
	err := m.mailengine.SendWithSSL(
		m.from,
		targetUserEmailAddress,
		"Welcome to Etherniti. Activate your account",
		emailBody,
	)
	return err
}

func (m *ApiEmailer) SendLoginEmail(loginRequest *protocol.LoginRequest) error {

	//generate target email address
	targetUserEmailAddress := &mail.Address{Name: "", Address: loginRequest.Email}

	//generate email body
	emailBody := newLoginEmailTemplate

	//common for all emails
	emailBody = strings.Replace(emailBody, "{{title}}", "Etherniti - Account activation", -1)
	emailBody = strings.Replace(emailBody, "{{unsubscribe_url}}", "https://cloud.etherniti.org/unsuscribe/{{email_hash}}", -1)
	emailBody = strings.Replace(emailBody, "{{email_hash}}", loginRequest.Email, -1)

	emailBody = strings.Replace(emailBody, "{{username}}", "aaa", -1)
	emailBody = strings.Replace(emailBody, "{{time}}", "aaa", -1)
	emailBody = strings.Replace(emailBody, "{{location}}", "aaa", -1)
	emailBody = strings.Replace(emailBody, "{{device}}", "aaa", -1)
	emailBody = strings.Replace(emailBody, "{{ip}}", "aaa", -1)
	emailBody = strings.Replace(emailBody, "{{lockdown_url}}", "aaa", -1)

	//security simple
	emailBody = strings.Replace(emailBody, "{{pgp}}", "406193B0B0639ECD19C42E92BE3FCCB5AD67564C1A371C34A69D11F380E7FB6B", -1)

	//send email
	err := m.mailengine.SendWithSSL(
		m.from,
		targetUserEmailAddress,
		"A new login was detect to your Etherniti account",
		emailBody,
	)
	return err
}

//todo add google markup
//guide: https://developers.google.com/gmail/markup/registering-with-google
//examples. https://developers.google.com/gmail/markup/reference/one-click-action
func (m *ApiEmailer) SendRecoveryEmail(recoveryRequest *protocol.RecoveryRequest) error {

	//generate target email address
	targetUserEmailAddress := &mail.Address{Address: recoveryRequest.Email}

	//generate email body
	emailBody := recoverEmailTemplate

	//common for all emails
	emailBody = strings.Replace(emailBody, "{{title}}", "Etherniti - Account recovery instructions", -1)
	emailBody = strings.Replace(emailBody, "{{unsubscribe_url}}", "https://cloud.etherniti.org/unsuscribe/{{email_hash}}", -1)
	emailBody = strings.Replace(emailBody, "{{email_hash}}", recoveryRequest.Email, -1)

	emailBody = strings.Replace(emailBody, "{{username}}", "aaa", -1)
	emailBody = strings.Replace(emailBody, "{{time}}", "aaa", -1)
	emailBody = strings.Replace(emailBody, "{{location}}", "aaa", -1)
	emailBody = strings.Replace(emailBody, "{{device}}", "aaa", -1)
	emailBody = strings.Replace(emailBody, "{{ip}}", "aaa", -1)
	emailBody = strings.Replace(emailBody, "{{lockdown_url}}", "aaa", -1)
	emailBody = strings.Replace(emailBody, "{{recovery_url}}", "aaa", -1)

	//security simple
	emailBody = strings.Replace(emailBody, "{{pgp}}", "406193B0B0639ECD19C42E92BE3FCCB5AD67564C1A371C34A69D11F380E7FB6B", -1)

	//send email
	err := m.mailengine.SendWithSSL(
		m.from,
		targetUserEmailAddress,
		"Etherniti account recovery instructions",
		emailBody,
	)
	return err
}
