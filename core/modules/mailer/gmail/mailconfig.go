// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package gmail

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/mail"
	"net/smtp"
	"strings"
	"sync"
	"time"

	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/logger"
)

type MailServerConfig struct {
	username    string
	password    string
	emailServer string
	auth        smtp.Auth
}

var (
	packageInstance           *MailServerConfig
	packageInstanceThreadSafe *MailServerConfig
	packageInitInstance       *MailServerConfig
	once                      sync.Once
	opts                      = config.GetDefaultOpts()
)

/*
thread safe initialization. it only writes content to variable packageInitInstance once and
since this variable cannot be accesses by other code outside this file, it is completely thread safe
*/
func init() {
	packageInitInstance = newInternalServerConfig(opts)
}

func GetMailServerConfigInstanceInit() *MailServerConfig {
	return packageInitInstance
}

func newInternalServerConfig(opts *config.EthernitiOptions) *MailServerConfig {
	conf := new(MailServerConfig)
	conf.username = opts.GetEmailUsername()
	conf.password = opts.GetEmailPassword()
	conf.emailServer = opts.GetEmailServer()
	conf.auth = smtp.PlainAuth("",
		conf.username,
		conf.password,
		opts.GetEmailServerOnly(),
	)
	return conf
}

/*
@deprecated use GetMailServerConfigInstanceInit() instead

warning: this is a non-thread safe implementation of singleton.
use at your own risk knowning variable io access
*/
func GetMailServerConfigInstance() *MailServerConfig {
	//not thread safe code
	if packageInstance == nil {
		packageInstance = newInternalServerConfig(opts)
	}
	return packageInstance
}

/*
@deprecated use GetMailServerConfigInstanceInit() instead
*/
func GetMailServerConfigInstanceThreadSafe() *MailServerConfig {
	//thread safe code
	once.Do(func() {
		packageInstanceThreadSafe = newInternalServerConfig(opts)
	})

	return packageInstanceThreadSafe
}

func (c *MailServerConfig) SendWithSSL(fromMail, toMail *mail.Address, subject, body string) error {

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = fromMail.String()
	headers["To"] = toMail.String()
	headers["Subject"] = subject

	headers["MIME-Version"] = "1.0"
	headers["X-Send-Timestamp"] = fmt.Sprintf(time.Now().Format(time.RFC850))
	headers["Content-Type"] = "text/html;charset=UTF-8"

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Connect to the SMTP Server
	servername := "smtp.gmail.com:465"

	host, _, _ := net.SplitHostPort(servername)

	auth := smtp.PlainAuth("",
		c.username,
		c.password,
		host,
	)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	conn, dialErr := tls.Dial("tcp", servername, tlsconfig)
	if dialErr != nil {
		logger.Debug("failed to dial tcp email server", dialErr)
		return dialErr
	}

	emailClient, clientErr := smtp.NewClient(conn, host)
	if clientErr != nil {
		logger.Error("failed to create a client connection", clientErr)
		return clientErr
	}

	// Auth
	if err := emailClient.Auth(auth); err != nil {
		logger.Error("failed to authenticate against email server", err)
		return err
	}

	// To && From
	if err := emailClient.Mail(fromMail.Address); err != nil {
		logger.Debug("failed to send email", err)
		return err
	}

	if err := emailClient.Rcpt(toMail.Address); err != nil {
		logger.Error("failed to rcpt email", err)
		return err
	}

	// Data
	w, err := emailClient.Data()
	if err != nil {
		logger.Error("failed to get email data", err)
		return err
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		logger.Error("failed to write email message", err)
		return err
	}

	err = w.Close()
	if err != nil {
		logger.Error("failed to close email connection", err)
		return err
	}

	emailClient.Quit()
	return nil
}

func (c *MailServerConfig) SendInsecure(fromMail, toMail *mail.Address, subject, msg string) error {

	headers := make(map[string]string)
	headers["From"] = fromMail.String()
	headers["To"] = toMail.String()
	headers["Subject"] = subject

	headers["MIME-Version"] = "1.0"
	headers["X-Send-Timestamp"] = fmt.Sprintf(time.Now().Format(time.RFC850))
	headers["Content-Type"] = "text/html;charset=UTF-8"

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + msg //+ base64.StdEncoding.EncodeToString([]byte(msg))

	err := smtp.SendMail(c.emailServer, c.auth, fromMail.Address, []string{toMail.Address}, []byte(message))
	if err != nil {
		logger.Error("An error occurred while sending email: %s" + err.Error())
	}
	return err
}

func encodeRFC2047(title string) string {
	// use mail's rfc2047 to encode any string
	addr := mail.Address{Name: title}
	encoded := strings.Trim(addr.String(), " <@>")
	return encoded
}
