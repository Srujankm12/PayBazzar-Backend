# Use a lightweight base image
FROM golang:1.24.5-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the application
RUN go build -o main .

# Use a smaller runtime image
FROM alpine:latest

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Expose port (change if needed)
EXPOSE 8080

# Command to run the application
CMD ["./main"]
