# Dockerfile for ISO-8583 Authorizer

# Build stage
FROM golang:1.21-alpine AS builder
RUN apk add --no-cache git build-base
WORKDIR /app

# Fetch dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build application
COPY . .
RUN go build -o bin/server ./cmd/server

# Final stage
FROM alpine:latest
RUN apk add --no-cache ca-certificates
WORKDIR /root/
COPY --from=builder /app/bin/server .

# Expose port (ajuste conforme sua configuração)
EXPOSE 9000

# Execute o binário
ENTRYPOINT ["./server"]
