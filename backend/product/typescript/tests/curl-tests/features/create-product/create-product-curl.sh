#!/bin/bash

curl -X POST \
   -H "Content-Type: application/json" \
   -H "X-REQUEST-ID: requestId" \
   -d '{}' \
   http://localhost:10003/api/v1products

echo ""

curl -X POST \
   -H "Content-Type: application/json" \
   -H "X-REQUEST-ID: requestId" \
   -d '{"userId": 1, "name": "name1", "description": "description1", "stock": 1}' \
   http://localhost:10003/api/v1/products