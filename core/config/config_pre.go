// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

// +build pre

package config

import (
	"github.com/zerjioang/etherniti/core/logger"

	"github.com/zerjioang/etherniti/core/util/str"
)

// openssl genrsa -out server.key 2048
// openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
// Country name (2 letter code) [AU]:ES
// State or Province name (full name) [Some-State]:Biscay
// Locality name (eg, city) []:Bilbao
// Organization name (eg, company) [Internet Widgits Pty Ltd]:Etherniti Project
// Organizational Unit name (eg, section) []:Etherniti CyberSecurity Team
// Common name (e.g. server FQDN or YOUR name) []:localhost
// Email Address []:

const (
	certPem = `-----BEGIN CERTIFICATE-----
MIICkjCCAhigAwIBAgIJANdeA9flJMlnMAoGCCqGSM49BAMCMIGGMQswCQYDVQQG
EwJFUzEPMA0GA1UECAwGQmlzY2F5MQ8wDQYDVQQHDAZCaWxiYW8xGjAYBgNVBAoM
EUV0aGVybml0aSBQcm9qZWN0MSUwIwYDVQQLDBxFdGhlcm5pdGkgQ3liZXJTZWN1
cml0eSBUZWFtMRIwEAYDVQQDDAlsb2NhbGhvc3QwHhcNMTkwMjI1MTAzNjAzWhcN
MjkwMjIyMTAzNjAzWjCBhjELMAkGA1UEBhMCRVMxDzANBgNVBAgMBkJpc2NheTEP
MA0GA1UEBwwGQmlsYmFvMRowGAYDVQQKDBFFdGhlcm5pdGkgUHJvamVjdDElMCMG
A1UECwwcRXRoZXJuaXRpIEN5YmVyU2VjdXJpdHkgVGVhbTESMBAGA1UEAwwJbG9j
YWxob3N0MHYwEAYHKoZIzj0CAQYFK4EEACIDYgAEdfDTiXG01evJo+H8iELSM5rB
u73ZODqPMpPyqxTXA8/r+juHqs+65USA+VsCk5wMaTWxTq4nMTonOu3+zqkySa1F
/uOp7HBbTGclLreYiRn0tsjKML8Hvoj8sHPm/Wdzo1AwTjAdBgNVHQ4EFgQU4LQl
3Q8zGmnYBQicOBloJirneBgwHwYDVR0jBBgwFoAU4LQl3Q8zGmnYBQicOBloJirn
eBgwDAYDVR0TBAUwAwEB/zAKBggqhkjOPQQDAgNoADBlAjEAmXDYP8TNLczRAmoq
5ijranWuzCXD4vJZFs84XO4/J/sh5Pz+TcCZFFChAODmuWd5AjB/PgnS1lMBJsEY
MfAwQl1+hKBNFvv0i5fsIM00QSgK/Eys3wfWf4nROAH4V/T+T98=
-----END CERTIFICATE-----`

	keyPem = `-----BEGIN EC PARAMETERS-----
BgUrgQQAIg==
-----END EC PARAMETERS-----
-----BEGIN EC PRIVATE KEY-----
MIGkAgEBBDBZL5r9/Rkt1T7RP86URs4HnRtJlfmN24mRgrxYNmiw09hNvOau6r4v
N9OXUhEAg3SgBwYFK4EEACKhZANiAAR18NOJcbTV68mj4fyIQtIzmsG7vdk4Oo8y
k/KrFNcDz+v6O4eqz7rlRID5WwKTnAxpNbFOricxOic67f7OqTJJrUX+46nscFtM
ZyUut5iJGfS2yMowvwe+iPywc+b9Z3M=
-----END EC PRIVATE KEY-----
`
	environmentName = "staging"
)

func init() {
	logger.Debug("loading staging config module data")
	//hardcoded cert content as bytes
	certPemBytes = str.UnsafeBytes(certPem)
	//hardcoded key content as bytes
	keyPemBytes = str.UnsafeBytes(keyPem)
}

//check if profiling is enabled or not
// preproduction and production profiling
// is always disabled
func IsProfilingEnabled() bool {
	return false
}

func IsDevelopment() bool {
	logger.Debug("checking if current server environment is development")
	return false
}

// read swagger url domain: ip or FQDN
// no port allowed in pre and prod. default port is used
// 80 and 443
func GetSwaggerAddressWithPort(opts EthernitiOptions) string {
	logger.Debug("reading swagger address port from env")
	return opts.SwaggerAddress
}

func Env() string {
	logger.Debug("reading server environment name")
	return environmentName
}

// setup server config
func Setup(opts *EthernitiOptions) error {
	logger.Debug("loading additional staging setup config")
	return nil
}
