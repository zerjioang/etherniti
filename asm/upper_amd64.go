//+build !noasm
//+build !appengine

package hex

import (
	"unsafe"
)

//go:noescape
func _add(a, b int) (c int)

func Add(a, b int) int {
	return _add(a, b)
}

//go:noescape
func _toUpper(src unsafe.Pointer) (result unsafe.Pointer)

func ToUpper(src []byte) []byte {
	ptr := unsafe.Pointer(&src)
	result := _toUpper(ptr)
	return *(*[]byte)(result)
	// return src
}

//go:noescape
func _toLower(src unsafe.Pointer) (result unsafe.Pointer)

func ToLower(src []byte) []byte {
	ptr := unsafe.Pointer(&src)
	result := _toLower(ptr)
	return *(*[]byte)(result)
}
