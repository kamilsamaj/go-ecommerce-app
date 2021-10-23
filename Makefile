BINARY_NAME="accolade-cli"
XC_OS ?= linux darwin windows
XC_ARCH ?= amd64 arm64
BIN="./bin"
VERSION ?= "local"
BUILD_TIME=$(shell date -u "+%FT%H:%M:%SZ")

.PHONY: all build test clean

default: all

all: test build

test:
	go test -v -cover ./...

check:
	pre-commit run --all-files

gocritic:
	gocritic check -enableAll ./...

build:
	@for OS in $(XC_OS); do \
		for ARCH in $(XC_ARCH); do \
			echo Building $$OS/$$ARCH to $(BIN)/$(BINARY_NAME)-$$OS-$$ARCH ; \
			CGO_ENABLED=0 \
			GOOS=$$OS \
			GOARCH=$$ARCH \
			go build \
			-ldflags="-s -w -X 'github.com/konciergeMD/$(BINARY_NAME)/cmd.version=$(VERSION)' -X 'github.com/konciergeMD/$(BINARY_NAME)/cmd.buildTime=$(BUILD_TIME)'" \
			-o=$(BIN)/$(BINARY_NAME)-$$OS-$$ARCH ; \
		done ; \
	done

clean:
	rm -rfv $(BIN)
