# Use the official Go image as the base image
FROM golang:1.24.1-bullseye AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Install which command
RUN apt-get update && apt-get install -y --no-install-recommends debianutils

# Copy the source code
COPY . .

RUN go install github.com/go-delve/delve/cmd/dlv@latest

ENV SHELL=/bin/bash

RUN go build -o /app/ma ./cmd/ma

FROM alpine:3.21.3 AS test

RUN apk add --no-cache bats

COPY --from=builder /app /app

ENV PATH="/usr/local/bin:${PATH}"

RUN mv /app/ma /usr/local/bin/ma && \
    chmod +x /usr/local/bin/ma && \
    ma version
