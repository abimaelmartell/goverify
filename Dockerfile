FROM golang:1.7.4-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git

RUN mkdir -p $GOPATH/src/github.com/abimaelmartell/goverify

WORKDIR "$GOPATH/src/github.com/abimaelmartell/goverify"

COPY . .

EXPOSE 8080

RUN go build

ENTRYPOINT ./goverify
