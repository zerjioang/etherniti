// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package solc

type CompilerMode string

const (
	SingleRawFile    = "single-raw"
	SingleBase64File = "single-base64"
	GitMode          = "git"
	ZipMode          = "zip"
	TargzMode        = "targz"
)
