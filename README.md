[![GoDoc](https://godoc.org/github.com/go-ng/sort?status.svg)](https://pkg.go.dev/github.com/go-ng/sort?tab=doc)

# About

This package copies parts the original Go's package `sort`, but provides a generics-based implementation instead of an interface-argument based one, which results in better performance. Currently only [`sort.Slice`](https://pkg.go.dev/sort#Slice), [`sort.Sort`](https://pkg.go.dev/sort#Sort) and [`sort.Search`](https://pkg.go.dev/sort#Search) are copied&improved.

In the most of the cases this is a drop-in replacement for `sort.Slice` and `sort.Sort`, which only requires to change the import path from `sort` to `github.com/go-ng/sort`.

![latencies](https://raw.githubusercontent.com/go-ng/docs/main/sort/performance_sort.png "latencies")

#### `Slice` performance:
time:
```
name             old time/op    new time/op    delta
Slice/1-16         31.8ns ± 2%     4.6ns ± 4%   -85.60%  (p=0.000 n=9+9)
Slice/10-16         171ns ± 1%      61ns ±11%   -64.37%  (p=0.000 n=9+10)
Slice/100-16       3.43µs ± 5%    2.98µs ± 4%   -13.05%  (p=0.000 n=10+10)
Slice/1000-16      60.5µs ± 4%    57.4µs ± 2%    -5.19%  (p=0.000 n=9+9)
Slice/10000-16      809µs ± 5%     762µs ± 2%    -5.81%  (p=0.000 n=10+10)
Slice/100000-16    10.3ms ± 5%     9.6ms ± 5%    -7.44%  (p=0.000 n=10+9)
```
space:
```
name             old alloc/op   new alloc/op   delta
Slice/1-16          24.0B ± 0%      0.0B       -100.00%  (p=0.000 n=10+10)
Slice/10-16         56.0B ± 0%      0.0B       -100.00%  (p=0.000 n=10+10)
Slice/100-16        56.0B ± 0%      0.0B       -100.00%  (p=0.000 n=10+10)
Slice/1000-16       56.0B ± 0%      0.0B       -100.00%  (p=0.000 n=10+10)
Slice/10000-16      56.0B ± 0%      0.0B       -100.00%  (p=0.000 n=10+10)
Slice/100000-16     56.0B ± 0%      0.0B       -100.00%  (p=0.000 n=10+10)

name             old allocs/op  new allocs/op  delta
Slice/1-16           1.00 ± 0%      0.00       -100.00%  (p=0.000 n=10+10)
Slice/10-16          2.00 ± 0%      0.00       -100.00%  (p=0.000 n=10+10)
Slice/100-16         2.00 ± 0%      0.00       -100.00%  (p=0.000 n=10+10)
Slice/1000-16        2.00 ± 0%      0.00       -100.00%  (p=0.000 n=10+10)
Slice/10000-16       2.00 ± 0%      0.00       -100.00%  (p=0.000 n=10+10)
Slice/100000-16      2.00 ± 0%      0.00       -100.00%  (p=0.000 n=10+10)
```

#### `Sort` performance:
time:
```
name            old time/op    new time/op    delta
Sort/1-16         23.9ns ± 2%     4.3ns ± 6%   -82.15%  (p=0.000 n=8+9)
Sort/10-16         141ns ± 6%      71ns ± 9%   -49.52%  (p=0.000 n=10+10)
Sort/100-16       3.41µs ± 6%    3.01µs ± 4%   -11.81%  (p=0.000 n=9+9)
Sort/1000-16      61.8µs ± 3%    57.1µs ±10%    -7.67%  (p=0.001 n=9+9)
Sort/10000-16      833µs ± 6%     768µs ± 5%    -7.76%  (p=0.000 n=10+10)
Sort/100000-16    10.5ms ± 3%     9.3ms ± 3%   -11.50%  (p=0.000 n=10+9)
```
space:
```
name            old alloc/op   new alloc/op   delta
Sort/1-16          24.0B ± 0%      0.0B       -100.00%  (p=0.000 n=10+10)
Sort/10-16         24.0B ± 0%      0.0B       -100.00%  (p=0.000 n=10+10)
Sort/100-16        24.0B ± 0%      0.0B       -100.00%  (p=0.000 n=10+10)
Sort/1000-16       24.0B ± 0%      0.0B       -100.00%  (p=0.000 n=10+10)
Sort/10000-16      24.0B ± 0%      0.0B       -100.00%  (p=0.000 n=10+10)
Sort/100000-16     24.0B ± 0%      0.0B       -100.00%  (p=0.000 n=10+10)

name            old allocs/op  new allocs/op  delta
Sort/1-16           1.00 ± 0%      0.00       -100.00%  (p=0.000 n=10+10)
Sort/10-16          1.00 ± 0%      0.00       -100.00%  (p=0.000 n=10+10)
Sort/100-16         1.00 ± 0%      0.00       -100.00%  (p=0.000 n=10+10)
Sort/1000-16        1.00 ± 0%      0.00       -100.00%  (p=0.000 n=10+10)
Sort/10000-16       1.00 ± 0%      0.00       -100.00%  (p=0.000 n=10+10)
Sort/100000-16      1.00 ± 0%      0.00       -100.00%  (p=0.000 n=10+10)
```

#### `SearchOrdered` performance:

```
name                    old time/op    new time/op    delta
size-1/Search-16         3.90ns ± 6%    2.34ns ± 3%  -40.14%  (p=0.000 n=9+10)
size-2/Search-16         3.93ns ± 3%    2.35ns ± 3%  -40.19%  (p=0.000 n=10+10)
size-4/Search-16         7.52ns ± 3%    3.40ns ± 2%  -54.74%  (p=0.000 n=9+8)
size-8/Search-16         7.75ns ± 5%    4.02ns ± 4%  -48.07%  (p=0.000 n=10+10)
size-16/Search-16        9.64ns ± 3%    5.09ns ± 6%  -47.20%  (p=0.000 n=10+10)
size-32/Search-16        12.0ns ± 3%     5.8ns ± 5%  -51.54%  (p=0.000 n=9+9)
size-64/Search-16        16.5ns ± 6%     6.7ns ± 4%  -59.12%  (p=0.000 n=10+9)
size-128/Search-16       19.1ns ± 5%     7.4ns ± 5%  -61.11%  (p=0.000 n=9+10)
size-256/Search-16       19.1ns ± 3%     8.2ns ± 4%  -57.22%  (p=0.000 n=9+9)
size-512/Search-16       21.2ns ± 2%     9.1ns ± 8%  -57.17%  (p=0.000 n=9+9)
size-1024/Search-16      23.4ns ± 3%    12.8ns ±17%  -45.35%  (p=0.000 n=10+10)
size-2048/Search-16      26.0ns ± 3%    23.4ns ± 7%  -10.13%  (p=0.000 n=9+9)
size-4096/Search-16      28.5ns ± 9%    25.9ns ± 1%   -9.23%  (p=0.000 n=9+8)
size-8192/Search-16      30.8ns ± 8%    28.3ns ± 5%   -8.02%  (p=0.000 n=9+10)
size-16384/Search-16     32.8ns ± 3%    29.5ns ± 3%  -10.16%  (p=0.000 n=10+9)
size-32768/Search-16     36.1ns ±10%    30.2ns ± 4%  -16.46%  (p=0.000 n=10+9)
size-65536/Search-16     37.8ns ± 6%    32.1ns ± 7%  -15.17%  (p=0.000 n=9+10)
size-131072/Search-16    40.3ns ± 4%    32.4ns ± 4%  -19.70%  (p=0.000 n=10+10)
size-262144/Search-16    42.1ns ± 6%    34.0ns ± 5%  -19.23%  (p=0.000 n=10+10)
size-524288/Search-16    44.9ns ± 3%    34.9ns ± 3%  -22.40%  (p=0.000 n=10+10)
```