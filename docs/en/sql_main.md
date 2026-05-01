---
outline: deep
---

# API / SQL / Main

::: info **Info**
This page documents the native `database/sql` integration for UUIDs. Both `Scan` and `Value` methods make UUID work seamlessly with SQL databases, including `NULL` support for nullable columns.
:::

## Scan
Implements the `sql.Scanner` interface. Decodes a UUID from a database value — accepts string, byte slice, or nil (NULL). Returns an error if the value cannot be parsed.
```go
import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    "github.com/mikhaildadaev/uuid"
)
var uu uuid.UUID
err := uu.Scan("018f3c14-8000-0000-0000-000000000001")
if err != nil {
    fmt.Println(err)
}
fmt.Println(uu.String())
```
Output
```text
018f3c14-8000-0000-0000-000000000001
```

## Value
Implements the `driver.Valuer` interface. Encodes the UUID into a value suitable for database storage. Returns nil for a null UUID.
```go
import (
    "database/sql"
    "fmt"
    "github.com/mikhaildadaev/uuid"
)
uu, err := uuid.Parse("018f3c14-8000-0000-0000-000000000001")
if err != nil {
    fmt.Println(err)
}
val, err := uu.Value()
if err != nil {
    fmt.Println(err)
}
fmt.Printf("%s\n", val)
```
Output
```text
018f3c14-8000-0000-0000-000000000001
```