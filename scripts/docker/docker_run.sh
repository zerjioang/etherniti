#!/bin/bash

#
# Copyright Etherniti Project. All Rights Reserved.
# SPDX-License-Identifier: GNU GPL v3
#

#source ./colors.sh

echo "downloading etherniti/proxy-oss:$(arch)-latest docker image"
docker pull etherniti/proxy-oss:$(arch)-latest

echo "stopping previous proxy containers..."
docker stop api && docker rm api

echo "deploying proxy container"
docker run \
    -d \
    -t \
    -p 80:8080 \
    -p 443:4430 \
    --cpus=4 \
    --hostname apollo-arm \
    --memory=2gb \
    --restart unless-stopped \
    --name api \
    --log-opt max-size=20m \
    --log-opt max-file=5 \
    --log-opt labels=production_status \
    -e X_ETHERNITI_SSL_CERT_FILE=/home/sergio/go/src/github.com/zerjioang/etherniti/secrets/letsencrypt/etc/certs/live/etherniti.org/fullchain.pem  \
    -e X_ETHERNITI_SSL_KEY_FILE=/home/sergio/go/src/github.com/zerjioang/etherniti/secrets/letsencrypt/etc/certs/live/etherniti.org/privkey.pem \
    -e X_ETHERNITI_LOG_LEVEL="warn" \
    -e X_ETHERNITI_LISTENING_MODE="http" \
    -e X_ETHERNITI_LISTENING_PORT="8080" \
    -e X_ETHERNITI_HTTPS_PORT="4430" \
    -e X_ETHERNITI_DEBUG_SERVER=true \
    -e X_ETHERNITI_HIDE_SERVER_DATA_IN_CONSOLE=true \
    -e X_ETHERNITI_TOKEN_SECRET="your-secret-jwt-key" \
    -e X_ETHERNITI_ENABLE_HTTPS_REDIRECT=false \
    -e X_ETHERNITI_ENABLE_LOGGING=true \
    -e X_ETHERNITI_USE_UNIQUE_REQUEST_ID=false \
    -e X_ETHERNITI_ENABLE_SECURITY=true \
    -e X_ETHERNITI_ENABLE_ANALYTICS=true \
    -e X_ETHERNITI_ENABLE_CORS=true \
    -e X_ETHERNITI_ENABLE_CACHE=true \
    -e X_ETHERNITI_ENABLE_RATE_LIMIT=false \
    -e X_ETHERNITI_BLOCK_TOR_CONNECTIONS=false \
    -e X_ETHERNITI_LISTENING_INTERFACE="0.0.0.0" \
    -e X_ETHERNITI_LISTENING_ADDRESS="0.0.0.0" \
    -e X_ETHERNITI_SWAGGER_ADDRESS="proxy.etherniti.org" \
    -e X_ETHERNITI_TOKEN_EXPIRATION=600 \
    -e X_ETHERNITI_RATELIMIT=10 \
    -e X_ETHERNITI_RATE_LIMIT_UNITS=10 \
    -e X_ETHERNITI_INFURA_TOKEN="153d442fd9d1449abb44150e515d1060" \
    -v $(pwd)/volumes/data:/root/.etherniti:rw \
    -v /etc/letsencrypt:/etc/letsencrypt \
    etherniti/proxy-oss:$(arch)-latest

echo "etherniti proxy container deployed"
docker ps