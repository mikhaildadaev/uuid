---
outline: deep
---

# API / SQL / Methods

::: info **Info**
This page documents the native `database/sql` integration for UUIDs. Both `Scan` and `Value` methods make UUID work seamlessly with SQL databases, including `NULL` support for nullable columns.
:::

## UUID Scan
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

## UUID Value
Implements the `driver.Valuer` interface. Encodes the UUID into a value suitable for database storage. Returns `nil` for a null UUID.
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

## NULLUUID Scan
Implements the `sql.Scanner` interface for nullable UUID columns. Accepts `nil` for SQL NULL or a valid UUID string/bytes. Sets `Valid` to `true` when a UUID is present, `false` for NULL.
```go
import (
    "database/sql"
    "fmt"
    "github.com/mikhaildadaev/uuid"
)
var nu uuid.NullUUID
if err := nu.Scan(nil); err != nil {
    fmt.Println(err)
}
fmt.Println("Valid:", nu.Valid)
if err := nu.Scan("01968727-8c7e-8000-87cb-bdba4f634d9f"); err != nil {
    fmt.Println(err)
}
fmt.Println("Valid:", nu.Valid)
fmt.Println("UUID:", nu.UUID)
```
Output
```text
Valid: false
Valid: true
UUID: 01968727-8c7e-8000-87cb-bdba4f634d9f
```

## NULLUUID Value
Implements the `driver.Valuer` interface for nullable UUID columns. Returns `nil` when `Valid` is `false`, or the UUID string when `Valid` is `true`.
```go
import (
    "database/sql"
    "fmt"
    "github.com/mikhaildadaev/uuid"
)
var nu uuid.NullUUID
value, _ := nu.Value()
fmt.Println("NULL value:", value)
nu.Scan("01968727-8c7e-8000-87cb-bdba4f634d9f")
value, _ = nu.Value()
fmt.Println("UUID value:", value)
```
Output
```text
NULL value: <nil>
UUID value: 01968727-8c7e-8000-87cb-bdba4f634d9f
```