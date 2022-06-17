// This code is a modified version of the file sort/search.go
// file from Go sorce code with license header:
//
// > Copyright 2017 The Go Authors. All rights reserved.
// > Use of this source code is governed by a BSD-style
// > license that can be found in the LICENSE file.
//
// Use of this modified code is allowed with respect to this license.
// The "LICENSE file" could be found in this directory
// with name "GOLANG-LICENSE"
//
// Any rights to the modifications themselves are waived and are
// available also under CC-0 1.0 license.
//                                             -- Dmitrii Okunev

package sort

import (
	"sort"

	"golang.org/x/exp/constraints"
)

// Search is just an alias to standard sort.Search.
func Search(n int, f func(int) bool) int {
	return sort.Search(n, f)
}

// START OF COPIED AND MODIFIED CODE //

func SearchOrdered[T constraints.Ordered](s []T, v T) int {
	// Define f(-1) == false and f(n) == true.
	// Invariant: f(i-1) == false, f(j) == true.
	i, j := 0, len(s)
	for i < j {
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
		// i ≤ h < j
		if s[h] < v {
			i = h + 1 // preserves f(i-1) == false
		} else {
			j = h // preserves f(j) == true
		}
	}
	// i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.
	return i
}

// END OF COPIED AND MODIFIED CODE //
