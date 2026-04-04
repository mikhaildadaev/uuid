[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/mikhaildadaev/uuid/blob/main/LICENSE.md)
[![Go Reference](https://pkg.go.dev/badge/github.com/mikhaildadaev/uuid.svg)](https://pkg.go.dev/github.com/mikhaildadaev/uuid)
[![Go Report Card](https://goreportcard.com/badge/github.com/mikhaildadaev/uuid)](https://goreportcard.com/report/github.com/mikhaildadaev/uuid)
[![Go Version](https://img.shields.io/github/go-mod/go-version/mikhaildadaev/uuid)](https://github.com/mikhaildadaev/uuid)
[![CI](https://github.com/mikhaildadaev/uuid/actions/workflows/ci.yml/badge.svg)](https://github.com/mikhaildadaev/uuid/actions/workflows/ci.yml)

# UUID Generator

A high-performance, zero-allocation UUID library for Go supporting versions 1 through 8, NULL UUID, and full serialization (Binary, JSON, Text) with SQL database integration.

## Features

- Generate UUIDv1..v8 and null UUID;
- Parse and validate standard and null UUID;
- Serialize and deserialize via binary/text/json;
- Extract metadata UUID;
- Fast, allocator-free core routines;

## Installation

```bash
go get github.com/mikhaildadaev/uuid
```

## Quick API

### Functions
#### API
- uuid.Null()
- uuid.Parse(uuid)
- uuid.V1()
- uuid.V2(posix)
- uuid.V3(namespace, name)
- uuid.V4()
- uuid.V5(namespace, name)
- uuid.V6()
- uuid.V7()
- uuid.V8(node)

### Methods
#### API
- uuid.Bytes()
- uuid.Equal(other)
- uuid.Info()
- uuid.IsZero()
- uuid.Node()
- uuid.Posix()
- uuid.Sequence()
- uuid.String()
- uuid.Timestamp()
- uuid.Validate()
- uuid.Variant()
- uuid.Version()
- nulluuid.IsZero()
- nulluuid.String()
- nulluuid.Validate()
#### Marshal
- uuid.MarshalBinary()
- uuid.MarshalJSON()
- uuid.MarshalText()
- uuid.UnmarshalBinary(data)
- uuid.UnmarshalJSON(data)
- uuid.UnmarshalText(data)
- nulluuid.MarshalBinary()
- nulluuid.MarshalJSON()
- nulluuid.MarshalText()
- nulluuid.UnmarshalBinary(data)
- nulluuid.UnmarshalJSON(data)
- nulluuid.UnmarshalText(data)
#### SQL
- uuid.Scan(src)
- uuid.Value()
- nulluuid.Scan(src)
- nulluuid.Value()

## Performance

|  Version | Operations | Time (ns/op) | Memory (B/op) | Allocs (allocs/op) |
|----------|------------|--------------|---------------|--------------------|
|  **V1**  |      13.5M |        85.46 |             0 |                  0 |
|  **V2**  |      32.1M |        36.61 |             0 |                  0 |
|  **V3**  |       9.9M |       117.30 |             0 |                  0 |
|  **V4**  |      26.7M |        44.28 |             0 |                  0 |
|  **V5**  |       7.7M |       152.60 |             0 |                  0 |
|  **V6**  |      13.7M |        85.78 |             0 |                  0 |
|  **V7**  |      10.7M |       109.50 |             0 |                  0 |
|  **V8**  |      10.2M |       109.10 |             0 |                  0 |

*Benchmarked on Intel Core i9-9880H (2.30 GHz)*

## Usage

```go
import (
    "fmt"
    "log"
    "github.com/mikhaildadaev/uuid"
)

func main() {
    name := "github.com/mikhaildadaev/uuid"
    node := 1995
    posix := 0
    uid := "01968727-8c7e-8000-87cb-bdba4f634d9f"
    // Generate all UUID versions
    fmt.Println(uuid.V1())
    fmt.Println(uuid.V2(posix))
    fmt.Println(uuid.V3(uuid.NameSpaceURL,name))
    fmt.Println(uuid.V4())
    fmt.Println(uuid.V5(uuid.NameSpaceURL,name))
    fmt.Println(uuid.V6())
    fmt.Println(uuid.V7())
    fmt.Println(uuid.V8(node))
    // Parse existing UUID
    parsed, err := uuid.Parse(uid)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Parsed:", parsed)
    // Check if UUID is zero
    fmt.Println("Is zero:", parsed.IsZero())
    // Using NullUUID for database NULL values
    var nu uuid.NullUUID
    if err := nu.Scan(nil); err != nil {
        log.Fatal(err)
    }
    fmt.Println("NULL valid:", nu.Valid)
    if err := nu.Scan(uid); err != nil {
        log.Fatal(err)
    }
    fmt.Println("UUID valid:", nu.Valid)
}
```

## Limits

- V2: `posix` should be 0..255;
- V3/V5: `name` should be 0..36 symbols;
- V8: `node` should be 0..16383;

## Tests and Benchmarks

Run:

```bash
go test ./...
go test -bench=. ./...
go test -cover ./...
go test -race ./...
```
