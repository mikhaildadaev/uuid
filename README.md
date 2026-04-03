[![Go Reference](https://pkg.go.dev/badge/github.com/mikhaildadaev/uuid.svg)](https://pkg.go.dev/github.com/mikhaildadaev/uuid)
[![Go Report Card](https://goreportcard.com/badge/github.com/mikhaildadaev/uuid)](https://goreportcard.com/report/github.com/mikhaildadaev/uuid)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/mikhaildadaev/uuid/blob/main/LICENSE.md)
# UUID Generator

Go package for generating, parsing, validating, and serializing UUIDs (versions 1-8) with optional NULL support.

## License

[MIT License](https://github.com/mikhaildadaev/uuid/blob/main/LICENSE.md)

## Links

- [GoDoc](https://pkg.go.dev/github.com/mikhaildadaev/uuid)
- [GitHub](https://github.com/mikhaildadaev/uuid)
- [Report](https://goreportcard.com/report/github.com/mikhaildadaev/uuid)

## Installation

```bash
go get github.com/mikhaildadaev/uuid
```

## Usage

```go
import (
    "fmt"
    "github.com/mikhaildadaev/uuid"
)

func main() {

    name := "github.com/mikhaildadaev/uuid"
    node := 1995
    posix := 0

    u1 := uuid.V1()
    u2 := uuid.V2(posix)
    u3 := uuid.V3(uuid.NameSpaceURL,name)
    u4 := uuid.V4()
    u5 := uuid.V5(uuid.NameSpaceURL,name)
    u6 := uuid.V6()
    u7 := uuid.V7()
    u8 := uuid.V8(node)

    fmt.Println(u1.String())
    fmt.Println(u2.String())
    fmt.Println(u3.String())
    fmt.Println(u4.String())
    fmt.Println(u5.String())
    fmt.Println(u6.String())
    fmt.Println(u7.String())
    fmt.Println(u8.String())
}
```

## Quick API

### Functions
- uuid.Null()
- uuid.Parse(uuid)
- uuid.V1()
- uuid.V2(posix)
- uuid.V3(namespace)
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

- uuid.MarshalBinary()
- uuid.MarshalJSON()
- uuid.MarshalText()
- uuid.UnmarshalBinary()
- uuid.UnmarshalJSON()
- uuid.UnmarshalText()

## Features

- Generate UUIDv1..v8 и null UUID;
- Parse and validate standard and null UUID;
- Serialize and deserialize via binary/text/json;
- Extract metadata UUID;
- Fast, allocator-free core routines;

## Limits

- V2: `posix` should be 0..255;
- V3/V5: `name` should be 0..36 symbols;
- V8: `node` should be 0..16383;

## Tests and Benchmarks

Run:

```bash
go test ./...
go test -bench=.
```

|  Version | Time (ns/op) | Memory | Allocs |
|----------|--------------|--------|--------|
|  **V1**  |        85.46 |    0 B |      0 |
|  **V2**  |        36.61 |    0 B |      0 |
|  **V3**  |       117.30 |    0 B |      0 |
|  **V4**  |        44.28 |    0 B |      0 |
|  **V5**  |       152.60 |    0 B |      0 |
|  **V6**  |        85.78 |    0 B |      0 |
|  **V7**  |       109.50 |    0 B |      0 |
|  **V8**  |       109.10 |    0 B |      0 |

## Example

```go
import (
    "fmt"
    "github.com/mikhaildadaev/uuid"
)

func main() {

    node := 1995

    u := uuid.V4()
    fmt.Println("UUID v4:", u)

    u2, err := uuid.Parse("01968727-8c7e-8000-87cb-bdba4f634d9f")
    if err != nil {
        panic(err)
    }

    fmt.Println("Parsed:", u2)
}
``` 
