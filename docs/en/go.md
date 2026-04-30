---
outline: deep
---

# Go

::: info Info
The latest stable version of `uuid` is **v1.26.11**.
:::

## Get Started
```bash
go get github.com/mikhaildadaev/uuid
```

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

> **Note** 
>
> Strings `name` that are 0 bytes or larger than 512 bytes return a null UUID to prevent allocation.

## Quick Navigation
- [Benchmarks](/en/benchmarks) - Core, file, and network performance data.
- **API**
    - **Core**
        - [Main](/en/core_main) — Telemetry setup, configuration, and standard logger adapter.
    - **Marshal**
        - [Main](/en/marshal_main) — Creating a file sink and basic setup.
    - **SQL**
        - [Main](/en/sql_main) — Creating an http sink and basic setup.breaker.

