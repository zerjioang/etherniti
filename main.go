// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package main

import (
	"github.com/zerjioang/etherniti/core/constants"
	"github.com/zerjioang/etherniti/core/listener"
	"github.com/zerjioang/etherniti/core/util/banner"

	"github.com/zerjioang/etherniti/core/logger"

	"github.com/labstack/gommon/log"
	"github.com/zerjioang/etherniti/core/config"
)

var (
	// build commit hash value
	Build = config.EnvironmentName
)

func init() {
	banner.Commit = Build
	logger.Info("system running with pointers size of: ", constants.PointerSize, "bits")
}

// generate build sha1: git rev-parse --short HEAD
// compile passing -ldflags "-X main.Build <build sha1>"
// example: go build -ldflags "-X main.Build a1064bc" example.go
func main() {

	// 1 read environment variables
	envars := map[string]interface{}{}
	config.SetDefaults(envars)
	config.Read(envars)

	// 2 get listening mode from env vars
	logger.Info("starting etherniti proxy")
	mode := config.ServiceListeningMode()

	config.Setup()

	// 3 run listener
	err := listener.FactoryListener(mode).Listen()
	if err != nil {
		log.Fatal("failed to execute http server:", err)
	}
}
