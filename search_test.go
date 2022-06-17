package sort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func testSlice(size uint) sort.IntSlice {
	rng := rand.New(rand.NewSource(0))
	s := make(sort.IntSlice, size)
	for idx := range s {
		s[idx] = rng.Intn(int(size))
	}
	sort.Sort(s)
	return s
}

func TestSearchOrdered(t *testing.T) {
	s := testSlice(1000)
	v := 100
	idx0 := SearchOrdered(s, v)
	idx1 := sort.Search(len(s), func(i int) bool {
		return s[i] >= v
	})
	if idx0 != idx1 {
		t.Fatalf("%d != %d", idx0, idx1)
	}
}

func Benchmark(b *testing.B) {
	for size := uint(1); size < 1024*1024; size *= 2 {
		b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
			s := testSlice(size)
			b.Run("Search", func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					sort.Search(len(s), func(i int) bool {
						return s[i] >= i%int(size)
					})
				}
			})
			b.Run("SearchOrdered", func(b *testing.B) {
				b.ReportAllocs()
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					SearchOrdered(s, i%int(size))
				}
			})
		})
	}
}
