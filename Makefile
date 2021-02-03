.PHONY: build

all: pre build docker

pre:
	chmod -R +x scripts

build: pre
	./scripts/app_build.sh

docker: pre
	./scripts/docker_build.sh