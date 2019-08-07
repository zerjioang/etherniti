package config

import (
	"strconv"
	"strings"

	"github.com/zerjioang/etherniti/core/config/env"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/modules/fastime"
	"github.com/zerjioang/etherniti/core/modules/hashset"
	"github.com/zerjioang/etherniti/shared/def/listener"
	"github.com/zerjioang/etherniti/thirdparty/gommon/log"
)

type EthernitiOptions struct {
	//environment variables
	envData *env.EnvConfig

	//log level string
	logLevelStr string
	logLevel    log.Lvl

	//swagger listening address
	swaggerAddress string
	// http service listening address
	listeningAddress string
	listeningPort    int
	httpInterface    string
	mode             string

	// allowed cors domains
	AllowedCorsOriginList hashset.HashSetWORM
	AllowedHostnames      hashset.HashSetWORM
	// user configured values
	BlockTorConnections bool
	// worker configuration
	MaxWorker int
	MaxQueue  int
	//custom endpoints for eth mainnets
	RopstenCustomEndpoint string
	RinkebyCustomEndpoint string
	KovanCustomEndpoint   string
	MainnetCustomEndpoint string
	//users management configuration
	useFirebaseManagement   bool
	checkUsersEmailValidity bool
	MinPasswordLen          int
	webAuthNEnabled         bool
}

var (
	// default etherniti proxy options
	defaultOptions = &EthernitiOptions{
		logLevelStr:             "debug",
		logLevel:                log.DEBUG,
		swaggerAddress:          "127.0.0.1",
		listeningAddress:        "127.0.0.1",
		listeningPort:           8080,
		httpInterface:           "127.0.0.1",
		mode:                    "http",
		BlockTorConnections:     false,
		MaxWorker:               4,
		MaxQueue:                200,
		RopstenCustomEndpoint:   "",
		RinkebyCustomEndpoint:   "",
		KovanCustomEndpoint:     "",
		MainnetCustomEndpoint:   "",
		useFirebaseManagement:   false,
		checkUsersEmailValidity: false,
		MinPasswordLen:          6,
		webAuthNEnabled:         false,
	}
	//default token expiration time when users does not provide one
	// 10 minute
	defaultTokenExpirationTime = 10 * fastime.Minute
)

func init() {
	defaultOptions.Init()
}

func (eo *EthernitiOptions) Init() {
	// load environment variables once
	logger.Debug("reading environment configuration")
	eo.envData = env.New()
	// read current os environment variables
	eo.envData.Load()

	// preload env config from readed data
	eo.preload()
}
func (eo *EthernitiOptions) conditionalOverwrite(readed, fallback string) string {
	if readed != "" {
		return readed
	}
	return fallback
}

func (eo *EthernitiOptions) preload() {
	logger.Debug("preloading proxy configuration")
	eo.logLevelStr = eo.conditionalOverwrite(eo.envData.String(XEthernitiLogLevel), eo.logLevelStr)
	//resolve current logger level from string
	eo.logLevel = eo.LogLevel()

	//load swagger ui address
	logger.Debug("reading swagger address from env")
	eo.swaggerAddress = eo.conditionalOverwrite(eo.envData.String(XEthernitiSwaggerAddress), eo.swaggerAddress)

	//service listening options
	logger.Debug("reading requested listening ip address from env")
	eo.listeningAddress = eo.conditionalOverwrite(eo.envData.String(XEthernitiListeningAddress), eo.listeningAddress)
	logger.Debug("reading requested listening port from env")
	eo.listeningPort = eo.envData.Int(XEthernitiListeningPort, 8080)
	logger.Debug("reading requested listening interface address from env")
	eo.httpInterface = eo.conditionalOverwrite(eo.envData.String(XEthernitiListeningInterface), eo.httpInterface)

	//service listening mode
	eo.mode = eo.conditionalOverwrite(eo.envData.String(XEthernitiListeningMode), eo.mode)

	// load CORS data
	eo.AllowedCorsOriginList = hashset.NewHashSetWORM()
	eo.AllowedCorsOriginList.LoadFromRaw(CorsFile, "\n")

	// load hostnames data
	eo.AllowedHostnames = hashset.NewHashSetWORM()
	eo.AllowedHostnames.LoadFromRaw(HostsFile, "\n")

	eo.BlockTorConnections = eo.resolveBlockTorConnections()
	eo.MaxWorker = eo.envData.Int(XEthernitiMaxWorkers, 4)
	eo.MaxQueue = eo.envData.Int(XEthernitiMaxQueue, 200)
	// load if exists custom endpoints for public mainnets
	eo.RopstenCustomEndpoint = eo.envData.String(XEthernitiRopstenEndpoint)
	eo.RinkebyCustomEndpoint = eo.envData.String(XEthernitiRinkebyEndpoint)
	eo.KovanCustomEndpoint = eo.envData.String(XEthernitiKovanEndpoint)
	eo.MainnetCustomEndpoint = eo.envData.String(XEthernitiMainnetEndpoint)

	//load users management configuration data
	eo.useFirebaseManagement = eo.envData.Bool(XEthernitiUsersFirebase, false)     //disabled by default
	eo.checkUsersEmailValidity = eo.envData.Bool(XEthernitiUsersCheckEmail, false) //disabled by default
	eo.MinPasswordLen = eo.envData.Int(XEthernitiMinPasswordLength, 6)             //6 chars by default
	eo.webAuthNEnabled = eo.envData.Bool(XEthernitiEnableWebAuthN, false)          //disabled by default
}

func (eo *EthernitiOptions) resolveBlockTorConnections() bool {
	v := eo.envData.Lower(XEthernitiBlockTorConnections)
	return v == "true"
}

func (eo *EthernitiOptions) LogLevelStr() string {
	return eo.logLevelStr
}
func (eo *EthernitiOptions) LogLevel() log.Lvl {
	eo.logLevelStr = strings.ToLower(eo.logLevelStr)
	switch eo.logLevelStr {
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

func (eo *EthernitiOptions) EnableLoggingStr() string {
	return eo.envData.String(XEthernitiEnableLogging)
}
func (eo *EthernitiOptions) EnableLogging() bool {
	logger.Debug("reading logging level from env")
	v := eo.envData.Lower(XEthernitiEnableLogging)
	return v == "true"
}
func (eo *EthernitiOptions) EnableSecureMode() bool {
	logger.Debug("reading secure mode from env")
	v := eo.envData.Lower(XEthernitiEnableSecurity)
	return v == "true"
}
func (eo *EthernitiOptions) EnableCors() bool {
	logger.Debug("reading cors mode from env")
	v := eo.envData.Lower(XEthernitiEnableCors)
	return v == "true"
}
func (eo *EthernitiOptions) EnableRateLimit() bool {
	logger.Debug("reading rate limit mode from env")
	v := eo.envData.Lower(XEthernitiEnableRateLimit)
	return v == "true"
}
func (eo *EthernitiOptions) EnableCompression() bool {
	logger.Debug("reading compression mode from env")
	v := eo.envData.Lower(XEthernitiEnableGzip)
	return v == "true"
}
func (eo *EthernitiOptions) EnableServerCache() bool {
	logger.Debug("reading server cache mode from env")
	v := eo.envData.Lower(XEthernitiEnableServerCache)
	return v == "true"
}
func (eo *EthernitiOptions) EnableAnalytics() bool {
	logger.Debug("reading analytics mode from env")
	v := eo.envData.Lower(XEthernitiEnableAnalytics)
	return v == "true"
}
func (eo *EthernitiOptions) EnableMetrics() bool {
	logger.Debug("reading metrics mode env")
	v := eo.envData.Lower(XEthernitiEnableMetrics)
	return v == "true"
}
func (eo *EthernitiOptions) UseUniqueRequestId() bool {
	logger.Debug("reading unique request id from env")
	v := eo.envData.Lower(XEthernitiUseUniqueRequestId)
	return v == "true"
}

func (eo *EthernitiOptions) RateLimit() uint32 {
	logger.Debug("reading ratelimit uint32 from env")
	v, found := eo.envData.Read(XEthernitiRateLimit)
	if found && v != nil {
		return v.(uint32)
	}
	return 60
}

func (eo *EthernitiOptions) RateLimitStr() string {
	return strconv.Itoa(int(eo.RateLimit()))
}

func (eo *EthernitiOptions) RateLimitUnitsFt() fastime.Duration {
	logger.Debug("reading rate limit units from env")
	v, found := eo.envData.Read(XEthernitiRateLimitUnitsFt)
	if found && v != nil {
		return v.(fastime.Duration)
	}
	return 1 * fastime.Minute
}

func (eo *EthernitiOptions) TokenSecret() string {
	logger.Debug("reading token secret from env")
	return eo.envData.String(XEthernitiTokenSecret)
}

func (eo *EthernitiOptions) DebugServer() bool {
	logger.Debug("reading debug mode from env")
	v := eo.envData.Lower(XEthernitiDebugServer)
	return v == "true"
}

func (eo *EthernitiOptions) HideServerData() bool {
	logger.Debug("reading debug mode from env")
	v := eo.envData.Lower(XEthernitiHideServerDataInConsole)
	return v == "true"
}

func (eo *EthernitiOptions) TokenExpiration() fastime.Duration {
	logger.Debug("reading token expiration from env")
	v := eo.envData.Int(XEthernitiTokenExpiration, -1)
	if v == -1 {
		// error while reading value
		// return default
		return defaultTokenExpirationTime
	}
	return fastime.Duration(v) * fastime.Second
}

func (eo *EthernitiOptions) GetSwaggerAddress() string {
	return eo.swaggerAddress
}

func (eo *EthernitiOptions) GetListeningAddress() string {
	return eo.listeningAddress
}

func (eo *EthernitiOptions) GetListeningAddressWithPort() string {
	return eo.GetListeningAddress() + ":" + eo.GetListeningPortStr()
}

func (eo *EthernitiOptions) GetListeningPort() int {
	return eo.listeningPort
}

func (eo *EthernitiOptions) GetListeningPortStr() string {
	return strconv.Itoa(eo.listeningPort)
}

func (eo *EthernitiOptions) GetHttpInterface() string {
	return eo.httpInterface
}

//simply converts http requests into https
func (eo *EthernitiOptions) GetRedirectUrl(host string, path string) string {
	return "https://" + eo.GetListeningAddressWithPort() + path
}

// get SSL certificate cert.pem from proper source:
// hardcoded value or from local storage file
func (eo *EthernitiOptions) GetCertPem() []byte {
	logger.Debug("getting .pem cert data")
	return certPemBytes
}

// get SSL certificate key.pem from proper source:
// hardcoded value or from local storage file
func (eo *EthernitiOptions) GetKeyPem() []byte {
	logger.Debug("getting .pem key data")
	return keyPemBytes
}

func (eo *EthernitiOptions) InfuraToken() string {
	logger.Debug("reading infura token")
	return eo.envData.String(XEthernitiInfuraToken)
}

func (eo *EthernitiOptions) IsHttpMode() bool {
	logger.Debug("checking if http mode is enabled")
	return eo.mode == "http"
}

func (eo *EthernitiOptions) IsHttpsMode() bool {
	logger.Debug("checking if https mode is enabled")
	return eo.mode == "https"
}

func (eo *EthernitiOptions) IsUnixSocketMode() bool {
	logger.Debug("checking if socket mode is enabled")
	return eo.mode == "socket"
}

func (eo *EthernitiOptions) IsWebSocketMode() bool {
	logger.Debug("checking if socket mode is enabled")
	return eo.mode == "wsm"
}

func (eo *EthernitiOptions) GetEmailUsername() string {
	logger.Debug("reading email username from env")
	return eo.envData.String(XEthernitiEmailUsername)
}

func (eo *EthernitiOptions) GetEmailPassword() string {
	logger.Debug("reading email password from env")
	return eo.envData.String(XEthernitiGmailAccessToken)
}
func (eo *EthernitiOptions) GetEmailServer() string {
	logger.Debug("reading email server name and port from env")
	return eo.envData.String(XEthernitiEmailServer)
}
func (eo *EthernitiOptions) GetEmailServerOnly() string {
	logger.Debug("reading email server name from env")
	return eo.envData.String(XEthernitiEmailServerOnly)
}

// sendgrid service configuration
func (eo *EthernitiOptions) SendGridApiKey() string {
	logger.Debug("reading sendgrid api key from env")
	return eo.envData.String(XEthernitiSendgridApiKey)
}

func (eo *EthernitiOptions) ServiceListeningModeStr() string {
	logger.Debug("reading service listening mode")
	return eo.mode
}
func (eo *EthernitiOptions) ServiceListeningMode() listener.ServiceType {
	v := eo.ServiceListeningModeStr()
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

func GetDefaultOpts() *EthernitiOptions {
	return defaultOptions
}
