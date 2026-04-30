---
outline: deep
---

# API / Запись по сети / Основное

::: info Информация
`SinkHttp` — готовый к промышленной эксплуатации HTTP-синк с батчингом, circuit breaker, дедупликацией, повтором и выборкой. Ваш сервис никогда не блокируется при сетевой доставке.
:::

## NewSinkHttp
Создаёт HTTP-синк для отправки логов на удалённый сервер со встроенными `Batching`, `Circuit Breaker`, `Deduplication`, `Retry`, `Sampling`.
```go
sinkHttp := ulog.NewSinkHttp("http://localhost:8080/logs",
    ulog.WithHttpBatch(100, 5*time.Second),
    ulog.WithHttpCircuitBreaker(10, 10*time.Second),
    ulog.WithHttpDedupWindow(5*time.Second),
    ulog.WithHttpHeader("Authorization", "Bearer token"),
    ulog.WithHttpRetry(3, time.Second),
    ulog.WithHttpSampleRate(100),
    ulog.WithHttpTimeout(30*time.Second),
)
defer sinkHttp.Close()
telemetry := ulog.NewTelemetry(
    ulog.WithFormat(ulog.FormatJson),
    ulog.WithMode(ulog.ModeAsync, sinkHttp, 10000),
)
defer telemetry.Close()
telemetry.Error(ulog.DataLog,
    ulog.String("message", "payment failed"),
    ulog.String("service", "billing"),
)
telemetry.Sync()
```