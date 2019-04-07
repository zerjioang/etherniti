// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package config

import (
	"os"
	"strings"

	"github.com/zerjioang/etherniti/core/logger"

	"github.com/zerjioang/etherniti/shared/def/listener"
)

var (
	//cert content as bytes readed from filesystem
	certPemBytes []byte
	//key content as bytes readed from filesystem
	keyPemBytes  []byte
	gopath       = os.Getenv("GOPATH")
	ResourcesDir = gopath + "/src/github.com/zerjioang/etherniti/resources"
	// define internal folders
	ResourcesDirInternal         = ResourcesDir + "/internal"
	ResourcesDirInternalSecurity = ResourcesDirInternal + "/security"
	ResourcesDirRoot             = ResourcesDir + "/root"
	ResourcesDirSwagger          = ResourcesDir + "/swagger"
	// define internal files
	ResourcesDirPHP       = ResourcesDirRoot + "/phpinfo.php"
	BlacklistedDomainFile = ResourcesDirInternalSecurity + "/domains.json"
	PhishingDomainFile    = ResourcesDirInternalSecurity + "/phishing.json"
)

//read environment variables
func Read(env map[string]interface{}) {
	logger.Debug("reading environment variables")
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if len(pair) == 2 {
			env[pair[0]] = pair[1]
		}
	}
}

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
	return listeningMode == "http"
}

func IsSocketMode() bool {
	logger.Debug("checking if socket mode is enabled")
	return listeningMode == "socket"
}

func IsProfilingEnabled() bool {
	logger.Debug("checking if profiling mode is enabled")
	return false
}

func ServiceListeningMode() listener.ServiceType {
	logger.Debug("reading service listening mode")
	switch listeningMode {
	case "http":
		return listener.HttpMode
	case "socket":
		return listener.UnixMode
	default:
		return listener.UnknownMode
	}
}
