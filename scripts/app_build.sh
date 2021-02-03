#!/bin/bash

cd "$(dirname "$0")" || exit 1
go build -o ../bin/myapp ../cmd/myapp