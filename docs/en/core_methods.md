---
outline: deep
---

# API / Core / Methods

::: info **Info**
This page documents all instance methods available on UUID and NullUUID — from basic operations like **String** and **Version** to advanced metadata extraction like **Timestamp**, **Node**, and **Sequence**. Every method works with all UUID versions, V1 through V8.
:::

## NULLUUID IsZero
Returns `true` if the NullUUID is invalid (SQL NULL) or the underlying UUID is the zero value.
```go
import "github.com/mikhaildadaev/uuid"
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
Returns the canonical text form (8-4-4-4-12) for a valid UUID, or `00000000-0000-0000-0000-000000000000` for a null value.
```go
import "github.com/mikhaildadaev/uuid"
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
Validates the NullUUID. Returns `nil` for both null and valid UUIDs (null is considered valid in SQL context). Returns an error only if the UUID is present but malformed.
```go
import "github.com/mikhaildadaev/uuid"
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

## UUID Equal
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

## UUID Info
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

## UUID IsZero
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

## UUID Node
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

## UUID Posix
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

## UUID Sequence
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

## UUID String
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

## UUID Timestamp
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

## UUID Validate
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

## UUID Variant
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

## UUID Version
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
