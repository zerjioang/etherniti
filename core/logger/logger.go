// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package logger

import "github.com/labstack/gommon/log"

// global logger for internal procedures
var (
	// custom internal logger with custom format
	customLog     *log.Logger
	defaultHeader = `{"time":"${time_unix_nano}","level":"${level}","prefix":"${prefix}"}`
)

func init() {
	// configure error log
	customLog = log.New("internal")
	customLog.SetHeader(defaultHeader)
	customLog.SetLevel(log.INFO)
}

// custom warn format logger
func Warn(i ...interface{}) {
	customLog.Warn(i...)
}

// custom error format logger
func Error(i ...interface{}) {
	customLog.Error(i...)
}

// custom info format logger
func Info(i ...interface{}) {
	customLog.Info(i...)
}
