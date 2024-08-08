#!/bin/bash

curl -X PATCH \
    -H "Content-Type: application/json" \
    -H "X-REQUEST-ID: requestId" \
    -d '{}' \
    http://localhost:10003/api/v1/products

echo ""

curl -X PATCH \
    -H "Content-Type: application/json" \
    -H "X-REQUEST-ID: requestId" \
    -d '{"id": 1, "name": "name edit", "description": "description edit"}' \
    http://localhost:10003/api/v1/products