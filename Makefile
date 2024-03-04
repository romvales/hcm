include .env

export $(shell sed 's/=.*//' .env)
export ROOTDIR=$(shell pwd)
export GODIR=${ROOTDIR}/src/goServer

run:
	@cd ${GODIR} && \
		[ -e ./main ] && ./main || go run cmd/main.go

protobuild:
	@./protogen.sh

build:
	@cd ${GODIR} && \
		go build cmd/main.go

dev:
	@make clean && make run

test:
	@make clean
	@cd ${GODIR} && go test -v ./internal/core/hcmcore/...

clean:
	@cd ${GODIR} && [ -e ./main ] && rm ./main || echo &> /dev/null;
	@cd ${GODIR} && go clean -testcache;