FROM golang:1.14-alpine

ENV GO111MODULE=on
ENV CGO_ENABLED=0

WORKDIR /go/src/github.com/kons16/book-shelf-server

COPY . .
RUN apk add --no-cache && go get -u github.com/pilu/fresh github.com/rubenv/sql-migrate/...
