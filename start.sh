#!/bin/bash
echo "Starting VaultGuard API server..."
echo "HTTP Server will start on: $HTTP_SERVER_ADDRESS"
echo "GRPC Server will start on: $GRPC_SERVER_ADDRESS"
./main
