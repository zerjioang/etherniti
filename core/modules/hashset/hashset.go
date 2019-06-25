// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package hashset

type HashSet map[string]struct{}
type HashUint32Set map[uint32]struct{}

func NewHashSet() HashSet {
	return HashSet{}
}

func NewHashSetPtr() *HashSet {
	return new(HashSet)
}
