# 1. Используем минимальный образ с Go
FROM golang:1.23-alpine AS builder

WORKDIR /app

# 2. Устанавливаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# 3. Копируем исходный код и компилируем бинарник
COPY . .
RUN go build -o main ./cmd/main.go

# 4. Создаем финальный образ
FROM alpine:latest

WORKDIR /app

# 5. Копируем бинарник из builder'а
COPY --from=builder /app/main .
COPY ./migrations ./migrations

# 7. Запуск приложения
CMD ["/app/main"]