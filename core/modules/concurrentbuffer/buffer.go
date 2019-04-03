// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package concurrentbuffer

import (
	"bytes"
	"io"
	"sync"
)

type ConcurrentBuffer struct {
	b bytes.Buffer
	// mutual-exclusion lock
	m *sync.Mutex
}

func (b *ConcurrentBuffer) Read(p []byte) (n int, err error) {
	b.m.Lock()
	n, err = b.b.Read(p)
	b.m.Unlock()
	return
}
func (b *ConcurrentBuffer) Write(p []byte) (n int, err error) {
	b.m.Lock()
	n, err = b.b.Write(p)
	b.m.Unlock()
	return
}
func (b *ConcurrentBuffer) String() string {
	b.m.Lock()
	raw := b.b.String()
	b.m.Unlock()
	return raw
}
func (b *ConcurrentBuffer) Bytes() []byte {
	b.m.Lock()
	raw := b.b.Bytes()
	b.m.Unlock()
	return raw
}
func (b *ConcurrentBuffer) Cap() int {
	b.m.Lock()
	raw := b.b.Cap()
	b.m.Unlock()
	return raw
}
func (b *ConcurrentBuffer) Grow(n int) {
	b.m.Lock()
	b.b.Grow(n)
	b.m.Unlock()
}
func (b *ConcurrentBuffer) Len() int {
	b.m.Lock()
	raw := b.b.Len()
	b.m.Unlock()
	return raw
}
func (b *ConcurrentBuffer) Next(n int) []byte {
	b.m.Lock()
	raw := b.b.Next(n)
	b.m.Unlock()
	return raw
}
func (b *ConcurrentBuffer) ReadByte() (c byte, err error) {
	b.m.Lock()
	c, err = b.b.ReadByte()
	b.m.Unlock()
	return
}
func (b *ConcurrentBuffer) ReadBytes(delim byte) (line []byte, err error) {
	b.m.Lock()
	line, err = b.b.ReadBytes(delim)
	b.m.Unlock()
	return
}
func (b *ConcurrentBuffer) ReadFrom(r io.Reader) (n int64, err error) {
	b.m.Lock()
	n, err = b.b.ReadFrom(r)
	b.m.Unlock()
	return
}
func (b *ConcurrentBuffer) ReadRune() (r rune, size int, err error) {
	b.m.Lock()
	r, size, err = b.b.ReadRune()
	b.m.Unlock()
	return
}
func (b *ConcurrentBuffer) ReadString(delim byte) (line string, err error) {
	b.m.Lock()
	line, err = b.b.ReadString(delim)
	b.m.Unlock()
	return
}
func (b *ConcurrentBuffer) Reset() {
	b.m.Lock()
	b.b.Reset()
	b.m.Unlock()
}
func (b *ConcurrentBuffer) Truncate(n int) {
	b.m.Lock()
	b.b.Truncate(n)
	b.m.Unlock()
}
func (b *ConcurrentBuffer) UnreadByte() error {
	b.m.Lock()
	raw := b.b.UnreadByte()
	b.m.Unlock()
	return raw
}
func (b *ConcurrentBuffer) UnreadRune() error {
	b.m.Lock()
	raw := b.b.UnreadRune()
	b.m.Unlock()
	return raw
}
func (b *ConcurrentBuffer) WriteByte(c byte) error {
	b.m.Lock()
	raw := b.b.WriteByte(c)
	b.m.Unlock()
	return raw
}
func (b *ConcurrentBuffer) WriteRune(r rune) (n int, err error) {
	b.m.Lock()
	n, err = b.b.WriteRune(r)
	b.m.Unlock()
	return
}
func (b *ConcurrentBuffer) WriteString(s string) (n int, err error) {
	b.m.Lock()
	n, err = b.b.WriteString(s)
	b.m.Unlock()
	return
}
func (b *ConcurrentBuffer) WriteTo(w io.Writer) (n int64, err error) {
	b.m.Lock()
	n, err = b.b.WriteTo(w)
	b.m.Unlock()
	return
}

// constructor like function for concurrent buffer
func NewConcurrentBuffer() ConcurrentBuffer {
	cb := ConcurrentBuffer{}
	cb.m = new(sync.Mutex)
	return cb
}

// constructor like function for concurrent buffer
func NewConcurrentBufferPtr() *ConcurrentBuffer {
	cb := new(ConcurrentBuffer)
	cb.m = new(sync.Mutex)
	return cb
}