// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

// +build prod

package config

import (
	"io/ioutil"
	"os"

	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/go-hpc/thirdparty/gommon/log"
)

var (
	//swagger.json injected params
	SwaggerApiDomain = "proxy.etherniti.org"
)

func init() {
	logger.Debug("loading production config module data")
	logger.Info("loading production ssl crypto material for https")
	certPath := os.Getenv("X_ETHERNITI_SSL_CERT_FILE")
	certPemBytes = loadCertBytes(certPath)

	keyPath := os.Getenv("X_ETHERNITI_SSL_KEY_FILE")
	keyPemBytes = loadCertBytes(keyPath)
}

func loadCertBytes(path string) []byte {
	logger.Debug("loading cert file data from fs")
	certData, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("failed to load production HTTPS SSL certificate data")
	}
	if len(certData) == 0 {
		log.Fatal("failed to load production HTTPS SSL certificate data. Empty content found on given file")
	}
	return certData
}

//check if profiling is enabled or not
// preproduction and production profiling
// is always disabled
func IsProfilingEnabled() bool {
	return false
}

func IsDevelopment() bool {
	logger.Debug("checking if current server environment is development")
	return false
}

// read swagger url domain: ip or FQDN
// no port allowed in pre and prod. default port is used
// 80 and 443
func GetSwaggerAddressWithPort(opts EthernitiOptions) string {
	logger.Debug("reading swagger address port from env")
	return opts.SwaggerAddress
}

func Env() string {
	logger.Debug("reading server environment name")
	return "production"
}

// setup server config
func Setup(opts *EthernitiOptions) error {
	logger.Debug("loading additional production setup config")
	return nil
}
