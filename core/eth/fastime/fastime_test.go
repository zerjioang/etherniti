// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package fastime

import (
	"testing"
	"time"
)

func TestFastTime(t *testing.T) {

	t.Run("duration", func(t *testing.T) {
		var d Duration
		d = Nanosecond * 200
		if d.Nanoseconds() != 200 {
			t.Error("failed to get nanoseconds")
		}
	})

	t.Run("now", func(t *testing.T) {
		tm1 := Now()
		if tm1.sec > 0 {

		}
	})
	t.Run("add", func(t *testing.T) {
		tm2 := Now()
		u := tm2.Add(Nanosecond * 200)
		t.Log(u.Unix())
		t.Log(tm2.Unix())
		if u.Unix()-tm2.Unix() != 0 {
			t.Error("failed to add time")
		}
	})
	t.Run("unix", func(t *testing.T) {
		tm2 := Now()
		u := tm2.Unix()
		if u > 0 {

		}
	})
}

func TestStandardTime(t *testing.T) {

	t.Run("standard-now", func(t *testing.T) {
		tm3 := time.Now()
		t.Log(tm3)
	})
	t.Run("standard-now-unix", func(t *testing.T) {
		tm4 := time.Now()
		u := tm4.Unix()
		t.Log(u)
	})
}