// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

// +build prod

package config

import (
	"io/ioutil"
	"os"

	"github.com/labstack/gommon/log"
)

const (
	HttpPort            = ":80"
	HttpsPort           = ":443"
	HttpAddress         = "0.0.0.0:80"
	HttpsAddress        = "0.0.0.0:443"
	DebugServer         = false
	HideServerData      = true
	TokenSecret         = `IoHrlEV4vl9GViynFBHsgJ6qDxkWULgz98UQrO4m`
	EnableHttpsRedirect = false
	UseUniqueRequestId  = false
	EnableRateLimit     = true
	BlockTorConnections = true
	EnableLogging       = true
	LogLevel            = log.ERROR
)

var (
	//cert content as bytes readed from filesystem
	fsCertBytes []byte
	//key content as bytes readed from filesystem
	fsKeyBytes []byte
	// allowed cors domains
	// allowed cors domains
	AllowedCorsOriginList = []string{
		"*",
		"api.etherniti.org",
		"proxy.etherniti.org",
	}
	//allowed hostnames
	AllowedHostnames = []string{
		"api.etherniti.org",
		"proxy.etherniti.org",
	}
	//swagger.json injected params
	SwaggerApiDomain = "dev-proxy.etherniti.org"
)

func init() {
	log.Info("loading production ssl crypto material for https")
	certPath := os.Getenv("X_ETHERNITI_SSL_CERT_FILE")
	fsCertBytes = loadCertBytes(certPath)

	keyPath := os.Getenv("X_ETHERNITI_SSL_KEY_FILE")
	fsKeyBytes = loadCertBytes(keyPath)
}

func loadCertBytes(path string) []byte {
	certData, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("failed to load production HTTPS SSL certificate data")
	}
	if len(certData) == 0 {
		log.Fatal("failed to load production HTTPS SSL certificate data. Empty content found on given file")
	}
	return certData
}

//simply converts http requests into https
func GetRedirectUrl(host string, path string) string {
	return "https://" + host + path
}

// get SSL certificate cert.pem from proper source:
// hardcoded value or from local storage file
func GetCertPem() []byte {
	return fsCertBytes
}

// get SSL certificate key.pem from proper source:
// hardcoded value or from local storage file
func GetKeyPem() []byte {
	return fsKeyBytes
}
