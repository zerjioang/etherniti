// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package fastime

import (
	"testing"
	"time"
)

func BenchmarkFastTime(b *testing.B) {

	b.Run("fastime-now", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = Now()
		}
	})
	b.Run("fastime-now-unix", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = Now().Unix()
		}
	})
	b.Run("fastime-now-nanos", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = Now().Nanos()
		}
	})
	b.Run("fastime-from-time-1", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		nt := time.Now()
		for n := 0; n < b.N; n++ {
			_ = FromTime(nt.Nanosecond(), nt.Unix())
		}
	})
	b.Run("fastime-from-time-2", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		nt := time.Now()
		ns := nt.Nanosecond()
		milis := nt.Unix()
		for n := 0; n < b.N; n++ {
			_ = FromTime(ns, milis)
		}
	})
	b.Run("fastime-from-time-3", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			nt := time.Now()
			ns := nt.Nanosecond()
			milis := nt.Unix()
			_ = FromTime(ns, milis)
		}
	})
}

func BenchmarkStandardTime(b *testing.B) {

	b.Run("standard-now", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = time.Now()
		}
	})
	b.Run("standard-now-unix", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = time.Now().Unix()
		}
	})
	b.Run("standard-now-nanos", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = time.Now().Nanosecond()
		}
	})
}