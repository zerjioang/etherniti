// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/zerjioang/etherniti/core/modules/fastime"
	"github.com/zerjioang/etherniti/core/modules/hashset"
	"github.com/zerjioang/etherniti/thirdparty/gommon/log"

	"github.com/zerjioang/etherniti/core/logger"

	"github.com/zerjioang/etherniti/shared/def/listener"
)

var (
	//cert content as bytes
	certPemBytes []byte
	//key content as bytes
	keyPemBytes []byte
	//system gopath
	gopath = os.Getenv("GOPATH")
	// current executable process id
	Pid = os.Getpid()
	//resources base dir
	ResourcesDir = gopath + "/src/github.com/zerjioang/etherniti/testdata"
	// define internal folders
	ResourcesDirInternal         = ResourcesDir + "/internal"
	ResourcesDirInternalSecurity = ResourcesDirInternal + "/security"
	ResourcesDirInternalBots     = ResourcesDirInternal + "/bots"
	ResourcesDirInternalTor      = ResourcesDirInternal + "/tor"
	ResourcesDirInternalTokens   = ResourcesDirInternal + "/tokens"
	ResourcesDirInternalBadIps   = ResourcesDirInternal + "/badips"
	ResourcesDirInternalCors     = ResourcesDirInternal + "/cors"
	ResourcesDirInternalHosts    = ResourcesDirInternal + "/hosts"
	ResourcesDirInternalEmail    = ResourcesDirInternal + "/templates/mail"
	ResourcesDirRoot             = ResourcesDir + "/root"
	ResourcesIndexHtml           = ResourcesDirRoot + "/index.html"
	ResourcesDirSwagger          = ResourcesDir + "/swagger"
	// define internal files
	ResourcesDirPHP       = ResourcesDirRoot + "/phpinfo.php"
	BlacklistedDomainFile = ResourcesDirInternalSecurity + "/domains.json"
	TokenListFile         = ResourcesDirInternalTokens + "/list.json"
	PhishingDomainFile    = ResourcesDirInternalSecurity + "/phishing.json"
	AntiBotsFile          = ResourcesDirInternalBots + "/bots.json"
	TorExitFile           = ResourcesDirInternalTor + "/tor_exit"
	TorAllFile            = ResourcesDirInternalTor + "/tor_all"
	BadIpsFile            = ResourcesDirInternalBadIps + "/list_any_5"
	CorsFile              = ResourcesDirInternalCors + "/cors"
	HostsFile             = ResourcesDirInternalHosts + "/hosts"
)

var (
	// single point of object access in an object-oriented application
	proxyEnv *EnvConfig
	// allowed cors domains
	AllowedCorsOriginList hashset.HashSetMutex
	AllowedHostnames      hashset.HashSetMutex
	// user configured values
	BlockTorConnections bool
	// worker configuration
	MaxWorker int
	MaxQueue  int
)

func init() {
	// load environment variables once
	logger.Debug("reading environment configuration")
	proxyEnv = newEnvironment()
	// read current os environment variables
	proxyEnv.readEnvironmentData()

	// load CORS data
	AllowedCorsOriginList = hashset.NewHashSet()
	AllowedCorsOriginList.LoadFromRaw(CorsFile, "\n")

	// load hostnames data
	AllowedHostnames = hashset.NewHashSet()
	AllowedHostnames.LoadFromRaw(HostsFile, "\n")

	// preload env config from readed data
	preload()
}

func preload() {
	logger.Debug("preloading proxy configuration")
	BlockTorConnections = resolveBlockTorConnections()
	//worker configuration
	e := GetEnvironment()
	MaxWorker = e.Int(XEthernitiMaxWorkers, 4)
	MaxQueue = e.Int(XEthernitiMaxQueue, 200)
}

func resolveBlockTorConnections() bool {
	v := GetEnvironment().Lower(XEthernitiBlockTorConnections)
	return v == "true"
}

func LogLevelStr() string {
	return GetEnvironment().String(XEthernitiLogLevel)
}
func LogLevel() log.Lvl {
	value := GetEnvironment().String(XEthernitiLogLevel)
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

func EnableLoggingStr() string {
	return GetEnvironment().String(XEthernitiEnableLogging)
}
func EnableLogging() bool {
	logger.Debug("reading logging level from env")
	v := GetEnvironment().Lower(XEthernitiEnableLogging)
	return v == "true"
}
func EnableSecureMode() bool {
	logger.Debug("reading secure mode from env")
	v := GetEnvironment().Lower(XEthernitiEnableSecurity)
	return v == "true"
}
func EnableCors() bool {
	logger.Debug("reading cors mode from env")
	v := GetEnvironment().Lower(XEthernitiEnableCors)
	return v == "true"
}
func EnableRateLimit() bool {
	logger.Debug("reading rate limit mode from env")
	v := GetEnvironment().Lower(XEthernitiEnableRateLimit)
	return v == "true"
}
func EnableCompression() bool {
	logger.Debug("reading compression mode from env")
	v := GetEnvironment().Lower(XEthernitiEnableGzip)
	return v == "true"
}
func EnableServerCache() bool {
	logger.Debug("reading server cache mode from env")
	v := GetEnvironment().Lower(XEthernitiEnableServerCache)
	return v == "true"
}
func EnableAnalytics() bool {
	logger.Debug("reading analytics mode from env")
	v := GetEnvironment().Lower(XEthernitiEnableAnalytics)
	return v == "true"
}
func EnableMetrics() bool {
	logger.Debug("reading metrics mode env")
	v := GetEnvironment().Lower(XEthernitiEnableMetrics)
	return v == "true"
}
func UseUniqueRequestId() bool {
	logger.Debug("reading unique request id from env")
	v := GetEnvironment().Lower(XEthernitiUseUniqueRequestId)
	return v == "true"
}

func RateLimit() uint32 {
	logger.Debug("reading ratelimit uint32 from env")
	v, found := GetEnvironment().Read(XEthernitiRateLimit)
	if found && v != nil {
		return v.(uint32)
	}
	return 60
}

func RateLimitStr() string {
	return strconv.Itoa(int(RateLimit()))
}

func RateLimitUnitsFt() fastime.Duration {
	logger.Debug("reading rate limit units from env")
	v, found := GetEnvironment().Read(XEthernitiRateLimitUnitsFt)
	if found && v != nil {
		return v.(fastime.Duration)
	}
	return 1 * fastime.Minute
}

func TokenSecret() string {
	logger.Debug("reading token secret from env")
	return GetEnvironment().String(XEthernitiTokenSecret)
}

func DebugServer() bool {
	logger.Debug("reading debug mode from env")
	v := GetEnvironment().Lower(XEthernitiDebugServer)
	return v == "true"
}

func HideServerData() bool {
	logger.Debug("reading debug mode from env")
	v := GetEnvironment().Lower(XEthernitiHideServerDataInConsole)
	return v == "true"
}

func TokenExpiration() fastime.Duration {
	logger.Debug("reading token expiration from env")
	v, found := GetEnvironment().Read(XEthernitiTokenExpiration)
	if found && v != nil {
		return v.(fastime.Duration)
	}
	return 100 * fastime.Hour
}

func GetSwaggerAddress() string {
	logger.Debug("reading swagger address from env")
	return GetEnvironment().String(XEthernitiSwaggerAddress)
}

func GetListeningAddress() string {
	logger.Debug("reading requested listening ip address from env")
	return GetEnvironment().String(XEthernitiListeningAddress)
}

func GetListeningAddressWithPort() string {
	return GetListeningAddress() + ":" + GetListeningPortStr()
}

func GetListeningPort() int {
	logger.Debug("reading requested listening port from env")
	return GetEnvironment().Int(XEthernitiListeningPort, 8080)
}

func GetListeningPortStr() string {
	logger.Debug("reading requested listening port from env")
	return GetEnvironment().String(XEthernitiListeningPort)
}

func GetHttpInterface() string {
	logger.Debug("reading requested listening interface address from env")
	return GetEnvironment().String(XEthernitiListeningInterface)
}

//simply converts http requests into https
func GetRedirectUrl(host string, path string) string {
	return "https://" + GetListeningAddressWithPort() + path
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

func InfuraToken() string {
	logger.Debug("reading infura token")
	return GetEnvironment().String(XEthernitiInfuraToken)
}

func IsHttpMode() bool {
	logger.Debug("checking if http mode is enabled")
	return GetEnvironment().String(XEthernitiListeningMode) == "http"
}

func IsHttpsMode() bool {
	logger.Debug("checking if https mode is enabled")
	return GetEnvironment().String(XEthernitiListeningMode) == "https"
}

func IsUnixSocketMode() bool {
	logger.Debug("checking if socket mode is enabled")
	return GetEnvironment().String(XEthernitiListeningMode) == "socket"
}

func IsWebSocketMode() bool {
	logger.Debug("checking if socket mode is enabled")
	return GetEnvironment().String(XEthernitiListeningMode) == "ws"
}

func GetEmailUsername() string {
	logger.Debug("reading email username from env")
	return GetEnvironment().String(XEthernitiEmailUsername)
}

func GetEmailPassword() string {
	logger.Debug("reading email password from env")
	return GetEnvironment().String(XEthernitiGmailAccessToken)
}
func GetEmailServer() string {
	logger.Debug("reading email server name and port from env")
	return GetEnvironment().String(XEthernitiEmailServer)
}
func GetEmailServerOnly() string {
	logger.Debug("reading email server name from env")
	return GetEnvironment().String(XEthernitiEmailServerOnly)
}

// sendgrid service configuration
func SendGridApiKey() string {
	logger.Debug("reading sendgrid api key from env")
	return GetEnvironment().String(XEthernitiSendgridApiKey)
}

func ServiceListeningModeStr() string {
	logger.Debug("reading service listening mode")
	return GetEnvironment().String(XEthernitiListeningMode)
}
func ServiceListeningMode() listener.ServiceType {
	v := ServiceListeningModeStr()
	switch v {
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
