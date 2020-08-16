package queue

type elementType int
type Heap struct {
	data []elementType
	size int
	max  int
}

func NewHeap(size int) *Heap {
	return &Heap{
		data: make([]elementType, size+1),
		max:  size + 1,
	}
}

func (h *Heap) isFull() bool {
	return h.size == h.max
}
func (h *Heap) isEmpty() bool {
	return h.size == 0
}

func (h *Heap) Enqueue(e elementType) {
	if h.isFull() {
		panic("heap is full")
	}
	h.size++
	h.data[h.size] = e
	l := h.size
	for i := h.size / 2; i >= 1; i = i / 2 {
		if e <= h.data[i] {
			break
		}
		h.data[l] = h.data[i]
		l = i
	}
	h.data[l] = e
}

func (h *Heap) PopMax() elementType {
	if h.isEmpty() {
		panic("heap is empty")
	}
	max := h.data[1]
	e := h.data[h.size]
	half := h.size / 2
	i := 1
	for i <= half {
		l := i * 2
		r := l + 1
		m := l
		if l < h.size && h.data[l] < h.data[r] {
			m = r
		}
		if h.data[m] <= e {
			break
		}
		h.data[i] = h.data[m]
		i = m
	}
	h.data[i] = e
	h.size--
	return max
}

func (h *Heap) Max() elementType {
	if h.isEmpty() {
		panic("heap is empty")
	}
	return h.data[1]
}
