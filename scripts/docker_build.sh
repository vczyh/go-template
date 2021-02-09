#!/bin/bash

cd "$(dirname $0)" || exit 1

docker build -f ../build/package/Dockerfile .. -t myapp

# PUSH
docker login harbor.local.com -u vczyh -p Vczyh1221
docker tag myapp harbor.local.com/demo/myapp

docker push harbor.local.com/demo/myapp
