FROM golang:1.7.4-alpine

MAINTAINER me@abimaelmartell.com

RUN mkdir -p $GOPATH/src/github.com/abimaelmartell/goverify

WORKDIR "$GOPATH/src/github.com/abimaelmartell/goverify"

COPY . .

RUN go get

ENTRYPOINT ["/go/bin/goverify"]

EXPOSE 8080