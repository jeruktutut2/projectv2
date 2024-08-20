#!/bin/bash

curl -X POST \
    -H "Content-Type: application/json" \
    -H "X-REQUEST-ID: requestId" \
    -H "X-SESSION-USER-ID: sessionId" \
    -d '{}' \
    http://localhost:10002/api/v1/users/logout

echo ""

curl -X POST \
    -H "Content-Type: application/json" \
    -H "X-REQUEST-ID: requestId" \
    -H "X-SESSION-USER-ID: sessionId" \
    -d '{}' \
    http://localhost:10002/api/v1/users/logout