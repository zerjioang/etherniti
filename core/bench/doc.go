// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0
package bench

/*
go test -tags "dev oss" -bench=. -benchmem -cpu=1,2,4
go test -tags "dev oss" -bench=. -benchmem -benchtime=5s -cpu=1,2,4 -memprofile memprofile.out -cpuprofile profile.out
go tool -web pprof profile.out
go tool -web pprof memprofile.out
go tool pprof -http=localhost:6060 memprofile.out

package functions performance:

BenchmarkPi/calculate-score-12                  26	      48355870 ns/op	   0.00 MB/s	   67922 B/op	      28 allocs/op
BenchmarkPi/get-score-12               	1000000000	         0.251 ns/op	3985.16 MB/s	       0 B/op	       0 allocs/op
BenchmarkPi/get-bench-time-12          	1000000000	         0.255 ns/op	3924.03 MB/s	       0 B/op	       0 allocs/op

after caching calculate-score results, since we only execute once at bootime, we get
*/
