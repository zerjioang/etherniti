// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package lib

import (
	"os"
	"reflect"
	"testing"
)

func TestNewOfuscator(t *testing.T) {
	tests := []struct {
		name string
		want Ofuscator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOfuscator(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOfuscator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOfuscator_Start(t *testing.T) {
	type fields struct {
		rootPath       string
		mainFilePath   string
		internalMapper map[string]OfusEntry
		pathCounter    uint64
		packageCounter uint64
	}
	type args struct {
		path         string
		mainFilePath string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			of := &Ofuscator{
				rootPath:       tt.fields.rootPath,
				mainFilePath:   tt.fields.mainFilePath,
				internalMapper: tt.fields.internalMapper,
				pathCounter:    tt.fields.pathCounter,
				packageCounter: tt.fields.packageCounter,
			}
			of.Start(tt.args.path, tt.args.mainFilePath)
		})
	}
}

func TestOfuscator_ofuscate(t *testing.T) {
	type fields struct {
		rootPath       string
		mainFilePath   string
		internalMapper map[string]OfusEntry
		pathCounter    uint64
		packageCounter uint64
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
			of := &Ofuscator{
				rootPath:       tt.fields.rootPath,
				mainFilePath:   tt.fields.mainFilePath,
				internalMapper: tt.fields.internalMapper,
				pathCounter:    tt.fields.pathCounter,
				packageCounter: tt.fields.packageCounter,
			}
			if err := of.ofuscate(); (err != nil) != tt.wantErr {
				t.Errorf("Ofuscator.ofuscate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOfuscator_visitor(t *testing.T) {
	type fields struct {
		rootPath       string
		mainFilePath   string
		internalMapper map[string]OfusEntry
		pathCounter    uint64
		packageCounter uint64
	}
	type args struct {
		file string
		info os.FileInfo
		err  error
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			of := &Ofuscator{
				rootPath:       tt.fields.rootPath,
				mainFilePath:   tt.fields.mainFilePath,
				internalMapper: tt.fields.internalMapper,
				pathCounter:    tt.fields.pathCounter,
				packageCounter: tt.fields.packageCounter,
			}
			if err := of.visitor(tt.args.file, tt.args.info, tt.args.err); (err != nil) != tt.wantErr {
				t.Errorf("Ofuscator.visitor() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOfuscator_addAsProcessableFile(t *testing.T) {
	type fields struct {
		rootPath       string
		mainFilePath   string
		internalMapper map[string]OfusEntry
		pathCounter    uint64
		packageCounter uint64
	}
	type args struct {
		file             string
		parent           string
		ofuscatedPkgName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			of := &Ofuscator{
				rootPath:       tt.fields.rootPath,
				mainFilePath:   tt.fields.mainFilePath,
				internalMapper: tt.fields.internalMapper,
				pathCounter:    tt.fields.pathCounter,
				packageCounter: tt.fields.packageCounter,
			}
			of.addAsProcessableFile(tt.args.file, tt.args.parent, tt.args.ofuscatedPkgName)
		})
	}
}

func TestOfuscator_addAsEntryPoint(t *testing.T) {
	type fields struct {
		rootPath       string
		mainFilePath   string
		internalMapper map[string]OfusEntry
		pathCounter    uint64
		packageCounter uint64
	}
	type args struct {
		file string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			of := &Ofuscator{
				rootPath:       tt.fields.rootPath,
				mainFilePath:   tt.fields.mainFilePath,
				internalMapper: tt.fields.internalMapper,
				pathCounter:    tt.fields.pathCounter,
				packageCounter: tt.fields.packageCounter,
			}
			of.addAsEntryPoint(tt.args.file)
		})
	}
}

func TestOfuscator_addAsNoProcessableFile(t *testing.T) {
	type fields struct {
		rootPath       string
		mainFilePath   string
		internalMapper map[string]OfusEntry
		pathCounter    uint64
		packageCounter uint64
	}
	type args struct {
		file     string
		basename string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			of := &Ofuscator{
				rootPath:       tt.fields.rootPath,
				mainFilePath:   tt.fields.mainFilePath,
				internalMapper: tt.fields.internalMapper,
				pathCounter:    tt.fields.pathCounter,
				packageCounter: tt.fields.packageCounter,
			}
			of.addAsNoProcessableFile(tt.args.file, tt.args.basename)
		})
	}
}

func TestOfuscator_mapPackageName(t *testing.T) {
	type fields struct {
		rootPath       string
		mainFilePath   string
		internalMapper map[string]OfusEntry
		pathCounter    uint64
		packageCounter uint64
	}
	type args struct {
		basedir  string
		basename string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		want1  string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			of := &Ofuscator{
				rootPath:       tt.fields.rootPath,
				mainFilePath:   tt.fields.mainFilePath,
				internalMapper: tt.fields.internalMapper,
				pathCounter:    tt.fields.pathCounter,
				packageCounter: tt.fields.packageCounter,
			}
			got, got1 := of.mapPackageName(tt.args.basedir, tt.args.basename)
			if got != tt.want {
				t.Errorf("Ofuscator.mapPackageName() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Ofuscator.mapPackageName() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOfuscator_skipPath(t *testing.T) {
	type fields struct {
		rootPath       string
		mainFilePath   string
		internalMapper map[string]OfusEntry
		pathCounter    uint64
		packageCounter uint64
	}
	type args struct {
		file string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			of := Ofuscator{
				rootPath:       tt.fields.rootPath,
				mainFilePath:   tt.fields.mainFilePath,
				internalMapper: tt.fields.internalMapper,
				pathCounter:    tt.fields.pathCounter,
				packageCounter: tt.fields.packageCounter,
			}
			if got := of.skipPath(tt.args.file); got != tt.want {
				t.Errorf("Ofuscator.skipPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOfuscator_isWhiteListed(t *testing.T) {
	type fields struct {
		rootPath       string
		mainFilePath   string
		internalMapper map[string]OfusEntry
		pathCounter    uint64
		packageCounter uint64
	}
	type args struct {
		filename string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			of := Ofuscator{
				rootPath:       tt.fields.rootPath,
				mainFilePath:   tt.fields.mainFilePath,
				internalMapper: tt.fields.internalMapper,
				pathCounter:    tt.fields.pathCounter,
				packageCounter: tt.fields.packageCounter,
			}
			if got := of.isWhiteListed(tt.args.filename); got != tt.want {
				t.Errorf("Ofuscator.isWhiteListed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOfuscator_counterToName(t *testing.T) {
	type fields struct {
		rootPath       string
		mainFilePath   string
		internalMapper map[string]OfusEntry
		pathCounter    uint64
		packageCounter uint64
	}
	type args struct {
		pathCounter uint64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			of := Ofuscator{
				rootPath:       tt.fields.rootPath,
				mainFilePath:   tt.fields.mainFilePath,
				internalMapper: tt.fields.internalMapper,
				pathCounter:    tt.fields.pathCounter,
				packageCounter: tt.fields.packageCounter,
			}
			if got := of.counterToName(tt.args.pathCounter); got != tt.want {
				t.Errorf("Ofuscator.counterToName() = %v, want %v", got, tt.want)
			}
		})
	}
}
