// Copyright gaethway
// SPDX-License-Identifier: Apache License 2.0

// +build prod

package config

import "github.com/labstack/gommon/log"

const (
	HttpAddress  = "0.0.0.0:80"
	HttpsAddress = "0.0.0.0:443"
	HttpsPort    = ""
)

const (
	CertPem     = ``
	KeyPem      = ``
	TokenSecret = `
IoHrlEV4vl9GViynFBHs
gJ6qDxkWULgz98UQrO4m
MfArd0xUhFOzC5Wfm9Pb
AAw0NP4Mes0ABN5o2NZQ
p4a2st3ziuhdqyYAVPUu
W0wWNkmIvvqNW5ifp1no
YRagObZKweSOD6H9hujc
02KJeEhiESHQ3slPMKU9
QHjieaOOKreDERoKUFqh
sXM4zHcNESu2ijyfTZmX
`
	EnableHttpsRedirect = false
	UseUniqueRequestId = false
	EnableRateLimit = true
	BlockTorConnections = true
	EnableLogging = true
	LogLevel = log.ERROR
)

//simply converts http requests into https
func GetRedirectUrl(host string, path string) string {
	return "https://" + host + path
}
