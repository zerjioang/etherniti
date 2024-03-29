// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package profile

import (
	"testing"
)

func BenchmarkConnectionProfile(b *testing.B) {
	b.Run("create-profile-empty", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = NewConnectionProfile()
		}
	})
	b.Run("create-profile", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = NewDefaultConnectionProfile()
		}
	})
	b.Run("valid-false", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		profile := NewConnectionProfile()
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = profile.Valid()
		}
	})
	b.Run("valid-true", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)

		profile := NewConnectionProfile()
		profile.Id = "test-id"
		profile.RpcEndpoint = "node-test-address"
		profile.Address = "test-account"
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = profile.Valid()
		}
	})
	b.Run("get-secret", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		profile := NewConnectionProfile()
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			profile.Secret()
		}
	})
	b.Run("create-token", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		profile := NewDefaultConnectionProfile()
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_, _ = CreateConnectionProfileToken(profile)
		}
	})
	b.Run("parse-token", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_, _ = ParseConnectionProfileToken(testToken)
		}
	})
}
