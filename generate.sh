protoc -I/usr/local/include -I. -I$GOPATH/src  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. api/proto/health.proto
protoc -I/usr/local/include -I. -I$GOPATH/src  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. api/proto/health.proto
protoc -I/usr/local/include -I. -I$GOPATH/src  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. api/proto/dog.proto
protoc --doc_out=html,index.html:./docs api/proto/dog.proto