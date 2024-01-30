```sh
clang++ -O3 -march=native -o main main.cpp
./main
curl -i  http://localhost:8080/?q=1234567890
	
wrk -t4 -c80 -d10s http://localhost:8080/?q=1234567890
# Running 10s test @ http://localhost:8080/?q=1234567890
#   4 threads and 80 connections
#   Thread Stats   Avg      Stdev     Max   +/- Stdev
#     Latency     1.23ms   10.66ms 210.90ms   99.25%
#     Req/Sec    27.70k     6.62k   48.87k    71.00%
#   1102346 requests in 10.00s, 69.17MB read
# Requests/sec: 110204.55
# Transfer/sec:      6.92MB


g++ -O3 -std=c++17 main.cpp -o main
# Running 10s test @ http://localhost:8080/?q=1234567890
#   4 threads and 80 connections
#   Thread Stats   Avg      Stdev     Max   +/- Stdev
#     Latency     1.15ms    7.20ms 418.23ms   99.72%
#     Req/Sec    29.17k     3.16k   39.84k    84.75%
#   1161041 requests in 10.00s, 72.86MB read
# Requests/sec: 116056.45
# Transfer/sec:      7.28MB

g++ -std=c++17 main.cpp -o main

wrk -t4 -c80 -d10s http://localhost:8080/?q=1234567890

# Running 10s test @ http://localhost:8080/?q=1234567890
#   4 threads and 80 connections
#   Thread Stats   Avg      Stdev     Max   +/- Stdev
#     Latency     1.33ms    7.76ms 214.02ms   99.61%
#     Req/Sec    17.50k     3.77k   33.24k    73.25%
#   696621 requests in 10.10s, 43.71MB read
# Requests/sec:  68967.48
# Transfer/sec:      4.33MB



```


```sh
 ab -n 2208160 -c 80 http://localhost:8080/?q=1234567890
```
