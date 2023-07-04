package buffer

// RingGrowing is a growing ring buffer.
// Not thread safe.
type RingGrowing[T any] struct {
	data     []T
	n        int // Size of Data
	beg      int // First available element
	readable int // Number of data items available
}

// NewRingGrowing constructs a new RingGrowing instance with provided parameters.
func NewRingGrowing[T any](initialSize int) *RingGrowing[T] {
	return &RingGrowing[T]{
		data: make([]T, initialSize),
		n:    initialSize,
	}
}

func (r *RingGrowing[T]) ReadOne() (data T, ok bool) {
	var empty T
	if r.readable == 0 {
		return empty, false
	}
	r.readable--
	element := r.data[r.beg]
	r.data[r.beg] = empty // Remove reference to the object to help GC
	if r.beg == r.n-1 {
		// Was the last element
		r.beg = 0
	} else {
		r.beg++
	}
	return element, true
}

func (r *RingGrowing[T]) WriteOne(data T) {
	if r.readable == r.n {
		// Time to grow
		newN := r.n * 2
		newData := make([]T, newN)
		to := r.beg + r.readable
		if to <= r.n {
			copy(newData, r.data[r.beg:to])
		} else {
			copied := copy(newData, r.data[r.beg:])
			copy(newData[copied:], r.data[:(to%r.n)])
		}
		r.beg = 0
		r.data = newData
		r.n = newN
	}
	r.data[(r.readable+r.beg)%r.n] = data
	r.readable++
}
