.PHONY: clean critic security lint test build run

APP_NAME = apiserver
BUILD_DIR = $(PWD)/build
MIGRATIONS_FOLDER = $(PWD)/platform/migrations
DATABASE_URL = mysql://root:123456@tcp(localhost:3306)/mysqldb?multiStatements=true

migrate.up:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" up

migrate.down:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" down


docker.fiber.build:
	docker build -t fiber .

docker.fiber: docker.fiber.build
	docker run --rm -d \
		--name cgapp-fiber \
		--network dev-network \
		-p 5000:5000 \
		fiber

docker.mysql:
	docker run -itd \
		--name go-mysql \
		-e MYSQL_ROOT_PASSWORD=123456 \
		-e MYSQL_USER=user \
		-e MYSQL_PASSWORD=password \
		-e MYSQL_DATABASE=mysqldb \
		-p 3306:3306 \
		mysql:8.0

docker.redis:
	docker run --rm -d \
		--name cgapp-redis \
		--network dev-network \
		-p 6379:6379 \
		redis

swag:
	swag init