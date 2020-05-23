FROM golang:1.13

ENV GO111MODULE on
ENV GOPROXY https://proxy.golang.org

RUN git clone https://github.com/sstephenson/bats.git /tmp/bats \
  && /tmp/bats/install.sh /usr/local

RUN curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin v1.17.1

