package api

import (
	"github.com/armon/go-radix"
	"testing"
)
func TestCreateRadix(t *testing.T){
	// Create a tree
	r := radix.New()
	/*
	1.163.34.119
	1.172.104.133
	1.41.132.176
	100.1.197.216
	*/
	r.Insert("1.163.34.119", nil)
	r.Insert("1.172.104.133", nil)
	r.Insert("1.41.132.176", nil)

	// Find the longest prefix match
	_, _ = r.Get("1.41.132.176")
}

// BenchmarkRadixResolve-4   	30000000	        39.7 ns/op	  25.22 MB/s	       0 B/op	       0 allocs/op
func BenchmarkRadixResolve(b *testing.B){
	// Create a tree
	r := radix.New()
	/*
	1.163.34.119
	1.172.104.133
	1.41.132.176
	100.1.197.216
	*/
	r.Insert("1.163.34.119", nil)
	r.Insert("1.172.104.133", nil)
	r.Insert("1.41.132.176", nil)

	b.ReportAllocs()
	b.SetBytes(1)
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		_, _ = r.Get("1.41.132.176")
	}
}
// BenchmarkMapResolve-4   	100000000	        11.9 ns/op	  84.20 MB/s	       0 B/op	       0 allocs/op
func BenchmarkMapResolve(b *testing.B){
	// Create a tree
	mapper := make(map[string]interface{})
	/*
	1.163.34.119
	1.172.104.133
	1.41.132.176
	100.1.197.216
	*/
	mapper["1.163.34.119"] = nil
	mapper["1.172.104.133"] = nil
	mapper["1.41.132.176"] = nil

	b.ReportAllocs()
	b.SetBytes(1)
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		_, _ = mapper["1.41.132.176"]
	}
}