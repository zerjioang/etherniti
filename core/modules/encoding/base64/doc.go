// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package base64

/*
http://0x80.pl/notesen/2015-12-27-base64-encoding.html
http://0x80.pl/articles/index.html#base64-algorithm-update
*/

/*

Initial performance:

BenchmarkBase64/go/encode-decode-4         	 3000000	       434 ns/op	   2.30 MB/s	     176 B/op	       4 allocs/op
BenchmarkBase64/go/encode-4                	10000000	       166 ns/op	   6.00 MB/s	      96 B/op	       2 allocs/op
BenchmarkBase64/custom/encode-4            	10000000	       162 ns/op	   6.14 MB/s	      96 B/op	       2 allocs/op

Adding support for streaming operations

BenchmarkBase64/go/encode-decode-4         	 3000000	       380 ns/op	   2.63 MB/s	     176 B/op	       4 allocs/op
BenchmarkBase64/go/encode-4                	10000000	       167 ns/op	   5.95 MB/s	      96 B/op	       2 allocs/op
BenchmarkBase64/custom/encode-string-4     	10000000	       165 ns/op	   6.05 MB/s	      96 B/op	       2 allocs/op
BenchmarkBase64/custom/encode-stream-4     	10000000	       126 ns/op	   7.92 MB/s	       0 B/op	       0 allocs/op

*/
