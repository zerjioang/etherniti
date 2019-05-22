// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package config

import (
	"os"
	"strings"

	"github.com/zerjioang/etherniti/core/logger"
)

func GetEnvironment() *EnvConfig {
	return proxyEnv
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

//readEnvironmentData environment variables
func (c *EnvConfig) readEnvironmentData() {
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
func (c EnvConfig) String(key string) string {
	v, ok := c.Read(key)
	if ok {
		str, valid := v.(string)
		if valid {
			return str
		}
	}
	return ""
}

func (c EnvConfig) Read(key string) (interface{}, bool) {
	v, ok := c.data[key]
	return v, ok
}
