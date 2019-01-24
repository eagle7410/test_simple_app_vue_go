#!/usr/bin/env bash
curl -X POST -v "http://localhost:6060/v2/user" -H "accept: application/xml" -H "Content-Type: application/json" -d "{\"email\": \"test@gg.com\", \"password\": \"123\"}"
