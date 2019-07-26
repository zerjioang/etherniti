package config

import (
	"errors"

	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/core/util/id"
)

// generate proxy service admin identity
func GenerateAdmin(opts *EthernitiOptions) (string, string, error) {
	if opts == nil {
		logger.Error("failed to generate admin identity")
		return "", "", errors.New("failed to generate proxy admin identity data because could not load proxy configuration options")
	} else {
		if !opts.Admin.LoadedFromEnv {
			logger.Debug("generating admin identity")
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
