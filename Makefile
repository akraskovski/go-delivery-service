.PHONY: lint
lint:
	$(go env GOPATH)/bin/golangci-lint run

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux go build -o main -v ./cmd/deliveryservice

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build
