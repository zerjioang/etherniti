// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package config

/*

Initial package performance:

go test -bench=. -benchmem -cpu=1,2,4
go test -bench=. -benchmem -benchtime=5s -cpu=1,2,4 -memprofile memprofile.out -cpuprofile profile.out
go tool -web pprof profile.out
go tool -web pprof memprofile.out
go tool pprof -http=localhost:6060 memprofile.out

for production deployments, an SSL certificate information is required.
in order to have a modular an extensible design, this information is provided via env variables as following:

* X_ETHERNITI_SSL_CERT_FILE: /path/to/cert/file.pem
* X_ETHERNITI_SSL_KEY_FILE: /path/to/cert/key.pem


BenchmarkCommon/BlockTorConnections-4         	 1000000	      1014 ns/op	   0.99 MB/s	     184 B/op	       2 allocs/op
*/
