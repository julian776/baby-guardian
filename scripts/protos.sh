#!/bin/bash

proto_path="./protos"

protoc --proto_path=${proto_path} \
  --go_out=${proto_path}/gen/go \
  --go_opt=paths=source_relative \
  signal.proto

cd ${proto_path}/gen/go

go mod tidy
