package mail

import (
	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/go-hpc/util/str"
)

var (
	ConfirmEmailTemplate  string
	NewLoginEmailTemplate string
	RecoverEmailTemplate  string
)

func init() {
	logger.Debug("loading email templates...")
	ConfirmEmailTemplate = str.ReadFileAsString(config.ResourcesDirInternalEmail + "/confirm.html")
	NewLoginEmailTemplate = str.ReadFileAsString(config.ResourcesDirInternalEmail + "/login.html")
	RecoverEmailTemplate = str.ReadFileAsString(config.ResourcesDirInternalEmail + "/recover.html")
}
