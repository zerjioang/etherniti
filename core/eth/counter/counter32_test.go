// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package counter

import "testing"

func TestCounter(t *testing.T) {
	t.Run("atomic-uint32-instantiate", func(t *testing.T) {
		_ = NewCounter32()
	})
	t.Run("atomic-uint32-get", func(t *testing.T) {
		var c1 Count32
		value := c1.Get()
		t.Log(value)
	})
	t.Run("atomic-uint32-add", func(t *testing.T) {
		var c2 Count32
		value := c2.Increment()
		t.Log(value)
	})
	t.Run("atomic-uint32-get-add", func(t *testing.T) {
		var c3 Count32
		v1 := c3.Get()
		if v1 == 0 {

		}
		v2 := c3.Increment()
		if v2 == 1 {

		}
	})
}
