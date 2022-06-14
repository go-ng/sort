package sort

import (
	"fmt"
	"math/rand"
	"runtime"
	"sort"
	"testing"
)

func BenchmarkSortGenerics(b *testing.B) {
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

			cs := make([]intSlice, slices)
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
				Sort(cs[idx])
			}
		})
	}
}

func BenchmarkStdSort(b *testing.B) {
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

			cs := make([]sort.IntSlice, slices)
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
				sort.Sort(cs[idx])
			}
		})
	}
}

type intSlice []int

func (s intSlice) Less(i, j int) bool {
	return s[i] < s[j]
}
