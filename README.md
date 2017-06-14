# stringsvc1
The purpose of this project was to learn go-kit.  This is my implementation of go-kit's stringsvc example with some additional features and REST API and grpc support.  

# How to get proto to work

* Get protoc from https://github.com/google/protobuf/releases

* go get -u golang.org/x/net/http2
* go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
* protoc stringsvc1.proto --go_out=plugins=grpc:.

# Install protoc3 on Ubuntu
## Make sure you grab the latest version
curl -OL https://github.com/google/protobuf/releases/download/v3.3.0/protoc-3.3.0-linux-x86_64.zip

## Unzip
unzip protoc-3.3.0-linux-x86_64.zip -d protoc3

## Move only protoc* to /usr/bin/
sudo mv protoc3/bin/protoc /usr/bin/protoc

# To test

* curl -X POST -d '{"input_string": "robert"}' http://localhost:8080/uppercase
* curl -X GET -d '{"input_string": "robert"}' http://localhost:8080/count
