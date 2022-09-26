.PHONY: all launch run_tests check_coverage

all: launch

TARGET = ./project/main.go

PKG = ./...

create_env:
	go get -u golang.org/x/lint/golint
	go get -u honnef.co/go/tools/cmd/staticcheck
	go get -u github.com/kisielk/errcheck

check:
	golint ${PKG}
	go vet ${PKG}
	staticcheck ${PKG}
	errcheck ${PKG}

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
