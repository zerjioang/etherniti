// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package main

import (
	"runtime"

	"github.com/zerjioang/etherniti/core/util/banner"

	"github.com/zerjioang/etherniti/core/bus"

	"github.com/zerjioang/etherniti/core/cmd"
	"github.com/zerjioang/etherniti/core/logger"
)

var (
	//build-time variables
	Commit  = ""
	Edition = ""

	// build commit hash value
	notifier = make(chan error, 1)
)

func init() {
	//pass build-time variables to banner package
	// in order to print in the welcome banner
	banner.Commit = Commit
	banner.Edition = Edition
}

// generate build sha1: git rev-parse --short HEAD
// compile passing -ldflags "-X main.Build <build sha1>"
// example: go build -ldflags "-X main.Build a1064bc" example.go
func main() {
	/*
		The GOMAXPROCS variable limits the number of operating system threads
		that can execute user-level Go code simultaneously. There is no limit
		to the number of threads that can be blocked in system calls on behalf
		of Go code; those do not count against the GOMAXPROCS limit. This package's
		GOMAXPROCS function queries and changes the limit.
	*/
	//set as default value
	max := runtime.NumCPU()
	logger.Info("setting GOMAXPROCS value to ", max)
	runtime.GOMAXPROCS(max)
	//run the server
	cmd.RunServer(notifier)
	err := <-notifier
	if err != nil {
		logger.Error("failed to execute etherniti proxy: ", err)
	}
	logger.Info("shutting down remaining modules")
	// finish graceful shutdown
	bus.SharedBus().Emit(bus.PowerOffEvent)
	bus.SharedBus().Shutdown()
	logger.Info("all systems securely shutdown. exiting")
}
