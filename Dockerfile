FROM golang:1.24-alpine AS builder
WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download
COPY . .

FROM golang:1.24-alpine
WORKDIR /app

COPY --from=builder /go/bin/air /usr/local/bin/
COPY . .

EXPOSE 8080
