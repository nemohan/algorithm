package misc

import ()

type Heap struct {
	keys []int
	num  int
}

func NewHeap(src []int) *Heap {
	num := len(src)
	h := &Heap{
		keys: make([]int, num+1, num+1),
		num:  num,
	}
	copy(h.keys[1:], src)
	h.construct()
	return h
}

// 1, 2 , 3, 4, 5
func (heap *Heap) construct() {
	if len(heap.keys) <= 2 {
		return
	}
	keys := heap.keys
	mid := (heap.num / 2)
	for i := mid; i >= 1; i-- {
		isHeap := false
		v := keys[i]
		j := i
		k := i
		for !isHeap && j < heap.num {
			j = j * 2
			if j < heap.num {
				if keys[j] < keys[j+1] {
					j = j + 1
				}
			} else if j > heap.num {
				break
			}
			if v < keys[j] {
				keys[k] = keys[j]
			} else {
				isHeap = true
			}
			k = j
		}
		keys[k] = v
	} //end for
}

func (heap *Heap) shiftDown() {
	keys := heap.keys
	v := keys[1]
	i := 1
	k := i
	j := i
	for j < heap.num {
		j = j * 2
		if j < heap.num {
			if keys[j] < keys[j+1] {
				j = j + 1
			}
		} else if j > heap.num {
			break
		}
		if v < keys[j] {
			keys[k] = keys[j]
		} else {
			break
		}
		k = j
	}
	keys[k] = v
}

func (heap *Heap) PopMax() int {
	keys := heap.keys
	max := heap.keys[1]
	keys[1] = keys[heap.num]
	heap.num--
	heap.shiftDown()
	return max
}

func (heap *Heap) Insert(key int) {
	if len(heap.keys) == heap.num+1 {
		heap.keys = append(heap.keys, key)
	} else {
		heap.keys[heap.num+1] = key
	}
	heap.num++
	keys := heap.keys
	mid := (heap.num / 2)
	j := heap.num
	i := mid
	for ; i > 1; i = i / 2 {
		if keys[i] >= key {
			return
		}
		//down
		keys[j] = keys[i]
		j = i
	}
	keys[i] = key
}

func (heap *Heap) Max() int {
	return heap.keys[1]
}

func (heap *Heap) Min() {

}

func (heap *Heap) Sort() {

}
