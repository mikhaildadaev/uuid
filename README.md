[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/mikhaildadaev/uuid/blob/main/LICENSE.md)
[![Go Version](https://img.shields.io/github/go-mod/go-version/mikhaildadaev/uuid)](https://github.com/mikhaildadaev/uuid)
[![Go Reference](https://pkg.go.dev/badge/github.com/mikhaildadaev/uuid.svg)](https://pkg.go.dev/github.com/mikhaildadaev/uuid)
[![Go Report Card](https://goreportcard.com/badge/github.com/mikhaildadaev/uuid)](https://goreportcard.com/report/github.com/mikhaildadaev/uuid)
[![CI](https://github.com/mikhaildadaev/uuid/actions/workflows/ci.yml/badge.svg)](https://github.com/mikhaildadaev/uuid/actions/workflows/ci.yml)

# UUID

A high-performance, zero-allocation platform for UUID versions 1-8.

## Get Started
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
- Generate UUIDv1..v8 and null UUID;
- Parse and validate standard and null UUID;
- Serialize and deserialize via Binary/Json/Text;
- Extract metadata UUID;
- Fast, allocator-free core routines;

## Limits
- V2: `posix` should be `0..255`;
- V3/V5: `name` should be `1..512` symbols (byte length). 
- V8: `node` should be `0..16383`;
> **Note** 
>
> Strings `name` that are 0 bytes or larger than 512 bytes return a null UUID to prevent allocation.

## Benchmarks

### Core Performance

#### MultiThread
|  Version | Operations | Time (ns/op) | Memory (B/op) | Allocs (allocs/op) |
|----------|------------|--------------|---------------|--------------------|
|  **V1**  |      34.1M |        29.29 |             0 |                  0 |
|  **V2**  |      59.6M |        16.79 |             0 |                  0 |
|  **V3**  |      84.7M |        11.81 |             0 |                  0 |
|  **V4**  |      18.9M |        52.88 |             0 |                  0 |
|  **V5**  |      51.2M |        19.51 |             0 |                  0 |
|  **V6**  |      32.5M |        30.75 |             0 |                  0 |
|  **V7**  |      22.5M |        44.45 |             0 |                  0 |
|  **V8**  |      32.0M |        31.21 |             0 |                  0 |

#### SingleThread
|  Version | Operations | Time (ns/op) | Memory (B/op) | Allocs (allocs/op) |
|----------|------------|--------------|---------------|--------------------|
|  **V1**  |      13.5M |        85.46 |             0 |                  0 |
|  **V2**  |      32.1M |        36.61 |             0 |                  0 |
|  **V3**  |       9.9M |       117.30 |             7 |                  0 |
|  **V4**  |      26.7M |        44.28 |             0 |                  0 |
|  **V5**  |       7.7M |       152.60 |             7 |                  0 |
|  **V6**  |      13.7M |        85.78 |             0 |                  0 |
|  **V7**  |      10.7M |       109.50 |             0 |                  0 |
|  **V8**  |      10.2M |       109.10 |             0 |                  0 |

> **Note**
> 
> - All benchmarks measure pure generation and parsing overhead.
> - Zero allocations (`0 B/op`, `0 allocs/op`) indicate that no heap memory is allocated during hot-path operations. The 7 B/op in single-threaded V3/V5 benchmarks is an artifact of the Go testing framework's string allocation for the benchmark loop itself, not an allocation within the UUID generation hot path. This is confirmed by `0 allocs/op` and the absence of this memory in multi-threaded benchmarks.
> - The `Info` benchmarks (e.g., `Benchmark_VX_Info`) include logging of generated UUIDs via `b.Logf`, which adds measurable overhead and is primarily used for debugging and validation, not for performance comparisons.
> - Real-world performance may vary based on CPU frequency, memory latency, and concurrency level.
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
    fmt.Println(uuid.V1())
    fmt.Println(uuid.V2(posix))
    fmt.Println(uuid.V3(uuid.NameSpaceURL,name))
    fmt.Println(uuid.V4())
    fmt.Println(uuid.V5(uuid.NameSpaceURL,name))
    fmt.Println(uuid.V6())
    fmt.Println(uuid.V7())
    fmt.Println(uuid.V8(node))
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
- uuid.V1()
- uuid.V2(posix)
- uuid.V3(namespace, name)
- uuid.V4()
- uuid.V5(namespace, name)
- uuid.V6()
- uuid.V7()
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