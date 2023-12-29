```sh
pip install fastapi uvicorn
uvicorn main:app --host localhost --port 8080 --log-level critical
curl -i  http://localhost:8080/?q=Hello

wrk -t4 -c80 -d10s http://localhost:8080/?q=1234567890
# Running 10s test @ http://localhost:8080/?q=1234567890
#   4 threads and 80 connections
#   Thread Stats   Avg      Stdev     Max   +/- Stdev
#     Latency    20.28ms   45.18ms 725.16ms   97.74%
#     Req/Sec     1.37k   359.94     2.61k    63.16%
#   54313 requests in 10.01s, 7.77MB read
# Requests/sec:   5423.83
# Transfer/sec:    794.51KB

uvicorn main:app --reload --host 0.0.0.0 --port 8080

uvicorn main:app --reload



```


```sh
 ab -n 2208160 -c 80 http://localhost:8080/?q=1234567890
```
