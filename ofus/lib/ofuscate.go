package lib

import (
	"bytes"
	"log"
	"os"
	"path"
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
	fileMap map[string]string
	pathCounter int
)
type Ofuscator struct {

}

func NewOfuscator() Ofuscator {
	return Ofuscator{}
}

func (of Ofuscator) Start(path string) {
	pathMap = make(map[string]string)
	fileMap = make(map[string]string)
	//scan the source code
	err := filepath.Walk(path, visitor)
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
	createErr := os.MkdirAll("out", os.ModePerm)
	if createErr != nil {
		return createErr
	}
	for k, v := range pathMap {
		log.Printf("processing file %s moving to %s\n", k, v)
		err := copyFile(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func visitor(file string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	skip := strings.Contains(file, "/.idea") ||
		strings.Contains(file, "/.git") ||
		strings.Contains(file, "/docs") ||
		strings.Contains(file, "/vendor") ||
		strings.Contains(file, "/resources") ||
		strings.Contains(file, "/scripts")
	if !skip {
		if !info.IsDir() {
			//add item to pathmap
			// get base path
			basedir, basename := path.Split(file)
			log.Println("processing", basedir)

			if strings.HasPrefix(basename, ".") || isWhiteListed(basename) {
				//file is considered as hidden file. keep it that way
				log.Println("skipping ofuscation of file", file)
				pathMap[file] = "out/" + basename
			} else {
				ofuscatedPathName := ofuscatePath(pathCounter)
				// add prefix
				ofuscatedPathName = "out/" + ofuscatedPathName
				// add extension
				extension := filepath.Ext(file)
				// processing basename
				ofuscatedPathName = ofuscatedPathName + extension
				pathMap[file] = ofuscatedPathName
				pathCounter++
			}
		}
	}
	return nil
}

func isWhiteListed(filename string) bool {
	return filename == "AUTHORS" ||
		filename == "LICENSE" ||
		filename == "VERSION" ||
		filename == "MAINTAINERS.md" ||
		filename == "CONTRIBUTING.md" ||
		filename == "Gopkg.toml" ||
		filename == "Gopkg.lock" ||
		filename == "README.md"
}

func ofuscatePath(pathCounter int) string {
	var buf bytes.Buffer
	encode(uint64(pathCounter), &buf, alphabet)
	return buf.String()
}