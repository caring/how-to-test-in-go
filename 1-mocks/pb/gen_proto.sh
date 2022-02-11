#!/bin/bash
if [ -f "$(command -v protoc)" ] && [ -f "$(command -v protoc-gen-go-grpc)" ]; then
    VER=$(protoc --version)
    PBDIR="1-mocks/pb/"
    echo "Using protoc version: $VER"
    protoc \
      --proto_path=$PBDIR \
      --plugin=grpc \
      --go_out=$PBDIR --go_opt=paths=source_relative \
      --go-grpc_out=$PBDIR --go-grpc_opt=paths=source_relative \
      $PBDIR*.proto
else
    echo "Error: protoc was not found. Please check that it is installed."
fi