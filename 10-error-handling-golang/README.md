## 10-error-handling-golang

Implements error handling on gRPC using Golang.

Reference: https://avi.im/grpc-errors/

### Installing the protocol compiler plugins for Go

```shell
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

$ export PATH="$PATH:$(go env GOPATH)/bin"
```

### Generating source files from .proto definition

```shell
$ protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    --go-grpc_opt=require_unimplemented_servers=false \
    calculator/calculator.proto
```

### Starting Server

`$ go run server/server.go`

### Running Client

`$ go run client/client.go`

Output:

```
Starting client...
2021/07/13 14:02:31 Sqrt of 16 is: 4.000000
2021/07/13 14:02:31 Sqrt of 9 is: 3.000000
2021/07/13 14:02:31 Sqrt of 1231 is: 35.085610
2021/07/13 14:02:31 Error message from server: Radicand must be a positive number, got: -4
2021/07/13 14:02:31 Error code from server: InvalidArgument
2021/07/13 14:02:31 Sent a invalid argument
2021/07/13 14:02:31 Sqrt of 25 is: 5.000000
```
