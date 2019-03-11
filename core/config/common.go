// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package config

import (
	"os"
	"strings"

	"github.com/zerjioang/etherniti/shared/def/listener"
)

var (
	//cert content as bytes readed from filesystem
	certPemBytes []byte
	//key content as bytes readed from filesystem
	keyPemBytes         []byte
	gopath              = os.Getenv("GOPATH")
	ResourcesDir        = gopath + "/src/github.com/zerjioang/etherniti/resources"
	ResourcesDirRoot    = ResourcesDir + "/root"
	ResourcesDirPHP     = ResourcesDir + "/root/phpinfo.php"
	ResourcesDirSwagger = ResourcesDir + "/swagger"
)

//read environment variables
func Read(env map[string]interface{}) {
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
	return certPemBytes
}

// get SSL certificate key.pem from proper source:
// hardcoded value or from local storage file
func GetKeyPem() []byte {
	return keyPemBytes
}

func IsHttpMode() bool {
	return listeningMode == "http"
}

func IsSocketMode() bool {
	return listeningMode == "socket"
}

func IsProfilingEnabled() bool {
	return true
}

func ServiceListeningMode() listener.ServiceType {
	switch listeningMode {
	case "http":
		return listener.HttpMode
	case "socket":
		return listener.UnixMode
	default:
		return listener.UnknownMode
	}
}
