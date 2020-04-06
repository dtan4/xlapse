NAME    := s3url

LDFLAGS  := -ldflags="-s -w"

.DEFAULT_GOAL := bin/$(NAME)

export GO111MODULE=on

bin/$(NAME):
	go build $(LDFLAGS) -o bin/$(NAME)

test:
	go test -coverpkg=./... -v ./...
