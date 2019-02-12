package profile

/*
initial package performance

BenchmarkConnectionProfile/instantiate-4         	2000000000	         0.34 ns/op	2937.51 MB/s	       0 B/op	       0 allocs/op
BenchmarkConnectionProfile/valid-false-4         	300000000	         5.44 ns/op	 183.84 MB/s	       0 B/op	       0 allocs/op
BenchmarkConnectionProfile/valid-true-4          	300000000	         5.99 ns/op	 166.94 MB/s	       0 B/op	       0 allocs/op
BenchmarkConnectionProfile/get-secret-4          	2000000000	         0.34 ns/op	2903.84 MB/s	       0 B/op	       0 allocs/op
BenchmarkConnectionProfile/create-token-4        	  200000	      8784 ns/op	   0.11 MB/s	    4033 B/op	      36 allocs/op
BenchmarkConnectionProfile/parse-token-4         	  100000	     21726 ns/op	   0.05 MB/s	    6398 B/op	     111 allocs/op

*/
