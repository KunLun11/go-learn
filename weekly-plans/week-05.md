# Неделя 5 — PostgreSQL, sqlc, auth, P2 готов

[← Неделя 4](week-04.md) · [Оглавление](README.md) · [Неделя 6 →](week-06.md)

> **Цель:** довести P2 до полноценного CRUD-сервиса с БД, авторизацией и интеграционными тестами.
> **Проект:** P2 `tasks`, этап 2 — Postgres, миграции, слои, JWT, testcontainers.
> **Бюджет:** ~20 часов (неделя самая плотная).

## Итоги недели

- [ ] Понимаю разницу `database/sql`, `pgx` и `pgxpool`; умею настроить пул.
- [ ] Пишу схему, индексы, транзакции; понимаю уровни изоляции на практике.
- [ ] Миграции через `goose`, накат и откат.
- [ ] `sqlc`: генерация по SQL, обработка `NULL`, `ErrNoRows`, keyset-пагинация.
- [ ] Слои: handler → service → store; DTO отдельно от домена.
- [ ] JWT access/refresh, bcrypt/argon2id, защищённые ручки.
- [ ] Интеграционные тесты с `testcontainers-go`.
- [ ] P2 поднимается с нуля одной командой, тесты зелёные с `-race`.

---

## День 1 — Postgres и пул соединений

- **Теория (40 мин):** `database/sql` (интерфейс, драйвер) vs `pgx` (нативный),
  `pgxpool`, DSN, `MaxConns`, `MinConns`, `MaxConnLifetime`, `MaxConnIdleTime`,
  `Ping`, health-check БД.
- **Практика (90 мин):**
  - `docker run postgres:16` или compose-файл.
  - Подключение через `pgxpool.New(ctx, dsn)`; `Ping`.
  - Тест: исчерпание пула при `MaxConns=1` и параллельных запросах.
  - Настройка пула из env, разумные дефолты.
- **Закрепление (30 мин):** [pgx README](https://github.com/jackc/pgx),
  `go doc github.com/jackc/pgx/v5/pgxpool`.
- **Проект (30 мин):** `internal/store` с `*pgxpool.Pool`; подключение на старте,
  закрытие на shutdown; `/readyz` проверяет БД.
- **Заметки:** почему пул меньше, чем «100 на всякий случай».

---

## День 2 — Схема, индексы, транзакции

- **Теория (40 мин):** нормализация в меру, типы (`uuid`, `timestamptz`, `text` vs `varchar`),
  индексы под запросы, составные индексы, транзакции, `READ COMMITTED` vs
  `REPEATABLE READ`, `SELECT ... FOR UPDATE`.
- **Практика (90 мин):**
  - Схема `users(id, email unique, password_hash, created_at)`,
    `tasks(id, user_id, title, status, created_at, updated_at)`.
  - Индекс под `tasks(user_id, created_at desc, id desc)` для keyset-пагинации.
  - Транзакция: создать задачу и запись аудита атомарно.
  - Воспроизведи lost update и почини через `FOR UPDATE` или версионирование.
- **Закрепление (30 мин):** PostgreSQL docs — Transactions, Indexes (выборочно).
- **Проект (30 мин):** миграции `goose`: `0001_init.sql`, `0002_tasks.sql`.
- **Заметки:** какие индексы реально нужны под твои запросы.

---

## День 3 — sqlc и слой доступа

- **Теория (40 мин):** идея sqlc: SQL → типизированный Go; `sqlc.yaml`,
  `emit_json_tags`, `pgtype` для `NULL`, `sql.ErrNoRows`/`pgx.ErrNoRows`,
  batch-вставка, `COPY`, пагинация.
- **Практика (90 мин):**
  - Установи sqlc, настрой `sqlc.yaml`, сгенерируй код по `queries/*.sql`.
  - CRUD-запросы; отдельный запрос со `LEFT JOIN` и nullable-полями.
  - Keyset-пагинация: `WHERE (created_at, id) < ($1, $2) ORDER BY ... LIMIT $3`.
  - Обработай `ErrNoRows` в сервисе → 404.
- **Закрепление (30 мин):** [sqlc — Getting Started (PostgreSQL)](https://docs.sqlc.dev/en/latest/tutorials/getting-started-postgresql.html);
  [goose](https://github.com/pressly/goose).
- **Проект (30 мин):** репозиторий `TaskStore`/`UserStore` поверх sqlc,
  интерфейсы на стороне сервиса.
- **Заметки:** `NULL` из БД: `pgtype.Text` vs `*string` — сравни.

---

## День 4 — Слои, валидация, ошибки API

- **Теория (40 мин):** handler → service → store, зачем DTO, маппинг доменных ошибок
  в HTTP, `validator`, идемпотентность, различие 400/401/403/404/409/422.
- **Практика (90 мин):**
  - DTO с `validate:"required,email"` и т.п.
  - Сервис возвращает доменные ошибки (`ErrNotFound`, `ErrConflict`).
  - Handler маппит их в RFC 7807.
  - Один сквозной сценарий: регистрация → логин → создание задачи → список.
- **Закрепление (30 мин):** [go-playground/validator](https://github.com/go-playground/validator);
  перечитай свой RFC 7807-хелпер.
- **Проект (30 мин):** полный CRUD задач с проверкой владельца.
- **Заметки:** где заканчивается валидация формы и начинается бизнес-правило.

---

## День 5 — Аутентификация и авторизация

- **Теория (40 мин):** пароли: bcrypt vs argon2id; JWT: claims, подпись, срок жизни,
  access/refresh, почему не хранить чувствительное в payload; rate limit на логин;
  CSRF не актуален для Bearer, но актуален для cookie.
- **Практика (90 мин):**
  - `bcrypt.GenerateFromPassword`, сравнение.
  - `golang-jwt/jwt/v5`: выпуск/проверка, `sub`, `exp`, `iat`.
  - Middleware аутентификации: парсит `Authorization: Bearer`, кладёт user в контекст.
  - Проверка владельца ресурса в сервисе (403).
  - Лимит попыток логина (in-memory или Redis — пока in-memory).
- **Закрепление (30 мин):** [golang-jwt](https://github.com/golang-jwt/jwt),
  `go doc golang.org/x/crypto/bcrypt`.
- **Проект (30 мин):** `POST /auth/register`, `POST /auth/login`, `POST /auth/refresh`.
- **Заметки:** почему access-токен короткий, а refresh — с ротацией.

---

## День 6 — Интеграционные тесты

- **Теория (40 мин):** `testcontainers-go`, `TestMain`, поднятие Postgres,
  применение миграций в тестах, изоляция данных (truncate/транзакция),
  параллельные тесты и общая БД.
- **Практика (90 мин):**
  - `TestMain`: старт контейнера, миграции, `defer terminate`.
  - Хелпер `newTestDB(t)` + truncate перед каждым тестом.
  - Тесты: регистрация, дубликат email (409), доступ к чужой задаче (403).
  - `t.Parallel` там, где безопасно.
  - `go test ./... -race -count=1`.
- **Закрепление (30 мин):** [testcontainers-go quickstart](https://golang.testcontainers.org/quickstart/).
- **Проект (30 мин):** `docker-compose.yml` для разработки: app + postgres;
  отдельный сервис/скрипт для миграций.
- **Заметки:** почему тестовая БД должна быть настоящей, а не SQLite.

---

## День 7 — Сборка P2 и ретроспектива

- **Теория (40 мин):** перечитай требования P2 и найди недоделки.
- **Практика (90 мин):**
  - Прогон с нуля: `docker compose down -v && docker compose up --build`.
  - Smoke-сценарий curl-ом целиком.
  - Проверка: миграции, health, ready, graceful shutdown.
  - `go vet`, линтер, покрытие.
- **Закрепление (30 мин):** сравни свой код со структурой из основного плана.
- **Проект (30 мин):** README: схема БД, ручки, переменные окружения, как тестировать.
- **Ретроспектива (20 мин):** что бы вынес в отдельный пакет; что упростил бы.

---

## Где решать и закреплять

| Ресурс | Что брать | Когда |
|---|---|---|
| [pgx](https://github.com/jackc/pgx) | pool, транзакции, типы | дни 1–2 |
| [goose](https://github.com/pressly/goose) | миграции | дни 1–2 |
| [sqlc docs](https://docs.sqlc.dev/) | генерация, nullable, пагинация | день 3 |
| [validator](https://github.com/go-playground/validator) | теги валидации | день 4 |
| [golang-jwt](https://github.com/golang-jwt/jwt) | access/refresh | день 5 |
| [testcontainers-go](https://golang.testcontainers.org/) | интеграционные тесты | день 6 |
| PostgreSQL docs | транзакции, индексы | дни 1–2 |

## Проект: P2 — критерии готовности

- [ ] `users` и `tasks`, CRUD с фильтрами и keyset-пагинацией.
- [ ] Миграции с нуля, `goose up/down`.
- [ ] `pgxpool` с настройками из env.
- [ ] `sqlc`-запросы вместо ручного `Scan`.
- [ ] JWT access/refresh, bcrypt/argon2id, проверка владельца.
- [ ] Валидация и RFC 7807.
- [ ] Интеграционные тесты на реальном Postgres.
- [ ] Compose для локального запуска.
- [ ] README, `.env.example`, Makefile.

## Типичные ошибки недели

- `SELECT *` и рассинхрон со структурой.
- `rows.Close()` забыт, `rows.Err()` не проверен.
- Пул на 200 соединений при 4 CPU.
- `ErrNoRows` проглочен → пустой ответ вместо 404.
- SQL-инъекции через конкатенацию вместо параметров.
- JWT-секрет в репозитории; нет проверки `exp`.
- Пароль в логах или в ответе API.
- Тесты на SQLite «потому что быстрее».
- `context.Background()` вместо `ctx` в запросах к БД.

## Гейт недели

```bash
docker compose down -v && docker compose up -d --build
go run ./cmd/api &
curl -s -X POST localhost:8080/auth/register -d '{"email":"a@b.c","password":"secret123"}'
curl -s -X POST localhost:8080/auth/login -d '{"email":"a@b.c","password":"secret123"}'
curl -s -X POST localhost:8080/tasks -H "Authorization: Bearer $TOKEN" -d '{"title":"x"}'
go test ./... -race -count=1
```

- [ ] Развёртывание с нуля одной командой.
- [ ] Все тесты зелёные, включая интеграционные.
- [ ] Линтер чист, секретов в коде нет.

## Самопроверка

1. Когда `pgx` лучше `database/sql`, а когда наоборот?
2. Как выбрать размер пула соединений?
3. Чем keyset-пагинация лучше offset при больших таблицах?
4. Как обрабатывать `NULL` из БД без потери смысла?
5. Где проверять владельца ресурса: в handler или service?
6. Как связаны `/readyz` и доступность БД?
7. Почему тестовая БД должна быть тем же движком, что и прод?
8. Что произойдёт с транзакцией при отмене `context`?

## Записать в заметки

- `go-gotchas.md`: `ErrNoRows`, `NULL`, пул, `rows.Close`, `context` в запросах.
- `learning-log.md`: какой слой оказался самым «толстым» и почему.
