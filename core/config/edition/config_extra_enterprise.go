// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

// +build pro

package edition

import (
	"github.com/zerjioang/etherniti/core/logger"
)

// additional configuration setup for Pro/Subscription (pro) edition
func ExtraSetup() error {
	logger.Debug("adding addition configuration for Pro edition")
	return nil
}