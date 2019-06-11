//+build !noasm
//+build !appengine

package string

//go:noescape
func _isNumericArray(buf, len, res unsafe.Pointer)

//go:noescape
func _lowerCase(buf, len, res unsafe.Pointer)
//go:noescape
func _isDigit(b byte) (result byte)

func IsNumericArray(src []byte) bool {
	var result byte
	r := _isNumericArray(
		unsafe.Pointer(&src[0]),
		unsafe.Pointer(len(src)),
		unsafe.Pointer(&result),
	)
	return result == 0x1
}

func LowerCase(src []byte) {
	_lowerCase(
		unsafe.Pointer(&src[0]),
		unsafe.Pointer(len(src)),
		unsafe.Pointer(&src)
	)
}

func IsDigit(b byte) bool {
	r := _isDigit(b)
	return r == 0x1
}