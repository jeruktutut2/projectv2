#!/bin/bash

curl -X PATCH \
   -H "Content-Type: application/json" \
   -H "X-REQUEST-ID: requestId" \
   -d '{}' \
   http://127.0.0.1:10005/api/v1/order

echo ""

curl -X PATCH \
   -H "Content-Type: application/json" \
   -H "X-REQUEST-ID: requestId" \
   -d '{"id": 4, "orderItems": [{"productId": 1, "quantity": 2}, {"productId": 2, "quantity": 3}, {"productId": 3, "quantity": 1}]}' \
   http://127.0.0.1:10005/api/v1/order