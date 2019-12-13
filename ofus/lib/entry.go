// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package lib

import (
	"os"
	"strings"
)

type OfusEntry struct {
	originalPath         string
	ofuscatedPackageName string
	ofuscatedFilename    string
	extension            string
	parentDir            string
	idx                  uint64
	ofuscate             bool
}

func (entry OfusEntry) OfuscatedBasePath() string {
	if entry.ofuscatedPackageName == "" {
		return "out/src/elf/"
	}
	return "out/src/elf/" + entry.ofuscatedPackageName
}

func (entry OfusEntry) OfuscatedFilePath() string {
	isGoTestFile := strings.LastIndex(strings.ToLower(entry.originalPath), "_test.go") != -1
	if isGoTestFile {
		// process test file
		ofusName := "out/src/elf/" + entry.ofuscatedPackageName + "/" + entry.ofuscatedFilename
		if entry.ofuscatedPackageName == "" {
			ofusName = "out/src/elf/" + entry.ofuscatedFilename
		}
		return ofusName + "_test.go"
	} else {
		ofusName := "out/src/elf/" + entry.ofuscatedPackageName + "/" + entry.ofuscatedFilename + entry.extension
		if entry.ofuscatedPackageName == "" {
			ofusName = "out/src/elf/" + entry.ofuscatedFilename + entry.extension
		}
		return ofusName
	}
}

func (entry OfusEntry) createDstOfuscatedDir() error {
	outpath := entry.OfuscatedBasePath()
	return os.MkdirAll(outpath, os.ModePerm)
}
