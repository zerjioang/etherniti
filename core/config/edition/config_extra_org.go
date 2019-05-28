// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

// +build org

package edition

import (
	"github.com/zerjioang/etherniti/core/logger"
)

// additional configuration setup for etherniti.org environment
func ExtraSetup() error {
	logger.Debug("adding addition configuration for Etherniti Public Proxy edition")
	return nil
}