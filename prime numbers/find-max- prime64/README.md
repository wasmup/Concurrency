```sh
time go run .
# 18446744073709551557

# real    0m3.440s
# user    0m25.392s
# sys     0m0.126s

time go run -race .
# 18446744073709551557

# real    0m42.697s
# user    5m18.210s
# sys     0m0.804s


```