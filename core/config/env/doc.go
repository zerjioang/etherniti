package env

/*
initial package performance

BenchmarkEnvironment/get-env-12         	           2000000	       675 ns/op	   1.48 MB/s	     240 B/op	       4 allocs/op
BenchmarkEnvironment/get-env-parallel-12         	   2000000	       737 ns/op	   1.36 MB/s	     240 B/op	       4 allocs/op
BenchmarkEnvironment/read-key-all-12             	     20000	     75598 ns/op	   0.01 MB/s	   16625 B/op	     353 allocs/op
BenchmarkEnvironment/read-key-env-12             	 100000000	      14.7 ns/op	  68.15 MB/s	       0 B/op	       0 allocs/op
BenchmarkEnvironment/read-key-env-parallel-12    	2000000000	      0.44 ns/op	2277.47 MB/s	       0 B/op	       0 allocs/op

After using an internal variable to check whether env variables were readed or not

BenchmarkEnvironment/get-env-12         	           2000000	         723 ns/op	   1.38 MB/s	     264 B/op	       4 allocs/op
BenchmarkEnvironment/get-env-parallel-12         	   2000000	         775 ns/op	   1.29 MB/s	     264 B/op	       4 allocs/op
BenchmarkEnvironment/read-key-all-12             	 100000000	        13.7 ns/op	  73.22 MB/s	       0 B/op	       0 allocs/op
BenchmarkEnvironment/read-key-all-parallel-12    	  50000000	        25.6 ns/op	  39.13 MB/s	       0 B/op	       0 allocs/op
BenchmarkEnvironment/read-key-env-12             	 100000000	        15.2 ns/op	  65.80 MB/s	       0 B/op	       0 allocs/op
BenchmarkEnvironment/read-key-env-parallel-12    	2000000000	        0.43 ns/op	2349.82 MB/s	       0 B/op	       0 allocs/op

Minor tweaking

BenchmarkEnvironment/new-env-12      		    	  2000000	         722 ns/op	   1.38 MB/s	     264 B/op	       4 allocs/op
BenchmarkEnvironment/new-env-parallel-12         	  2000000	        1084 ns/op	   0.92 MB/s	     264 B/op	       4 allocs/op
BenchmarkEnvironment/read-key-all-12             	100000000	        13.8 ns/op	  72.59 MB/s	       0 B/op	       0 allocs/op
BenchmarkEnvironment/read-key-all-parallel-12    	 50000000	        26.6 ns/op	  37.65 MB/s	       0 B/op	       0 allocs/op
BenchmarkEnvironment/read-key-env-12             	100000000	        14.7 ns/op	  68.25 MB/s	       0 B/op	       0 allocs/op
BenchmarkEnvironment/read-key-env-parallel-12      2000000000	        0.83 ns/op	1199.17 MB/s	       0 B/op	       0 allocs/op

*/
