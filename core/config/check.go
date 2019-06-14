package config

import (
	"errors"
	"strings"

	"github.com/zerjioang/etherniti/core/modules/entropy"

	"github.com/zerjioang/etherniti/shared/def/listener"

	"github.com/zerjioang/etherniti/core/logger"
)

const (
	infuraKeyErr     = "invalid infura token provided. Make sure your environment has correctly setup key X_ETHERNITI_INFURA_TOKEN"
	infuraKeyLenErr  = "invalid infura token provided. Make sure your provided infura key contains 32 chars on it. Check X_ETHERNITI_INFURA_TOKEN"
	logLevelErr      = "invalid log level provided. Make sure your environment has correctly setup key X_ETHERNITI_LOG_LEVEL. Allowed values are: debug, info, warn, error, off"
	listeningModeErr = "invalid listening mode provided. Make sure your environment has correctly setup key X_ETHERNITI_LISTENING_MODE. Allowed values are: http, https, socket"
)

func CheckConfiguration(opts *EthernitiOptions) error {
	logger.Info("checking etherniti proxy server configuration before full startup")

	// check log level
	if opts.LogLevelStr() == "" {
		logger.Error(logLevelErr)
		return errors.New(logLevelErr)
	}

	// check log enabled
	if opts.EnableLoggingStr() == "" {
		logger.Warn("proxy logging status is not defined. Make sure your environment has correctly setup key ", XEthernitiEnableLogging, ". Allowed values are: true, false")
	}

	if opts.EnableSecureMode() {
		logger.Info("[INFO] enabling secure mode")
	} else {
		logger.Warn("[WARNING] secure mode is disabled")
	}

	if !opts.EnableCors() {
		logger.Warn("[WARNING] CORS is disabled")
	}

	if !opts.EnableRateLimit() {
		logger.Warn("[WARNING] rate limit is disabled")
	} else {
		logger.Warn("[WARNING] rate limit is enabled")
	}

	// check infura token
	if opts.InfuraToken() == "" {
		logger.Warn(infuraKeyErr)
		logger.Warn("infura provider is disabled until valid token is provided")
	}
	if len(opts.InfuraToken()) != 32 {
		logger.Error(infuraKeyLenErr)
		return errors.New(infuraKeyLenErr)
	}
	if opts.BlockTorConnections == false {
		logger.Warn("[WARNING] block of tor based connections is disabled")
	}

	// check gmail smtp
	if opts.GetEmailUsername() == "" {
		logger.Warn("[WARNING] Missing GMAIL username. Some or all emailing features wont work as expected")
	}
	if opts.GetEmailPassword() == "" {
		logger.Warn("[WARNING] Missing GMAIL password or access token. Some or all emailing features wont work as expected")
	}
	if opts.GetEmailServer() == "" {
		logger.Warn("[WARNING] Missing GMAIL server configuration. Some or all emailing features wont work as expected")
	}
	if opts.GetEmailServerOnly() == "" {
		logger.Warn("[WARNING] Missing GMAIL configuration. Some or all emailing features wont work as expected")
	}

	// check sendgrid api key
	if opts.SendGridApiKey() == "" {
		logger.Warn("[WARNING] SendGrid API KEY not found. Make sure your environment has correctly setup key ", XEthernitiSendgridApiKey)
	}

	// proxy listener configuration checks

	// check swagger address
	if opts.GetSwaggerAddress() == "" {
		logger.Warn("[WARNING] missing swagger address. Make sure your environment has correctly setup key ", XEthernitiSwaggerAddress)
		return errors.New("missing swagger address. Make sure your environment has correctly setup key " + XEthernitiSwaggerAddress)
	}
	// check listening address
	if opts.GetListeningAddress() == "" {
		logger.Warn("[WARNING] missing http listening address. Make sure your environment has correctly setup key ", XEthernitiListeningAddress)
		return errors.New("missing http listening address. Make sure your environment has correctly setup key " + XEthernitiListeningAddress)
	}
	// check listening port
	if opts.GetListeningPort() < 1024 {
		logger.Warn("[WARNING] selected listening port may require privileged access. Make sure your environment has correctly setup key ", XEthernitiListeningPort)
		return errors.New("selected listening port may require privileged access. Make sure your environment has correctly setup key " + XEthernitiListeningPort)
	}
	// check listening port string
	if opts.GetListeningPortStr() == "" {
		msg := "selected listening port is not set. Make sure your environment has correctly setup key " + XEthernitiListeningPort
		logger.Warn(msg)
		return errors.New(msg)
	}
	// check listening mode
	if opts.ServiceListeningModeStr() == "" {
		logger.Error(listeningModeErr)
		return errors.New(listeningModeErr)
	}

	if opts.ServiceListeningMode() == listener.UnknownMode {
		logger.Error(listeningModeErr)
		return errors.New(listeningModeErr)
	}

	// check http listening interface
	if opts.GetHttpInterface() == "" {
		logger.Warn("[WARNING] missing http listening interface. Make sure your environment has correctly setup key ", XEthernitiListeningInterface)
		return errors.New("missing http listening interface. Make sure your environment has correctly setup key " + XEthernitiListeningInterface)
	}
	//check eth mainnets and testnets endpoints
	if opts.RopstenCustomEndpoint == "" {
		logger.Warn("[WARNING] missing Ropsten network custom endpoint. Default infura endpoint will be used if token provided")
	} else if !isValidUrl(opts.RopstenCustomEndpoint) {
		msg := "ropsten network custom endpoint is not a valid URL"
		logger.Error(msg)
		return errors.New(msg)
	}
	if opts.RinkebyCustomEndpoint == "" {
		logger.Warn("[WARNING] missing Rinkeby network custom endpoint. Default infura endpoint will be used if token provided")
	} else if !isValidUrl(opts.RinkebyCustomEndpoint) {
		msg := "rinkeby network custom endpoint is not a valid URL"
		logger.Error(msg)
		return errors.New(msg)
	}
	if opts.KovanCustomEndpoint == "" {
		logger.Warn("[WARNING] missing Kovan network custom endpoint. Default infura endpoint will be used if token provided")
	} else if !isValidUrl(opts.KovanCustomEndpoint) {
		msg := "kovan network custom endpoint is not a valid URL"
		logger.Error(msg)
		return errors.New(msg)
	}
	if opts.MainnetCustomEndpoint == "" {
		logger.Warn("[WARNING] missing Mainnet network custom endpoint. Default infura endpoint will be used if token provided")
	} else if !isValidUrl(opts.MainnetCustomEndpoint) {
		msg := "mainnet network custom endpoint is not a valid URL"
		logger.Error(msg)
		return errors.New(msg)
	}

	//check current system entropy
	if entropy.HasCriticalValue() {
		logger.Warn("[WARNING] current system entropy available bytes are in critical value. we recommend you to increase ssytem entropy for security purposes")
	}
	// check worker module config
	return checkWorkerModule(opts)
}

func checkWorkerModule(opts *EthernitiOptions) error {
	// worker module configuration
	if opts.MaxWorker <= 0 {
		logger.Warn("[WARNING] Invalid ", XEthernitiMaxWorkers, " value found. It must be bigger than 0")
	}
	if opts.MaxQueue <= 0 {
		logger.Warn("[WARNING] Invalid ", XEthernitiMaxQueue, " value found. It must be bigger than 0")
	}
	return nil
}

func isValidUrl(url string) bool {
	if url == "" {
		return false
	}
	if !(strings.HasPrefix(url, "https://") || strings.HasPrefix(url, "http://")) {
		return false
	}
	return true
}
