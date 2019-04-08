// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package hashset

/*
Initial Benchmarking

BenchmarkHashSet/instantiate-4         							30000000	        59.0 ns/op	      48 B/op	       1 allocs/op
BenchmarkHashSet/add/simple-4          							100000000	        16.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashSet/add/10000-items-4     	    					2000              830986 ns/op	   39367 B/op	    9900 allocs/op
BenchmarkHashSet/contains/simple-4     							200000000	        6.09 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashSet/contains/10000-items/contains-first-4         	100000000	        21.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashSet/contains/10000-items/contains-middle-4        	50000000	        24.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashSet/contains/10000-items/contains-last-4          	50000000	        28.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashSet/count-0-4                                     	2000000000	        0.33 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashSet/count-10000-4                                 	2000000000	        0.40 ns/op	       0 B/op	       0 allocs/op

*/
