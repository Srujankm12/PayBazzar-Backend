# Use Go base image
FROM golang:1.24.5-alpine

# Set working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy everything (source code)
COPY . .

# Build the Go app located in cmd/
RUN go build -o main ./cmd

# Expose your app port (change if needed)
EXPOSE 8080

# Run the built binary
CMD ["./main"]
