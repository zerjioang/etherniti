package config

import (
	"strconv"
	"strings"

	"github.com/zerjioang/etherniti/core/config/env"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/modules/fastime"
	"github.com/zerjioang/etherniti/core/modules/hashset"
	"github.com/zerjioang/etherniti/core/util/str"
	"github.com/zerjioang/etherniti/shared/def/listener"
	"github.com/zerjioang/etherniti/thirdparty/gommon/log"
)

type EthernitiAdminOptions struct {
	Key           string
	Secret        string
	LoadedFromEnv bool
}
type EthernitiOptions struct {
	//environment variables
	envData *env.EnvConfig

	//log level string
	LogLevelStr string
	LogLevel    log.Lvl

	// proxy manager configuration
	Admin EthernitiAdminOptions

	// feature configurations/activation
	LoggingEnabled     bool
	CORSEnabled        bool
	SecureModeEnabled  bool
	CompressionEnabled bool
	RateLimitEnabled   bool
	ServerCacheEnabled bool
	AnalyticsEnabled   bool
	MetricsEnabled     bool
	UniqueIdsEnabled   bool

	//swagger listening address
	SwaggerAddress string
	// http service listening address
	ListeningAddress string
	ListeningPort    int
	HttpInterface    string
	ListeningModeStr string
	ListeningMode    listener.ServiceType
	// automatic browser open configuration
	OpenBrowserOnSuccess bool

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
		LogLevelStr:             "warn",
		LogLevel:                log.WARN,
		LoggingEnabled:          true,
		CORSEnabled:             true,
		SecureModeEnabled:       true,
		CompressionEnabled:      true,
		RateLimitEnabled:        true,
		ServerCacheEnabled:      true,
		AnalyticsEnabled:        true,
		MetricsEnabled:          true,
		UniqueIdsEnabled:        true,
		SwaggerAddress:          "127.0.0.1",
		ListeningAddress:        "127.0.0.1",
		ListeningPort:           8080,
		HttpInterface:           "127.0.0.1",
		ListeningModeStr:        "http",
		ListeningMode:           listener.HttpMode,
		OpenBrowserOnSuccess:    true,
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
func (eo *EthernitiOptions) conditionalOverwriteBool(readed string, fallback bool) bool {
	if readed == "true" {
		return true
	} else if readed == "false" {
		return false
	} else {
		return fallback
	}
}

func (eo *EthernitiOptions) preload() {
	logger.Debug("preloading proxy configuration")
	eo.LogLevelStr = eo.conditionalOverwrite(eo.envData.String(XEthernitiLogLevel), eo.LogLevelStr)
	//resolve current logger level from string
	eo.LogLevel = eo.logLevelResolver()
	logger.Debug("updating log level to specified level: ", eo.LogLevel)
	logger.Level(eo.LogLevel)

	// load env variables to enable/disable modules
	logger.Debug("reading log status from env")
	eo.LoggingEnabled = eo.conditionalOverwriteBool(eo.envData.Lower(XEthernitiEnableLogging), eo.LoggingEnabled)
	logger.Debug("reading cors mode from env")
	eo.CORSEnabled = eo.conditionalOverwriteBool(eo.envData.Lower(XEthernitiEnableCors), eo.CORSEnabled)
	logger.Debug("reading secure listening mode from env")
	eo.SecureModeEnabled = eo.conditionalOverwriteBool(eo.envData.Lower(XEthernitiEnableSecurity), eo.SecureModeEnabled)
	logger.Debug("reading compression listening mode from env")
	eo.CompressionEnabled = eo.conditionalOverwriteBool(eo.envData.Lower(XEthernitiEnableGzip), eo.CompressionEnabled)
	logger.Debug("reading rate limit listening mode from env")
	eo.RateLimitEnabled = eo.conditionalOverwriteBool(eo.envData.Lower(XEthernitiEnableRateLimit), eo.RateLimitEnabled)
	logger.Debug("reading server cache listening mode from env")
	eo.ServerCacheEnabled = eo.conditionalOverwriteBool(eo.envData.Lower(XEthernitiEnableServerCache), eo.ServerCacheEnabled)
	logger.Debug("reading analytics listening mode from env")
	eo.AnalyticsEnabled = eo.conditionalOverwriteBool(eo.envData.Lower(XEthernitiEnableAnalytics), eo.AnalyticsEnabled)
	logger.Debug("reading metrics listening mode env")
	eo.MetricsEnabled = eo.conditionalOverwriteBool(eo.envData.Lower(XEthernitiEnableMetrics), eo.MetricsEnabled)
	logger.Debug("reading unique request id from env")
	eo.UniqueIdsEnabled = eo.conditionalOverwriteBool(eo.envData.Lower(XEthernitiUseUniqueRequestId), eo.UniqueIdsEnabled)

	//load swagger ui address
	logger.Debug("reading swagger address from env")
	eo.SwaggerAddress = eo.conditionalOverwrite(eo.envData.String(XEthernitiSwaggerAddress), eo.SwaggerAddress)

	//service listening options
	logger.Debug("reading requested listening ip address from env")
	eo.ListeningAddress = eo.conditionalOverwrite(eo.envData.String(XEthernitiListeningAddress), eo.ListeningAddress)
	logger.Debug("reading requested listening port from env")
	eo.ListeningPort = eo.envData.Int(XEthernitiListeningPort, 8080)
	logger.Debug("reading requested listening interface address from env")
	eo.HttpInterface = eo.conditionalOverwrite(eo.envData.String(XEthernitiListeningInterface), eo.HttpInterface)

	//service listening ListeningMode
	eo.ListeningModeStr = eo.conditionalOverwrite(eo.envData.String(XEthernitiListeningMode), eo.ListeningModeStr)
	eo.ListeningMode = eo.ServiceListeningModeResolver()

	// load browser automatic opening mode
	eo.OpenBrowserOnSuccess = eo.conditionalOverwriteBool(eo.envData.Lower(XEthernitiAutoOpenBrowser), eo.OpenBrowserOnSuccess)

	// load CORS data
	eo.AllowedCorsOriginList = hashset.NewHashSetWORM()
	eo.AllowedCorsOriginList.LoadFromRaw(CorsFile, "\n")

	// load hostnames data
	eo.AllowedHostnames = hashset.NewHashSetWORM()
	eo.AllowedHostnames.LoadFromRaw(HostsFile, "\n")

	eo.BlockTorConnections = eo.conditionalOverwriteBool(eo.envData.Lower(XEthernitiBlockTorConnections), eo.BlockTorConnections)
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
	logger.Debug("reading debug listening mode from env")
	v := eo.envData.Lower(XEthernitiDebugServer)
	return v == "true"
}

func (eo *EthernitiOptions) HideServerData() bool {
	logger.Debug("reading debug listening mode from env")
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

func (eo *EthernitiOptions) GetListeningAddressWithPort() string {
	return eo.ListeningAddress + ":" + eo.GetListeningPortStr()
}
func (eo *EthernitiOptions) GetURI() string {
	return "http://" + eo.GetListeningAddressWithPort()
}

func (eo *EthernitiOptions) GetListeningPortStr() string {
	return strconv.Itoa(eo.ListeningPort)
}

func (eo *EthernitiOptions) GetHttpInterface() string {
	return eo.HttpInterface
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
	logger.Debug("checking if http listening mode is enabled")
	return eo.ListeningMode == listener.HttpMode
}

func (eo *EthernitiOptions) IsHttpsMode() bool {
	logger.Debug("checking if https listening mode is enabled")
	return eo.ListeningMode == listener.HttpsMode
}

func (eo *EthernitiOptions) IsUnixSocketMode() bool {
	logger.Debug("checking if socket listening mode is enabled")
	return eo.ListeningMode == listener.UnixMode
}

func (eo *EthernitiOptions) IsWebSocketMode() bool {
	logger.Debug("checking if socket listening mode is enabled")
	return eo.ListeningMode == listener.WebsocketMode
}

func (eo *EthernitiOptions) IsSecureWebSocketMode() bool {
	logger.Debug("checking if secure socket listening mode is enabled")
	return eo.ListeningMode == listener.SecureWebsocketMode
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

func (eo *EthernitiOptions) ServiceListeningModeResolver() listener.ServiceType {
	eo.ListeningModeStr = str.ToLowerAscii(eo.ListeningModeStr)
	switch eo.ListeningModeStr {
	case "http":
		return listener.HttpMode
	case "https":
		return listener.HttpsMode
	case "socket", "unix", "ipc":
		return listener.UnixMode
	case "ws", "websocket":
		return listener.WebsocketMode
	case "wss", "securewebsocket":
		return listener.SecureWebsocketMode
	default:
		return listener.UnknownMode
	}
}

func (eo *EthernitiOptions) logLevelResolver() log.Lvl {
	eo.LogLevelStr = strings.ToLower(eo.LogLevelStr)
	switch eo.LogLevelStr {
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

func GetDefaultOpts() *EthernitiOptions {
	return defaultOptions
}
