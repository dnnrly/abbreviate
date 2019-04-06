CURL_BIN ?= curl
GO_BIN ?= go
LINT_BIN ?= gometalinter
PACKR_BIN ?= packr2
GORELEASER_BIN ?= goreleaser

PUBLISH_PARAM ?=
TMP_DIR=/tmp

PACKR_VERSION = 2.1.0

export PATH := ./bin:$(PATH)

install: deps

build:
	$(PACKR_BIN) build -mod vendor .

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
	# $(GO_BIN) get ./...
	$(CURL_BIN) -L https://git.io/vp6lP | sh
ifeq ($(GO111MODULE),on)
	$(GO_BIN) mod tidy
endif

./bin:
	mkdir ./bin

build-deps: ./bin
	curl -L https://github.com/gobuffalo/packr/releases/download/v${PACKR_VERSION}/packr_${PACKR_VERSION}_linux_amd64.tar.gz -o packr_${PACKR_VERSION}_linux_amd64.tar.gz
	tar -xvf ./packr_${PACKR_VERSION}_linux_amd64.tar.gz packr2
	mv packr2 ./bin/

deps: build-deps test-deps

test:
	$(GO_BIN) test -mod vendor ./...

acceptance-test:
	bats --tap acceptance.bats

ci-test:
	$(GO_BIN) test -mod vendor -race -coverprofile=coverage.txt -covermode=atomic ./...

lint:
	$(LINT_BIN) --vendor ./... --deadline=1m --skip=internal

release: clean build acceptance-test
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
