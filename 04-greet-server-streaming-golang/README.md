## 04-greet-server-streaming-golang

Implements a server-streaming RPC using Golang.

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
    greetpb/greet.proto
```

### Starting Server

`$ go run server/server.go`

### Running Client

`$ go run client/client.go`

Output:

```
Starting client...
2021/06/11 10:24:07 Stream Response: 0 - Hello John Days
2021/06/11 10:24:08 Stream Response: 1 - Hello John Days
2021/06/11 10:24:09 Stream Response: 2 - Hello John Days
2021/06/11 10:24:10 Stream Response: 3 - Hello John Days
2021/06/11 10:24:11 Stream Response: 4 - Hello John Days
2021/06/11 10:24:12 Stream Response: 5 - Hello John Days
2021/06/11 10:24:13 Stream Response: 6 - Hello John Days
2021/06/11 10:24:14 Stream Response: 7 - Hello John Days
2021/06/11 10:24:15 Stream Response: 8 - Hello John Days
2021/06/11 10:24:16 Stream Response: 9 - Hello John Days
```
