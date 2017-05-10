# How to get proto to work

* Get protoc from https://github.com/google/protobuf/releases

* go get -u golang.org/x/net/http2
* go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
* protoc stringsvc1.proto --go_out=plugins=grpc:.
