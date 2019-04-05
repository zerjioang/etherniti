// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package stringbank

import (
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStringbank(t *testing.T) {
	sb := Stringbank{}

	s1 := sb.Save("hello")
	s2 := sb.Save("goodbye")
	s3 := sb.Save("cheese")

	assert.Equal(t, "hello", sb.Get(s1))
	assert.Equal(t, "goodbye", sb.Get(s2))
	assert.Equal(t, "cheese", sb.Get(s3))
}

func TestStringbankSize(t *testing.T) {
	sb := Stringbank{}
	assert.Zero(t, sb.Size())
	sb.Save("hello")
	assert.Equal(t, stringbankSize, sb.Size())
}

func TestPackageBank(t *testing.T) {
	s1 := Save("hello")
	s2 := Save("goodbye")
	s3 := Save("cheese")

	assert.Equal(t, "hello", s1.String())
	assert.Equal(t, "goodbye", s2.String())
	assert.Equal(t, "cheese", s3.String())
}

func TestLengths(t *testing.T) {
	tests := []struct {
		len int
	}{
		{1},
		{127},
		{128},
		{254},
		{255},
		{256},
		{0xFFFFFFFFFF},
	}

	for _, test := range tests {
		t.Run(strconv.Itoa(test.len), func(t *testing.T) {
			buf := make([]byte, 10)

			l := writeLength(test.len, buf)
			assert.Equal(t, l, spaceForLength(test.len))
			len, lenlen := readLength(buf)
			assert.Equal(t, l, lenlen)
			assert.Equal(t, test.len, len)
		})
	}
}

func TestGC(t *testing.T) {
	sb := Stringbank{}
	for i := 0; i < 10000000; i++ {
		sb.Save(strconv.Itoa(i))
	}
	runtime.GC()

	start := time.Now()
	runtime.GC()
	assert.True(t, time.Since(start) < 1000*time.Microsecond)
	runtime.KeepAlive(sb)
}

func ExampleSave() {
	i := Save("hello")
	fmt.Println(i)
	// Output: hello
}

func ExampleStringbank() {
	sb := Stringbank{}
	i := sb.Save("goodbye")
	fmt.Println(sb.Get(i))
	// Output: goodbye
}

func TestIndex_String(t *testing.T) {
	tests := []struct {
		name string
		i    Index
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.String(); got != tt.want {
				t.Errorf("Index.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSave(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name string
		args args
		want Index
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Save(tt.args.val); got != tt.want {
				t.Errorf("Save() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewStringbank(t *testing.T) {
	tests := []struct {
		name string
		want Stringbank
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStringbank(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStringbank() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringbank_Size(t *testing.T) {
	type fields struct {
		current     []byte
		allocations [][]byte
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stringbank{
				current:     tt.fields.current,
				allocations: tt.fields.allocations,
			}
			if got := s.Size(); got != tt.want {
				t.Errorf("Stringbank.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringbank_Get(t *testing.T) {
	type fields struct {
		current     []byte
		allocations [][]byte
	}
	type args struct {
		index int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stringbank{
				current:     tt.fields.current,
				allocations: tt.fields.allocations,
			}
			if got := s.Get(tt.args.index); got != tt.want {
				t.Errorf("Stringbank.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringbank_Save(t *testing.T) {
	type fields struct {
		current     []byte
		allocations [][]byte
	}
	type args struct {
		tocopy string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stringbank{
				current:     tt.fields.current,
				allocations: tt.fields.allocations,
			}
			if got := s.Save(tt.args.tocopy); got != tt.want {
				t.Errorf("Stringbank.Save() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringbank_reserve(t *testing.T) {
	type fields struct {
		current     []byte
		allocations [][]byte
	}
	type args struct {
		l int
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantIndex int
		wantData  []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stringbank{
				current:     tt.fields.current,
				allocations: tt.fields.allocations,
			}
			gotIndex, gotData := s.reserve(tt.args.l)
			if gotIndex != tt.wantIndex {
				t.Errorf("Stringbank.reserve() gotIndex = %v, want %v", gotIndex, tt.wantIndex)
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("Stringbank.reserve() gotData = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}

func Test_spaceForLength(t *testing.T) {
	type args struct {
		len int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spaceForLength(tt.args.len); got != tt.want {
				t.Errorf("spaceForLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_writeLength(t *testing.T) {
	type args struct {
		len int
		buf []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := writeLength(tt.args.len, tt.args.buf); got != tt.want {
				t.Errorf("writeLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readLength(t *testing.T) {
	type args struct {
		buf []byte
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := readLength(tt.args.buf)
			if got != tt.want {
				t.Errorf("readLength() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("readLength() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestExampleSave(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ExampleSave()
		})
	}
}

func TestExampleStringbank(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ExampleStringbank()
		})
	}
}
