// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package base64

/*

Initial performance:

BenchmarkBase64/go/encode-decode-4         	 3000000	       434 ns/op	   2.30 MB/s	     176 B/op	       4 allocs/op
BenchmarkBase64/go/encode-4                	10000000	       166 ns/op	   6.00 MB/s	      96 B/op	       2 allocs/op
BenchmarkBase64/custom/encode-4            	10000000	       162 ns/op	   6.14 MB/s	      96 B/op	       2 allocs/op

*/
