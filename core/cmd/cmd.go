package cmd

import (
	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/handlers"
	"github.com/zerjioang/etherniti/core/listener"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/shared/constants"
)

func init() {
	handlers.LoadIndexConstants()
	logger.Info("system running with pointers size of: ", constants.PointerSize, " bits")
}

func RunServer() error {
	// setup current execution environment
	config.Setup()

	// 2 get listening mode
	logger.Info("starting etherniti proxy listener with requested mode")
	mode := config.ServiceListeningMode()

	// 3 run listener
	return listener.FactoryListener(mode).Listen()
}
