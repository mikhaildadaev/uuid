---
outline: deep
---

# API / Интеграция с SQL / Методы

::: info **Информация**
На этой странице описана нативная интеграция с `database/sql` для UUID. Методы `Scan` и `Value` позволяют UUID беспрепятственно работать с базами данных SQL, включая поддержку `NULL` для nullable-колонок.
:::

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
fmt.Println("UUID:", nu.UUID)
uuidV8String := "01968727-8c7e-8000-87cb-bdba4f634d9f"
if err := nu.Scan(uuidV8String); err != nil {
    fmt.Println(err)
}
fmt.Println("Valid:", nu.Valid)
fmt.Println("UUID:", nu.UUID)
```
Output
```text
Valid: false
UUID: 00000000-0000-0000-0000-000000000000
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
fmt.Printf(value)
uuidV8String := "01968727-8c7e-8000-87cb-bdba4f634d9f"
nu.Scan(uuidV8String)
value, _ = nu.Value()
fmt.Printf(value)
```
Output
```text
<nil>
01968727-8c7e-8000-87cb-bdba4f634d9f
```

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
uuidV8String := "01968727-8c7e-8000-87cb-bdba4f634d9f"
err := uu.Scan(uuidV8String)
if err != nil {
    fmt.Println(err)
}
fmt.Println(uu.String())
```
Output
```text
01968727-8c7e-8000-87cb-bdba4f634d9f
```

## UUID Value
Реализует интерфейс `driver.Valuer`. Кодирует UUID в значение, пригодное для хранения в базе данных. Возвращает nil для null UUID.
```go
import (
    "database/sql"
    "fmt"
    "github.com/mikhaildadaev/uuid"
)
uuidV8String := "01968727-8c7e-8000-87cb-bdba4f634d9f"
uu, err := uuid.Parse(uuidV8String)
if err != nil {
    fmt.Println(err)
}
value, err := uu.Value()
if err != nil {
    fmt.Println(err)
}
fmt.Println(value)
```
Output
```text
01968727-8c7e-8000-87cb-bdba4f634d9f
```
