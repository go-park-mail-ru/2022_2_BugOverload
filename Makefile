.PHONY: all clear create_env check launch build launch_docker run_tests check_coverage check_full_coverage

all: check build run_tests

LINTERS_CONFIG = ./configs/.golangci.yml

PKG = ./...

clear:
	sudo rm -rf main coverage.html coverage.out c.out

create_env:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.49.0
	${GOPATH}/bin/golangci-lint

check:
	${GOPATH}/bin/golangci-lint run --config=${LINTERS_CONFIG}
	go fmt ${PKG}

launch_debug:
	go run ./cmd/debug/main.go --config-path ./cmd/debug/configs/config.toml

launch_prod:
	go run ./cmd/prod/main.go --config-path ./cmd/prod/configs/config.toml

build:
	go build cmd/debug/main.go ${TARGET}

launch_docker:
	sudo docker run -it --net=host -v "$(shell pwd):/project" --rm  andeo1812/golang_web

run_tests:
	go test -race ${PKG} -cover -coverpkg ${PKG}

get_coverage:
	go test ${PKG} -coverprofile coverage.out
	go tool cover -html coverage.out -o coverage.html

get_coverage_stat:
	go test -race -coverpkg=${PKG} -coverprofile=c.out ${PKG}
	go tool cover -func=c.out

create_doc:
	swag init --parseDependency --parseInternal --parseDepth 1 -g ./cmd/debug/main.go -o docs
