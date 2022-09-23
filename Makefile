.PHONY: all launch run_tests check_coverage

all: launch

TARGET = ./project/main.go

launch:
	go run ${TARGET} ${ARGS}

build:
	go build ${TARGET}

launch_race:
	go run -race ${TARGET}

run_tests:
	go test -race ./... -cover -coverpkg ./...

check_coverage:
	go test ./... -coverprofile coverage.out
	go tool cover -html coverage.out -o coverage.html
