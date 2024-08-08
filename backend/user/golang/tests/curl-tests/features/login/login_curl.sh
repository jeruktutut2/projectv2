#!/bin/bash

curl -X POST \
    -H "Content-Type: application/json" \
    -H "X-REQUEST-ID: requestId" \
    -d '{}' \
    http://localhost:10002/api/v1/users/login

curl -X POST \
    -H "Content-Type: application/json" \
    -H "X-REQUEST-ID: requestId" \
    -d '{"email": "email1@email.com", "password": "password@A1"}' \
    http://localhost:10002/api/v1/users/login