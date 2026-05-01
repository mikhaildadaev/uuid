---
outline: deep
---

# API / 核心 / 方法

::: info **关于**
本页记录了 UUID 和 NullUUID 上可用的所有实例方法——从 **String** 和 **Version** 等基本操作，到 **Timestamp**、**Node** 和 **Sequence** 等高级元数据提取。每个方法适用于所有 UUID 版本，V1 至 V8。
:::

## Bytes
将 UUID 作为 16 字节切片返回。
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
比较两个 UUID 是否相等。如果两个 UUID 表示相同的值，则返回 true。
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
返回 UUID 的可读摘要：版本、变体、时间戳、序列、节点和 POSIX。
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
如果 UUID 为零值（所有 16 字节均为零），则返回 true。
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
返回 UUID 版本 V1、V2、V6 和 V8 的节点标识符。其他版本返回 0。
```go
import (
    "fmt"
    "github.com/mikhaildadaev/uuid"
)
node := 1995
u8 := uuid.NewV8(node)
fmt.Println(u8.Node())
```
Output
```text
1995
```

## Posix
返回 UUID 版本 V2 的 POSIX UID/GID。其他版本返回 0。
```go
import (
    "fmt"
    "github.com/mikhaildadaev/uuid"
)
posix := 501
u2 := uuid.NewV2(posix)
fmt.Println(u2.Posix())
```
Output
```text
501
```

## Sequence
返回 UUID 版本 V1、V2、V6 和 V7 的时钟序列字段。其他版本返回 0。
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
返回 UUID 的规范文本形式：8-4-4-4-12 十六进制数字，用连字符分隔。
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
返回嵌入在 UUID 版本 V1、V2、V6 和 V7 中的时间戳。其他版本返回零时间。
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
验证 UUID，成功时返回 UUID 版本，失败时返回错误。
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
返回 UUID 变体号。有效的 UUID 返回 1（RFC 4122 标准）。
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
返回 UUID 版本号（1 到 8）。对于 null UUID，返回 0。
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
