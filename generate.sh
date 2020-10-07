#!/bin/bash

cd protos

protoc -I . --grpc-gateway_out ../internal/gen \
     --grpc-gateway_opt logtostderr=true \
     --grpc-gateway_opt generate_unbound_methods=true \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
--go_out=plugins=grpc:../internal/gen ./echo.proto

protoc -I . --swagger_out ../doc/swagger --swagger_opt logtostderr=true \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
./echo.proto
