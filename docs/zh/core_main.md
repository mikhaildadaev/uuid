---
outline: deep
---

# API / 核心 / 主要

::: info 关于
核心 是 `ulog` 的基础。在这里，您将学习如何创建遥测实例、配置所有选项，并了解每种数据类型和字段构造函数。
:::

## NewTelemetry
创建包含所有配置选项的遥测实例
```go
ctx := context.Background()
ctx = context.WithValue(ctx, "node_id", "123-abc")
ctx = context.WithValue(ctx, "trace_id", "abc-123")
telemetry := ulog.NewTelemetry(
    ulog.WithExtractor("node_id", "trace_id"),
    ulog.WithFormat(ulog.FormatJson),
    ulog.WithLevel(ulog.LevelDebug),
    ulog.WithMode(ulog.ModeAsync, ulog.DefaultWriterOut, 1000),
    ulog.WithTheme(ulog.ThemeLight),
)
defer telemetry.Close()
telemetry.InfoWithContext(ctx, ulog.DataLog, 
    ulog.String("message", "text"),
)
telemetry.InfoWithContext(ctx, ulog.DataMetric, 
    ulog.String("name", "payments"),
    ulog.Float64("value", 99.99),
)
telemetry.InfoWithContext(ctx, ulog.DataTrace,
    ulog.String("name", "payment_processing"),
    ulog.Int64("duration", 150),
    ulog.String("span_id", "span-456"),
)
telemetry.Sync()
telemetry.SetExtractor()
telemetry.SetFormat(ulog.FormatText)
telemetry.SetLevel(ulog.LevelDebug)
telemetry.SetMode(ulog.ModeSync, ulog.DefaultWriterOut)
telemetry.SetTheme(ulog.ThemeDark)
telemetry.Info(ulog.DataLog,
	ulog.String("message", "text"),
)
telemetry.Info(ulog.DataMetric,
	ulog.String("name", "payments"),
	ulog.Float64("value", 99.99),
)
telemetry.Info(ulog.DataTrace,
	ulog.String("name", "payment_processing"),
	ulog.Int64("duration", 150),
	ulog.String("span_id", "span-456"),
)
telemetry.Sync()
```
Output:
```json
{"level":"info","type":"log","message":"text","node_id":"123-abc","trace_id":"abc-123"}
{"level":"info","type":"metric","name":"payments","value":99.99,"node_id":"123-abc","trace_id":"abc-123"}
{"level":"info","type":"trace","name":"payment_processing","duration":150,"span_id":"span-456","node_id":"123-abc","trace_id":"abc-123"}
```
```text
[INFO] type="log" message="text"
[INFO] type="metric" name="payments" value=99.99
[INFO] type="trace" name="payment_processing" duration=150 span_id="span-456"
```
