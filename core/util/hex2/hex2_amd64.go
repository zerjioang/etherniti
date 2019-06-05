//+build !noasm
//+build !appengine

package hex2

import (
"unsafe"
)

//go:noescape
func _Hex2(vec1, vec2, vec3, result unsafe.Pointer)

func Hex2(f1, f2, f3 *[8]float32) [8]float32 {

	_f4 := [8]float32{}

	_Hex2(unsafe.Pointer(f1), unsafe.Pointer(f2), unsafe.Pointer(f3), unsafe.Pointer(&_f4))

	return _f4
}
