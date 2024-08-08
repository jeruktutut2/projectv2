#!/bin/bash

curl -X DELETE \
    -H "Content-Type: application/json" \
    -H "X-REQUEST-ID: requestId" \
    -d '{"id": "2", "userId": 1, "productId": 2}' \
    http://localhost:10004/api/v1/carts