---
outline: deep
---

# API / SQL 集成 / 主要

::: info **关于**
本页记录了 UUID 的原生 `database/sql` 集成。`Scan` 和 `Value` 方法使 UUID 能够与 SQL 数据库无缝协作，包括对可空列的 `NULL` 支持。
:::

## Scan
实现 `sql.Scanner` 接口。从数据库值解码 UUID——接受字符串、字节切片或 nil (NULL)。如果值无法解析则返回错误。
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