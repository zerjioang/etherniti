// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package lib

import "testing"

func TestOfusEntry_OfuscatedBasePath(t *testing.T) {
	type fields struct {
		originalPath         string
		ofuscatedPackageName string
		ofuscatedFilename    string
		extension            string
		parentDir            string
		idx                  uint64
		ofuscate             bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entry := OfusEntry{
				originalPath:         tt.fields.originalPath,
				ofuscatedPackageName: tt.fields.ofuscatedPackageName,
				ofuscatedFilename:    tt.fields.ofuscatedFilename,
				extension:            tt.fields.extension,
				parentDir:            tt.fields.parentDir,
				idx:                  tt.fields.idx,
				ofuscate:             tt.fields.ofuscate,
			}
			if got := entry.OfuscatedBasePath(); got != tt.want {
				t.Errorf("OfusEntry.OfuscatedBasePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOfusEntry_OfuscatedFilePath(t *testing.T) {
	type fields struct {
		originalPath         string
		ofuscatedPackageName string
		ofuscatedFilename    string
		extension            string
		parentDir            string
		idx                  uint64
		ofuscate             bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entry := OfusEntry{
				originalPath:         tt.fields.originalPath,
				ofuscatedPackageName: tt.fields.ofuscatedPackageName,
				ofuscatedFilename:    tt.fields.ofuscatedFilename,
				extension:            tt.fields.extension,
				parentDir:            tt.fields.parentDir,
				idx:                  tt.fields.idx,
				ofuscate:             tt.fields.ofuscate,
			}
			if got := entry.OfuscatedFilePath(); got != tt.want {
				t.Errorf("OfusEntry.OfuscatedFilePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOfusEntry_createDstOfuscatedDir(t *testing.T) {
	type fields struct {
		originalPath         string
		ofuscatedPackageName string
		ofuscatedFilename    string
		extension            string
		parentDir            string
		idx                  uint64
		ofuscate             bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entry := OfusEntry{
				originalPath:         tt.fields.originalPath,
				ofuscatedPackageName: tt.fields.ofuscatedPackageName,
				ofuscatedFilename:    tt.fields.ofuscatedFilename,
				extension:            tt.fields.extension,
				parentDir:            tt.fields.parentDir,
				idx:                  tt.fields.idx,
				ofuscate:             tt.fields.ofuscate,
			}
			if err := entry.createDstOfuscatedDir(); (err != nil) != tt.wantErr {
				t.Errorf("OfusEntry.createDstOfuscatedDir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
