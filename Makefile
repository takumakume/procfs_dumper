VERSION ?= $(shell git describe --tag --abbrev=0)

build: export GO111MODULE=on
build:
	go build --ldflags "-s -w -X main.version=$(VERSION)" ./cmd/procfs_dump

docker: docker_build
	docker run --security-opt=seccomp:unconfined -it -v ${PWD}:/go/src procfs_dumper bash

docker_build:
	docker build -t procfs_dumper .
