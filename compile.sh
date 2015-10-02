#!/bin/bash

./make.sh
mkdir -p build
GOOS=linux GOARCH=arm GOARM=7 go build -o build/arm main.go