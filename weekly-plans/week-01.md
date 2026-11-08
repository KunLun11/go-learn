# Неделя 1 — Синтаксис, инструменты, P1 v0

[← Оглавление](README.md) · [Неделя 2 →](week-02.md)

> **Цель:** писать простые программы без гугления базового синтаксиса; освоить тулинг;
> получить работающий скелет CLI-проекта.
> **Проект:** P1 `urlcheck`, этап 1 из 3 (последовательная версия).
> **Бюджет:** ~18 часов (2,5–3 ч в день).

## Итоги недели

- [ ] Установлена актуальная версия Go, `go version` работает.
- [ ] `go.mod` с нормальным именем модуля (не `main`) и актуальной версией языка.
- [ ] Понимаю `go mod init/tidy/get`, зачем `go.sum`, где кэш модулей.
- [ ] Свободно пишу: переменные, типы, `if/for/switch`, функции, структуры, методы.
- [ ] Предсказываю вывод кода со срезами, мапами и строками.
- [ ] P1 v0: читает URL из файла/флагов, делает последовательные HTTP-запросы с таймаутом,
      печатает таблицу.
- [ ] Есть тесты на парсинг и отчёт, `gofmt` и `go vet` чистые.

---

## День 1 — Тулинг, модуль, первый запуск

- **Теория (40 мин):** установка/обновление Go, `go env`, `GOPATH` и кэш модулей,
  `go run` vs `go build` vs `go install`, зачем `go.sum`.
- **Практика (90 мин):**
  - `go version`, `go env GOMODCACHE GOPATH GOPROXY`.
  - В проекте: `go mod edit -module github.com/<you>/go-learn`,
    `go mod edit -go=<актуальная версия>`, затем `go mod tidy`.
  - Убери `rsc.io/quote` из `main.go`, напиши `hello` с аргументами командной строки.
  - `go build -o bin/hello .` и `go run .`.
- **Закрепление (30 мин):** [A Tour of Go](https://go.dev/tour/) — «Basics», пакеты,
  импорты, экспортируемые имена. Поиграйся в [Go Playground](https://go.dev/play/).
- **Проект (30 мин):** создай каркас:
  ```
  cmd/urlcheck/main.go
  internal/checker/
  internal/report/
  ```
  Пока пустые функции и `main`, который печатает версию.
- **Заметки:** что такое `GO111MODULE`, `GOTOOLCHAIN`, почему модуль не должен называться `main`.

---

## День 2 — Переменные, типы, управление потоком

- **Теория (40 мин):** базовые типы, zero values, явные конверсии, `const`/`iota`,
  `var` vs `:=`, области видимости, `if`/`for`/`switch`, `range`. В Go нет `while` и тернарника.
- **Практика (90 мин):**
  - FizzBuzz, сумма цифр, палиндром, простые числа до N.
  - Калькулятор на `switch`, обработка деления на ноль.
  - Таблица типов: `int` vs `int64`, `float64`, `byte`, `rune`, `string`, `bool`.
- **Закрепление (30 мин):** [Tour of Go](https://go.dev/tour/) — «Flow control statements».
  [Exercism Go](https://exercism.org/tracks/go): концепции *Basics* и *Numbers, strings,
  booleans* (Lasagna, Vehicle Purchase, Cars Assemble или аналоги).
- **Проект (30 мин):** парсинг аргументов через `flag`: `-o table|json`, `-timeout`,
  `-c` (пока игнорируется). Выведи распарсенные значения.
- **Заметки:** почему неявные конверсии запрещены и как это спасает от багов.

---

## День 3 — Функции, defer, panic/recover, указатели

- **Теория (40 мин):** функции, множественные возвраты, именованные результаты,
  variadic, `defer` (порядок, вычисление аргументов), `panic`/`recover`,
  указатели `&`/`*`, когда передавать значение, а когда указатель.
- **Практика (90 мин):**
  - `minMax([]int) (int, int)`.
  - Безопасное деление: `(int, error)` вместо panic.
  - Функция с `defer`, которая логирует длительность.
  - Поменять два значения местами через указатели.
- **Закрепление (30 мин):** Tour of Go — «More types» (указатели).
  Exercism: упражнения на строки (*Party Robot*, *Welcome to Tech Palace!* или аналоги).
- **Проект (30 мин):** в `internal/checker`:
  - `type Result struct { URL string; Status int; Latency time.Duration; Err error }`
  - `func Check(ctx context.Context, url string, timeout time.Duration) Result` —
    пока заглушка, реальный HTTP будет завтра.
- **Заметки:** почему `defer` в цикле — отдельная история; когда `panic` оправдан.

---

## День 4 — Срезы, мапы, строки

- **Теория (40 мин):** массив vs срез, `len`/`cap`, `append`, `copy`, `make`,
  алиасинг и переиспользование массива; мапы: `comma-ok`, `delete`, порядок обхода;
  строки как UTF-8, `rune` vs `byte`, `strings.Builder`.
- **Практика (90 мин) — обязательно с выводом на бумаге:**
  1. `s := make([]int, 2, 4)`, `append`, распечатай `len/cap` на каждом шаге.
  2. Два среза из одного массива: измени один — посмотри, что стало со вторым.
  3. `append` внутри функции — вернулся ли новый срез?
  4. Удали дубликаты из среза, разверни срез, проверь палиндром по рунам.
  5. Мапа: подсчёт слов, `comma-ok`, удаление ключа во время обхода.
- **Закрепление (30 мин):** Tour of Go — «Slices», «Maps».
  [Go by Example](https://gobyexample.com/) — Slices, Maps, Strings and Runes.
  Exercism: *Bird Watcher*, *Card Tricks*, *Gross Store*, *Logs, Logs, Logs!*.
- **Проект (30 мин):** чтение списка URL из файла, обрезка пробелов, пропуск пустых
  строк и дубликатов через мапу.
- **Заметки:** три правила срезов: длина, ёмкость, общий массив.

---

## День 5 — Структуры, методы, embedding, теги

- **Теория (40 мин):** структуры, zero value структуры, методы, value vs pointer receiver,
  embedding, теги структур, видимость полей.
- **Практика (90 мин):**
  - `type Config`, `type Result`, методы `String()` и `Validate()`.
  - Pointer receiver меняет поле, value receiver — нет: докажи кодом.
  - Вложенная структура через embedding и обращение к «повышенным» полям.
  - JSON-теги: `json:"url"`, `omitempty`, `json:"-"`.
- **Закрепление (30 мин):** Go by Example — Structs, Methods, Struct Embedding.
  Exercism: концепция *Structs*.
- **Проект (30 мин):** реальный HTTP-запрос в `Check`:
  - `http.Client{Timeout: ...}` (не `http.DefaultClient`),
  - `ctx`, `http.NewRequestWithContext`,
  - `resp.Body.Close()`, чтение и закрытие тела.
  - Результаты складываются в `[]Result`.
- **Заметки:** почему `http.Client` создаётся один раз, а не на каждый запрос.

---

## День 6 — Пакеты, структура проекта, документация

- **Теория (40 мин):** пакеты и импорты, экспорт, `internal/`, циклы импортов,
  `go doc`, комментарии к экспортируемым сущностям, `cmd/` и `internal/`.
- **Практика (90 мин):**
  - Разложи код: `cmd/urlcheck`, `internal/checker`, `internal/report`.
  - Убедись, что `internal/` нельзя импортировать снаружи модуля.
  - Добавь doc-комментарии ко всем экспортируемым типам и функциям.
  - `go doc ./internal/checker` — проверь, как это читается.
- **Закрепление (30 мин):** [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments) —
  разделы про naming, doc comments, package names.
- **Проект (30 мин):** вывод таблицы: колонки URL, STATUS, LATENCY, ERROR.
  Выравнивание через `text/tabwriter`.
- **Заметки:** правила именования пакетов; почему не `utils` и не `common`.

---

## День 7 — Тесты, форматирование, ретроспектива

- **Теория (40 мин):** `testing`, `func TestXxx(t *testing.T)`, табличные тесты,
  `t.Run`, `t.Fatalf` vs `t.Errorf`, `go test -v`, `-run`, `-cover`, `-race`.
- **Практика (90 мин):**
  - Тесты на парсинг URL-файла (включая пустые строки и мусор).
  - Табличный тест на форматирование отчёта.
  - `go test ./... -cover`, `go test ./... -race`.
  - `gofmt -l .`, `go vet ./...` — добиться пустого вывода.
- **Закрепление (30 мин):** [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests) —
  главы *Hello, World* и *Integers*. Сравни со своим кодом.
- **Проект (30 мин):** Makefile с целями `build`, `test`, `lint`, `run`.
  Первый коммит и тег `v0.1.0` (локально).
- **Ретроспектива (20 мин):** 3 строки в `learning-log.md`; что далось тяжелее всего.

---

## Где решать и закреплять

| Ресурс | Что брать | Когда |
|---|---|---|
| [A Tour of Go](https://go.dev/tour/) | Basics, Flow control, More types (slices, maps, pointers) | дни 1–4 |
| [Go by Example](https://gobyexample.com/) | Slices, Maps, Structs, Methods, JSON | дни 4–5 |
| [Exercism Go](https://exercism.org/tracks/go) | Basics, Numbers/Strings, Slices, Maps, Structs | дни 2–5 |
| [Go Playground](https://go.dev/play/) | быстрые эксперименты с `len/cap`, алиасингом | каждый день |
| [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests) | первые главы, чтобы увидеть стиль тестов | день 7 |
| LeetCode / Codewars (8–7 kyu) | разминка на синтаксис, 2–3 задачи в день | по желанию |

## Проект: P1 v0 — что должно быть сделано

```
urlcheck -o table -timeout 3s urls.txt
URL                  STATUS  LATENCY  ERROR
https://go.dev       200     120ms    -
https://example.com  200     340ms    -
```

- [ ] Чтение URL из файла и/или аргументов.
- [ ] Флаги `-o`, `-timeout`, `-c` (принимается, но не используется).
- [ ] Последовательные запросы, таймаут на каждый.
- [ ] Таблица и JSON-вывод.
- [ ] Exit code `1`, если хотя бы один запрос не удался.
- [ ] Тесты + Makefile + README.

## Типичные ошибки недели

- `http.DefaultClient` без таймаута.
- Забытый `resp.Body.Close()`.
- Ожидание, что `append` всегда меняет исходный срез.
- `range` по мапе с надеждой на порядок.
- Смешивание `byte` и `rune` при работе со строками.
- Модуль называется `main` — потом больно импортировать свои пакеты.

## Гейт недели

```bash
gofmt -l .            # пусто
go vet ./...          # пусто
go test ./... -race   # зелёный
go build -o bin/urlcheck ./cmd/urlcheck
./bin/urlcheck -timeout 2s https://go.dev
```

- [ ] Все пункты «Итоги недели» закрыты.
- [ ] README объясняет запуск за 30 секунд.

## Самопроверка

1. Чем `go run` отличается от `go build`?
2. Что хранится в `go.sum` и зачем?
3. Почему нельзя неявно сложить `int` и `int64`?
4. Чем `len` строки отличается от количества символов?
5. Что произойдёт при `append` сверх ёмкости?
6. Когда receiver должен быть указателем?

## Записать в заметки

- `learning-log.md`: что освоил, где залип.
- `go-gotchas.md`: алиасинг срезов, `cap`, порядок мапы, `len` vs `rune`.
