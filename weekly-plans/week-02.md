# Неделя 2 — Интерфейсы, ошибки, generics, тесты, P1 v1

[← Неделя 1](week-01.md) · [Оглавление](README.md) · [Неделя 3 →](week-03.md)

> **Цель:** писать идиоматичный Go, а не «Python с другим синтаксисом»; освоить
> интерфейсы, обработку ошибок и полноценное тестирование.
> **Проект:** P1 `urlcheck`, этап 2 (интерфейсы, ошибки, тесты, линтер).
> **Бюджет:** ~18 часов.

## Итоги недели

- [ ] Понимаю неявную реализацию интерфейсов и ловушку nil-интерфейса.
- [ ] Умею проектировать минимальные интерфейсы на стороне потребителя.
- [ ] Ошибки: свои типы, `%w`, `errors.Is/As/Join`, без сравнения строк.
- [ ] Знаю, зачем functional options и embedding, и когда их не надо.
- [ ] Понимаю generics и ограничения (`comparable`, `~`), не тащу их везде.
- [ ] Пишу табличные тесты, сабтесты, бенчмарки; настроен `golangci-lint`.
- [ ] P1 v1: отчёт через интерфейс, ошибки типизированы, README и линтер в порядке.

---

## День 1 — Интерфейсы

- **Теория (40 мин):** интерфейс как набор методов, неявная реализация,
  `io.Reader`/`io.Writer`/`fmt.Stringer`, `any`, type assertion, type switch,
  nil-интерфейс с типизированным nil внутри.
- **Практика (90 мин):**
  - `Reporter` с методом `Render(io.Writer, []Result) error` и две реализации:
    table и JSON. Проверь, что обе удовлетворяют интерфейсу.
  - Свой `io.Writer`: счётчик байт и writer, который пишет в `bufio.Writer`.
  - Воспроизведи ловушку:
    ```go
    var err error = (*MyError)(nil)
    fmt.Println(err == nil) // ?
    ```
- **Закрепление (30 мин):** Tour of Go — Methods and interfaces.
  [Effective Go](https://go.dev/doc/effective_go) — раздел Interfaces.
- **Проект (30 мин):** вынеси рендеринг отчёта в `internal/report`, `main` зависит
  от интерфейса, а не от конкретного типа.
- **Заметки:** почему интерфейс объявляют там, где используют, а не где реализуют.

---

## День 2 — Ошибки

- **Теория (40 мин):** `error` как значение, sentinel-ошибки, свои типы,
  `Unwrap`, `fmt.Errorf("%w")`, `errors.Is`, `errors.As`, `errors.Join`,
  когда `panic` допустим (программная ошибка, а не бизнес-отказ).
- **Практика (90 мин):**
  - `var ErrInvalidURL = errors.New("invalid url")`.
  - `type HTTPError struct { Status int; URL string }` + `Error()` + `Unwrap()`.
  - Оберни ошибку слоя checker в main, проверь `errors.Is`/`As`.
  - Несколько ошибок сразу — `errors.Join`.
  - Убери все `fmt.Errorf("...: %v")` в пользу `%w`.
- **Закрепление (30 мин):** Go Blog — [Working with Errors in Go 1.13](https://go.dev/blog/go1.13-errors)
  (актуально и позже). Go by Example — Errors.
- **Проект (30 мин):** классификация ошибок: DNS, timeout, HTTP 5xx, отмена контекста.
  Отдельные счётчики в итоговом отчёте.
- **Заметки:** правило «оборачивай один раз, проверяй через `Is/As`, не парси строки».

---

## День 3 — Композиция, embedding, functional options

- **Теория (40 мин):** композиция вместо наследования, embedding структур и интерфейсов,
  декораторы-обёртки, functional options, почему геттеры/сеттеры не нужны по умолчанию.
- **Практика (90 мин):**
  - Обёртка `loggingReporter{next Reporter}` — логирует вызов и делегирует.
  - `type Option func(*Config)` и `NewClient(urls []string, opts ...Option)`.
  - Пять опций: `WithTimeout`, `WithConcurrency`, `WithRetries`, `WithOutput`, `WithRateLimit`.
  - Сравни с конфиг-структурой: где опции лучше, а где хуже.
- **Закрепление (30 мин):** [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments) —
  разделы про embedding и functional options; Go by Example — Variadic Functions.
- **Проект (30 мин):** переведи конфиг P1 на functional options, дефолты — в одном месте.
- **Заметки:** «принимай интерфейсы, возвращай структуры» — что это значит на практике.

---

## День 4 — Generics

- **Теория (40 мин):** параметры типов, constraints, `any`, `comparable`, `~int`,
  пакеты `slices`, `maps`, `cmp` (Go 1.21+), когда generics **не** нужны.
- **Практика (90 мин):**
  - `Map[T, U]`, `Filter[T]`, `Reduce[T]`, `Unique[T comparable]`.
  - `Min[T cmp.Ordered]`, `Max`.
  - Ограничение `~int | ~int64` и свой интерфейсный constraint.
  - Перепиши одно место, где generics явно лишние, обратно на конкретный тип.
- **Закрепление (30 мин):** Tour of Go — Generics. `go doc slices`, `go doc maps`.
  Exercism: концепция *Generics*.
- **Проект (30 мин):** используй `slices.SortFunc` и `maps.Keys` в отчёте
  (сортировка результатов по статусу/задержке).
- **Заметки:** два вопроса перед generics: «сколько типов реально?», «не проще ли интерфейс?».

---

## День 5 — Тестирование всерьёз

- **Теория (40 мин):** табличные тесты, сабтесты, `t.Helper`, `t.Cleanup`, `t.TempDir`,
  `t.Parallel`, фейки vs моки, `testify/require`, golden-файлы, бенчмарки, `benchstat`.
- **Практика (90 мин):**
  - Табличный тест checker с фейковым HTTP-сервером (`httptest.NewServer`).
  - Тест отчёта: golden-файл в `testdata/`.
  - Фейковый `Reporter` и проверка, что main вызывает его один раз.
  - Бенчмарк парсинга и рендера: `go test -bench . -benchmem`.
  - `go install golang.org/x/perf/cmd/benchstat@latest`, сравни два прогона.
- **Закрепление (30 мин):** Learn Go with Tests — главы про mocking и dependency injection.
- **Проект (30 мин):** тесты на все ветки: успех, таймаут, 5xx, невалидный URL, пустой файл.
- **Заметки:** почему фейк лучше мока; что тестировать на границах пакета.

---

## День 6 — Линтер, стиль, структура

- **Теория (40 мин):** `gofmt`/`goimports`, `go vet`, `staticcheck`, `golangci-lint`,
  типичные линты (unused, errcheck, govet, revive, gocritic), как читать отчёт линтера.
- **Практика (90 мин):**
  - Установи `goimports` и `golangci-lint`.
  - Создай `.golangci.yml` с разумным набором линтеров.
  - Прогони и исправь всё, включая `errcheck` и `unused`.
  - Настрой VS Code Go: format on save, organize imports.
- **Закрепление (30 мин):** [golangci-lint docs](https://golangci-lint.run/),
  [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments) целиком.
- **Проект (30 мин):** приведи имена, комментарии и структуру пакетов в порядок.
- **Заметки:** какие линты ловят реальные баги, а какие — вкусовщина.

---

## День 7 — Рефакторинг, README, ретроспектива

- **Теория (40 мин):** перечитай свой код на неделю назад и найди три вещи, которые
  сделал бы иначе.
- **Практика (90 мин):**
  - Раздели `main` на `run() error` + `main()` с `os.Exit`.
  - Убедись, что все ошибки возвращаются, а не логируются и не теряются.
  - Проверь, что нет глобального состояния и `init()`.
  - Прогони `go test ./... -race -cover`.
- **Закрепление (30 мин):** сравни свой P1 с примерами из Effective Go; выпиши различия.
- **Проект (30 мин):** README: установка, флаги, примеры, ограничения; Makefile `lint`.
- **Ретроспектива (20 мин):** что было самым сложным; что занести в `go-gotchas.md`.

---

## Где решать и закреплять

| Ресурс | Что брать | Когда |
|---|---|---|
| [A Tour of Go](https://go.dev/tour/) | Methods and interfaces, Generics | дни 1, 4 |
| [Effective Go](https://go.dev/doc/effective_go) | Interfaces, Errors, Embedding | дни 1–3 |
| [Go by Example](https://gobyexample.com/) | Errors, Interfaces, Variadic, Generics | дни 1–4 |
| [Exercism Go](https://exercism.org/tracks/go) | Interfaces, Errors, Generics | дни 1–4 |
| [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests) | Mocking, DI, Structs/Methods | дни 3, 5 |
| [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments) | стиль и нейминг | дни 3, 6 |

## Проект: P1 v1 — что добавляется

- [ ] `Reporter` как интерфейс, реализации table/JSON.
- [ ] Обёртка-логгер над репортёром (композиция).
- [ ] Конфиг через functional options.
- [ ] Типизированные ошибки, `errors.Is/As`, без `%v`-обёрток.
- [ ] `run() error` + `main()` с корректным exit code.
- [ ] Табличные тесты, `httptest`, golden-файл, бенчмарк.
- [ ] `.golangci.yml`, чистый линт, README и Makefile.

## Типичные ошибки недели

- Интерфейс из 6 методов «на всякий случай».
- `err == ErrX` вместо `errors.Is`.
- Потеря контекста ошибки: `return err` без пояснения или, наоборот, `%v` вместо `%w`.
- nil-интерфейс с nil-указателем внутри.
- Generics ради generics.
- Тесты, которые проверяют реализацию, а не поведение.

## Гейт недели

```bash
go build ./...
go test ./... -race -cover
go vet ./...
golangci-lint run
gofmt -l .
```

- [ ] Все пункты «Итоги недели» закрыты.
- [ ] В коде нет `fmt.Errorf("...: %v", err)`.
- [ ] Линтер зелёный без `//nolint`-костылей.

## Самопроверка

1. Почему интерфейс лучше объявлять в пакете-потребителе?
2. Как проверить, что тип реализует интерфейс, ещё на этапе компиляции?
3. Чем `errors.Is` отличается от `errors.As` и когда нужен `errors.Join`?
4. Когда `panic` допустим, а когда это баг дизайна?
5. Что не так с `var e error = (*MyErr)(nil); e != nil`?
6. Зачем functional options, если есть структура конфига?
7. Когда generics ухудшают код?

## Записать в заметки

- `go-gotchas.md`: nil-интерфейс, `%w` vs `%v`, `errors.Is/As`.
- `learning-log.md`: какие линты поймали реальные проблемы.
