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

after adding some sync pools and replacing json serializer with *bytes.buffer and manual marshalling

BenchmarkWrapper/to-success-4         	 3000000	       525 ns/op	   1.90 MB/s	      26 B/op	       3 allocs/op
BenchmarkWrapper/to-success-pool-4    	 3000000	       602 ns/op	   1.66 MB/s	      30 B/op	       4 allocs/op
BenchmarkWrapper/to-error-4           	20000000	       108 ns/op	   9.18 MB/s	       0 B/op	       0 allocs/op
BenchmarkWrapper/to-error-pool-4      	10000000	       176 ns/op	   5.68 MB/s	       3 B/op	       1 allocs/op
BenchmarkWrapper/send-success-4       	  500000	      3127 ns/op	   0.32 MB/s	     432 B/op	       8 allocs/op
BenchmarkWrapper/send-success-pool-4  	  500000	      3182 ns/op	   0.31 MB/s	     432 B/op	       9 allocs/op
BenchmarkWrapper/send-success-blob-4  	 1000000	      1073 ns/op	   0.93 MB/s	     184 B/op	       2 allocs/op
BenchmarkWrapper/success-4            	 1000000	      1702 ns/op	   0.59 MB/s	     248 B/op	       7 allocs/op
BenchmarkWrapper/error-str-4          	 1000000	      1199 ns/op	   0.83 MB/s	     224 B/op	       4 allocs/op
BenchmarkWrapper/error-4              	 1000000	      1348 ns/op	   0.74 MB/s	     224 B/op	       4 allocs/op
BenchmarkWrapper/error-code-4         	 1000000	      1164 ns/op	   0.86 MB/s	     192 B/op	       3 allocs/op
BenchmarkWrapper/stack-error-4        	 1000000	      1306 ns/op	   0.77 MB/s	     224 B/op	       4 allocs/op
*/
