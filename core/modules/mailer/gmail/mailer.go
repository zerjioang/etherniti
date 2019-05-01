// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package gmail

import (
	"encoding/json"
	"net/mail"

	"github.com/zerjioang/etherniti/core/modules/mailer/model"
)

var (
	defaultGmailer *ApiEmailer
)

type ApiEmailer struct {
	maxLimit     int
	currentLimit int
	mailengine   *MailServerConfig
	from         *mail.Address
}

func NewApiEmailer() *ApiEmailer {
	em := new(ApiEmailer)
	// currently sent email amount
	em.currentLimit = 0
	//by default, google free email services only allow to send a maximum of 2000 emails per day
	em.maxLimit = 2000
	//default email sending address
	em.from = &mail.Address{Name: "Etherniti", Address: "noreply@etherniti.org"}
	//initialize default mail engine
	em.mailengine = GetMailServerConfigInstanceInit()
	return em
}

func init() {
	defaultGmailer = NewApiEmailer()
}

func SendWithGmail(data *model.Maildata) (json.RawMessage, error) {
	//send email
	err := defaultGmailer.mailengine.SendWithSSL(
		defaultGmailer.from,
		&mail.Address{Name: data.To.User, Address: data.To.Email},
		data.Subject,
		data.Htmltext,
	)
	return nil, err
}
