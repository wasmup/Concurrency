```sh
go run .


wrk -t4 -c80 -d10s http://localhost:8080/?q=1234567890

# Running 10s test @ http://localhost:8080/?q=1234567890
#   4 threads and 80 connections
#   Thread Stats   Avg      Stdev     Max   +/- Stdev
#     Latency   837.69us    1.12ms  22.72ms   87.55%
#     Req/Sec    35.68k     6.12k   50.53k    68.75%
#   1422329 requests in 10.02s, 180.41MB read
# Requests/sec: 141910.88
# Transfer/sec:     18.00MB



```


```sh
 ab -n 2208160 -c 80 http://localhost:8080/?q=1234567890
```
