package handlers

/*
initial package performance


Index controller

BenchmarkIndex/instantiation-4         	2000000000	         0.35 ns/op	2891.74 MB/s	       0 B/op	       0 allocs/op
BenchmarkIndex/status-4                	  200000	      9157 ns/op	   0.11 MB/s	    2144 B/op	      17 allocs/op
BenchmarkIndex/integrity-4             	    2000	    937990 ns/op	   0.00 MB/s	    6510 B/op	      97 allocs/op

After using specific model instead of maps

BenchmarkIndexMethods/instantiation-4         	2000000000	         0.36 ns/op	2796.09 MB/s	       0 B/op	       0 allocs/op
BenchmarkIndexMethods/status-4                	  200000	      7278 ns/op	   0.14 MB/s	     144 B/op	       3 allocs/op
BenchmarkIndexMethods/integrity-4             	    2000	    981535 ns/op	   0.00 MB/s	    6128 B/op	      92 allocs/op

After using sync.pools and optimizing with pproof status method we get a boost of x5

BenchmarkIndexMethods/status-4                	 1000000	      1271 ns/op	   0.79 MB/s	     601 B/op	       8 allocs/op


 */
