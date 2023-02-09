#!/bin/bash
export LANG=zh_CN.UTF-8

ifndef GIT_BRANCH
	GIT_BRANCH=`git symbolic-ref --short -q HEAD`
endif

ifndef GIT_HASH
	GIT_HASH=`git rev-parse --short HEAD`
endif

ifndef BUILD_TIME
	BUILD_TIME=`date '+%Y-%m-%dT%H:%M:%S'`
endif

ENVARG=CGO_ENABLED=0 GOPROXY=https://goproxy.cn,direct
LINUXARG=GOOS=linux GOARCH=amd64
BUILDARG=-mod=mod -ldflags " -s -X main.buildTime=${BUILD_TIME} -X main.gitHash=${GIT_BRANCH}:${GIT_HASH}"

dep:
	cd src; ${ENVARG} go get ./...; cd -

updep:
	cd src; ${ENVARG} go get -u ./...; go mod tidy; cd -

p:
	mkdir -p lib/proto
	rm -rf lib/proto/*

	cd ./; protoc -I ./protocol --go_out=. common.proto; cd -
	cd ./; protoc -I ./protocol --go-grpc_out=. common_service.proto; cd -
	cd ./; protoc -I ./protocol --go_out=. confuse_api.proto; cd -
#	cd src; protoc -I ../protocol --gofast_out=plugins=grpc:. voucher_admin.proto; cd -

	ls ./lib/proto/*/*.pb.go | xargs sed -i -e "s@\"lib/proto/@\"confuse/lib/proto/@"
	ls ./lib/proto/*/*.pb.go | xargs sed -i -e "s/,omitempty//"
	ls ./lib/proto/*/*.pb.go | xargs sed -i -e "s/json:\"\([a-zA-Z_-]*\)\"/json:\"\1\" form:\"\1\"/g"
	ls ./lib/proto/*/*.pb.go | xargs sed -i -e "/force omitempty/{n;s/json:\"\([a-zA-Z_-]*\)\"/json:\"\1,omitempty\"/g;}"

	rm -f ./lib/proto/*/*.pb.go-e

.PHONY: api
api:
	cd ./; ${ENVARG} go build ${BUILDARG} -o ./bin/api main.go;

linux_api:
	cd ./; ${ENVARG} ${LINUXARG} go build ${BUILDARG} -o ./lbin/api main.go;

all: p api

linux_all: p linux_api

clean:
	rm -fr bin/*
	rm -fr lbin/*