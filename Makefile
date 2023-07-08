REPOSITORY="superm4n"
PROJECT="book-server-backend"
TAG="v0.0.1"
CONTAINER_PORT=8080

build:
	docker build . -t ${REPOSITORY}/${PROJECT}:${TAG}

push:
	docker push ${REPOSITORY}/${PROJECT}:${TAG}

run:
	docker run -d -p ${CONTAINER_PORT}:${CONTAINER_PORT} ${REPOSITORY}/${PROJECT}:${TAG}
