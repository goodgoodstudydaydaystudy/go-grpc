#!/bin/sh

find server -iname "*.proto" -exec protoc --go_out=plugins=grpc:../../ {} \;
find logic -iname "*.proto" -exec protoc --go_out=plugins=grpc:../../ {} \;
