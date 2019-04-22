.PHONY: all clean goget env en-stop do-test test build image run logs stop

NAME    = goteca
VERSION = 1.0.0

all: image

clean:
	rm -f $(NAME)

goget:
	go get ./...

env: ## Run environment that will be need to test
	docker-compose up -d

env-stop: ## Kill test envinronment
	-docker-compose down -v

do-test: ## Execute all tests
	MONGO_URL=localhost:27017 \
	go test -p=1 $$(go list ./... | grep -v /vendor/)

test: env-stop env do-test env-stop ## Run the envinroment, do tests then kill the environment

build: clean test
	CGO_ENABLED=0 go build -v -a -installsuffix cgo -o $(NAME) ./cmd/server

image: build
	docker build -t=$(NAME):$(VERSION) .

run:
	docker run -d --name $(NAME) -v MONGO_URL="127.0.0.1:27017" -p 8080:8080 $(NAME):$(VERSION)

logs:
	docker logs -f $(NAME)

stop:
	docker rm -vf $(NAME)