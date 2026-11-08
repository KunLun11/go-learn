# Неделя 4 — Стандартная библиотека и HTTP, скелет P2

[← Неделя 3](week-03.md) · [Оглавление](README.md) · [Неделя 5 →](week-05.md)

> **Цель:** уверенно писать HTTP-сервисы на stdlib, понимать, что именно добавляют фреймворки.
> **Проект:** P2 `tasks`, этап 1 — каркас API, конфиг, middleware, graceful shutdown.
> **Бюджет:** ~18 часов.

## Итоги недели

- [ ] Свободно работаю с `io`, `os`, `bufio`, `embed`, `path/filepath`.
- [ ] Знаю подводные камни `encoding/json`: теги, `omitempty`, `null` vs `[]`, стриминг.
- [ ] Понимаю `time`: длительности, layout, UTC, таймеры.
- [ ] Пишу HTTP-сервер на `net/http` с роутингом Go 1.22+ и middleware.
- [ ] Настроен HTTP-клиент: таймауты, `Transport`, переиспользование соединений.
- [ ] Логи структурированные (`log/slog`), конфиг из env с валидацией.
- [ ] P2 отвечает на `/healthz`, корректно останавливается по SIGTERM, есть тесты.

---

## День 1 — io, файлы, JSON

- **Теория (40 мин):** `io.Reader/Writer/Closer`, `io.Copy`, `bufio.Scanner/Writer`,
  `os.Open/Create/ReadFile`, `os/fs`, `embed`, `filepath.Join/WalkDir`.
  `encoding/json`: теги, `omitempty`, `-`, указатели, `null` vs `[]`,
  `json.RawMessage`, `Encoder/Decoder` для стримов.
- **Практика (90 мин):**
  - JSON-утилита: marshaling `nil`-среза и пустого среза — увидеть `null` и `[]`.
  - Кастомный `MarshalJSON` для `time.Time` в формате RFC 3339.
  - Стриминговый декодер массива объектов без загрузки всего в память.
  - `embed` шаблона/статики в бинарник.
  - Копирование файла через `io.Copy` с `bufio`.
- **Закрепление (30 мин):** Go by Example — JSON, Reading Files, Writing Files, Embed.
- **Проект (30 мин):** конфиг P2 из env: `HTTP_ADDR`, `DATABASE_URL`, `JWT_SECRET`,
  `LOG_LEVEL`; валидация на старте, понятные ошибки.
- **Заметки:** `omitempty` не различает `0` и «не задано» — как жить.

---

## День 2 — time и утилиты

- **Теория (40 мин):** `time.Duration`, `time.Time`, layout-строка, `UTC`,
  `time.Timer/Ticker`, монотонные часы, сравнение через `Equal`.
  `math/rand/v2`, `crypto/rand` для безопасности, базовые хеши.
- **Практика (90 мин):**
  - Парсер дат в нескольких форматах с понятными ошибками.
  - Генератор ID (UUID или ULID через библиотеку).
  - `Ticker` с корректной остановкой.
  - Сравнение `time.Time` через `==` vs `Equal` — воспроизведи разницу.
- **Закрепление (30 мин):** Go by Example — Time, Epoch, Time Formatting/Parsing;
  `go doc math/rand/v2`.
- **Проект (30 мин):** middleware для замера длительности запроса; `request_id`
  на основе `crypto/rand`.
- **Заметки:** всегда храни время в UTC, локаль — только на границе.

---

## День 3 — HTTP-сервер и роутинг

- **Теория (40 мин):** `http.Handler`/`HandlerFunc`, `ServeMux`,
  паттерны Go 1.22+: `"GET /tasks/{id}"`, `r.PathValue("id")`,
  коды ответов, заголовки, `http.MaxBytesReader`, `http.Error`.
- **Практика (90 мин):**
  - Роуты: `GET /healthz`, `POST /tasks`, `GET /tasks/{id}`.
  - Хелперы `writeJSON`, `writeProblem` (RFC 7807).
  - `r.PathValue` и валидация ID.
  - Лимит тела запроса и `Decoder.DisallowUnknownFields()`.
- **Закрепление (30 мин):** [Go Blog — Routing Enhancements for Go 1.22](https://go.dev/blog/routing-enhancements);
  `go doc net/http.ServeMux`.
- **Проект (30 мин):** каркас `cmd/api`, `internal/httpapi/router.go`, in-memory store
  (заглушка до недели 5).
- **Заметки:** какие ограничения у stdlib-роутера и когда нужен `chi`.

---

## День 4 — HTTP-клиент

- **Теория (40 мин):** `http.Client` с таймаутом, `http.Transport`
  (`MaxIdleConns`, `MaxIdleConnsPerHost`, `IdleConnTimeout`),
  `NewRequestWithContext`, обязательное `Body.Close()`, почему не `DefaultClient`.
  Ретраи только для идемпотентных запросов.
- **Практика (90 мин):**
  - Клиент с общим `Transport`, таймаутом и лимитом.
  - Чтение тела с ограничением (`io.LimitReader`).
  - Отмена запроса по контексту; тест с `httptest.NewServer`.
  - Ретраи с backoff поверх клиента (переиспользуй код P1).
- **Закрепление (30 мин):** `go doc net/http.Client`, `go doc net/http.Transport`;
  Go by Example — HTTP Client.
- **Проект (30 мин):** внутренний клиент для будущих вебхуков P3 — пока с тестами.
- **Заметки:** утечки соединений из-за незакрытого Body.

---

## День 5 — slog и конфигурация

- **Теория (40 мин):** `log/slog`: уровни, `slog.New` с `JSONHandler`,
  `With`/`WithGroup`, request-scoped логгер, `slog.SetDefault`.
  12-factor конфиг: только env, без секретов в репозитории.
- **Практика (90 мин):**
  - Логгер с JSON-хендлером и уровнем из env.
  - Middleware: кладёт `request_id` в контекст и логгер в контекст.
  - Логи без секретов и без интерполяции пользовательских строк в сообщение
    (использовать атрибуты).
  - `LOG_LEVEL=debug` включает подробности.
- **Закрепление (30 мин):** `go doc log/slog`; [12factor.net](https://12factor.net/config).
- **Проект (30 мин):** единый логгер в `internal/platform` или `internal/config`;
  логирование старта с версией и адресом.
- **Заметки:** почему `log.Fatal` в библиотеке — плохо.

---

## День 6 — Middleware и graceful shutdown

- **Теория (40 мин):** цепочка middleware, `http.Handler`-обёртки, порядок,
  recover-middleware, таймаут на запрос, CORS, `server.Shutdown(ctx)`,
  `signal.NotifyContext`.
- **Практика (90 мин):**
  - Middleware: recover, request ID, access log, timeout, CORS.
  - `srv := &http.Server{ReadHeaderTimeout, ReadTimeout, WriteTimeout, IdleTimeout}`.
  - `Shutdown` с таймаутом 10 секунд по SIGTERM.
  - Проверь: в момент остановки in-flight запросы доходят до конца.
- **Закрепление (30 мин):** `go doc net/http.Server.Shutdown`; паттерн из
  `net/http` examples; перечитай свой P1 `run()`.
- **Проект (30 мин):** собери `router` из всех middleware; `/healthz` и `/readyz`.
- **Заметки:** какие таймауты сервера за что отвечают.

---

## День 7 — Тесты HTTP и ретроспектива

- **Теория (40 мин):** `httptest.NewRecorder` vs `httptest.NewServer`,
  тестирование middleware, тестовые фикстуры, `t.Parallel`.
- **Практика (90 мин):**
  - Тесты всех роутов: успех, плохой JSON, неизвестное поле, лимит тела.
  - Тест middleware request ID и recover.
  - Тест graceful shutdown (запусти сервер, пошли запрос, отмени контекст).
  - `go test ./... -race -cover`.
- **Закрепление (30 мин):** Go by Example — HTTP Server; `go doc net/http/httptest`.
- **Проект (30 мин):** README P2: как запускать, переменные окружения, примеры curl.
- **Ретроспектива (20 мин):** что упростить в структуре пакетов до добавления БД.

---

## Где решать и закреплять

| Ресурс | Что брать | Когда |
|---|---|---|
| [Go by Example](https://gobyexample.com/) | JSON, Files, Time, HTTP Server, HTTP Client | дни 1–4 |
| [Go Blog — Routing 1.22](https://go.dev/blog/routing-enhancements) | паттерны `ServeMux` | день 3 |
| `go doc net/http`, `log/slog`, `encoding/json` | контракты API | дни 1–6 |
| [12factor.net](https://12factor.net/) | конфиг и логи | день 5 |
| [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests) | HTTP handlers, JSON | дни 3, 7 |
| [RFC 7807](https://www.rfc-editor.org/rfc/rfc7807) | формат ошибок API | день 3 |

## Проект: P2, этап 1 — что должно работать

```bash
curl -s localhost:8080/healthz
curl -s localhost:8080/readyz
curl -s -X POST localhost:8080/tasks -d '{"title":"learn go"}'
curl -s localhost:8080/tasks
```

- [ ] `cmd/api` с конфигом из env и валидацией.
- [ ] Роутер с паттернами `METHOD /path/{id}`.
- [ ] Middleware: recover, request ID, access log, timeout, CORS.
- [ ] In-memory store как временная заглушка.
- [ ] Единый формат ошибок RFC 7807.
- [ ] Graceful shutdown по SIGTERM.
- [ ] `/healthz` и `/readyz`.
- [ ] Тесты на роуты и middleware.
- [ ] README и `.env.example`.

## Типичные ошибки недели

- `http.DefaultClient` / `http.DefaultServeMux` в проде.
- Нет `ReadHeaderTimeout` → Slowloris.
- `json.NewDecoder(r.Body)` без лимита и без `DisallowUnknownFields`.
- Забытый `Body.Close()` в клиенте.
- Логирование через `fmt.Printf` и склейку строк.
- Секреты в коде или в `README`.
- `os.Exit` посреди обработчика — не даёт сработать `defer`.
- Отсутствие таймаута на `Shutdown`.

## Гейт недели

```bash
go test ./... -race -cover
go vet ./...
golangci-lint run

go run ./cmd/api &
curl -i localhost:8080/healthz
curl -i -X POST localhost:8080/tasks -d '{"title":"x"}'
kill -TERM %1   # сервер завершается корректно, в логе видно shutdown
```

- [ ] Все пункты «Итоги недели» закрыты.
- [ ] В логах JSON с `request_id`.
- [ ] Ни одного `fmt.Println` в `internal/`.

## Самопроверка

1. Чем `httptest.NewRecorder` отличается от `httptest.NewServer`?
2. Зачем `ReadHeaderTimeout`, если есть `ReadTimeout`?
3. Почему `json.Marshal` на nil-срезе даёт `null` и как получить `[]`?
4. Что делает `DisallowUnknownFields` и когда это плохо?
5. Почему `http.Client` создают один раз на приложение?
6. Как работает `Server.Shutdown` и что будет с in-flight запросами?
7. Чем `/healthz` отличается от `/readyz`?

## Записать в заметки

- `go-gotchas.md`: `omitempty`, nil-срез в JSON, Body.Close, таймауты сервера.
- `learning-log.md`: какие middleware оказались неочевидными.
