// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package lib

import "os"

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
	if entry.ofuscatedPackageName == "" {
		return "out/src/elf/" + entry.ofuscatedFilename + entry.extension
	}
	return "out/src/elf/" + entry.ofuscatedPackageName + "/" + entry.ofuscatedFilename + entry.extension
}

func (entry OfusEntry) createDstOfuscatedDir() error {
	outpath := entry.OfuscatedBasePath()
	return os.MkdirAll(outpath, os.ModePerm)
}
