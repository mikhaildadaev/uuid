---
outline: deep
---

# API / Marshal / Основное

::: info Информация
На этой странице описана сериализация и десериализация UUID в форматы **Binary**, **JSON** и **Text**. Каждый метод работает со всеми версиями UUID — от V1 до V8, включая Null.
:::

## MarshalBinary
Кодирует UUID в 16-байтовое бинарное представление.
```go
import "github.com/mikhaildadaev/uuid"
uu, err := uuid.Parse("018f3c14-8000-0000-0000-000000000001")
if err != nil {
    fmt.Println(err)
}
data, err := uu.MarshalBinary()
if err != nil {
    fmt.Println(err)
}
fmt.Printf("%x\n", data)
```
Output:
```text
018f3c1480000000000000000000000001
```

## MarshalJson
Кодирует UUID в строку JSON в стандартном формате 8-4-4-4-12.
```go
import "github.com/mikhaildadaev/uuid"
uu, err := uuid.Parse("018f3c14-8000-0000-0000-000000000001")
if err != nil {
    fmt.Println(err)
}
data, err := uu.MarshalJSON()
if err != nil {
    fmt.Println(err)
}
fmt.Println(string(data))
```
Output:
```text
018f3c14-8000-0000-0000-000000000001
```

## MarshalText
Кодирует UUID в каноническую текстовую форму: 32 шестнадцатеричные цифры с дефисами.
```go
import "github.com/mikhaildadaev/uuid"
uu, err := uuid.Parse("018f3c14-8000-0000-0000-000000000001")
if err != nil {
    fmt.Println(err)
}
data, err := uu.MarshalText()
if err != nil {
    fmt.Println(err)
}
fmt.Println(string(data))
```
Output:
```text
018f3c14-8000-0000-0000-000000000001
```

## UnmarshalBinary
Декодирует UUID из 16-байтового бинарного представления. Возвращает ошибку, если длина входных данных не равна 16 байтам.
```go
import "github.com/mikhaildadaev/uuid"
var u uuid.UUID
bin := []byte{0x01, 0x8f, 0x3c, 0x14, 0x00, 0x00, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}
err := u.UnmarshalBinary(bin)
if err != nil {
    fmt.Println(err)
}
fmt.Println(u.String())
```
Output:
```text
018f3c14-0000-8000-0000-000000000001
```

## UnmarshalJson
Декодирует UUID из JSON-строки. Принимает как строки в кавычках, так и без кавычек; возвращает ошибку при неверном формате.
```go
import "github.com/mikhaildadaev/uuid"
var u uuid.UUID
err := u.UnmarshalJSON([]byte(`"018f3c14-0000-8000-0000-000000000001"`))
if err != nil {
    fmt.Println(err)
}
fmt.Println(u.String())
```
Output:
```text
018f3c14-0000-8000-0000-000000000001
```

## UnmarshalText
Декодирует UUID из канонической текстовой формы (32 шестнадцатеричные цифры с дефисами). Регистронезависимый; возвращает ошибку при неверном формате.
```go
import "github.com/mikhaildadaev/uuid"
var u uuid.UUID
err := u.UnmarshalText([]byte("018f3c14-0000-8000-0000-000000000001"))
if err != nil {
    fmt.Println(err)
}
fmt.Println(u.String())
```
Output:
```text
018f3c14-0000-8000-0000-000000000001
```