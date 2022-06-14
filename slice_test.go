package sort

import (
	"fmt"
	"math/rand"
	"runtime"
	"sort"
	"testing"
)

func intsEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for idx := range a {
		if a[idx] != b[idx] {
			return false
		}
	}
	return true
}

func TestSlice(t *testing.T) {
	a := make([]int, 1000)
	for idx := range a {
		a[idx] = rand.Intn(100)
	}

	cmp0 := make([]int, 1000)
	copy(cmp0, a)
	Slice(cmp0, func(i, j int) bool {
		return cmp0[i] < cmp0[j]
	})

	cmp1 := make([]int, 1000)
	copy(cmp1, a)
	sort.Slice(cmp1, func(i, j int) bool {
		return cmp1[i] < cmp1[j]
	})

	if !intsEqual(cmp1, cmp0) {
		t.Fatalf("%v != %v", cmp1, cmp0)
	}
}

func BenchmarkSliceGenerics(b *testing.B) {
	for _, size := range []uint{1, 10, 100, 1000, 10000, 100000} {
		b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
			rng := rand.New(rand.NewSource(0))

			slices := 1000/size + 30
			vs := make([][]int, slices)
			for idx := range vs {
				vs[idx] = make([]int, size)
				s := vs[idx]
				for idx := range s {
					s[idx] = rng.Intn(int(size)/2 + 1)
				}
			}

			cs := make([][]int, slices)
			for idx := range cs {
				cs[idx] = make([]int, size)
			}

			runtime.GC()
			runtime.GC()
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				idx := uint(i) % slices
				if idx == 0 {
					b.StopTimer()
					for idx := range cs {
						copy(cs[idx], vs[idx])
					}
					b.StartTimer()
				}
				s := cs[idx]
				Slice(s, func(i, j int) bool {
					return s[i] < s[j]
				})
			}
		})
	}
}

func BenchmarkStdSlice(b *testing.B) {
	for _, size := range []uint{1, 10, 100, 1000, 10000, 100000} {
		b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
			rng := rand.New(rand.NewSource(0))

			slices := 1000/size + 30
			vs := make([][]int, slices)
			for idx := range vs {
				vs[idx] = make([]int, size)
				s := vs[idx]
				for idx := range s {
					s[idx] = rng.Intn(int(size)/2 + 1)
				}
			}

			cs := make([][]int, slices)
			for idx := range cs {
				cs[idx] = make([]int, size)
			}

			runtime.GC()
			runtime.GC()
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				idx := uint(i) % slices
				if idx == 0 {
					b.StopTimer()
					for idx := range cs {
						copy(cs[idx], vs[idx])
					}
					b.StartTimer()
				}
				s := cs[idx]
				sort.Slice(s, func(i, j int) bool {
					return s[i] < s[j]
				})
			}
		})
	}
}
