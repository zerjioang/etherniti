// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package util

import (
	json2 "encoding/json"
	"testing"
)

func TestGetJsonBytes(t *testing.T) {
	t.Run("get-bytes-nil", func(t *testing.T) {
		GetJsonBytes(nil)
	})
	t.Run("ToLowerAscii", func(t *testing.T) {
		val := "Hello World, This is AWESOME"
		converted := ToLowerAscii(val)
		t.Log(val)
		t.Log(converted)
		if converted != "hello world, this is awesome" {
			t.Error("failed to lowercase")
		}
	})
	t.Run("ToLowerAscii-ua", func(t *testing.T) {
		val := "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:61.0) Gecko/20100101 Firefox/61.0"
		converted := ToLowerAscii(val)
		t.Log(val)
		t.Log(converted)
		if converted != "mozilla/5.0 (x11; ubuntu; linux x86_64; rv:61.0) gecko/20100101 firefox/61.0" {
			t.Error("failed to lowercase")
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
	b.Run("fast-marshal-example", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)

		type testStruct struct {
			Id      int
			Message string
		}
		test := testStruct{Id: 23554675, Message: "this is a test struct"}
		for i := 0; i < b.N; i++ {
			_, _ = FastMarshal(test)
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
			_, _ = json2.Marshal(test)
		}
	})
}
