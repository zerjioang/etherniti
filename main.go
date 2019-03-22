// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package main

import (
	"github.com/zerjioang/etherniti/core/constants"
	"github.com/zerjioang/etherniti/core/listener"

	"github.com/zerjioang/etherniti/core/logger"

	"github.com/labstack/gommon/log"
	"github.com/zerjioang/etherniti/core/config"

	"github.com/zerjioang/etherniti/core/util"
)

var (
	// build commit hash value
	Build string
)

func init() {
	util.Commit = Build
	println(util.WelcomeBanner())
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
		log.Error("failed to execute http server:", err)
	}
}
