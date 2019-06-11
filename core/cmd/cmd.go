// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package cmd

import (
	"errors"
	"sync/atomic"

	"github.com/zerjioang/etherniti/core/config/edition"
	"github.com/zerjioang/etherniti/core/listener"

	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/controllers"
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
		// load etherniti proxy configuration
		opts := config.GetDefaultOpts()

		// setup current execution environment
		err := config.Setup(opts)
		if err != nil {
			// env error configuration found
			notifier <- err
			return
		}
		// 2 setup additional
		// run additional extra configuration depending on compiled edition
		extraErr := edition.ExtraSetup()
		if extraErr != nil {
			// edition error configuration found
			notifier <- extraErr
			return
		}
		// 3 check proxy server configuration
		configErr := config.CheckConfiguration(opts)
		if configErr != nil {
			// proxy configuration error configuration found
			notifier <- configErr
			return
		}
		// 4 get listening mode
		logger.Info("starting etherniti proxy listener with requested mode")
		mode := opts.ServiceListeningMode()

		// 5 update value
		serverStarted.Store(true)

		// 6 run listener
		go listener.FactoryListener(mode).Listen(notifier)
	} else {
		notifier <- errAlreadyStarted
	}
}
