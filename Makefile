NAME := remote-file-to-s3-function

LDFLAGS  := -ldflags="-s -w"

.DEFAULT_GOAL := build-distributor build-downloader

export GO111MODULE=on

build-distributor:
	cd distributor; \
	go build $(LDFLAGS) -o ../bin/$(NAME)-distributor

build-downloader:
	cd downloader; \
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build $(LDFLAGS) -o ../bin/$(NAME)-downloader

test:
	go test -coverpkg=./... -coverprofile=coverage.txt -v ./...
