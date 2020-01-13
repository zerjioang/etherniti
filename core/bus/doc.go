package bus

/*

Initial package performance:

go test -pibench=. -benchmem -benchtime=5s -cpu=1,2,4

BenchmarkBus/get-bus            10000000000              0.33 ns/op     3052.68 MB/s           0 B/op          0 allocs/op
BenchmarkBus/get-bus-2          10000000000              0.33 ns/op     3064.29 MB/s           0 B/op          0 allocs/op
BenchmarkBus/get-bus-4          10000000000              0.34 ns/op     2925.96 MB/s           0 B/op          0 allocs/op

*/
