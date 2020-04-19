FROM golang:1.14-alpine

RUN apk update && apk upgrade && apk add --no-cache bash

LABEL maintainer="Luuk Verweij <luuk_verweij@msn.com>"

WORKDIR /app

ADD . .

RUN go mod download

ENV CGO_ENABLED 0
RUN go install ./cmd/dots/main.go

CMD go test ./internal/...