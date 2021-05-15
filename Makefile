.PHONY: build
build:
		go build -o z_api -v .
		./z_api

.PHONY: test
test:
		go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build