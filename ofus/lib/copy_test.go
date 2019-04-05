// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package lib

import "testing"

func Test_copyFolder(t *testing.T) {
	type args struct {
		source string
		dest   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := copyFolder(tt.args.source, tt.args.dest); (err != nil) != tt.wantErr {
				t.Errorf("copyFolder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_copyFile(t *testing.T) {
	type args struct {
		source string
		dest   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := copyFile(tt.args.source, tt.args.dest); (err != nil) != tt.wantErr {
				t.Errorf("copyFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
