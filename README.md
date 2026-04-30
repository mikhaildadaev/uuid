[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/mikhaildadaev/uuid/blob/main/LICENSE.md)
[![Go Version](https://img.shields.io/github/go-mod/go-version/mikhaildadaev/uuid)](https://github.com/mikhaildadaev/uuid)
[![Go Reference](https://pkg.go.dev/badge/github.com/mikhaildadaev/uuid.svg)](https://pkg.go.dev/github.com/mikhaildadaev/uuid)
[![Go Report Card](https://goreportcard.com/badge/github.com/mikhaildadaev/uuid)](https://goreportcard.com/report/github.com/mikhaildadaev/uuid)
[![CI](https://github.com/mikhaildadaev/uuid/actions/workflows/ci.yml/badge.svg)](https://github.com/mikhaildadaev/uuid/actions/workflows/ci.yml)

# UUID

A high-performance, zero-allocation platform for UUID versions 1-8.

## Go

> **Info**
> 
> The latest stable version of `uuid` is **v1.26.11**.

```bash
go get github.com/mikhaildadaev/uuid
```

### Run Test
```bash
go test ./...
go test -bench=. ./...
go test -cover ./...
go test -race ./...
```

### Key Features
- **UUIDv1..v8 & Null** — Generate all UUID versions and SQL-compatible null UUID;
- **Parse & Validate** — Parse and validate standard and null UUID strings;
- **Metadata Extraction** — Extract timestamp, sequence, node, POSIX, variant, and version;
- **Serialization** — Full Binary, JSON, and Text marshal/unmarshal support;
- **SQL Integration** — Native database/sql driver with Scan/Value interfaces;
- **Zero-Allocation** — Fast, allocator-free core routines on all hot paths.

### Limits
- **V2**: `posix` should be `0..255`;
- **V3/V5**: `name` should be `1..512` symbols (byte length);
- **V8**: `node` should be `0..16383`;

> **Note** 
>
> Strings `name` that are 0 bytes or larger than 512 bytes return a null UUID to prevent allocation.

## Benchmarks
> **Info**
> 
> The best way to compare libraries is to run benchmarks in your own environment with your own workload. Each project has unique requirements — latency, throughput, memory usage, and integration complexity — and no single test can cover them all.
> 
> I recommend that you test `uuid` alongside other libraries and choose the tool that best suits your needs.

### Core Performance
Pure generation and serialization overhead. Benchmarks write to `io.Discard` — no I/O involved.

#### MultiThread
| Version | Operations | Time (ns/op) | Memory (B/op) | Allocs (allocs/op) |
|---------|------------|--------------|---------------|--------------------|
| **V1**  |      34.1M |        29.29 |             0 |                  0 |
| **V2**  |      59.6M |        16.79 |             0 |                  0 |
| **V3**  |      84.7M |        11.81 |             0 |                  0 |
| **V4**  |      18.9M |        52.88 |             0 |                  0 |
| **V5**  |      51.2M |        19.51 |             0 |                  0 |
| **V6**  |      32.5M |        30.75 |             0 |                  0 |
| **V7**  |      22.5M |        44.45 |             0 |                  0 |
| **V8**  |      32.0M |        31.21 |             0 |                  0 |

#### SingleThread
| Version | Operations | Time (ns/op) | Memory (B/op) | Allocs (allocs/op) |
|---------|------------|--------------|---------------|--------------------|
| **V1**  |      13.5M |        85.46 |             0 |                  0 |
| **V2**  |      32.1M |        36.61 |             0 |                  0 |
| **V3**  |       9.9M |       117.30 |             7 |                  0 |
| **V4**  |      26.7M |        44.28 |             0 |                  0 |
| **V5**  |       7.7M |       152.60 |             7 |                  0 |
| **V6**  |      13.7M |        85.78 |             0 |                  0 |
| **V7**  |      10.7M |       109.50 |             0 |                  0 |
| **V8**  |      10.2M |       109.10 |             0 |                  0 |

> **Note**
> 
> All benchmarks measure pure generation and parsing overhead. Benchmarks marked as zero-allocation (`0 B/op`, `0 allocs/op`) perform no heap allocations in hot-path operations — minor artifacts in single-threaded V3/V5 are from the Go testing framework itself, not from UUID generation. Real-world performance depends on CPU, memory latency, and concurrency.
>
> *Benchmarked on Intel Core i9-9880H (2.30 GHz)*

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
    fmt.Println(uuid.NewV1())
    fmt.Println(uuid.NewV2(posix))
    fmt.Println(uuid.NewV3(uuid.NameSpaceURL,name))
    fmt.Println(uuid.NewV4())
    fmt.Println(uuid.NewV5(uuid.NameSpaceURL,name))
    fmt.Println(uuid.NewV6())
    fmt.Println(uuid.NewV7())
    fmt.Println(uuid.NewV8(node))
    parsed, err := uuid.Parse(uid)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Parsed:", parsed)
    fmt.Println("Is zero:", parsed.IsZero())
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

## Quick API

### Functions
- uuid.Null()
- uuid.Parse(uuid)
- uuid.NewV1()
- uuid.NewV2(posix)
- uuid.NewV3(namespace, name)
- uuid.NewV4()
- uuid.NewV5(namespace, name)
- uuid.NewV6()
- uuid.NewV7()
- uuid.V8(node)

### Methods
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
- uuid.MarshalJson()
- uuid.MarshalText()
- uuid.UnmarshalBinary(data)
- uuid.UnmarshalJson(data)
- uuid.UnmarshalText(data)
- nulluuid.MarshalBinary()
- nulluuid.MarshalJson()
- nulluuid.MarshalText()
- nulluuid.UnmarshalBinary(data)
- nulluuid.UnmarshalJson(data)
- nulluuid.UnmarshalText(data)
#### SQL
- uuid.Scan(src)
- uuid.Value()
- nulluuid.Scan(src)
- nulluuid.Value()