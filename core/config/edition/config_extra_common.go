// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package edition

import "github.com/zerjioang/etherniti/shared/constants"

// check if active edition is opensource
// atomic/thread-safe
func IsOpenSource() bool {
	return Edition() == constants.OpenSource
}

// check if active edition is pro
// atomic/thread-safe
func IsEnterprise() bool {
	return Edition() == constants.Enterprise
}
