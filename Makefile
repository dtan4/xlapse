NAME := image-to-s3-function

LDFLAGS  := -ldflags="-s -w"

.DEFAULT_GOAL := build

export GO111MODULE=on

build:
	go build $(LDFLAGS) -o bin/$(NAME)

test:
	go test -coverpkg=./... -v ./...
