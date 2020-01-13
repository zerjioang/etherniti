// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

// +build dev

package config

import (
	"net/http"
	"runtime"

	"github.com/zerjioang/etherniti/core/config/edition"

	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/go-hpc/util/str"

	_ "net/http/pprof" //adds 2,5Mb to final executable when imported

	"github.com/zerjioang/go-hpc/thirdparty/gommon/log"
)

// openssl req -newkey rsa:2048 \
//  -new -nodes -x509 \
//  -days 3650 \
//  -out cert.pem \
//  -keyout key.pem \
//  -subj "/C=ES/ST=Biscay/L=Bilbao/O=Etherniti/OU=Cyber Unit/CN=127.0.0.1

const (
	// openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
	// certificate signed for 127.0.0.1
	certPem = `-----BEGIN CERTIFICATE-----
MIICyjCCAlCgAwIBAgIUaDW7PwiRieONBAxj2l57fJ2nToIwCgYIKoZIzj0EAwIw
gZsxCzAJBgNVBAYTAkVTMQ8wDQYDVQQIDAZCaXNjYXkxDzANBgNVBAcMBkJpbGJh
bzESMBAGA1UECgwJRXRoZXJuaXRpMRYwFAYDVQQLDA1DeWJlcnNlY3VyaXR5MRIw
EAYDVQQDDAkxMjcuMC4wLjExKjAoBgkqhkiG9w0BCQEWG2N5YmVyc2VjdXJpdHlA
ZXRoZXJuaXRpLm9yZzAeFw0xOTA5MTUxOTUyMzJaFw0yOTA5MTIxOTUyMzJaMIGb
MQswCQYDVQQGEwJFUzEPMA0GA1UECAwGQmlzY2F5MQ8wDQYDVQQHDAZCaWxiYW8x
EjAQBgNVBAoMCUV0aGVybml0aTEWMBQGA1UECwwNQ3liZXJzZWN1cml0eTESMBAG
A1UEAwwJMTI3LjAuMC4xMSowKAYJKoZIhvcNAQkBFhtjeWJlcnNlY3VyaXR5QGV0
aGVybml0aS5vcmcwdjAQBgcqhkjOPQIBBgUrgQQAIgNiAATTFv0BHZZLNTUD0GXq
/3aueh7HW4GXbE186ERf57nE/Sf/F7ZLUs150Nlh/IZ8zAkCB8Fnaci5+C0Te7Ur
GgzIAuNwkEscUtYxW/YisLpjof7ZmgkOoxd8zAR7b2Vo3DajUzBRMB0GA1UdDgQW
BBQPVPtmu22NeQx0tOzoIzgmyIBrPzAfBgNVHSMEGDAWgBQPVPtmu22NeQx0tOzo
IzgmyIBrPzAPBgNVHRMBAf8EBTADAQH/MAoGCCqGSM49BAMCA2gAMGUCMDnihM83
iJEwae0ORhPDzRPllH4ajsx5zV3L5ogOQWDaWCWrreoOtm4509LhDoBhkgIxAIQG
oCjNcgx0rP+/Y5zqIdlmfIVWlj4oKZhd4gJeSV5qkqVS8qqkAy3TCFWyy3aiYQ==
-----END CERTIFICATE-----

`

	// openssl genrsa -out server.key 4096
	// openssl ecparam -genkey -name secp384r1 -out server.key
	keyPem = `-----BEGIN EC PARAMETERS-----
BgUrgQQAIg==
-----END EC PARAMETERS-----
-----BEGIN EC PRIVATE KEY-----
MIGkAgEBBDBPC3gEQ9iLy1erH3+oZnCqXpOqMwxBLnlSY9nXBype9Iw0gDgLljxV
w8g67Pd4+iagBwYFK4EEACKhZANiAATTFv0BHZZLNTUD0GXq/3aueh7HW4GXbE18
6ERf57nE/Sf/F7ZLUs150Nlh/IZ8zAkCB8Fnaci5+C0Te7UrGgzIAuNwkEscUtYx
W/YisLpjof7ZmgkOoxd8zAR7b2Vo3DY=
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
func GetSwaggerAddressWithPort(opts EthernitiOptions) string {
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
