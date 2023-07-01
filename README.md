# book-server-backend

## Run server
first export the mongodb cluster uri
```bash
export MONGODB_URI="mongodb+srv://<username>:<password>@<mongo_cluster>/?retryWrites=true&w=majority"
go run main.go
```

## APIs
* `[POST]`:`/apis/v1/author` - Adds a author to the database. The request body contains 
```json
{
    "name": "<author_name>",
    "email": "<author_email>"
}
```
* `[GET]`:`apis/v1/author/{name}` - Gets the author's information. The author is identifies by the `name`.

## Architecture
![Archetecture](./docs/resource/archetecture.jpg)