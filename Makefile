NAME := xlapse

LDFLAGS  := -ldflags="-s -w"

.DEFAULT_GOAL := build

export GO111MODULE=on

build: build-distributor build-downloader build-gif-distributor build-gif-maker

build-distributor:
	cd function/distributor; \
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build $(LDFLAGS) -o ../../bin/$(NAME)-distributor

build-downloader:
	cd function/downloader; \
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build $(LDFLAGS) -o ../../bin/$(NAME)-downloader

build-gif-distributor:
	cd function/gif-distributor; \
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build $(LDFLAGS) -o ../../bin/$(NAME)-gif-distributor

build-gif-maker:
	cd function/gif-maker; \
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build $(LDFLAGS) -o ../../bin/$(NAME)-gif-maker

protoc-go:
	protoc --go_out=paths=source_relative:. types/v1/*.proto

test:
	go test -coverpkg=./... -coverprofile=coverage.txt -v ./...
