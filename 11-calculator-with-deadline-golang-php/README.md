## 11-calculator-with-deadline-golang-php

Implements a gRPC call with deadline using Golang for the server and PHP for the client.

Reference: https://grpc.io/blog/deadlines/

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
Factorial of 1 = 1
Factorial of 2 = 2
Factorial of 3 = 6
Factorial of 4 = 24
Factorial of 5 = 120
Factorial of 6 = 720
Factorial of 7 = 5040
Factorial of 8 = 40320
Factorial of 9 = 362880
Error when calculating the factorial of: 10. Error: 4, Deadline Exceeded
Error when calculating the factorial of: 11. Error: 4, Deadline Exceeded
Error when calculating the factorial of: 12. Error: 4, Deadline Exceeded
Error when calculating the factorial of: 13. Error: 4, Deadline Exceeded
Error when calculating the factorial of: 14. Error: 4, Deadline Exceeded
Error when calculating the factorial of: 15. Error: 4, Deadline Exceeded
Error when calculating the factorial of: 16. Error: 4, Deadline Exceeded
Error when calculating the factorial of: 17. Error: 4, Deadline Exceeded
Error when calculating the factorial of: 18. Error: 4, Deadline Exceeded
Error when calculating the factorial of: 19. Error: 4, Deadline Exceeded
Error when calculating the factorial of: 20. Error: 4, Deadline Exceeded
```
