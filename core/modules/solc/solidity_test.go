// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package solc

import (
	"os/exec"
	"reflect"
	"testing"
)

const (
	testSource = `
pragma solidity >0.0.0;
contract test {
   /// @notice Will multiply ` + "`a`" + ` by 7.
   function multiply(uint a) public returns(uint d) {
       return a * 7;
   }
}
`
)

func skipWithoutSolc(t *testing.T) {
	if _, err := exec.LookPath("solc"); err != nil {
		t.Skip(err)
	}
}

func TestSolidityCompiler(t *testing.T) {
	skipWithoutSolc(t)

	contracts, err := CompileSolidityString(testSource)
	if err != nil {
		t.Fatalf("error compiling source. result %v: %v", contracts, err)
	}
	if len(contracts) != 1 {
		t.Errorf("one contract expected, got %d", len(contracts))
	}
	c, ok := contracts["test"]
	if !ok {
		c, ok = contracts["<stdin>:test"]
		if !ok {
			t.Fatal("info for contract 'test' not present in result")
		}
	}
	if c.Code == "" {
		t.Error("empty code")
	}
	if c.Info.Source != testSource {
		t.Error("wrong source")
	}
	if c.Info.CompilerVersion == "" {
		t.Error("empty version")
	}
}

func TestSolidityCompileError(t *testing.T) {
	skipWithoutSolc(t)

	contracts, err := CompileSolidityString(testSource[4:])
	if err == nil {
		t.Errorf("error expected compiling source. got none. result %v", contracts)
	}
	t.Logf("error: %v", err)
}

func TestSolidity_makeArgs(t *testing.T) {
	type fields struct {
		Path        string
		Version     string
		FullVersion string
		Major       int
		Minor       int
		Patch       int
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Solidity{
				Path:        tt.fields.Path,
				Version:     tt.fields.Version,
				FullVersion: tt.fields.FullVersion,
				Major:       tt.fields.Major,
				Minor:       tt.fields.Minor,
				Patch:       tt.fields.Patch,
			}
			if got := s.makeArgs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solidity.makeArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolidityVersion(t *testing.T) {
	tests := []struct {
		name    string
		want    *Solidity
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SolidityVersion()
			if (err != nil) != tt.wantErr {
				t.Errorf("SolidityVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SolidityVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompileSolidityString(t *testing.T) {
	type args struct {
		source string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]*Contract
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CompileSolidityString(tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("CompileSolidityString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CompileSolidityString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompileSolidity(t *testing.T) {
	type args struct {
		sourcefiles []string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]*Contract
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CompileSolidity(tt.args.sourcefiles...)
			if (err != nil) != tt.wantErr {
				t.Errorf("CompileSolidity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CompileSolidity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolidity_run(t *testing.T) {
	type fields struct {
		Path        string
		Version     string
		FullVersion string
		Major       int
		Minor       int
		Patch       int
	}
	type args struct {
		cmd    *exec.Cmd
		source string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]*Contract
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Solidity{
				Path:        tt.fields.Path,
				Version:     tt.fields.Version,
				FullVersion: tt.fields.FullVersion,
				Major:       tt.fields.Major,
				Minor:       tt.fields.Minor,
				Patch:       tt.fields.Patch,
			}
			got, err := s.run(tt.args.cmd, tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("Solidity.run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solidity.run() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseCombinedJSON(t *testing.T) {
	type args struct {
		combinedJSON    []byte
		source          string
		languageVersion string
		compilerVersion string
		compilerOptions string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]*Contract
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseCombinedJSON(tt.args.combinedJSON, tt.args.source, tt.args.languageVersion, tt.args.compilerVersion, tt.args.compilerOptions)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseCombinedJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseCombinedJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_skipWithoutSolc(t *testing.T) {
	type args struct {
		t *testing.T
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			skipWithoutSolc(tt.args.t)
		})
	}
}
