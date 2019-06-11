package auth

/*

Initial package performance:

BenchmarkResponse/instantiate-12         	2000000000	         0.33 ns/op	3048.56 MB/s	       0 B/op	       0 allocs/op
BenchmarkResponse/json-12                	 5000000	       355 ns/op	   2.81 MB/s	      72 B/op	       4 allocs/op

After replacing json encoder by custom implementation

BenchmarkResponse/instantiate-12         	2000000000	         0.26 ns/op	3892.86 MB/s	       0 B/op	       0 allocs/op
BenchmarkResponse/json-12                	50000000	        26.1 ns/op	  38.27 MB/s	       0 B/op	       0 allocs/op

after using unsafe to convert string to bytes

BenchmarkResponse/json-12                     5000000	         355 ns/op	   2.81 MB/s	      72 B/op	       4 allocs/op
BenchmarkResponse/json-12                	100000000	        21.5 ns/op	  46.52 MB/s	       0 B/op	       0 allocs/op

we also add a method for writing directly to a writer interface

BenchmarkResponse/instantiate-12         	2000000000	         0.26 ns/op	3802.33 MB/s	       0 B/op	       0 allocs/op
BenchmarkResponse/json-12                	100000000	        21.3 ns/op	  47.02 MB/s	       0 B/op	       0 allocs/op
BenchmarkResponse/writer/nil-12          	2000000000	         1.90 ns/op	 525.32 MB/s	       0 B/op	       0 allocs/op
BenchmarkResponse/writer/bytes-buffer-12 	100000000	        21.9 ns/op	  45.70 MB/s	       0 B/op	       0 allocs/op

In conclusion, we updated from 5000000 ops to 100000000 ops, having a 20x speedup
*/
