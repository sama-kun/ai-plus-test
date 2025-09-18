

# Employee Service API

Простой сервис на Go для управления сотрудниками с PostgreSQL, Docker и Swagger.

---

## Запуск сервера

### Локально

```bash
go run ./cmd/main.go
```

Сервер будет доступен на:

```
http://localhost:5000
```

### Через Docker

```bash
docker-compose up --build
```

---

## API

* `POST /employee` — создать сотрудника
* `GET /employee` — получить список сотрудников

---

## Swagger UI

Документация доступна по адресу:

```
http://localhost:5000/swagger/index.html
```

Обновить Swagger:

```bash
swag init -g cmd/main.go -o ./docs
```

---

## Тестирование

```bash
go test ./...
```

