package project

/*
Initial package performance:

go test -bench=. -benchmem -benchtime=5s -cpu=1,2,4

BenchmarkProjectController/create-controller              200000             45512 ns/op           0.02 MB/s       14124 B/op         59 allocs/op
BenchmarkProjectController/create-controller-2            200000             47167 ns/op           0.02 MB/s       14125 B/op         59 allocs/op
BenchmarkProjectController/create-controller-4            200000             45893 ns/op           0.02 MB/s       14132 B/op         59 allocs/op

*/
