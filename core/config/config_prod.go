// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

// +build !dev
// +build !pre
// +build prod

package config

import (
	"io/ioutil"
	"os"

	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/thirdparty/gommon/log"
)

var (
	//swagger.json injected params
	SwaggerApiDomain = "proxy.etherniti.org"
)

func init() {
	logger.Debug("loading productin config module data")
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

// setup server config
func Setup() error {
	logger.Debug("loading additional production setup config")
	return nil
}
