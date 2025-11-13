#!/bin/bash
set -e

trap 'echo "Error occurred on line $LINENO with exit code $?"' ERR

GOOS=js GOARCH=wasm go build -o main.wasm main.go

echo "Build successful: main.wasm created"