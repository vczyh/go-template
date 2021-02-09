#!/bin/bash

cd "$(dirname $0)" || exit 1

export CGO_ENABLED=0

go test ../pkg...
