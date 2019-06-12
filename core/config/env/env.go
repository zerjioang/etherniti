// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package env

import (
	"os"
	"strconv"
	"strings"
	"sync/atomic"

	"github.com/zerjioang/etherniti/core/logger"
)

// configuration data type
type EnvConfig struct {
	l    atomic.Value
	data map[string]interface{}
}

// Constructor like function for new Env config data wrappers
func New() *EnvConfig {
	logger.Debug("creating new enviroment data")
	cfg := new(EnvConfig)
	cfg.data = make(map[string]interface{})
	cfg.l.Store(false)
	return cfg
}

//read environment ariables
func (c *EnvConfig) Load() {
	if c.l.Load() == false {
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
	//mark the env config as readed from env
	c.l.Store(true)
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

func (c EnvConfig) Lower(key string) string {
	return strings.ToLower(c.String(key))
}

func (c EnvConfig) Int(key string, fallback int) int {
	v, ok := c.Read(key)
	if ok {
		str, ok := v.(string)
		if ok {
			//successfully converted from interface to string
			//now convert from string to int
			num, err := strconv.Atoi(str)
			if err == nil {
				return num
			}
		}
	}
	return fallback
}

func (c EnvConfig) Read(key string) (interface{}, bool) {
	v, ok := c.data[key]
	return v, ok
}
