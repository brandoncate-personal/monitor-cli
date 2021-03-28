FROM golang:1.15-alpine AS builder

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /tmp/app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./out/monitor ./cmd/exe

# Start fresh from a smaller image
FROM alpine:3.13 
RUN apk add ca-certificates

COPY --from=builder /tmp/app/out/monitor /app/monitor

# Run the binary program produced by `go install`
ENTRYPOINT ["/app/monitor"]