protoc -Iproto --go_out=. --go_opt=module=github.com/puffyguy/grpcUnary --go-grpc_out=. --go-grpc_opt=module=github.com/puffyguy/grpcUnary proto/calculator.proto
