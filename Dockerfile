# Build Stage
FROM golang:1.23-alpine AS builder

# Set environment variables
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Install dependencies
RUN apk update && apk add --no-cache git

# Create and set working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project directory
COPY . .

# Build the application
RUN go build -o server ./cmd

# Final Stage
FROM alpine:latest

# Set working directory in the final image
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/server .

# Expose the application's port
EXPOSE 3000

# Run the binary
CMD ["./server"]
