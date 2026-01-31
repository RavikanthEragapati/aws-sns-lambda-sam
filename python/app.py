
import time

cold_start = True
init_time = time.time()

def handler(event, context):
    global cold_start

    start = time.time()

    print("Language=Python")
    print(f"ColdStart={cold_start}")
    print(f"InitDurationMs={(time.time() - init_time) * 1000:.2f}")

    cold_start = False

    print(f"ExecutionDurationMs={(time.time() - start) * 1000:.2f}")
