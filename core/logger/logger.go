// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package logger

import "github.com/zerjioang/go-hpc/thirdparty/gommon/log"

// global logger for internal procedures
var (
	// custom internal logger with custom format
	customLog     *log.Logger
	defaultHeader = `{"prefix":"${prefix}","time":"${time_unix_nano}","level":"${level}"}`
)

func init() {
	// configure error log
	customLog = log.New("internal")
	customLog.SetHeader(defaultHeader)
	customLog.SetLevel(log.DEBUG)
}

func Enabled(status bool) {
	if status {
		customLog.SetLevel(log.DEBUG)
	} else {
		customLog.SetLevel(log.OFF)
	}
}

func Level(v log.Lvl) {
	customLog.SetLevel(v)
}

// custom warn format logger
func Warn(i ...interface{}) {
	customLog.Warn(i...)
}

// custom error format logger
func Error(i ...interface{}) {
	customLog.Error(i...)
}

// custom error format logger
func Fatal(i ...interface{}) {
	customLog.Fatal(i...)
}

// custom info format logger
func Info(i ...interface{}) {
	customLog.Info(i...)
}

// custom info format logger
func Debug(i ...interface{}) {
	customLog.Debug(i...)
}
