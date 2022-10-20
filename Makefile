.PHONY: all clear generate-api-doc check launch build run-tests

all: check build run-all-tests

LINTERS_CONFIG = ./configs/.golangci.yml

PKG = ./...

SERVICE_APP = app
SERVICE_DYNAMODB_ADMIN = dynamodb-admin
SERVICE_LOCALSTACK =localstack

# develop
clear:
	sudo rm -rf main coverage.html coverage.out c.out *.log data

create-env-linters:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.49.0
	${GOPATH}/bin/golangci-lint

create-env-S3:
	./scripts/set_env_S3.sh

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

fill-S3-tests:
	./scripts/fill_test_data_S3.sh

fill-S3:
	./scripts/fill_data_S3.sh ${SOURCE}

# production
prod-mode:
	go run ./cmd/prod/main.go --config-path ./cmd/prod/configs/config.toml

# infrastructure
launch:
	docker-compose up --remove-orphans &
	sleep 6
	make fill-S3

stop:
	docker-compose kill
	docker-compose down

compose-log:
	docker-compose logs -f

app-restart:
	docker-compose restart  $(SERVICE_APP)

S3-restart:
	docker-compose restart  $(SERVICE_LOCALSTACK)

# OLD
docker-launch:
	docker run -it --net=host -v "$(shell pwd):/project" --rm  andeo1812/golang_web
