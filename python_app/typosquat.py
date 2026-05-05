#!/usr/bin/python3

import hashlib
import request
import sys

url = "https://api.github.com/events"
response = request.get(url)

hash = hashlib.md5(response.status_code.encode()).hexdigest()

# Accessing response data
print(f"Status Code: {hash}")

sys.exit(0)
