.PHONY: build
build:
	go build -v ./cmd/shortener

.PHONY: lint
lint:
	golangci-lint run ./...

.DEFAULT_GOAL := build