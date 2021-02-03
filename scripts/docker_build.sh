#!/bin/bash

cd "$(dirname "$0")" || exit 1
docker build -f ../build/package/Dockerfile .. -t myapp
