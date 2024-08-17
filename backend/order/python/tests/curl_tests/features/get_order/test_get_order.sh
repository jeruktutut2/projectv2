#!/bin/bash

curl -X GET \
   -H "Content-Type: application/json" \
   -H "X-REQUEST-ID: requestId" \
   -d '{}' \
   http://127.0.0.1:10005/api/v1/order?orderId=2