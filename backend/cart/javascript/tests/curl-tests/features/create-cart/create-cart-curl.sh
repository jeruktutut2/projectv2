#!/bin/bash

# empty request body
curl -X POST \
    -H "Content-Type: application/json" \
    -H "X-REQUEST-ID: requestId" \
    -d '{}' \
    http://localhost:10004/api/v1/carts

echo ""

# success
curl -X POST \
    -H "Content-Type: application/json" \
    -H "X-REQUEST-ID: requestId" \
    -d '{"userId": 1, "productId": 2, "quantity": 1}' \
    http://localhost:10004/api/v1/carts