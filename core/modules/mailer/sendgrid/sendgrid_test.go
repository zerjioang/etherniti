// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package sendgrid

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zerjioang/etherniti/core/modules/mailer/model"
)

const (
	testdata = `{
  "personalizations": [
    {
      "to": [
        {
          "name": "User 01",
          "email": "user01@etherniti.org"
        }
      ],
      "subject": "Hello, World!"
    }
  ],
  "from": {
    "name": "Etherniti",
    "email": "noreply@etherniti.org"
  },
  "content": [
    {
      "type": "text/plain",
      "value": "Hello, World!"
    },
    {
      "type": "text/html",
      "value": "<h1>Hello, World!</h1"
    }
  ],
  "categories": ["noreply", "auth"]
}`
)

func TestSendGridSendEmail(t *testing.T) {
	t.Run("send-email", func(t *testing.T) {
		data := &model.Maildata{
			From:      model.MailAddress{User: "Etherniti", Email: "noreply@etherniti.org"},
			To:        model.MailAddress{User: "User 02", Email: "user02@etherniti.org"},
			Subject:   "SendGrid Mail Delivery Test",
			Plaintext: "this is a SendGrid Mail Delivery Test",
			Htmltext:  "<h1>this is a SendGrid Mail Delivery Test</h1>",
		}
		response, err := SendGridMailDelivery(data)
		assert.Nil(t, err)
		assert.NotNil(t, response)
	})
}
