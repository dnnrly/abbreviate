FROM golang:1.12

ADD scripts/release/entrypoint.sh /usr/bin/entrypoint.sh
RUN chmod +x /usr/bin/entrypoint.sh

ENTRYPOINT [entrypoint.sh]

