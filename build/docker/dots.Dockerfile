# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from golang:1.13-alpine base image
FROM golang:1.13-alpine

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

# Add Maintainer Info
LABEL maintainer="Luuk Verweij <luuk_verweij@msn.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

ADD . .

# Build the Go app
ENV CGO_ENABLED 0
RUN go install ./cmd/dots

# Run the executable
CMD ["dots"]