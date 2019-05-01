// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package mailer

import (
	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/util/str"
)

var (
	confirmEmailTemplate  string
	newLoginEmailTemplate string
	recoverEmailTemplate  string
)

func init() {
	logger.Debug("loading email templates...")
	confirmEmailTemplate = str.ReadFileAsString(config.ResourcesDirInternalEmail + "/confim.html")
	newLoginEmailTemplate = str.ReadFileAsString(config.ResourcesDirInternalEmail + "/login.html")
	recoverEmailTemplate = str.ReadFileAsString(config.ResourcesDirInternalEmail + "/recover.html")
}
