# Docker build
* First change the db cluster url
  * Go to `Dockerfile` then set `ENV` variable as `MONGODB_URI`
* Build docker container
    ```bash
    docker build . -t superm4n/book-server-backend:v0.0.1
    ```
  You can rename the `container name` and `tag name` for the server
* You can push this container to docker container repository by applying
  ```bash
  docker push superm4n/book-server-backend:v0.0.1    
  ```
  You have to specify the correct image name that you have build in the previous step.