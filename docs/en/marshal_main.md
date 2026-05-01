---
outline: deep
---

# API / Marshal / Main

::: info Info
This page documents serialization and deserialization of UUIDs to and from **Binary**, **JSON**, and **Text** formats. Every marshal/unmarshal method works with all UUID versions — V1 through V8 and Null.
:::

## MarshalBinary
Encodes the UUID into its 16-byte binary representation.
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
Encodes the UUID into a JSON string in the standard 8-4-4-4-12 format.
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
Encodes the UUID into its canonical text form: 32 hex digits with hyphens.
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
Decodes a UUID from its 16-byte binary representation. An error is returned if the input is not exactly 16 bytes long.
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
Decodes a UUID from a JSON-encoded string. Accepts both quoted and unquoted forms; returns an error if the format is invalid.
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
Decodes a UUID from its canonical text form (32 hex digits with hyphens). Case-insensitive; returns an error if the format is invalid.
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