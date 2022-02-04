#!/bin/bash

protoc -I=./pb finsvc.proto \
    --go_out=plugins=grpc:./finsvc/pb \
    --js_out=import_style=commonjs,binary:./web/pages \
    --grpc-web_out=import_style=typescript,mode=grpcweb:./web/pages

