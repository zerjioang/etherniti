package profile

import (
	"testing"
	"time"

	"github.com/zerjioang/etherniti/core/util"
)

func BenchmarkConnectionProfile(b *testing.B) {
	b.Run("instantiate", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		for n := 0; n < b.N; n++ {
			_ = NewConnectionProfile()
		}
	})
	b.Run("valid-false", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		profile := NewConnectionProfile()
		for n := 0; n < b.N; n++ {
			_ = profile.Valid()
		}
	})
	b.Run("valid-true", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		// run the Fib function b.N times
		profile := NewConnectionProfile()
		profile.ConnectionId = "test-id"
		profile.NodeAddress = "node-test-address"
		profile.Account = "test-account"
		for n := 0; n < b.N; n++ {
			_ = profile.Valid()
		}
	})
	b.Run("get-secret", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		profile := NewConnectionProfile()
		for n := 0; n < b.N; n++ {
			profile.Secret()
		}
	})
	b.Run("create-token", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		now := time.Now()
		profile := ConnectionProfile{
			ConnectionId: util.GenerateUUID(),
			NodeAddress:  "http://127.0.0.1:8454",
			Mode:         "http",
			Port:         8454,
			Account:      "0x0",
			//standard claims
			Id:        util.GenerateUUID(),
			Issuer:    "etherniti",
			ExpiresAt: now.Add(10 * time.Minute).Unix(),
			NotBefore: now.Unix(),
			IssuedAt:  now.Unix(),
		}
		for n := 0; n < b.N; n++ {
			_, _ = CreateConnectionProfileToken(profile)
		}
	})
	b.Run("parse-token", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		for n := 0; n < b.N; n++ {
			_, _ = ParseConnectionProfileToken(testToken)
		}
	})
}
