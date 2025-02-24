FROM golang:1.24-alpine AS builder

WORKDIR /app

# Download Go modules
COPY go.mod ./
RUN go mod download

# Copy the source code
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-api

FROM alpine:latest

WORKDIR /app

# Copy binary from builder
COPY --from=builder /go-api .

# Copy database directory
COPY --from=builder /app/db ./db

# Expose the port
EXPOSE 5006

# Run binary
CMD ["./go-api"]