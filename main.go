// Copyright gaethway
// SPDX-License-Identifier: Apache License 2.0

package main

import (
	"fmt"
	"log"

	"github.com/zerjioang/gaethway/core/util"

	"github.com/zerjioang/gaethway/core"
)

func init() {
	fmt.Println(util.WelcomeBanner())
}

func main() {
	err := core.NewDeployer().Run()
	if err != nil {
		log.Fatal(err)
	}
}
