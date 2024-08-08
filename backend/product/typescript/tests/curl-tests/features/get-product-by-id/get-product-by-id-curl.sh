#!/bin/bash

curl -X GET \
    -H "Content-Type: application/json" \
    -H "X-REQUEST-ID: requestId" \
    -d '{}' \
    http://localhost:10003/api/v1/products

echo ""

curl -X GET \
    -H "Content-Type: application/json" \
    -H "X-REQUEST-ID: requestId" \
    -d '{"id": 2}' \
    http://localhost:10003/api/v1/products