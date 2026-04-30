---
outline: deep
---

# API / HTTP 接收器 / 主要

::: info 关于
`SinkHttp` — 生产就绪的 HTTP 接收器，内置批处理、断路器、去重、重试和采样。您的服务在网络传输期间永远不会阻塞。
:::

## NewSinkHttp
创建一个 HTTP 接收器，用于将日志发送到远程端点，内置 `Batching`、`Circuit Breaker`、`Deduplication`、`Retry`、`Sampling`。
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
