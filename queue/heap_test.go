package queue

import (
	"math/rand"
	"sort"
	"testing"
)

func genTestInput() []int {
	keySet := make(map[int]bool, 0)
	keys := make([]int, 0)
	for i := 0; i < 200; i++ {
		//j := rand.Int31n(500)
		j := rand.Intn(500)
		if _, ok := keySet[j]; ok {
			continue
		}
		keySet[j] = false
		keys = append(keys, j)
	}
	return keys
}
func TestHeapEnqueue(t *testing.T) {
	h := NewHeap(10)
	for i := 0; i < 10; i++ {
		h.Enqueue(elementType(i))
		m := int(h.Max())
		if m != i {
			t.Errorf("TestHeapEnqueue want:%d got:%d\n", i, m)
			break
		}
	}
}

func TestHeapPop(t *testing.T) {
	input := genTestInput()
	h := NewHeap(len(input))
	for _, v := range input {
		h.Enqueue(elementType(v))
	}
	sort.Ints(input)
	for i := len(input) - 1; i >= 0; i-- {
		m := int(h.PopMax())
		if m != input[i] {
			t.Errorf("TestHeapPop %v want:%d got:%d\n", input, input[i], m)
			break
		}
	}
}
