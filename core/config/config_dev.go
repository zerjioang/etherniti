// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

// +build dev !dev
// +build !pre
// +build !prod

package config

import (
	"net/http"
	"runtime"

	"github.com/zerjioang/etherniti/core/config/edition"

	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/util/str"

	_ "net/http/pprof" //adds 2,5Mb to final executable when imported

	"github.com/zerjioang/etherniti/thirdparty/gommon/log"
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
	logger.Debug("loading development config module data")
	//hardcoded cert content as bytes
	certPemBytes = str.UnsafeBytes(certPem)
	//hardcoded key content as bytes
	keyPemBytes = str.UnsafeBytes(keyPem)
}

//check if profiling is enabled or not
func IsProfilingEnabled(opts *EthernitiOptions) bool {
	logger.Debug("checking if profiling ListeningMode is enabled")
	v, found := opts.envData.Read("X_ETHERNITI_ENABLE_PROFILER")
	return found && v == "true"
}

func IsDevelopment() bool {
	logger.Debug("checking if current server environment is development")
	return true
}

// allow swagger ui access via ip and port or domain and port.
// development only
func GetSwaggerAddressWithPort(opts *EthernitiOptions) string {
	logger.Debug("reading swagger address with port from env")
	return opts.SwaggerAddress + ":" + opts.GetListeningPortStr()
}

func Env() string {
	logger.Debug("reading server environment name")
	return "development"
}

// setup server config
func Setup(opts *EthernitiOptions) error {
	logger.Debug("loading additional development setup config")
	// enable profile ListeningMode if requested
	if edition.IsEnterprise() && IsProfilingEnabled(opts) {
		go runProfiler()
	}
	return nil
}

// There are 7 places you can get profiles in the default webserver: the ones mentioned above
//
// http://localhost:6060/debug/pprof/
// http://localhost:6060/debug/pprof/goroutine
// http://localhost:6060/debug/pprof/heap
// http://localhost:6060/debug/pprof/threadcreate
// http://localhost:6060/debug/pprof/block
// http://localhost:6060/debug/pprof/mutex
//
// and also 2 more: the CPU profile and the CPU trace.
//
// http://localhost:6060/debug/pprof/profile?seconds=15
// http://localhost:6060/debug/pprof/trace?seconds=15
//
// run in the web
//
// go tool pprof -http=localhost:6061 profile.out
func runProfiler() {
	go func() {
		logger.Info("starting go profiler on 127.0.0.1:6060...")
		runtime.SetBlockProfileRate(1)
		log.Error(http.ListenAndServe("127.0.0.1:6060", nil))
	}()
}
