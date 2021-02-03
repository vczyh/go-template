#!/bin/bash

cd $(dirname $0)
docker build -f ../build/package/Dockerfile .. -t myapp
