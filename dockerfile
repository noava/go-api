FROM golang:1.24

WORKDIR /app

# Download Go modules
COPY go.mod ./
RUN go mod download

# Copy the source code
COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-api

# Expose the port
EXPOSE 5006

# Run
CMD ["/go-api"]