# go-health-service

Минимальный HTTP-сервис на Go с эндпоинтом проверки состояния. Учебный проект для отработки CI/CD.

## Возможности

- `GET /health` возвращает `{"status":"ok"}`
- Порт настраивается через переменную окружения `PORT` (по умолчанию `8080`)

## Требования

- Go 1.22+

## Запуск

```bash
go run .
curl http://localhost:8080/health
```

## Тесты

```bash
go test ./... -v
```

## Сборка

```bash
go build -o bin/go-health-service .
```

## CI/CD

Статус: в разработке (GitHub Actions).

## Лицензия

MIT
