# Указываем базовый образ, который содержит Golang
FROM golang:1.23.1-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY /etc/letsencrypt/live/tonpredictionmarket.com/fullchain.pem .
COPY /etc/letsencrypt/live/tonpredictionmarket.com/privkey.pem .

RUN go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/fullchain.pem .
COPY --from=builder /app/privkey.pem .

EXPOSE 8443

CMD ["./main"]