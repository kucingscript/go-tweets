FROM golang:1.24-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/app ./cmd/main.go

FROM scratch
WORKDIR /app

COPY --from=builder /app/bin/app .

EXPOSE 8080

CMD ["./app"]