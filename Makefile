.PHONY: all clear run-linter run-tests build

all: run-linter run-tests build

LINTER_CFG = ./configs/.golangci.yml

PKG = ./internal/... ./pkg/...
PKG_INTERNAL = ./internal/...

PKG_LINTERS = ./...

MICROSERVICE_DIR=$(PWD)/internal

create-env:
	go mod download

run-linter:
	$(GOPATH)/bin/golangci-lint run $(PKG_LINTERS) --config=$(LINTER_CFG)
	go fmt $(PKG)

build:
	rm -f cmd/api/api_bin cmd/image/image_bin cmd/warehouse/warehouse_bin cmd/auth/auth_bin cmd/filldb/filldb_bin
	go build cmd/api/main.go
	mv main cmd/api/api_bin
	go build cmd/image/main.go
	mv main cmd/image/image_bin
	go build cmd/warehouse/main.go
	mv main cmd/warehouse/warehouse_bin
	go build cmd/auth/main.go
	mv main cmd/auth/auth_bin
	go build cmd/filldb/main.go
	mv main cmd/filldb/filldb_bin

run-tests:
	go test -race $(PKG_INTERNAL) -cover -coverpkg $(PKG_INTERNAL)

get-stat-coverage:
	go test -race -coverpkg=${PKG_INTERNAL} -coverprofile=c.out ${PKG_INTERNAL}
	cat c.out | fgrep -v "easyjson" | fgrep -v "mock" | fgrep -v "dev" | fgrep -v "test.go" | fgrep -v "docs" |  fgrep -v "testing.go" | fgrep -v ".pb.go" | fgrep -v "config" > c2.out
	go tool cover -func=c2.out
	go tool cover -html c2.out -o coverage.html

api-doc-generate:
	swag init --parseDependency --parseInternal --parseDepth 1 -g ./cmd/api/main.go -o docs

# mocks, easyjson
generate:
	go generate ${PKG}

proto-generate:
	protoc --proto_path=${MICROSERVICE_DIR}/image/delivery/grpc/protobuf image.proto --go_out=plugins=grpc:${MICROSERVICE_DIR}/image/delivery/grpc/protobuf
	protoc --proto_path=${MICROSERVICE_DIR}/warehouse/delivery/grpc/protobuf warehouse.proto --go_out=plugins=grpc:${MICROSERVICE_DIR}/warehouse/delivery/grpc/protobuf
	protoc --proto_path=${MICROSERVICE_DIR}/auth/delivery/grpc/protobuf auth.proto --go_out=plugins=grpc:${MICROSERVICE_DIR}/auth/delivery/grpc/protobuf


# S3
IMAGES = /home/webapps/images
S3_ENDPOINT = http://localhost:4566

# Example: make fill-S3-slow IMAGES=/home/webapps/images S3_ENDPOINT=http://localhost:4566
fill-S3-slow:
	./scripts/fill_data_S3.sh ${IMAGES} ${S3_ENDPOINT}

fill-S3-fast:
	./scripts/fill_data_S3_fast.sh ${IMAGES} ${S3_ENDPOINT}


# Migrations
MIGRATIONS_DIR = scripts/migrations
DB_URL := $(shell echo postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=${POSTGRES_SSLMODE})

fill-db: export POSTGRES_HOST=localhost
fill-db:
	./cmd/filldb/filldb_bin --config-path=./cmd/filldb/configs/prod.toml --data-path=./test/newdata

# Production BEGIN ----------------------------------------------------

prod-create-env:
	sudo cp /etc/letsencrypt/live/movie-gate.ru/fullchain.pem ./cmd/api/
	sudo cp /etc/letsencrypt/live/movie-gate.ru/privkey.pem ./cmd/api/

prod-restart:
	docker-compose -f docker-compose.production.yml restart warehouse
	docker-compose -f docker-compose.production.yml restart image
	docker-compose -f docker-compose.production.yml restart auth
	docker-compose -f docker-compose.production.yml restart api

prod-deploy:
	mkdir -p --mode=777 logs/database/main
	make prod-create-env
	docker-compose -f docker-compose.production.yml up -d main_db admin_db localstack
	sleep 5
	make fill-db
	docker-compose -f docker-compose.production.yml up -d image warehouse auth api
	docker-compose -f docker-compose.production.yml up -d monitor_db prometheus node_exporter grafana
	sleep 5
	make fill-S3-slow IMAGES=/home/webapps/images S3_ENDPOINT=http://localhost:4566

# Production END ------------------------------------------------------


stop:
	docker-compose kill
	docker-compose down


# Debug BEGIN ---------------------------------------------------------
#IMAGES=/home/andeo/Загрузки/images S3_ENDPOINT=http://localhost:4566

debug-fill-db: export POSTGRES_HOST=localhost
debug-fill-db:
	./cmd/filldb/filldb_bin --config-path=./cmd/filldb/configs/debug.toml --data-path=./test/newdata

debug-base-deploy:
	docker-compose up -d main_db admin_db localstack

debug-app-deploy:
	docker-compose up -d image warehouse auth api

debug-monitoring-deploy:
	docker-compose up -d monitor_db prometheus node_exporter grafana

debug-deploy:
	make clear
	mkdir -p --mode=777 logs/database/main
	make debug-base-deploy
	sleep 1
	make build
	make debug-app-deploy
	make debug-monitoring-deploy

debug-app-restart:
	docker-compose restart warehouse
	docker-compose restart image
	docker-compose restart auth
	docker-compose restart api

run-perf-tests:
	go run ./perf_test/perftest.go

# Debug END -----------------------------------------------------------

send-bin:
	scp ./cmd/filldb/filldb_bin root@95.163.242.214:/home/webapps/movie-gate.ru/backend/2022_2_BugOverload/cmd/filldb/filldb_bin
	scp ./cmd/api/api_bin root@95.163.242.214:/home/webapps/movie-gate.ru/backend/2022_2_BugOverload/cmd/api/api_bin
	scp ./cmd/image/image_bin root@95.163.242.214:/home/webapps/movie-gate.ru/backend/2022_2_BugOverload/cmd/image/image_bin
	scp ./cmd/warehouse/warehouse_bin root@95.163.242.214:/home/webapps/movie-gate.ru/backend/2022_2_BugOverload/cmd/warehouse/warehouse_bin
	scp ./cmd/auth/auth_bin root@95.163.242.214:/home/webapps/movie-gate.ru/backend/2022_2_BugOverload/cmd/auth/auth_bin

# Debug END -----------------------------------------------------------

clear:
	sudo rm -rf main coverage.html coverage.out c.out *.log logs/ c2.out fullchain.pem privkey.pem cmd/*/*_bin
