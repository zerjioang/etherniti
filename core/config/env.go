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
	logger.Debug("accessing to environment configuration")
	doOnce.Do(func() {
		logger.Debug("reading environment configuration")
		globalCfg = newEnvironment()
		globalCfg.SetDefaults()
		// override default values with user provided data
		globalCfg.read()
	})
	return *globalCfg
}

func GetEnvironmentPtr() *EnvConfig {
	logger.Debug("accessing to environment configuration")
	doOnce.Do(func() {
		logger.Debug("reading environment configuration")
		globalCfg = newEnvironment()
		globalCfg.SetDefaults()
		// override default values with user provided data
		globalCfg.read()
	})
	return globalCfg
}

func ReadEnvironment(key string) (interface{}, bool) {
	logger.Debug("reading environment key")
	v, ok := GetEnvironment().data[key]
	return v, ok
}

func ReadEnvironmentString(key string) string {
	logger.Debug("reading environment key as string")
	v, found := ReadEnvironment(key)
	if !found {
		return ""
	}
	return v.(string)
}

// configuration data type
type EnvConfig struct {
	data map[string]interface{}
}

// Constructor like function for new Env config data wrappers
func newEnvironment() *EnvConfig {
	logger.Debug("creating new enviroment data")
	cfg := new(EnvConfig)
	cfg.data = make(map[string]interface{})
	return cfg
}

//read environment variables
func (c *EnvConfig) read() {
	logger.Debug("reading environment variables from current operating system")
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if len(pair) == 2 {
			k := pair[0]
			v := pair[1]
			logger.Debug(k, " = ", v)
			c.data[k] = v
		}
	}
}
