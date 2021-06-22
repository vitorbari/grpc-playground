## 08-greet-bi-directional-streaming-golang

Implements a bi-directional streaming RPC using Golang.

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
Starting client streaming...
Sending req: greeting:{first_name:"John"  last_name:"Days"}
Received: Hello John

Sending req: greeting:{first_name:"Bob"  last_name:"Bar"}
Received: Hello Bob

Sending req: greeting:{first_name:"Alice"  last_name:"Ger"}
Received: Hello Alice

```
