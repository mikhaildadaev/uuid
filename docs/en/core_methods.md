---
outline: deep
---

# API / Core / Methods

::: info **Info**
The Core is the foundation of `uuid`. Here you'll learn how to create every UUID version — from the SQL-compatible `NULL` to the latest timestamp-based V7 and vendor-specific V8. Start with **NewV4** for general use or **NewV7** for sortable identifiers.
:::

## NewNull
Creates a UUID that represents a SQL-compatible `NULL` value.
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
Creates a UUID version 1 based on the current timestamp and the local machine's MAC address (or a random node if unavailable).
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
Creates a UUID version 2 (DCE Security) using the current timestamp, the local machine's MAC address, and a POSIX UID/GID.
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
Creates a UUID version 3 by hashing a namespace identifier and a name with MD5.
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
Creates a UUID version 4 using cryptographically secure random numbers.
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
Creates a UUID version 5 by hashing a namespace identifier and a name with SHA-1.
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
Creates a UUID version 6 (field-compatible with UUIDv1) based on the current timestamp and the local machine's MAC address (or a random node if unavailable).
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
Creates a UUID version 7 (timestamp-based, lexicographically sortable) using the current Unix timestamp in milliseconds and random bits.
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
Creates a UUID version 8 (vendor-specific, custom) using a node identifier and the current timestamp.
```go
import "github.com/mikhaildadaev/uuid"
u8 := uuid.NewV8(node)
fmt.Println(u8.Version())
```
Output:
```text
8
```