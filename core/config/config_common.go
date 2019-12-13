package config

import (
	"bytes"
	"errors"
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/util/fs"
	"github.com/zerjioang/etherniti/core/util/id"
)

var (
	isDockerContainerDetected bool
)

func init() {
	logger.Info("checking if application is running under docker environment")
	isDockerContainerDetected = bytes.Contains(fs.ReadAll(" /proc/self/cgroup"), []byte(":/docker/")) ||
		bytes.Contains(fs.ReadAll(" /proc/1/cgroup"), []byte(":/docker/")) ||
		bytes.Contains(fs.ReadAll(" /proc/1/cgroup"), []byte(":/docker/")) ||
		fs.Exists("/.dockerenv")
}

// generate proxy service admin identity
func GenerateAdmin(opts *EthernitiOptions) (string, string, error) {
	if opts == nil {
		logger.Error("failed to generate admin identity")
		return "", "", errors.New("failed to generate proxy admin identity data because could not load proxy configuration options")
	} else {
		if !opts.Admin.LoadedFromEnv {
			logger.Debug("generating admin identity since not provided by environment options")
			accessKey := id.GenerateUUIDFromEntropy()
			accessSecret := id.GenerateUUIDFromEntropy()
			logger.Warn("proxy admin: key = ", accessKey, ", secret = ", accessSecret)
			opts.Admin.Key = accessKey
			opts.Admin.Secret = accessSecret
			return accessKey, accessSecret, nil
		} else {
			logger.Info("using admin identity loaded from env")
			return opts.Admin.Key, opts.Admin.Secret, nil
		}
	}
}

// ls -ali / | sed '2!d' |awk {'print $1'}
// inside docker alpine: 6554197
// inside docker ubuntu:latest 7884073
// inside host: 2
func IsDocker() bool {
	return isDockerContainerDetected
}
