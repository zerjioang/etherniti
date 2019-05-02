// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package e2e

/*
go test -bench=. -race -benchmem -benchtime=2s -cpu=1,2,4 -memprofile memprofile.out -cpuprofile profile.out
go tool pprof -http=localhost:6060 profile.out

initial performance benchmark

BenchmarkWalletController/new-entropy                    1000000              2047 ns/op           0.49 MB/s         367 B/op          6 allocs/op
BenchmarkWalletController/new-entropy-2                  2000000              1856 ns/op           0.54 MB/s         367 B/op          6 allocs/op
BenchmarkWalletController/new-entropy-4                  2000000              1795 ns/op           0.56 MB/s         367 B/op          6 allocs/op

*/
