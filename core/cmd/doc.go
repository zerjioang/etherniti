// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package cmd

/*
go test -bench=. -benchmem -cpu=1,2,4
go test -bench=. -benchmem -benchtime=5s -cpu=1,2,4 -memprofile memprofile.out -cpuprofile profile.out
go tool -web pprof profile.out
go tool -web pprof memprofile.out
go tool pprof -http=localhost:6060 memprofile.out
*/
