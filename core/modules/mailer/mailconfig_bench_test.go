// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package mailer

import (
	"testing"
)

func BenchmarkGetMailServerConfigInstance(b *testing.B) {
	b.SetBytes(1)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetMailServerConfigInstance()
	}
}

func BenchmarkGetMailServerConfigInstanceThreadSafe(b *testing.B) {
	b.SetBytes(1)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetMailServerConfigInstanceThreadSafe()
	}
}

func BenchmarkGetMailServerConfigInstanceInit(b *testing.B) {
	b.SetBytes(1)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetMailServerConfigInstanceInit()
	}
}
