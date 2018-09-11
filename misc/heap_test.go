package misc

import (
	"testing"
)

func TestHeap(t *testing.T) {
	h := NewHeap([]int{1, 2, 3, 4, 5})
	if max := h.Max(); max != 5 {
		t.Fatalf("expected max:%d got:%d\n", 5, max)
	}

}

func TestHeapPopMax(t *testing.T) {
	src := []int{1, 2, 3, 4, 5}
	h := NewHeap([]int{1, 2, 3, 4, 5})
	for i := len(src) - 1; i >= 0; i-- {
		max := h.PopMax()
		if max != src[i] {
			t.Fatalf("expected max:%d got:%d\n", src[i], max)
		}
	}
}

func TestHeapSort(t *testing.T) {

}

func TestHeapInsert(t *testing.T) {
	src := []int{1, 2, 3, 4, 5}
	h := NewHeap(src)
	h.Insert(6)
	if max := h.Max(); max != 6 {
		t.Fatalf("expected max:%d got:%d\n", 6, max)
	}
}
