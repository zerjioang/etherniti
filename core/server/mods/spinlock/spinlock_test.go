// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package spinlock

import "testing"

func TestSpinLock_TyLock(t *testing.T) {
	type fields struct {
		state int32
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SpinLock{
				state: tt.fields.state,
			}
			if got := s.TyLock(); got != tt.want {
				t.Errorf("SpinLock.TyLock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpinLock_lock(t *testing.T) {
	type fields struct {
		state int32
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SpinLock{
				state: tt.fields.state,
			}
			if got := s.lock(); got != tt.want {
				t.Errorf("SpinLock.lock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpinLock_IsLocked(t *testing.T) {
	type fields struct {
		state int32
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SpinLock{
				state: tt.fields.state,
			}
			if got := s.IsLocked(); got != tt.want {
				t.Errorf("SpinLock.IsLocked() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpinLock_Unlock(t *testing.T) {
	type fields struct {
		state int32
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SpinLock{
				state: tt.fields.state,
			}
			s.Unlock()
		})
	}
}
