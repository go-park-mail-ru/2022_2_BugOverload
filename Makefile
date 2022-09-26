.PHONY: all launch run_tests check_coverage

all: check launch

TARGET = ./project/main.go
ARGS= ./configs/webserver.txt

PKG = ./...

clear:
	sudo rm -rf main

create_env:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.49.0
	${GOPATH}/bin/golangci-lint

check:
	golangci-lint run
	go fmt ${PKG}

launch:
	go run ${TARGET} ${ARGS}

build:
	go build ${TARGET}

docker_launch:
	sudo docker run -it --net=host -v "$(shell pwd):/project" --rm  andeo1812/golang_web


launch_race:
	go run -race ${TARGET}

run_tests:
	go test -race ./... -cover -coverpkg ./...

check_coverage:
	go test ./... -coverprofile coverage.out
	go tool cover -html coverage.out -o coverage.html
