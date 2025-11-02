# ---------- Build stage ----------
FROM golang:1.24.5-alpine AS builder

# Faster builds: install build deps
RUN apk add --no-cache git

# Where your Go module root will be inside the container
WORKDIR /src

# Only copy go.mod/sum first to cache deps
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of your source
COPY . .

# Build target: the path to the package with "package main"
# Example: ./cmd/server or ./backend/cmd/api or . (for root)
ARG BUILD_TARGET=./cmd/server

# Build a small, static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" -o /out/app ${BUILD_TARGET}

# ---------- Runtime stage ----------
# Use scratch for tiniest image; switch to alpine if you need shell/ca-certs
FROM gcr.io/distroless/base-debian12

# Optional: add a non-root user (distroless already runs as nonroot by default)
USER nonroot:nonroot

WORKDIR /app
COPY --from=builder /out/app /app/app

# Change if your app listens elsewhere
EXPOSE 8080

# Pass envs at runtime (from Actions step)
ENV GIN_MODE=release

# Run
ENTRYPOINT ["/app/app"]
