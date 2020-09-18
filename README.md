# proto-golang
This guide describes how to use the protocol buffer language to structure your protocol buffer data, including .proto file syntax and how to generate data access classes from your .proto files.

## install go package
go install google.golang.org/protobuf/cmd/protoc-gen-go

## protoc example
protoc --proto_path=src --go_out=build/gen --go_opt=paths=source_relative src/foo.proto src/bar/baz.proto