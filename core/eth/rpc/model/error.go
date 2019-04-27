// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package model

import "strconv"

// EthError - ethereum error
type EthError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err EthError) Error() string {
	return "an error occurred: " + err.Message + " with code " + strconv.Itoa(err.Code)
}