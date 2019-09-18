package httpawn

import (
	"reflect"
	"unsafe"
)

func resolveHttpPath(raw []byte, start uint8) string {
	return "/"
}
func resolveHttpMethod(raw []byte) (httpMethod, uint8) {
	// GET, HEAD, POST, PUT, DELETE, CONNECT, OPTIONS, TRACE, PATCH
	r1 := raw[1]
	r0 := raw[0]
	if r0 == 'P' {
		// http method can be:
		// POST, PUT, PATCH
		switch r1 {
		case 'O': //POST
			return POST, 4
		case 'U': //PUT
			return PUT, 3
		case 'A': //PATCH
			return PATCH, 5
		}
	} else {
		// we use only first char to detect http method
		// we also asume http client is honest and follows http specification
		switch r0 {
		case 'G': //GET
			return GET, 3
		case 'H': //HEAD
			return HEAD, 4
		case 'D': //DELETE
			return DELETE, 6
		case 'C': //CONNECT
			return CONNECT, 7
		case 'O': //OPTIONS
			return OPTIONS, 7
		case 'T': //TRACE
			return TRACE, 5
		}
	}
	return UNKNOWN, 0
}

// A hack until issue golang/go#2632 is fixed.
// See: https://github.com/golang/go/issues/2632
func BytesToString(b *[]byte) string {
	return *(*string)(unsafe.Pointer(b))
}

func StringToBytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}
