package sort

import (
	"sort"
)

// Search is just an alias to standard sort.Search.
func Search(n int, f func(int) bool) int {
	return sort.Search(n, f)
}
