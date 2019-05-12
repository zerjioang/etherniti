// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package profile

/*
initial package performance

BenchmarkConnectionProfile/create-profile-empty-4         	2000000000	         0.43 ns/op	2308.66 MB/s	       0 B/op	       0 allocs/op
BenchmarkConnectionProfile/create-profile-4         	      1000000	      1117 ns/op	   0.89 MB/s	      64 B/op	       2 allocs/op
BenchmarkConnectionProfile/valid-false-4                  	200000000	         7.09 ns/op	 140.95 MB/s	       0 B/op	       0 allocs/op
BenchmarkConnectionProfile/valid-true-4                   	200000000	         8.19 ns/op	 122.09 MB/s	       0 B/op	       0 allocs/op
BenchmarkConnectionProfile/get-secret-4                   	2000000000	         0.45 ns/op	2242.62 MB/s	       0 B/op	       0 allocs/op
BenchmarkConnectionProfile/create-token-4                 	  200000	      9184 ns/op	   0.11 MB/s	    3649 B/op	      36 allocs/op
BenchmarkConnectionProfile/parse-token-4                  	   50000	     25569 ns/op	   0.04 MB/s	    5929 B/op	     103 allocs/op

after replacing time with fasttime

BenchmarkConnectionProfile/create-profile-empty-4         	2000000000	         0.40 ns/op	2481.14 MB/s	       0 B/op	       0 allocs/op
BenchmarkConnectionProfile/create-profile-4               	 1000000	      1223 ns/op	   0.82 MB/s	      64 B/op	       2 allocs/op
BenchmarkConnectionProfile/valid-false-4                  	300000000	         5.87 ns/op	 170.29 MB/s	       0 B/op	       0 allocs/op
BenchmarkConnectionProfile/valid-true-4                   	300000000	         5.86 ns/op	 170.78 MB/s	       0 B/op	       0 allocs/op
BenchmarkConnectionProfile/get-secret-4                   	2000000000	         0.41 ns/op	2463.78 MB/s	       0 B/op	       0 allocs/op
BenchmarkConnectionProfile/create-token-4                 	  200000	     10693 ns/op	   0.09 MB/s	    3713 B/op	      36 allocs/op
BenchmarkConnectionProfile/parse-token-4                  	   50000	     23192 ns/op	   0.04 MB/s	    5929 B/op	     103 allocs/op
*/
