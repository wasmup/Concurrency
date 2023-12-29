from fastapi import FastAPI
from fastapi.responses import PlainTextResponse
from concurrent.futures import ThreadPoolExecutor
import base64
import time

app = FastAPI()

request_count = 0
elapsed_time = 0
executor = ThreadPoolExecutor(max_workers=2)

def print_stats():
    global request_count, elapsed_time
    while True:
        time.sleep(1)
        with executor:
            print(f"Requests: {request_count}, Total Elapsed Time: {elapsed_time}")

@app.get("/")
async def read_root(q: str = "Hi"):
    t0 = time.time()
    encoded_q = base64.b64encode(q.encode()).decode()
    elapsed = time.time() - t0

    global request_count, elapsed_time
    with executor:
        request_count += 1
        elapsed_time += elapsed

    return PlainTextResponse(encoded_q)

if __name__ == "__main__":
    executor.submit(print_stats)
    import uvicorn
    uvicorn.run(app, host="localhost", port=8080)
