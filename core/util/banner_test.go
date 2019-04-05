// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package util

import "testing"

func TestWelcomeBanner(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WelcomeBanner(); got != tt.want {
				t.Errorf("WelcomeBanner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getBannerFromTemplate(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBannerFromTemplate(); got != tt.want {
				t.Errorf("getBannerFromTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}
