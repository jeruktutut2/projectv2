#!/bin/bash

curl -X POST \
    -H "Content-Type: application/json" \
    -H "X-REQUEST-ID: requestId" \
    -d '{}' \
    http://localhost:10002/api/v1/users/register

echo ""

curl -X POST \
    -H "Content-Type: application/json" \
    -H "X-REQUEST-ID: requestId" \
    -d '{"username": "username1", "email": "email1@email.com", "password": "password@A1", "confirmpassword": "password@A1", "utc": "+0800"}' \
    http://localhost:10002/api/v1/users/register