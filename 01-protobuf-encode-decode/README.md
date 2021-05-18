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
