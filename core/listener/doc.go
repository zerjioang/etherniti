// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package listener

/*

initial package benchmarking: go test -pibench=. -benchmem -benchtime=2s -cpu=1,2,4

BenchmarkFactoryListener/factory-http                   2000000000               2.72 ns/op     2942.84 MB/s           0 B/op          0 allocs/op
BenchmarkFactoryListener/factory-http-2                 2000000000               2.69 ns/op     2976.69 MB/s           0 B/op          0 allocs/op
BenchmarkFactoryListener/factory-http-4                 1000000000               3.41 ns/op     2346.31 MB/s           0 B/op          0 allocs/op
BenchmarkFactoryListener/factory-https                  1000000000               3.33 ns/op     2401.81 MB/s           0 B/op          0 allocs/op
BenchmarkFactoryListener/factory-https-2                2000000000               3.34 ns/op     2391.96 MB/s           0 B/op          0 allocs/op
BenchmarkFactoryListener/factory-https-4                1000000000               3.43 ns/op     2331.65 MB/s           0 B/op          0 allocs/op
BenchmarkFactoryListener/factory-unix                   50000000                57.8 ns/op       553.54 MB/s          32 B/op          1 allocs/op
BenchmarkFactoryListener/factory-unix-2                 50000000                60.1 ns/op       532.89 MB/s          32 B/op          1 allocs/op
BenchmarkFactoryListener/factory-unix-4                 50000000                58.6 ns/op       546.16 MB/s          32 B/op          1 allocs/op
BenchmarkFactoryListener/factory-other                  50000000                55.9 ns/op       572.84 MB/s          32 B/op          1 allocs/op
BenchmarkFactoryListener/factory-other-2                50000000                55.5 ns/op       576.47 MB/s          32 B/op          1 allocs/op
BenchmarkFactoryListener/factory-other-4                50000000                54.6 ns/op       585.69 MB/s          32 B/op          1 allocs/op

*/
