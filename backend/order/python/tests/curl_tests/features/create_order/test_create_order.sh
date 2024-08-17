#!/bin/bash

curl -X POST \
   -H "Content-Type: application/json" \
   -H "X-REQUEST-ID: requestId" \
   -d '{}' \
   http://127.0.0.1:10005/api/v1/order

echo ""

curl -X POST \
   -H "Content-Type: application/json" \
   -H "X-REQUEST-ID: requestId" \
   -d '{"userId": 1, "orderItems": [{"productId": 1, "quantity": 1}, {"productId": 2, "quantity": 2}]}' \
   http://127.0.0.1:10005/api/v1/order