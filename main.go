// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package main

import (
	"github.com/zerjioang/etherniti/core/cmd"
	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/util/banner"
	"github.com/zerjioang/etherniti/thirdparty/gommon/log"
)

var (
	// build commit hash value
	Build    = config.GetEnvironmentName()
	notifier = make(chan error, 1)
)

func init() {
	banner.Commit = Build
}

// generate build sha1: git rev-parse --short HEAD
// compile passing -ldflags "-X main.Build <build sha1>"
// example: go build -ldflags "-X main.Build a1064bc" example.go
func main() {
	cmd.RunServer(notifier)
	err := <-notifier
	if err != nil {
		log.Fatal("failed to execute etherniti proxy:", err)
	}
	<-notifier
}
