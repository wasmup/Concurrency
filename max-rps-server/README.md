```sh
go run .

# Requests: 2_208_160, Total Elapsed Time: 3.486776405s

wrk -t4 -c80 -d10s http://localhost:8080/?q=1234567890

# Running 10s test @ http://localhost:8080/?q=1234567890
#   4 threads and 80 connections
#   Thread Stats   Avg      Stdev     Max   +/- Stdev
#     Latency   649.36us    0.96ms  11.91ms   87.57%
#     Req/Sec    55.29k     8.26k   86.93k    69.25%
#   2208101 requests in 10.04s, 280.07MB read
# Requests/sec: 219_875.56
# Transfer/sec:     27.89MB



```


```sh
 ab -n 2208160 -c 80 http://localhost:8080/?q=1234567890
# This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
# Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
# Licensed to The Apache Software Foundation, http://www.apache.org/

# Benchmarking localhost (be patient)
# Completed 220816 requests
# Completed 441632 requests
# Completed 662448 requests
# Completed 883264 requests
# Completed 1104080 requests
# Completed 1324896 requests
# Completed 1545712 requests
# Completed 1766528 requests
# Completed 1987344 requests
# Completed 2208160 requests
# Finished 2208160 requests


# Server Software:        
# Server Hostname:        localhost
# Server Port:            8080

# Document Path:          /?q=1234567890
# Document Length:        16 bytes

# Concurrency Level:      80
# Time taken for tests:   122.136 seconds
# Complete requests:      2208160
# Failed requests:        0
# Total transferred:      293685280 bytes
# HTML transferred:       35330560 bytes
# Requests per second:    18079.47 [#/sec] (mean)
# Time per request:       4.425 [ms] (mean)
# Time per request:       0.055 [ms] (mean, across all concurrent requests)
# Transfer rate:          2348.21 [Kbytes/sec] received

# Connection Times (ms)
#               min  mean[+/-sd] median   max
# Connect:        0    2   0.7      2      16
# Processing:     1    2   0.7      2      26
# Waiting:        0    1   0.7      1      16
# Total:          2    4   0.8      4      39

# Percentage of the requests served within a certain time (ms)
#   50%      4
#   66%      5
#   75%      5
#   80%      5
#   90%      5
#   95%      6
#   98%      7
#   99%      7
#  100%     39 (longest request)
```
