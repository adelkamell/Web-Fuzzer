FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /web-fuzzer ./cmd/web-fuzzer

FROM alpine:latest
COPY --from=builder /web-fuzzer /web-fuzzer
ENTRYPOINT ["/web-fuzzer"]