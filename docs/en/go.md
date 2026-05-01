---
outline: deep
---

# Go
```bash
go get github.com/mikhaildadaev/uuid
```

::: info **Info**
The latest stable version of `uuid` is **v1.26.11**.
:::

## Run Test 
```bash
go test ./...
go test -bench=. ./...
go test -cover ./...
go test -race ./...
```

## Key Features
- **UUIDv1..v8 & Null** — Generate all UUID versions and SQL-compatible null UUID;
- **Parse & Validate** — Parse and validate standard and null UUID strings;
- **Metadata Extraction** — Extract timestamp, sequence, node, POSIX, variant, and version;
- **Serialization** — Full Binary, JSON, and Text marshal/unmarshal support;
- **SQL Integration** — Native database/sql driver with Scan/Value interfaces;
- **Zero-Allocation** — Fast, allocator-free core routines on all hot paths.

## Limits
- **V2**: `posix` should be `0..255`;
- **V3/V5**: `name` should be `1..512` symbols (byte length);
- **V8**: `node` should be `0..16383`;

::: tip **Note** 
Strings `name` that are 0 bytes or larger than 512 bytes return a null UUID to prevent allocation.
:::