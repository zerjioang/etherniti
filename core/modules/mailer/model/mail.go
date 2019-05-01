// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package model

import "encoding/json"

// current supported methods are:
// * SendGridMailDelivery
// * GmailDelivery
type EmailSenderMecanism func(maildata *Maildata) (json.RawMessage, error)

type MailAddress struct {
	User  string
	Email string
}

type Maildata struct {
	From      MailAddress
	To        MailAddress
	Cc        []MailAddress
	Bcc       []MailAddress
	Subject   string
	Plaintext string
	Htmltext  string
}
