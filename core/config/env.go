// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package config

import (
	"os"
	"strings"
	"sync"

	"github.com/zerjioang/etherniti/core/logger"
)

var doOnce sync.Once

// single point of object access in an object-oriented application
var globalCfg *EnvConfig

func GetEnvironment() EnvConfig {
	doOnce.Do(func() {
		logger.Debug("reading environment configuration")
		globalCfg = newEnvironment()
		SetDefaults(globalCfg)
		// override default values with user provided data
		globalCfg.read()
	})
	logger.Debug("accessing to environment configuration")
	return *globalCfg
}

func ReadEnvironment(key string) interface{} {
	return GetEnvironment().data[key]
}

// configuration data type
type EnvConfig struct {
	data map[string]interface{}
}

// Constructor like function for new Env config data wrappers
func newEnvironment() *EnvConfig {
	cfg := new(EnvConfig)
	cfg.data = make(map[string]interface{})
	return cfg
}

//read environment variables
func (env EnvConfig) read() {
	logger.Debug("reading environment variables from current operating system")
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if len(pair) == 2 {
			k := pair[0]
			v := pair[1]
			logger.Debug(k, " = ", v)
			env.data[k] = v
		}
	}
}
