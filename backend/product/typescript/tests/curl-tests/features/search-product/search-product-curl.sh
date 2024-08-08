#!/bin/bash

curl -X GET \
    -H "Content-Type: application/json" \
    -H "X-REQUEST-ID: requestId" \
    -d '{}' \
    http://localhost:10003/api/v1/products/search

echo ""

curl -X GET \
    -H "Content-Type: application/json" \
    -H "X-REQUEST-ID: requestId" \
    -d '{"keyword": "name"}' \
    http://localhost:10003/api/v1/products/search