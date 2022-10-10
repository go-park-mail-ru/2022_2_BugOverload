.PHONY: all clear create_env check launch build launch_docker run_tests check_coverage check_full_coverage

all: check build run_tests

TARGET = ./cmd/defaultlaunch/main.go
ARGS= --config-path ./configs/config.toml
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

launch:
	go run ${TARGET} ${ARGS}

build:
	go build ${TARGET}

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
