package heap

import "testing"

func TestHeap(t *testing.T) {
	var (
		values          []interface{} = []interface{}{1, 0, 2, 9, 3, 8, 4, 7, 5, 6}               // Values to add to the heap
		expPoppedValues []interface{} = []interface{}{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}               // Values expected after popping the heap to empty and after revers sorting via greaterOrEqual
		expSorted       []interface{} = []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}               // Values expected after sorting
		expTop          interface{}   = -(1<<(32<<(^uint(0)>>63)-1) - 1)                          // Value expected from peek; will update as heap changes; source: strconv.IntSize
		expSize         int           = len(values) >> 1                                          // Initialize with only part of the values; will update as the heap changes
		less            Lesser        = func(x, y interface{}) bool { return x.(int) < y.(int) }  // Less-than comparison of integers
		greaterOrEqual  Lesser        = func(x, y interface{}) bool { return y.(int) <= x.(int) } // Greater-than-or-equal comparison of integers; ironically satisfies the Lesser definition
		h               *Heap         = New(less, values[:expSize]...)
	)

	// Update expected top
	for _, value := range values[:expSize] {
		if expTop.(int) < value.(int) {
			expTop = value
		}
	}

	// Test size
	if rec := h.Size(); expSize != rec {
		t.Errorf("\nexpected %d\nreceived %d\n", expSize, rec)
		return
	}

	// Test peek, push, and size
	for _, value := range values[expSize:] {
		h.Push(value)
		expSize++
		if recSize := h.Size(); expSize != recSize {
			t.Errorf("\nexpected %d\nreceived %d\n", expSize, recSize)
			return
		}

		if expTop.(int) < value.(int) {
			expTop = value
		}

		if recTop := h.Peek(); expTop != recTop {
			t.Errorf("\nexpected %d\nreceived %d\n", expTop, recTop)
			return
		}
	}

	// Test copy
	var cpy *Heap = h.Copy() // Copy of h; will be updated
	if !h.Equals(cpy) {
		t.Errorf("\nexpected %v\nreceived %v\n", h, cpy)
		return
	}

	// Test clear
	cpy.Clear()
	if h.Equals(cpy) {
		t.Errorf("\nexpected inequality (%v)\nreceived %v\n", h, cpy)
		return
	}

	// Test values
	cpy.Push(h.Values()...)
	if !h.Equals(cpy) {
		t.Errorf("\nexpected %v\nreceived %v\n", h, cpy)
		return
	}

	// Test sorted
	var recSorted []interface{} = cpy.Sorted()
	if len(expSorted) != len(recSorted) {
		t.Errorf("\nexpected %d\nreceived %d\n", len(expSorted), len(recSorted))
		return
	}

	for i, exp := range expSorted {
		if exp != recSorted[i] {
			t.Errorf("\nexpected %d\nreceived %d\n", exp, recSorted[i])
		}
	}

	// Test set-less and sorted
	var recRevSorted []interface{} = cpy.SetLess(greaterOrEqual).Sorted()
	if len(expPoppedValues) != len(recRevSorted) {
		t.Errorf("\nexpected %d\nreceived %d\n", len(expPoppedValues), len(recRevSorted))
		return
	}

	for i, exp := range expPoppedValues {
		if exp != recRevSorted[i] {
			t.Errorf("\nexpected %d\nreceived %d\n", exp, recRevSorted[i])
			return
		}
	}

	// Test pop and size
	for _, exp := range expPoppedValues {
		if rec := h.Pop(); exp != rec {
			t.Errorf("\nexpected %d\nreceived %d\n", expPoppedValues, rec)
			return
		}

		expSize--
		if recSize := h.Size(); expSize != recSize {
			t.Errorf("\nexpected %d\nreceived %d\n", expSize, recSize)
			return
		}
	}
}

func TestNextPow2(t *testing.T) {
	expNextPow2 := map[int]int{
		-1: 0, // Undefined case: n < 0

		0: 1,
		1: 1,

		2: 2,

		3: 4,
		4: 4,

		5: 8,
		7: 8,
		8: 8,

		9:  16,
		15: 16,
		16: 16,

		17: 32,
		31: 32,
		32: 32,

		33: 64,
		63: 64,
		64: 64,

		65:  128,
		127: 128,
		128: 128,

		129: 256,
		255: 256,
		256: 256,
	}

	for n, exp := range expNextPow2 {
		if rec := nextPow2(n); exp != rec {
			t.Errorf("\nexpected %d\nreceived %d\n", exp, rec)
		}
	}
}
