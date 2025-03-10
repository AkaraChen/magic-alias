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

FROM ubuntu:24.04 AS test

RUN apt-get update && apt-get install -y bats

COPY --from=builder /app /app

RUN mv /app/ma /bin/ma && \
    chmod +x /bin/ma && \
    ma -h
