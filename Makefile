.PHONY: all clear generate-api-doc check launch build run-tests

all: check build run-all-tests

LINTERS_CONFIG = ./configs/.golangci.yml

PKG = ./...

SERVICE_MAIN = main

# develop
clear:
	sudo rm -rf main coverage.html coverage.out c.out *.log data

create-env:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.49.0
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.49.0
	${GOPATH}/bin/golangci-lint
	pip install -y migrate

env:
	export AWS_REGION=us-east-1 && export AWS_PROFILE=default && export AWS_ACCESS_KEY_ID=foo && export AWS_SECRET_ACCESS_KEY=bar &&
	export POSTGRES_DB=mgdb && export POSTGRES_USER=mguser && export POSTGRES_PASSWORD=mgpass &&
	export POSTGRESQL_URL='postgres://mguser:mgpass@localhost:5432/mgdb?sslmode=disable'

check:
	${GOPATH}/bin/golangci-lint run --config=${LINTERS_CONFIG}
	go fmt ${PKG}

debug-mode:
	go run ./cmd/debug/main.go --config-path ./cmd/debug/configs/config.toml

build:
	go build cmd/debug/main.go ${TARGET}

run-all-tests:
	go test -race ${PKG} -cover -coverpkg ${PKG}

run-tests:
	go test -race ./cmd/debug/tests/ -cover -coverpkg $(PKG)

get-coverage:
	go test ${PKG} -coverprofile coverage.out
	go tool cover -html coverage.out -o coverage.html

get-stat-coverage:
	go test -race -coverpkg=${PKG} -coverprofile=c.out ${PKG}
	go tool cover -func=c.out

generate-api-doc:
	swag init --parseDependency --parseInternal --parseDepth 1 -g ./cmd/debug/main.go -o docs

# Example: make fill-S3 IMAGES=/home/andeo/Загрузки/images S3_ENDPOINT=http://localhost:4566
fill-S3:
	./scripts/fill_data_S3.sh ${IMAGES} ${S3_ENDPOINT} &

# Example: make set-format TARGET=/home/andeo/Загрузки/images FORMAT=jpeg
set-format:
	./scripts/set_format.sh ${TARGET} ${FORMAT}

open-last-log:
	./scripts/print_last_log.sh

# production
prod-mode:
	go run ./cmd/prod/main.go --config-path ./cmd/prod/configs/config.toml

# infrastructure
# Example: make prod-deploy IMAGES=/home/andeo/Загрузки/images S3_ENDPOINT=http://localhost:4566
prod-deploy:
	docker-compose -f docker-compose.yml -f docker-compose.production.yml up -d
	sleep 30
	make fill-S3 ${IMAGES} ${S3_ENDPOINT}

debug-deploy:
	docker-compose up -d
	sleep 30
	make fill-S3 ${IMAGES} ${S3_ENDPOINT}

stop:
	docker-compose kill
	docker-compose down

logs:
	docker-compose logs -f

main-debug-restart:
	docker-compose restart $(SERVICE_MAIN)

main-prod-restart:
	docker-compose -f docker-compose.yml -f docker-compose.production.yml restart $(SERVICE_MAIN)


# OLD
docker-launch:
	docker run -it --net=host -v "$(shell pwd):/project" --rm  andeo1812/golang_web
