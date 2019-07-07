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

func NewHashUint32Set() HashUint32Set {
	return HashUint32Set{}
}

func NewHashUint32SetPtr() *HashUint32Set {
	return new(HashUint32Set)
}
