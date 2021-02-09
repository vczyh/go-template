#!/bin/bash

cd "$(dirname $0)" || exit 1

export GO111MODULE=on
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

export GOPROXY=https://goproxy.io,direct

go build -o ../bin/myapp ../cmd/myapp
