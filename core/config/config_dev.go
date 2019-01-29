// Copyright gaethway
// SPDX-License-Identifier: Apache License 2.0

// +build dev

package config

const (
	DevelopmentAddress = "localhost"
	HttpsPort          = ":4430"
	HttpPort           = ":8080"
	HttpAddress        = DevelopmentAddress + HttpPort
	HttpsAddress       = DevelopmentAddress + HttpsPort
)

//simply converts http requests into https
func GetRedirectUrl(host string, path string) string {
	return "https://" + HttpsAddress + path
}
