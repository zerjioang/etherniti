package string

/*
Initial package performance:

go test -bench=. -benchmem -benchtime=5s -cpu=1,2,4

BenchmarkString/create-empty            10000000000              0.31 ns/op     3181.85 MB/s           0 B/op          0 allocs/op
BenchmarkString/create-empty-2          10000000000              0.31 ns/op     3177.60 MB/s           0 B/op          0 allocs/op
BenchmarkString/create-empty-4          10000000000              0.31 ns/op     3222.27 MB/s           0 B/op          0 allocs/op
BenchmarkString/create-with-data        10000000000              0.34 ns/op     2959.60 MB/s           0 B/op          0 allocs/op
BenchmarkString/create-with-data-2      10000000000              0.34 ns/op     2932.14 MB/s           0 B/op          0 allocs/op
BenchmarkString/create-with-data-4      10000000000              0.33 ns/op     2987.18 MB/s           0 B/op          0 allocs/op
BenchmarkString/last-index/standard     500000000               11.7 ns/op        85.28 MB/s           0 B/op          0 allocs/op
BenchmarkString/last-index/standard-2   1000000000              12.1 ns/op        82.44 MB/s           0 B/op          0 allocs/op
BenchmarkString/last-index/standard-4   1000000000              11.6 ns/op        86.06 MB/s           0 B/op          0 allocs/op
BenchmarkString/last-index/custom       500000000               13.2 ns/op        75.83 MB/s           0 B/op          0 allocs/op
BenchmarkString/last-index/custom-2     500000000               12.2 ns/op        81.76 MB/s           0 B/op          0 allocs/op
BenchmarkString/last-index/custom-4     1000000000              12.1 ns/op        82.54 MB/s           0 B/op          0 allocs/op
BenchmarkString/to-bytes/standard       1000000000               7.48 ns/op      133.71 MB/s           0 B/op          0 allocs/op
BenchmarkString/to-bytes/standard-2     1000000000               7.49 ns/op      133.48 MB/s           0 B/op          0 allocs/op
BenchmarkString/to-bytes/standard-4     1000000000               7.73 ns/op      129.30 MB/s           0 B/op          0 allocs/op
BenchmarkString/to-bytes/custom         10000000000              0.49 ns/op     2036.08 MB/s           0 B/op          0 allocs/op
BenchmarkString/to-bytes/custom-2       10000000000              0.49 ns/op     2043.30 MB/s           0 B/op          0 allocs/op
BenchmarkString/to-bytes/custom-4       10000000000              0.48 ns/op     2070.74 MB/s           0 B/op          0 allocs/op
BenchmarkString/chart-at/standard       2000000000               3.59 ns/op      278.44 MB/s           0 B/op          0 allocs/op
BenchmarkString/chart-at/standard-2     2000000000               3.61 ns/op      276.76 MB/s           0 B/op          0 allocs/op
BenchmarkString/chart-at/standard-4     2000000000               3.52 ns/op      284.39 MB/s           0 B/op          0 allocs/op
BenchmarkString/chart-at/custom         10000000000              0.32 ns/op     3151.47 MB/s           0 B/op          0 allocs/op
BenchmarkString/chart-at/custom-2       10000000000              0.32 ns/op     3117.36 MB/s           0 B/op          0 allocs/op
BenchmarkString/chart-at/custom-4       10000000000              0.31 ns/op     3184.27 MB/s           0 B/op          0 allocs/op
BenchmarkString/length/standard         10000000000              0.32 ns/op     3101.67 MB/s           0 B/op          0 allocs/op
BenchmarkString/length/standard-2       10000000000              0.31 ns/op     3199.25 MB/s           0 B/op          0 allocs/op
BenchmarkString/length/standard-4       10000000000              0.31 ns/op     3217.94 MB/s           0 B/op          0 allocs/op
BenchmarkString/length/custom           10000000000              0.35 ns/op     2821.80 MB/s           0 B/op          0 allocs/op
BenchmarkString/length/custom-2         10000000000              0.35 ns/op     2850.31 MB/s           0 B/op          0 allocs/op
BenchmarkString/length/custom-4         10000000000              0.36 ns/op     2809.50 MB/s           0 B/op          0 allocs/op
BenchmarkString/is-empty/standard       10000000000              0.34 ns/op     2953.35 MB/s           0 B/op          0 allocs/op
BenchmarkString/is-empty/standard-2     10000000000              0.34 ns/op     2908.05 MB/s           0 B/op          0 allocs/op
BenchmarkString/is-empty/standard-4     10000000000              0.35 ns/op     2831.15 MB/s           0 B/op          0 allocs/op
BenchmarkString/is-empty/custom         10000000000              0.32 ns/op     3173.14 MB/s           0 B/op          0 allocs/op
BenchmarkString/is-empty/custom-2       10000000000              0.32 ns/op     3096.65 MB/s           0 B/op          0 allocs/op
BenchmarkString/is-empty/custom-4       10000000000              0.31 ns/op     3185.26 MB/s           0 B/op          0 allocs/op
BenchmarkString/to-lowercase/standard                   300000000               29.0 ns/op        34.46 MB/s           0 B/op          0 allocs/op
BenchmarkString/to-lowercase/standard-2                 200000000               30.4 ns/op        32.88 MB/s           0 B/op          0 allocs/op
BenchmarkString/to-lowercase/standard-4                 300000000               29.5 ns/op        33.86 MB/s           0 B/op          0 allocs/op
BenchmarkString/to-lowercase/custom                     500000000               16.4 ns/op        60.98 MB/s           0 B/op          0 allocs/op
BenchmarkString/to-lowercase/custom-2                   500000000               16.1 ns/op        62.19 MB/s           0 B/op          0 allocs/op
BenchmarkString/to-lowercase/custom-4                   500000000               15.5 ns/op        64.47 MB/s           0 B/op          0 allocs/op
BenchmarkString/reverse/custom                          300000000               20.9 ns/op        47.90 MB/s           0 B/op          0 allocs/op
BenchmarkString/reverse/custom-2                        300000000               20.2 ns/op        49.42 MB/s           0 B/op          0 allocs/op
BenchmarkString/reverse/custom-4                        300000000               21.1 ns/op        47.49 MB/s           0 B/op          0 allocs/op
BenchmarkString/title-case/custom                       200000000               32.0 ns/op        31.21 MB/s           0 B/op          0 allocs/op
BenchmarkString/title-case/custom-2                     200000000               31.6 ns/op        31.69 MB/s           0 B/op          0 allocs/op
BenchmarkString/title-case/custom-4                     200000000               33.6 ns/op        29.73 MB/s           0 B/op          0 allocs/op
BenchmarkString/count-byte-match/custom                 300000000               21.5 ns/op        46.45 MB/s           0 B/op          0 allocs/op
BenchmarkString/count-byte-match/custom-2               300000000               20.7 ns/op        48.39 MB/s           0 B/op          0 allocs/op
BenchmarkString/count-byte-match/custom-4               300000000               20.8 ns/op        48.17 MB/s           0 B/op          0 allocs/op

*/
