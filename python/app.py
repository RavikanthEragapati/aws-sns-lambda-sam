import json
import time

cold_start = True
init_time = time.time()

def handler(event, context):
    global cold_start

    start = time.time()

    # Print full SNS event
    print("SNS Event:")
    print(json.dumps(event, indent=2))

    # Print message(s)
    for record in event.get("Records", []):
        message = record.get("Sns", {}).get("Message")
        print(f"SNS Message: {message}")

    print("Language=Python")
    print(f"ColdStart={cold_start}")
    print(f"InitDurationMs={(time.time() - init_time) * 1000:.2f}")

    cold_start = False

    print(f"ExecutionDurationMs={(time.time() - start) * 1000:.2f}")
