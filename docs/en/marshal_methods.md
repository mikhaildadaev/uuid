---
outline: deep
---

# API / Marshal / Methods

::: info **Info**
This page documents serialization and deserialization of UUIDs to and from **Binary**, **Json**, and **Text** formats. Every marshal/unmarshal method works with all UUID versions — V1 through V8 and Null.
:::

## NULLUUID MarshalBinary
Encodes the NullUUID into its 16-byte binary representation. Returns all zeros for a null value.
```go
import "github.com/mikhaildadaev/uuid"
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
Encodes the NullUUID into a Json string. Returns null for a null value.
```go
import "github.com/mikhaildadaev/uuid"
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
Encodes the NullUUID into its canonical text form. Returns an empty string for a null value.
```go
import "github.com/mikhaildadaev/uuid"
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
Decodes a NullUUID from its 16-byte binary representation. Sets Valid based on the input.
```go
import "github.com/mikhaildadaev/uuid"
var nu uuid.NullUUID
data := []byte{}
if err := nu.UnmarshalBinary(data); err != nil {
	fmt.Println(err)
}
fmt.Println("Valid:", nu.Valid)
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
Valid: true
UUID: 01968727-8c7e-8000-87cb-bdba4f634d9f
```

## NULLUUID UnmarshalJson
Decodes a NullUUID from a Json-encoded string. Accepts null for SQL NULL or a valid UUID string.
```go
import "github.com/mikhaildadaev/uuid"
var nu uuid.NullUUID
data := []byte(`null`)
if err := nu.UnmarshalJson(data); err != nil {
	fmt.Println(err)
}
fmt.Println("Valid:", nu.Valid)
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
Valid: true
UUID: 01968727-8c7e-8000-87cb-bdba4f634d9f
```

## NULLUUID UnmarshalText
Decodes a NullUUID from its canonical text form. Accepts an empty string for SQL NULL.
```go
import "github.com/mikhaildadaev/uuid"
var nu uuid.NullUUID
data := []byte("")
if err := nu.UnmarshalText(data); err != nil {
	fmt.Println(err)
}
fmt.Println("Valid:", nu.Valid)
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
Valid: true
UUID: 01968727-8c7e-8000-87cb-bdba4f634d9f
```

## UUID MarshalBinary
Encodes the UUID into its 16-byte binary representation.
```go
import "github.com/mikhaildadaev/uuid"
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
Encodes the UUID into a Json string in the standard 8-4-4-4-12 format.
```go
import "github.com/mikhaildadaev/uuid"
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
Encodes the UUID into its canonical text form: 32 hex digits with hyphens.
```go
import "github.com/mikhaildadaev/uuid"
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
Decodes a UUID from its 16-byte binary representation. An error is returned if the input is not exactly 16 bytes long.
```go
import "github.com/mikhaildadaev/uuid"
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
Decodes a UUID from a Json-encoded string. Accepts both quoted and unquoted forms; returns an error if the format is invalid.
```go
import "github.com/mikhaildadaev/uuid"
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
Decodes a UUID from its canonical text form (32 hex digits with hyphens). Case-insensitive; returns an error if the format is invalid.
```go
import "github.com/mikhaildadaev/uuid"
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
