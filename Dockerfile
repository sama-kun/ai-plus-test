FROM golang:1.24-alpine AS builder

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
COPY --from=builder /app/config ./config


# 6. Запуск приложения
CMD ["/app/main"]
