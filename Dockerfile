# Build
FROM golang:1.25.2-alpine AS builder
WORKDIR /app
ENV CGO_ENABLED=0
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o imob ./cmd/api

# Runtime
FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/imob /imob
EXPOSE 8000
ENTRYPOINT ["/imob"]
