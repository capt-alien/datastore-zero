# Stage 1: Build with CGO enabled
FROM golang:1.24.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Important: CGO_ENABLED=1 for sqlite3 support
RUN CGO_ENABLED=1 GOOS=linux go build -o datastore-zero cmd/main.go

# Stage 2: Minimal Debian runtime
FROM debian:bookworm-slim

# Install required packages (CA certs + optional sqlite3 CLI for debug)
RUN apt-get update && apt-get install -y \
    ca-certificates \
    sqlite3 \
 && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /app/datastore-zero .

# Optional: create the /data directory in case you're persisting
RUN mkdir -p /app/data

EXPOSE 8080

CMD ["./datastore-zero"]
