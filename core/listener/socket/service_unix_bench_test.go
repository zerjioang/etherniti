// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package socket

import "testing"

func BenchmarkUnixSocketListener(b *testing.B) {
	b.Run("instantiation", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		for n := 0; n < b.N; n++ {
			NewSocketListener()
		}
	})
	b.Run("runmode-foreground", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		l := NewSocketListener()
		for n := 0; n < b.N; n++ {
			l.RunMode("", false)
		}
	})
	b.Run("runmode-background", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		l := NewSocketListener()
		for n := 0; n < b.N; n++ {
			l.RunMode("", true)
		}
	})
}
