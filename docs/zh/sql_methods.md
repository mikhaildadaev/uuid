---
outline: deep
---

# API / SQL 集成 / 方法

::: info **关于**
本页记录了 UUID 的原生 `database/sql` 集成。`Scan` 和 `Value` 方法使 UUID 能够与 SQL 数据库无缝协作，包括对可空列的 `NULL` 支持。
:::

## UUID Scan
实现 `sql.Scanner` 接口。从数据库值解码 UUID——接受字符串、字节切片或 nil (NULL)。如果值无法解析则返回错误
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
实现 `driver.Valuer` 接口。将 UUID 编码为适合数据库存储的值。对于空 UUID 返回 nil。
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
为可空 UUID 列实现 `sql.Scanner` 接口。接受 `nil` 表示 SQL NULL，或有效的 UUID 字符串/字节。当 UUID 存在时将 `Valid` 设置为 `true`，为 NULL 时设置为 `false`。
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
为可空 UUID 列实现 `driver.Valuer` 接口。当 `Valid` 为 `false` 时返回 `nil`，当 `Valid` 为 `true` 时返回 UUID 字符串。
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