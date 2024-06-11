
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

test: post-createLists post-createItems put-doneItems

# Создать списки
post-createLists:
	curl -i -X POST localhost:8000/api/lists/ \
	-H "Content-Type: application/json" \
	-d '{"title": "Покупка продуктов", "description": "На праздник"}';
	curl -i -X POST localhost:8000/api/lists/ \
	-H "Content-Type: application/json" \
	-d '{"title": "Список домашнего задания", "description": "По курсам"}';
	curl -i -X POST localhost:8000/api/lists/ \
	-H "Content-Type: application/json" \
	-d '{"title": "Встречи"}'


# Создать задания
post-createItems:
	curl -i -X POST localhost:8000/api/lists/1/items/ \
	-H "Content-Type: application/json" \
	-d '{"title": "Хлеб", "description": "свежий"}';
	curl -i -X POST localhost:8000/api/lists/1/items/ \
	-H "Content-Type: application/json" \
	-d '{"title": "Яблоки", "description": "2кг"}';
	curl -i -X POST localhost:8000/api/lists/1/items/ \
	-H "Content-Type: application/json" \
	-d '{"title": "Молоко", "description": "3,2%"}';
	curl -i -X POST localhost:8000/api/lists/2/items/ \
	-H "Content-Type: application/json" \
	-d '{"title": "Математика", "description": "2 задачи"}';
	curl -i -X POST localhost:8000/api/lists/2/items/ \
	-H "Content-Type: application/json" \
	-d '{"title": "Информатика"}';
	curl -i -X POST localhost:8000/api/lists/3/items/ \
	-H "Content-Type: application/json" \
	-d '{"title": "Обсудить задачи", "description": "C отделом планирования"}';
	curl -i -X POST localhost:8000/api/lists/3/items/ \
	-H "Content-Type: application/json" \
	-d '{"title": "С другом"}'

# Отметить выполнено
put-doneItems:
	curl -i -X PUT localhost:8000/api/items/1 \
	-H "Content-Type: application/json" \
	-d '{"done": true, "description": "вчерашний"}';
	curl -i -X PUT localhost:8000/api/items/4 \
	-H "Content-Type: application/json" \
	-d '{"done": true}';
	curl -i -X PUT localhost:8000/api/items/7 \
	-H "Content-Type: application/json" \
	-d '{"done": true}';


# Удалить несколько задач
delete-Items:
	curl -i -X DELETE localhost:8000/api/items/2 \
	-H "Content-Type: application/json";
	curl -i -X DELETE localhost:8000/api/items/4 \
	-H "Content-Type: application/json";

# Запрос всех листа 1
get-AllItems1:
	curl -i -X GET localhost:8000/api/lists/1/items/ \
	-H "Content-Type: application/json";

# Запрос всех листа 1 с фильтрацией выполнено
get-AllItems1Done:
	curl -i -X GET localhost:8000/api/lists/1/items/?done=true \
	-H "Content-Type: application/json";

# Запрос всех листа 1 с фильтрацией не выполнено и пагинацией
get-AllItems1DonePag:
	curl -i -X GET "localhost:8000/api/lists/1/items/?done=false&limit=1&offset=0" \
	-H "Content-Type: application/json";
