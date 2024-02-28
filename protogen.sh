#!/bin/sh -

export PROJECT_DIR="$(pwd)"
export GO_SERVERPATH="$(pwd)/src/goServer"
export HCMCORE_PATH=$GO_SERVERPATH/internal/core/pb

cd $GO_SERVERPATH

protoc \
  -I=$PROJECT_DIR/proto \
  --go_out=$HCMCORE_PATH \
  --go_opt=paths=source_relative \
  --go-grpc_out=$HCMCORE_PATH \
  --go-grpc_opt=paths=source_relative \
  $PROJECT_DIR/proto/hcmcore.proto