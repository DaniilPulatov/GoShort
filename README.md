# GOshort

**GoShort** — это простой сервис сокращения URL, написанный на Go.

## 🔧 Стек технологий

- **Golang** — основной язык разработки
- **Gin** — HTTP-фреймворк для маршрутизации и middleware
- **PostgreSQL** — СУБД для хранения URL
- **uber/fx** — фреймворк для внедрения зависимостей (DI)

## 🚀 API

- Базовый префикс: `/api/v1`

| Метод | Путь             | Описание                         |
|-------|------------------|----------------------------------|
| POST  | `/shorten`       | Создаёт короткую ссылку         |
| GET   | `/:token`        | Редирект по токену на оригинал  |

### Пример запроса на сокращение

POST /api/v1/shorten
Content-Type: application/json

{
    "url": "https://example.com",
    "identifier":"optional_identifier"
}

### Пример ответа

{
    "short_url": "https://short.ly/optional_or_generated_identifier"
}


## Конфигурация

The following environment variables need to be set in your `.env` file:

| Variable         | Default Value                          | Description                                                                 |
|------------------|----------------------------------------|-----------------------------------------------------------------------------|
| `HOST`          | `0.0.0.0`                             | The host address the application will bind to (0.0.0.0 for all interfaces)  |
| `PORT`          | `9999`                                | The port number the application will listen on                              |
| `DATABASE_URL`  | `postgres://user:1234@localhost:5432/url_db?sslmode=disable` | PostgreSQL connection URL with credentials |
| `BASE_URL`      | `http://localhost:9999`               | Base URL for the application (used for generating absolute URLs)            |
| `MIGRATIONS_DIR`| `internal/migrations`                 | Directory where database migration files are stored                         |

### Пример `.env` файла

```env
HOST=0.0.0.0
PORT=9999
DATABASE_URL=postgres://user:1234@localhost:5432/url_db?sslmode=disable
BASE_URL=http://localhost:9999
MIGRATIONS_DIR=internal/migrations