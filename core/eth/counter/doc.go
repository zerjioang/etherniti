// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package counter

/*

go test -bench=. -benchmem -benchtime=2s -cpu=1,2,4

this package implements an concurrency safe atomic uint32 data structure

Initial performance:

BenchmarkCounterPtr/instantiate                 3000000000               0.31 ns/op     3225.37 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/instantiate-2               3000000000               0.31 ns/op     3235.25 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/instantiate-4               3000000000               0.31 ns/op     3180.21 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/add                         500000000                6.03 ns/op      165.73 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/add-2                       500000000                5.91 ns/op      169.25 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/add-4                       500000000                5.94 ns/op      168.25 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/get                         3000000000               0.33 ns/op     2998.91 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/get-2                       3000000000               0.34 ns/op     2969.39 MB/s           0 B/op          0 allocs/op
BenchmarkCounterPtr/get-4                       3000000000               0.34 ns/op     2950.34 MB/s           0 B/op          0 allocs/op

As can be seen non pointer based implementation is much slower ,and thus, is use is not recommended
*/
