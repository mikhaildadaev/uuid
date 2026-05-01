---
outline: deep
---

# API / Marshal / 主要

::: info 关于
本页记录了 UUID 与 **Binary**、**JSON** 和 **Text** 格式之间的序列化和反序列化。每个方法适用于所有 UUID 版本 — V1 至 V8 以及 Null。
:::

## MarshalBinary
将 UUID 编码为其 16 字节的二进制表示。
```go
import "github.com/mikhaildadaev/uuid"
uu, err := uuid.Parse("018f3c14-8000-0000-0000-000000000001")
if err != nil {
    fmt.Println(err)
}
data, err := uu.MarshalBinary()
if err != nil {
    fmt.Println(err)
}
fmt.Printf("%x\n", data)
```
Output:
```text
018f3c1480000000000000000000000001
```

## MarshalJson
将 UUID 编码为标准 8-4-4-4-12 格式的 JSON 字符串。
```go
import "github.com/mikhaildadaev/uuid"
uu, err := uuid.Parse("018f3c14-8000-0000-0000-000000000001")
if err != nil {
    fmt.Println(err)
}
data, err := uu.MarshalJSON()
if err != nil {
    fmt.Println(err)
}
fmt.Println(string(data))
```
Output:
```text
018f3c14-8000-0000-0000-000000000001
```

## MarshalText
将 UUID 编码为其规范的文本形式：32 个十六进制数字，用连字符分隔。
```go
import "github.com/mikhaildadaev/uuid"
uu, err := uuid.Parse("018f3c14-8000-0000-0000-000000000001")
if err != nil {
    fmt.Println(err)
}
data, err := uu.MarshalText()
if err != nil {
    fmt.Println(err)
}
fmt.Println(string(data))
```
Output:
```text
018f3c14-8000-0000-0000-000000000001
```

## UnmarshalBinary
从 16 字节的二进制表示解码 UUID。输入如果不是 16 字节将返回错误。
```go
import "github.com/mikhaildadaev/uuid"
var u uuid.UUID
bin := []byte{0x01, 0x8f, 0x3c, 0x14, 0x00, 0x00, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}
err := u.UnmarshalBinary(bin)
if err != nil {
    fmt.Println(err)
}
fmt.Println(u.String())
```
Output:
```text
018f3c14-0000-8000-0000-000000000001
```

## UnmarshalJson
从 JSON 编码的字符串解码 UUID。接受带引号和不带引号的形式；格式无效时返回错误。
```go
import "github.com/mikhaildadaev/uuid"
var u uuid.UUID
err := u.UnmarshalJSON([]byte(`"018f3c14-0000-8000-0000-000000000001"`))
if err != nil {
    fmt.Println(err)
}
fmt.Println(u.String())
```
Output:
```text
018f3c14-0000-8000-0000-000000000001
```

## UnmarshalText
从规范的文本形式（32 个十六进制数字，用连字符分隔）解码 UUID。不区分大小写；格式无效时返回错误。
```go
import "github.com/mikhaildadaev/uuid"
var u uuid.UUID
err := u.UnmarshalText([]byte("018f3c14-0000-8000-0000-000000000001"))
if err != nil {
    fmt.Println(err)
}
fmt.Println(u.String())
```
Output:
```text
018f3c14-0000-8000-0000-000000000001
```