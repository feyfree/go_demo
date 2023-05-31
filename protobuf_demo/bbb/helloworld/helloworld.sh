# protoc --proto_path=protobuf_demo/bbb/helloworld --go_out=. helloworld.proto

protoc --proto_path=protobuf_demo/bbb/helloworld --go_out=. --go-grpc_out=. helloworld.proto