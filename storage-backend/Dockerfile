# Используем официальный образ Go
FROM golang:1.23-alpine AS builder

# Устанавливаем зависимости
WORKDIR /app
COPY go.mod .
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем бинарник
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/main main.go

# Итоговый образ (маленький)
FROM alpine:latest
WORKDIR /root/

# Копируем бинарник из builder-образа
COPY --from=builder /go/bin/main /root/main

# Порт и запуск
EXPOSE 8080
CMD ["./main"]
