// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package main

import (
	"github.com/zerjioang/etherniti/core/handlers"
	"github.com/zerjioang/etherniti/core/listener"
	"github.com/zerjioang/etherniti/core/util/banner"
	"github.com/zerjioang/etherniti/shared/constants"

	"github.com/zerjioang/etherniti/core/logger"

	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/thirdparty/gommon/log"
)

var (
	// build commit hash value
	Build = config.EnvironmentName
)

func init() {
	banner.Commit = Build
	handlers.LoadIndexConstants()
	logger.Info("system running with pointers size of: ", constants.PointerSize, " bits")
}

// generate build sha1: git rev-parse --short HEAD
// compile passing -ldflags "-X main.Build <build sha1>"
// example: go build -ldflags "-X main.Build a1064bc" example.go
func main() {

	// setup current execution environment
	config.Setup()

	// 2 get listening mode
	logger.Info("starting etherniti proxy listener with requested mode")
	mode := config.ServiceListeningMode()

	// 3 run listener
	err := listener.FactoryListener(mode).Listen()
	if err != nil {
		log.Fatal("failed to execute etherniti proxy:", err)
	}
}
