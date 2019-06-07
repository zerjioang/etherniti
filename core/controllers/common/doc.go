package common

/*
initial package performance

different data appending methods performance:

BenchmarkDatabaseController/append-12         	30000000	        77.5 ns/op	  12.90 MB/s	      80 B/op	       1 allocs/op
BenchmarkDatabaseController/append-2-12       	20000000	       109 ns/op	   9.17 MB/s	     160 B/op	       2 allocs/op
BenchmarkDatabaseController/append-3-12       	10000000	       132 ns/op	   7.52 MB/s	     224 B/op	       3 allocs/op
BenchmarkDatabaseController/append-4-12       	20000000	        73.5 ns/op	  13.61 MB/s	      96 B/op	       2 allocs/op
BenchmarkDatabaseController/append-5-12       	300000000	         4.57 ns/op	 219.03 MB/s	       0 B/op	       0 allocs/op
BenchmarkDatabaseController/append-6-12       	30000000	        77.7 ns/op	  12.87 MB/s	      80 B/op	       1 allocs/op

*/
