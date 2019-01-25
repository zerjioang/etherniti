// Copyright MethW
// SPDX-License-Identifier: Apache License 2.0

package util

import (
	"github.com/satori/go.uuid"
)

func GenerateUUID() string {
	u := uuid.NewV4()
	return u.String()
}
