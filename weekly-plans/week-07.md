# Неделя 7 — Docker, VPS, observability, CI

[← Неделя 6](week-06.md) · [Оглавление](README.md) · [Неделя 8 →](week-08.md)

> **Цель:** превратить учебный код в сервис, который можно эксплуатировать.
> **Проект:** P3 в проде: HTTPS, метрики, CI, runbook.
> **Бюджет:** ~18 часов.

## Итоги недели

- [ ] Multi-stage Dockerfile, статический бинарник, non-root, маленький образ.
- [ ] Версия и commit пробрасываются через `-ldflags`.
- [ ] Compose «как прод»: healthchecks, секреты, миграции отдельным шагом.
- [ ] Сервис задеплоен на VPS по HTTPS через reverse proxy.
- [ ] Метрики Prometheus, JSON-логи, `/healthz`, `/readyz`, pprof под защитой.
- [ ] Нагрузочный тест и профилирование с отчётом «до/после».
- [ ] CI собирает, тестирует с `-race`, линтит и публикует образ.
- [ ] Есть runbook на одну страницу.

---

## День 1 — Docker-образ

- **Теория (40 мин):** multi-stage build, `CGO_ENABLED=0`, статическая линковка,
  `-ldflags "-s -w"`, non-root пользователь, `.dockerignore`,
  distroless/scratch, кэширование слоёв.
- **Практика (90 мин):**
  - Dockerfile: builder на `golang:alpine` → runtime `gcr.io/distroless/static`.
  - `-X main.version`, `-X main.commit`, `-X main.date` через build args.
  - `USER nonroot`, `EXPOSE`, `ENTRYPOINT`.
  - Собери и проверь размер: цель < 30 МБ.
  - `docker run` локально с env-переменными.
- **Закрепление (30 мин):** [Docker multi-stage docs](https://docs.docker.com/build/building/multi-stage/),
  [distroless](https://github.com/GoogleContainerTools/distroless).
- **Проект (30 мин):** два Dockerfile или один multi-target для `api` и `worker`.
- **Заметки:** почему `scratch` без CA-сертификатов ломает HTTPS-клиенты.

---

## День 2 — Compose «как прод»

- **Теория (40 мин):** healthcheck, `depends_on: condition: service_healthy`,
  секреты через env-файл, отдельный job для миграций, рестарт-политики.
- **Практика (90 мин):**
  - Compose: postgres, redis, api, worker, migrate (one-shot), caddy.
  - Healthcheck для postgres (`pg_isready`) и redis (`redis-cli ping`).
  - Миграции как отдельный сервис, который завершается.
  - `restart: unless-stopped`, лимиты памяти.
  - Запуск с нуля: `docker compose down -v && up -d --build`.
- **Закрепление (30 мин):** Docker Compose docs — healthcheck, depends_on.
- **Проект (30 мин):** `.env.example` без секретов + README «как поднять».
- **Заметки:** почему миграции не должны гоняться из двух реплик API.

---

## День 3 — VPS и TLS

- **Теория (40 мин):** systemd unit vs Docker на VPS, reverse proxy (Caddy/Nginx),
  автоматический TLS, firewall (`ufw`), отдельный пользователь, ротация логов,
  бэкапы БД.
- **Практика (90 мин):**
  - Подними VPS, создай non-root пользователя, `ufw allow 22,80,443`.
  - Caddy в Compose: `reverse_proxy api:8080`, автоматический сертификат.
  - Проверь HTTPS и редирект с HTTP.
  - `pg_dump` по cron + проверка восстановления на копии.
  - Ограничь доступ к внутренним портам.
- **Закрепление (30 мин):** [Caddy docs](https://caddyserver.com/docs/),
  systemd unit examples.
- **Проект (30 мин):** деплой P3, домен или nip.io, README с инструкцией.
- **Заметки:** что именно открыто наружу и почему.

---

## День 4 — Observability

- **Теория (40 мин):** три столпа: логи, метрики, трейсы (трейсы — опционально).
  `prometheus/client_golang`: Counter, Histogram, Gauge.
  SLI: latency, error rate, saturation. `/metrics`, `/healthz`, `/readyz`.
  pprof: только localhost или BasicAuth.
- **Практика (90 мин):**
  - Middleware с `http_request_duration_seconds` (Histogram) и счётчиком ошибок.
  - Метрики очереди: длина, число ретраев, длительность задачи.
  - `promhttp.Handler()` на отдельном порту или под auth.
  - pprof на `127.0.0.1:6060` или за BasicAuth.
  - Structlog: `job_id`, `request_id`, `user_id` во всех сообщениях.
- **Закрепление (30 мин):** [Prometheus Go guide](https://prometheus.io/docs/guides/go-application/),
  `go doc net/http/pprof`.
- **Проект (30 мин):** README: список метрик и что они значат.
- **Заметки:** какие метрики реально помогут в инциденте.

---

## День 5 — Нагрузка и профилирование

- **Теория (40 мин):** `k6`/`hey`, RPS, p50/p95/p99, профили CPU/heap/block/mutex,
  `GOMEMLIMIT`, `GOGC`, когда оптимизировать, а когда масштабировать.
- **Практика (90 мин):**
  - `hey -z 30s -c 50` или k6-сценарий на `POST /jobs` и `GET /jobs/{id}`.
  - Сними CPU-профиль под нагрузкой, найди топ-3 функции.
  - Heap-профиль: где растёт память.
  - Одна осознанная оптимизация (например, пул буферов или `sync.Pool`).
  - Повтори нагрузку, сравни p95 и RPS.
- **Закрепление (30 мин):** [k6 docs](https://k6.io/docs/), [hey](https://github.com/rakyll/hey),
  Go Blog — pprof.
- **Проект (30 мин):** `docs/loadtest.md`: методика, цифры, выводы.
- **Заметки:** что оказалось узким местом: CPU, БД, очередь или сеть.

---

## День 6 — CI

- **Теория (40 мин):** GitHub Actions: workflow, jobs, matrix, кэш модулей,
  сервисы (postgres/redis) для интеграционных тестов, артефакты, секреты.
- **Практика (90 мин):**
  - Workflow: `gofmt -l`, `go vet`, `golangci-lint`, `go test -race`, build.
  - Сервис-контейнеры для интеграционных тестов.
  - Кэш `~/go/pkg/mod` и `~/.cache/go-build`.
  - Сборка и пуш Docker-образа (можно в GHCR).
  - Защита main: обязательный зелёный CI.
- **Закрепление (30 мин):** [GitHub Actions docs](https://docs.github.com/en/actions).
- **Проект (30 мин):** бейдж CI в README; версия образа по тегу/коммиту.
- **Заметки:** какие шаги самые долгие и что кэшировать.

---

## День 7 — Учения и runbook

- **Теория (40 мин):** failure modes прода: упал Postgres, упал Redis, кончилось
  место на диске, утёк образ, не применились миграции.
- **Практика (90 мин):**
  - Останови Postgres — проверь `/readyz` и поведение API.
  - Останови Redis — проверь деградацию (API жив, задачи не ставятся).
  - Перезапусти контейнеры — данные на месте, задачи не потеряны.
  - Откат на предыдущий образ.
  - Проверь восстановление из бэкапа.
- **Закрепление (30 мин):** запиши шаги, которые реально выполнял.
- **Проект (30 мин):** `RUNBOOK.md`: симптомы → действия → проверка.
- **Ретроспектива (20 мин):** что бы автоматизировал в первую очередь.

---

## Где решать и закреплять

| Ресурс | Что брать | Когда |
|---|---|---|
| [Docker docs](https://docs.docker.com/build/building/multi-stage/) | multi-stage, distroless | день 1 |
| [Caddy](https://caddyserver.com/docs/) | reverse proxy и TLS | день 3 |
| [Prometheus Go guide](https://prometheus.io/docs/guides/go-application/) | метрики | день 4 |
| [k6](https://k6.io/docs/) / [hey](https://github.com/rakyll/hey) | нагрузка | день 5 |
| [GitHub Actions](https://docs.github.com/en/actions) | CI | день 6 |
| Go Blog — pprof | профилирование | день 5 |
| [12factor.net](https://12factor.net/) | конфиг, логи, процессы | дни 2–4 |

## Проект: P3 в проде — критерии готовности

- [ ] Образ < 30 МБ, non-root, версия внутри.
- [ ] Compose с healthchecks и миграциями.
- [ ] HTTPS через Caddy, наружу открыты только 22/80/443.
- [ ] Метрики и логи с request/job id.
- [ ] pprof недоступен из интернета.
- [ ] CI зелёный, образ публикуется.
- [ ] Отчёт по нагрузке с цифрами.
- [ ] Runbook и инструкция по откату.
- [ ] Бэкап БД и проверенное восстановление.

## Типичные ошибки недели

- `FROM golang` в runtime-слое → образ на сотни мегабайт.
- Запуск от root.
- Секреты в `docker-compose.yml` или в образе.
- Миграции из каждой реплики API одновременно.
- `/metrics` открыт в интернет.
- pprof доступен снаружи.
- CI без `-race` и без интеграционных тестов.
- Нагрузочный тест без фиксации методики — цифры несравнимы.
- «Оптимизация» без профиля.

## Гейт недели

```bash
docker build -t jobs:test .
docker image ls jobs:test          # размер
docker compose up -d --build
curl -sI https://<домен>/healthz   # 200 + TLS
curl -s https://<домен>/metrics | head
gh workflow view                   # или вкладка Actions: зелёный
```

- [ ] Все пункты «Итоги недели» закрыты.
- [ ] Сервис переживает перезапуск и падение Redis.
- [ ] Runbook проверен руками.

## Самопроверка

1. Зачем `CGO_ENABLED=0` и что ломается без него?
2. Как пробросить версию в бинарник?
3. Почему миграции — отдельный шаг, а не `init()` приложения?
4. Чем `/healthz` отличается от `/readyz` и как это влияет на прокси?
5. Какие метрики нужны для алерта на деградацию?
6. Почему pprof нельзя открывать наружу?
7. Что показывает p95 и почему среднее врёт?
8. Что произойдёт при SIGTERM во время in-flight задачи?

## Записать в заметки

- `go-gotchas.md`: образы, миграции, метрики, pprof.
- `learning-log.md`: цифры нагрузки до/после и вывод.
