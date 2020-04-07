NAME := remote-file-to-s3-function

LDFLAGS  := -ldflags="-s -w"

.DEFAULT_GOAL := build

export GO111MODULE=on

build: build-distributor build-downloader

build-distributor:
	cd distributor; \
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build $(LDFLAGS) -o ../bin/$(NAME)-distributor

build-downloader:
	cd downloader; \
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build $(LDFLAGS) -o ../bin/$(NAME)-downloader

test:
	go test -coverpkg=./... -coverprofile=coverage.txt -v ./...
