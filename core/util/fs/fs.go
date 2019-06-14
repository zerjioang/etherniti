package fs

import (
	"io/ioutil"
	"os"
)

var (
	empty = []byte("")
)

func Exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func ReadAll(path string) []byte {
	content, err := ioutil.ReadFile(path)
	if err == nil && content != nil && len(content) > 0 {
		return content
	}
	return empty
}
