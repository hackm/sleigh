# Meta Info
NAME := sleigh
VERSION := v0.0.1
GO_PACKAGES = $(shell go list ./... | grep -v vendor)
GO_FILES = $(shell find . -name "*.go" | grep -v vendor | uniq)

## Setup

## Test
test:
	go test $(GO_PACKAGES)

## Lint
lint:
	go vet $(GO_PACKAGES)

.PHONY: test lint
