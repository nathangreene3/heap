package heap

// A Heap is a binary tree together with the maximum heap property. That is,
// each parent item is greater than its children.
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

// Clean any cached values. This cannot be reversed.
func (h *Heap) Clean() *Heap {
	h.values = h.values[:h.size]
	return h
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

// Restore any cached values.
func (h *Heap) Restore() *Heap {
	return h.Push(h.values[h.size:]...)
}

// Size of the heap.
func (h *Heap) Size() int {
	return h.size
}

// Sorted pops the heap repeatedly, then returns a copy of the values sorted
// least to greatest.
func (h *Heap) Sorted() []interface{} {
	size := h.size
	for 0 < h.size {
		h.Pop()
	}

	return append(make([]interface{}, 0, size), h.values[:size]...)
}
