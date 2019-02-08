// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package util

const (
	bannerTemplate = `
        __  .__                        .__  __  .__ 
  _____/  |_|  |__   ___________  ____ |__|/  |_|__|
_/ __ \   __\  |  \_/ __ \_  __ \/    \|  \   __\  |
\  ___/|  | |   Y  \  ___/|  | \/   |  \  ||  | |  |
 \___  >__| |___|  /\___  >__|  |___|  /__||__| |__|
     \/          \/     \/           \/             
                 
                 etherniti loading...
                 arch: $GOARCH
                 ver : $VER
`
	version = "0.0.1"
)

var (
	banner = ""
)

func init(){
	banner = getBannerFromTemplate()
}
func WelcomeBanner() string {
	return banner
}
func getBannerFromTemplate() string {
	banner = strings.Replace(banner, "$GOARCH", os.GOARCH, -1)
	banner = strings.Replace(banner, "$VER", version, -1)
	return banner
}

