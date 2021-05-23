.PHONY: build
build:
		go build -o z_api -v ./
		./z_api

.PHONY: test
test:
#go test -v -race -timeout 30s ./...
		go test -v -race -timeout 30s ./... | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''

.DEFAULT_GOAL := build