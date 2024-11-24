BIN := bin/umlcoder
VERSION := $(shell git describe --tags --abbrev=0)
LDFLAGS := "-s -w -X main.version=$(VERSION)"
export GO111MODULE=on

all: test clean build

.PHONY: build
build:
	go build -trimpath -ldflags=$(LDFLAGS) -o $(BIN) ./cmd/umlcoder

.PHONY: install
install:
	go install -trimpath -ldflags=$(LDFLAGS) ./cmd/umlcoder

.PHONY: clean
clean:
	rm -rf $(BIN)
	go clean

.PHONY: test
test:
	go test -race ./... -count=1
