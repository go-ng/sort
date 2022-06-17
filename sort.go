// This code is a modified version of the file sort/sort.go
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

// START OF COPIED AND MODIFIED CODE //

// insertionSort_sort sorts s[a:b] using insertion sort.
func insertionSort_sort[E any, S Interface[E]](s S, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && s.Less(j, j-1); j-- {
			s[j], s[j-1] = s[j-1], s[j]
		}
	}
}

// siftDown_sort implements the heap property on s[lo:hi].
// first is an offset into the array where the root of the heap lies.
func siftDown_sort[E any, S Interface[E]](s S, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && s.Less(first+child, first+child+1) {
			child++
		}
		if !s.Less(first+root, first+child) {
			return
		}
		s[first+root], s[first+child] = s[first+child], s[first+root]
		root = child
	}
}

func heapSort_sort[E any, S Interface[E]](s S, a, b int) {
	first := a
	lo := 0
	hi := b - a

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown_sort(s, i, hi, first)
	}

	// Pop elements, largest first, into end of s.
	for i := hi - 1; i >= 0; i-- {
		s[first], s[first+i] = s[first+i], s[first]
		siftDown_sort(s, lo, i, first)
	}
}

// Quicksort, loosely following Bentley and McIlroy,
// ``Engineering a Sort Function,'' SP&E November 1993.

// medianOfThree_sort moves the median of the three values s[m0], s[m1], s[m2] into s[m1].
func medianOfThree_sort[E any, S Interface[E]](s S, m1, m0, m2 int) {
	// sort 3 elements
	if s.Less(m1, m0) {
		s[m1], s[m0] = s[m0], s[m1]
	}
	// s[m0] <= s[m1]
	if s.Less(m2, m1) {
		s[m2], s[m1] = s[m1], s[m2]
		// s[m0] <= s[m2] && s[m1] < s[m2]
		if s.Less(m1, m0) {
			s[m1], s[m0] = s[m0], s[m1]
		}
	}
	// now s[m0] <= s[m1] <= s[m2]
}

func swapRange_sort[E any, S Interface[E]](s S, a, b, n int) {
	for i := 0; i < n; i++ {
		s[a+i], s[b+i] = s[b+i], s[a+i]
	}
}

func doPivot_sort[E any, S Interface[E]](s S, lo, hi int) (midlo, midhi int) {
	m := int(uint(lo+hi) >> 1) // Written like this to avoid integer overflow.
	if hi-lo > 40 {
		// Tukey'p ``Ninther,'' median of three medians of three.
		p := (hi - lo) / 8
		medianOfThree_sort(s, lo, lo+p, lo+2*p)
		medianOfThree_sort(s, m, m-p, m+p)
		medianOfThree_sort(s, hi-1, hi-1-p, hi-1-2*p)
	}
	medianOfThree_sort(s, lo, m, hi-1)

	// Invariants are:
	//	s[lo] = pivot (set up by ChoosePivot)
	//	s[lo < i < a] < pivot
	//	s[a <= i < b] <= pivot
	//	s[b <= i < c] unexamined
	//	s[c <= i < hi-1] > pivot
	//	s[hi-1] >= pivot
	pivot := lo
	a, c := lo+1, hi-1

	for ; a < c && s.Less(a, pivot); a++ {
	}
	b := a
	for {
		for ; b < c && !s.Less(pivot, b); b++ { // s[b] <= pivot
		}
		for ; b < c && s.Less(pivot, c-1); c-- { // s[c-1] > pivot
		}
		if b >= c {
			break
		}
		// s[b] > pivot; s[c-1] <= pivot
		s[b], s[c-1] = s[c-1], s[b]
		b++
		c--
	}
	// If hi-c<3 then there are duplicates (by property of median of nine).
	// Let's be a bit more conservative, and set border to 5.
	protect := hi-c < 5
	if !protect && hi-c < (hi-lo)/4 {
		// Lets test some points for equality to pivot
		dups := 0
		if !s.Less(pivot, hi-1) { // s[hi-1] = pivot
			s[c], s[hi-1] = s[hi-1], s[c]
			c++
			dups++
		}
		if !s.Less(b-1, pivot) { // s[b-1] = pivot
			b--
			dups++
		}
		// m-lo = (hi-lo)/2 > 6
		// b-lo > (hi-lo)*3/4-1 > 8
		// ==> m < b ==> s[m] <= pivot
		if !s.Less(m, pivot) { // s[m] = pivot
			s[m], s[b-1] = s[b-1], s[m]
			b--
			dups++
		}
		// if at least 2 points are equal to pivot, assume skewed distribution
		protect = dups > 1
	}
	if protect {
		// Protect against a lot of duplicates
		// Add invariant:
		//	s[a <= i < b] unexamined
		//	s[b <= i < c] = pivot
		for {
			for ; a < b && !s.Less(b-1, pivot); b-- { // s[b] == pivot
			}
			for ; a < b && s.Less(a, pivot); a++ { // s[a] < pivot
			}
			if a >= b {
				break
			}
			// s[a] == pivot; s[b-1] < pivot
			s[a], s[b-1] = s[b-1], s[a]
			a++
			b--
		}
	}
	// Swap pivot into middle
	s[pivot], s[b-1] = s[b-1], s[pivot]
	return b - 1, c
}

func quickSort_sort[E any, S Interface[E]](s S, a, b, maxDepth int) {
	for b-a > 12 { // Use ShellSort for slices <= 12 elements
		if maxDepth == 0 {
			heapSort_sort(s, a, b)
			return
		}
		maxDepth--
		mlo, mhi := doPivot_sort(s, a, b)
		// Avoiding recursion on the larger subproblem guarantees
		// a stack depth of at most lg(b-a).
		if mlo-a < b-mhi {
			quickSort_sort(s, a, mlo, maxDepth)
			a = mhi // i.e., quickSort(s, mhi, b)
		} else {
			quickSort_sort(s, mhi, b, maxDepth)
			b = mlo // i.e., quickSort(s, a, mlo)
		}
	}
	if b-a > 1 {
		// Do ShellSort pass with gap 6
		// It could be written in this simplified form cause b-a <= 12
		for i := a + 6; i < b; i++ {
			if s.Less(i, i-6) {
				s[i], s[i-6] = s[i-6], s[i]
			}
		}
		insertionSort_sort(s, a, b)
	}
}

// Sort is a faster implementation of sort.Sort.
func Sort[E any, S Interface[E]](s S) {
	n := len(s)
	quickSort_sort(s, 0, n, maxDepth(n))
}

// END OF COPIED AND MODIFIED CODE //
