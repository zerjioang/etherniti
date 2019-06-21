// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package sendgrid

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/zerjioang/etherniti/core/util/str"

	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/modules/httpclient"
	"github.com/zerjioang/etherniti/core/modules/mailer/model"
)

/*
Sendgrid API email sender
More info at: https://sendgrid.com/docs/API_Reference/Web_API_v3/Mail/index.html

curl --request POST \
--url https://api.sendgrid.com/v3/mail/send \
--header "Authorization: Bearer $SENDGRID_API_KEY" \
--header 'Content-Type: application/json' \
--data '{
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
      "value": "Hello, World!"
    }
  ],
  "categories": ["noreply", "auth"]
}
'

*/

const (
	sendgridUrl = "https://api.sendgrid.com/v3/mail/send"
	payloadBody = `{
  "personalizations": [
    {
      "to": [
        {
          "name": "{{user}}",
          "email": "{{email}}"
        }
      ],
      "subject": "{{subject}}"
    }
  ],
  "from": {
    "name": "Etherniti",
    "email": "noreply@etherniti.org"
  },
  "content": [
    {
      "type": "text/plain",
      "value": "{{plaintext}}"
    },
    {
      "type": "text/html",
      "value": "{{htmltext}}"
    }
  ],
  "categories": ["noreply", "auth"]
}`
)

var (
	cfg                      = config.GetDefaultOpts()
	apiKey                   = ""
	defaultRequestHeader     http.Header
	defaultSendGridApiClient *http.Client
	noApiKeyErr              = errors.New("no SENDGRID_API_KEY was defined")
	spaceRemover             = regexp.MustCompile(`\s+`)
)

func init() {
	logger.Debug("reading SENDGRID_API_KEY")
	apiKey = cfg.SendGridApiKey()
	//generate header used in all sendgrid request
	defaultRequestHeader = http.Header{
		"Content-Type":  []string{httpclient.ApplicationJson},
		"Authorization": []string{"Bearer " + apiKey},
	}
	logger.Debug("creating sendgrid api client")
	defaultSendGridApiClient = &http.Client{
		Timeout: time.Second * 3,
		Transport: &http.Transport{
			TLSHandshakeTimeout: 3 * time.Second,
		},
	}
}

func buildSendGridPayload(maildata *model.Maildata) string {
	logger.Debug("building email payload data in sendgrid format")
	current := payloadBody
	current = strings.Replace(current, "{{user}}", maildata.To.User, 1)
	current = strings.Replace(current, "{{email}}", maildata.To.Email, 1)
	current = strings.Replace(current, "{{subject}}", maildata.Subject, 1)
	current = strings.Replace(current, "{{plaintext}}", maildata.Plaintext, 1)
	// we need to clean html tex. remove new lines
	maildata.Htmltext = cleanHtml(maildata.Htmltext)
	current = strings.Replace(current, "{{htmltext}}", maildata.Htmltext, 1)
	return current
}
func SendGridMailDelivery(data *model.Maildata) (json.RawMessage, error) {
	if apiKey == "" {
		logger.Error("aborting email delivery because no api key was defined in current environment variables")
		return nil, noApiKeyErr
	} else {
		emailStr := buildSendGridPayload(data)
		logger.Debug("sending email via sendgrid api")
		return httpclient.MakePost(defaultSendGridApiClient, sendgridUrl, defaultRequestHeader, str.UnsafeBytes(emailStr))
	}
}

func cleanHtml(html string) string {
	html = strings.Replace(html, "\t", "", -1)
	html = strings.Replace(html, "\r", "", -1)
	html = strings.Replace(html, "\n", "", -1)
	html = spaceRemover.ReplaceAllString(html, " ")
	html = strings.Replace(html, "> <", "><", -1)
	html = strings.Replace(html, "\"", "'", -1)
	return html
}
