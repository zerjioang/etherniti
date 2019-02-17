package logger

import "github.com/labstack/gommon/log"

// global logger for errors
var (
	ErrorLog      = log.New("-")
	defaultHeader = `{"time":"${time_rfc3339_nano}","level":"${level}","prefix":"${prefix}",` +
		`"file":"${short_file}","line":"${line}"}`
)

func init() {
	ErrorLog.set
}
