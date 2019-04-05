// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package handlers

import (
	"reflect"
	"testing"
)

func TestNewRopstenController(t *testing.T) {
	tests := []struct {
		name string
		want EthereumPublicController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRopstenController(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRopstenController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRinkebyController(t *testing.T) {
	tests := []struct {
		name string
		want EthereumPublicController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRinkebyController(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRinkebyController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewKovanController(t *testing.T) {
	tests := []struct {
		name string
		want EthereumPublicController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewKovanController(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewKovanController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMainNetController(t *testing.T) {
	tests := []struct {
		name string
		want EthereumPublicController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMainNetController(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMainNetController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewInfuraController(t *testing.T) {
	tests := []struct {
		name string
		want EthereumPublicController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInfuraController(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInfuraController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewQuikNodeController(t *testing.T) {
	tests := []struct {
		name string
		want EthereumPublicController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewQuikNodeController(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQuikNodeController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPrivateNetController(t *testing.T) {
	tests := []struct {
		name string
		want EthereumPrivateController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPrivateNetController(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPrivateNetController() = %v, want %v", got, tt.want)
			}
		})
	}
}
