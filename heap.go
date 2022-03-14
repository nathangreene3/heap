package heap

import (
	"fmt"
	"math/bits"
)

// A Heap is a binary tree together with the maximum heap property.
// That is, each parent item is greater than its children.
type Heap struct {
	less   Lesser
	values []interface{}
	size   int
}

// Lesser defines the less-than comparison between two values.
type Lesser func(_, _ interface{}) bool

// A New heap.
func New(less Lesser, values ...interface{}) *Heap {
	h := Heap{
		less:   less,
		values: make([]interface{}, 0, len(values)),
	}

	return h.Push(values...)
}

// Clean frees up space that is no longer needed.
func (h *Heap) Clean() *Heap {
	if h.size < cap(h.values)>>1 {
		// 2*size < cap --> reduce cap to 2^n >= len for minimal n
		h.values = append(make([]interface{}, 0, nextPow2(h.size)), h.values[:h.size]...)
	}

	return h
}

// Clear removes all values from the heap.
func (h *Heap) Clear() *Heap {
	h.size = 0
	return h
}

// Contains determines if a value is in the heap.
func (h *Heap) Contains(value interface{}) bool {
	for i := 0; i < h.size; i++ {
		if h.values[i] == value {
			return true
		}
	}

	return false
}

// Copy returns a copy of a heap.
func (h *Heap) Copy() *Heap {
	cpy := Heap{
		less:   h.less,
		values: append(make([]interface{}, 0, nextPow2(h.size)), h.values[:h.size]...),
		size:   h.size,
	}

	return &cpy
}

// Equals determines if two heaps have equal values.
func (h *Heap) Equals(heap *Heap) bool {
	if h == heap {
		return true
	}

	if h.size != heap.size {
		return false
	}

	for i := 0; i < h.size; i++ {
		if h.values[i] != heap.values[i] {
			return false
		}
	}

	return true
}

// Peek returns the top of the heap without altering the heap.
func (h *Heap) Peek() interface{} {
	if h.size == 0 {
		return nil
	}

	return h.values[0]
}

// Pop the top of the heap.
func (h *Heap) Pop() interface{} {
	if h.size == 0 {
		return nil
	}

	// pop
	h.size--
	h.values[0], h.values[h.size] = h.values[h.size], h.values[0]

	// sift down
	for i := 0; i < h.size; {
		j := (i << 1) + 1
		if j < h.size {
			if k := j + 1; k < h.size && h.less(h.values[j], h.values[k]) {
				j = k
			}

			if h.less(h.values[i], h.values[j]) {
				h.values[i], h.values[j] = h.values[j], h.values[i]
			}
		}

		i = j
	}

	return h.values[h.size]
}

// Push values onto the heap.
func (h *Heap) Push(values ...interface{}) *Heap {
	for i := 0; i < len(values); i++ {
		// push
		if len(h.values) == h.size {
			h.values = append(h.values, values[i])
		} else {
			h.values[h.size] = values[i]
		}

		h.size++

		// sift up
		for j := h.size - 1; 0 < j; {
			k := (j - 1) >> 1
			if !h.less(h.values[k], h.values[j]) {
				break
			}

			h.values[j], h.values[k] = h.values[k], h.values[j]
			j = k
		}
	}

	return h
}

// SetLess updates the less-than comparison function. The heap will be
// updated to reflect any changes.
func (h *Heap) SetLess(less Lesser) *Heap {
	size := h.size
	for ; h.size != 0; h.Pop() {
	}

	h.less = less
	h.Push(h.values[:size]...)
	return h
}

// Size of the heap.
func (h *Heap) Size() int {
	return h.size
}

// Sorted returns a sorted copy of the values.
func (h *Heap) Sorted() []interface{} {
	values := make([]interface{}, 0, h.size)
	for ; 0 < h.size; h.Pop() {
	}

	values = append(values, h.values[:cap(values)]...)
	h.Push(values...)
	return values
}

// String returns a representation of a heap.
func (h *Heap) String() string {
	return fmt.Sprintf("{values: %v size: %d}", h.values, h.size)
}

// Values returns a copy of the values on the heap. They will not be
// sorted.
func (h *Heap) Values() []interface{} {
	return append(make([]interface{}, 0, h.size), h.values[:h.size]...)
}

// --------------------------------------------------------------------
// Helpers
// --------------------------------------------------------------------

// nextPow2 returns the next power of two greater than or equal to a
// given number.
// 	n < 0 --> 0 *Undefined, but safe to call
// 	n = 0 --> 1
// 	n > 0 --> 2^m such that 2^m >= n for  minimal m >= 0
func nextPow2(n int) int {
	if -1 < n && n < 2 {
		return 1
	}

	return 1 << (bits.Len(uint(n-1)) - 1) << 1
}
