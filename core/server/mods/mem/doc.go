// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package mem

/*
initial package performance:

BenchmarkMemStatus/instantiate-struct-4         	2000000000	         0.34 ns/op	2973.27 MB/s	       0 B/op	       0 allocs/op
BenchmarkMemStatus/instantiate-ptr-4            	2000000000	         0.34 ns/op	2939.79 MB/s	       0 B/op	       0 allocs/op
BenchmarkMemStatus/instantiate-internal-4       	 2000000	      1099 ns/op	   0.91 MB/s	    6184 B/op	       3 allocs/op
BenchmarkMemStatus/start-4                      	  1000000	      1563 ns/op	   0.64 MB/s	     524 B/op	       1 allocs/op
PASS

But now we have to make sure about concurrent access

*/
