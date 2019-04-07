// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package trycatch

import "testing"

func BenchmarkError(b *testing.B) {
	b.Run("generate-nil", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = Nil()
		}
	})
	b.Run("generate-default", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = New("default")
		}
	})
	b.Run("none", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		stackErr := New("default")
		for n := 0; n < b.N; n++ {
			_ = stackErr.None()
		}
	})
	b.Run("error", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		stackErr := New("default")
		for n := 0; n < b.N; n++ {
			_ = stackErr.Error()
		}
	})
}
