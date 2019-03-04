// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

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
)

type Ofuscator struct {
	rootPath       string
	mainFilePath   string
	internalMapper map[string]OfusEntry
	pathCounter    uint64
	packageCounter uint64
}

func NewOfuscator() Ofuscator {
	return Ofuscator{}
}

func (of *Ofuscator) Start(path string, mainFilePath string) {
	// set the analysis root path
	of.rootPath = path
	of.mainFilePath = mainFilePath

	of.internalMapper = make(map[string]OfusEntry)

	//scan the source code
	err := filepath.Walk(path, of.visitor)
	if err != nil {
		log.Println(err)
	}
	// once scanned, start ofuscating
	ofusErr := of.ofuscate()
	if ofusErr != nil {
		log.Println(ofusErr)
	}
}

func (of *Ofuscator) ofuscate() error {
	createErr := os.MkdirAll("out", os.ModePerm)
	if createErr != nil {
		return createErr
	}
	for k, v := range of.internalMapper {
		// read ofuscated base dir path, and create if not exists
		createErr := v.createDstOfuscatedDir()
		if createErr != nil {
			return createErr
		}
		// read ofuscated file path
		dst := v.OfuscatedFilePath()
		if v.ofuscate {
			log.Printf("processing file %s moving to %s\n", k, dst)
			err := copyFile(k, dst)
			if err != nil {
				return err
			}
		} else {
			log.Printf("skipping ofuscation of file %s", v.originalPath)
		}
	}
	return nil
}

func (of *Ofuscator) visitor(file string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	skip := of.skipPath(file)
	isDir := info.IsDir()
	isMainFile := file == of.mainFilePath
	if !skip {
		if isDir {
			log.Println("skipping directory:", file)
		} else {
			if isMainFile {
				//do not ofuscate program entry point
				log.Println("entry point detected at", file)
				of.addAsEntryPoint(file)
			} else {
				basedir, basename := path.Split(file)
				if strings.HasPrefix(basename, ".") || of.isWhiteListed(basename) {
					of.addAsNoProcessableFile(file, basename)
				} else {
					// calculate mapping between current path and package name
					parent, ofuscatedPgkName := of.mapPackageName(basedir, basename)
					of.addAsProcessableFile(file, parent, ofuscatedPgkName)
				}
			}
		}
	}
	return nil
}

func (of *Ofuscator) addAsProcessableFile(file string, parent string, ofuscatedPkgName string) {
	//ofuscate current file name
	//ofuscate filename
	ofuscatedName := of.counterToName(of.pathCounter)
	of.pathCounter++

	// add extension
	extension := filepath.Ext(file)

	entryItem := OfusEntry{
		originalPath:         file,
		ofuscatedFilename:    ofuscatedName,
		ofuscatedPackageName: ofuscatedPkgName,
		extension:            extension,
		parentDir:            parent,
		idx:                  of.pathCounter,
		ofuscate:             true,
	}

	// processing basename
	of.internalMapper[file] = entryItem
}

func (of *Ofuscator) addAsEntryPoint(file string) {
	extension := filepath.Ext(file)
	ofuscatedName := of.counterToName(of.pathCounter)
	of.pathCounter++
	ofuscatedName = "main_" + ofuscatedName

	entryItem := OfusEntry{
		originalPath:         file,
		ofuscatedFilename:    ofuscatedName,
		ofuscatedPackageName: "main",
		extension:            extension,
		parentDir:            "",
		idx:                  of.pathCounter,
		ofuscate:             true,
	}

	of.internalMapper[file] = entryItem
}

func (of *Ofuscator) addAsNoProcessableFile(file string, basename string) {
	//file is considered as hidden file. keep it that way
	log.Println("skipping ofuscation of file", file)

	entryItem := OfusEntry{
		originalPath:         file,
		ofuscatedFilename:    basename,
		ofuscatedPackageName: "",
		extension:            "",
		parentDir:            "",
		idx:                  0,
		ofuscate:             true,
	}

	of.internalMapper[file] = entryItem
}

func (of *Ofuscator) mapPackageName(basedir string, basename string) (string, string) {
	var readedEntry OfusEntry
	var parent, ofuscatedPgkName string
	var ok bool
	// check if this path was already ofuscated before
	readedEntry, ok = of.internalMapper[basedir]
	ofuscatedPgkName = readedEntry.ofuscatedPackageName
	valid := ok && len(ofuscatedPgkName) > 0
	if !valid {
		//read current package name from path
		parent = filepath.Base(basedir)
		//generate an ofuscated name for current package
		ofuscatedPgkName = of.counterToName(of.packageCounter)
		of.packageCounter++
		if len(parent) > 0 {
			//link current package name with ofuscated name for ofuscation later
			readedEntry.ofuscatedPackageName = ofuscatedPgkName
			of.internalMapper[basedir] = readedEntry
			log.Println("packagename:", parent, "ofuscated:", ofuscatedPgkName)
		}
	}
	log.Println("basedir:", basedir, "basename:", basename, "pkg:", ofuscatedPgkName)
	return parent, ofuscatedPgkName
}

// concurrency safe
func (of Ofuscator) skipPath(file string) bool {
	return strings.Contains(file, "/.idea") ||
		strings.Contains(file, "/.git") ||
		strings.Contains(file, "/docs") ||
		strings.Contains(file, "/vendor") ||
		strings.Contains(file, "/out") ||
		strings.Contains(file, "/ofus") ||
		strings.Contains(file, "/resources") ||
		strings.Contains(file, "/scripts")
}

// concurrency safe
func (of Ofuscator) isWhiteListed(filename string) bool {
	return filename == "AUTHORS" ||
		filename == "LICENSE" ||
		filename == "VERSION" ||
		filename == "MAINTAINERS.md" ||
		filename == "CONTRIBUTING.md" ||
		filename == "Gopkg.toml" ||
		filename == "Gopkg.lock" ||
		filename == "README.md"
}

// concurrency safe
func (of Ofuscator) counterToName(pathCounter uint64) string {
	var buf bytes.Buffer
	encode(pathCounter, &buf, alphabet)
	return buf.String()
}
