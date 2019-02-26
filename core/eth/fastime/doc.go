// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package fastime

/*
Initial package performance:

BenchmarkFastTime/fastime-now-4         			30000000	        53.7 ns/op	  18.63 MB/s	       0 B/op	       0 allocs/op
BenchmarkFastTime/fastime-now-unix-4    			30000000	        54.3 ns/op	  18.40 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardTime/standard-now-4    			30000000	        49.8 ns/op	  20.07 MB/s	       0 B/op	       0 allocs/op
BenchmarkStandardTime/standard-now-unix-4         	30000000	        49.3 ns/op	  20.29 MB/s	       0 B/op	       0 allocs/op

## Package scape analysis results

```bash
/usr/local/go/bin/go \
	test -c -gcflags '-m -m -l' \
	-o /tmp/___fastime_test_go github.com/zerjioang/etherniti/core/eth/fastime
```

```
github.com/zerjioang/etherniti/core/eth/fastime
# github.com/zerjioang/etherniti/core/eth/fastime [github.com/zerjioang/etherniti/core/eth/fastime.test]
core/eth/fastime/fastime_bench_test.go:8:24: leaking param: b
core/eth/fastime/fastime_bench_test.go:8:24:    from b (passed to call[argument escapes]) at core/eth/fastime/fastime_bench_test.go:10:7
core/eth/fastime/fastime_bench_test.go:10:23: func literal escapes to heap
core/eth/fastime/fastime_bench_test.go:10:23:   from .sink (assigned to top level variable) at <unknown line number>
core/eth/fastime/fastime_bench_test.go:10:23: func literal escapes to heap
core/eth/fastime/fastime_bench_test.go:10:23:   from &(func literal) (address-of) at core/eth/fastime/fastime_bench_test.go:10:23
core/eth/fastime/fastime_bench_test.go:10:23:   from .sink (assigned to top level variable) at <unknown line number>
core/eth/fastime/fastime_bench_test.go:17:28: func literal escapes to heap
core/eth/fastime/fastime_bench_test.go:17:28:   from .sink (assigned to top level variable) at <unknown line number>
core/eth/fastime/fastime_bench_test.go:17:28: func literal escapes to heap
core/eth/fastime/fastime_bench_test.go:17:28:   from &(func literal) (address-of) at core/eth/fastime/fastime_bench_test.go:17:28
core/eth/fastime/fastime_bench_test.go:17:28:   from .sink (assigned to top level variable) at <unknown line number>
core/eth/fastime/fastime_bench_test.go:10:28: BenchmarkFastTime.func1 b does not escape
core/eth/fastime/fastime_bench_test.go:17:33: BenchmarkFastTime.func2 b does not escape
core/eth/fastime/fastime_bench_test.go:26:28: leaking param: b
core/eth/fastime/fastime_bench_test.go:26:28:   from b (passed to call[argument escapes]) at core/eth/fastime/fastime_bench_test.go:28:7
core/eth/fastime/fastime_bench_test.go:28:24: func literal escapes to heap
core/eth/fastime/fastime_bench_test.go:28:24:   from .sink (assigned to top level variable) at <unknown line number>
core/eth/fastime/fastime_bench_test.go:28:24: func literal escapes to heap
core/eth/fastime/fastime_bench_test.go:28:24:   from &(func literal) (address-of) at core/eth/fastime/fastime_bench_test.go:28:24
core/eth/fastime/fastime_bench_test.go:28:24:   from .sink (assigned to top level variable) at <unknown line number>
core/eth/fastime/fastime_bench_test.go:35:29: func literal escapes to heap
core/eth/fastime/fastime_bench_test.go:35:29:   from .sink (assigned to top level variable) at <unknown line number>
core/eth/fastime/fastime_bench_test.go:35:29: func literal escapes to heap
core/eth/fastime/fastime_bench_test.go:35:29:   from &(func literal) (address-of) at core/eth/fastime/fastime_bench_test.go:35:29
core/eth/fastime/fastime_bench_test.go:35:29:   from .sink (assigned to top level variable) at <unknown line number>
core/eth/fastime/fastime_bench_test.go:28:29: BenchmarkStandardTime.func1 b does not escape
core/eth/fastime/fastime_bench_test.go:35:34: BenchmarkStandardTime.func2 b does not escape
core/eth/fastime/fastime_test.go:8:19: leaking param: t
core/eth/fastime/fastime_test.go:8:19:  from t (passed to call[argument escapes]) at core/eth/fastime/fastime_test.go:10:7
core/eth/fastime/fastime_test.go:10:23: func literal escapes to heap
core/eth/fastime/fastime_test.go:10:23:         from .sink (assigned to top level variable) at <unknown line number>
core/eth/fastime/fastime_test.go:10:23: func literal escapes to heap
core/eth/fastime/fastime_test.go:10:23:         from &(func literal) (address-of) at core/eth/fastime/fastime_test.go:10:23
core/eth/fastime/fastime_test.go:10:23:         from .sink (assigned to top level variable) at <unknown line number>
core/eth/fastime/fastime_test.go:16:28: func literal escapes to heap
core/eth/fastime/fastime_test.go:16:28:         from .sink (assigned to top level variable) at <unknown line number>
core/eth/fastime/fastime_test.go:16:28: func literal escapes to heap
core/eth/fastime/fastime_test.go:16:28:         from &(func literal) (address-of) at core/eth/fastime/fastime_test.go:16:28
core/eth/fastime/fastime_test.go:16:28:         from .sink (assigned to top level variable) at <unknown line number>
core/eth/fastime/fastime_test.go:10:28: TestFastTime.func1 t does not escape
core/eth/fastime/fastime_test.go:16:33: TestFastTime.func2 t does not escape
core/eth/fastime/fastime_test.go:25:23: leaking param: t
core/eth/fastime/fastime_test.go:25:23:         from t (passed to call[argument escapes]) at core/eth/fastime/fastime_test.go:27:7
core/eth/fastime/fastime_test.go:27:24: func literal escapes to heap
core/eth/fastime/fastime_test.go:27:24:         from .sink (assigned to top level variable) at <unknown line number>
core/eth/fastime/fastime_test.go:27:24: func literal escapes to heap
core/eth/fastime/fastime_test.go:27:24:         from &(func literal) (address-of) at core/eth/fastime/fastime_test.go:27:24
core/eth/fastime/fastime_test.go:27:24:         from .sink (assigned to top level variable) at <unknown line number>
core/eth/fastime/fastime_test.go:31:29: func literal escapes to heap
core/eth/fastime/fastime_test.go:31:29:         from .sink (assigned to top level variable) at <unknown line number>
core/eth/fastime/fastime_test.go:31:29: func literal escapes to heap
core/eth/fastime/fastime_test.go:31:29:         from &(func literal) (address-of) at core/eth/fastime/fastime_test.go:31:29
core/eth/fastime/fastime_test.go:31:29:         from .sink (assigned to top level variable) at <unknown line number>
core/eth/fastime/fastime_test.go:29:4: t.common escapes to heap
core/eth/fastime/fastime_test.go:29:4:  from t.common (passed to call[argument escapes]) at core/eth/fastime/fastime_test.go:29:8
core/eth/fastime/fastime_test.go:27:29: leaking param: t
core/eth/fastime/fastime_test.go:27:29:         from t.common (dot of pointer) at core/eth/fastime/fastime_test.go:29:4
core/eth/fastime/fastime_test.go:27:29:         from t.common (address-of) at core/eth/fastime/fastime_test.go:29:4
core/eth/fastime/fastime_test.go:27:29:         from t.common (passed to call[argument escapes]) at core/eth/fastime/fastime_test.go:29:8
core/eth/fastime/fastime_test.go:29:8: tm3 escapes to heap
core/eth/fastime/fastime_test.go:29:8:  from ... argument (arg to ...) at core/eth/fastime/fastime_test.go:29:8
core/eth/fastime/fastime_test.go:29:8:  from *(... argument) (indirection) at core/eth/fastime/fastime_test.go:29:8
core/eth/fastime/fastime_test.go:29:8:  from ... argument (passed to call[argument content escapes]) at core/eth/fastime/fastime_test.go:29:8
core/eth/fastime/fastime_test.go:34:4: t.common escapes to heap
core/eth/fastime/fastime_test.go:34:4:  from t.common (passed to call[argument escapes]) at core/eth/fastime/fastime_test.go:34:8
core/eth/fastime/fastime_test.go:31:34: leaking param: t
core/eth/fastime/fastime_test.go:31:34:         from t.common (dot of pointer) at core/eth/fastime/fastime_test.go:34:4
core/eth/fastime/fastime_test.go:31:34:         from t.common (address-of) at core/eth/fastime/fastime_test.go:34:4
core/eth/fastime/fastime_test.go:31:34:         from t.common (passed to call[argument escapes]) at core/eth/fastime/fastime_test.go:34:8
core/eth/fastime/fastime_test.go:34:8: u escapes to heap
core/eth/fastime/fastime_test.go:34:8:  from ... argument (arg to ...) at core/eth/fastime/fastime_test.go:34:8
core/eth/fastime/fastime_test.go:34:8:  from *(... argument) (indirection) at core/eth/fastime/fastime_test.go:34:8
core/eth/fastime/fastime_test.go:34:8:  from ... argument (passed to call[argument content escapes]) at core/eth/fastime/fastime_test.go:34:8
core/eth/fastime/fastime_test.go:29:8: TestStandardTime.func1 ... argument does not escape
core/eth/fastime/fastime_test.go:34:8: TestStandardTime.func2 ... argument does not escape
<autogenerated>:1: (*Duration).Nanoseconds .this does not escape
<autogenerated>:1: (*FastTime).Add .this does not escape
<autogenerated>:1: (*FastTime).Unix .this does not escape
# github.com/zerjioang/etherniti/core/eth/fastime.test
/tmp/go-build640814093/b001/_testmain.go:48:42: testdeps.TestDeps literal escapes to heap
/tmp/go-build640814093/b001/_testmain.go:48:42:         from testdeps.TestDeps literal (passed to call[argument escapes]) at $WORK/b001/_testmain.go:48:24
```

As can be seen our fastime variables tm1 and tm2 are not scaped to Heap.
*/
