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

	"github.com/zerjioang/etherniti/core/modules/fastime"
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

//set default environment variables value for current context
func (c *EnvConfig) SetDefaults() {
	env := c.data
	env["X_ETHERNITI_LOG_LEVEL"] = "warn"
	env["X_ETHERNITI_ENVIRONMENT_NAME"] = "production"
	env["X_ETHERNITI_HTTP_PORT"] = "80"
	env["X_ETHERNITI_HTTPS_PORT"] = "443"
	env["X_ETHERNITI_DEBUG_SERVER"] = false
	env["X_ETHERNITI_HIDE_SERVER_DATA_IN_CONSOLE"] = true
	env["X_ETHERNITI_TOKEN_SECRET"] = `IoHrlEV4vl9GViynFBHsgJ6qDxkWULgz98UQrO4m`
	env["X_ETHERNITI_ENABLE_HTTPS_REDIRECT"] = true
	env["X_ETHERNITI_USE_UNIQUE_REQUEST_ID"] = false
	env["X_ETHERNITI_ENABLE_CORS"] = true
	env["X_ETHERNITI_ENABLE_SECURITY"] = true
	env["X_ETHERNITI_ENABLE_ANALYTICS"] = true
	env["X_ETHERNITI_ENABLE_METRICS"] = true
	env["X_ETHERNITI_ENABLE_CACHE"] = true
	env["X_ETHERNITI_ENABLE_RATELIMIT"] = true
	env["X_ETHERNITI_ENABLE_PROFILER"] = false
	env["X_ETHERNITI_BLOCK_TOR_CONNECTIONS"] = false
	env["X_ETHERNITI_ENABLE_LOGGING"] = true

	//for 'local development' deployment
	env["X_ETHERNITI_LISTENING_MODE"] = "https"
	env["X_ETHERNITI_LISTENING_INTERFACE"] = "0.0.0.0"
	env["X_ETHERNITI_LISTENING_ADDRESS"] = env["X_ETHERNITI_LISTENING_INTERFACE"].(string) + ":" + env["X_ETHERNITI_HTTP_PORT"].(string)
	env["X_ETHERNITI_SWAGGER_ADDRESS"] = "proxy.etherniti.org"

	//connection profile params
	env["X_ETHERNITI_TOKEN_EXPIRATION"] = 10 * fastime.Minute

	//rate limit units must be the same in both variables
	env["X_ETHERNITI_RATE_LIMIT_UNITS"] = 10 * time.Second
	env["X_ETHERNITI_RATE_LIMIT_UNITS_FT"] = 10 * time.Second

	// ratelimit configuration
	env["X_ETHERNITI_RATE_LIMIT"] = 10
	env["X_ETHERNITI_RATE_LIMIT_STR"] = "10"
}

// setup server config
func Setup() {
	logger.Debug("loading additional production setup config")
}
