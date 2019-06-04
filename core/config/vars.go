package config

// this file contains all environment variable names allowed by
// current etherniti version

const (
	XEthernitiLogLevel                = "X_ETHERNITI_LOG_LEVEL"
	XEthernitiEnableLogging           = "X_ETHERNITI_ENABLE_LOGGING"
	XEthernitiEnableSecurity          = "X_ETHERNITI_ENABLE_SECURITY"
	XEthernitiEnableCors              = "X_ETHERNITI_ENABLE_CORS"
	XEthernitiEnableRateLimit         = "X_ETHERNITI_ENABLE_RATE_LIMIT"
	XEthernitiEnableAnalytics         = "X_ETHERNITI_ENABLE_ANALYTICS"
	XEthernitiEnableGzip              = "X_ETHERNITI_ENABLE_GZIP"
	XEthernitiEnableServerCache       = "X_ETHERNITI_ENABLE_SERVER_CACHE"
	XEthernitiEnableMetrics           = "X_ETHERNITI_ENABLE_METRICS"
	XEthernitiBlockTorConnections     = "X_ETHERNITI_BLOCK_TOR_CONNECTIONS"
	XEthernitiUseUniqueRequestId      = "X_ETHERNITI_USE_UNIQUE_REQUEST_ID"
	XEthernitiRateLimit               = "X_ETHERNITI_RATE_LIMIT"
	XEthernitiRateLimitUnitsFt        = "X_ETHERNITI_RATE_LIMIT_UNITS_FT"
	XEthernitiRateLimitUnits          = "X_ETHERNITI_RATE_LIMIT_UNITS"
	XEthernitiTokenSecret             = "X_ETHERNITI_TOKEN_SECRET"
	XEthernitiDebugServer             = "X_ETHERNITI_DEBUG_SERVER"
	XEthernitiHideServerDataInConsole = "X_ETHERNITI_HIDE_SERVER_DATA_IN_CONSOLE"
	XEthernitiTokenExpiration         = "X_ETHERNITI_TOKEN_EXPIRATION"
	XEthernitiSwaggerAddress          = "X_ETHERNITI_SWAGGER_ADDRESS"

	// proxy listener configuration
	XEthernitiListeningAddress   = "X_ETHERNITI_LISTENING_ADDRESS"
	XEthernitiListeningPort      = "X_ETHERNITI_LISTENING_PORT"
	XEthernitiListeningInterface = "X_ETHERNITI_LISTENING_INTERFACE"
	XEthernitiListeningMode      = "X_ETHERNITI_LISTENING_MODE"

	// infura service configuration
	XEthernitiInfuraToken = "X_ETHERNITI_INFURA_TOKEN"

	//gmail email service configuration
	XEthernitiEmailUsername    = "X_ETHERNITI_EMAIL_USERNAME"
	XEthernitiGmailAccessToken = "X_ETHERNITI_GMAIL_ACCESS_TOKEN"
	XEthernitiEmailServer      = "X_ETHERNITI_EMAIL_SERVER"
	XEthernitiEmailServerOnly  = "X_ETHERNITI_EMAIL_SERVER_ONLY"

	//sendgrid email service configuration
	XEthernitiSendgridApiKey = "X_ETHERNITI_SENDGRID_API_KEY"

	// worker module configuration
	XEthernitiMaxWorkers = "X_ETHERNITI_MAX_WORKERS"
	XEthernitiMaxQueue   = "X_ETHERNITI_MAX_QUEUE"
)
