// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

// +build oss

package edition

import (
	"github.com/zerjioang/etherniti/core/logger"
	"github.com/zerjioang/etherniti/shared/constants"
)

// get current edition details
// atomic/thread-safe
func Edition() constants.Edition {
	return constants.OpenSource
}

// additional configuration setup for open source (oss) edition
func ExtraSetup() error {
	logger.Debug("adding addition configuration for Open Source Community edition")
	return nil
}
