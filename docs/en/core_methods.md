---
outline: deep
---

# API / Core / Methods

::: info **Info**
This page documents all instance methods available on UUID and NullUUID — from basic operations like **String** and **Version** to advanced metadata extraction like **Timestamp**, **Node**, and **Sequence**. Every method works with all UUID versions, V1 through V8.
:::

## Bytes
Returns the UUID as a 16-byte slice.
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

## Equal
Compares two UUIDs for equality. Returns true if both UUIDs represent the same value.
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

## Info
Returns a human-readable summary of the UUID: version, variant, timestamp, sequence, node, and POSIX.
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

## IsZero
Returns true if the UUID is the zero value (all 16 bytes are zero).
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

## Node
Returns the node identifier for UUID versions V1, V2, V6, and V8. Returns 0 for other versions.
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

## Posix
Returns the POSIX UID/GID for UUID version V2. Returns 0 for other versions.
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

## Sequence
Returns the clock sequence field for UUID versions V1, V2, V6, and V7. Returns 0 for other versions.
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

## String
Returns the UUID in its canonical text form: 8-4-4-4-12 hex digits with hyphens.
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

## Timestamp
Returns the timestamp embedded in UUID versions V1, V2, V6, and V7. Returns a zero time for other versions.
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

## Validate
Validates the UUID and returns the UUID version on success, or an error on failure.
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

## Variant
Returns the UUID variant number. Valid UUIDs return 1 (RFC 4122 standard).
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

## Version
Returns the UUID version number (1 through 8). Returns 0 for a null UUID.
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
