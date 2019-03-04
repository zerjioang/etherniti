// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package protocol

type IntegrityResponse struct {
	Message   string `json:"message"`
	Millis    string `json:"millis"`
	Hash      string `json:"hash"`
	Signature string `json:"signature"`
}
