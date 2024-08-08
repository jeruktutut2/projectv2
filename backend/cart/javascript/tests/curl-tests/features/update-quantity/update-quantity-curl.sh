#!/bin/bash

# empty request
curl -X PATCH \
    -H "Content-Type: application/json" \
    -H "X-REQUEST-ID: requestId" \
    -d '{}' \
    http://localhost:10004/api/v1/carts/update-quantity

echo ""

# success
curl -X PATCH \
    -H "Content-Type: application/json" \
    -H "X-REQUEST-ID: requestId" \
    -d '{"userId": 1, "productId": 2, "quantity": 2}' \
    http://localhost:10004/api/v1/carts/update-quantity