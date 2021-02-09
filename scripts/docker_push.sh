#!/bin/bash

cd "$(dirname $0)" || exit 1

docker login harbor.local.com -u vczyh -p Vczyh1221

docker tag myapp harbor.local.com/demo/myapp:7
docker push harbor.local.com/demo/myapp:7
