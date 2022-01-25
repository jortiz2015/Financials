#!/bin/bash

protoc -I=./pb finsvc.proto \
    --go_out=plugins=grpc:./finsvc/pb #\
#    --js_out=import_style=commonjs,binary:./path/to/desired/typescript/output \
#    --grpc-web_out=import_style=typescript,mode=grpcweb:./path/to/desired/typescript/output

