// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package util

import (
	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/release"
	"runtime"
	"strings"
)

const (
	bannerTemplate = `
        __  .__                        .__  __  .__ 
  _____/  |_|  |__   ___________  ____ |__|/  |_|__|
_/ __ \   __\  |  \_/ __ \_  __ \/    \|  \   __\  |
\  ___/|  | |   Y  \  ___/|  | \/   |  \  ||  | |  |
 \___  >__| |___|  /\___  >__|  |___|  /__||__| |__|
     \/          \/     \/           \/             
                 
etherniti loading...

arch              : $GOARCH
goroot            : $GOROOT
goversion         : $GO_VERSION
gocompiler        : $GO_COMPILER
etherniti version : $VER
etherniti commit  : $COMMIT
`
)

var (
	banner = ""
	// compilation commit
	Commit = ""
)

// thread safe init function
func init() {
	banner = getBannerFromTemplate()
}

// return welcome banner asci art message
func WelcomeBanner() string {
	banner = strings.Replace(banner, "$COMMIT", Commit, 1)
	return banner
}

// inject runtime variables in welcome banner message
func getBannerFromTemplate() string {
	banner = strings.Replace(bannerTemplate, "$GOARCH", runtime.GOARCH, 1)
	banner = strings.Replace(banner, "$GOARCH", runtime.GOOS, 1)
	banner = strings.Replace(banner, "$GOROOT", runtime.GOROOT(), 1)
	banner = strings.Replace(banner, "$GO_VERSION", runtime.Version(), 1)
	banner = strings.Replace(banner, "$GO_COMPILER", runtime.Compiler, 1)
	banner = strings.Replace(banner, "$VER", release.Version, 1)
	return banner
}
