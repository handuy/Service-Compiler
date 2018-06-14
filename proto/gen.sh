#!/usr/bin/env bash

protoc --proto_path=. -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/gogo/protobuf/protobuf \
--gogoslick_out=. \
--micro_out=.  \
base.proto