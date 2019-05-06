// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package id

type UniqueId [32]byte

func (uid UniqueId) String() string {
	return string(uid[:])
}
