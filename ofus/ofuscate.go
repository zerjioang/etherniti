package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyz"
)
var (
	//alphabet byte array
	alphabetRaw = []byte(alphabet)

	pathMap map[string]string
	pathCounter int
)
func main() {
	pathMap = make(map[string]string)
	//scan the source code
	err := filepath.Walk(".", visitor)
	if err != nil {
		log.Println(err)
	}
	// one scanner, start ofuscating
	ofusErr := ofuscate()
	if ofusErr != nil {
		log.Println(ofusErr)
	}
}

func ofuscate() error {
	for k, v := range pathMap {
		log.Printf("processing dir %s \t\t moving to %s\n", k, v)
		err := copyFolder(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func visitor(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	skip := path == "." ||
		path == ".git" ||
		path == ".idea" ||
		path == "docs" ||
		path == "resources" ||
		path == "scripts" ||
		path == "vendor" ||
		path == "ofus" ||
		strings.HasPrefix(path,".git/") ||
		strings.HasPrefix(path,"vendor/") ||
		strings.HasPrefix(path,"resources/") ||
		strings.HasPrefix(path,"scripts/")
	if !skip {
		if info.IsDir() {
			//add item to pathmap
			ofuscatedPathName := ofuscatePath(pathCounter)
			// add prefix
			ofuscatedPathName = "out/" + ofuscatedPathName
			pathMap[path] = ofuscatedPathName
			fmt.Println(path, ofuscatedPathName, info.Size())
			pathCounter++
		}
	}
	return nil
}

func ofuscatePath(pathCounter int) string {
	firstSize := len(alphabetRaw)
	if pathCounter < firstSize {
		//path can be encoded as single character
		return string(alphabetRaw[pathCounter])
	} else {
		idx := pathCounter/firstSize - 1 //arrays start a 0
		current := string(alphabetRaw[idx])
		return current + ofuscatePath(pathCounter-firstSize)
	}
}