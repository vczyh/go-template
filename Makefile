all: test build docker

.PHONY: pre
pre:
	chmod -R +x scripts

.PHONY: test
test: pre
	scripts/app_test.sh

.PHONY: build
build: pre
	scripts/app_build.sh

.PHONY: docker
docker: pre
	scripts/docker_build.sh