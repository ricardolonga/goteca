.PHONY: all clean goget env env-ip en-stop do-test test build image run logs stop

NAME    = goteca
VERSION = 1.0.0

all: image

clean:
	rm -f $(NAME)

goget:
	go get ./...

env:
	docker-compose up -d

env-ip:
	echo $$(docker inspect -f '{{.NetworkSettings.Networks.goteca_default.IPAddress}}' goteca_mongodb_1)

env-stop:
	-docker-compose down -v

do-test:
	MONGO_URL=$$(docker inspect -f '{{.NetworkSettings.Networks.goteca_default.IPAddress}}' goteca_mongodb_1) \
	go test $$(go list ./... | grep -v vendor)

test: env-stop env do-test env-stop ## Run the envinroment, do tests then kill the environment

build: clean test
	CGO_ENABLED=0 go build -v -a -installsuffix cgo -o $(NAME) .

image: build
	docker build -t=$(NAME):$(VERSION) .

run:
	docker run -d --name $(NAME) -e MONGO_URL="172.20.0.2:27017" -p 8080:8080 $(NAME):$(VERSION)

logs:
	docker logs -f $(NAME)

stop:
	docker rm -vf $(NAME)