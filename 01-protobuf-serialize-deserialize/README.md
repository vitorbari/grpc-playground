## Generating source files from .proto definition

`$ mkdir -p generated/php`

`$ protoc customers.proto --php_out=./generated/php`

## Using generated source-code to encode / decode data

`$ php src/php/customers.php | jq .`

Output:

```json
{
  "customers": [
    {
      "id": 1,
      "name": "John"
    },
    {
      "id": 2,
      "name": "Bob"
    }
  ]
}
```

## Using protoc --decode_raw

`$ php src/php/serialize-stdout.php | protoc --decode_raw`

Output:

```
1 {
  1: 1
  2: "John"
}
1 {
  1: 2
  2: "Bob"
}
```
