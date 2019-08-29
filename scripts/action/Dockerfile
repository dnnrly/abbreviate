FROM golang:1.12

ENV GO111MODULE on

RUN go install github.com/gobuffalo/packr/v2/packr2
RUN git clone https://github.com/sstephenson/bats.git /tmp/bats \
  && /tmp/bats/install.sh /usr/local

RUN apt-get install curl \
  && curl https://git.io/vp6lP | sh

