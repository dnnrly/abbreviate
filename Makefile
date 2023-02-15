CURL_BIN ?= curl
GO_BIN ?= go
LINT_BIN ?= golangci-lint
GORELEASER_BIN ?= goreleaser

PUBLISH_PARAM?=
GO_MOD_PARAM?=
TMP_DIR=?/tmp

BASE_DIR=$(shell pwd)

export PATH := $(BASE_DIR)/bin:$(PATH)

.PHONY: install
install: deps

.PHONY: build
build:
	$(GO_BIN) build $(GO_MOD_PARAM)

.PHONY: clean
clean:
	rm -f abbreviate

.PHONY: clean-deps
clean-deps:
	rm -rf ./bin
	rm -rf ./tmp
	rm -rf ./share

test-deps: ./bin/godog
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

./bin/godog: ./bin ./tmp
	curl --fail -L -o ./tmp/godog.tar.gz https://github.com/cucumber/godog/releases/download/v0.11.0/godog-v0.11.0-linux-amd64.tar.gz
	tar -xf ./tmp/godog.tar.gz -C ./tmp
	cp ./tmp/godog-v0.11.0-linux-amd64/godog ./bin

build-deps: ./bin/goreleaser

deps: build-deps test-deps

.PHONY: test
test:
	$(GO_BIN) test $(GO_MOD_PARAM) ./...

.PHONY: acceptance-test
acceptance-test:
	cd test && godog -t @Acceptance

.PHONY: ci-test
ci-test:
	$(GO_BIN) test $(GO_MOD_PARAM) -race -coverprofile=coverage.txt -covermode=atomic ./...

.PHONY: lint
lint:
	$(LINT_BIN) run

.PHONY: release
release: clean
	$(GORELEASER_BIN) $(PUBLISH_PARAM)

.PHONY: update
update:
	$(GO_BIN) get -u
ifeq ($(GO111MODULE),on)
	$(GO_BIN) mod tidy
endif
	make test
	make install
ifeq ($(GO111MODULE),on)
	$(GO_BIN) mod tidy
endif

