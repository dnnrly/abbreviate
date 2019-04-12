CURL_BIN ?= curl
GO_BIN ?= go
LINT_BIN ?= gometalinter
PACKR_BIN ?= ./bin/packr2
GORELEASER_BIN ?= goreleaser

PUBLISH_PARAM?=
GO_MOD_PARAM?=-mod vendor
TMP_DIR=?/tmp

PACKR_VERSION = 2.1.0
BASE_DIR=$(shell pwd)

export PATH := ./bin:$(PATH)

install: deps

build:
	$(PACKR_BIN)
	$(GO_BIN) build $(GO_MOD_PARAM)

clean:
	$(PACKR_BIN) clean
	rm -f abbreviate
	rm -rf dist

clean-deps:
	rm -rf ./bin
	rm -rf ./tmp
	rm -rf ./libexec
	rm -rf ./share
	rm packr_${PACKR_VERSION}_linux_amd64.tar.gz

./bin/bats:
	git clone https://github.com/sstephenson/bats.git ./tmp/bats
	./tmp/bats/install.sh .

test-deps: ./bin/bats
	$(CURL_BIN) -L https://git.io/vp6lP | sh
	$(GO_BIN) get ./...
ifeq ($(GO111MODULE),on)
	$(GO_BIN) mod tidy
endif

./bin:
	mkdir ./bin

./bin/packr2: ./bin
	cd vendor/github.com/gobuffalo/packr/v2/packr2; go build -pkgdir $(BASE_DIR)/vendor -o $(BASE_DIR)/bin/packr2

build-deps: ./bin/packr2

deps: build-deps test-deps

test:
	$(GO_BIN) test $(GO_MOD_PARAM) ./...

acceptance-test:
	bats --tap acceptance.bats

ci-test:
	$(GO_BIN) test $(GO_MOD_PARAM) -race -coverprofile=coverage.txt -covermode=atomic ./...

lint:
	$(LINT_BIN) --vendor ./... --deadline=1m --skip=internal

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
