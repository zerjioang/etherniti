// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

// +build !dev
// +build pre
// +build !prod

package config

import (
	"time"

	"github.com/zerjioang/etherniti/core/util/str"

	"github.com/zerjioang/etherniti/core/eth/fastime"
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
)

func init() {
	//hardcoded cert content as bytes
	certPemBytes = str.UnsafeBytes(certPem)
	//hardcoded key content as bytes
	keyPemBytes = str.UnsafeBytes(keyPem)
}

//set default environment variables value for current context
func (c *EnvConfig) SetDefaults() {
	env := c.data
	env["X_ETHERNITI_LOG_LEVEL"] = "debug"
	env["X_ETHERNITI_ENVIRONMENT_NAME"] = "beta-stage"
	env["X_ETHERNITI_HTTP_PORT"] = "8080"
	env["X_ETHERNITI_HTTPS_PORT"] = "4430"
	env["X_ETHERNITI_DEBUG_SERVER"] = false
	env["X_ETHERNITI_HIDE_SERVER_DATA_IN_CONSOLE"] = true
	env["X_ETHERNITI_TOKEN_SECRET"] = "t0k3n-s3cr3t-h3r3"
	env["X_ETHERNITI_ENABLE_HTTPS_REDIRECT"] = false
	env["X_ETHERNITI_USE_UNIQUE_REQUEST_ID"] = false
	env["X_ETHERNITI_ENABLE_CORS"] = true
	env["X_ETHERNITI_ENABLE_SECURITY"] = true
	env["X_ETHERNITI_ENABLE_ANALYTICS"] = true
	env["X_ETHERNITI_ENABLE_METRICS"] = true
	env["X_ETHERNITI_ENABLE_CACHE"] = true
	env["X_ETHERNITI_ENABLE_RATELIMIT"] = false
	env["X_ETHERNITI_ENABLE_PROFILER"] = false
	env["X_ETHERNITI_BLOCK_TOR_CONNECTIONS"] = false
	env["X_ETHERNITI_ENABLE_LOGGING"] = true

	//for 'local development' deployment
	env["X_ETHERNITI_LISTENING_MODE"] = "http" // http or socket
	env["X_ETHERNITI_LISTENING_INTERFACE"] = "0.0.0.0"
	env["X_ETHERNITI_LISTENING_ADDRESS"] = env["X_ETHERNITI_LISTENING_INTERFACE"].(string) + ":" + env["X_ETHERNITI_HTTP_PORT"].(string)
	env["X_ETHERNITI_SWAGGER_ADDRESS"] = "proxy.etherniti.org"

	//connection profile params
	env["X_ETHERNITI_TOKEN_EXPIRATION"] = 10 * fastime.Minute

	//rate limit units must be the same in both variables
	//rate limit units must be the same in both variables
	env["X_ETHERNITI_RATE_LIMIT_UNITS"] = 5 * time.Second
	env["X_ETHERNITI_RATE_LIMIT_UNITS_FT"] = 5 * fastime.Second
	env["X_ETHERNITI_RATE_LIMIT"] = 10
	env["X_ETHERNITI_RATE_LIMIT_STR"] = "10"
}

// setup server config
func Setup() {
}
