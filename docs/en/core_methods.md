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
uu, err := uuid.Parse("018f3c14-8000-0000-0000-000000000001")
if err != nil {
    fmt.Println(err)
}
fmt.Printf("%x\n", uu.Bytes())
```
Output
```text
018f3c1480000000000000000000000001
```

## Equal
Compares two UUIDs for equality. Returns true if both UUIDs represent the same value.
```go
import (
    "fmt"
    "github.com/mikhaildadaev/uuid"
)
u1, _ := uuid.Parse("018f3c14-8000-0000-0000-000000000001")
u2, _ := uuid.Parse("018f3c14-8000-0000-0000-000000000002")
fmt.Println(u1.Equal(u2))
```
Output
```text
false
```

## Info
Returns a human-readable summary of the UUID: version, variant, timestamp, sequence, node, and POSIX.
```go
import (
    "fmt"
    "github.com/mikhaildadaev/uuid"
)
uu, err := uuid.Parse("018f3c14-8000-0000-0000-000000000001")
if err != nil {
    fmt.Println(err)
}
fmt.Println(uu.Info())
```
Output
```text
V7 time-sortable Unix-epoch 2026-04-01 00:00:00.000 +0000 UTC variant10 node1995 posix0 sequence0
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
uu, err := uuid.Parse("018f3c14-8000-0000-0000-000000000001")
if err != nil {
    fmt.Println(err)
}
fmt.Println(uu.Node())
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
uu := uuid.NewV2(1000)
fmt.Println(uu.Posix())
```
Output
```text
1000
```

## Sequence
Returns the clock sequence field for UUID versions V1, V2, V6, and V7. Returns 0 for other versions.
```go
import (
    "fmt"
    "github.com/mikhaildadaev/uuid"
)
uu, err := uuid.Parse("018f3c14-8000-0000-0000-000000000001")
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
uu, err := uuid.Parse("018f3c14-8000-0000-0000-000000000001")
if err != nil {
    fmt.Println(err)
}
fmt.Println(uu.String())
```
Output
```text
018f3c14-8000-0000-0000-000000000001
```

## Timestamp
Returns the timestamp embedded in UUID versions V1, V2, V6, and V7. Returns a zero time for other versions.
```go
import (
    "fmt"
    "github.com/mikhaildadaev/uuid"
)
uu, err := uuid.Parse("018f3c14-8000-0000-0000-000000000001")
if err != nil {
    fmt.Println(err)
}
fmt.Println(uu.Timestamp())
```
Output
```text
2026-04-01 00:00:00.000 +0000 UTC
```

## Validate
Validates the UUID and returns the UUID version on success, or an error on failure.
```go
import (
    "fmt"
    "github.com/mikhaildadaev/uuid"
)
uu, err := uuid.Parse("018f3c14-8000-0000-0000-000000000001")
if err != nil {
    fmt.Println(err)
}
fmt.Println(uu.Validate())
```
Output
```text
7 <nil>
```

## Variant
Returns the UUID variant number. Valid UUIDs return 10 (RFC 4122 standard).
```go
import (
    "fmt"
    "github.com/mikhaildadaev/uuid"
)
uu, err := uuid.Parse("018f3c14-8000-0000-0000-000000000001")
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
uu, err := uuid.Parse("018f3c14-8000-0000-0000-000000000001")
if err != nil {
    fmt.Println(err)
}
fmt.Println(uu.Version())
```
Output
```text
8
```
