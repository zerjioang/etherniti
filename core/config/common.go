// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package config

const (
	//profile key http header key
	HttpProfileHeaderkey = "X-Etherniti-Profile"
)

var (
	//cert content as bytes readed from filesystem
	certPemBytes []byte
	//key content as bytes readed from filesystem
	keyPemBytes []byte
)

//simply converts http requests into https
func GetRedirectUrl(host string, path string) string {
	return "https://" + ListeningAddress + path
}

// get SSL certificate cert.pem from proper source:
// hardcoded value or from local storage file
func GetCertPem() []byte {
	return certPemBytes
}

// get SSL certificate key.pem from proper source:
// hardcoded value or from local storage file
func GetKeyPem() []byte {
	return keyPemBytes
}

func IsHttpMode() bool {
	return ListeningMode == "http"
}

func IsSocketMode() bool {
	return ListeningMode == "socket"
}
