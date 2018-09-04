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

func sortTest(t *testing.T) {
	src := []int{10, 2, 4, 6, 8}
	dst := []int{2, 4, 6, 8, 10}
	tmp := SelectSort(src)
	checkResult(dst, tmp, t)

	src = []int{1, 2, 3, 4, 6}
	checkResult(src, SelectSort(src), t)
}

func TestSelectSort(t *testing.T) {
	sortTest(t)
}

func TestBubbleSort(t *testing.T) {
	sortTest(t)
}
