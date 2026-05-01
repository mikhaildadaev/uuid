---
outline: deep
---

# Go

::: info 关于
`uuid` 的最新稳定版本是 **v1.26.11**.
:::

```bash
go get github.com/mikhaildadaev/uuid
```

## Run Test 
```bash
go test ./...
go test -bench=. ./...
go test -cover ./...
go test -race ./...
```

## Key Features
- **UUIDv1..v8 和 Null** — 生成所有版本的 UUID 和兼容 SQL 的 null UUID；
- **解析和验证** — 解析和验证标准和 null UUID 字符串；
- **元数据提取** — 提取时间戳、序列号、节点、POSIX、变体和版本；
- **序列化** — 完整的 Binary、JSON 和 Text 编组/解组支持；
- **SQL 集成** — 原生的 database/sql 驱动，支持 Scan/Value 接口；
- **零分配** — 所有热路径上快速、无内存分配的核心例程。

## Limits
- **V2**: `posix` 应为 `0..255`;
- **V3/V5**: `name` 应为 `1..512` 个符号（字节长度）； 
- **V8**: `node` 应为 `0..16383`;

> **注意** 
>
> 长度为 0 字节或超过 512 字节的 `name` 字符串将返回 null UUID，以防止内存分配。
