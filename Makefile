.PHONY: build run test cover
build:
	mkdir -p ./build
	go build -v -o build/service ./cmd/service

run:
	go run ./cmd/service/main.go

test:
	go test -v -race -coverprofile=cover.out ./...

cover:
	go tool cover -html=cover.out

.DEFAULT_GOAL := build
