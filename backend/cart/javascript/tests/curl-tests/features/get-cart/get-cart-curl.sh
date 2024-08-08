#!/bin/bash

# success
curl -X GET \
    -H "Content-Type: application/json" \
    -H "X-REQUEST-ID: requestId" \
    -d '{"userId": 1}' \
    http://localhost:10004/api/v1/carts