---
outline: deep
---

# API / SinkHttp / Main

::: info Info
`SinkHttp` — production-ready HTTP sink with batching, circuit breaker, deduplication, retry, and sampling. Your service never blocks during network delivery.
:::

## NewSinkHttp
Creates an HTTP sink for sending logs to a remote endpoint with `Batching`, `Circuit Breaker`, `Deduplication`, `Retry`, `Sampling` built in
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