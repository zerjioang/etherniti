// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package main

import (
	"fmt"

	"github.com/zerjioang/etherniti/core/util"

	"github.com/zerjioang/etherniti/core"
)

func init() {
	fmt.Println(util.WelcomeBanner())
}

func main() {
	core.NewDeployer().Run()
}
