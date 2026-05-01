---
outline: deep
---

# API / Ядро / Конструкторы

::: info **Информация**
На этой странице описано создание каждой версии UUID: от совместимого с SQL `NULL` до основанной на метке времени V7 и пользовательской V8. Начните с **NewV4** для общего использования или **NewV7** для сортируемых идентификаторов.
:::

## NewNull
Создаёт UUID, совместимый с SQL-значением `NULL`.
```go
import "github.com/mikhaildadaev/uuid"
un := uuid.NewNull()
fmt.Println(un.IsZero())
fmt.Println(un.String())
```
Output:
```text
true
00000000-0000-0000-0000-000000000000
```

## NewV1
Создаёт UUID версии 1 на основе текущей метки времени и MAC-адреса локальной машины (или случайного узла, если MAC недоступен).
```go
import "github.com/mikhaildadaev/uuid"
u1 := uuid.NewV1()
fmt.Println(u1.Version())
```
Output:
```text
1
```

## NewV2
Создаёт UUID версии 2 (DCE Security) с использованием текущей метки времени, MAC-адреса локальной машины и POSIX UID/GID.
```go
import "github.com/mikhaildadaev/uuid"
u2 := uuid.NewV2(posix)
fmt.Println(u2.Version())
```
Output:
```text
2
```

## NewV3
Создаёт UUID версии 3 путём хеширования идентификатора пространства имён и имени с помощью MD5.
```go
import "github.com/mikhaildadaev/uuid"
u3 := uuid.NewV3(uuid.NameSpaceDNS, name)
fmt.Println(u3.Version())
```
Output:
```text
3
```

## NewV4
Создаёт UUID версии 4 с использованием криптографически безопасных случайных чисел.
```go
import "github.com/mikhaildadaev/uuid"
u4 := uuid.NewV4()
fmt.Println(u4.Version())
```
Output:
```text
4
```

## NewV5
Создаёт UUID версии 5 путём хеширования идентификатора пространства имён и имени с помощью SHA-1.
```go
import "github.com/mikhaildadaev/uuid"
u5 := uuid.NewV5(uuid.NameSpaceDNS, name)
fmt.Println(u5.Version())
```
Output:
```text
5
```

## NewV6
Создаёт UUID версии 6 (полевая совместимость с UUIDv1) на основе текущей метки времени и MAC-адреса локальной машины (или случайного узла, если MAC недоступен).
```go
import "github.com/mikhaildadaev/uuid"
u6 := uuid.NewV6()
fmt.Println(u6.Version())
```
Output:
```text
6
```

## NewV7
Создаёт UUID версии 7 (на основе метки времени, лексикографически сортируемый) с использованием текущей метки времени Unix в миллисекундах и случайных битов.
```go
import "github.com/mikhaildadaev/uuid"
u7 := uuid.NewV7()
fmt.Println(u7.Version())
```
Output:
```text
7
```

## NewV8
Создаёт UUID версии 8 (пользовательский, зависящий от поставщика) с использованием идентификатора узла и текущей метки времени.
```go
import "github.com/mikhaildadaev/uuid"
u8 := uuid.NewV8(node)
fmt.Println(u8.Version())
```
Output:
```text
8
```

uuid.Bytes()
uuid.Equal(other)
uuid.Info()
uuid.IsZero()
uuid.Node()
uuid.Posix()
uuid.Sequence()
uuid.String()
uuid.Timestamp()
uuid.Validate()
uuid.Variant()
uuid.Version()
nulluuid.IsZero()
nulluuid.String()
nulluuid.Validate()