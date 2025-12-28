#!/bin/bash

echo "Waiting for services to be healthy..."
sleep 15

echo "Testing Complete Data..."
curl -s -X POST http://localhost:8080/enrichment \
-H "Content-Type: application/json" \
-d '{"customer_code": "123e4567-e89b-12d3-a456-426614174000"}'

echo -e "\n\nTesting Missing Address..."
curl -s -X POST http://localhost:8080/enrichment \
-H "Content-Type: application/json" \
-d '{"customer_code": "123e4567-e89b-12d3-a456-426614174001"}'

echo -e "\n\nTesting Missing Suggestions..."
curl -s -X POST http://localhost:8080/enrichment \
-H "Content-Type: application/json" \
-d '{"customer_code": "123e4567-e89b-12d3-a456-426614174002"}'

echo -e "\n\nTesting Missing PII (Random UUID)..."
curl -s -X POST http://localhost:8080/enrichment \
-H "Content-Type: application/json" \
-d '{"customer_code": "123e4567-e89b-12d3-a456-426614174999"}'
