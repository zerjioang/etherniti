// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package config

import (
	"os"
	"strings"

	"github.com/zerjioang/etherniti/core/eth/fastime"
	"github.com/zerjioang/etherniti/core/modules/hashset"
	"github.com/zerjioang/etherniti/thirdparty/gommon/log"

	"github.com/zerjioang/etherniti/core/logger"

	"github.com/zerjioang/etherniti/shared/def/listener"
)

var (
	//cert content as bytes
	certPemBytes []byte
	//key content as bytes
	keyPemBytes  []byte
	gopath       = os.Getenv("GOPATH")
	ResourcesDir = gopath + "/src/github.com/zerjioang/etherniti/testdata"
	// define internal folders
	ResourcesDirInternal         = ResourcesDir + "/internal"
	ResourcesDirInternalSecurity = ResourcesDirInternal + "/security"
	ResourcesDirInternalBots     = ResourcesDirInternal + "/bots"
	ResourcesDirInternalTor      = ResourcesDirInternal + "/tor"
	ResourcesDirInternalBadIps   = ResourcesDirInternal + "/badips"
	ResourcesDirInternalCors     = ResourcesDirInternal + "/cors"
	ResourcesDirInternalHosts    = ResourcesDirInternal + "/hosts"
	ResourcesDirInternalEmail    = ResourcesDirInternal + "/templates/mail"
	ResourcesDirLanding          = ResourcesDir + "/landing"
	ResourcesIndexHtml           = ResourcesDirLanding + "/index.html"
	ResourcesDirRoot             = ResourcesDir + "/root"
	ResourcesDirSwagger          = ResourcesDir + "/swagger"
	// define internal files
	ResourcesDirPHP       = ResourcesDirRoot + "/phpinfo.php"
	BlacklistedDomainFile = ResourcesDirInternalSecurity + "/domains.json"
	PhishingDomainFile    = ResourcesDirInternalSecurity + "/phishing.json"
	AntiBotsFile          = ResourcesDirInternalBots + "/bots.json"
	TorExitFile           = ResourcesDirInternalTor + "/tor_exit"
	TorAllFile            = ResourcesDirInternalTor + "/tor_all"
	BadIpsFile            = ResourcesDirInternalBadIps + "/list_any_5"
	CorsFile              = ResourcesDirInternalCors + "/cors"
	HostsFile             = ResourcesDirInternalHosts + "/hosts"
)

var (
	// allowed cors domains
	AllowedCorsOriginList hashset.HashSetMutex
	AllowedHostnames      hashset.HashSetMutex
	// user configured values
	BlockTorConnections bool
)

func init() {
	AllowedCorsOriginList = hashset.NewHashSet()
	AllowedCorsOriginList.LoadFromRaw(CorsFile, "\n")
	AllowedHostnames = hashset.NewHashSet()
	AllowedHostnames.LoadFromRaw(HostsFile, "\n")
	BlockTorConnections = resolveBlockTorConnections()
}

func resolveBlockTorConnections() bool {
	v, found := ReadEnvironment("X_ETHERNITI_BLOCK_TOR_CONNECTIONS")
	return found && v == true
}

func LogLevel() log.Lvl {
	value := ReadEnvironmentString("X_ETHERNITI_LOG_LEVEL")
	value = strings.ToLower(value)
	switch value {
	case "debug":
		return log.DEBUG
	case "info":
		return log.INFO
	case "warn":
		return log.WARN
	case "error":
		return log.ERROR
	case "off":
		return log.OFF
	default:
		return log.DEBUG
	}
}

func EnableLogging() bool {
	logger.Debug("reading logging level from env")
	v, found := ReadEnvironment("X_ETHERNITI_ENABLE_LOGGING")
	return found && v == true
}
func EnableSecureMode() bool {
	logger.Debug("reading secure mode from env")
	v, found := ReadEnvironment("X_ETHERNITI_ENABLE_SECURITY")
	return found && v == true
}
func EnableCors() bool {
	logger.Debug("reading cors mode from env")
	v, found := ReadEnvironment("X_ETHERNITI_ENABLE_CORS")
	return found && v == true
}
func EnableRateLimit() bool {
	logger.Debug("reading rate limit mode from env")
	v, found := ReadEnvironment("X_ETHERNITI_ENABLE_RATE_LIMIT")
	return found && v == true
}
func EnableAnalytics() bool {
	logger.Debug("reading analytics mode from env")
	v, found := ReadEnvironment("X_ETHERNITI_ENABLE_ANALYTICS")
	return found && v == true
}
func EnableMetrics() bool {
	logger.Debug("reading metrics mode env")
	v, found := ReadEnvironment("X_ETHERNITI_ENABLE_METRICS")
	return found && v == true
}
func UseUniqueRequestId() bool {
	logger.Debug("reading unique request id from env")
	v, found := ReadEnvironment("X_ETHERNITI_USE_UNIQUE_REQUEST_ID")
	return found && v == true
}

func RateLimit() uint32 {
	logger.Debug("reading ratelimit uint32 from env")
	v, found := ReadEnvironment("X_ETHERNITI_RATE_LIMIT")
	if found && v != nil {
		return v.(uint32)
	}
	return 100
}

func RateLimitUnitsFt() fastime.Duration {
	logger.Debug("reading rate limit units from env")
	v, found := ReadEnvironment("X_ETHERNITI_RATE_LIMIT_UNITS_FT")
	if found && v != nil {
		return v.(fastime.Duration)
	}
	return 100 * fastime.Hour
}

func RateLimitUnitsStr() string {
	logger.Debug("reading rate limit string from env")
	return ReadEnvironmentString("X_ETHERNITI_RATE_LIMIT_UNITS")
}

func TokenSecret() string {
	logger.Debug("reading token secret from env")
	return ReadEnvironmentString("X_ETHERNITI_TOKEN_SECRET")
}

func DebugServer() bool {
	logger.Debug("reading debug mode from env")
	v, found := ReadEnvironment("X_ETHERNITI_DEBUG_SERVER")
	return found && v == true
}

func TokenExpiration() fastime.Duration {
	logger.Debug("reading token expiration from env")
	v, found := ReadEnvironment("X_ETHERNITI_TOKEN_EXPIRATION")
	if found && v != nil {
		return v.(fastime.Duration)
	}
	return 100 * fastime.Hour
}

func GetSwaggerAddress() string {
	logger.Debug("reading swagger address from env")
	return ReadEnvironmentString("X_ETHERNITI_SWAGGER_ADDRESS")
}

func GetEnvironmentName() string {
	logger.Debug("reading etherniti environment name env")
	return ReadEnvironmentString("X_ETHERNITI_ENVIRONMENT_NAME")
}

func GetListeningAddress() string {
	logger.Debug("reading requested listening ip address from env")
	return ReadEnvironmentString("X_ETHERNITI_LISTENING_ADDRESS")
}

func GetHttpInterface() string {
	logger.Debug("reading requested listening interface address from env")
	return ReadEnvironmentString("X_ETHERNITI_HTTP_LISTEN_INTERFACE")
}

//simply converts http requests into https
func GetRedirectUrl(host string, path string) string {
	return "https://" + GetListeningAddress() + path
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
	return ReadEnvironmentString("X_ETHERNITI_LISTENING_MODE") == "http"
}

func IsSocketMode() bool {
	logger.Debug("checking if socket mode is enabled")
	return ReadEnvironmentString("X_ETHERNITI_LISTENING_MODE") == "socket"
}

func IsProfilingEnabled() bool {
	logger.Debug("checking if profiling mode is enabled")
	v, found := ReadEnvironment("X_ETHERNITI_ENABLE_PROFILER")
	return found && v == true
}

func GetEmailUsername() string {
	logger.Debug("reading email username from env")
	return ReadEnvironmentString("X_ETHERNITI_EMAIL_USERNAME")
}

func GetEmailPassword() string {
	logger.Debug("reading email password from env")
	return ReadEnvironmentString("X_ETHERNITI_GMAIL_ACCESS_TOKEN")
}
func GetEmailServer() string {
	logger.Debug("reading email server name and port from env")
	return ReadEnvironmentString("X_ETHERNITI_EMAIL_SERVER")
}
func GetEmailServerOnly() string {
	logger.Debug("reading email server name from env")
	return ReadEnvironmentString("X_ETHERNITI_EMAIL_SERVER_ONLY")
}

// sendgrid service configuration
func SendGridApiKey() string {
	logger.Debug("reading sendgrid api key from env")
	return ReadEnvironmentString("SENDGRID_API_KEY")
}

func ServiceListeningMode() listener.ServiceType {
	logger.Debug("reading service listening mode")
	switch ReadEnvironmentString("X_ETHERNITI_LISTENING_MODE") {
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
	logger.Debug("checking if current server environment is development")
	return GetEnvironmentName() == "development"
}
