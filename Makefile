.PHONY: build
build:
	go build -v ./cmd/delivery_service

.DEFAULT_GOAL := build
