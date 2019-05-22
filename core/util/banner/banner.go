// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package banner

import (
	"runtime"
	"strings"

	"github.com/zerjioang/etherniti/core/logger"

	"github.com/zerjioang/etherniti/shared/constants"
)

const (
	bannerTemplate = `
        __  .__                        .__  __  .__ 
  _____/  |_|  |__   ___________  ____ |__|/  |_|__|
_/ __ \   __\  |  \_/ __ \_  __ \/    \|  \   __\  |
\  ___/|  | |   Y  \  ___/|  | \/   |  \  ||  | |  |
 \___  >__| |___|  /\___  >__|  |___|  /__||__| |__|
     \/          \/     \/           \/

  Official Page  :        https://www.etherniti.org
  Github Page    :     https://github.com/etherniti
  Issues         : https://github.com/etherniti/rfc

  Build information:

  arch             : $GOARCH
  go/root          : $GOROOT
  go/version       : $GO_VERSION
  go/compiler      : $GO_COMPILER
  proxy/version    : $VER
  proxy/commit     : $COMMIT
  proxy/edition    : $EDITION

`
)

var (
	banner = ""
	// compilation built-in: commit
	Commit = ""
	// compilation built-in: edition
	// proxy edition: oss or pro
	Edition = ""
)

// thread safe init function
func init() {
	logger.Debug("loading banner module data")
	banner = getBannerFromTemplate()
}

// return welcome banner asci art message
func WelcomeBanner() string {
	logger.Debug("reading welcome banner")
	banner = strings.Replace(banner, "$COMMIT", Commit, 1)
	banner = strings.Replace(banner, "EDITION", Edition, 1)
	return banner
}

// inject runtime variables in welcome banner message
func getBannerFromTemplate() string {
	logger.Debug("generating welcome banner from template")
	banner = strings.Replace(bannerTemplate, "$GOARCH", runtime.GOARCH, 1)
	banner = strings.Replace(banner, "$GOARCH", runtime.GOOS, 1)
	banner = strings.Replace(banner, "$GOROOT", runtime.GOROOT(), 1)
	banner = strings.Replace(banner, "$GO_VERSION", runtime.Version(), 1)
	banner = strings.Replace(banner, "$GO_COMPILER", runtime.Compiler, 1)
	banner = strings.Replace(banner, "$VER", constants.Version, 1)
	return banner
}
