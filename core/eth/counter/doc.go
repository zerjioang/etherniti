// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package counter

/*
this package implements an concurrency safe atomic uint32 data structure

Initial performance:

BenchmarkCounterPtr/instantiate-4         	2000000000	         0.32 ns/op	3121.65 MB/s	       0 B/op	       0 allocs/op
BenchmarkCounterPtr/add-4                 	200000000	         5.98 ns/op	 167.33 MB/s	       0 B/op	       0 allocs/op
BenchmarkCounterPtr/get-4                 	2000000000	         0.32 ns/op	3099.45 MB/s	       0 B/op	       0 allocs/op

BenchmarkCounter/instantiate-4            	2000000000	         0.32 ns/op	3156.49 MB/s	       0 B/op	       0 allocs/op
BenchmarkCounter/add-4                    	100000000	        21.8 ns/op	  45.83 MB/s	       4 B/op	       1 allocs/op
BenchmarkCounter/get-4                    	100000000	        15.5 ns/op	  64.58 MB/s	       4 B/op	       1 allocs/op

As can be seen non pointer based implementation is much slower ,and thus, is use is not recommended
*/
