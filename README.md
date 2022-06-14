[![GoDoc](https://godoc.org/github.com/go-ng/sort?status.svg)](https://pkg.go.dev/github.com/go-ng/sort?tab=doc)

# About

This package copies parts the original Go's package `sort`, but provides a generics-based implementation instead of an interface-argument based one, which results in better performance. Currently only [`sort.Slice`](https://pkg.go.dev/sort#Slice) and [`sort.Sort`](https://pkg.go.dev/sort#Sort) are copied&improved.

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
