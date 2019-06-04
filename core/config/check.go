package config

import (
	"errors"

	"github.com/zerjioang/etherniti/shared/def/listener"

	"github.com/zerjioang/etherniti/core/logger"
)

const (
	infuraKeyErr     = "invalid infura token provided. Make sure your environment has correctly setup key X_ETHERNITI_INFURA_TOKEN"
	infuraKeyLenErr  = "invalid infura token provided. Make sure your provided infura key contains 32 chars on it. Check X_ETHERNITI_INFURA_TOKEN"
	logLevelErr      = "invalid log level provided. Make sure your environment has correctly setup key X_ETHERNITI_LOG_LEVEL. Allowed values are: debug, info, warn, error, off"
	listeningModeErr = "invalid listening mode provided. Make sure your environment has correctly setup key X_ETHERNITI_LISTENING_MODE. Allowed values are: http, https, socket"
)

func CheckConfiguration() error {
	logger.Info("checking etherniti proxy server configuration before full startup")

	// check log level
	if LogLevelStr() == "" {
		logger.Error(logLevelErr)
		return errors.New(logLevelErr)
	}

	// check log enabled
	if EnableLoggingStr() == "" {
		logger.Warn("proxy logging status is not defined. Make sure your environment has correctly setup key ", XEthernitiEnableLogging, ". Allowed values are: true, false")
	}

	if EnableSecureMode() {
		logger.Info("[INFO] enabling secure mode")
	} else {
		logger.Warn("[WARNING] secure mode is disabled")
	}

	if !EnableCors() {
		logger.Warn("[WARNING] CORS is disabled")
	}

	if !EnableRateLimit() {
		logger.Warn("[WARNING] rate limit is disabled")
	} else {
		logger.Warn("[WARNING] rate limit is enabled")
	}

	// check infura token
	if InfuraToken() == "" {
		logger.Error(infuraKeyErr)
		return errors.New(infuraKeyErr)
	}
	if len(InfuraToken()) != 32 {
		logger.Error(infuraKeyLenErr)
		return errors.New(infuraKeyLenErr)
	}
	if BlockTorConnections == false {
		logger.Warn("[WARNING] block of tor based connections is disabled")
	}

	// check gmail smtp
	if GetEmailUsername() == "" {
		logger.Warn("[WARNING] Missing GMAIL username. Some or all emailing features wont work as expected")
	}
	if GetEmailPassword() == "" {
		logger.Warn("[WARNING] Missing GMAIL password or access token. Some or all emailing features wont work as expected")
	}
	if GetEmailServer() == "" {
		logger.Warn("[WARNING] Missing GMAIL server configuration. Some or all emailing features wont work as expected")
	}
	if GetEmailServerOnly() == "" {
		logger.Warn("[WARNING] Missing GMAIL configuration. Some or all emailing features wont work as expected")
	}

	// check sendgrid api key
	if SendGridApiKey() == "" {
		logger.Warn("[WARNING] SendGrid API KEY not found. Make sure your environment has correctly setup key ", XEthernitiSendgridApiKey)
	}

	// proxy listener configuration checks

	// check swagger address
	if GetSwaggerAddress() == "" {
		logger.Warn("[WARNING] missing swagger address. Make sure your environment has correctly setup key ", XEthernitiSwaggerAddress)
		return errors.New("missing swagger address. Make sure your environment has correctly setup key " + XEthernitiSwaggerAddress)
	}
	// check listening address
	if GetListeningAddress() == "" {
		logger.Warn("[WARNING] missing http listening address. Make sure your environment has correctly setup key ", XEthernitiListeningAddress)
		return errors.New("missing http listening address. Make sure your environment has correctly setup key " + XEthernitiListeningAddress)
	}
	// check listening port
	if GetListeningPort() < 1024 {
		logger.Warn("[WARNING] selected listening port may require privileged access. Make sure your environment has correctly setup key ", XEthernitiListeningPort)
		return errors.New("selected listening port may require privileged access. Make sure your environment has correctly setup key " + XEthernitiListeningPort)
	}
	// check listening port string
	if GetListeningPortStr() == "" {
		msg := "selected listening port is not set. Make sure your environment has correctly setup key " + XEthernitiListeningPort
		logger.Warn(msg)
		return errors.New(msg)
	}
	// check listening mode
	if ServiceListeningModeStr() == "" {
		logger.Error(listeningModeErr)
		return errors.New(listeningModeErr)
	}

	if ServiceListeningMode() == listener.UnknownMode {
		logger.Error(listeningModeErr)
		return errors.New(listeningModeErr)
	}

	// check http listening interface
	if GetHttpInterface() == "" {
		logger.Warn("[WARNING] missing http listening interface. Make sure your environment has correctly setup key ", XEthernitiListeningInterface)
		return errors.New("missing http listening interface. Make sure your environment has correctly setup key " + XEthernitiListeningInterface)
	}

	// check worker module config
	return checkWorkerModule()
}

func checkWorkerModule() error {
	// worker module configuration
	if MaxWorker <= 0 {
		logger.Warn("[WARNING] Invalid ", XEthernitiMaxWorkers, " value found. It must be bigger than 0")
	}
	if MaxQueue <= 0 {
		logger.Warn("[WARNING] Invalid ", XEthernitiMaxQueue, " value found. It must be bigger than 0")
	}
	return nil
}
