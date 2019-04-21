// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package base58

import (
	"github.com/zerjioang/etherniti/core/modules/encoding"
)

const (
	base62Alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

var (
	// Base62 represents bytes as a common-62 number [0-9A-Za-z].
	Base62 = encoding.NewEncoding(base62Alphabet)
)
