#!/bin/bash

curl -X DELETE \
   -H "Content-Type: application/json" \
   -H "X-REQUEST-ID: requestId" \
   -d '{"id": 4}' \
   http://127.0.0.1:10005/api/v1/order