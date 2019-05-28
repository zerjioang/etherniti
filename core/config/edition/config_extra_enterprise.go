// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

// +build enterprise

package edition

import (
	"github.com/zerjioang/etherniti/core/logger"
)


// additional configuration setup for Enterprise/Subscription (enterprise) edition
func ExtraSetup() error {
	logger.Debug("adding addition configuration for Enterprise edition")
	return nil
}