// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package str

import (
	"encoding/json"
	"strings"
	"testing"
)

func BenchmarkStringUtils(b *testing.B) {

	b.Run("to-lower-std", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		val := "Hello World, This is AWESOME"
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = strings.ToLower(val)
		}
	})
	b.Run("ToLowerAscii", func(b *testing.B) {
		b.Run("empty", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			val := ""
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = ToLowerAscii(val)
			}
		})
		b.Run("with-content", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			val := "Hello World, This is AWESOME"
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = ToLowerAscii(val)
			}
		})
	})
	b.Run("ToLowerAscii-bytes", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		val := "Hello World, This is AWESOME"
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = ToLowerAscii(val)
		}
	})

	b.Run("len-std", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		val := "Hello World, This is AWESOME"
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = len(val)
		}
	})
	b.Run("len-custom", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		val := "Hello World, This is AWESOME"
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = strLen(val)
		}
	})
}

func BenchmarkGetJsonBytes(b *testing.B) {
	b.Run("get-bytes-nil", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		for i := 0; i < b.N; i++ {
			GetJsonBytes(nil)
		}
	})
	b.Run("std-marshal-example", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)

		type testStruct struct {
			Id      int
			Message string
		}
		test := testStruct{Id: 23554675, Message: "this is a test struct"}
		for i := 0; i < b.N; i++ {
			_, _ = StdMarshal(test)
		}
	})
	b.Run("std-json-go", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)

		type testStruct struct {
			Id      int
			Message string
		}
		test := testStruct{Id: 23554675, Message: "this is a test struct"}
		for i := 0; i < b.N; i++ {
			_, _ = json.Marshal(test)
		}
	})
}
