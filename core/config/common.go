// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package config

import (
	"os"

	"github.com/zerjioang/etherniti/core/logger"

	"github.com/zerjioang/etherniti/shared/def/listener"
)

var (
	//cert content as bytes readed from filesystem
	certPemBytes []byte
	//key content as bytes readed from filesystem
	keyPemBytes  []byte
	gopath       = os.Getenv("GOPATH")
	ResourcesDir = gopath + "/src/github.com/zerjioang/etherniti/testdata"
	// define internal folders
	ResourcesDirInternal         = ResourcesDir + "/internal"
	ResourcesDirInternalSecurity = ResourcesDirInternal + "/security"
	ResourcesDirInternalBots     = ResourcesDirInternal + "/bots"
	ResourcesDirLanding          = ResourcesDir + "/landing"
	ResourcesIndexHtml         = ResourcesDirLanding + "/index.html"
	ResourcesDirRoot             = ResourcesDir + "/root"
	ResourcesDirSwagger          = ResourcesDir + "/swagger"
	// define internal files
	ResourcesDirPHP       = ResourcesDirRoot + "/phpinfo.php"
	BlacklistedDomainFile = ResourcesDirInternalSecurity + "/domains.json"
	PhishingDomainFile    = ResourcesDirInternalSecurity + "/phishing.json"
	AntiBotsFile          = ResourcesDirInternalBots + "/bots.json"
)

//simply converts http requests into https
func GetRedirectUrl(host string, path string) string {
	return "https://" + ListeningAddress + path
}

// get SSL certificate cert.pem from proper source:
// hardcoded value or from local storage file
func GetCertPem() []byte {
	logger.Debug("getting .pem cert data")
	return certPemBytes
}

// get SSL certificate key.pem from proper source:
// hardcoded value or from local storage file
func GetKeyPem() []byte {
	logger.Debug("getting .pem key data")
	return keyPemBytes
}

func IsHttpMode() bool {
	logger.Debug("checking if http mode is enabled")
	return ReadEnvironment("X_ETHERNITI_LISTENING_MODE") == "http"
}

func IsSocketMode() bool {
	logger.Debug("checking if socket mode is enabled")
	return ReadEnvironment("X_ETHERNITI_LISTENING_MODE") == "socket"
}

func IsProfilingEnabled() bool {
	logger.Debug("checking if profiling mode is enabled")
	return ReadEnvironment("X_ETHERNITI_ENABLE_PROFILER") == true
}

func ServiceListeningMode() listener.ServiceType {
	logger.Debug("reading service listening mode")
	switch ReadEnvironment("X_ETHERNITI_LISTENING_MODE") {
	case "http":
		return listener.HttpMode
	case "https":
		return listener.HttpsMode
	case "socket":
		return listener.UnixMode
	default:
		return listener.UnknownMode
	}
}

func IsDevelopment() bool {
	return EnvironmentName == "development"
}