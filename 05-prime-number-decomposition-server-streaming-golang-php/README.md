## 05-prime-number-decomposition-server-streaming-golang-php

Implements a gRPC streaming-server and a PHP client.

### Generating source files from .proto definition

```shell
$ protoc --go_out=server/calculator --go_opt=paths=source_relative \
    --go-grpc_out=server/calculator --go-grpc_opt=paths=source_relative \
    --go-grpc_opt=require_unimplemented_servers=false \
    --php_out=client/calculator \
    --grpc_out=client/calculator \
    --plugin=protoc-gen-grpc=bins/opt/grpc_php_plugin \
    calculator.proto
```


### Starting Server

`$ cd server && go run main.go`

### Running Client

`$ php client/run.php`

Output:

```
2,2,2,2,3,5,5
```
