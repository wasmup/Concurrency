import threading
import time

def fib(n: int) -> int:
  if n == 1:
    return 0
  if n == 2:
    return 1
  return fib(n - 1) + fib(n - 2)

def print_fib(n: int) -> None:
  print(f"fib({n}) = {fib(n)}")

a,b = 41,42
start = time.time()

t1 = threading.Thread(target=print_fib, args=(a,))
t1.start()

print_fib(b)

t1.join() 

print(f'took {time.time() - start:.4f} seconds.')

# 102334155
# 165580141
# took 83.4180 seconds.