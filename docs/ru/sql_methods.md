---
outline: deep
---

# API / Интеграция с SQL / Методы

::: info **Информация**
На этой странице описана нативная интеграция с `database/sql` для UUID. Методы `Scan` и `Value` позволяют UUID беспрепятственно работать с базами данных SQL, включая поддержку `NULL` для nullable-колонок.
:::

## UUID Scan
Реализует интерфейс `sql.Scanner`. Декодирует UUID из значения базы данных — принимает строку, байтовый срез или nil (NULL). Возвращает ошибку, если значение не может быть разобрано.
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
Реализует интерфейс `driver.Valuer`. Кодирует UUID в значение, пригодное для хранения в базе данных. Возвращает nil для null UUID.
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
Реализует интерфейс `sql.Scanner` для nullable-колонок UUID. Принимает `nil` для SQL NULL или валидную строку/байты UUID. Устанавливает `Valid` в `true,` когда UUID присутствует, и в `false` для NULL.
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
Реализует интерфейс `driver.Valuer` для nullable-колонок UUID. Возвращает `nil`, когда `Valid` равен `false`, или строку UUID, когда `Valid` равен `true`.
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