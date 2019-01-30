// Copyright gaethway
// SPDX-License-Identifier: Apache License 2.0

// +build prod

package config

const (
	HttpAddress  = "0.0.0.0:80"
	HttpsAddress = "0.0.0.0:443"
	HttpsPort    = ""
)

const (
	CertPem = ``
	KeyPem  = ``
)

//simply converts http requests into https
func GetRedirectUrl(host string, path string) string {
	return "https://" + host + path
}
