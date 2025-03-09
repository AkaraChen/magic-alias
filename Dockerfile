# Use the official Go image as the base image
FROM golang:1.24.1-bullseye

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

ENV SHELL=/bin/bash
