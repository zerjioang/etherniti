package string

import (
	"bytes"
	"reflect"
	"unsafe"
)

const (
	level = 32
)

type String struct {
	// Data here is analogous to the C string
	Data  unsafe.Pointer
	Len   int
	start uintptr
}

func New() String {
	return String{}
}

func NewWith(data []byte) String {
	s := String{}
	s.Data = unsafe.Pointer(&data)
	s.start = uintptr(unsafe.Pointer(&data[0]))
	s.Len = len(data)
	return s
}

func (c String) CharAt(index int) byte {
	return *(*byte)((unsafe.Pointer)(c.start + uintptr(index)))
}

func (c *String) Bytes() []byte {
	return *(*[]byte)((unsafe.Pointer)(c.Data))
}

// LastIndex returns the index of the last instance of substr in s, or -1 if substr is not present in s.
func (c *String) LastIndex(item []byte) int {
	inputSize := len(item)
	switch {
	case inputSize == 0:
		return c.Len
	case inputSize == 1:
		// todo
		return c.LastIndexOfByte(item[0])
	case inputSize == c.Len:
		return bytes.Compare(c.Bytes(), item)
	case inputSize > c.Len:
		return -1
	}
	return -1
}

// Length returns the length of the string
func (c String) Length() int {
	return c.Len
}

// IsEmpty returns whether the string is empty or not
func (c String) IsEmpty() bool {
	return c.Len == 0
}

// String returns native implementation of string value
func (c *String) String() string {
	header := (*reflect.SliceHeader)(c.Data)
	strHeader := &reflect.StringHeader{
		Data: header.Data,
		Len:  header.Len,
	}
	return *(*string)(unsafe.Pointer(strHeader))
}

// LowerCase returns a lowercase version of the string
func (c String) LowerCase() {
	for i := 0; i < c.Len; i++ {
		// get char at current index
		ptri := uintptr(i)
		char := *(*byte)((unsafe.Pointer)(c.start + ptri))
		if char >= 'A' && char <= 'Z' {
			*(*byte)((unsafe.Pointer)(c.start + ptri)) = char | level
		}
	}
}

// UpperCase returns an uppercase version of the string
func (c String) UpperCase() {
	for i := 0; i < c.Len; i++ {
		// get char at current index
		ptri := uintptr(i)
		char := *(*byte)((unsafe.Pointer)(c.start + ptri))
		if char >= 'a' && char <= 'z' {
			*(*byte)((unsafe.Pointer)(c.start + ptri)) = char &^ level
		}
	}
}

func (c *String) Reverse() {
	for i := 0; i < c.Len/2; i++ {
		a := *(*byte)((unsafe.Pointer)(c.start + uintptr(i)))
		*(*byte)((unsafe.Pointer)(c.start + uintptr(i))) = *(*byte)((unsafe.Pointer)(c.start + uintptr(c.Len-1-i)))
		*(*byte)((unsafe.Pointer)(c.start + uintptr(c.Len-1-i))) = a
	}
}

func (c *String) CountByte(b byte) int {
	counter := 0
	for i := 0; i < c.Len; i++ {
		c := *(*byte)((unsafe.Pointer)(c.start + uintptr(i)))
		if c^b == 0x0 {
			counter++
		}
	}
	return counter
}

func (c String) TitleCase() {
	//check if first byte is ascii, put uppercase if it is
	char := *(*byte)((unsafe.Pointer)(c.start + uintptr(0)))
	if char >= 'a' && char <= 'z' {
		*(*byte)((unsafe.Pointer)(c.start + uintptr(0))) = char &^ level
	}
	for i := 1; i < c.Len; i++ {
		// get previous char
		previous := *(*byte)((unsafe.Pointer)(c.start + uintptr(i-1)))
		// get current char
		ptri := uintptr(i)
		char := *(*byte)((unsafe.Pointer)(c.start + ptri))
		// put it lowercase, just in case
		*(*byte)((unsafe.Pointer)(c.start + ptri)) = char | level

		// check if previous char is separator
		if previous == ' ' || previous == '-' || previous == '.' {
			*(*byte)((unsafe.Pointer)(c.start + ptri)) = char &^ level
		}
	}
}

// String returns native implementation of string value
func (c String) Compare(b String) int {
	return 0
}

// return the index of first matching byte
func (c *String) LastIndexOfByte(b byte) int {
	var found byte = 0xff
	var i int
	for i = 0; i < c.Len && found != 0; i++ {
		found = *(*byte)((unsafe.Pointer)(c.start + uintptr(i))) ^ b
	}
	return i - 1
}

func (c String) Contains(subStr []byte) bool {
	var i int
	if subStr == nil || len(subStr) == 0 {
		return false
	}
	subStrStart := uintptr(unsafe.Pointer(&subStr[0]))
	lastSubStrIdx := 0
	currentSubstrMatch := *(*byte)((unsafe.Pointer)(subStrStart + uintptr(0)))
	s := len(subStr)
	for i = 0; i < c.Len && lastSubStrIdx < s; i++ {
		current := *(*byte)((unsafe.Pointer)(c.start + uintptr(i)))
		if current == currentSubstrMatch {
			//we matched first char, lets check if following char matches too
			lastSubStrIdx++
			if lastSubStrIdx < s {
				currentSubstrMatch = *(*byte)((unsafe.Pointer)(subStrStart + uintptr(lastSubStrIdx)))
			}
		} else {
			lastSubStrIdx = 0
			currentSubstrMatch = *(*byte)((unsafe.Pointer)(subStrStart + uintptr(lastSubStrIdx)))
		}
	}
	return lastSubStrIdx == s
}
