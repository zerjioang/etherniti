// Copyright go-phc (https://github.com/zerjioang/go-hpc)
// SPDX-License-Identifier: Apache License 2.0

package mail

import (
	"errors"
	"strings"

	"github.com/zerjioang/go-hpc/lib/mailer/model"
)

var (
	errNoDeliveryMethodSet = errors.New("failed to deliver message to the client because no delivery method was set")
)

//todo add google markup
//guide: https://developers.google.com/gmail/markup/registering-with-google
//examples. https://developers.google.com/gmail/markup/reference/one-click-action
func SendActivationEmail(registerRequest *model.AuthMailRequest, emailBody string, sender model.EmailSenderMechanism) error {

	//generate target email address
	targetUserEmailAddress := model.MailAddress{User: registerRequest.Username, Email: registerRequest.Email}

	emailBody = strings.Replace(emailBody, "{{title}}", "Welcome to {{appname}}", -1)
	emailBody = strings.Replace(emailBody, "{{msgtitle}}", "Welcome to {{appname}}", -1)
	emailBody = strings.Replace(emailBody, "{{headerlogo}}", "https://user-images.githubusercontent.com/6706342/56867116-de523b80-69e1-11e9-86f5-1aa694aeed3f.jpg", -1)
	emailBody = strings.Replace(emailBody, "{{founderlogo}}", "https://avatars3.githubusercontent.com/u/6706342?s=60&v=4", -1)
	emailBody = strings.Replace(emailBody, "{{footnote}}", `"Cause real live is far away from movies...or not so much!"`, -1)
	emailBody = strings.Replace(emailBody, "{{title}}", "Etherniti - Account activation", -1)
	emailBody = strings.Replace(emailBody, "{{email_hash}}", registerRequest.Email, -1)
	emailBody = strings.Replace(emailBody, "{{unsubscribe_url}}", "https://cloud.etherniti.org/unsuscribe/{{email_hash}}", -1)
	emailBody = strings.Replace(emailBody, "{{confirm_url}}", "https://cloud.etherniti.org/auth/confirm/{{confirm_code}}", 2)
	emailBody = strings.Replace(emailBody, "{{confirm_code}}", registerRequest.Confirmation, -1)
	emailBody = strings.Replace(emailBody, "{{username}}", registerRequest.Username, -1)

	//common for all emails
	emailBody = applyDefaults(emailBody)

	//generate email data object
	data := &model.Maildata{
		From:      model.MailAddress{User: "Etherniti", Email: "noreply@etherniti.org"},
		To:        targetUserEmailAddress,
		Subject:   "Etherniti - Account activation",
		Plaintext: "Etherniti - Account activation",
		Htmltext:  emailBody,
	}
	if sender != nil {
		// send the email
		_, err := sender(data)
		return err
	}
	return errNoDeliveryMethodSet
}
func applyDefaults(emailBody string) string {
	emailBody = strings.Replace(emailBody, "{{app_homepage}}", "https://dashboard.etherniti.org", -1)
	emailBody = strings.Replace(emailBody, "{{appname}}", "Etherniti", -1)
	emailBody = strings.Replace(emailBody, "{{appicon}}", "https://avatars3.githubusercontent.com/u/47393730?s=30&v=4", -1)

	//security simple
	emailBody = strings.Replace(emailBody, "{{pgp}}", "406193B0B0639ECD19C42E92BE3FCCB5AD67564C1A371C34A69D11F380E7FB6B", -1)
	emailBody = strings.Replace(emailBody, "{{appname}}", "Etherniti", -1)
	emailBody = strings.Replace(emailBody, "{{domain}}", "www.etherniti.org", -1)
	emailBody = strings.Replace(emailBody, "{{applocation}}", "Bilbao, Basque Country", -1)
	return emailBody
}

func SendLoginEmail(loginRequest *model.AuthMailRequest, emailBody string, sender model.EmailSenderMechanism) error {

	//generate target email address
	targetUserEmailAddress := model.MailAddress{User: "", Email: loginRequest.Email}

	emailBody = strings.Replace(emailBody, "{{title}}", "Etherniti - New Login", -1)
	emailBody = strings.Replace(emailBody, "{{unsubscribe_url}}", "https://cloud.etherniti.org/unsuscribe/{{email_hash}}", -1)
	emailBody = strings.Replace(emailBody, "{{email_hash}}", loginRequest.Email, -1)

	emailBody = strings.Replace(emailBody, "{{username}}", "aaa", -1)
	emailBody = strings.Replace(emailBody, "{{time}}", "aaa", -1)
	emailBody = strings.Replace(emailBody, "{{location}}", "aaa", -1)
	emailBody = strings.Replace(emailBody, "{{device}}", "aaa", -1)
	emailBody = strings.Replace(emailBody, "{{ip}}", "aaa", -1)
	emailBody = strings.Replace(emailBody, "{{lockdown_url}}", "aaa", -1)

	//common for all emails
	emailBody = applyDefaults(emailBody)

	//generate email data object
	data := &model.Maildata{
		From:      model.MailAddress{User: "Etherniti", Email: "noreply@etherniti.org"},
		To:        targetUserEmailAddress,
		Subject:   "Etherniti - New Login",
		Plaintext: "A new login was detect to your Etherniti account",
		Htmltext:  emailBody,
	}
	if sender != nil {
		// send the email
		_, err := sender(data)
		return err
	}
	return errNoDeliveryMethodSet
}

//todo add google markup
//guide: https://developers.google.com/gmail/markup/registering-with-google
//examples. https://developers.google.com/gmail/markup/reference/one-click-action
func SendRecoveryEmail(recoveryRequest *model.AuthMailRequest, emailBody string, sender model.EmailSenderMechanism) error {

	//generate target email address
	targetUserEmailAddress := model.MailAddress{Email: recoveryRequest.Email}

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

	//common for all emails
	emailBody = applyDefaults(emailBody)

	//generate email data object
	data := &model.Maildata{
		From:      model.MailAddress{User: "Etherniti", Email: "noreply@etherniti.org"},
		To:        targetUserEmailAddress,
		Subject:   "Etherniti - Account recovery instructions",
		Plaintext: "Etherniti - Account recovery instructions",
		Htmltext:  emailBody,
	}
	if sender != nil {
		// send the email
		_, err := sender(data)
		return err
	}
	return errNoDeliveryMethodSet
}
