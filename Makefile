CURL_BIN ?= curl
GO_BIN ?= go
LINT_BIN ?= golangci-lint
GORELEASER_BIN ?= goreleaser

PUBLISH_PARAM?=
GO_MOD_PARAM?=
TMP_DIR=?/tmp

BASE_DIR=$(shell pwd)

export PATH := ./bin:$(PATH)

install: deps

build:
	$(GO_BIN) build $(GO_MOD_PARAM)

clean:
	rm -f abbreviate
	rm -rf dist

clean-deps:
	rm -rf ./bin
	rm -rf ./tmp
	rm -rf ./libexec
	rm -rf ./share

./bin/bats:
	git clone https://github.com/sstephenson/bats.git ./tmp/bats
	./tmp/bats/install.sh .

test-deps: ./bin/bats
	$(CURL_BIN) -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b ./bin v1.21.0
	$(GO_BIN) get ./...
ifeq ($(GO111MODULE),on)
	$(GO_BIN) mod tidy
endif

./bin:
	mkdir ./bin

./tmp:
	mkdir ./tmp

./bin/goreleaser: ./bin ./tmp
	$(CURL_BIN) --fail -L -o ./tmp/goreleaser.tar.gz https://github.com/goreleaser/goreleaser/releases/download/v0.117.2/goreleaser_Linux_x86_64.tar.gz
	gunzip -f ./tmp/goreleaser.tar.gz
	tar -C ./bin -xvf ./tmp/goreleaser.tar

build-deps: ./bin/goreleaser

deps: build-deps test-deps

test:
	$(GO_BIN) test $(GO_MOD_PARAM) ./...

acceptance-test:
	bats --tap acceptance.bats

ci-test:
	$(GO_BIN) test $(GO_MOD_PARAM) -race -coverprofile=coverage.txt -covermode=atomic ./...

lint:
	$(LINT_BIN) run

release: clean
	$(GORELEASER_BIN) $(PUBLISH_PARAM)

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
