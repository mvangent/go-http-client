.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	staticcheck ./...
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

test: vet
	go test -v ./... -covermode=count -coverprofile=coverage.out
.PHONY:test		

test_report: test
	go tool cover -func=coverage.out -o=coverage.out
.PHONY:test_report

build: vet
	go build ./...
.PHONY:build
