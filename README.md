[![Go Reference](https://pkg.go.dev/badge/github.com/mikhaildadaev/uuid.svg)](https://pkg.go.dev/github.com/mikhaildadaev/uuid)
[![Go Report Card](https://goreportcard.com/badge/github.com/mikhaildadaev/uuid)](https://goreportcard.com/report/github.com/mikhaildadaev/uuid)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/mikhaildadaev/uuid/blob/main/LICENSE.md)
# UUID Generator / Генератор UUID

Go package for generating, parsing, validating, and serializing UUIDs (versions 1-8) with optional NULL support. / Go Пакет для генерации, разбора, валидации и сериализации UUID (версии 1-8), включая поддержу NULL-UUID.

---

## Installation / Установка

```bash
go get github.com/mikhaildadaev/uuid
```

## Usage / Использование

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

## Quick API / Быстрый API

- uuid.V1(), uuid.V2(t), uuid.V3(ns, name), uuid.V4(), uuid.V5(ns, name), uuid.V6(), uuid.V7(), uuid.V8(node)
- uuid.Parse(string) -> UUID or error
- uuid.Null() -> UUIDNULL
- u.Validate(), u.IsZero(), u.Equal(other), u.MarshalBinary(), u.UnmarshalBinary(), u.MarshalJSON(), u.UnmarshalJSON(), u.MarshalText(), u.UnmarshalText()
- u.Version(), u.Variant(), u.Timestamp(), u.Info(), u.Node(), u.POS(), u.Sequence()

## Features / Возможности

- Generate UUIDv1..v8 и null UUID / Генерация UUIDv1..v8 и пустых UUID
- Parse and validate standard and null UUID | Парсинг и валидация стандартного и нулевого UUID
- Serialize and deserialize via binary|text|json / Сериализация и десериализация в бинарный, текстовый и JSON форматы.
- Extract metadata (version|variant|timestamp|node|POS|sequence) / Получение метаинформации (версия, вариант, таймстамп, узел, POS, последовательность) 
- Fast, allocator-free core routines / Быстрые основные процедуры, не требующие выделения ресурсов

## Limits / Ограничения

- V2: `posix` should be 0..255 / должно быть 0..255
- V3/V5: `name` should be 0..36 symbols / должно быть 0..36 символов
- V8: `node` should be 0..16383 / должно быть 0..16383

## Benchmarks / Производительность

| Version | Time (ns/op) | Memory | Allocs |
|---------|--------------|--------|--------|
| **V1**  |        85.46 |    0 B |      0 |
| **V2**  |        36.61 |    0 B |      0 |
| **V3**  |       117.30 |    0 B |      0 |
| **V4**  |        44.28 |    0 B |      0 |
| **V5**  |       152.60 |    0 B |      0 |
| **V6**  |        85.78 |    0 B |      0 |
| **V7**  |       109.50 |    0 B |      0 |
| **V8**  |       109.10 |    0 B |      0 |

Benchmarks are available in `uuid_bench_test.go` / Бенчмарки доступны в файле `uuid_bench_test.go`.

---

## Example / Пример

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

## Tests and CI / Тесты и CI

Run:

```bash
go test ./...
go test -bench=.
```

---

## License / Лицензия

MIT License - `LICENSE.md`

## Links / Ссылки

- [GoDoc](https://pkg.go.dev/github.com/mikhaildadaev/uuid)
- [GitHub](https://github.com/mikhaildadaev/uuid)
- [Report](https://goreportcard.com/report/github.com/mikhaildadaev/uuid)
