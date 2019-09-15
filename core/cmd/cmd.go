// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package cmd

import (
	"errors"
	"sync/atomic"

	"github.com/zerjioang/etherniti/core/modules/browser"

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

		// 0 generate root/superadmin proxy identity for management
		_, _, adminErr := config.GenerateAdmin(opts)
		if adminErr != nil {
			// env error configuration found
			notifier <- adminErr
			return
		}
		// 1 setup current execution environment
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

		// 5 update value
		serverStarted.Store(true)

		// 6 run listener
		listenerHandler := listener.FactoryListener(opts.ListeningMode)
		go listenerHandler.Listen(notifier)

		// 7 open web browser if requested on desktop computer
		if opts.OpenBrowserOnSuccess && browser.HasGraphicInterface() {
			uri := opts.GetURI()
			logger.Debug("opening URI in local web browser")
			logger.Debug("opening URI: ", uri)
			navErr := browser.OpenBrowser(uri)
			if navErr != nil {
				logger.Error("failed to open web browser on local navigator due to: ", navErr)
			}
		}
	} else {
		notifier <- errAlreadyStarted
	}
}
