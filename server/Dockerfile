# Stage 1: Build
FROM golang:1.23.1-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Копируем сертификаты в директорию проекта для доступа
COPY ./certs/fullchain.pem ./certs/fullchain.pem
COPY ./certs/privkey.pem ./certs/privkey.pem

RUN go build -o main .

# Stage 2: Create final image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/certs/fullchain.pem ./certs/fullchain.pem
COPY --from=builder /app/certs/privkey.pem ./certs/privkey.pem

EXPOSE 8443

CMD ["./main"]
