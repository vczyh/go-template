.PHONY: build

all: pre build

pre:
	chmod -R +x scripts

build:
	./scripts/app_build.sh

docker:
	./scripts/docker_build.sh