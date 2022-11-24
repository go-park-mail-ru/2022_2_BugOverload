.PHONY: all clear generate-api-doc check launch build run-tests migrate-debug-up DB_URL

all: check build run-tests

LINTERS_CONFIG = ./configs/.golangci.yml

PKG = ./...

SERVICE_MAIN = main

MICROSERVICE_DIR=$(PWD)/internal

# ENV
create-env-lint:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.49.0
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.49.0
	${GOPATH}/bin/golangci-lint

migrate-install:
	curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz
	sudo mv migrate /bin
	rm README.md LICENSE

env:
	export AWS_REGION=us-east-1 && export AWS_PROFILE=default && export AWS_ACCESS_KEY_ID=foo &&
	export AWS_SECRET_ACCESS_KEY=bar && export POSTGRES_HOST=main_db &&
	export POSTGRES_DB=mgdb && export POSTGRES_USER=mguser && export POSTGRES_PASSWORD=mgpass &&
	export POSTGRES_PORT=5432 && export POSTGRES_SSLMODE=disable && export POSTGRES_HOST_DEV=localhost

# Development
check:
	golangci-lint run --config=${LINTERS_CONFIG}
	go fmt ${PKG}

debug-mode:
	go run ./cmd/debug/main.go --config-path ./cmd/debug/configs/config.toml

image-service-launch:
	go run ./cmd/image/main.go --config-path ./cmd/image/configs/config.toml

build:
	go build cmd/debug/main.go ${TARGET}

run-all-tests:
	go test -race ${PKG} -cover -coverpkg ${PKG}

run-tests:
	go test -race ${PKG} -cover -coverpkg $(PKG)

get-stat-coverage:
	go test -race -coverpkg=${PKG} -coverprofile=c.out ${PKG}
	cat c.out | fgrep -v "easyjson" | fgrep -v "mock" | fgrep -v "dev" | fgrep -v "test.go" | fgrep -v "docs" |  fgrep -v "testing.go" | fgrep -v ".pb.go" | fgrep -v "config" > c2.out
	go tool cover -func=c2.out
	go tool cover -html c2.out -o coverage.html

api-doc-generate:
	swag init --parseDependency --parseInternal --parseDepth 1 -g ./cmd/debug/main.go -o docs

mocks-generate:
	go generate ${PKG}

proto-generate:
	protoc --proto_path=${MICROSERVICE_DIR}/image/delivery/grpc/protobuf image.proto --go_out=plugins=grpc:${MICROSERVICE_DIR}/image/delivery/grpc/protobuf

# Example: make fill-S3 IMAGES=/home/andeo/Загрузки/images S3_ENDPOINT=http://localhost:4566
fill-S3-slow:
	./scripts/fill_data_S3.sh ${IMAGES} ${S3_ENDPOINT} &

fill-S3-fast:
	./scripts/fill_data_S3_fast.sh ${IMAGES} ${S3_ENDPOINT} &

dev-fill-db:
	go run ./cmd/filldb/main.go --config-path ./cmd/filldb/configs/config.toml --data-path ./test/newdata

# production
prod-mode:
	go run ./cmd/prod/main.go --config-path ./cmd/prod/configs/config.toml

# infrastructure
# Example: make prod-deploy IMAGES=/home/andeo/Загрузки/images S3_ENDPOINT=http://localhost:4566
prod-create-env:
	sudo cp /etc/letsencrypt/live/movie-gate.online/fullchain.pem .
	sudo cp /etc/letsencrypt/live/movie-gate.online/privkey.pem .

prod-deploy:
	make prod-create-env
	docker-compose -f docker-compose.yml -f docker-compose.production.yml up -d
	sleep 2
	make reboot-db-debug
	sleep 30
	make fill-S3-slow ${IMAGES} ${S3_ENDPOINT}

debug-deploy:
	docker-compose up -d
	sleep 1
	make reboot-db-debug
	sleep 1
	make fill-S3-fast ${IMAGES} ${S3_ENDPOINT}

stop:
	docker-compose kill
	docker-compose down

logs:
	docker-compose logs -f

reboot-db-debug:
	docker-compose exec $(SERVICE_MAIN) make -C project  reboot-db COUNT=3

main-debug-restart:
	docker-compose restart $(SERVICE_MAIN)

main-prod-restart:
	make prod-create-env
	docker-compose -f docker-compose.yml -f docker-compose.production.yml restart $(SERVICE_MAIN)

# Example: make infro-command COMMAND=run-all-tests
infro-command:
	docker-compose exec $(SERVICE_MAIN) make -C project  ${COMMAND}

# Migrations
MIGRATIONS_DIR = scripts/migrations
DB_URL := $(shell echo postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=$(POSTGRES_SSLMODE))

migrate-up:
	migrate -source file://${MIGRATIONS_DIR} -database ${DB_URL} up ${COUNT}

migrate-down:
	migrate -source file://${MIGRATIONS_DIR} -database ${DB_URL} down ${COUNT}

migrate-force:
	migrate -source file://${MIGRATIONS_DIR} -database ${DB_URL} force ${COUNT}

reboot-db:
	echo 'y' | migrate -source file://${MIGRATIONS_DIR}  -database ${DB_URL} down
	migrate -source file://${MIGRATIONS_DIR}  -database ${DB_URL} up ${COUNT}
	make dev-fill-db

# Utils
clear:
	sudo rm -rf main coverage.html coverage.out c.out *.log logs/ c2.out fullchain.pem privkey.pem

open-last-log:
	./scripts/print_last_log.sh

# Example: make set-format TARGET=/home/andeo/Загрузки/images FORMAT=jpeg
set-format:
	./scripts/set_format.sh ${TARGET} ${FORMAT}
