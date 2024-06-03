.DEFAULT_GOAL := build-all

export PROJECT := "us-soccer-go"
export PACKAGE := "us-soccer-go-test"
export COMPOSE_PROJECT := "app"

VERSION=$(shell git describe --tags --always --abbrev=0 --match=v* 2> /dev/null | sed -r "s:^v::g" || echo 0)
VERSION_FULL=$(shell git describe --tags --always --dirty --match=v* 2> /dev/null | sed -r "s:^v::g" || echo 0)

# general
build-all: go-fetch go-build
	@echo

up: go-upgrade-deps
	@echo


docker:
	docker compose \
		--project-name ${COMPOSE_PROJECT} \
		--file docker-compose.yml \
		up \
		--remove-orphans \
		--build \
		--timeout 0 ${COMPOSE_ARGS}

docker-clean:
	docker compose \
		--project-name ${COMPOSE_PROJECT} \
		--file docker-compose.yml \
		down \
		--volumes \
		--remove-orphans \
		--rmi local --timeout 1

docker-build:
	docker build \
		--tag ${PROJECT} \
		--force-rm .


# backend
go-prepare: go-fetch
	go generate -x ./...

go-fetch:
	go mod download
	go mod tidy

go-upgrade-deps:
	go get -u ./...
	go mod tidy

go-upgrade-deps-patch:
	go get -u=patch ./...
	go mod tidy

go-dlv: go-prepare
	dlv debug \
		--headless --listen=:2345 \
		--api-version=2 --log \
		--allow-non-terminal-interactive \
		${PACKAGE} -- --debug

go-debug: go-prepare
	go run ${PACKAGE} --debug

go-debug-fast:
	go run ${PACKAGE} --debug

go-build: go-prepare go-fetch
	CGO_ENABLED=1 \
	go build \
		-ldflags '-s -w -extldflags=-static' \
		-tags=netgo,osusergo,static_build \
		-installsuffix netgo \
		-buildvcs=false \
		-trimpath \
		-o ${PROJECT} \
		${PACKAGE}