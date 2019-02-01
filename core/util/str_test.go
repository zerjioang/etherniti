package util

import (
	json2 "encoding/json"
	"testing"
)

func TestGetJsonBytes(t *testing.T) {
	t.Run("get-bytes-nil", func(t *testing.T) {
		GetJsonBytes(nil)
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
			Id int
			Message string
		}
		test := testStruct{Id:23554675, Message:"this is a test struct"}
		for i := 0; i < b.N; i++ {
			_,_ = FastMarshal(test)
		}
	})
	b.Run("std-marshal-example", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)

		type testStruct struct {
			Id int
			Message string
		}
		test := testStruct{Id:23554675, Message:"this is a test struct"}
		for i := 0; i < b.N; i++ {
			_,_ = StdMarshal(test)
		}
	})
	b.Run("std-json-go", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)

		type testStruct struct {
			Id int
			Message string
		}
		test := testStruct{Id:23554675, Message:"this is a test struct"}
		for i := 0; i < b.N; i++ {
			_,_ = json2.Marshal(test)
		}
	})
}
