package bip39

/*
initial package performance

BenchmarkBip39/bip39-generate-4         	  200000	     10377 ns/op	   0.10 MB/s	    2512 B/op	      62 allocs/op

after some minor changes

BenchmarkBip39/bip39-generate-4         	  200000	      7832 ns/op	   0.13 MB/s	    2544 B/op	      63 allocs/op
BenchmarkBip39/is-valid-4   		      	  500000	      3797 ns/op	   0.26 MB/s	     384 B/op	       1 allocs/op


*/
