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