package config

import (
	"errors"

	"github.com/zerjioang/etherniti/core/logger"
)

const (
	infuraKeyErr     = "invalid infura token provided. Make sure your environment has correctly setup key X_ETHERNITI_INFURA_TOKEN"
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
		logger.Warn("proxy logging status is not defined. Make sure your environment has correctly setup key X_ETHERNITI_ENABLE_LOGGING. Allowed values are: true, false")
	}

	// check infura token
	if InfuraToken() == "" {
		logger.Error(infuraKeyErr)
		return errors.New(infuraKeyErr)
	}
	if BlockTorConnections == false {
		logger.Warn("[WARNING] block of tor based connections is disabled")
	}

	// check gmail smtp
	if GetEmailPassword() == "" {
		logger.Warn("[WARNING] Missing gmail username. Some or all emailing features wont work as expected")
	}
	if GetEmailPassword() == "" {
		logger.Warn("[WARNING] Missing gmail password or access token. Some or all emailing features wont work as expected")
	}
	if GetEmailServer() == "" {
		logger.Warn("[WARNING] Missing gmail server configuration. Some or all emailing features wont work as expected")
	}
	if GetEmailServerOnly() == "" {
		logger.Warn("[WARNING] Missing gmail configuration. Some or all emailing features wont work as expected")
	}

	// check sendgrid api key
	if SendGridApiKey() == "" {
		logger.Warn("[WARNING] SendGrid API KEY not found")
	}

	// check listening mode
	if ServiceListeningModeStr() == "" {
		logger.Error(listeningModeErr)
		return errors.New(listeningModeErr)
	}
	return nil
}
