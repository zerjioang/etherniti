// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package cmd

import (
	"errors"
	"sync/atomic"

	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/controllers"
	"github.com/zerjioang/etherniti/core/listener"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/shared/constants"
)

var (
	// serverStarted atomic.Value
	serverStarted     atomic.Value
	errAlreadyStarted = errors.New("already started")
)

func init() {
	controllers.LoadIndexConstants()
	logger.Info("system running with pointers size of: ", constants.PointerSize, " bits")
	serverStarted.Store(false)
}

func RunServer(notifier chan error) {
	logger.Debug("running etherniti main server")
	// 1 read value
	if !serverStarted.Load().(bool) {
		// setup current execution environment
		config.Setup()

		// 2 get listening mode
		logger.Info("starting etherniti proxy listener with requested mode")
		mode := config.ServiceListeningMode()

		// 4 update value
		serverStarted.Store(true)

		// 3 run listener
		go listener.FactoryListener(mode).Listen(notifier)
	} else {
		notifier <- errAlreadyStarted
	}
}
