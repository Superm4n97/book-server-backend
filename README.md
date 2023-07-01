# book-server-backend

## Database
* Create a mongodb database naming `book-server`
* Create the following collections
  * book
  * author

## Run server
first export the mongodb cluster uri
```bash
export MONGODB_URI="mongodb+srv://<username>:<password>@<mongo_cluster>/?retryWrites=true&w=majority"
go run main.go
```