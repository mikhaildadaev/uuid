---
outline: deep
---

# API / 核心 / 主要

::: warning 警告
本页正在开发中
:::

## NewNull
创建一个与 SQL 兼容的 `NULL` 值 UUID。
```go
import "github.com/mikhaildadaev/uuid"
un := uuid.NewNull()
fmt.Println(un.IsZero())
fmt.Println(un.String())
```
Output:
```text
true
00000000-0000-0000-0000-000000000000
```

## NewV1
创建一个基于当前时间戳和本地机器 MAC 地址（如果不可用则为随机节点）的 UUID 版本 1。
```go
import "github.com/mikhaildadaev/uuid"
u1 := uuid.NewV1()
fmt.Println(u1.Version())
```
Output:
```text
1
```

## NewV2
创建一个 UUID 版本 2（DCE 安全），使用当前时间戳、本地机器的 MAC 地址和 POSIX UID/GID。
```go
import "github.com/mikhaildadaev/uuid"
u2 := uuid.NewV2(posix)
fmt.Println(u2.Version())
```
Output:
```text
2
```

## NewV3
通过使用 MD5 哈希命名空间标识符和名称来创建 UUID 版本 3。
```go
import "github.com/mikhaildadaev/uuid"
u3 := uuid.NewV3(uuid.NameSpaceDNS, name)
fmt.Println(u3.Version())
```
Output:
```text
3
```

## NewV4
使用加密安全的随机数创建 UUID 版本 4。
```go
import "github.com/mikhaildadaev/uuid"
u4 := uuid.NewV4()
fmt.Println(u4.Version())
```
Output:
```text
4
```

## NewV5
通过使用 SHA-1 哈希命名空间标识符和名称来创建 UUID 版本 5。
```go
import "github.com/mikhaildadaev/uuid"
u5 := uuid.NewV5(uuid.NameSpaceDNS, name)
fmt.Println(u5.Version())
```
Output:
```text
5
```

## NewV6
创建一个基于当前时间戳和本地机器 MAC 地址（或随机节点，如果 MAC 不可用）的 UUID 版本 6（与 UUIDv1 字段兼容）。
```go
import "github.com/mikhaildadaev/uuid"
u6 := uuid.NewV6()
fmt.Println(u6.Version())
```
Output:
```text
6
```

## NewV7
创建一个 UUID 版本 7（基于时间戳，按字典顺序可排序），使用当前 Unix 时间戳（以毫秒为单位）和随机位。
```go
import "github.com/mikhaildadaev/uuid"
u7 := uuid.NewV7()
fmt.Println(u7.Version())
```
Output:
```text
7
```

## NewV8
创建一个 UUID 版本 8（供应商特定，自定义），使用节点标识符和当前时间戳。
```go
import "github.com/mikhaildadaev/uuid"
u8 := uuid.NewV8(node)
fmt.Println(u8.Version())
```
Output:
```text
8
```
