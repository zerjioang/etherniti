// Copyright gaethway
// SPDX-License-Identifier: Apache License 2.0

package util

import (
	"github.com/satori/go.uuid"
)

func GenerateUUID() string {
	return uuid.NewV4().String()
}
