BIN := "./bin/calendar"
DOCKER_IMG="calendar:develop"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

build: build_calendar build_scheduler build_sender

build_calendar:
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd/calendar

build_scheduler:
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd/calendar_scheduler

build_sender:
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd/calendar_sender

run: build
	$(BIN) -config ./configs/config.toml

db:
	docker-compose up -d postgresql

pull:
	docker-compose pull

docker-build:
	docker-compose build --pull

up:
	docker-compose up -d

down:
	docker-compose down --remove-orphans

clear:
	docker-compose down -v --remove-orphans

build-img:
	docker build \
		--build-arg=LDFLAGS="$(LDFLAGS)" \
		-t $(DOCKER_IMG) \
		-f build/Dockerfile .

run-img: build-img
	docker run $(DOCKER_IMG)

version: build
	$(BIN) version

test:
	go test -race ./internal/... #./pkg/...

remove-lint-deps:
	rm $(which golangci-lint)

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.50.1

lint: install-lint-deps
	golangci-lint run ./...

integration-tests: integration-tests-run integration-tests-clear

integration-tests-run:
	docker compose -f ./docker-compose.test.yml run integration_tests go test

integration-tests-clear:
	docker compose -f ./docker-compose.test.yml down -v --remove-orphans

generate:
	go generate ./api

.PHONY: build run build-img run-img version test lint
