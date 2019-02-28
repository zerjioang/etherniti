// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

// +build !dev
// +build !pre
// +build prod

package config

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/zerjioang/etherniti/core/eth/fastime"
	"github.com/zerjioang/etherniti/core/logger"
)

const (
	EnvironmentName     = "production"
	HttpPort            = "80"
	HttpsPort           = "443"
	HttpAddress         = HttpListenInterface + HttpPort
	HttpsAddress        = HttpListenInterface + HttpsPort
	DebugServer         = false
	HideServerData      = true
	TokenSecret         = `IoHrlEV4vl9GViynFBHsgJ6qDxkWULgz98UQrO4m`
	EnableHttpsRedirect = false
	UseUniqueRequestId  = false
	EnableCors          = true
	EnableCache         = true
	EnableRateLimit     = true
	BlockTorConnections = true
	EnableLogging       = true
	LogLevel            = log.ERROR

	// production required listen mode
	HttpListenInterface = "0.0.0.0"
	ListeningAddress    = HttpListenInterface + ":" + HttpPort
	SwaggerAddress    = "http://dev-proxy.etherniti.org"

	//connection profile params
	TokenExpiration = 10 * fastime.Minute

	//rate limit units must be the same in both variables
	RateLimitUnits   = 1 * time.Hour
	RateLimitUnitsFt = 1 * fastime.Hour
	// ratelimit configuration
	RateLimit    = 100
	RateLimitStr = "100"
)

var (
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
	logger.Info("loading production ssl crypto material for https")
	certPath := os.Getenv("X_ETHERNITI_SSL_CERT_FILE")
	certPemBytes = loadCertBytes(certPath)

	keyPath := os.Getenv("X_ETHERNITI_SSL_KEY_FILE")
	keyPemBytes = loadCertBytes(keyPath)
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
