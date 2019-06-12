package asm

import "unsafe"

//go:noescape
func _Z7bin2hexPhi(vec1, vec2, result unsafe.Pointer)

func Hex2(someObj Object) {

	_Z7bin2hexPhi(someObj.GetVec1(), someObj.GetVec2(), someObj.GetResult()))
}