CURL_BIN ?= curl
GO_BIN ?= go
LINT_BIN ?= gometalinter
PACKR_BIN ?= packr2

export PATH := ./bin:$(PATH)

install: deps

build:
	$(PACKR_BIN) build

deps:
	$(GO_BIN) get ./...
	$(GO_BIN) get github.com/gobuffalo/packr/v2/packr2
	$(CURL_BIN) -L https://git.io/vp6lP | sh
ifeq ($(GO111MODULE),on)
	$(GO_BIN) mod tidy
endif

test:
	$(GO_BIN) test ./...

ci-test:
	$(GO_BIN) test -race  -coverprofile=coverage.txt -covermode=atomic ./...

lint:
	$(LINT_BIN) --vendor ./... --deadline=1m --skip=internal

update:
	$(GO_BIN) get -u -tags ${TAGS}
ifeq ($(GO111MODULE),on)
	$(GO_BIN) mod tidy
endif
	make test
	make install
ifeq ($(GO111MODULE),on)
	$(GO_BIN) mod tidy
endif
