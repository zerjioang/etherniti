package profile

/*
initial package performance

BenchmarkConnectionProfile/instantiate-4         	2000000000	         0.34 ns/op	2937.51 MB/s	       0 B/op	       0 allocs/op
BenchmarkConnectionProfile/create-profile-4          1000000	      1808 ns/op	   0.55 MB/s	     128 B/op	       4 allocs/op
BenchmarkConnectionProfile/valid-false-4         	300000000	         5.44 ns/op	 183.84 MB/s	       0 B/op	       0 allocs/op
BenchmarkConnectionProfile/valid-true-4          	300000000	         5.99 ns/op	 166.94 MB/s	       0 B/op	       0 allocs/op
BenchmarkConnectionProfile/get-secret-4          	2000000000	         0.34 ns/op	2903.84 MB/s	       0 B/op	       0 allocs/op
BenchmarkConnectionProfile/create-token-4        	  200000	      8784 ns/op	   0.11 MB/s	    4033 B/op	      36 allocs/op
BenchmarkConnectionProfile/parse-token-4         	  100000	     21726 ns/op	   0.05 MB/s	    6398 B/op	     111 allocs/op

// after optimizing package to use fastime instead of time:

BenchmarkConnectionProfile/instantiate-empty-profile-4         	2000000000	         0.33 ns/op	3072.33 MB/s	       0 B/op	       0 allocs/op
BenchmarkConnectionProfile/create-profile-4                    	 2000000	       901 ns/op	   1.11 MB/s	      64 B/op	       2 allocs/op
BenchmarkConnectionProfile/valid-false-4                       	300000000	         6.38 ns/op	 156.66 MB/s	       0 B/op	       0 allocs/op
BenchmarkConnectionProfile/valid-true-4                        	200000000	         6.63 ns/op	 150.89 MB/s	       0 B/op	       0 allocs/op
BenchmarkConnectionProfile/get-secret-4                        	2000000000	         0.33 ns/op	3068.21 MB/s	       0 B/op	       0 allocs/op
BenchmarkConnectionProfile/create-token-4                      	  200000	      7956 ns/op	   0.13 MB/s	    3553 B/op	      34 allocs/op
BenchmarkConnectionProfile/parse-token-4                       	  100000	     20064 ns/op	   0.05 MB/s	    6350 B/op	     107 allocs/op
*/
