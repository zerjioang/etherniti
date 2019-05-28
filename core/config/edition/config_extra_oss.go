// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

// +build oss

package edition

import (
	"github.com/zerjioang/etherniti/core/logger"
)


// additional configuration setup for open source (oss) edition
func ExtraSetup() error {
	logger.Debug("adding addition configuration for Open Source Community edition")
	return nil
}