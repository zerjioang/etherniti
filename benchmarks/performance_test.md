# Performance test

## Test: Get Account Balance via nodejs web3

```bash
Iterations  1 { totalMilliseconds: 26.992991999723017,  averageMillisecondsPerTask: 26.96292999992147 }
Iterations  5 { totalMilliseconds: 9.277198999654502,  averageMillisecondsPerTask: 1.8478354000486434 }
Iterations  10 { totalMilliseconds: 24.057543999981135,  averageMillisecondsPerTask: 2.399410399980843 }
Iterations  50 { totalMilliseconds: 69.18710500001907,  averageMillisecondsPerTask: 1.377910760026425 }
Iterations  100 { totalMilliseconds: 129.0320410002023,  averageMillisecondsPerTask: 1.2843921100348235 }
Iterations  500 { totalMilliseconds: 452.7258150000125,  averageMillisecondsPerTask: 0.9004029679885134 }
Iterations  1000 { totalMilliseconds: 947.9371759998612,  averageMillisecondsPerTask: 0.9418804899957031 }
Iterations  2000 { totalMilliseconds: 1504.9332340001129,  averageMillisecondsPerTask: 0.7482629249994643 }
```

## Result returned by getBalance()

```bash
Result: Ether: 99.98695186
```

## Test: Get Account Balance via Etherniti Proxy

```bash
ab -n 10 -c 10 -H "X-Etherniti-Profile:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbmRwb2ludCI6IkhUVFA6Ly8xMjcuMC4wLjE6NzU0NSIsImFkZHJlc3MiOiIweDNERUIxODk0REMyZDNlMUI0YTA3M2Y1MjBlNTE2QzJERjZmNDVCODgiLCJzb3VyY2UiOjIxMzA3MDY0MzMsInZlcnNpb24iOiIwLjAuNiIsImV4cCI6MTU1NjUwMjE1MCwianRpIjoiZmQ0NjQzZmItNjk4My00MzI1LWIzNzctMTJmOWRmZDY2M2IxIiwiaWF0IjoxNTU2MTQyMTUwLCJpc3MiOiJldGhlcm5pdGkub3JnIiwibmJmIjoxNTU2MTQyMTUwLCJ2YWxpZGl0eSI6ZmFsc2V9.bwOFdtZBJ6oLhQtNwo_IQTPnOMf2edQGQfDeKQEhNuI" http://127.0.0.1:8080/v1/private/balance/0x3DEB1894DC2d3e1B4a073f520e516C2DF6f45B88
```

### Results with cache disabled and logging debug

```bash
Benchmarking 127.0.0.1 (be patient)
Completed 200 requests
Completed 400 requests
Completed 600 requests
Completed 800 requests
Completed 1000 requests
Completed 1200 requests
Completed 1400 requests
Completed 1600 requests
Completed 1800 requests
Completed 2000 requests
Finished 2000 requests


Server Software:        Apache/2.0.54
Server Hostname:        127.0.0.1
Server Port:            8080

Document Path:          /v1/private/balance/0x3DEB1894DC2d3e1B4a073f520e516C2DF6f45B88
Document Length:        65 bytes

Concurrency Level:      500
Time taken for tests:   6.039 seconds
Complete requests:      2000
Failed requests:        0
Total transferred:      1316000 bytes
HTML transferred:       130000 bytes
Requests per second:    331.18 [#/sec] (mean)
Time per request:       1509.760 [ms] (mean)
Time per request:       3.020 [ms] (mean, across all concurrent requests)
Transfer rate:          212.81 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0  131 337.8      0    1034
Processing:    27 1213 381.7   1369    1715
Waiting:       27 1213 381.7   1369    1715
Total:         43 1344 512.4   1406    2725

Percentage of the requests served within a certain time (ms)
  50%   1406
  66%   1511
  75%   1533
  80%   1555
  90%   2018
  95%   2342
  98%   2601
  99%   2665
 100%   2725 (longest request)
```

### Results with cache enabled and logging debug

```bash
Benchmarking 127.0.0.1 (be patient)
Completed 200 requests
Completed 400 requests
Completed 600 requests
Completed 800 requests
Completed 1000 requests
Completed 1200 requests
Completed 1400 requests
Completed 1600 requests
Completed 1800 requests
Completed 2000 requests
Finished 2000 requests


Server Software:        Apache/2.0.54
Server Hostname:        127.0.0.1
Server Port:            8080

Document Path:          /v1/private/balance/0x3DEB1894DC2d3e1B4a073f520e516C2DF6f45B88
Document Length:        65 bytes

Concurrency Level:      500
Time taken for tests:   0.466 seconds
Complete requests:      2000
Failed requests:        0
Total transferred:      1398000 bytes
HTML transferred:       130000 bytes
Requests per second:    4288.07 [#/sec] (mean)
Time per request:       116.602 [ms] (mean)
Time per request:       0.233 [ms] (mean, across all concurrent requests)
Transfer rate:          2927.11 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    2   3.8      0      14
Processing:     0   38  11.6     39     253
Waiting:        0   38  11.6     39     253
Total:          0   40  11.9     40     253

Percentage of the requests served within a certain time (ms)
  50%     40
  66%     43
  75%     45
  80%     46
  90%     53
  95%     58
  98%     64
  99%     73
 100%    253 (longest request)

```

### Results returned by Etherniti

```bash
curl -H "X-Etherniti-Profile:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbmRwb2ludCI6IkhUVFA6Ly8xMjcuMC4wLjE6NzU0NSIsImFkZHJlc3MiOiIweDNERUIxODk0REMyZDNlMUI0YTA3M2Y1MjBlNTE2QzJERjZmNDVCODgiLCJzb3VyY2UiOjIxMzA3MDY0MzMsInZlcnNpb24iOiIwLjAuNiIsImV4cCI6MTU1NjUwMjE1MCwianRpIjoiZmQ0NjQzZmItNjk4My00MzI1LWIzNzctMTJmOWRmZDY2M2IxIiwiaWF0IjoxNTU2MTQyMTUwLCJpc3MiOiJldGhlcm5pdGkub3JnIiwibmJmIjoxNTU2MTQyMTUwLCJ2YWxpZGl0eSI6ZmFsc2V9.bwOFdtZBJ6oLhQtNwo_IQTPnOMf2edQGQfDeKQEhNuI" http://127.0.0.1:8080/v1/private/balance/0x3DEB1894DC2d3e1B4a073f520e516C2DF6f45B88  
    
{"id":0,"code":200,"msg":"balance","result":99986951860000000000}
```