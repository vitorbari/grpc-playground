## 13-crud-api-mongodb-golang

Implements a gRPC server to expose CRUD operations for a blog. 

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
    blogpb/blog.proto
```

### Starting Server

`$ docker-compose up`

### Testing

`$ evans -r --host localhost -p 50051 repl`

```

  ______
 |  ____|
 | |__    __   __   __ _   _ __    ___
 |  __|   \ \ / /  / _. | | '_ \  / __|
 | |____   \ V /  | (_| | | | | | \__ \
 |______|   \_/    \__,_| |_| |_| |___/

 more expressive universal gRPC client


blog.BlogService@localhost:50051> call CreateBlog
blog::id (TYPE_STRING) =>
blog::author_id (TYPE_STRING) => vitorbari
blog::title (TYPE_STRING) => My First Blog Post
blog::content (TYPE_STRING) => this it the content of post 1
{
  "blog": {
    "authorId": "vitorbari",
    "content": "this it the content of post 1",
    "id": "61806040a838e0152aa0fd15",
    "title": "My First Blog Post"
  }
}

blog.BlogService@localhost:50051> call ReadBlog
blog_id (TYPE_STRING) => 61806040a838e0152aa0fd15
{
  "blog": {
    "authorId": "vitorbari",
    "content": "this it the content of post 1",
    "id": "61806040a838e0152aa0fd15",
    "title": "My First Blog Post"
  }
}

blog.BlogService@localhost:50051> call CreateBlog
blog::id (TYPE_STRING) =>
blog::author_id (TYPE_STRING) => john.days
blog::title (TYPE_STRING) => Another Post
blog::content (TYPE_STRING) => foo
{
  "blog": {
    "authorId": "john.days",
    "content": "foo",
    "id": "61806067a838e0152aa0fd16",
    "title": "Another Post"
  }
}

blog.BlogService@localhost:50051> call ListBlogs
{
  "blog": {
    "authorId": "vitorbari",
    "content": "this it the content of post 1",
    "id": "61806040a838e0152aa0fd15",
    "title": "My First Blog Post"
  }
}
{
  "blog": {
    "authorId": "john.days",
    "content": "foo",
    "id": "61806067a838e0152aa0fd16",
    "title": "Another Post"
  }
}

blog.BlogService@localhost:50051> call UpdateBlog
blog::id (TYPE_STRING) => 61806040a838e0152aa0fd15
blog::author_id (TYPE_STRING) => vitor.edited
blog::title (TYPE_STRING) => Changed title of post 1
blog::content (TYPE_STRING) => new content
{
  "blog": {
    "authorId": "vitor.edited",
    "content": "new content",
    "id": "61806040a838e0152aa0fd15",
    "title": "Changed title of post 1"
  }
}

blog.BlogService@localhost:50051> call ListBlogs
{
  "blog": {
    "authorId": "vitor.edited",
    "content": "new content",
    "id": "61806040a838e0152aa0fd15",
    "title": "Changed title of post 1"
  }
}
{
  "blog": {
    "authorId": "john.days",
    "content": "foo",
    "id": "61806067a838e0152aa0fd16",
    "title": "Another Post"
  }
}

blog.BlogService@localhost:50051> call DeleteBlog
blog_id (TYPE_STRING) => 61806067a838e0152aa0fd16
{
  "blogId": "61806067a838e0152aa0fd16"
}

blog.BlogService@localhost:50051> call ReadBlog
blog_id (TYPE_STRING) => 61806067a838e0152aa0fd16
command call: rpc error: code = NotFound desc = Blog not found: 61806067a838e0152aa0fd16
```


