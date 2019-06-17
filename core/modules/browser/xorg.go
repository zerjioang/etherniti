package browser

import (
	"bytes"
	"os/exec"
)

var (
	hasUI bool
	xorg = []byte("Xorg")
	xwayland = []byte("Xwayland")
)
func init(){
	hasUI = detectUI()
}

// this function will detect if current environment has a gui or not
func detectUI() bool {
	// ps -e | grep X
	// 2941 tty1     00:00:00 Xwayland
	// 5192 tty2     00:07:29 Xorg
	c := exec.Command("bash", "-c", "ps -e | grep X")
	out, err := c.Output()
	if err == nil {
		hasUI = bytes.Contains(out, xorg) || bytes.Contains(out, xwayland)
	}
	return hasUI
}

func HasGraphicInterface() bool {
	return hasUI
}
