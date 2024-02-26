#!/bin/bash

# Install protoc-gen-go if not already installed
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# Add protoc-gen-go to the PATH
export PATH="$PATH:$(go env GOPATH)/bin"

# Define common paths
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
PROTO_DIR="$SCRIPT_DIR/../proto"
GO_OUT_DIR="$SCRIPT_DIR/go_out"
SERVER_DIR="$SCRIPT_DIR/../internal/quiz/generated"

# Define specific file names
PROTO_FILES="$PROTO_DIR/quiz.proto"
GO_OUT_FILES="$GO_OUT_DIR/quiz.pb.go $GO_OUT_DIR/quiz_grpc.pb.go"

# Generate Go code from the Protocol Buffers files
protoc --go_out="$GO_OUT_DIR" --go_opt=paths=source_relative \
        --go-grpc_out="$GO_OUT_DIR" --go-grpc_opt=paths=source_relative \
        -I="$PROTO_DIR" \
        "$PROTO_FILES"

# Move the generated files to SERVER_DIR
mkdir -p "$SERVER_DIR"
mv $GO_OUT_FILES "$SERVER_DIR"

# Remove the go_out directory and its contents after generating code
rm -rf GO_OUT_DIR
