FROM golang:1.11.2 as packr

RUN go get github.com/gobuffalo/packr/v2/packr2

FROM golang:1.11.2 as builder

COPY --from=packr /go/bin/packr2 /go/bin

RUN mkdir -p /go/src/github.com/dnnrly/abbreviate
ADD . /go/src/github.com/dnnrly/abbreviate
WORKDIR /go/src/github.com/dnnrly/abbreviate

RUN make clean
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 packr2 build -ldflags '-w -s' -a -installsuffix cgo

FROM scratch

COPY --from=builder /go/src/github.com/dnnrly/abbreviate/abbreviate /abbreviate

ENTRYPOINT ["/abbreviate"]

