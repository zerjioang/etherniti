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
	ListeningMode       = "http" // http or socket
	HttpListenInterface = "0.0.0.0"
	ListeningAddress    = HttpListenInterface + ":" + HttpPort
	SwaggerAddress      = "dev-proxy.etherniti.org"

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

//set default environment variables value for current context
func SetDefaults(env map[string]interface{}) map[string]interface{} {
	env["X_ETHERNITI_ENVIRONMENT_NAME"] = "beta-stage"
	env["X_ETHERNITI_HTTP_PORT"] = "8080"
	env["X_ETHERNITI_HTTPS_PORT"] = "4430"
	env["X_ETHERNITI_DEBUG_SERVER"] = false
	env["X_ETHERNITI_HIDE_SERVER_DATA_IN_CONSOLE"] = true
	env["X_ETHERNITI_TOKEN_SECRET"] = "t0k3n-s3cr3t-h3r3"
	env["X_ETHERNITI_ENABLE_HTTPS_REDIRECT"] = false
	env["X_ETHERNITI_USE_UNIQUE_REQUEST_ID"] = false
	env["X_ETHERNITI_ENABLE_CORS"] = true
	env["X_ETHERNITI_ENABLE_CACHE"] = true
	env["X_ETHERNITI_ENABLE_RATELIMIT"] = false
	env["X_ETHERNITI_BLOCK_TOR_CONNECTIONS"] = false
	env["X_ETHERNITI_ENABLE_LOGGING"] = true
	env["X_ETHERNITI_LOG_LEVEL"] = log.DEBUG

	//for 'local development' deployment
	env["X_ETHERNITI_LISTENING_MODE"] = "http" // http or socket
	env["X_ETHERNITI_LISTENING_INTERFACE"] = "0.0.0.0"
	env["X_ETHERNITI_LISTENING_ADDRESS"] = HttpListenInterface + ":" + HttpPort
	env["X_ETHERNITI_SWAGGER_ADDRESS"] = "dev-proxy.etherniti.org"

	//connection profile params
	env["X_ETHERNITI_TOKEN_EXPIRATION"] = 10 * fastime.Minute

	//rate limit units must be the same in both variables
	env["X_ETHERNITI_RATE_LIMIT_UNITS"] = 10 * time.Second
	env["X_ETHERNITI_RATE_LIMIT_UNITS_FT"] = 10 * time.Second

	// ratelimit configuration
	env["X_ETHERNITI_RATE_LIMIT"] = 10
	env["X_ETHERNITI_RATE_LIMIT_STR"] = "10"

	return env
}
