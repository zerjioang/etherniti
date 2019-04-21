// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package api

/*
package functions performance:

BenchmarkWrapper/send-success-4         	  200000	      6131 ns/op	   0.16 MB/s	     289 B/op	       5 allocs/op
BenchmarkWrapper/send-success-blob-4    	  500000	      2565 ns/op	   0.39 MB/s	      20 B/op	       1 allocs/op
BenchmarkWrapper/success-4              	  500000	      3777 ns/op	   0.26 MB/s	     261 B/op	       5 allocs/op
BenchmarkWrapper/error-str-4            	  500000	      3634 ns/op	   0.28 MB/s	     280 B/op	       5 allocs/op
BenchmarkWrapper/error-4                	  500000	      3625 ns/op	   0.28 MB/s	     266 B/op	       4 allocs/op
BenchmarkWrapper/error-code-4           	  300000	      3649 ns/op	   0.27 MB/s	     348 B/op	       4 allocs/op
BenchmarkWrapper/stack-error-4          	  500000	      3708 ns/op	   0.27 MB/s	     298 B/op	       5 allocs/op

after using Blob in responses instead of JsonBlob:

BenchmarkWrapper/send-success-4         	  200000	      6148 ns/op	   0.16 MB/s	     289 B/op	       5 allocs/op
BenchmarkWrapper/send-success-blob-4    	  500000	      2539 ns/op	   0.39 MB/s	      20 B/op	       1 allocs/op
BenchmarkWrapper/success-4              	  300000	      3782 ns/op	   0.26 MB/s	     241 B/op	       5 allocs/op
BenchmarkWrapper/error-str-4            	  500000	      3677 ns/op	   0.27 MB/s	     280 B/op	       5 allocs/op
BenchmarkWrapper/error-4                	  500000	      3642 ns/op	   0.27 MB/s	     266 B/op	       4 allocs/op
BenchmarkWrapper/error-code-4           	  500000	      3619 ns/op	   0.28 MB/s	     266 B/op	       4 allocs/op
BenchmarkWrapper/stack-error-4          	  500000	      3695 ns/op	   0.27 MB/s	     298 B/op	       5 allocs/op

*/
