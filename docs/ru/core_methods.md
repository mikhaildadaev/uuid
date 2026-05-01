---
outline: deep
---

# API / Ядро / Методы

::: info **Информация**
На этой странице описаны все методы экземпляров UUID и NullUUID — от базовых операций, таких как **String** и **Version**, до расширенного извлечения метаданных, таких как **Timestamp**, **Node** и **Sequence**. Каждый метод работает со всеми версиями UUID, от V1 до V8.
:::


## NULLUUID IsZero
Возвращает `true`, если NullUUID недействителен (SQL NULL) или базовый UUID является нулевым значением.
```go
import (
    "fmt"
    "github.com/mikhaildadaev/uuid"
)
var nu uuid.NullUUID
fmt.Println(nu.IsZero())
nu.Scan("01968727-8c7e-8000-87cb-bdba4f634d9f")
fmt.Println(nu.IsZero())
```
Output
```text
true
false
```

## NULLUUID String
Возвращает каноническую текстовую форму (8-4-4-4-12) для валидного UUID или `00000000-0000-0000-0000-000000000000` для null-значения.
```go
import (
    "fmt"
    "github.com/mikhaildadaev/uuid"
)
var nu uuid.NullUUID
fmt.Println(nu.String())
nu.Scan("01968727-8c7e-8000-87cb-bdba4f634d9f")
fmt.Println(nu.String())
```
Output
```text
00000000-0000-0000-0000-000000000000
01968727-8c7e-8000-87cb-bdba4f634d9f
```

## NULLUUID Validate
Проверяет NullUUID. Возвращает `nil` как для null, так и для валидных UUID (null считается валидным в контексте SQL). Возвращает ошибку только если UUID присутствует, но имеет неверный формат.
```go
import (
    "fmt"
    "github.com/mikhaildadaev/uuid"
)
var nu uuid.NullUUID
fmt.Println(nu.Validate())
nu.Scan("01968727-8c7e-8000-87cb-bdba4f634d9f")
fmt.Println(nu.Validate())
```
Output
```text
<nil>
<nil>
```

## UUID Bytes
Возвращает UUID в виде 16-байтового среза.
```go
import (
    "fmt"
    "github.com/mikhaildadaev/uuid"
)
uuidV8String := "01968727-8c7e-8000-87cb-bdba4f634d9f"
uu, err := uuid.Parse(uuidV8String)
if err != nil {
    fmt.Println(err)
}
fmt.Printf("%x\n", uu.Bytes())
```
Output
```text
019687278c7e800087cbbdba4f634d9f
```

## UUID Equal
Сравнивает два UUID на равенство. Возвращает true, если оба UUID представляют одно и то же значение.
```go
import (
    "fmt"
    "github.com/mikhaildadaev/uuid"
)
uuidV8String := "01968727-8c7e-8000-87cb-bdba4f634d9f"
uu, err := uuid.Parse(uuidV8String)
if err != nil {
	fmt.Println(err)
}
other := uu
fmt.Println(uu.Equal(other))
```
Output
```text
true
```

## UUID Info
Возвращает удобочитаемую сводку UUID: версию, вариант, метку времени, последовательность, узел и POSIX.
```go
import (
    "fmt"
    "github.com/mikhaildadaev/uuid"
)
uuidV8String := "01968727-8c7e-8000-87cb-bdba4f634d9f"
uu, err := uuid.Parse(uuidV8String)
if err != nil {
    fmt.Println(err)
}
fmt.Println(uu.Info())
```
Output
```text
UUID: 01968727-8c7e-8000-87cb-bdba4f634d9f
VAR.: RFC4122
VER.: 8
FORM: TTTTTTTT-TTTT-8SSS-VNNN-RRRRRRRRRRRR
INFO: TIME (1-milliseconds interval since 1970-01-01) + SEQUENCE (0-4095) + NODE (0-16383) + RANDOM
TIME: 1746024238206 [2025-04-30 14:43:58.206 +0000 UTC]
SEQ.: 0
NODE: 1995
RAND: bdba4f634d9f
```

## UUID IsZero
Возвращает true, если UUID является нулевым значением (все 16 байт равны нулю).
```go
import (
    "fmt"
    "github.com/mikhaildadaev/uuid"
)
uu := uuid.NewNull()
fmt.Println(uu.IsZero())
```
Output
```text
true
```

## UUID Node
Возвращает идентификатор узла для UUID версий V1, V2, V6 и V8. Возвращает 0 для других версий.
```go
import (
    "fmt"
    "github.com/mikhaildadaev/uuid"
)
uuidNode := 1995
u8 := uuid.NewV8(uuidNode)
fmt.Println(u8.Node())
```
Output
```text
1995
```

## UUID Posix
Возвращает POSIX UID/GID для UUID версии V2. Возвращает 0 для других версий.
```go
import (
    "fmt"
    "github.com/mikhaildadaev/uuid"
)
uuidPosType := 0
uuidPosValue := 501
u2 := uuid.NewV2(uuidPosType, uuidPosValue)
fmt.Println(u2.Posix())
```
Output
```text
UID 501
```

## UUID Sequence
Возвращает поле тактовой последовательности для UUID версий V1, V2, V6 и V7. Возвращает 0 для других версий.
```go
import (
    "fmt"
    "github.com/mikhaildadaev/uuid"
)
uuidV8String := "01968727-8c7e-8000-87cb-bdba4f634d9f"
uu, err := uuid.Parse(uuidV8String)
if err != nil {
    fmt.Println(err)
}
fmt.Println(uu.Sequence())
```
Output
```text
0
```

## UUID String
Возвращает UUID в канонической текстовой форме: 8-4-4-4-12 шестнадцатеричных цифр с дефисами.
```go
import (
    "fmt"
    "github.com/mikhaildadaev/uuid"
)
uuidV8String := "01968727-8c7e-8000-87cb-bdba4f634d9f"
uu, err := uuid.Parse(uuidV8String)
if err != nil {
    fmt.Println(err)
}
fmt.Println(uu.String())
```
Output
```text
01968727-8c7e-8000-87cb-bdba4f634d9f
```

## UUID Timestamp
Возвращает метку времени, встроенную в UUID версий V1, V2, V6 и V7. Возвращает нулевое время для других версий.
```go
import (
    "fmt"
    "github.com/mikhaildadaev/uuid"
)
uuidV8String := "01968727-8c7e-8000-87cb-bdba4f634d9f"
uu, err := uuid.Parse(uuidV8String)
if err != nil {
    fmt.Println(err)
}
fmt.Println(uu.Timestamp())
```
Output
```text
1746024238206
```

## UUID Validate
Проверяет UUID и возвращает версию UUID при успехе или ошибку при неудаче.
```go
import (
    "fmt"
    "github.com/mikhaildadaev/uuid"
)
uuidV8String := "01968727-8c7e-8000-87cb-bdba4f634d9f"
uu, err := uuid.Parse(uuidV8String)
if err != nil {
    fmt.Println(err)
}
fmt.Println(uu.Validate())
```
Output
```text
<nil>
```

## UUID Variant
Возвращает номер варианта UUID. Действительные UUID возвращают 1 (стандарт RFC 4122).
```go
import (
    "fmt"
    "github.com/mikhaildadaev/uuid"
)
uuidV8String := "01968727-8c7e-8000-87cb-bdba4f634d9f"
uu, err := uuid.Parse(uuidV8String)
if err != nil {
    fmt.Println(err)
}
fmt.Println(uu.Variant())
```
Output
```text
1
```

## UUID Version
Возвращает номер версии UUID (от 1 до 8). Возвращает 0 для null UUID.
```go
import (
    "fmt"
    "github.com/mikhaildadaev/uuid"
)
uuidV8String := "01968727-8c7e-8000-87cb-bdba4f634d9f"
uu, err := uuid.Parse(uuidV8String)
if err != nil {
    fmt.Println(err)
}
fmt.Println(uu.Version())
```
Output
```text
8
```
