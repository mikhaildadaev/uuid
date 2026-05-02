---
outline: deep
---

# API / 序列化 / 方法

::: info **关于**
本页记录了 UUID 与 **Binary**、**Json** 和 **Text** 格式之间的序列化和反序列化。每个方法适用于所有 UUID 版本 — V1 至 V8 以及 Null。
:::

## NULLUUID MarshalBinary
将 NullUUID 编码为其 16 字节的二进制表示。对于空值返回全零。
```go
var nu uuid.NullUUID
data, err := nu.MarshalBinary()
if err != nil {
    fmt.Println(err)
}
fmt.Printf("%x\n", data)
uuidV8String := "01968727-8c7e-8000-87cb-bdba4f634d9f"
nu.Scan(uuidV8String)
data, err = nu.MarshalBinary()
if err != nil {
    fmt.Println(err)
}
fmt.Printf("%x\n", data)
```
Output:
```text

019687278c7e800087cbbdba4f634d9f
```

## NULLUUID MarshalJson
将 NullUUID 编码为 Json 字符串。对于空值返回 null。
```go
var nu uuid.NullUUID
data, err := nu.MarshalJson()
if err != nil {
    fmt.Println(err)
}
fmt.Println(string(data))
uuidV8String := "01968727-8c7e-8000-87cb-bdba4f634d9f"
nu.Scan(uuidV8String)
data, err := nu.MarshalJson()
if err != nil {
    fmt.Println(err)
}
fmt.Println(string(data))
```
Output:
```text
null
"01968727-8c7e-8000-87cb-bdba4f634d9f"
```

## NULLUUID MarshalText
将 NullUUID 编码为其规范的文本形式。对于空值返回空字符串。
```go
var nu uuid.NullUUID
data, err := nu.MarshalText()
if err != nil {
    fmt.Println(err)
}
fmt.Println(string(data))
uuidV8String := "01968727-8c7e-8000-87cb-bdba4f634d9f"
nu.Scan(uuidV8String)
data, err := nu.MarshalText()
if err != nil {
    fmt.Println(err)
}
fmt.Println(string(data))
```
Output:
```text

01968727-8c7e-8000-87cb-bdba4f634d9f
```

## NULLUUID UnmarshalBinary
从 16 字节的二进制表示解码 NullUUID。根据输入设置 Valid。
```go
var nu uuid.NullUUID
data := []byte{}
if err := nu.UnmarshalBinary(data); err != nil {
	fmt.Println(err)
}
fmt.Println("Valid:", nu.Valid)
fmt.Println("UUID:", nu.UUID)
uuidV8Binary := []byte{0x01, 0x96, 0x87, 0x27, 0x8c, 0x7e, 0x80, 0x00, 0x87, 0xcb, 0xbd, 0xba, 0x4f, 0x63, 0x4d, 0x9f}
err := nu.UnmarshalBinary(uuidV8Binary)
if err != nil {
    fmt.Println(err)
}
fmt.Println("Valid:", nu.Valid)
fmt.Println("UUID:", nu.UUID)
```
Output:
```text
Valid: false
UUID: 00000000-0000-0000-0000-000000000000
Valid: true
UUID: 01968727-8c7e-8000-87cb-bdba4f634d9f
```

## NULLUUID UnmarshalJson
从 Json 编码的字符串解码 NullUUID。接受 null 表示 SQL NULL 或有效的 UUID 字符串。
```go
var nu uuid.NullUUID
data := []byte(`null`)
if err := nu.UnmarshalJson(data); err != nil {
	fmt.Println(err)
}
fmt.Println("Valid:", nu.Valid)
fmt.Println("UUID:", nu.UUID)
uuidV8Json := []byte(`"01968727-8c7e-8000-87cb-bdba4f634d9f"`)
err := nu.UnmarshalJson(uuidV8Json)
if err != nil {
    fmt.Println(err)
}
fmt.Println("Valid:", nu.Valid)
fmt.Println("UUID:", nu.UUID)
```
Output:
```text
Valid: false
UUID: 00000000-0000-0000-0000-000000000000
Valid: true
UUID: 01968727-8c7e-8000-87cb-bdba4f634d9f
```

## NULLUUID UnmarshalText
从规范的文本形式解码 NullUUID。接受空字符串表示 SQL NULL。
```go
var nu uuid.NullUUID
data := []byte("")
if err := nu.UnmarshalText(data); err != nil {
	fmt.Println(err)
}
fmt.Println("Valid:", nu.Valid)
fmt.Println("UUID:", nu.UUID)
uuidV8Text := []byte("01968727-8c7e-8000-87cb-bdba4f634d9f")
err := nu.UnmarshalText(uuidV8Text)
if err != nil {
    fmt.Println(err)
}
fmt.Println("Valid:", nu.Valid)
fmt.Println("UUID:", nu.UUID)
```
Output:
```text
Valid: false
UUID: 00000000-0000-0000-0000-000000000000
Valid: true
UUID: 01968727-8c7e-8000-87cb-bdba4f634d9f
```

## UUID MarshalBinary
将 UUID 编码为其 16 字节的二进制表示。
```go
uuidV8String := "01968727-8c7e-8000-87cb-bdba4f634d9f"
uu, err := uuid.Parse(uuidV8String)
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
019687278c7e800087cbbdba4f634d9f
```

## UUID MarshalJson
将 UUID 编码为标准 8-4-4-4-12 格式的 Json 字符串。
```go
uuidV8String := "01968727-8c7e-8000-87cb-bdba4f634d9f"
uu, err := uuid.Parse(uuidV8String)
if err != nil {
    fmt.Println(err)
}
data, err := uu.MarshalJson()
if err != nil {
    fmt.Println(err)
}
fmt.Println(string(data))
```
Output:
```text
"01968727-8c7e-8000-87cb-bdba4f634d9f"
```

## UUID MarshalText
将 UUID 编码为其规范的文本形式：32 个十六进制数字，用连字符分隔。
```go
uuidV8String := "01968727-8c7e-8000-87cb-bdba4f634d9f"
uu, err := uuid.Parse(uuidV8String)
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
01968727-8c7e-8000-87cb-bdba4f634d9f
```

## UUID UnmarshalBinary
从 16 字节的二进制表示解码 UUID。输入如果不是 16 字节将返回错误。
```go
var uu uuid.UUID
uuidV8Binary := []byte{0x01, 0x96, 0x87, 0x27, 0x8c, 0x7e, 0x80, 0x00, 0x87, 0xcb, 0xbd, 0xba, 0x4f, 0x63, 0x4d, 0x9f}
err := uu.UnmarshalBinary(uuidV8Binary)
if err != nil {
    fmt.Println(err)
}
fmt.Println(uu.String())
```
Output:
```text
01968727-8c7e-8000-87cb-bdba4f634d9f
```

## UUID UnmarshalJson
从 Json 编码的字符串解码 UUID。接受带引号和不带引号的形式；格式无效时返回错误。
```go
var uu uuid.UUID
uuidV8Json := []byte(`"01968727-8c7e-8000-87cb-bdba4f634d9f"`)
err := uu.UnmarshalJson(uuidV8Json)
if err != nil {
    fmt.Println(err)
}
fmt.Println(uu.String())
```
Output:
```text
01968727-8c7e-8000-87cb-bdba4f634d9f
```

## UUID UnmarshalText
从规范的文本形式（32 个十六进制数字，用连字符分隔）解码 UUID。不区分大小写；格式无效时返回错误。
```go
var uu uuid.UUID
uuidV8Text := []byte("01968727-8c7e-8000-87cb-bdba4f634d9f")
err := uu.UnmarshalText(uuidV8Text)
if err != nil {
    fmt.Println(err)
}
fmt.Println(uu.String())
```
Output:
```text
01968727-8c7e-8000-87cb-bdba4f634d9f
```
