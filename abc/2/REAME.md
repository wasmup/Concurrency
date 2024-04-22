## REAME.md

**Concurrent Execution Time Measurement**

This Go program demonstrates concurrent execution time measurement using synchronization primitives. This version measures the execution time after a wake-up signal is received.

**Functionality:**

* Spawns multiple goroutines (workers) to simulate some work.
* Captures the start time before spawning workers (`t0`).
* Measures the execution time for each worker after receiving a wake-up signal.
* Captures the wake-up time (`t1`) just before signaling the workers.

**Example Output:**

```sh
Worker 0 duration since Wake up call Broadcast: 4.712µs duration since t0: 22.909µs
Worker 1 duration since Wake up call Broadcast: 4.903µs duration since t0: 23.128µs
Worker 2 duration since Wake up call Broadcast: 5.075µs duration since t0: 23.296µs
Worker 3 duration since Wake up call Broadcast: 5.269µs duration since t0: 23.49µs
Worker 4 duration since Wake up call Broadcast: 5.442µs duration since t0: 23.666µs
Worker 5 duration since Wake up call Broadcast: 5.611µs duration since t0: 23.834µs
Worker 6 duration since Wake up call Broadcast: 5.779µs duration since t0: 24.003µs
Worker 7 duration since Wake up call Broadcast: 5.954µs duration since t0: 24.178µs
Worker 8 duration since Wake up call Broadcast: 6.249µs duration since t0: 24.473µs
Worker 9 duration since Wake up call Broadcast: 3.806µs duration since t0: 20.936µs
```