# !/bin/sh
apk update && apk add musl-dev gcc protobuf protoc
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
mkdir -p pb

export PATH="$PATH:$(go env GOPATH)/bin"
echo "compiling protobuf.."

protoc  --go_out=./app/pb  --proto_path=./protos \
        --go_opt=paths=source_relative  --go-grpc_out=./app/pb --go-grpc_opt=paths=source_relative \
        protos/*.proto
