// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package api

/*
package functions performance:

BenchmarkRadixResolve-4       	30000000	        47.2 ns/op	  21.19 MB/s	       0 B/op	       0 allocs/op
BenchmarkMapResolve-4         	100000000	        11.4 ns/op	  87.45 MB/s	       0 B/op	       0 allocs/op
BenchmarkMapUint32Resolve-4   	300000000	         6.57 ns/op	 152.26 MB/s	       0 B/op	       0 allocs/op       0 allocs/op

## Bad bot list

BenchmarkBadBot/first-item-access-4         	2000000000	         0.54 ns/op	1846.87 MB/s	       0 B/op	       0 allocs/op
BenchmarkBadBot/last-item-access-4          	2000000000	         0.35 ns/op	2885.50 MB/s	       0 B/op	       0 allocs/op

*/
