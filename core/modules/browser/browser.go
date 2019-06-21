package browser

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/zerjioang/etherniti/core/logger"
)

// open a link in your favorite browser
func OpenBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		logger.Error("failed to open browser window: ", err)
	}
}
