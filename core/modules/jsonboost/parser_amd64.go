package jsonboost

import (
	"github.com/zerjioang/etherniti/core/logger"
	"unsafe"
)

//go:noescape
// parts(char* src, char* dst, int start, int end, int steps){
func _parts(src, dst unsafe.Pointer, start, end, steps int) (slice int)

//go:noescape
func _lookup(json, key, keyLen, jsonLen, result unsafe.Pointer) (error int)

func Lookup(json string, key string) string {

	if json == "" {
		logger.Error("INVALID_JSON_LEN error")
		return ""
	}
	if key == "" {
		logger.Error("INVALID_KEY_LEN error")
		return ""
	}

	// 1. convert strings to unsafe.pointers
	keyRaw := StringToBytes(key)
	jsonRaw := StringToBytes(json)
	var resultRaw string
	resultPtr := unsafe.Pointer(&resultRaw)

	errorCode := _lookup(
		unsafe.Pointer(&jsonRaw[0]), // json key we want to read. simple o dot formatted &keyRaw[0]
		unsafe.Pointer(&keyRaw[0]), // json raw we want to parse looking for key content &jsonRaw[0]
		unsafe.Pointer(uintptr(len(json))), // length of the key json
		unsafe.Pointer(uintptr(len(key))), // length of the raw json
		resultPtr,
	)
	switch errorCode {
	case 3:
		logger.Error("INVALID_JSON_DATA error. make sure json content is valid and starts with [,{ and ends with ],} characters")
		return ""
	}
	//Get the string pointer by address
	stringPtr := (*string)(resultPtr)

	//Get the value at that pointer
	newData := *stringPtr
	return newData
}

//go:noescape
func _add(a, b int) (c int)

func Add(a, b int) int {
	return _add(a, b)
}