# grpc-playground

A playground repo for gRPC (and protobuf) stuff.

## Catalog

Here you can find the labs organized by different criterias:

### By Concept

| Concept                     | Labs                   | Docs                                                                         |
|-----------------------------|------------------------|------------------------------------------------------------------------------|
| Unary RPC                   | [02](./02-greet-service-golang), [03](./03-calculator-service-golang-php), [10](./10-error-handling-golang), [11](./11-calculator-with-deadline-golang-php), [12](./12-greet-service-with-reflection-golang), [13](./13-crud-api-mongodb-golang) | https://grpc.io/docs/what-is-grpc/core-concepts/#unary-rpc                   |
| Server Streaming RPC        | [04](./04-greet-server-streaming-golang), [05](./05-prime-number-decomposition-server-streaming-golang-php)                 | https://grpc.io/docs/what-is-grpc/core-concepts/#server-streaming-rpc        |
| Client Streaming RPC        | [06](./06-greet-client-streaming-golang), [07](./07-calculator-average-client-streaming-golang-php)                 | https://grpc.io/docs/what-is-grpc/core-concepts/#client-streaming-rpc        |
| Bidirectional streaming RPC | [08](./08-greet-bi-directional-streaming-golang), [09](./09-calculator-maximum-bi-directional-streaming-golang-php)                 | https://grpc.io/docs/what-is-grpc/core-concepts/#bidirectional-streaming-rpc |
| Deadlines/Timeouts          | [11](./11-calculator-with-deadline-golang-php)                     | https://grpc.io/docs/what-is-grpc/core-concepts/#deadlines                   |
| Error Handling              | [10](./10-error-handling-golang)                     | https://www.grpc.io/docs/guides/error/                                       |
| Reflection                  | [12](./12-greet-service-with-reflection-golang), [13](./13-crud-api-mongodb-golang)                 | https://github.com/grpc/grpc/blob/master/doc/server-reflection.md            |

### By Language

| Language | Labs                                       | Docs                                      |
|----------|--------------------------------------------|-------------------------------------------|
| Go       | [02](./02-greet-service-golang), [03](./03-calculator-service-golang-php), [04](./04-greet-server-streaming-golang), [05](./05-prime-number-decomposition-server-streaming-golang-php), [06](./06-greet-client-streaming-golang), [07](./07-calculator-average-client-streaming-golang-php), [08](./08-greet-bi-directional-streaming-golang), [09](./09-calculator-maximum-bi-directional-streaming-golang-php), [10](./10-error-handling-golang), [11](./11-calculator-with-deadline-golang-php), [12](./12-greet-service-with-reflection-golang), [13](./13-crud-api-mongodb-golang) | https://grpc.io/docs/languages/go/basics/ |
| PHP      | [01](01-protobuf-serialize-deserialize-php), [03](./03-calculator-service-golang-php), [05](./05-prime-number-decomposition-server-streaming-golang-php), [07](./07-calculator-average-client-streaming-golang-php), [09](./09-calculator-maximum-bi-directional-streaming-golang-php), [11](./11-calculator-with-deadline-golang-php)                     | https://grpc.io/docs/languages/php/       |

## TODO

Topics to cover:

- Middlewares: https://github.com/grpc-ecosystem/go-grpc-middleware
- gRPC-Gateway: https://github.com/grpc-ecosystem/grpc-gateway
