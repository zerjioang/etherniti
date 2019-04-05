// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package concurrentbuffer

import (
	"bytes"
	"io"
	"reflect"
	"sync"
	"testing"
)

func TestConcurrentBuffer_Read(t *testing.T) {
	type fields struct {
		b bytes.Buffer
		m *sync.Mutex
	}
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantN   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &ConcurrentBuffer{
				b: tt.fields.b,
				m: tt.fields.m,
			}
			gotN, err := b.Read(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConcurrentBuffer.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("ConcurrentBuffer.Read() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func TestConcurrentBuffer_Write(t *testing.T) {
	type fields struct {
		b bytes.Buffer
		m *sync.Mutex
	}
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantN   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &ConcurrentBuffer{
				b: tt.fields.b,
				m: tt.fields.m,
			}
			gotN, err := b.Write(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConcurrentBuffer.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("ConcurrentBuffer.Write() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func TestConcurrentBuffer_String(t *testing.T) {
	type fields struct {
		b bytes.Buffer
		m *sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &ConcurrentBuffer{
				b: tt.fields.b,
				m: tt.fields.m,
			}
			if got := b.String(); got != tt.want {
				t.Errorf("ConcurrentBuffer.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConcurrentBuffer_Bytes(t *testing.T) {
	type fields struct {
		b bytes.Buffer
		m *sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &ConcurrentBuffer{
				b: tt.fields.b,
				m: tt.fields.m,
			}
			if got := b.Bytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConcurrentBuffer.Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConcurrentBuffer_Cap(t *testing.T) {
	type fields struct {
		b bytes.Buffer
		m *sync.Mutex
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
			b := &ConcurrentBuffer{
				b: tt.fields.b,
				m: tt.fields.m,
			}
			if got := b.Cap(); got != tt.want {
				t.Errorf("ConcurrentBuffer.Cap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConcurrentBuffer_Grow(t *testing.T) {
	type fields struct {
		b bytes.Buffer
		m *sync.Mutex
	}
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &ConcurrentBuffer{
				b: tt.fields.b,
				m: tt.fields.m,
			}
			b.Grow(tt.args.n)
		})
	}
}

func TestConcurrentBuffer_Len(t *testing.T) {
	type fields struct {
		b bytes.Buffer
		m *sync.Mutex
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
			b := &ConcurrentBuffer{
				b: tt.fields.b,
				m: tt.fields.m,
			}
			if got := b.Len(); got != tt.want {
				t.Errorf("ConcurrentBuffer.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConcurrentBuffer_Next(t *testing.T) {
	type fields struct {
		b bytes.Buffer
		m *sync.Mutex
	}
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &ConcurrentBuffer{
				b: tt.fields.b,
				m: tt.fields.m,
			}
			if got := b.Next(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConcurrentBuffer.Next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConcurrentBuffer_ReadByte(t *testing.T) {
	type fields struct {
		b bytes.Buffer
		m *sync.Mutex
	}
	tests := []struct {
		name    string
		fields  fields
		wantC   byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &ConcurrentBuffer{
				b: tt.fields.b,
				m: tt.fields.m,
			}
			gotC, err := b.ReadByte()
			if (err != nil) != tt.wantErr {
				t.Errorf("ConcurrentBuffer.ReadByte() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotC != tt.wantC {
				t.Errorf("ConcurrentBuffer.ReadByte() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}

func TestConcurrentBuffer_ReadBytes(t *testing.T) {
	type fields struct {
		b bytes.Buffer
		m *sync.Mutex
	}
	type args struct {
		delim byte
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantLine []byte
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &ConcurrentBuffer{
				b: tt.fields.b,
				m: tt.fields.m,
			}
			gotLine, err := b.ReadBytes(tt.args.delim)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConcurrentBuffer.ReadBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotLine, tt.wantLine) {
				t.Errorf("ConcurrentBuffer.ReadBytes() = %v, want %v", gotLine, tt.wantLine)
			}
		})
	}
}

func TestConcurrentBuffer_ReadFrom(t *testing.T) {
	type fields struct {
		b bytes.Buffer
		m *sync.Mutex
	}
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantN   int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &ConcurrentBuffer{
				b: tt.fields.b,
				m: tt.fields.m,
			}
			gotN, err := b.ReadFrom(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConcurrentBuffer.ReadFrom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("ConcurrentBuffer.ReadFrom() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func TestConcurrentBuffer_ReadRune(t *testing.T) {
	type fields struct {
		b bytes.Buffer
		m *sync.Mutex
	}
	tests := []struct {
		name     string
		fields   fields
		wantR    rune
		wantSize int
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &ConcurrentBuffer{
				b: tt.fields.b,
				m: tt.fields.m,
			}
			gotR, gotSize, err := b.ReadRune()
			if (err != nil) != tt.wantErr {
				t.Errorf("ConcurrentBuffer.ReadRune() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotR != tt.wantR {
				t.Errorf("ConcurrentBuffer.ReadRune() gotR = %v, want %v", gotR, tt.wantR)
			}
			if gotSize != tt.wantSize {
				t.Errorf("ConcurrentBuffer.ReadRune() gotSize = %v, want %v", gotSize, tt.wantSize)
			}
		})
	}
}

func TestConcurrentBuffer_ReadString(t *testing.T) {
	type fields struct {
		b bytes.Buffer
		m *sync.Mutex
	}
	type args struct {
		delim byte
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantLine string
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &ConcurrentBuffer{
				b: tt.fields.b,
				m: tt.fields.m,
			}
			gotLine, err := b.ReadString(tt.args.delim)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConcurrentBuffer.ReadString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotLine != tt.wantLine {
				t.Errorf("ConcurrentBuffer.ReadString() = %v, want %v", gotLine, tt.wantLine)
			}
		})
	}
}

func TestConcurrentBuffer_Reset(t *testing.T) {
	type fields struct {
		b bytes.Buffer
		m *sync.Mutex
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &ConcurrentBuffer{
				b: tt.fields.b,
				m: tt.fields.m,
			}
			b.Reset()
		})
	}
}

func TestConcurrentBuffer_Truncate(t *testing.T) {
	type fields struct {
		b bytes.Buffer
		m *sync.Mutex
	}
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &ConcurrentBuffer{
				b: tt.fields.b,
				m: tt.fields.m,
			}
			b.Truncate(tt.args.n)
		})
	}
}

func TestConcurrentBuffer_UnreadByte(t *testing.T) {
	type fields struct {
		b bytes.Buffer
		m *sync.Mutex
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &ConcurrentBuffer{
				b: tt.fields.b,
				m: tt.fields.m,
			}
			if err := b.UnreadByte(); (err != nil) != tt.wantErr {
				t.Errorf("ConcurrentBuffer.UnreadByte() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConcurrentBuffer_UnreadRune(t *testing.T) {
	type fields struct {
		b bytes.Buffer
		m *sync.Mutex
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &ConcurrentBuffer{
				b: tt.fields.b,
				m: tt.fields.m,
			}
			if err := b.UnreadRune(); (err != nil) != tt.wantErr {
				t.Errorf("ConcurrentBuffer.UnreadRune() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConcurrentBuffer_WriteByte(t *testing.T) {
	type fields struct {
		b bytes.Buffer
		m *sync.Mutex
	}
	type args struct {
		c byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &ConcurrentBuffer{
				b: tt.fields.b,
				m: tt.fields.m,
			}
			if err := b.WriteByte(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("ConcurrentBuffer.WriteByte() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConcurrentBuffer_WriteRune(t *testing.T) {
	type fields struct {
		b bytes.Buffer
		m *sync.Mutex
	}
	type args struct {
		r rune
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantN   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &ConcurrentBuffer{
				b: tt.fields.b,
				m: tt.fields.m,
			}
			gotN, err := b.WriteRune(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConcurrentBuffer.WriteRune() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("ConcurrentBuffer.WriteRune() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func TestConcurrentBuffer_WriteString(t *testing.T) {
	type fields struct {
		b bytes.Buffer
		m *sync.Mutex
	}
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantN   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &ConcurrentBuffer{
				b: tt.fields.b,
				m: tt.fields.m,
			}
			gotN, err := b.WriteString(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConcurrentBuffer.WriteString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("ConcurrentBuffer.WriteString() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func TestConcurrentBuffer_WriteTo(t *testing.T) {
	type fields struct {
		b bytes.Buffer
		m *sync.Mutex
	}
	tests := []struct {
		name    string
		fields  fields
		wantN   int64
		wantW   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &ConcurrentBuffer{
				b: tt.fields.b,
				m: tt.fields.m,
			}
			w := &bytes.Buffer{}
			gotN, err := b.WriteTo(w)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConcurrentBuffer.WriteTo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("ConcurrentBuffer.WriteTo() = %v, want %v", gotN, tt.wantN)
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("ConcurrentBuffer.WriteTo() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestNewConcurrentBuffer(t *testing.T) {
	tests := []struct {
		name string
		want ConcurrentBuffer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConcurrentBuffer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConcurrentBuffer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewConcurrentBufferPtr(t *testing.T) {
	tests := []struct {
		name string
		want *ConcurrentBuffer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConcurrentBufferPtr(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConcurrentBufferPtr() = %v, want %v", got, tt.want)
			}
		})
	}
}
