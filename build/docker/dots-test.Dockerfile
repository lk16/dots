FROM golang:1.14-alpine

RUN apk update && apk upgrade && apk add --no-cache bash

LABEL maintainer="Luuk Verweij <luuk_verweij@msn.com>"

RUN  wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.25.0

WORKDIR /app

ADD . .

RUN go mod download

ENV CGO_ENABLED 0
RUN golangci-lint run ./... && \
go install ./cmd/dots/main.go && \
go test -coverprofile=coverage.out ./internal/...