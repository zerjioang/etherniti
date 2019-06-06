// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package edition

import "github.com/zerjioang/etherniti/shared/constants"

var (
	isOpen, isPro, isValid bool
)
// check if active edition is opensource
func init(){
	e := Edition()
	isOpen = e == constants.OpenSource
	isPro = e == constants.Enterprise
	isValid = e != constants.Unknown
}
// atomic/thread-safe
func IsOpenSource() bool {
	return isOpen
}

// check if active edition is pro
// atomic/thread-safe
func IsEnterprise() bool {
	return isPro
}

// check if active edition is valid or not
// atomic/thread-safe
func IsValidEdition() bool {
	return isValid
}
