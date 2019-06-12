//+build !noasm
//+build !appengine

package string

//go:noescape
func _ip_to_int(buf unsafe.Pointer) (result uint32)

//go:noescape
func _ip_to_int2(buf, size unsafe.Pointer) (result uint32)

func IpToInt(ip []byte) uint32 {
	var result uint32
	r := _ip_to_int(
		unsafe.Pointer(&ip[0]),
	)
	return r
}

func IpToInt2(ip []byte) uint32 {
	var result uint32
	r := _ip_to_int2(
		unsafe.Pointer(&ip[0]),
		unsafe.Pointer(uintptr(len(ip))),
	)
	return r
}
