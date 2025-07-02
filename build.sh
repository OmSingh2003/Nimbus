#!/bin/bash
# Build script for Render deployment

echo "Installing dependencies..."
go mod tidy

echo "Building application..."
go build -o main main.go

echo "Build completed successfully!"
