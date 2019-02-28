// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package config

const (
	//profile key http header key
	HttpProfileHeaderkey = "X-Etherniti-Profile"

	// integrity related digital signature p256 private key
	IntegrityPrivateKeyPem = `-----BEGIN PRIVATE KEY-----
MGgCAQEEHC6J2ZYpSrUrIba8+msP0+ZmUnGV8DGYsRk/O7egBwYFK4EEACGhPAM6
AARePCmlTNeJjy78d51zYcIEiTpeGa7PJiDiiV/8FeVNsJIJwG6EEjhd8M8n65Yn
pd8dt/7HEahbvw==
-----END PRIVATE KEY-----`

	// integrity related digital signature p256 public key
	IntegrityPublicKeyPem = `-----BEGIN PUBLIC KEY-----
ME4wEAYHKoZIzj0CAQYFK4EEACEDOgAEXjwppUzXiY8u/Hedc2HCBIk6XhmuzyYg
4olf/BXlTbCSCcBuhBI4XfDPJ+uWJ6XfHbf+xxGoW78=
-----END PUBLIC KEY-----`
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
