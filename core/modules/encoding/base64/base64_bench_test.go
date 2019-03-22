// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package base64

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func BenchmarkBase64(b *testing.B) {
	b.Run("iterations", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		for n := 0; n < b.N; n++ {
			msg := "Hello, world"
			encoded := base64.StdEncoding.EncodeToString([]byte(msg))
			_, err := base64.StdEncoding.DecodeString(encoded)
			if err != nil {
				fmt.Println("decode error:", err)
				return
			}
		}
	})
}
