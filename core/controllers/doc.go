// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package controllers

/*
initial package performance

Index controller

BenchmarkIndex/instantiation-4         	2000000000	         0.35 ns/op	2891.74 MB/s	       0 B/op	       0 allocs/op
BenchmarkIndex/metrics-4                	200000	      9157 ns/op	   0.11 MB/s	    2144 B/op	      17 allocs/op
BenchmarkIndex/integrity-4             	      2000	    937990 ns/op	   0.00 MB/s	    6510 B/op	      97 allocs/op

After using specific model instead of maps

BenchmarkIndexMethods/instantiation-4         	2000000000	      0.36 ns/op	2796.09 MB/s	       0 B/op	       0 allocs/op
BenchmarkIndexMethods/metrics-4                	  200000	      7278 ns/op	   0.14 MB/s	     144 B/op	       3 allocs/op
BenchmarkIndexMethods/integrity-4             	    2000	    981535 ns/op	   0.00 MB/s	    6128 B/op	      92 allocs/op

After using sync.pools and optimizing with pproof status method we get a boost of x5

BenchmarkIndexMethods/instantiation-4         	    2000	    975627 ns/op	   0.00 MB/s	    7810 B/op	     107 allocs/op
BenchmarkIndexMethods/metrics-4                	10000000	       189 ns/op	   5.28 MB/s	      64 B/op	       1 allocs/op
BenchmarkIndexMethods/metrics-reload-4         	 1000000	      1033 ns/op	   0.97 MB/s	     164 B/op	       7 allocs/op
BenchmarkIndexMethods/integrity-4             	10000000	       174 ns/op	   5.73 MB/s	      64 B/op	       1 allocs/op
BenchmarkIndexMethods/integrity-reload-4      	    2000	    932014 ns/op	   0.00 MB/s	    6135 B/op	      92 allocs/op

finally, removing all references to defer and using mutexes as pointer we get:

BenchmarkIndexMethods/instantiation-4         	     1000	   1001855 ns/op	   0.00 MB/s	    7655 B/op	     105 allocs/op
BenchmarkIndexMethods/metrics-4                   50000000	      20.2 ns/op	  49.53 MB/s	       0 B/op	       0 allocs/op
BenchmarkIndexMethods/metrics-reload-4         	  2000000	       980 ns/op	   1.02 MB/s	     158 B/op	       6 allocs/op
BenchmarkIndexMethods/integrity-4             	100000000	      19.8 ns/op	  50.51 MB/s	       0 B/op	       0 allocs/op
BenchmarkIndexMethods/integrity-reload-4      	     2000	    931280 ns/op	   0.00 MB/s	    6017 B/op	      88 allocs/op

BenchmarkIndexMethods/instantiation-4         	     2000	     1009697 ns/op	   0.00 MB/s	    7536 B/op	     102 allocs/op
BenchmarkIndexMethods/metrics-4                	100000000	        19.8 ns/op	  50.58 MB/s	       0 B/op	       0 allocs/op
BenchmarkIndexMethods/metrics-reload-4            3000000	         902 ns/op	   1.11 MB/s	      50 B/op	       8 allocs/op
BenchmarkIndexMethods/integrity-4             	100000000	        21.1 ns/op	  47.39 MB/s	       0 B/op	       0 allocs/op
BenchmarkIndexMethods/integrity-reload-4      	     2000	      954523 ns/op	   0.00 MB/s	    6016 B/op	      88 allocs/op

We achieve our goal, have a high performance status and integrity calls, that will only send new information on each updated round.
*/
