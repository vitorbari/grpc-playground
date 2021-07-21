## 12-greet-service-with-reflection-golang

Implements a gRPC server with support to reflection using Golang.

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

### Testing

#### Using fullstorydev/grpcurl

[grpcurl installation instructions](https://github.com/fullstorydev/grpcurl#installation).

`$ grpcurl -plaintext localhost:50051 list`

```
greet.GreetService
grpc.reflection.v1alpha.ServerReflection
```

`$ grpcurl -plaintext localhost:50051 describe greet.GreetService`

```
greet.GreetService is a service:
service GreetService {
  rpc Greet ( .greet.GreetRequest ) returns ( .greet.GreetResponse );
}
```

`$ grpcurl -plaintext localhost:50051 describe greet.GreetRequest`

```
greet.GreetRequest is a message:
message GreetRequest {
  .greet.Greeting greeting = 1;
}
```

`$ grpcurl -plaintext localhost:50051 describe greet.GreetResponse`

```
greet.GreetResponse is a message:
message GreetResponse {
  string result = 1;
}
```

`$ grpcurl -plaintext localhost:50051 describe greet.GreetService.Greet`

```
greet.GreetService.Greet is a method:
rpc Greet ( .greet.GreetRequest ) returns ( .greet.GreetResponse );
```

Invoking:

`$ grpcurl -plaintext -d '{"greeting": {"first_name": "John", "last_name": "Days"}}' localhost:50051 greet.GreetService.Greet`

```
{
  "result": "Hello John Days"
}
```


#### Using ktr0731/evans

[evans installation instructions](https://github.com/ktr0731/evans#installation).

`$ evans -r --host localhost -p 50051 repl`

```

  ______
 |  ____|
 | |__    __   __   __ _   _ __    ___
 |  __|   \ \ / /  / _. | | '_ \  / __|
 | |____   \ V /  | (_| | | | | | \__ \
 |______|   \_/    \__,_| |_| |_| |___/

 more expressive universal gRPC client


greet.GreetService@localhost:50051> show package
+-------------------------+
|         PACKAGE         |
+-------------------------+
| greet                   |
| grpc.reflection.v1alpha |
+-------------------------+

greet.GreetService@localhost:50051> package greet

greet@localhost:50051> show service
+--------------+-------+--------------+---------------+
|   SERVICE    |  RPC  | REQUEST TYPE | RESPONSE TYPE |
+--------------+-------+--------------+---------------+
| GreetService | Greet | GreetRequest | GreetResponse |
+--------------+-------+--------------+---------------+

greet@localhost:50051> show message
+---------------+
|    MESSAGE    |
+---------------+
| GreetRequest  |
| GreetResponse |
+---------------+

greet@localhost:50051> desc GreetRequest
+----------+-------------------------+----------+
|  FIELD   |          TYPE           | REPEATED |
+----------+-------------------------+----------+
| greeting | TYPE_MESSAGE (Greeting) | false    |
+----------+-------------------------+----------+

greet@localhost:50051> service GreetService

greet.GreetService@localhost:50051> call Greet
greeting::first_name (TYPE_STRING) => John
greeting::last_name (TYPE_STRING) => Days
{
  "result": "Hello John Days"
}
```
