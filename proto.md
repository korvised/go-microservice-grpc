```shell
protoc -Igreet/proto --go_out=. --go-grpc_out=. --go_opt=module=github.com/korvised/go-microservice-grpc --go-grpc_opt=module=github.com/korvised/go-microservice-grpc greet/proto/dummy.proto
```