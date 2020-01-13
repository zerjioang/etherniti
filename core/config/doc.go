// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package config

/*

Initial package performance:

go test -pibench=. -benchmem -cpu=1,2,4
go test -pibench=. -benchmem -benchtime=5s -cpu=1,2,4 -memprofile memprofile.out -cpuprofile profile.out
go tool -web pprof profile.out
go tool -web pprof memprofile.out
go tool pprof -http=localhost:6060 memprofile.out

for production deployments, an SSL certificate information is required.
in order to have a modular an extensible design, this information is provided via env variables as following:

* X_ETHERNITI_SSL_CERT_FILE: /path/to/cert/file.pem
* X_ETHERNITI_SSL_KEY_FILE: /path/to/cert/key.pem


BenchmarkGetEnvironment/get-env-4         	 1000000	      1393 ns/op	   0.72 MB/s	       0 B/op	       0 allocs/op
BenchmarkGetEnvironment/get-env-ptr-4     	 1000000	      1264 ns/op	   0.79 MB/s	       0 B/op	       0 allocs/op
BenchmarkGetEnvironment/get-env-parallel-4         	 1000000	      1757 ns/op	   0.57 MB/s	       0 B/op	       0 allocs/op
BenchmarkGetEnvironment/get-env-ptr-parallel-4     	 1000000	      1703 ns/op	   0.59 MB/s	       0 B/op	       0 allocs/op
BenchmarkGetEnvironment/read-key-env-4             	 1000000	      1233 ns/op	   0.81 MB/s	       0 B/op	       0 allocs/op
BenchmarkGetEnvironment/read-key-ptr-4             	 1000000	      1221 ns/op	   0.82 MB/s	       0 B/op	       0 allocs/op

Optimizations: 1 - use func() init to load environment data only once (at proxy boot time). func init ensured thread safety so no problem
Speedup: x500 -> from 1000000 to 500000000. and 0 allocs

BenchmarkCommon/BlockTorConnections             10000000000              0.41 ns/op     2464.90 MB/s           0 B/op          0 allocs/op
BenchmarkCommon/BlockTorConnections-2           10000000000              0.40 ns/op     2486.18 MB/s           0 B/op          0 allocs/op
BenchmarkCommon/BlockTorConnections-4           10000000000              0.40 ns/op     2475.94 MB/s           0 B/op          0 allocs/op
BenchmarkGetRedirectUrl/redirect                 5000000              1426 ns/op           0.70 MB/s          48 B/op          1 allocs/op
BenchmarkGetRedirectUrl/redirect-2               5000000              1356 ns/op           0.74 MB/s          48 B/op          1 allocs/op
BenchmarkGetRedirectUrl/redirect-4               5000000              1399 ns/op           0.71 MB/s          48 B/op          1 allocs/op
BenchmarkGetRedirectUrl/cert-pem                10000000              1186 ns/op           0.84 MB/s           0 B/op          0 allocs/op
BenchmarkGetRedirectUrl/cert-pem-2              10000000              1241 ns/op           0.81 MB/s           0 B/op          0 allocs/op
BenchmarkGetRedirectUrl/cert-pem-4              10000000              1360 ns/op           0.74 MB/s           0 B/op          0 allocs/op
BenchmarkGetRedirectUrl/key-pem                  5000000              1197 ns/op           0.83 MB/s           0 B/op          0 allocs/op
BenchmarkGetRedirectUrl/key-pem-2               10000000              1182 ns/op           0.85 MB/s           0 B/op          0 allocs/op
BenchmarkGetRedirectUrl/key-pem-4                5000000              1571 ns/op           0.64 MB/s           0 B/op          0 allocs/op
BenchmarkGetEnvironment/get-env                 10000000000              0.45 ns/op     2201.39 MB/s           0 B/op          0 allocs/op
BenchmarkGetEnvironment/get-env-2               10000000000              0.43 ns/op     2333.64 MB/s           0 B/op          0 allocs/op
BenchmarkGetEnvironment/get-env-4               10000000000              0.41 ns/op     2452.59 MB/s           0 B/op          0 allocs/op
BenchmarkGetEnvironment/get-env-parallel        3000000000               2.65 ns/op      376.96 MB/s           0 B/op          0 allocs/op
BenchmarkGetEnvironment/get-env-parallel-2      10000000000              1.36 ns/op      732.70 MB/s           0 B/op          0 allocs/op
BenchmarkGetEnvironment/get-env-parallel-4      10000000000              0.86 ns/op     1166.32 MB/s           0 B/op          0 allocs/op
BenchmarkGetEnvironment/read-key-env            500000000               19.3 ns/op        51.87 MB/s           0 B/op          0 allocs/op
BenchmarkGetEnvironment/read-key-env-2          300000000               19.9 ns/op        50.28 MB/s           0 B/op          0 allocs/op
BenchmarkGetEnvironment/read-key-env-4          500000000               19.7 ns/op        50.70 MB/s           0 B/op          0 allocs/op
BenchmarkGetEnvironment/read-key-ptr            500000000               19.4 ns/op        51.53 MB/s           0 B/op          0 allocs/op
BenchmarkGetEnvironment/read-key-ptr-2          500000000               19.2 ns/op        52.21 MB/s           0 B/op          0 allocs/op
BenchmarkGetEnvironment/read-key-ptr-4          500000000               19.0 ns/op        52.52 MB/s           0 B/op          0 allocs/op

*/
