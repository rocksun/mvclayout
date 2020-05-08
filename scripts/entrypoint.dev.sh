#!/bin/bash

set -e

# export GOPROXY=https://goproxy.cn
# export GO111MODULE=on

go get github.com/rocksun/CompileDaemon

CompileDaemon --verbose --build="go build -o main cmd/api/main.go" --command=./main --command-stop