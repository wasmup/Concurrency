all:
	sudo apt install libbase64-dev libcpprest-dev

	g++ -o main main.cpp -lbase64 -lcpprest

	g++ -o main main.cpp -lcrypto -lcpprest
	
	g++ -O3 -flto -march=native -o main main.cpp -lcrypto -lcpprest
	./main
	wrk -t4 -c80 -d10s http://localhost:8080/?q=1234567890
# Running 10s test @ http://localhost:8080/?q=1234567890
#   4 threads and 80 connections
#   Thread Stats   Avg      Stdev     Max   +/- Stdev
#     Latency     1.84ms    2.14ms  39.02ms   90.94%
#     Req/Sec    13.39k     3.20k   22.05k    65.00%
#   534036 requests in 10.04s, 46.86MB read
# Requests/sec:  53207.96
# Transfer/sec:      4.67MB
	g++ -O3 -fprofile-generate -o main main.cpp -lcrypto -lcpprest
./main  # Run the application with typical workloads
g++ -O3 -fprofile-use -o main main.cpp -lcrypto -lcpprest

	
	clang++ -o main main.cpp -lcpprest -lcrypto

	wrk -t4 -c80 -d10s http://localhost:8080/?q=1234567890
# Running 10s test @ http://localhost:8080/?q=1234567890
#   4 threads and 80 connections
#   Thread Stats   Avg      Stdev     Max   +/- Stdev
#     Latency     2.34ms    2.99ms  57.41ms   91.25%
#     Req/Sec    11.21k     2.31k   17.76k    67.00%
#   447121 requests in 10.04s, 39.23MB read
# Requests/sec:  44520.64
# Transfer/sec:      3.91MB

	clang++ -O3 -flto -o main main.cpp -lcpprest -lcrypto

# -march=native: Optimizes for the specific architecture of your processor.
	clang++ -O3 -march=native -o main main.cpp -lcpprest -lcrypto
	./main
	curl -i  http://localhost:8080/?q=1234567890
	
	wrk -t4 -c80 -d10s http://localhost:8080/?q=1234567890
# Running 10s test @ http://localhost:8080/?q=1234567890
#   4 threads and 80 connections
#   Thread Stats   Avg      Stdev     Max   +/- Stdev
#     Latency     1.46ms    1.66ms  36.64ms   91.54%
#     Req/Sec    16.49k     2.66k   22.83k    70.75%
#   657028 requests in 10.02s, 60.15MB read
# Requests/sec:  65585.11
# Transfer/sec:      6.00MB