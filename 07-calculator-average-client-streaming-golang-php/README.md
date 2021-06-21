## 07-calculator-average-client-streaming-golang-php

Implements a gRPC with Golang server and PHP streaming-client.

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
Numbers: Array
(
    [0] => 1
    [1] => 2
    [2] => 3
    [3] => 4
)
Response from RPC - Average = 2.5
```
