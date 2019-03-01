// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package main

import (
	"fmt"

	"github.com/zerjioang/etherniti/core/handlers"

	"github.com/zerjioang/etherniti/core/util"
)

var (
	// build commit hash value
	Build string
)

func init() {
	util.Commit = Build
	fmt.Println(util.WelcomeBanner())
}

//generate build sha1: git rev-parse --short HEAD
//compile passing -ldflags "-X main.Build <build sha1>"
// example: go build -ldflags "-X main.Build a1064bc" example.go
func main() {
	handlers.NewDeployer().Run()
}
