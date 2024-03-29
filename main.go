// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package main

import (
	"os"
	"runtime"

	"github.com/zerjioang/etherniti/core/controllers/index"

	"github.com/zerjioang/etherniti/core/bus"
	"github.com/zerjioang/etherniti/core/cmd"
	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/shared/notifier"
	"github.com/zerjioang/go-hpc/lib/fdlimit"
	"github.com/zerjioang/go-hpc/lib/pibench"

	"github.com/zerjioang/etherniti/util/banner"

	"github.com/olekukonko/tablewriter"
	"github.com/zerjioang/etherniti/core/logger"
)

var (
	//build-time variables
	// default values when no data is found

	// Version is the built-variable that indicates compiled version data
	Version = "latest"

	// Commit is the built-variable that indicates compiled code commit
	Commit = "latest"

	// Edition is the built-variable that indicates etherniti edition data
	Edition = "oss"

	// build commit hash value
	notifierChan = make(chan error, 1)
)

func init() {
	//pass build-time variables to banner package
	// in order to print in the welcome banner
	banner.Version = Version
	banner.Commit = Commit
	banner.Edition = Edition
	// show welcome banner
	println(banner.WelcomeBanner())
	// load application internal constants
	index.LoadIndexConstants()
	// detect if the application is executing inside of docker container or not
	isDocker := config.IsDocker()
	if isDocker {
		logger.Info("etherniti is running inside docker instance")
	}
	// hack 1
	// Increase resources limitations
	// adds logic to increase the soft limit on the max number of open files for the server process
	currentVal, raiseErr := fdlimit.RaiseMax()
	if raiseErr != nil {
		logger.Error(raiseErr)
	}
	logger.Info("current server fd limit established at: ", currentVal)

	// hack 2
	// #!/bin/bash
	// echo 2621440 >  /proc/sys/net/netfilter/nf_conntrack_max
}

func Launcher() {
	/*
		The GOMAXPROCS variable limits the number of operating system threads
		that can execute user-level Go code simultaneously. There is no limit
		to the number of threads that can be blocked in system calls on behalf
		of Go code; those do not count against the GOMAXPROCS limit. This package's
		GOMAXPROCS function queries and changes the limit.

		It is advised in most of the cases, to run all your goroutines on one core
		but if you need to divide goroutines among available CPU cores of your system,
		you can use GOMAXPROCS environment variable or call to runtime using function
		runtime.GOMAXPROCS(n) where n is the number of cores to use.
	*/
	//set as default value
	max := runtime.NumCPU()
	logger.Info("setting GOMAXPROCS value to ", max)
	runtime.GOMAXPROCS(max)

	// read server benchmark evaluation function based on montecarlo pi generator
	logger.Info("current server runtime benchmark score: ", pibench.GetScore(), " points")

	//run the server
	cmd.RunServer(notifierChan)
	err := <-notifierChan
	if err != nil {
		logger.Error("failed to execute etherniti proxy: ", err)
		//print error details in a table
		defer showErrorInformation(err)
	}
	logger.Info("shutting down remaining modules")
	// finish graceful shutdown
	bus.SharedBus().Emit(notifier.PowerOffEvent)
	bus.SharedBus().Shutdown()
	logger.Info("all systems securely shutdown. exiting")
}

func showErrorInformation(e error) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Error", "Description"})
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgRedColor, tablewriter.FgHiWhiteColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgRedColor, tablewriter.FgHiWhiteColor},
	)
	table.Append([]string{"failed to execute etherniti proxy", e.Error()})
	table.Render()
}

// generate build sha1: git rev-parse --short HEAD
// compile passing -ldflags "-X main.Build <build sha1>"
// example: go build -ldflags "-X main.Build a1064bc" example.go
func main() {
	Launcher()
}
