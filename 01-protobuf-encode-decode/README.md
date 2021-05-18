`$ mkdir generated`

`$ mkdir generated/php`

`$ protoc customers.proto --php_out=./generated/php`

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
