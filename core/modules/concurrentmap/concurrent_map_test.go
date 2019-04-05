// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package concurrentmap

import (
	"reflect"
	"testing"
)

func TestNewConcurrentMap(t *testing.T) {
	tests := []struct {
		name string
		want ConcurrentMap
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConcurrentMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConcurrentMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConcurrentMap_GetShard(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		m    ConcurrentMap
		args args
		want *ConcurrentMapShared
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.GetShard(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConcurrentMap.GetShard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConcurrentMap_MSet(t *testing.T) {
	type args struct {
		data map[string]interface{}
	}
	tests := []struct {
		name string
		m    ConcurrentMap
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.MSet(tt.args.data)
		})
	}
}

func TestConcurrentMap_Set(t *testing.T) {
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name string
		m    ConcurrentMap
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.Set(tt.args.key, tt.args.value)
		})
	}
}

func TestConcurrentMap_Upsert(t *testing.T) {
	type args struct {
		key   string
		value interface{}
		cb    UpsertCb
	}
	tests := []struct {
		name    string
		m       ConcurrentMap
		args    args
		wantRes interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := tt.m.Upsert(tt.args.key, tt.args.value, tt.args.cb); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("ConcurrentMap.Upsert() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestConcurrentMap_SetIfAbsent(t *testing.T) {
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name string
		m    ConcurrentMap
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.SetIfAbsent(tt.args.key, tt.args.value); got != tt.want {
				t.Errorf("ConcurrentMap.SetIfAbsent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConcurrentMap_Get(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name  string
		m     ConcurrentMap
		args  args
		want  interface{}
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.m.Get(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConcurrentMap.Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ConcurrentMap.Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestConcurrentMap_Count(t *testing.T) {
	tests := []struct {
		name string
		m    ConcurrentMap
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Count(); got != tt.want {
				t.Errorf("ConcurrentMap.Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConcurrentMap_Has(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		m    ConcurrentMap
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Has(tt.args.key); got != tt.want {
				t.Errorf("ConcurrentMap.Has() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConcurrentMap_Remove(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		m    ConcurrentMap
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.Remove(tt.args.key)
		})
	}
}

func TestConcurrentMap_Pop(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name       string
		m          ConcurrentMap
		args       args
		wantV      interface{}
		wantExists bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotV, gotExists := tt.m.Pop(tt.args.key)
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("ConcurrentMap.Pop() gotV = %v, want %v", gotV, tt.wantV)
			}
			if gotExists != tt.wantExists {
				t.Errorf("ConcurrentMap.Pop() gotExists = %v, want %v", gotExists, tt.wantExists)
			}
		})
	}
}

func TestConcurrentMap_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		m    ConcurrentMap
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.IsEmpty(); got != tt.want {
				t.Errorf("ConcurrentMap.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConcurrentMap_Iter(t *testing.T) {
	tests := []struct {
		name string
		m    ConcurrentMap
		want <-chan Tuple
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Iter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConcurrentMap.Iter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConcurrentMap_IterBuffered(t *testing.T) {
	tests := []struct {
		name string
		m    ConcurrentMap
		want <-chan Tuple
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.IterBuffered(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConcurrentMap.IterBuffered() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_snapshot(t *testing.T) {
	type args struct {
		m ConcurrentMap
	}
	tests := []struct {
		name      string
		args      args
		wantChans []chan Tuple
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotChans := snapshot(tt.args.m); !reflect.DeepEqual(gotChans, tt.wantChans) {
				t.Errorf("snapshot() = %v, want %v", gotChans, tt.wantChans)
			}
		})
	}
}

func Test_fanIn(t *testing.T) {
	type args struct {
		chans []chan Tuple
		out   chan Tuple
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fanIn(tt.args.chans, tt.args.out)
		})
	}
}

func TestConcurrentMap_Items(t *testing.T) {
	tests := []struct {
		name string
		m    ConcurrentMap
		want map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Items(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConcurrentMap.Items() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConcurrentMap_IterCb(t *testing.T) {
	type args struct {
		fn IterCb
	}
	tests := []struct {
		name string
		m    ConcurrentMap
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.IterCb(tt.args.fn)
		})
	}
}

func TestConcurrentMap_Keys(t *testing.T) {
	tests := []struct {
		name string
		m    ConcurrentMap
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Keys(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConcurrentMap.Keys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConcurrentMap_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		m       ConcurrentMap
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("ConcurrentMap.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConcurrentMap.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fnv32(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fnv32(tt.args.key); got != tt.want {
				t.Errorf("fnv32() = %v, want %v", got, tt.want)
			}
		})
	}
}
