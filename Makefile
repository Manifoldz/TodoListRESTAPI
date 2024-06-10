
PASSWORD=qwerty
HOST_PORT=5432

CONTAINER_NAME=todosrv-db

all:

build: 
	docker compose build todosrv

run:
	docker-compose up todosrv

pull-postgress:
	docker pull postgres

docker-run:
	docker run --name=$(CONTAINER_NAME) -e POSTGRES_PASSWORD=$(PASSWORD) -p $(HOST_PORT):5432 -d postgres

docker-stop:
	docker stop $(CONTAINER_NAME)

docker-start:
	docker start $(CONTAINER_NAME)

docker-rm: docker-stop
	docker rm $(CONTAINER_NAME)

migrate-up:
	migrate -path ./schema -database \
	'postgres://postgres:$(PASSWORD)@localhost:$(HOST_PORT)/postgres?sslmode=disable' up

migrate-down:
	migrate -path ./schema -database \
	'postgres://postgres:$(PASSWORD)@localhost:$(HOST_PORT)/postgres?sslmode=disable' down

connect:
	docker exec -it $(CONTAINER_NAME) /bin/bash

rebuild: migrate-down start
