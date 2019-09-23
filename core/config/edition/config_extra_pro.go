// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

// +build pro

package edition

import (
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/shared/constants"
)

// get current edition details
// atomic/thread-safe
func Edition() constants.Edition {
	return constants.Pro
}

// additional configuration setup for Pro/Subscription (pro) edition
func ExtraSetup() error {
	logger.Debug("adding addition configuration for enterprise edition")
	return nil
}
