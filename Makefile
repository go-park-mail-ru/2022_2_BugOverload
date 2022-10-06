.PHONY: all launch run_tests check_coverage

all: check build run_tests

TARGET = ./project/main.go
ARGS= :8088 ./configs/webserver.txt

PKG = ./...

clear:
	sudo rm -rf main
	sudo rm -rf coverage.html
	sudo rm -rf coverage.out
	sudo rm -rf c.out

create_env:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.49.0
	${GOPATH}/bin/golangci-lint

check:
	${GOPATH}/bin/golangci-lint run --config=linters_config/.golangci.yml
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

check_full_coverage:
	go test -race -coverpkg=./... -coverprofile=c.out ./... && go tool cover -func=c.out
