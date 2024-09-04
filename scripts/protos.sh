#!/bin/bash

proto_path="./protos"
go_path="./gen/go"
go_out_path="${proto_path}/${go_path}"

protoc --proto_path=${proto_path} \
  --go_out=${go_out_path} \
  --go_opt=paths=source_relative \
  --go-grpc_out=${go_out_path} \
  --go-grpc_opt=paths=source_relative \
  --grpc-gateway_out=${go_out_path} \
  --grpc-gateway_opt paths=source_relative \
  --grpc-gateway_opt generate_unbound_methods=true \
  signal.proto \
  analytics.proto \
  auth.proto \

cd ${go_out_path}

go mod tidy
