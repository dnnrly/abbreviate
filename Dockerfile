FROM golang:1.11.2 as builder

RUN mkdir -p /go/src/github.com/dnnrly/abbreviate
ADD . /go/src/github.com/dnnrly/abbreviate
WORKDIR /go/src/github.com/dnnrly/abbreviate

RUN make clean
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 build -ldflags '-w -s' -a -installsuffix cgo

FROM scratch

COPY --from=builder /go/src/github.com/dnnrly/abbreviate/abbreviate /abbreviate

ENTRYPOINT ["/abbreviate"]

