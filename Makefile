CURL_BIN ?= curl
GO_BIN ?= go
LINT_BIN ?= golangci-lint
GORELEASER_BIN ?= goreleaser

PUBLISH_PARAM?=
GO_MOD_PARAM?=
TMP_DIR=?/tmp

BASE_DIR=$(shell pwd)

TMP_DIR?=./tmp
BASE_DIR=$(shell pwd)
MAKEFILE_ABSPATH := $(CURDIR)/$(word $(words $(MAKEFILE_LIST)),$(MAKEFILE_LIST))
MAKEFILE_RELPATH := $(call MAKEFILE_ABSPATH)

export PATH := $(BASE_DIR)/bin:$(PATH)

.PHONY: help
help: ## print help message
	@echo "Usage: make <command>"
	@echo
	@echo "Available commands are:"
	@grep -E '^\S[^:]*:.*?## .*$$' $(MAKEFILE_RELPATH) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-4s\033[36m%-30s\033[0m %s\n", "", $$1, $$2}'
	@echo

.PHONY: build
build: ## build abbreviate
	$(GO_BIN) build $(GO_MOD_PARAM)

.PHONY: clean
clean: ## remove build artifacts
	rm -f abbreviate

.PHONY: clean-deps
clean-deps: ## remove dependencies
	rm -rf ./bin
	rm -rf ./tmp
	rm -rf ./share

.PHONY: test-deps
test-deps: ./bin/godog ## set up test dependencies
	$(CURL_BIN) -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b ./bin v1.21.0
	$(GO_BIN) get -t ./...

./bin:
	mkdir ./bin

./tmp:
	mkdir ./tmp

./bin/goreleaser: ./bin ./tmp
	$(CURL_BIN) --fail -L -o ./tmp/goreleaser.tar.gz https://github.com/goreleaser/goreleaser/releases/download/v0.117.2/goreleaser_Linux_x86_64.tar.gz
	gunzip -f ./tmp/goreleaser.tar.gz
	tar -C ./bin -xvf ./tmp/goreleaser.tar

.PHONY: build-deps
build-deps: ./bin/goreleaser ## set up build depenencies

deps: build-deps test-deps

.PHONY: test
test: ## run unit tests
	$(GO_BIN) test $(GO_MOD_PARAM) ./...

.PHONY: acceptance-test
acceptance-test: ## run acceptance tests
	cd test && go test -tags acceptance

.PHONY: ci-test
ci-test: ## run unit tests for CI
	$(GO_BIN) test $(GO_MOD_PARAM) -race -coverprofile=coverage.txt -covermode=atomic ./...

.PHONY: lint
lint: ## run linter
	$(LINT_BIN) run

.PHONY: release
release: clean ## create and publish release artifact
	$(GORELEASER_BIN) $(PUBLISH_PARAM)

.PHONY: update
update: ## update Go dependencies
	$(GO_BIN) get -u
ifeq ($(GO111MODULE),on)
	$(GO_BIN) mod tidy
endif
	make test
	make install
ifeq ($(GO111MODULE),on)
	$(GO_BIN) mod tidy
endif

