// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package api

import (
	"testing"

	"github.com/zerjioang/etherniti/core/util/ip"

	"github.com/armon/go-radix"
)

func BenchmarkRadixResolve(b *testing.B) {
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

	b.SetBytes(1)
	key := "1.41.132.176"
	// run the Fib function b.N times
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = r.Get(key)
	}
}

func BenchmarkMapResolve(b *testing.B) {
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

	b.SetBytes(1)
	// run the Fib function b.N times
	key := "1.41.132.176"
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = mapper[key]
	}
}

func BenchmarkMapUint32Resolve(b *testing.B) {
	// Create a tree
	mapper := make(map[uint32]interface{})
	/*
		1.163.34.119
		1.172.104.133
		1.41.132.176
		100.1.197.216
	*/
	mapper[ip.Ip2int("1.163.34.119")] = nil
	mapper[ip.Ip2int("1.172.104.133")] = nil
	mapper[ip.Ip2int("1.41.132.176")] = nil

	b.SetBytes(1)
	key := ip.Ip2int("1.41.132.176")
	// run the Fib function b.N times
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = mapper[key]
	}
}
