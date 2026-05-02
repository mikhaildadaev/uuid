---
outline: deep
---

# API / SQL / Methods

::: info **Info**
This page documents the native `database/sql` integration for UUIDs. Both `Scan` and `Value` methods make UUID work seamlessly with SQL databases, including `NULL` support for nullable columns.
:::

## NULLUUID Scan
Implements the `sql.Scanner` interface for nullable UUID columns. Accepts `nil` for SQL NULL or a valid UUID string/bytes. Sets `Valid` to `true` when a UUID is present, `false` for NULL.
```go
var nu uuid.NullUUID
if err := nu.Scan(nil); err != nil {
    fmt.Println(err)
}
fmt.Println("Valid:", nu.Valid)
fmt.Println("UUID:", nu.UUID)
uuidV8String := "01968727-8c7e-8000-87cb-bdba4f634d9f"
if err := nu.Scan(uuidV8String); err != nil {
    fmt.Println(err)
}
fmt.Println("Valid:", nu.Valid)
fmt.Println("UUID:", nu.UUID)
```
Output
```text
Valid: false
UUID: 00000000-0000-0000-0000-000000000000
Valid: true
UUID: 01968727-8c7e-8000-87cb-bdba4f634d9f
```

## NULLUUID Value
Implements the `driver.Valuer` interface for nullable UUID columns. Returns `nil` when `Valid` is `false`, or the UUID string when `Valid` is `true`.
```go
var nu uuid.NullUUID
value, _ := nu.Value()
fmt.Println(value)
uuidV8String := "01968727-8c7e-8000-87cb-bdba4f634d9f"
nu.Scan(uuidV8String)
value, _ = nu.Value()
fmt.Println(value)
```
Output
```text
<nil>
01968727-8c7e-8000-87cb-bdba4f634d9f
```

## UUID Scan
Implements the `sql.Scanner` interface. Decodes a UUID from a database value — accepts string, byte slice, or nil (NULL). Returns an error if the value cannot be parsed.
```go
var uu uuid.UUID
uuidV8String := "01968727-8c7e-8000-87cb-bdba4f634d9f"
err := uu.Scan(uuidV8String)
if err != nil {
    fmt.Println(err)
}
fmt.Println(uu.String())
```
Output
```text
01968727-8c7e-8000-87cb-bdba4f634d9f
```

## UUID Value
Implements the `driver.Valuer` interface. Encodes the UUID into a value suitable for database storage. Returns `nil` for a null UUID.
```go
uuidV8String := "01968727-8c7e-8000-87cb-bdba4f634d9f"
uu, err := uuid.Parse(uuidV8String)
if err != nil {
    fmt.Println(err)
}
value, err := uu.Value()
if err != nil {
    fmt.Println(err)
}
fmt.Println(value)
```
Output
```text
01968727-8c7e-8000-87cb-bdba4f634d9f
```
