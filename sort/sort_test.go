package sort

import (
	"testing"
)

func checkResult(expected []int, result []int, t *testing.T) {
	num := len(result)
	if len(expected) != num {
		t.Fatalf("expected:%v got:%v\n", expected, result)
	}
	for i := 0; i < num; i++ {
		if expected[i] != result[i] {
			t.Fatalf("expected:%v got:%v\n", expected, result)
		}
	}
}

func sortTest(t *testing.T, sortFunc func([]int) []int) {
	src := []int{10, 2, 4, 6, 8}
	dst := []int{2, 4, 6, 8, 10}
	tmp := sortFunc(src)
	checkResult(dst, tmp, t)

	src = []int{1, 2, 3, 4, 6}
	checkResult(src, sortFunc(src), t)

	src = []int{1, 2, 3, 4, 6}
	src1 := []int{6, 4, 3, 2, 1}
	checkResult(src, sortFunc(src1), t)
}

func TestSelectSort(t *testing.T) {
	sortTest(t, SelectSort)
}

func TestBubbleSort(t *testing.T) {
	sortTest(t, BubbleSort)
}
func TestInsertSort(t *testing.T) {
	sortTest(t, InsertSort)
}
