.PHONY: all clean goget test build image

NAME    = goteca
VERSION = 1.0.0

all: image

clean:
	rm -f $(NAME)

goget:
	go get ./...

test: goget
	go test ./...

build: clean test
	CGO_ENABLED=0 go build -v -a -installsuffix cgo -o $(NAME) .

image: build
	docker build -t=$(NAME):$(VERSION) .