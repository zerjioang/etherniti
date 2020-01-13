// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package cmd

/*
Initial package performance:

go test -pibench=. -benchmem -benchtime=5s -cpu=1,2,4

go test -pibench=. -benchmem -cpu=1,2,4
go test -pibench=. -benchmem -benchtime=5s -cpu=1,2,4 -memprofile memprofile.out -cpuprofile profile.out
go tool -web pprof profile.out
go tool -web pprof memprofile.out
go tool pprof -http=localhost:6060 memprofile.out

BenchmarkCmd/run-server-goroutines                 50000            168386 ns/op           0.01 MB/s       10208 B/op        103 allocs/op
BenchmarkCmd/run-server-goroutines-2               20000            327142 ns/op           0.00 MB/s       10208 B/op        103 allocs/op
BenchmarkCmd/run-server-goroutines-4               20000            392836 ns/op           0.00 MB/s       10209 B/op        103 allocs/op

*/
