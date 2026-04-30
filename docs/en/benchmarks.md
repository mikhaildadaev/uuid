---
outline: deep
---

# Benchmarks
::: info Info
The best way to compare libraries is to run benchmarks in **your own environment** with **your own workload**. Each project has unique requirements — latency, throughput, memory usage, and integration complexity — and no single test can cover them all.

I recommend that you test `ulog` alongside other libraries and choose the tool that best suits your needs.
:::

## Core Performance
Pure generation and serialization overhead. Benchmarks write to `io.Discard` — no I/O involved.

#### MultiThread
| Version | Operations | Time (ns/op) | Memory (B/op) | Allocs |
|---------|------------|--------------|---------------|--------|
| **V1**  |      34.1M |        29.29 |             0 |      0 |
| **V2**  |      59.6M |        16.79 |             0 |      0 |
| **V3**  |      84.7M |        11.81 |             0 |      0 |
| **V4**  |      18.9M |        52.88 |             0 |      0 |
| **V5**  |      51.2M |        19.51 |             0 |      0 |
| **V6**  |      32.5M |        30.75 |             0 |      0 |
| **V7**  |      22.5M |        44.45 |             0 |      0 |
| **V8**  |      32.0M |        31.21 |             0 |      0 |

#### SingleThread
| Version | Operations | Time (ns/op) | Memory (B/op) | Allocs |
|---------|------------|--------------|---------------|--------|
| **V1**  |      13.5M |        85.46 |             0 |      0 |
| **V2**  |      32.1M |        36.61 |             0 |      0 |
| **V3**  |       9.9M |       117.30 |             7 |      0 |
| **V4**  |      26.7M |        44.28 |             0 |      0 |
| **V5**  |       7.7M |       152.60 |             7 |      0 |
| **V6**  |      13.7M |        85.78 |             0 |      0 |
| **V7**  |      10.7M |       109.50 |             0 |      0 |
| **V8**  |      10.2M |       109.10 |             0 |      0 |

::: tip Note
All benchmarks measure pure generation and parsing overhead. Benchmarks marked as zero-allocation (`0 B/op`, `0 allocs/op`) perform no heap allocations in hot-path operations — minor artifacts in single-threaded V3/V5 are from the Go testing framework itself, not from UUID generation. Real-world performance depends on CPU, memory latency, and concurrency.

*Benchmarked on Intel Core i9-9880H (2.30 GHz).*
:::
