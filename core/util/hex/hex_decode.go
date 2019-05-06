// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package hex

import (
	"errors"
	"fmt"
	"reflect"
	"unsafe"
)

// ErrLength reports an attempt to decode an odd-length input
// using Decode or DecodeString.
// The stream-based Decoder returns io.ErrUnexpectedEOF instead of ErrLength.
var ErrLength = errors.New("encoding/hex: odd length hex string")

// InvalidByteError values describe errors resulting from an invalid byte in a hex string.
type InvalidByteError byte

func init() {
}

// DecodeString returns the bytes represented by the hexadecimal string s.
//
// DecodeString expects that src contains only hexadecimal
// characters and that src has even length.
// If the input is malformed, DecodeString returns
// the bytes decoded before the error.
func UnsafeDecodeString(s string) ([]byte, error) {
	// convert string to bytes using unsfe pointer
	//return *(*[]byte)(unsafe.Pointer(&data))
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	src := *(*[]byte)(unsafe.Pointer(&bh))
	// We can use the source slice itself as the destination
	// because the decode loop increments by one and then the 'seen' byte is not used anymore.
	n, err := decode(src, src)
	return src[:n], err
}

// todo: WIP
func decode(dst, src []byte) (int, error) {

	start := uintptr(unsafe.Pointer(&dst[0]))
	step := unsafe.Sizeof(dst[0])

	var i int
	for i = 0; i < len(src)/2; i++ {
		c := src[i*2]
		a, b, ok1, ok2 := fromHexCharDual(src[i*2], src[i*2+1])
		if !ok1 || !ok2 {
			return i, InvalidByteError(c)
		}
		*(*byte)((unsafe.Pointer)(start + step*uintptr(i))) = (a << 4) | b
	}
	if len(src)%2 == 1 {
		// Check for invalid char before reporting bad length,
		// since the invalid char (if present) is an earlier problem.
		c := src[i*2]
		if _, ok := fromHexChar(c); !ok {
			return i, InvalidByteError(c)
		}
		return i, ErrLength
	}
	return i, nil
}

// fromHexChar converts a hex character into its value and a success flag.
func fromHexChar(c byte) (byte, bool) {
	switch {
	case '0' <= c && c <= '9':
		return c - '0', true
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10, true
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10, true
	}

	return 0, false
}
// fromHexChar converts a hex character into its value and a success flag.
func fromHexCharDual(a byte, b byte) (a1 byte, b1 byte, ab bool, bb bool) {
	switch {
	case '0' <= a && a <= '9':
		a1 = a - '0'
		ab = true
	case 'a' <= a && a <= 'f':
		a1 = a - 'a' + 10
		ab = true
	case 'A' <= a && a <= 'F':
		a1 = a - 'A' + 10
		ab = true
	}

	switch {
	case '0' <= a && a <= '9':
		b1 = b - '0'
		bb = true
	case 'a' <= a && a <= 'f':
		b1 = b - 'a' + 10
		bb = true
	case 'A' <= a && a <= 'F':
		b1 = b - 'A' + 10
		bb = true
	}

	return
}
func (e InvalidByteError) Error() string {
	return fmt.Sprintf("encoding/hex: invalid byte: %#U", rune(e))
}
