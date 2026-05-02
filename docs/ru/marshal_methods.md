---
outline: deep
---

# API / Сериализация / Методы

::: info **Информация**
На этой странице описана сериализация и десериализация UUID в форматы **Binary**, **Json** и **Text**. Каждый метод работает со всеми версиями UUID — от V1 до V8, включая Null.
:::

## NULLUUID MarshalBinary
Кодирует NullUUID в 16-байтовое бинарное представление. Возвращает все нули для null-значения.
```go
var nu uuid.NullUUID
data, err := nu.MarshalBinary()
if err != nil {
    fmt.Println(err)
}
fmt.Printf("%x\n", data)
uuidV8String := "01968727-8c7e-8000-87cb-bdba4f634d9f"
nu.Scan(uuidV8String)
data, err = nu.MarshalBinary()
if err != nil {
    fmt.Println(err)
}
fmt.Printf("%x\n", data)
```
Output:
```text

019687278c7e800087cbbdba4f634d9f
```

## NULLUUID MarshalJson
Кодирует NullUUID в строку Json. Возвращает null для null-значения.
```go
var nu uuid.NullUUID
data, err := nu.MarshalJson()
if err != nil {
    fmt.Println(err)
}
fmt.Println(string(data))
uuidV8String := "01968727-8c7e-8000-87cb-bdba4f634d9f"
nu.Scan(uuidV8String)
data, err := nu.MarshalJson()
if err != nil {
    fmt.Println(err)
}
fmt.Println(string(data))
```
Output:
```text
null
"01968727-8c7e-8000-87cb-bdba4f634d9f"
```

## NULLUUID MarshalText
Кодирует NullUUID в каноническую текстовую форму. Возвращает пустую строку для null-значения.
```go
var nu uuid.NullUUID
data, err := nu.MarshalText()
if err != nil {
    fmt.Println(err)
}
fmt.Println(string(data))
uuidV8String := "01968727-8c7e-8000-87cb-bdba4f634d9f"
nu.Scan(uuidV8String)
data, err := nu.MarshalText()
if err != nil {
    fmt.Println(err)
}
fmt.Println(string(data))
```
Output:
```text

01968727-8c7e-8000-87cb-bdba4f634d9f
```

## NULLUUID UnmarshalBinary
Декодирует NullUUID из 16-байтового бинарного представления. Устанавливает Valid в зависимости от входных данных.
```go
var nu uuid.NullUUID
data := []byte{}
if err := nu.UnmarshalBinary(data); err != nil {
	fmt.Println(err)
}
fmt.Println("Valid:", nu.Valid)
fmt.Println("UUID:", nu.UUID)
uuidV8Binary := []byte{0x01, 0x96, 0x87, 0x27, 0x8c, 0x7e, 0x80, 0x00, 0x87, 0xcb, 0xbd, 0xba, 0x4f, 0x63, 0x4d, 0x9f}
err := nu.UnmarshalBinary(uuidV8Binary)
if err != nil {
    fmt.Println(err)
}
fmt.Println("Valid:", nu.Valid)
fmt.Println("UUID:", nu.UUID)
```
Output:
```text
Valid: false
UUID: 00000000-0000-0000-0000-000000000000
Valid: true
UUID: 01968727-8c7e-8000-87cb-bdba4f634d9f
```

## NULLUUID UnmarshalJson
Декодирует NullUUID из Json-строки. Принимает null для SQL NULL или валидную строку UUID.
```go
var nu uuid.NullUUID
data := []byte(`null`)
if err := nu.UnmarshalJson(data); err != nil {
	fmt.Println(err)
}
fmt.Println("Valid:", nu.Valid)
fmt.Println("UUID:", nu.UUID)
uuidV8Json := []byte(`"01968727-8c7e-8000-87cb-bdba4f634d9f"`)
err := nu.UnmarshalJson(uuidV8Json)
if err != nil {
    fmt.Println(err)
}
fmt.Println("Valid:", nu.Valid)
fmt.Println("UUID:", nu.UUID)
```
Output:
```text
Valid: false
UUID: 00000000-0000-0000-0000-000000000000
Valid: true
UUID: 01968727-8c7e-8000-87cb-bdba4f634d9f
```

## NULLUUID UnmarshalText
Декодирует NullUUID из канонической текстовой формы. Принимает пустую строку для SQL NULL.
```go
var nu uuid.NullUUID
data := []byte("")
if err := nu.UnmarshalText(data); err != nil {
	fmt.Println(err)
}
fmt.Println("Valid:", nu.Valid)
fmt.Println("UUID:", nu.UUID)
uuidV8Text := []byte("01968727-8c7e-8000-87cb-bdba4f634d9f")
err := nu.UnmarshalText(uuidV8Text)
if err != nil {
    fmt.Println(err)
}
fmt.Println("Valid:", nu.Valid)
fmt.Println("UUID:", nu.UUID)
```
Output:
```text
Valid: false
UUID: 00000000-0000-0000-0000-000000000000
Valid: true
UUID: 01968727-8c7e-8000-87cb-bdba4f634d9f
```

## UUID MarshalBinary
Кодирует UUID в 16-байтовое бинарное представление.
```go
uuidV8String := "01968727-8c7e-8000-87cb-bdba4f634d9f"
uu, err := uuid.Parse(uuidV8String)
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
019687278c7e800087cbbdba4f634d9f
```

## UUID MarshalJson
Кодирует UUID в строку Json в стандартном формате 8-4-4-4-12.
```go
uuidV8String := "01968727-8c7e-8000-87cb-bdba4f634d9f"
uu, err := uuid.Parse(uuidV8String)
if err != nil {
    fmt.Println(err)
}
data, err := uu.MarshalJson()
if err != nil {
    fmt.Println(err)
}
fmt.Println(string(data))
```
Output:
```text
"01968727-8c7e-8000-87cb-bdba4f634d9f"
```

## UUID MarshalText
Кодирует UUID в каноническую текстовую форму: 32 шестнадцатеричные цифры с дефисами.
```go
uuidV8String := "01968727-8c7e-8000-87cb-bdba4f634d9f"
uu, err := uuid.Parse(uuidV8String)
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
01968727-8c7e-8000-87cb-bdba4f634d9f
```

## UUID UnmarshalBinary
Декодирует UUID из 16-байтового бинарного представления. Возвращает ошибку, если длина входных данных не равна 16 байтам.
```go
var uu uuid.UUID
uuidV8Binary := []byte{0x01, 0x96, 0x87, 0x27, 0x8c, 0x7e, 0x80, 0x00, 0x87, 0xcb, 0xbd, 0xba, 0x4f, 0x63, 0x4d, 0x9f}
err := uu.UnmarshalBinary(uuidV8Binary)
if err != nil {
    fmt.Println(err)
}
fmt.Println(uu.String())
```
Output:
```text
01968727-8c7e-8000-87cb-bdba4f634d9f
```

## UUID UnmarshalJson
Декодирует UUID из Json-строки. Принимает как строки в кавычках, так и без; возвращает ошибку при неверном формате.
```go
var uu uuid.UUID
uuidV8Json := []byte(`"01968727-8c7e-8000-87cb-bdba4f634d9f"`)
err := uu.UnmarshalJson(uuidV8Json)
if err != nil {
    fmt.Println(err)
}
fmt.Println(uu.String())
```
Output:
```text
01968727-8c7e-8000-87cb-bdba4f634d9f
```

## UUID UnmarshalText
Декодирует UUID из канонической текстовой формы (32 шестнадцатеричные цифры с дефисами). Регистронезависимый; возвращает ошибку при неверном формате.
```go
var uu uuid.UUID
uuidV8Text := []byte("01968727-8c7e-8000-87cb-bdba4f634d9f")
err := uu.UnmarshalText(uuidV8Text)
if err != nil {
    fmt.Println(err)
}
fmt.Println(uu.String())
```
Output:
```text
01968727-8c7e-8000-87cb-bdba4f634d9f
```
