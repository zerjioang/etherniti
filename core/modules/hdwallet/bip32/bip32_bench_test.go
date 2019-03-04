// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package bip32

import (
	"log"
	"testing"
)

func BenchmarkBip32(b *testing.B) {
	b.Run("new-seed", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		for n := 0; n < b.N; n++ {
			// Generate a seed to determine all keys from.
			// This should be persisted, backed up, and secured
			_, _ = NewSeed()
		}
	})
	b.Run("new-master", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		for n := 0; n < b.N; n++ {
			// Generate a seed to determine all keys from.
			// This should be persisted, backed up, and secured
			seed, err := NewSeed()
			if err != nil {
				log.Fatalln("Error generating seed:", err)
			}

			// Create master private key from seed
			_, _ = NewMasterKey(seed)
		}
	})
	b.Run("new-child-key-600", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		for n := 0; n < b.N; n++ {
			// Generate a seed to determine all keys from.
			// This should be persisted, backed up, and secured
			seed, err := NewSeed()
			if err != nil {
				log.Fatalln("Error generating seed:", err)
			}

			_, _ = NewMasterKey(seed)
			master, _ := NewMasterKey(seed)
			_, _ = master.NewChildKey(600)
		}
	})
	b.Run("new-child-key-60000", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		for n := 0; n < b.N; n++ {
			// Generate a seed to determine all keys from.
			// This should be persisted, backed up, and secured
			seed, err := NewSeed()
			if err != nil {
				log.Fatalln("Error generating seed:", err)
			}

			_, _ = NewMasterKey(seed)
			master, _ := NewMasterKey(seed)
			_, _ = master.NewChildKey(60000)
		}
	})
	b.Run("example", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		for n := 0; n < b.N; n++ {
			// Generate a seed to determine all keys from.
			// This should be persisted, backed up, and secured
			seed, err := NewSeed()
			if err != nil {
				log.Fatalln("Error generating seed:", err)
			}

			// Create master private key from seed
			computerVoiceMasterKey, _ := NewMasterKey(seed)

			// Map departments to keys
			// There is a very small chance a given child index is invalid
			// If so your real program should handle this by skipping the index
			departmentKeys := map[string]*Key{}
			departmentKeys["a"], _ = computerVoiceMasterKey.NewChildKey(0)
			departmentKeys["b"], _ = computerVoiceMasterKey.NewChildKey(1)
			departmentKeys["c"], _ = computerVoiceMasterKey.NewChildKey(2)
			departmentKeys["d"], _ = computerVoiceMasterKey.NewChildKey(3)
		}
	})
}
