#!/bin/bash
set -e  # Exit on error

GOOS=js GOARCH=wasm go build -o main.wasm main.go

echo "Build successful: main.wasm created"